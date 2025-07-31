@echo off
echo Starting Master Guide Backend for local development...

REM 设置环境变量
set CONFIG_PATH=./configs/config.local.yaml

REM 启动数据库和Redis（如果Docker可用）
docker --version >nul 2>&1
if not errorlevel 1 (
    echo Starting PostgreSQL and Redis with Docker...
    docker-compose up postgres redis -d
    timeout /t 5 /nobreak >nul
) else (
    echo Docker not available, please start PostgreSQL and Redis manually
    echo PostgreSQL should be running on localhost:5432
    echo Redis should be running on localhost:6379
)

REM 运行应用
echo Starting application...
go run cmd/server/main.go

pause 