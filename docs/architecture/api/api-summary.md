# Master Guide API 设计文档 - 概述

## 1. API 概述

### 1.1 基础信息
- **API版本**: v1.0
- **基础URL**: `https://api.masterguide.com/v1`
- **认证方式**: JWT Bearer Token
- **数据格式**: JSON
- **字符编码**: UTF-8

### 1.2 通用响应格式
```json
{
  "code": 0,
  "message": "success",
  "data": {},
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 1.3 错误码定义
| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 409 | 资源冲突 |
| 422 | 数据验证失败 |
| 500 | 服务器内部错误 |

### 1.4 分页参数
```json
{
  "page": 1,
  "page_size": 20,
  "total": 100,
  "total_pages": 5
}
```

## 2. API 模块概览

### 2.1 核心业务模块
- **[认证相关 API](./auth-api.md)** - 用户注册、登录、身份切换等
- **[用户管理 API](./user-api.md)** - 用户资料、身份管理、偏好设置等
- **[大师管理 API](./mentor-api.md)** - 大师信息、推荐、评价等
- **[课程管理 API](./course-api.md)** - 课程创建、浏览、报名、学习等
- **[预约管理 API](./appointment-api.md)** - 预约创建、管理、状态更新等

### 2.2 社交互动模块
- **[社群管理 API](./circle-api.md)** - 圈子、动态、评论、点赞等
- **[评价管理 API](./review-api.md)** - 课程评价、大师评价等
- **[通知管理 API](./notification-api.md)** - 消息通知、设置等
- **[实时通信 API](./chat-api.md)** - 在线聊天、消息等

### 2.3 学习管理模块
- **[学习记录 API](./learning-api.md)** - 学习进度、作业、统计等
- **[学生管理 API](./student-api.md)** - 学生信息、进度管理、沟通等

### 2.4 商业运营模块
- **[收入管理 API](./income-api.md)** - 收入统计、提现、报告等
- **[支付管理 API](./payment-api.md)** - 支付订单、退款、统计等

### 2.5 系统服务模块
- **[文件上传 API](./upload-api.md)** - 文件上传、管理
- **[搜索 API](./search-api.md)** - 全局搜索、过滤
- **[统计 API](./stats-api.md)** - 系统统计、数据分析
- **[WebSocket 事件](./websocket-api.md)** - 实时事件推送

## 3. 错误处理

### 3.1 错误响应格式
```json
{
  "code": 400,
  "message": "请求参数错误",
  "errors": [
    {
      "field": "email",
      "message": "邮箱格式不正确"
    }
  ],
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.2 常见错误场景
- **参数验证失败**: 返回 422 状态码，包含具体字段错误信息
```json
{
  "code": 422,
  "message": "数据验证失败",
  "data": {
    "errors": {
      "email": ["邮箱格式不正确"],
      "password": ["密码长度至少6位"]
    }
  }
}
```

- **资源不存在**: 返回 404 状态码
```json
{
  "code": 404,
  "message": "课程不存在"
}
```

- **权限不足**: 返回 403 状态码
```json
{
  "code": 403,
  "message": "权限不足",
  "data": {
    "required_permission": "master_identity"
  }
}
```
- **认证失败**: 返回 401 状态码
```json
{
  "code": 401,
  "message": "用户名活密码错误"
}
```

- **业务逻辑错误**: 返回 409 状态码
```json
{
  "code": 409,
  "message": "已加入该圈子"
}
```

## 4. 限流说明

### 4.1 限流策略
- **认证接口**: 每分钟最多 5 次请求
- **普通接口**: 每分钟最多 100 次请求
- **文件上传**: 每分钟最多 10 次请求
- **搜索接口**: 每分钟最多 50 次请求
- **WebSocket连接**: 每分钟最多 10 次

### 4.2 限流响应
```json
{
  "code": 429,
  "message": "请求过于频繁，请稍后再试",
  "retry_after": 60,
  "timestamp": "2024-12-01T10:00:00Z"
}
```

## 5. 版本控制

### 5.1 版本策略
- **主版本号**: 不兼容的 API 变更
- **次版本号**: 向后兼容的功能性新增
- **修订版本号**: 向后兼容的问题修正

### 5.2 版本兼容性
- 支持同时运行多个 API 版本
- 新版本发布后，旧版本继续支持 12 个月
- 版本废弃前 6 个月开始发送通知

### 5.3 版本切换
- 通过 URL 路径指定版本: `/v1/`, `/v2/`
- 通过请求头指定版本: `Accept: application/vnd.masterguide.v1+json`
- 默认使用最新稳定版本

## 6. 安全规范

### 6.1 认证要求
- 除公开接口外，所有 API 都需要 JWT 认证
- Token 有效期为 24 小时
- 支持 Token 刷新机制

### 6.2 数据安全
- 所有敏感数据传输使用 HTTPS
- 密码等敏感信息不记录日志
- 支持请求签名验证

### 6.3 访问控制
- 基于角色的访问控制 (RBAC)
- 支持细粒度的权限控制
- 操作日志记录和审计
