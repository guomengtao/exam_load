#!/bin/bash
echo "🔍 Running go mod tidy..."
go mod tidy

echo "🧹 Running gofmt..."
gofmt -w .

echo "🔍 Running lint..."
golangci-lint run
if [ $? -ne 0 ]; then
  echo "❌ Lint failed"
  exit 1
fi

echo "🔨 Building..."
go build ./...
if [ $? -ne 0 ]; then
  echo "❌ Build failed"
  exit 1
fi

echo "✅ All checks passed."