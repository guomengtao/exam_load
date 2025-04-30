#!/bin/bash

# 定义宿主机路径和容器路径
HOST_PATH="/www/wwwroot/gin-go-test/static/vue-project/my-vue-app"
CONTAINER_PATH="/app"

# 容器名称
CONTAINER_NAME="vue-prod-container"

# Docker 镜像名称
IMAGE_NAME="vue-app-prod"

# 构建 Docker 镜像
echo "构建 Docker 镜像: $IMAGE_NAME"
docker build -t $IMAGE_NAME .

# 如果已经有容器存在，先删除它
echo "检查容器是否已存在并删除..."
docker rm -f $CONTAINER_NAME

# 运行容器并挂载宿主机目录到容器内的工作目录
echo "启动容器: $CONTAINER_NAME"
docker run -d \
  --restart always \
  -p 8089:8080 \
  --name $CONTAINER_NAME \
  -v $HOST_PATH:$CONTAINER_PATH \
  -w $CONTAINER_PATH \
  $IMAGE_NAME npm run serve

# 提示用户容器正在启动
echo "容器正在启动，访问 http://localhost:8089/ 查看应用"