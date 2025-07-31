# 用户管理 API

## 概述
用户管理API提供用户资料管理、身份管理、统计信息、偏好设置等核心用户功能。

## API 列表

### 3.1 获取用户信息
**GET** `/users/profile`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "phone": "13800138000",
      "status": "active",
      "created_at": "2024-12-01T10:00:00Z"
    },
    "current_identity": {
      "id": "uuid",
      "identity_type": "apprentice",
      "domain": "software_development",
      "status": "active",
      "profile": {
        "name": "张三",
        "avatar": "https://example.com/avatar.jpg",
        "bio": "热爱学习的新手",
        "skills": ["JavaScript", "Vue.js"],
        "experience_years": 1
      }
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.2 更新用户档案
**PUT** `/users/profile`

**请求参数**:
```json
{
  "name": "张三",
  "avatar": "https://example.com/avatar.jpg",
  "bio": "热爱学习的新手",
  "skills": ["JavaScript", "Vue.js"],
  "experience_years": 1,
  "hourly_rate": 100.00
}
```

**响应**:
```json
{
  "code": 0,
  "message": "用户档案更新成功",
  "data": {},
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.3 获取用户身份列表
**GET** `/users/identities`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "identities": [
      {
        "id": "uuid",
        "identity_type": "apprentice",
        "domain": "software_development",
        "status": "active",
        "profile": {
          "name": "张三",
          "avatar": "https://example.com/avatar.jpg"
        }
      },
      {
        "id": "uuid",
        "identity_type": "master",
        "domain": "ui_design",
        "status": "pending",
        "profile": {
          "name": "张三",
          "avatar": "https://example.com/avatar.jpg"
        }
      }
    ]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.4 添加新身份
**POST** `/users/identities`

**请求参数**:
```json
{
  "identity_type": "master",
  "domain": "ui_design",
  "name": "张三",
  "bio": "UI设计专家",
  "skills": ["Figma", "Sketch", "Adobe XD"],
  "experience_years": 5,
  "hourly_rate": 200.00
}
```

**响应**:
```json
{
  "code": 0,
  "message": "身份创建成功",
  "data": {
    "identity_id": "uuid",
    "status": "pending"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.5 更新身份信息
**PUT** `/users/identities/{identity_id}`

**请求参数**:
```json
{
  "name": "张三",
  "bio": "UI设计专家",
  "skills": ["Figma", "Sketch", "Adobe XD"],
  "experience_years": 5,
  "hourly_rate": 200.00
}
```

**响应**:
```json
{
  "code": 0,
  "message": "身份信息更新成功",
  "data": {},
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.6 获取用户学习统计
**GET** `/users/stats/learning`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total_courses": 12,
    "progress": 65,
    "completed_lessons": 8,
    "total_lessons": 15,
    "current_course": "Vue.js 进阶开发",
    "next_lesson": "组件通信与状态管理"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.7 获取用户教学统计
**GET** `/users/stats/teaching`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total_students": 8,
    "total_hours": 24,
    "total_earnings": 2400,
    "average_rating": 4.8,
    "completed_sessions": 12,
    "upcoming_sessions": 3
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.8 获取用户通用统计
**GET** `/users/stats/general`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "active_days": 7,
    "achievements": 3,
    "total_login_days": 15,
    "last_login_date": "2024-01-15",
    "streak_days": 5
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.9 获取用户成就列表
**GET** `/users/achievements`

**查询参数**:
- `identity_type`: 身份类型 (master/apprentice)

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "achievements": [
      {
        "id": "uuid",
        "name": "学习新手",
        "description": "完成第一门课程",
        "icon": "🎓"
      },
      {
        "id": "uuid",
        "name": "坚持不懈",
        "description": "连续学习7天",
        "icon": "🔥"
      }
    ]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.10 获取用户偏好
**GET** `/users/preferences`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "learning_style": "one-on-one",
    "time_preference": "flexible",
    "budget_range": "medium",
    "learning_goals": ["掌握前端开发", "提升编程技能"],
    "preferred_domains": ["软件开发", "前端开发"],
    "experience_level": "beginner",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.11 保存用户偏好
**PUT** `/users/preferences`

**请求参数**:
```json
{
  "learning_style": "one-on-one",
  "time_preference": "flexible",
  "budget_range": "medium",
  "learning_goals": ["掌握前端开发", "提升编程技能"],
  "preferred_domains": ["软件开发", "前端开发"],
  "experience_level": "beginner"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "用户偏好保存成功",
  "data": {},
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.12 获取推荐学习路径
**GET** `/users/recommended-learning-path`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "recommended_path": "one-on-one",
    "confidence": 0.7,
    "reasons": ["基于用户偏好推荐", "适合初学者"]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 3.13 获取学习路径统计
**GET** `/users/learning-path-stats`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total_users": 1250,
    "path_distribution": {
      "one-on-one": 45,
      "structured": 30,
      "browse": 20,
      "other": 5
    },
    "satisfaction_rates": {
      "one-on-one": 4.8,
      "structured": 4.6,
      "browse": 4.4,
      "other": 4.2
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

## 错误码

| 错误码 | 说明 |
|--------|------|
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 422 | 数据验证失败 |

## 注意事项

1. **身份管理**: 用户可以拥有多个身份，但同一时间只能激活一个身份
2. **权限控制**: 某些API需要特定身份类型才能访问
3. **数据验证**: 所有输入数据都会进行格式和业务规则验证
4. **隐私保护**: 敏感信息（如手机号）会进行脱敏处理 