# Master Guide - 大师指导平台

## 项目概述

Master Guide 是一个专为传统技艺传承和现代技能培养设计的移动端应用平台。该平台连接各领域大师与学习者，提供专业技艺培养计划、教学课程和实时线上指导服务。

### 核心理念
- **技艺传承**: 为传统技艺和现代技能提供专业的传承平台
- **师徒制**: 建立现代师徒关系，实现一对一专业指导
- **双重身份**: 支持用户在不同领域拥有大师和学徒双重身份
- **标准化**: 定义技艺里程碑和评估标准，为学习者提供清晰路径
- **社区评价**: 通过学徒和行业评价体系，确保大师资质

## 项目结构

```
master/
├── docs/                     # 项目文档
│   ├── PRD.md               # 产品需求文档
│   ├── architecture/        # 架构设计文档
│   │   ├── system-architecture-summary.md
│   │   ├── frontend-architecture.md
│   │   ├── backend-architecture.md
│   │   └── api-design.md
│   ├── requirements/        # 需求文档
│   │   └── user-stories.md  # 用户故事文档
│   ├── ui-design/          # UI设计文档
│   │   ├── ui-design-guidelines.md
│   │   ├── interaction-design.md
│   │   ├── component-library.md
│   │   ├── page-layout-design.md
│   │   ├── ui-prompt-generation.md
│   │   └── ui-prompt-templates.md
│   └── database/           # 数据库设计
│       ├── migrations/     # 数据库迁移文件
│       └── seeds/         # 数据库种子数据
├── frontend/               # 前端项目 (Vue 3 + TypeScript)
│   ├── src/
│   │   ├── assets/        # 静态资源
│   │   ├── components/    # 组件
│   │   ├── views/         # 页面
│   │   ├── router/        # 路由
│   │   ├── stores/        # 状态管理
│   │   ├── services/      # API服务
│   │   └── utils/         # 工具函数
│   ├── package.json
│   ├── vite.config.ts
│   └── README.md
├── backend/                # 后端项目 (Go + Gin)
│   ├── cmd/               # 应用入口
│   ├── internal/          # 内部包
│   ├── pkg/               # 公共包
│   └── README.md
├── web-bundles/           # Web打包配置
├── y/                     # 项目配置
├── .bmad-core/           # BMAD核心配置
├── .gitignore            # Git忽略文件
└── README.md             # 项目总览
```

## 技术架构

### 前端技术栈
- **Vue 3**: 渐进式JavaScript框架
- **TypeScript**: 类型安全的JavaScript超集
- **Vite**: 下一代前端构建工具
- **Vue Router 4**: 路由管理
- **Pinia**: 状态管理
- **Element Plus**: UI组件库
- **Socket.io-client**: 实时通信
- **Axios**: HTTP客户端

### 后端技术栈
- **Go 1.21**: 高性能编程语言
- **Gin**: 轻量级Web框架
- **PostgreSQL**: 关系型数据库
- **Redis**: 缓存和会话存储
- **GORM**: ORM框架
- **JWT**: 身份认证
- **WebSocket**: 实时通信

### 开发工具
- **Docker**: 容器化部署
- **GitHub Actions**: CI/CD流水线
- **Prometheus**: 性能监控
- **ELK Stack**: 日志管理

## 快速开始

### 环境要求
- **Node.js**: 18.0+
- **Go**: 1.21+
- **PostgreSQL**: 15.0+
- **Redis**: 7.0+
- **Docker**: 20.0+

### 开发环境搭建

#### 1. 克隆项目
```bash
git clone https://github.com/your-username/master-guide.git
cd master-guide
```

#### 2. 启动前端服务
```bash
cd frontend
npm install
npm run dev
```

#### 3. 启动后端服务
```bash
cd backend
go mod download
go run cmd/server/main.go
```

### 文档导航

#### 产品文档
- [产品需求文档](./docs/PRD.md) - 详细的产品功能需求
- [用户故事文档](./docs/requirements/user-stories.md) - 用户故事和验收标准

#### 技术文档
- [系统架构总览](./docs/architecture/system-architecture-summary.md) - 整体技术架构
- [前端架构设计](./docs/architecture/frontend-architecture.md) - Vue 3前端架构
- [后端架构设计](./docs/architecture/backend-architecture.md) - Go后端架构
- [API设计文档](./docs/architecture/api-design.md) - 接口设计规范

#### 设计文档
- [UI设计规范](./docs/ui-design/ui-design-guidelines.md) - 设计系统和组件规范
- [交互设计](./docs/ui-design/interaction-design.md) - 用户交互流程
- [组件库](./docs/ui-design/component-library.md) - UI组件库
- [页面布局设计](./docs/ui-design/page-layout-design.md) - 页面布局规范

## 功能特性

### 核心功能模块

#### 1. 身份管理系统
- **双重身份支持**: 用户可同时拥有大师和学徒身份
- **身份切换**: 在应用内可随时切换不同身份
- **权限管理**: 不同身份拥有不同的功能和权限

#### 2. 大师指导系统
- **大师搜索**: 按技能、经验等条件搜索
- **大师档案**: 详细的个人介绍、专业领域、评分
- **预约系统**: 直接预约指导课程
- **实时通信**: WebSocket支持实时指导

#### 3. 课程学习系统
- **精选课程**: 大师设计的专业课程
- **分类筛选**: 按领域分类（开发、设计、营销等）
- **学习进度**: 跟踪学习进度和完成状态
- **评价系统**: 对课程和大师进行评价

#### 4. 社群互动系统
- **领域圈子**: 按专业领域划分的交流社群
- **动态分享**: 学习心得、作品展示
- **实时互动**: 点赞、评论、私信功能
- **问答系统**: 专业问题讨论

## 开发规范

### 代码规范
- 使用TypeScript进行类型检查
- 遵循ESLint和Prettier代码规范
- 使用Composition API编写Vue组件
- 采用SCSS进行样式开发

### 提交规范
- 使用Conventional Commits规范
- 提交前进行代码检查和测试
- 重要功能需要代码审查

### 文档规范
- 所有API接口必须有文档
- 组件必须有使用说明
- 重要决策需要记录在文档中

## 部署指南

### 开发环境
```bash
# 前端开发服务器
cd frontend && npm run dev

# 后端开发服务器
cd backend && go run cmd/server/main.go
```

### 生产环境
```bash
# Docker部署
docker-compose up -d

# 查看服务状态
docker-compose ps
```

## 贡献指南

1. Fork项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开Pull Request

## 许可证

本项目采用 AGPL v3 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 联系方式

- 项目维护者: [Jack Zhu]
- 邮箱: [qqjack2012@gmail.com]
- 项目链接: [https://github.com/lighthought/master](https://github.com/lighthought/master)

---

**Master Guide** - 让技艺传承更简单，让学习更高效！