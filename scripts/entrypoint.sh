#!/bin/sh

/scripts/download_swagger.sh
/scripts/generate_config_js.sh > /www/swagger-ui/env.js
/scripts/generate_config_swagger.sh

echo "Start Server:"
ls -lart /app/bin
ls -lart /app/bin/server
./app/bin/server
