# 课程管理 API

## 概述
课程管理API提供课程创建、浏览、报名、学习等核心功能，支持大师创建课程和学生报名学习。

## API 列表

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
  },
  "timestamp": "2024-12-01T10:00:00Z"
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
  },
  "timestamp": "2024-12-01T10:00:00Z"
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

**响应**:
```json
{
  "code": 0,
  "message": "课程创建成功",
  "data": {
    "course_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 5.4 报名课程
**POST** `/courses/{course_id}/enroll`

**请求参数**:
```json
{
  "payment_method": "alipay",
  "user_info": {
    "name": "张三",
    "phone": "13800138000"
  }
}
```

**响应**:
```json
{
  "code": 0,
  "message": "报名成功",
  "data": {
    "enrollment_id": "uuid",
    "course_id": "uuid",
    "status": "enrolled",
    "payment_url": "https://example.com/payment"
  },
  "timestamp": "2024-12-01T10:00:00Z"
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
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 5.6 搜索课程
**GET** `/courses/search`

**查询参数**:
- `q`: 搜索关键词
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
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 5.7 获取推荐课程
**GET** `/courses/recommended`

**查询参数**:
- `user_id`: 用户ID（可选，用于个性化推荐）

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "courses": [
      {
        "id": "uuid",
        "title": "Vue.js 进阶开发",
        "description": "深入学习Vue.js高级特性",
        "cover_image": "https://example.com/cover.jpg",
        "price": 399.00,
        "duration_hours": 25,
        "difficulty": "intermediate",
        "student_count": 80,
        "rating": 4.9,
        "mentor": {
          "id": "uuid",
          "name": "王大师",
          "avatar": "https://example.com/avatar.jpg"
        },
        "recommendation_reason": "基于您的学习历史推荐"
      }
    ]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 5.8 获取用户已报名课程
**GET** `/courses/enrolled`

**查询参数**:
- `status`: 课程状态 (learning, completed, paused)
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
        },
        "enrollment_status": "learning",
        "progress_percentage": 65.5,
        "enrolled_at": "2024-12-01T10:00:00Z",
        "last_accessed_at": "2024-12-01T15:30:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 50,
      "total_pages": 3
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
| 404 | 课程不存在 |
| 409 | 已报名该课程 |
| 422 | 数据验证失败 |

## 注意事项

1. **权限控制**: 只有大师身份可以创建和编辑课程
2. **报名限制**: 已报名的课程不能重复报名
3. **课程状态**: 课程有草稿、发布、下架等状态
4. **学习进度**: 支持学习进度跟踪和续学功能