## 14. 实时通信 API

### 14.1 获取在线用户
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
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 14.2 获取聊天记录
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
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
