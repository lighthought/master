## 12. 通知管理 API

### 12.1 获取通知列表
**GET** `/notifications`

**查询参数**:
- `type`: 通知类型 (message, system, activity, reminder)
- `status`: 通知状态 (unread, read, all)
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "notifications": [
      {
        "id": "uuid",
        "type": "message",
        "title": "新消息",
        "content": "李大师给您发送了一条消息",
        "status": "unread",
        "sender": {
          "id": "uuid",
          "name": "李大师",
          "avatar": "https://example.com/avatar.jpg"
        },
        "related_data": {
          "message_id": "uuid",
          "chat_type": "direct"
        },
        "created_at": "2024-12-01T10:00:00Z",
        "read_at": null
      },
      {
        "id": "uuid",
        "type": "system",
        "title": "系统通知",
        "content": "您的课程《Go Web开发实战》有新的作业发布",
        "status": "read",
        "related_data": {
          "course_id": "uuid",
          "assignment_id": "uuid"
        },
        "created_at": "2024-12-01T09:00:00Z",
        "read_at": "2024-12-01T09:30:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 50,
      "total_pages": 3
    },
    "unread_count": 5
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 12.2 标记通知为已读
**PUT** `/notifications/{notification_id}/read`

**响应**:
```json
{
  "code": 0,
  "message": "通知已标记为已读",
  "data": {
    "notification_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 12.3 批量标记通知为已读
**PUT** `/notifications/read`

**请求参数**:
```json
{
  "notification_ids": ["uuid1", "uuid2", "uuid3"],
  "mark_all": false
}
```

**响应**:
```json
{
  "code": 0,
  "message": "通知已批量标记为已读",
  "data": {
    "marked_count": 3
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 12.4 删除通知
**DELETE** `/notifications/{notification_id}`

**响应**:
```json
{
  "code": 0,
  "message": "通知删除成功",
  "data": {
    "notification_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 12.5 批量删除通知
**DELETE** `/notifications`

**请求参数**:
```json
{
  "notification_ids": ["uuid1", "uuid2", "uuid3"],
  "delete_all": false
}
```

**响应**:
```json
{
  "code": 0,
  "message": "通知批量删除成功",
  "data": {
    "deleted_count": 3
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 12.6 获取未读通知数量
**GET** `/notifications/unread-count`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "unread_count": 5,
    "count_by_type": {
      "message": 2,
      "system": 1,
      "activity": 1,
      "reminder": 1
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 12.7 获取通知设置
**GET** `/notifications/settings`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "settings": {
      "email_notifications": {
        "enabled": true,
        "types": ["message", "system", "activity"]
      },
      "push_notifications": {
        "enabled": true,
        "types": ["message", "system"]
      },
      "in_app_notifications": {
        "enabled": true,
        "types": ["message", "system", "activity", "reminder"]
      },
      "quiet_hours": {
        "enabled": false,
        "start_time": "22:00",
        "end_time": "08:00"
      }
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 12.8 更新通知设置
**PUT** `/notifications/settings`

**请求参数**:
```json
{
  "email_notifications": {
    "enabled": true,
    "types": ["message", "system", "activity"]
  },
  "push_notifications": {
    "enabled": true,
    "types": ["message", "system"]
  },
  "in_app_notifications": {
    "enabled": true,
    "types": ["message", "system", "activity", "reminder"]
  },
  "quiet_hours": {
    "enabled": false,
    "start_time": "22:00",
    "end_time": "08:00"
  }
}
```

**响应**:
```json
{
  "code": 0,
  "message": "通知设置更新成功",
  "data": {
    "settings_updated": true
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 12.9 发送通知（系统）
**POST** `/notifications/send`

**请求参数**:
```json
{
  "user_ids": ["uuid1", "uuid2"],
  "type": "system",
  "title": "系统维护通知",
  "content": "系统将于今晚22:00-24:00进行维护",
  "related_data": {
    "maintenance_id": "uuid"
  }
}
```

**响应**:
```json
{
  "code": 0,
  "message": "通知发送成功",
  "data": {
    "sent_count": 2,
    "notification_ids": ["uuid1", "uuid2"]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
