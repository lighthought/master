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
  },
  "timestamp": "2024-12-01T10:00:00Z"
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
  },
  "timestamp": "2024-12-01T10:00:00Z"
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

**响应**:
```json
{
  "code": 0,
  "message": "预约状态更新成功",
  "data": {
    "appointment_id": "uuid",
    "status": "confirmed"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 6.4 获取预约详情
**GET** `/appointments/{appointment_id}`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "appointment": {
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
      "price": 200.00,
      "notes": "想请教Go并发编程的问题",
      "created_at": "2024-12-01T10:00:00Z",
      "updated_at": "2024-12-01T15:30:00Z"
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 6.5 取消预约
**DELETE** `//{appointment_id}`

**响应**:
```json
{
  "code": 0,
  "message": "预约取消成功",
  "data": {
    "appointment_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 6.6 获取大师预约统计
**GET** `/appointments/mentor-stats`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "stats": {
      "total_appointments": 150,
      "pending_appointments": 5,
      "confirmed_appointments": 120,
      "completed_appointments": 100,
      "cancelled_appointments": 10,
      "total_earnings": 30000.00,
      "average_rating": 4.8
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
