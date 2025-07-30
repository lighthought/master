# Master Guide 项目文档索引

## 📚 文档导航

### 🎯 产品文档
| 文档 | 描述 | 状态 |
|------|------|------|
| [产品需求文档](./PRD.md) | 详细的产品功能需求和技术要求 | ✅ 完成 |
| [用户故事文档](./requirements/user-stories.md) | 用户故事和验收标准 | ✅ 完成 |

### 🏗️ 架构设计
| 文档 | 描述 | 状态 |
|------|------|------|
| [系统架构总览](./architecture/system-architecture-summary.md) | 整体技术架构设计 | ✅ 完成 |
| [前端架构设计](./architecture/frontend-architecture.md) | Vue 3前端技术架构 | ✅ 完成 |
| [后端架构设计](./architecture/backend-architecture.md) | Go后端技术架构 | ✅ 完成 |
| [API设计文档](./architecture/api-design.md) | 接口设计规范 | ✅ 完成 |

### 🎨 设计文档
| 文档 | 描述 | 状态 |
|------|------|------|
| [UI设计规范](./ui-design/ui-design-guidelines.md) | 设计系统和组件规范 | ✅ 完成 |
| [交互设计](./ui-design/interaction-design.md) | 用户交互流程设计 | ✅ 完成 |
| [组件库](./ui-design/component-library.md) | UI组件库设计 | ✅ 完成 |
| [页面布局设计](./ui-design/page-layout-design.md) | 页面布局规范 | ✅ 完成 |
| [UI提示词生成](./ui-design/ui-prompt-generation.md) | UI设计提示词指南 | ✅ 完成 |
| [UI提示词模板](./ui-design/ui-prompt-templates.md) | UI设计模板 | ✅ 完成 |

### 🗄️ 数据库设计
| 文档 | 描述 | 状态 |
|------|------|------|
| [数据库迁移](./database/migrations/) | 数据库结构变更 | 🔄 进行中 |
| [种子数据](./database/seeds/) | 初始测试数据 | 🔄 进行中 |

## 🚀 快速开始

### 开发环境
```bash
# 克隆项目
git clone https://github.com/your-username/master-guide.git
cd master-guide

# 启动前端
cd frontend
npm install
npm run dev

# 启动后端
cd backend
go mod download
go run cmd/server/main.go
```

### 文档阅读顺序
1. [产品需求文档](./PRD.md) - 了解产品整体需求
2. [系统架构总览](./architecture/system-architecture-summary.md) - 了解技术架构
3. [用户故事文档](./requirements/user-stories.md) - 了解具体功能需求
4. [前端/后端架构设计](./architecture/) - 了解技术实现细节
5. [UI设计文档](./ui-design/) - 了解界面设计规范

## 📋 项目状态

### ✅ 已完成
- [x] 产品需求文档 (PRD)
- [x] 用户故事文档
- [x] 系统架构设计
- [x] 前端架构设计
- [x] 后端架构设计
- [x] UI设计规范
- [x] 前端项目基础结构
- [x] 项目文档索引

### 🔄 进行中
- [ ] 数据库迁移文件
- [ ] 后端项目实现
- [ ] 前端组件开发
- [ ] API接口实现

### 📅 计划中
- [ ] 单元测试
- [ ] 集成测试
- [ ] 部署配置
- [ ] 监控配置

## 🛠️ 技术栈

### 前端
- **Vue 3** + **TypeScript** + **Vite**
- **Element Plus** UI组件库
- **Pinia** 状态管理
- **Vue Router 4** 路由管理

### 后端
- **Go 1.21** + **Gin** 框架
- **PostgreSQL** 数据库
- **Redis** 缓存
- **GORM** ORM框架

### 开发工具
- **Docker** 容器化
- **GitHub Actions** CI/CD
- **ESLint** + **Prettier** 代码规范

## 📞 联系方式

- 项目维护者: [Your Name]
- 邮箱: [your.email@example.com]
- 项目链接: [https://github.com/your-username/master-guide](https://github.com/your-username/master-guide)

---

**最后更新**: 2024年12月
**文档版本**: v1.0.0