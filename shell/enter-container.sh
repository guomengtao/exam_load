#!/bin/bash

# 定义你的容器名称
CONTAINER_NAME="vue-prod-container"

# 检查容器是否存在并正在运行
if docker ps -a --format '{{.Names}}' | grep -q "^$CONTAINER_NAME$"; then
  echo "正在进入容器：$CONTAINER_NAME"
  docker exec -it "$CONTAINER_NAME" /bin/bash
else
  echo "容器 $CONTAINER_NAME 不存在或未运行"
fi