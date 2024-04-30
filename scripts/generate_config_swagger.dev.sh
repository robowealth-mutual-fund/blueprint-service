#!/bin/sh
sed -ie "s|{BASE_PATH}|$BASE_PATH|" ./www/swagger-ui/swagger.json
