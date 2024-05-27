#!/bin/sh

# 后台运行 ollama 程序
ollama serve &

# 等待 10 秒确保 server 启动
sleep 10

# 拉取镜像，并关闭 server
ollama pull ${MODEL} && \
ps aux | awk '/ollama serve/ && !/awk/ {print $2}' | xargs -I {} kill -9 {}