# Master Guide API 文档索引

## 概述
本文档是 Master Guide 平台 API 的完整索引，按功能模块分类组织，便于开发者快速查找和使用。

## 📋 文档结构

### 🔐 认证与用户管理
- **[API 概述](./api-summary.md)** - 基础信息、通用格式、错误处理
- **[认证相关 API](./auth-api.md)** - 注册、登录、身份切换、Token管理
- **[用户管理 API](./user-api.md)** - 用户资料、身份管理、统计、偏好设置

### 👨‍🏫 核心业务模块
- **[大师管理 API](./mentor-api.md)** - 大师信息、搜索、推荐、评价
- **[课程管理 API](./course-api.md)** - 课程创建、浏览、报名、学习
- **[预约管理 API](./appointment-api.md)** - 预约创建、管理、状态更新

### 💬 社交互动模块
- **[社群管理 API](./circle-api.md)** - 圈子、动态、评论、点赞
- **[评价管理 API](./review-api.md)** - 课程评价、大师评价
- **[通知管理 API](./notification-api.md)** - 消息通知、设置
- **[实时通信 API](./chat-api.md)** - 在线聊天、消息

### 📚 学习管理模块
- **[学习记录 API](./learning-api.md)** - 学习进度、作业、统计
- **[学生管理 API](./student-api.md)** - 学生信息、进度管理、沟通

### 💰 商业运营模块
- **[收入管理 API](./income-api.md)** - 收入统计、提现、报告
- **[支付管理 API](./payment-api.md)** - 支付订单、退款、统计

### 🔧 系统服务模块
- **[文件上传 API](./upload-api.md)** - 文件上传、管理
- **[搜索 API](./search-api.md)** - 全局搜索、过滤
- **[统计 API](./stats-api.md)** - 系统统计、数据分析
- **[WebSocket 事件](./websocket-api.md)** - 实时事件推送

## 🚀 快速开始

### 1. 环境准备
```bash
# API 基础URL
https://api.masterguide.com/v1

# 认证方式
Authorization: Bearer <your_jwt_token>
```

### 2. 认证流程
1. 调用 `POST /auth/register` 注册用户
2. 调用 `POST /auth/login` 获取 Token
3. 在后续请求中使用 Token

### 3. 示例请求
```bash
# 获取用户信息
curl -H "Authorization: Bearer <token>" \
     https://api.masterguide.com/v1/users/profile

# 获取大师列表
curl -H "Authorization: Bearer <token>" \
     https://api.masterguide.com/v1/mentors?domain=software_development
```

## 📖 使用指南

### 通用响应格式
所有 API 都遵循统一的响应格式：
```json
{
  "code": 0,
  "message": "success",
  "data": {},
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 分页参数
支持分页的 API 使用以下参数：
- `page`: 页码（从1开始）
- `page_size`: 每页数量（默认20，最大100）

### 错误处理
- 成功响应：`code: 0`
- 错误响应：`code: 非0值`，包含错误信息

## 🔧 开发工具

### API 测试
推荐使用以下工具进行 API 测试：
- [Postman](https://www.postman.com/)
- [Insomnia](https://insomnia.rest/)
- [curl](https://curl.se/)

### SDK 支持
- JavaScript/TypeScript SDK（开发中）
- Python SDK（计划中）
- Go SDK（计划中）

## 📞 技术支持

### 文档更新
- 本文档会随 API 版本更新
- 重大变更会提前通知
- 建议订阅更新通知

### 问题反馈
- 技术问题：提交 GitHub Issue
- 功能建议：通过反馈渠道提交
- 紧急问题：联系技术支持团队

## 📄 许可证

本文档遵循 AGPL v3 许可证，详见 [LICENSE](../LICENSE) 文件。

---

**最后更新**: 2024-12-01  
**版本**: v1.0  
**维护者**: Master Guide 开发团队 