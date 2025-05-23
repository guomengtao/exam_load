#!/bin/bash

if [ "$1" == "-b" ]; then
  echo "📦 仅构建 hello 可执行文件，不启动"
  go build -o hello main.go
else
  if pgrep hello > /dev/null; then
    echo "🛑 检测到已有 hello 实例，正在关闭..."
    pkill hello
  fi

  echo "🚀 正在构建并启动 hello..."
  go build -o hello main.go && ./hello &
fi