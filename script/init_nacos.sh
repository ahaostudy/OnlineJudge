#!/bin/bash

NACOS_HOST="nacos"
NACOS_PORT=8848
NAMESPACE="onlinejudge-default"

if [ -n "$1" ]; then
    CONFIG_PATH="$1"
else
    CONFIG_PATH="config/nacos_default_config.zip"
fi

curl -X POST "http://$NACOS_HOST:$NACOS_PORT/nacos/v1/console/namespaces" \
    -d "customNamespaceId=$NAMESPACE&namespaceName=$NAMESPACE&namespaceDesc=onlinejudge default config"

curl --location \
    --request POST "http://$NACOS_HOST:$NACOS_PORT/nacos/v1/cs/configs?import=true&namespace=$NAMESPACE" \
    --form "policy=OVERWRITE" \
    --form "file=@$CONFIG_PATH"

