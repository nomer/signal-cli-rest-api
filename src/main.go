package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
	"os/exec"
	"time"
	"errors"
	"flag"
	"bytes"
	"os"
	"encoding/base64"
	"encoding/json"
	"strings"
)


type GroupEntry struct{
	Name string `json:"name"`
	Id string `json:"id"`
	InternalId string `json:"internal_id"`
	Members []string `json:"members"`
	Active bool `json:"active"`
	Blocked bool `json:"blocked"`
}

func cleanupTmpFiles(paths []string) {
	for _, path := range paths {
		os.Remove(path)
	}
}

func send(c *gin.Context, attachmentTmpDir string, signalCliConfig string, number string, message string, 
		recipients []string, base64Attachments []string, base64EncodedGroupId string) {
	cmd := []string{"--config", signalCliConfig, "-u", number, "send", "-m", message}
	
	if base64EncodedGroupId != "" && len(recipients) > 0 {
		c.JSON(400, gin.H{"error": "Please specify either a group id or recipient(s) - but not both"})
		return
	}

	if len(recipients) > 0 {
		cmd = append(cmd, recipients...)
	} else {
		groupId, err := base64.StdEncoding.DecodeString(base64EncodedGroupId)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid group id"})
			return
		}
		
		cmd = append(cmd, []string{"-g", string(groupId)}...)
	}
	

	attachmentTmpPaths := []string{}
	for _, base64Attachment := range base64Attachments {
		u, err := uuid.NewV4()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		dec, err := base64.StdEncoding.DecodeString(base64Attachment)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		fType, err := filetype.Get(dec)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		attachmentTmpPath := attachmentTmpDir + u.String() + "." + fType.Extension
		attachmentTmpPaths = append(attachmentTmpPaths, attachmentTmpPath)

		f, err := os.Create(attachmentTmpPath)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer f.Close()

		if _, err := f.Write(dec); err != nil {
			cleanupTmpFiles(attachmentTmpPaths)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := f.Sync(); err != nil {
			cleanupTmpFiles(attachmentTmpPaths)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		f.Close()
	}

	if len(attachmentTmpPaths) > 0 {
		cmd = append(cmd, "-a")
		cmd = append(cmd , attachmentTmpPaths...)
	}

	_, err := runSignalCli(cmd)
	if err != nil {
		cleanupTmpFiles(attachmentTmpPaths)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, nil)
}

func getGroups(number string, signalCliConfig string) ([]GroupEntry, error) {
	groupEntries := []GroupEntry{}

	out, err := runSignalCli([]string{"--config", signalCliConfig, "-u", number, "listGroups", "-d"})
	if err != nil {
		return groupEntries, err
	}

	lines := strings.Split(out, "\n")
	for _, line := range lines {
		var groupEntry GroupEntry
		if line == "" {
			continue
		}
		
		idIdx := strings.Index(line, " Name: ")
		idPair := line[:idIdx]
		groupEntry.InternalId = strings.TrimPrefix(idPair, "Id: ")
		groupEntry.Id = base64.StdEncoding.EncodeToString([]byte(groupEntry.InternalId))
		lineWithoutId := strings.TrimLeft(line[idIdx:], " ")

		nameIdx := strings.Index(lineWithoutId, " Active: ")
		namePair := lineWithoutId[:nameIdx]
		groupEntry.Name = strings.TrimRight(strings.TrimPrefix(namePair, "Name: "), " ")
		lineWithoutName := strings.TrimLeft(lineWithoutId[nameIdx:], " ")

		activeIdx := strings.Index(lineWithoutName, " Blocked: ")
		activePair := lineWithoutName[:activeIdx]
		active := strings.TrimPrefix(activePair, "Active: ")
		if active == "true" {
			groupEntry.Active = true
		} else {
			groupEntry.Active = false
		}
		lineWithoutActive := strings.TrimLeft(lineWithoutName[activeIdx:], " ")

		blockedIdx := strings.Index(lineWithoutActive, " Members: ")
		blockedPair := lineWithoutActive[:blockedIdx]
		blocked := strings.TrimPrefix(blockedPair, "Blocked: ")
		if blocked == "true" {
			groupEntry.Blocked = true
		} else {
			groupEntry.Blocked = false
		}
		lineWithoutBlocked := strings.TrimLeft(lineWithoutActive[blockedIdx:], " ")

		membersPair := lineWithoutBlocked
		members := strings.TrimPrefix(membersPair, "Members: ")
		trimmedMembers := strings.TrimRight(strings.TrimLeft(members, "["), "]")
		trimmedMembersList := strings.Split(trimmedMembers, ",")
		for _, member := range trimmedMembersList {
			groupEntry.Members = append(groupEntry.Members, strings.Trim(member, " "))
		}

		groupEntries = append(groupEntries, groupEntry)
	}

	return groupEntries, nil
}

func runSignalCli(args []string) (string, error) {
	cmd := exec.Command("signal-cli", args...)
	var errBuffer bytes.Buffer
	var outBuffer bytes.Buffer
	cmd.Stderr = &errBuffer
	cmd.Stdout = &outBuffer

	err := cmd.Start()
	if err != nil {
		return "", err
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(60 * time.Second):
		err := cmd.Process.Kill()
		if err != nil {
			return "", err
		}
		return "", errors.New("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			return "", errors.New(errBuffer.String())
		}
	}
	return outBuffer.String(), nil
}

func main() {
	signalCliConfig := flag.String("signal-cli-config", "/home/.local/share/signal-cli/", "Config directory where signal-cli config is stored")
	attachmentTmpDir := flag.String("attachment-tmp-dir", "/tmp/", "Attachment tmp directory")
	flag.Parse()

	router := gin.Default()
	log.Info("Started Signal Messenger REST API")

	router.GET("/v1/about", func(c *gin.Context) {
		type About struct {
			SupportedApiVersions []string `json:"versions"`
		}

		about := About{SupportedApiVersions: []string{"v1", "v2"}}
		c.JSON(200, about)
	})

	router.POST("/v1/register/:number", func(c *gin.Context) {
		number := c.Param("number")

		type Request struct{
			UseVoice bool `json:"use_voice"`
		}

		var req Request

		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		if buf.String() != "" {
			err := json.Unmarshal(buf.Bytes(), &req)
			if err != nil {
				log.Error("Couldn't register number: ", err.Error())
				c.JSON(400, "Couldn't process request - invalid request.")
				return
			}
		} else {
			req.UseVoice = false
		}

		if number == "" {
			c.JSON(400, "Please provide a number")
			return
		}

		command := []string{"--config", *signalCliConfig, "-u", number, "register"}

		if req.UseVoice == true {
			command = append(command, "--voice")
		}

		_, err := runSignalCli(command)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(201, nil)
	})

	router.POST("/v1/register/:number/verify/:token", func(c *gin.Context) {
		number := c.Param("number")
		token := c.Param("token")

		if number == "" {
			c.JSON(400, "Please provide a number")
			return
		}

		if token == "" {
			c.JSON(400, gin.H{"error": "Please provide a verification code"})
			return
		}

		
		_, err := runSignalCli([]string{"--config", *signalCliConfig, "-u", number, "verify", token})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, nil)
	})

	router.POST("/v1/send", func(c *gin.Context) {
		type Request struct{
			Number string `json:"number"`
			Recipients []string `json:"recipients"`
			Message string `json:"message"`
			Base64Attachment string `json:"base64_attachment"`
			GroupId string `json:"group_id"`
		}
		var req Request
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(400, "Couldn't process request - invalid request")
			return
		}

		base64Attachments := []string{}
		if req.Base64Attachment != "" {
			base64Attachments = append(base64Attachments, req.Base64Attachment)
		}

		send(c, *signalCliConfig, *signalCliConfig, req.Number, req.Message, req.Recipients, base64Attachments, req.GroupId)
	})

	router.POST("/v2/send", func(c *gin.Context) {
		type Request struct{
			Number string `json:"number"`
			Recipients []string `json:"recipients"`
			Message string `json:"message"`
			Base64Attachments []string `json:"base64_attachments"`
			GroupId string `json:"group_id"`
		}
		var req Request
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(400, "Couldn't process request - invalid request")
			log.Error(err.Error())
			return
		}

		send(c, *attachmentTmpDir, *signalCliConfig, req.Number, req.Message, req.Recipients, req.Base64Attachments, req.GroupId)
	})

	router.GET("/v1/receive/:number", func(c *gin.Context) {
		number := c.Param("number")

		command := []string{"--config", *signalCliConfig, "-u", number, "receive", "-t", "1", "--json"}
		out, err := runSignalCli(command)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		
		out = strings.Trim(out, "\n")
		lines := strings.Split(out, "\n")
		
		jsonStr := "["
		for i, line := range lines {
			jsonStr += line
			if i != (len(lines) - 1) {
				jsonStr += ","
			}
		}
		jsonStr += "]"

		c.String(200, jsonStr)
	})

	router.POST("/v1/groups/:number", func(c *gin.Context) {
		number := c.Param("number")
		
		type Request struct{
			Name string `json:"name"`
			Members []string `json:"members"`
		}

		var req Request
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(400, "Couldn't process request - invalid request")
			log.Error(err.Error())
			return
		}


		cmd := []string{"--config", *signalCliConfig, "-u", number, "updateGroup", "-n", req.Name, "-m"}
		cmd = append(cmd, req.Members...)
		log.Info(cmd)
		out, err := runSignalCli(cmd)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		log.Info(out)

	})

	router.GET("/v1/groups/:number", func(c *gin.Context) { 
		number := c.Param("number")
		
		groups, err := getGroups(number, *signalCliConfig)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, groups)
	})


	router.DELETE("/v1/groups/:number/:groupid", func(c *gin.Context) {
		base64EncodedGroupId := c.Param("groupid")
		number := c.Param("number")

		if base64EncodedGroupId == "" {
			c.JSON(400, gin.H{"error": "Please specify a group id"})
			return
		}

		groupId, err := base64.StdEncoding.DecodeString(base64EncodedGroupId)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid group id"})
			return
		}

		_, err = runSignalCli([]string{"--config", *signalCliConfig, "-u", number, "quitGroup", "-g", string(groupId)})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	})

	router.Run()
}
