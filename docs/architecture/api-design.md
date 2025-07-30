# Master Guide API 设计文档

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

## 2. 认证相关 API

### 2.1 用户注册
**POST** `/auth/register`

**请求参数**:
```json
{
  "email": "user@example.com",
  "password": "password123",
  "phone": "13800138000",
  "primary_identity": {
    "identity_type": "apprentice",
    "domain": "software_development",
    "name": "张三"
  }
}
```

**响应**:
```json
{
  "code": 0,
  "message": "注册成功",
  "data": {
    "user_id": "uuid",
    "token": "jwt_token",
    "identities": [
      {
        "id": "uuid",
        "identity_type": "apprentice",
        "domain": "software_development",
        "status": "pending"
      }
    ]
  }
}
```

### 2.2 用户登录
**POST** `/auth/login`

**请求参数**:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "登录成功",
  "data": {
    "user_id": "uuid",
    "token": "jwt_token",
    "current_identity": {
      "id": "uuid",
      "identity_type": "apprentice",
      "domain": "software_development"
    },
    "identities": [
      {
        "id": "uuid",
        "identity_type": "apprentice",
        "domain": "software_development",
        "status": "active"
      }
    ]
  }
}
```

### 2.3 身份切换
**POST** `/auth/switch-identity`

**请求参数**:
```json
{
  "identity_id": "uuid"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "身份切换成功",
  "data": {
    "current_identity": {
      "id": "uuid",
      "identity_type": "master",
      "domain": "ui_design",
      "status": "active"
    }
  }
}
```

### 2.4 刷新Token
**POST** `/auth/refresh`

**响应**:
```json
{
  "code": 0,
  "message": "Token刷新成功",
  "data": {
    "token": "new_jwt_token"
  }
}
```

## 3. 用户管理 API

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
  }
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
  }
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

## 4. 大师管理 API

### 4.1 获取大师列表
**GET** `/mentors`

**查询参数**:
- `domain`: 专业领域
- `min_rating`: 最低评分
- `max_price`: 最高价格
- `is_online`: 是否在线
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "mentors": [
      {
        "id": "uuid",
        "identity_id": "uuid",
        "name": "李大师",
        "avatar": "https://example.com/avatar.jpg",
        "domain": "software_development",
        "rating": 4.8,
        "student_count": 150,
        "hourly_rate": 200.00,
        "is_online": true,
        "skills": ["Go", "Vue.js", "PostgreSQL"],
        "bio": "资深全栈开发工程师"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 100,
      "total_pages": 5
    }
  }
}
```

### 4.2 获取大师详情
**GET** `/mentors/{mentor_id}`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "mentor": {
      "id": "uuid",
      "identity_id": "uuid",
      "name": "李大师",
      "avatar": "https://example.com/avatar.jpg",
      "domain": "software_development",
      "rating": 4.8,
      "student_count": 150,
      "hourly_rate": 200.00,
      "is_online": true,
      "skills": ["Go", "Vue.js", "PostgreSQL"],
      "bio": "资深全栈开发工程师",
      "experience_years": 8,
      "courses": [
        {
          "id": "uuid",
          "title": "Go Web开发实战",
          "price": 299.00,
          "student_count": 50
        }
      ],
      "reviews": [
        {
          "id": "uuid",
          "rating": 5,
          "content": "老师讲解很详细，收获很大",
          "reviewer_name": "王同学",
          "created_at": "2024-12-01T10:00:00Z"
        }
      ]
    }
  }
}
```

### 4.3 搜索大师
**GET** `/mentors/search`

**查询参数**:
- `q`: 搜索关键词
- `domain`: 专业领域
- `min_rating`: 最低评分
- `max_price`: 最高价格
- `is_online`: 是否在线

## 5. 课程管理 API

### 5.1 获取课程列表
**GET** `/courses`

**查询参数**:
- `domain`: 专业领域
- `difficulty`: 难度级别
- `min_price`: 最低价格
- `max_price`: 最高价格
- `sort_by`: 排序方式 (rating, price, created_at)
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "courses": [
      {
        "id": "uuid",
        "title": "Go Web开发实战",
        "description": "从零开始学习Go Web开发",
        "cover_image": "https://example.com/cover.jpg",
        "price": 299.00,
        "duration_hours": 20,
        "difficulty": "intermediate",
        "student_count": 50,
        "rating": 4.8,
        "mentor": {
          "id": "uuid",
          "name": "李大师",
          "avatar": "https://example.com/avatar.jpg"
        }
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 100,
      "total_pages": 5
    }
  }
}
```

