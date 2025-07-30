# Master Guide 前端项目

## 项目简介

Master Guide 是一个技艺传承平台的前端项目，基于 Vue 3 + TypeScript + Element Plus 构建。

## 技术栈

- **Vue 3** - 渐进式JavaScript框架
- **TypeScript** - 类型安全的JavaScript超集
- **Vite** - 下一代前端构建工具
- **Vue Router 4** - 路由管理
- **Pinia** - 状态管理
- **Element Plus** - UI组件库
- **Socket.io-client** - 实时通信
- **Axios** - HTTP客户端

## 项目结构

```
frontend/
├── public/                    # 静态资源
├── src/
│   ├── assets/               # 资源文件
│   │   └── styles/          # 样式文件
│   ├── views/               # 页面组件
│   ├── router/              # 路由配置
│   ├── stores/              # 状态管理
│   ├── services/            # API服务
│   ├── utils/               # 工具函数
│   ├── types/               # TypeScript类型定义
│   ├── App.vue              # 根组件
│   └── main.ts              # 入口文件
├── .env                     # 环境变量
├── package.json
├── tsconfig.json
├── vite.config.ts
└── README.md
```

## 开发环境

### 环境要求

- Node.js >= 18.0.0
- npm >= 8.0.0

### 安装依赖

```bash
npm install
```

### 启动开发服务器

```bash
npm run dev
```

### 构建生产版本

```bash
npm run build
```

### 预览生产版本

```bash
npm run preview
```

### 代码检查

```bash
npm run lint
```

### 类型检查

```bash
npm run type-check
```

## 设计规范

### 色彩系统

- **主色调**: #FF6B35 (橙色)
- **辅助色**: #FFD93D (金黄色)
- **背景色**: #1A1A1A (深色主题)
- **文字色**: #FFFFFF (白色)

### 字体规范

- **中文字体**: PingFang SC, Hiragino Sans GB, Microsoft YaHei
- **英文字体**: SF Pro Display, Segoe UI, Roboto

### 响应式设计

- **移动端**: < 768px
- **平板端**: 768px - 1024px
- **桌面端**: > 1024px

## 开发规范

### 代码风格

- 使用 TypeScript 进行类型检查
- 遵循 ESLint 和 Prettier 代码规范
- 使用 Composition API 编写组件
- 采用 SCSS 进行样式开发

### 组件开发

- 组件命名使用 PascalCase
- 文件命名使用 kebab-case
- 样式使用 scoped 作用域
- 使用 TypeScript 定义 props 和 emits

### 状态管理

- 使用 Pinia 进行状态管理
- 按功能模块拆分 store
- 使用 TypeScript 定义 store 类型

## 部署

### 构建

```bash
npm run build
```

### 部署到服务器

将 `dist` 目录下的文件部署到 Web 服务器即可。

## 许可证

MIT License