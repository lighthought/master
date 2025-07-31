## 9. 学习记录 API

### 9.1 获取学习记录列表
**GET** `/learning-records`

**查询参数**:
- `course_id`: 课程ID（可选）
- `status`: 学习状态 (learning, completed, paused)
- `start_date`: 开始日期
- `end_date`: 结束日期
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "records": [
      {
        "id": "uuid",
        "course": {
          "id": "uuid",
          "title": "Go Web开发实战",
          "cover_image": "https://example.com/cover.jpg"
        },
        "mentor": {
          "id": "uuid",
          "name": "李大师",
          "avatar": "https://example.com/avatar.jpg"
        },
        "enrollment_date": "2024-12-01T10:00:00Z",
        "last_study_date": "2024-12-01T15:30:00Z",
        "total_study_time": 120,
        "progress_percentage": 65.5,
        "status": "learning",
        "current_chapter": "第三章：Web框架开发",
        "completed_chapters": ["第一章", "第二章"],
        "certificate_issued": false
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

### 9.2 获取学习记录详情
**GET** `/learning-records/{record_id}`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "record": {
      "id": "uuid",
      "course": {
        "id": "uuid",
        "title": "Go Web开发实战",
        "description": "从零开始学习Go Web开发",
        "cover_image": "https://example.com/cover.jpg",
        "duration_hours": 20
      },
      "mentor": {
        "id": "uuid",
        "name": "李大师",
        "avatar": "https://example.com/avatar.jpg"
      },
      "enrollment_date": "2024-12-01T10:00:00Z",
      "last_study_date": "2024-12-01T15:30:00Z",
      "total_study_time": 120,
      "progress_percentage": 65.5,
      "status": "learning",
      "current_chapter": "第三章：Web框架开发",
      "completed_chapters": ["第一章", "第二章"],
      "study_sessions": [
        {
          "id": "uuid",
          "start_time": "2024-12-01T14:00:00Z",
          "end_time": "2024-12-01T15:30:00Z",
          "duration_minutes": 90,
          "chapter": "第三章：Web框架开发",
          "notes": "学习了Gin框架的基本使用"
        }
      ],
      "assignments": [
        {
          "id": "uuid",
          "title": "Web API开发作业",
          "status": "submitted",
          "submitted_at": "2024-12-01T16:00:00Z",
          "score": 85
        }
      ],
      "certificate_issued": false,
      "certificate_url": null
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 9.3 更新学习进度
**PUT** `/learning-records/{record_id}/progress`

**请求参数**:
```json
{
  "progress_percentage": 70.0,
  "current_chapter": "第四章：数据库集成",
  "study_time_minutes": 30
}
```

**响应**:
```json
{
  "code": 0,
  "message": "学习进度更新成功",
  "data": {
    "record_id": "uuid",
    "progress_percentage": 70.0
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 9.4 提交作业
**POST** `/learning-records/{record_id}/assignments`

**请求参数**:
```json
{
  "title": "Web API开发作业",
  "content": "完成了用户管理API的开发",
  "attachment_urls": ["https://example.com/assignment.zip"]
}
```

**响应**:
```json
{
  "code": 0,
  "message": "作业提交成功",
  "data": {
    "assignment_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 9.5 获取学习统计
**GET** `/learning-records/stats`

**查询参数**:
- `period`: 统计周期 (week, month, year, all)

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "stats": {
      "total_courses": 5,
      "completed_courses": 2,
      "total_study_hours": 45.5,
      "average_progress": 68.2,
      "current_streak_days": 7,
      "total_assignments": 15,
      "completed_assignments": 12,
      "average_score": 85.6,
      "certificates_earned": 2
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 9.6 获取学习路径推荐
**GET** `/learning-records/recommended-path`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "path": {
      "current_level": "intermediate",
      "next_courses": [
        {
          "id": "uuid",
          "title": "Go微服务架构",
          "reason": "基于您当前的学习进度推荐",
          "estimated_duration": 25
        }
      ],
      "skills_to_develop": ["微服务设计", "服务治理", "容器化部署"],
      "estimated_completion_time": "3个月"
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
