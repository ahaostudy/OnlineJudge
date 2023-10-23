#!/bin/bash

services=(
    "cmd/gateway/main.go"
    "cmd/judge/main.go"
    "cmd/problem/main.go"
    "cmd/submit/main.go"
    "cmd/user/main.go"
    "cmd/chatgpt/main.go"
)

mkdir -p log

for service in "${services[@]}"; do

    echo "Starting $service..."

    log_name=$(basename $(dirname $service))
    >"log/${log_name}.log"
    go run $service >>"log/${log_name}.log" 2>&1 &

done

echo "All services started."
