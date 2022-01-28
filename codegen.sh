#!/bin/bash

curl https://www.bitmex.com/api/explorer/swagger.json -o swagger.json

java -jar swagger-codegen-cli-2.4.25.jar generate -i swagger.json -l go -o swagger
#java -jar openapi-generator-cli.jar generate -i swagger.json -g go -o openapi
