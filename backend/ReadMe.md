# Master Guide Backend

Master Guide 后端服务，基于 Go + Gin + GORM + PostgreSQL + Redis 构建的技艺传承平台后端系统。

## 技术栈

- **语言**: Go 1.21
- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: PostgreSQL 15
- **缓存**: Redis 7
- **容器化**: Docker & Docker Compose
- **反向代理**: Nginx
- **文档**: Swagger

## 功能特性

- 🔐 用户认证与授权
- 👥 双重身份管理（大师/学徒）
- 📚 课程管理
- 📅 预约系统
- 💬 社群互动
- 💰 收入管理
- 📊 数据统计
- 🔄 实时通信（WebSocket）
- 📁 文件上传
- 🔍 全局搜索

## 快速开始

### 环境要求

- Docker & Docker Compose
- Go 1.21+ (仅开发环境需要)

### 启动服务

#### Windows
```bash
# 双击运行
start.bat

# 或命令行运行
docker-compose up --build -d
```

#### Linux/Mac
```bash
# 添加执行权限
chmod +x start.sh

# 运行启动脚本
./start.sh

# 或直接使用docker-compose
docker-compose up --build -d
```

### 访问服务

- **API文档**: http://localhost:8080/swagger/
- **健康检查**: http://localhost:8080/health
- **API基础URL**: http://localhost:8080/api/v1

### 停止服务

```bash
docker-compose down
```

## 项目结构

```
backend/
├── cmd/                    # 应用入口
│   └── server/            # 主服务器
├── configs/               # 配置文件
├── internal/              # 内部包
│   ├── api/              # API层
│   │   ├── handlers/     # 处理器
│   │   ├── middleware/   # 中间件
│   │   └── routes/       # 路由
│   ├── models/           # 数据模型
│   ├── repository/       # 数据访问层
│   └── services/         # 业务服务层
├── pkg/                   # 公共包
│   ├── cache/            # 缓存管理
│   ├── config/           # 配置管理
│   ├── database/         # 数据库连接
│   └── logger/           # 日志管理
├── scripts/              # 数据库脚本
├── static/               # 静态文件
├── logs/                 # 日志文件
├── docker-compose.yml    # Docker编排
├── Dockerfile           # Docker镜像
├── nginx.conf           # Nginx配置
├── go.mod               # Go模块
└── README.md            # 项目文档
```

## 配置说明

### 环境变量

主要配置在 `configs/config.yaml` 文件中：

```yaml
server:
  port: 8080
  mode: debug

database:
  host: postgres
  port: 5432
  user: master_guide
  password: master_guide123
  dbname: master_guide

redis:
  host: redis
  port: 6379
```

### 数据库

PostgreSQL 数据库会自动初始化，包含以下主要表：

- `users` - 用户基础信息
- `user_identities` - 用户身份信息
- `courses` - 课程信息
- `appointments` - 预约信息
- `circles` - 圈子信息
- `posts` - 动态信息
- 等等...

## API 接口

### 认证相关
- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录
- `POST /api/v1/auth/refresh` - 刷新Token

### 用户管理
- `GET /api/v1/users/profile` - 获取用户资料
- `PUT /api/v1/users/profile` - 更新用户资料
- `GET /api/v1/users/identities` - 获取用户身份列表

### 大师管理
- `GET /api/v1/mentors` - 获取大师列表
- `GET /api/v1/mentors/:id` - 获取大师详情
- `GET /api/v1/mentors/search` - 搜索大师

### 课程管理
- `GET /api/v1/courses` - 获取课程列表
- `GET /api/v1/courses/:id` - 获取课程详情
- `POST /api/v1/courses` - 创建课程
- `POST /api/v1/courses/:id/enroll` - 报名课程

### 预约管理
- `GET /api/v1/appointments` - 获取预约列表
- `POST /api/v1/appointments` - 创建预约
- `PUT /api/v1/appointments/:id/status` - 更新预约状态

### 社群管理
- `GET /api/v1/circles` - 获取圈子列表
- `POST /api/v1/circles/:id/join` - 加入圈子
- `DELETE /api/v1/circles/:id/join` - 退出圈子

### 动态管理
- `GET /api/v1/posts` - 获取动态列表
- `POST /api/v1/posts` - 发布动态
- `POST /api/v1/posts/:id/like` - 点赞动态

## 开发指南

### 本地开发

1. 启动数据库和Redis：
```bash
docker-compose up postgres redis -d
```

2. 运行应用：
```bash
go run cmd/server/main.go
```

### 代码规范

- 使用 Go modules 管理依赖
- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 添加适当的注释和文档

### 测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./internal/api/handlers

# 运行测试并显示覆盖率
go test -cover ./...
```

## 部署

### 生产环境

1. 修改配置文件中的生产环境设置
2. 构建镜像：
```bash
docker build -t master-guide-backend .
```

3. 使用 Docker Compose 部署：
```bash
docker-compose -f docker-compose.prod.yml up -d
```

### 环境变量

生产环境建议使用环境变量覆盖配置：

```bash
export DB_HOST=your-db-host
export DB_PASSWORD=your-db-password
export JWT_SECRET=your-jwt-secret
```

## 监控与日志

### 日志

- 应用日志：`./logs/app.log`
- 访问日志：通过 Nginx 记录
- 错误日志：通过 Gin 中间件记录

### 健康检查

- 端点：`GET /health`
- 检查数据库连接
- 检查Redis连接
- 返回服务状态

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查 PostgreSQL 容器是否正常运行
   - 验证数据库配置信息

2. **Redis连接失败**
   - 检查 Redis 容器是否正常运行
   - 验证Redis配置信息

3. **端口冲突**
   - 检查 8080、5432、6379 端口是否被占用
   - 修改 docker-compose.yml 中的端口映射

### 日志查看

```bash
# 查看应用日志
docker-compose logs app

# 查看数据库日志
docker-compose logs postgres

# 查看Redis日志
docker-compose logs redis
```

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

本项目采用 AGPL v3 许可证，详见 [LICENSE](../LICENSE) 文件。

## 联系方式

- 项目维护者：Master Guide 开发团队
- 技术支持：提交 GitHub Issue
- 功能建议：通过反馈渠道提交
