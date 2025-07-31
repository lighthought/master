# 社群管理 API

## 概述
社群管理API提供圈子创建、动态发布、评论互动、点赞等社交功能，支持用户建立学习社群和分享学习心得。

## API 列表

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
  },
  "timestamp": "2024-12-01T10:00:00Z"
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
  },
  "timestamp": "2024-12-01T10:00:00Z"
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

**响应**:
```json
{
  "code": 0,
  "message": "动态发布成功",
  "data": {
    "post_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.4 点赞动态
**POST** `/posts/{post_id}/like`

**响应**:
```json
{
  "code": 0,
  "message": "点赞成功",
  "data": {
    "post_id": "uuid",
    "like_count": 16
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.5 取消点赞
**DELETE** `/posts/{post_id}/like`

**响应**:
```json
{
  "code": 0,
  "message": "取消点赞成功",
  "data": {
    "post_id": "uuid",
    "like_count": 15
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.6 获取评论列表
**GET** `/posts/{post_id}/comments`

**查询参数**:
- `page`: 页码
- `page_size`: 每页数量

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
          "name": "李同学",
          "avatar": "https://example.com/avatar.jpg"
        },
        "identity": {
          "identity_type": "apprentice",
          "domain": "software_development"
        },
        "content": "确实很棒！我也在学习这个",
        "like_count": 3,
        "is_liked": false,
        "created_at": "2024-12-01T11:00:00Z",
        "replies": [
          {
            "id": "uuid",
            "user": {
              "id": "uuid",
              "name": "王同学",
              "avatar": "https://example.com/avatar.jpg"
            },
            "content": "一起加油！",
            "like_count": 1,
            "is_liked": false,
            "created_at": "2024-12-01T11:30:00Z"
          }
        ]
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

### 7.7 发表评论
**POST** `/posts/{post_id}/comments`

**请求参数**:
```json
{
  "content": "确实很棒！我也在学习这个"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "评论发表成功",
  "data": {
    "comment_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.8 回复评论
**POST** `/comments/{comment_id}/replies`

**请求参数**:
```json
{
  "content": "一起加油！"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "回复发表成功",
  "data": {
    "reply_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.9 点赞评论
**POST** `/comments/{comment_id}/like`

**响应**:
```json
{
  "code": 0,
  "message": "点赞成功",
  "data": {
    "comment_id": "uuid",
    "like_count": 4
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.10 取消点赞评论
**DELETE** `/comments/{comment_id}/like`

**响应**:
```json
{
  "code": 0,
  "message": "取消点赞成功",
  "data": {
    "comment_id": "uuid",
    "like_count": 3
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.11 点赞回复
**POST** `/replies/{reply_id}/like`

**响应**:
```json
{
  "code": 0,
  "message": "点赞成功",
  "data": {
    "reply_id": "uuid",
    "like_count": 2
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.12 取消点赞回复
**DELETE** `/replies/{reply_id}/like`

**响应**:
```json
{
  "code": 0,
  "message": "取消点赞成功",
  "data": {
    "reply_id": "uuid",
    "like_count": 1
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.13 删除评论
**DELETE** `/comments/{comment_id}`

**响应**:
```json
{
  "code": 0,
  "message": "评论删除成功",
  "data": {
    "comment_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.14 删除回复
**DELETE** `/replies/{reply_id}`

**响应**:
```json
{
  "code": 0,
  "message": "回复删除成功",
  "data": {
    "reply_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.15 加入圈子
**POST** `/circles/{circle_id}/join`

**响应**:
```json
{
  "code": 0,
  "message": "加入圈子成功",
  "data": {
    "circle_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.16 退出圈子
**DELETE** `/circles/{circle_id}/join`

**响应**:
```json
{
  "code": 0,
  "message": "退出圈子成功",
  "data": {
    "circle_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 7.17 获取推荐圈子
**GET** `/circles/recommended`

**查询参数**:
- `user_id`: 用户ID（可选，用于个性化推荐）

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "circles": [
      {
        "id": "uuid",
        "name": "Vue.js开发圈",
        "description": "Vue.js前端开发技术交流",
        "domain": "software_development",
        "member_count": 800,
        "is_joined": false,
        "recommendation_reason": "基于您的学习兴趣推荐"
      }
    ]
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
| 409 | 已加入该圈子 |
| 422 | 数据验证失败 |

## 注意事项

1. **权限控制**: 圈子创建者和管理员有特殊权限
2. **内容审核**: 动态和评论内容会进行审核
3. **互动限制**: 防止恶意刷赞和评论
4. **隐私设置**: 支持公开和私密圈子 