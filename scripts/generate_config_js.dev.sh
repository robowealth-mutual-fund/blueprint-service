#!/bin/sh -eu

export API_URL=http://localhost:3002/swagger-ui

cat <<EOF
window.appConfig = {
    API_URL: "${API_URL}"
    }
EOF