# 认证相关 API

## 概述
认证相关API提供用户注册、登录、身份切换、Token管理等核心认证功能。

## API 列表

### 2.1 用户注册
**POST** `/auth/register`

**请求参数**:
```json
{
  "email": "user@example.com",
  "password": "password123",
  "phone": "13800138000",
  "primary_identity": {
    "identity_type": "apprentice",
    "domain": "software_development",
    "name": "张三"
  }
}
```

**响应**:
```json
{
  "code": 0,
  "message": "注册成功",
  "data": {
    "user_id": "uuid",
    "token": "jwt_token",
    "identities": [
      {
        "id": "uuid",
        "identity_type": "apprentice",
        "domain": "software_development",
        "status": "pending"
      }
    ]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 2.2 用户登录
**POST** `/auth/login`

**请求参数**:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "登录成功",
  "data": {
    "user_id": "uuid",
    "token": "jwt_token",
    "current_identity": {
      "id": "uuid",
      "identity_type": "apprentice",
      "domain": "software_development"
    },
    "identities": [
      {
        "id": "uuid",
        "identity_type": "apprentice",
        "domain": "software_development",
        "status": "active"
      }
    ]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 2.3 身份切换
**POST** `/auth/switch-identity`

**请求参数**:
```json
{
  "identity_id": "uuid"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "身份切换成功",
  "data": {
    "current_identity": {
      "id": "uuid",
      "identity_type": "master",
      "domain": "ui_design",
      "status": "active"
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 2.4 刷新Token
**POST** `/auth/refresh`

**响应**:
```json
{
  "code": 0,
  "message": "Token刷新成功",
  "data": {
    "token": "new_jwt_token"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 2.5 修改密码
**POST** `/auth/change-password`

**请求参数**:
```json
{
  "current_password": "old_password123",
  "new_password": "new_password123"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "密码修改成功",
  "data": {},
  "timestamp": "2024-12-01T10:00:00Z"
}
```

## 错误码

| 错误码 | 说明 |
|--------|------|
| 400 | 请求参数错误 |
| 401 | 用户名或密码错误 |
| 409 | 用户已存在 |
| 422 | 数据验证失败 |

## 注意事项

1. **密码安全**: 密码必须包含至少8个字符，包含大小写字母、数字和特殊字符
2. **邮箱验证**: 注册后需要验证邮箱才能激活账户
3. **Token安全**: Token有效期为24小时，建议定期刷新
4. **身份切换**: 只能切换到已激活的身份 