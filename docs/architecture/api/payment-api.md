# 支付管理 API

## 概述
支付管理API提供支付订单创建、状态查询、退款处理、支付统计等核心功能，支持多种支付方式和完整的支付流程。

## API 列表


### 13.1 创建支付订单
**POST** `/payments/orders`

**请求参数**:
```json
{
  "order_type": "course_enrollment",
  "order_id": "uuid",
  "amount": 299.00,
  "currency": "CNY",
  "payment_method": "alipay",
  "description": "Go Web开发实战课程报名",
  "metadata": {
    "course_id": "uuid",
    "course_title": "Go Web开发实战"
  }
}
```

**响应**:
```json
{
  "code": 0,
  "message": "支付订单创建成功",
  "data": {
    "order_id": "uuid",
    "payment_id": "uuid",
    "payment_url": "https://example.com/payment/gateway",
    "qr_code": "https://example.com/qr-code.png",
    "expires_at": "2024-12-01T11:00:00Z"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 13.2 查询支付状态
**GET** `/payments/orders/{order_id}/status`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "order_id": "uuid",
    "payment_id": "uuid",
    "status": "completed",
    "amount": 299.00,
    "currency": "CNY",
    "payment_method": "alipay",
    "paid_at": "2024-12-01T10:05:00Z",
    "transaction_id": "202412011234567890"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 13.3 获取支付历史
**GET** `/payments/history`

**查询参数**:
- `type`: 支付类型 (course_enrollment, appointment, refund)
- `status`: 支付状态 (pending, completed, failed, cancelled)
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
    "payments": [
      {
        "id": "uuid",
        "order_id": "uuid",
        "type": "course_enrollment",
        "amount": 299.00,
        "currency": "CNY",
        "payment_method": "alipay",
        "status": "completed",
        "description": "Go Web开发实战课程报名",
        "created_at": "2024-12-01T10:00:00Z",
        "paid_at": "2024-12-01T10:05:00Z",
        "transaction_id": "202412011234567890",
        "metadata": {
          "course_id": "uuid",
          "course_title": "Go Web开发实战"
        }
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

### 13.4 申请退款
**POST** `/payments/refunds`

**请求参数**:
```json
{
  "payment_id": "uuid",
  "amount": 299.00,
  "reason": "课程内容不符合预期",
  "description": "希望退款并选择其他课程"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "退款申请提交成功",
  "data": {
    "refund_id": "uuid",
    "status": "pending",
    "estimated_completion_time": "2024-12-03T10:00:00Z"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 13.5 查询退款状态
**GET** `/payments/refunds/{refund_id}/status`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "refund_id": "uuid",
    "payment_id": "uuid",
    "amount": 299.00,
    "status": "completed",
    "reason": "课程内容不符合预期",
    "created_at": "2024-12-01T10:00:00Z",
    "completed_at": "2024-12-01T14:00:00Z",
    "refund_transaction_id": "202412011234567890"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 13.6 获取支付方式列表
**GET** `/payments/methods`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "payment_methods": [
      {
        "id": "alipay",
        "name": "支付宝",
        "icon": "https://example.com/alipay-icon.png",
        "enabled": true,
        "min_amount": 0.01,
        "max_amount": 50000.00
      },
      {
        "id": "wechat",
        "name": "微信支付",
        "icon": "https://example.com/wechat-icon.png",
        "enabled": true,
        "min_amount": 0.01,
        "max_amount": 50000.00
      },
      {
        "id": "bank_card",
        "name": "银行卡",
        "icon": "https://example.com/bank-icon.png",
        "enabled": true,
        "min_amount": 1.00,
        "max_amount": 100000.00
      }
    ]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 13.7 获取支付统计
**GET** `/payments/stats`

**查询参数**:
- `period`: 统计周期 (day, week, month, quarter, year)
- `start_date`: 开始日期
- `end_date`: 结束日期

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "stats": {
      "total_amount": 50000.00,
      "total_transactions": 150,
      "successful_transactions": 145,
      "failed_transactions": 5,
      "refund_amount": 2000.00,
      "refund_count": 8,
      "payment_methods_distribution": {
        "alipay": 60,
        "wechat": 30,
        "bank_card": 10
      },
      "daily_stats": [
        {
          "date": "2024-12-01",
          "amount": 800.00,
          "transactions": 3
        }
      ]
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 13.8 支付回调处理
**POST** `/payments/webhook/{gateway}`

**请求参数**:
```json
{
  "order_id": "uuid",
  "transaction_id": "202412011234567890",
  "status": "success",
  "amount": 299.00,
  "signature": "abc123def456"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "回调处理成功",
  "data": {
    "processed": true
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```


## 错误码

| 错误码 | 说明 |
|--------|------|
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 订单不存在 |
| 409 | 订单状态不允许操作 |
| 422 | 数据验证失败 |

## 注意事项

1. **支付安全**: 所有支付相关操作都经过安全验证
2. **订单状态**: 支付订单有多种状态，需要按状态流转
3. **退款规则**: 退款需要符合平台规则和时间限制
4. **回调处理**: 支付网关回调需要验证签名 