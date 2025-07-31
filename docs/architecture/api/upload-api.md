## 15. 文件上传 API

### 15.1 上传文件
**POST** `/upload/file`

**请求参数**:
- `file`: 文件（multipart/form-data）
- `type`: 文件类型 (avatar, course_cover, post_image)

**响应**:
```json
{
  "code": 0,
  "message": "上传成功",
  "data": {
    "file_url": "https://example.com/uploads/file.jpg",
    "file_id": "uuid"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
