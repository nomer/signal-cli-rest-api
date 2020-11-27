#!/bin/sh

set -x
set -e

# Fix permissions to ensure backward compatibility
chown 1000:1000 -R /home/.local/share/signal-cli

# Show warning on docker exec
cat <<EOF >> /root/.bashrc
echo "WARNING: signal-cli-rest-api runs as signal-api (not as root!)" 
echo "Run 'su signal-api' before using signal-cli!"
EOF

# Start API as signal-api user
exec su -s /bin/sh -c "exec signal-cli-rest-api $@" signal-api
