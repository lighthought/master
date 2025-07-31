 @echo off
chcp 65001 >nul

echo 🚀 MasterGuide Backend 启动脚本 (Windows)
echo =========================================

REM 检查Go是否安装
go version >nul 2>&1
if errorlevel 1 (
    echo ❌ Go未安装，请先安装Go 1.21+
    pause
    exit /b 1
)

echo ✅ Go环境检查通过

REM 下载依赖
echo 📦 下载Go依赖...
go mod download
go mod tidy

REM 检查swag是否安装
swag --version >nul 2>&1
if errorlevel 1 (
    echo 📚 安装swag工具...
    go install github.com/swaggo/swag/cmd/swag@latest
)

REM 生成API文档
echo 📚 生成API文档...
swag init -g cmd/server/main.go -o docs

REM 检查Docker是否安装
docker --version >nul 2>&1
if errorlevel 1 (
    echo ⚠️ Docker未安装，将使用本地模式运行
    goto :local_run
)

REM 检查Docker Compose是否安装
docker-compose --version >nul 2>&1
if errorlevel 1 (
    echo ⚠️ Docker Compose未安装，将使用本地模式运行
    goto :local_run
)

echo ✅ Docker环境检查通过

REM 启动数据库和Redis
echo 🗄️ 启动PostgreSQL和Redis...
docker-compose up -d postgres redis

REM 等待数据库启动
echo ⏳ 等待数据库启动...
timeout /t 10 /nobreak >nul

REM 初始化数据库
echo 🔧 初始化数据库...
if exist "scripts\db_init.sql" (
    echo 请手动执行数据库初始化脚本: scripts\db_init.sql
) else (
    echo ⚠️ 数据库初始化脚本不存在
)

goto :choose_mode

:local_run
echo ⚠️ 本地模式：请确保PostgreSQL和Redis服务已启动

:choose_mode
REM 启动应用
echo 🚀 启动应用...
echo 选择启动方式：
echo 1. 本地运行 (go run)
echo 2. 构建后运行 (go build)
echo 3. 开发模式 (需要安装air)
set /p choice=请选择 (1-3): 

if "%choice%"=="1" (
    echo 🏃 本地运行...
    go run cmd/server/main.go
) else if "%choice%"=="2" (
    echo 🔨 构建应用...
    go build -o bin/server.exe cmd/server/main.go
    if errorlevel 1 (
        echo ❌ 构建失败
        pause
        exit /b 1
    )
    echo 🚀 启动应用...
    bin\server.exe
) else if "%choice%"=="3" (
    echo 🔧 检查开发工具...
    air --version >nul 2>&1
    if errorlevel 1 (
        echo 📦 安装air开发工具...
        go install github.com/cosmtrek/air@latest
    )
    echo 🔧 开发模式运行...
    air
) else (
    echo ❌ 无效选择，退出
    pause
    exit /b 1
)

pause 