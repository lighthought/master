## 16. 搜索 API

### 16.1 全局搜索
**GET** `/search`

**查询参数**:
- `q`: 搜索关键词
- `type`: 搜索类型 (mentors, courses, posts)
- `domain`: 专业领域
- `page`: 页码
- `page_size`: 每页数量

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "mentors": [...],
    "courses": [...],
    "posts": [...],
    "total_results": 150
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