### 5.2 获取课程详情
**GET** `/courses/{course_id}`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "course": {
      "id": "uuid",
      "title": "Go Web开发实战",
      "description": "从零开始学习Go Web开发",
      "cover_image": "https://example.com/cover.jpg",
      "price": 299.00,
      "duration_hours": 20,
      "difficulty": "intermediate",
      "student_count": 50,
      "rating": 4.8,
      "mentor": {
        "id": "uuid",
        "name": "李大师",
        "avatar": "https://example.com/avatar.jpg"
      },
      "contents": [
        {
          "id": "uuid",
          "title": "第一章：Go基础语法",
          "content_type": "video",
          "duration_minutes": 45,
          "order_index": 1
        }
      ],
      "reviews": [
        {
          "id": "uuid",
          "rating": 5,
          "content": "课程内容很实用",
          "reviewer_name": "王同学",
          "created_at": "2024-12-01T10:00:00Z"
        }
      ]
    }
  }
}
```

### 5.3 创建课程（大师身份）
**POST** `/courses`

**请求参数**:
```json
{
  "title": "Go Web开发实战",
  "description": "从零开始学习Go Web开发",
  "cover_image": "https://example.com/cover.jpg",
  "price": 299.00,
  "duration_hours": 20,
  "difficulty": "intermediate",
  "max_students": 100,
  "contents": [
    {
      "title": "第一章：Go基础语法",
      "content_type": "video",
      "content_url": "https://example.com/video1.mp4",
      "duration_minutes": 45,
      "order_index": 1
    }
  ]
}
```

### 5.4 报名课程
**POST** `/courses/{course_id}/enroll`

**响应**:
```json
{
  "code": 0,
  "message": "报名成功",
  "data": {
    "enrollment_id": "uuid",
    "course_id": "uuid",
    "status": "enrolled"
  }
}
```

### 5.5 获取学习进度
**GET** `/courses/{course_id}/progress`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "course_id": "uuid",
    "progress_percentage": 65.5,
    "status": "learning",
    "enrolled_at": "2024-12-01T10:00:00Z",
    "last_accessed_at": "2024-12-01T15:30:00Z",
    "completed_contents": ["uuid1", "uuid2"]
  }
}
```

## 6. 预约管理 API

### 6.1 创建预约
**POST** `/appointments`

**请求参数**:
```json
{
  "mentor_id": "uuid",
  "appointment_time": "2024-12-02T14:00:00Z",
  "duration_minutes": 60,
  "meeting_type": "video",
  "notes": "想请教Go并发编程的问题"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "预约创建成功",
  "data": {
    "appointment_id": "uuid",
    "status": "pending",
    "price": 200.00
  }
}
```

### 6.2 获取预约列表
**GET** `/appointments`

**查询参数**:
- `status`: 预约状态
- `type`: 预约类型 (student/mentor)
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "appointments": [
      {
        "id": "uuid",
        "mentor": {
          "id": "uuid",
          "name": "李大师",
          "avatar": "https://example.com/avatar.jpg"
        },
        "student": {
          "id": "uuid",
          "name": "王同学",
          "avatar": "https://example.com/avatar.jpg"
        },
        "appointment_time": "2024-12-02T14:00:00Z",
        "duration_minutes": 60,
        "meeting_type": "video",
        "status": "confirmed",
        "price": 200.00
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 50,
      "total_pages": 3
    }
  }
}
```

### 6.3 更新预约状态
**PUT** `/appointments/{appointment_id}/status`

**请求参数**:
```json
{
  "status": "confirmed"
}
```

## 7. 社群管理 API

### 7.1 获取圈子列表
**GET** `/circles`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "circles": [
      {
        "id": "uuid",
        "name": "Go开发交流圈",
        "description": "Go语言开发技术交流",
        "domain": "software_development",
        "member_count": 1250,
        "is_joined": true
      }
    ]
  }
}
```

### 7.2 获取圈子动态
**GET** `/circles/{circle_id}/posts`

**查询参数**:
- `post_type`: 动态类型
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "posts": [
      {
        "id": "uuid",
        "user": {
          "id": "uuid",
          "name": "王同学",
          "avatar": "https://example.com/avatar.jpg"
        },
        "identity": {
          "identity_type": "apprentice",
          "domain": "software_development"
        },
        "content": "今天学习了Go的并发编程，感觉收获很大！",
        "media_urls": ["https://example.com/image1.jpg"],
        "post_type": "text",
        "like_count": 15,
        "comment_count": 8,
        "created_at": "2024-12-01T10:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 100,
      "total_pages": 5
    }
  }
}
```

### 7.3 发布动态
**POST** `/circles/{circle_id}/posts`

**请求参数**:
```json
{
  "content": "今天学习了Go的并发编程，感觉收获很大！",
  "media_urls": ["https://example.com/image1.jpg"],
  "post_type": "text"
}
```

### 7.4 点赞动态
**POST** `/posts/{post_id}/like`

### 7.5 取消点赞
**DELETE** `/posts/{post_id}/like`

### 7.6 获取评论列表
**GET** `/posts/{post_id}/comments`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "comments": [
      {
        "id": "uuid",
        "user": {
          "id": "uuid",
          "name": "李大师",
          "avatar": "https://example.com/avatar.jpg"
        },
        "identity": {
          "identity_type": "master",
          "domain": "software_development"
        },
        "content": "很棒！Go的并发编程确实很强大",
        "like_count": 5,
        "created_at": "2024-12-01T10:30:00Z"
      }
    ]
  }
}
```

