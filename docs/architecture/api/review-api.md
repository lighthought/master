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

**响应**:
```json
{
  "code": 0,
  "message": "评价创建成功",
  "data": {
    "review_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 8.2 获取评价列表
**GET** `/reviews`

**查询参数**:
- `reviewed_id`: 被评价对象ID
- `review_type`: 评价类型 (course, mentor)
- `rating`: 评分筛选
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
        "reviewer": {
          "id": "uuid",
          "name": "王同学",
          "avatar": "https://example.com/avatar.jpg"
        },
        "reviewed_id": "uuid",
        "course_id": "uuid",
        "rating": 5,
        "content": "课程内容很实用，老师讲解很详细",
        "review_type": "course",
        "created_at": "2024-12-01T10:00:00Z"
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

### 8.3 获取评价详情
**GET** `/reviews/{review_id}`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "review": {
      "id": "uuid",
      "reviewer": {
        "id": "uuid",
        "name": "王同学",
        "avatar": "https://example.com/avatar.jpg"
      },
      "reviewed_id": "uuid",
      "course_id": "uuid",
      "rating": 5,
      "content": "课程内容很实用，老师讲解很详细",
      "review_type": "course",
      "created_at": "2024-12-01T10:00:00Z",
      "updated_at": "2024-12-01T10:00:00Z"
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 8.4 更新评价
**PUT** `/reviews/{review_id}`

**请求参数**:
```json
{
  "rating": 4,
  "content": "课程内容很实用，老师讲解很详细，但希望增加更多实践案例"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "评价更新成功",
  "data": {
    "review_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 8.5 删除评价
**DELETE** `/reviews/{review_id}`

**响应**:
```json
{
  "code": 0,
  "message": "评价删除成功",
  "data": {
    "review_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 8.6 获取评价统计
**GET** `/reviews/stats`

**查询参数**:
- `reviewed_id`: 被评价对象ID
- `review_type`: 评价类型 (course, mentor)

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "stats": {
      "total_reviews": 150,
      "average_rating": 4.6,
      "rating_distribution": {
        "5": 80,
        "4": 45,
        "3": 15,
        "2": 7,
        "1": 3
      }
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
