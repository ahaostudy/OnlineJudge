#!/bin/bash

# 检查是否提供了参数
if [ "$#" -eq 0 ]; then
    echo "请提供至少一个端口号作为参数。"
    exit 1
fi

# 循环遍历所有传入的端口号
for port in "$@"; do
    # 使用lsof命令找到给定端口的进程ID
    pid=$(lsof -ti tcp:$port)

    # 检查进程ID是否存在
    if [ -z "$pid" ]; then
        echo "没有找到运行在端口 $port 的进程。"
    else
        # 终止进程
        echo "正在关闭运行在端口 $port 的进程（PID：$pid）..."
        kill -9 $pid
    fi
done