### 7.7 发表评论
**POST** `/posts/{post_id}/comments`

**请求参数**:
```json
{
  "content": "很棒！Go的并发编程确实很强大"
}
```

## 8. 评价管理 API

### 8.1 创建评价
**POST** `/reviews`

**请求参数**:
```json
{
  "reviewed_id": "uuid",
  "course_id": "uuid",
  "rating": 5,
  "content": "课程内容很实用，老师讲解很详细",
  "review_type": "course"
}
```

### 8.2 获取评价列表
**GET** `/reviews`

**查询参数**:
- `reviewed_id`: 被评价者ID
- `review_type`: 评价类型
- `page`: 页码
- `page_size`: 每页数量

## 9. 实时通信 API

### 9.1 获取在线用户
**GET** `/chat/online-users`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "online_users": [
      {
        "user_id": "uuid",
        "name": "李大师",
        "avatar": "https://example.com/avatar.jpg",
        "is_online": true,
        "last_seen": "2024-12-01T10:00:00Z"
      }
    ]
  }
}
```

### 9.2 获取聊天记录
**GET** `/chat/messages`

**查询参数**:
- `target_id`: 目标用户ID
- `circle_id`: 圈子ID
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "messages": [
      {
        "id": "uuid",
        "from_user": {
          "id": "uuid",
          "name": "李大师",
          "avatar": "https://example.com/avatar.jpg"
        },
        "content": "你好，有什么问题需要帮助吗？",
        "type": "text",
        "created_at": "2024-12-01T10:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 50,
      "total": 100,
      "total_pages": 2
    }
  }
}
```

## 10. 文件上传 API

### 10.1 上传文件
**POST** `/upload/file`

**请求参数**:
- `file`: 文件（multipart/form-data）
- `type`: 文件类型 (avatar, course_cover, post_image)

**响应**:
```json
{
  "code": 0,
  "message": "上传成功",
  "data": {
    "file_url": "https://example.com/uploads/file.jpg",
    "file_id": "uuid"
  }
}
```

## 11. 搜索 API

### 11.1 全局搜索
**GET** `/search`

**查询参数**:
- `q`: 搜索关键词
- `type`: 搜索类型 (mentors, courses, posts)
- `domain`: 专业领域
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "mentors": [...],
    "courses": [...],
    "posts": [...],
    "total_results": 150
  }
}
```

## 12. 统计 API

### 12.1 获取用户统计
**GET** `/stats/user`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "learning_stats": {
      "enrolled_courses": 5,
      "completed_courses": 2,
      "total_study_hours": 45.5
    },
    "teaching_stats": {
      "total_students": 25,
      "total_income": 5000.00,
      "average_rating": 4.8
    }
  }
}
```

## 13. WebSocket 事件

### 13.1 连接认证
```json
{
  "event": "authenticate",
  "data": {
    "token": "jwt_token"
  }
}
```

### 13.2 消息事件
```json
{
  "event": "message",
  "data": {
    "id": "uuid",
    "from_user": {
      "id": "uuid",
      "name": "李大师"
    },
    "content": "你好！",
    "type": "text",
    "created_at": "2024-12-01T10:00:00Z"
  }
}
```

### 13.3 在线状态事件
```json
{
  "event": "user_status",
  "data": {
    "user_id": "uuid",
    "status": "online",
    "timestamp": "2024-12-01T10:00:00Z"
  }
}
```

## 14. 错误处理

### 14.1 验证错误
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

### 14.2 权限错误
```json
{
  "code": 403,
  "message": "权限不足",
  "data": {
    "required_permission": "master_identity"
  }
}
```

## 15. 限流说明

### 15.1 限流规则
- **认证接口**: 5次/分钟
- **普通接口**: 100次/分钟
- **文件上传**: 10次/分钟
- **WebSocket连接**: 10次/分钟

### 15.2 限流响应
```json
{
  "code": 429,
  "message": "请求过于频繁",
  "data": {
    "retry_after": 60
  }
}
```

## 16. 版本控制

### 16.1 版本策略
- 通过URL路径进行版本控制
- 当前版本: v1
- 向后兼容性保证

### 16.2 废弃通知
```json
{
  "code": 0,
  "message": "success",
  "data": {},
  "deprecation_warning": "此接口将在v2版本中废弃，请使用新的接口"
}
```

---

**文档版本**: v1.0  
**最后更新**: 2024-12-01  
**维护者**: Master Guide 技术团队 