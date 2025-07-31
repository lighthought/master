## 17. 统计 API

### 17.1 获取用户统计
**GET** `/stats/user`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "learning_stats": {
      "enrolled_courses": 5,
      "completed_courses": 2,
      "total_study_hours": 45.5
    },
    "teaching_stats": {
      "total_students": 25,
      "total_income": 5000.00,
      "average_rating": 4.8
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
