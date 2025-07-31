# 大师管理 API

## 概述
大师管理API提供大师信息查询、搜索、推荐、评价等核心功能，支持学徒寻找合适的大师进行学习。

## API 列表

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
  },
  "timestamp": "2024-12-01T10:00:00Z"
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
  },
  "timestamp": "2024-12-01T10:00:00Z"
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
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 4.4 获取推荐大师
**GET** `/mentors/recommended`

**查询参数**:
- `user_id`: 用户ID（可选，用于个性化推荐）

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
        "bio": "资深全栈开发工程师",
        "recommendation_reason": "基于您的学习偏好推荐"
      }
    ]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 4.5 获取大师评价
**GET** `/mentors/{mentor_id}/reviews`

**查询参数**:
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "reviews": [
      {
        "id": "uuid",
        "rating": 5,
        "content": "老师讲解很详细，收获很大",
        "reviewer_name": "王同学",
        "reviewer_avatar": "https://example.com/avatar.jpg",
        "created_at": "2024-12-01T10:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 50,
      "total_pages": 5
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

## 错误码

| 错误码 | 说明 |
|--------|------|
| 400 | 请求参数错误 |
| 404 | 大师不存在 |
| 422 | 数据验证失败 |

## 注意事项

1. **在线状态**: 大师的在线状态会实时更新，建议定期刷新
2. **评分系统**: 评分基于学生评价计算，最低1分，最高5分
3. **价格范围**: 价格单位为人民币元，支持按价格区间筛选
4. **推荐算法**: 推荐基于用户偏好、学习历史、大师评分等因素 