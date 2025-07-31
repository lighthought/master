# API 文档拆分说明

## 已完成拆分的模块

✅ **API 概述** - `api-summary.md`  
✅ **认证相关 API** - `auth-api.md`  
✅ **用户管理 API** - `user-api.md`  
✅ **大师管理 API** - `mentor-api.md`  
✅ **社群管理 API** - `community-api.md`  
✅ **支付管理 API** - `payment-api.md`  

## 待拆分的模块

### 核心业务模块
- [ ] **课程管理 API** - `course-api.md` (已有占位符)
- [ ] **预约管理 API** - `booking-api.md`
- [ ] **学习记录 API** - `learning-api.md`
- [ ] **学生管理 API** - `student-api.md`
- [ ] **收入管理 API** - `income-api.md`

### 社交互动模块
- [ ] **评价管理 API** - `review-api.md`
- [ ] **通知管理 API** - `notification-api.md`
- [ ] **实时通信 API** - `chat-api.md`

### 系统服务模块
- [ ] **文件上传 API** - `upload-api.md`
- [ ] **搜索 API** - `search-api.md`
- [ ] **统计 API** - `stats-api.md`
- [ ] **WebSocket 事件** - `websocket-api.md`

## 拆分步骤

1. **提取内容**: 从 `api-design.md` 中提取对应章节
2. **格式化**: 统一格式，添加概述和注意事项
3. **更新索引**: 在 `README.md` 中更新链接
4. **验证**: 检查链接和格式是否正确

## 文件结构

```
docs/architecture/api/
├── README.md                    # 主索引文件
├── api-summary.md              # API 概述
├── auth-api.md                 # 认证相关 API
├── user-api.md                 # 用户管理 API
├── mentor-api.md               # 大师管理 API
├── course-api.md               # 课程管理 API
├── booking-api.md              # 预约管理 API
├── community-api.md            # 社群管理 API
├── review-api.md               # 评价管理 API
├── learning-api.md             # 学习记录 API
├── income-api.md               # 收入管理 API
├── student-api.md              # 学生管理 API
├── notification-api.md         # 通知管理 API
├── payment-api.md              # 支付管理 API
├── chat-api.md                 # 实时通信 API
├── upload-api.md               # 文件上传 API
├── search-api.md               # 搜索 API
├── stats-api.md                # 统计 API
├── websocket-api.md            # WebSocket 事件
└── split-api-docs.md           # 拆分说明（本文件）
```

## 注意事项

1. **保持一致性**: 所有文档使用相同的格式和结构
2. **链接更新**: 确保所有内部链接正确
3. **版本控制**: 保持与原文档的版本同步
4. **完整性**: 确保所有API都被正确拆分

## 下一步行动

建议按以下顺序完成剩余模块的拆分：

1. 课程管理 API (核心功能)
2. 预约管理 API (核心功能)
3. 学习记录 API (学习管理)
4. 收入管理 API (商业运营)
5. 学生管理 API (学习管理)
6. 其他模块按优先级完成 