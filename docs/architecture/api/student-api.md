## 11. 学生管理 API

### 11.1 获取学生列表
**GET** `/students`

**查询参数**:
- `status`: 学生状态 (active, inactive, graduated)
- `course_id`: 课程ID（可选）
- `search`: 搜索关键词（姓名、邮箱）
- `sort_by`: 排序方式 (name, enrollment_date, progress)
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "students": [
      {
        "id": "uuid",
        "name": "王同学",
        "avatar": "https://example.com/avatar.jpg",
        "email": "wang@example.com",
        "phone": "138****1234",
        "enrollment_date": "2024-12-01T10:00:00Z",
        "status": "active",
        "total_courses": 3,
        "completed_courses": 1,
        "total_study_hours": 45.5,
        "average_progress": 68.2,
        "last_activity": "2024-12-01T15:30:00Z",
        "current_courses": [
          {
            "course_id": "uuid",
            "title": "Go Web开发实战",
            "progress_percentage": 65.5,
            "status": "learning"
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

### 11.2 获取学生详情
**GET** `/students/{student_id}`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "student": {
      "id": "uuid",
      "name": "王同学",
      "avatar": "https://example.com/avatar.jpg",
      "email": "wang@example.com",
      "phone": "138****1234",
      "bio": "热爱编程的在校学生",
      "enrollment_date": "2024-12-01T10:00:00Z",
      "status": "active",
      "learning_goals": ["掌握Go开发", "提升编程技能"],
      "preferred_learning_style": "hands-on",
      "courses": [
        {
          "course_id": "uuid",
          "title": "Go Web开发实战",
          "enrollment_date": "2024-12-01T10:00:00Z",
          "progress_percentage": 65.5,
          "status": "learning",
          "last_study_date": "2024-12-01T15:30:00Z",
          "total_study_time": 120,
          "assignments": [
            {
              "id": "uuid",
              "title": "Web API开发作业",
              "status": "submitted",
              "score": 85,
              "submitted_at": "2024-12-01T16:00:00Z"
            }
          ]
        }
      ],
      "appointments": [
        {
          "id": "uuid",
          "appointment_time": "2024-12-02T14:00:00Z",
          "status": "confirmed",
          "topic": "Go并发编程问题"
        }
      ],
      "reviews": [
        {
          "id": "uuid",
          "rating": 5,
          "content": "老师讲解很详细",
          "created_at": "2024-12-01T10:00:00Z"
        }
      ]
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 11.3 获取学生统计
**GET** `/students/stats`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "stats": {
      "total_students": 50,
      "active_students": 45,
      "inactive_students": 3,
      "graduated_students": 2,
      "new_students_this_month": 8,
      "average_progress": 68.2,
      "average_rating": 4.8,
      "top_performing_students": 12,
      "students_by_course": [
        {
          "course_id": "uuid",
          "course_title": "Go Web开发实战",
          "student_count": 25,
          "average_progress": 72.5
        }
      ]
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 11.4 发送消息给学生
**POST** `/students/{student_id}/messages`

**请求参数**:
```json
{
  "content": "你好，关于你的作业我有一些建议",
  "type": "text"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "消息发送成功",
  "data": {
    "message_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 11.5 获取与学生聊天记录
**GET** `/students/{student_id}/messages`

**查询参数**:
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
        "to_user": {
          "id": "uuid",
          "name": "王同学",
          "avatar": "https://example.com/avatar.jpg"
        },
        "content": "你好，关于你的作业我有一些建议",
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

### 11.6 更新学生学习进度
**PUT** `/students/{student_id}/courses/{course_id}/progress`

**请求参数**:
```json
{
  "progress_percentage": 70.0,
  "notes": "学生表现很好，建议进入下一章节"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "学习进度更新成功",
  "data": {
    "student_id": "uuid",
    "course_id": "uuid",
    "progress_percentage": 70.0
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 11.7 评价学生作业
**POST** `/students/{student_id}/assignments/{assignment_id}/grade`

**请求参数**:
```json
{
  "score": 85,
  "feedback": "代码结构清晰，但可以优化性能",
  "comments": "建议使用更高效的数据结构"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "作业评价成功",
  "data": {
    "assignment_id": "uuid",
    "score": 85
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 11.8 获取学生学习报告
**GET** `/students/{student_id}/report`

**查询参数**:
- `period`: 报告周期 (week, month, quarter, year)

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "report": {
      "student_id": "uuid",
      "student_name": "王同学",
      "period": "month",
      "study_time": 45.5,
      "courses_progress": [
        {
          "course_id": "uuid",
          "title": "Go Web开发实战",
          "progress_percentage": 65.5,
          "study_time": 30.0,
          "assignments_completed": 3,
          "average_score": 85.6
        }
      ],
      "strengths": ["编程逻辑清晰", "学习态度积极"],
      "areas_for_improvement": ["需要加强算法思维", "可以多练习项目实战"],
      "recommendations": ["建议学习数据结构", "可以尝试开源项目"]
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
