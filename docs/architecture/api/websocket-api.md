## 18. WebSocket 事件

### 18.1 连接认证
```json
{
  "event": "authenticate",
  "data": {
    "token": "jwt_token"
  }
}
```

### 18.2 消息事件
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

### 18.3 在线状态事件
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