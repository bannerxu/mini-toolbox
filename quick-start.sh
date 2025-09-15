#!/bin/bash

# 简化启动脚本 - 快速启动前后端服务
# 使用方法: ./quick-start.sh

FRONTEND_DIR="/Users/xgl/GolandProjects/mini-toolbox/frontend"
BACKEND_DIR="/Users/xgl/GolandProjects/mini-toolbox/backend"

echo "🚀 启动 Mini Toolbox..."

# 清理可能占用的端口
echo "🔧 清理端口..."
pkill -f "go run main.go" 2>/dev/null
pkill -f "vite" 2>/dev/null
sleep 1

# 启动后端
echo "⚡ 启动后端 (端口 8080)..."
cd "$BACKEND_DIR"
go run main.go &
BACKEND_PID=$!

# 启动前端
echo "⚡ 启动前端 (端口 3000)..."
cd "$FRONTEND_DIR"
npm run dev &
FRONTEND_PID=$!

# 等待启动
echo "⏳ 等待服务启动..."
sleep 5

echo ""
echo "✅ 服务启动完成！"
echo "📱 前端: http://localhost:3000"
echo "🔧 后端: http://localhost:8080"
echo "🔗 API: http://localhost:8080/api/v1"
echo ""
echo "💡 按 Ctrl+C 停止所有服务"

# 清理函数
cleanup() {
    echo ""
    echo "🛑 正在停止服务..."
    kill $BACKEND_PID 2>/dev/null
    kill $FRONTEND_PID 2>/dev/null
    pkill -f "go run main.go" 2>/dev/null
    pkill -f "vite" 2>/dev/null
    echo "✅ 服务已停止"
    exit 0
}

# 捕获信号
trap cleanup SIGINT SIGTERM

# 保持运行
wait