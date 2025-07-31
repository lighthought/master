#!/bin/bash

echo "Starting Master Guide Backend..."

# 检查Docker是否运行
if ! command -v docker &> /dev/null; then
    echo "Error: Docker is not installed"
    exit 1
fi

if ! docker info &> /dev/null; then
    echo "Error: Docker is not running"
    exit 1
fi

# 构建并启动服务
echo "Building and starting services..."
docker-compose up --build -d

echo ""
echo "Services started successfully!"
echo ""
echo "API Documentation: http://localhost:8080/swagger/"
echo "Health Check: http://localhost:8080/health"
echo ""
echo "Press Ctrl+C to stop services..."

# 等待中断信号
trap 'echo ""; echo "Stopping services..."; docker-compose down; echo "Services stopped."; exit 0' INT

# 保持脚本运行
while true; do
    sleep 1
done 