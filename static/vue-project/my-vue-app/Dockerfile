FROM node:16

# 设置工作目录
WORKDIR /app

# 复制项目文件（注意是宿主机的 my-vue-app 内容）
COPY . .

# 安装依赖
RUN npm install

# 暴露端口
EXPOSE 8080

# 开发环境运行方式
CMD ["npm", "run", "serve"]