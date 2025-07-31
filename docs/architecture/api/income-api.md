## 10. 收入管理 API

### 10.1 获取收入统计
**GET** `/income/stats`

**查询参数**:
- `period`: 统计周期 (week, month, quarter, year, all)
- `start_date`: 开始日期
- `end_date`: 结束日期

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "stats": {
      "total_income": 50000.00,
      "total_transactions": 150,
      "average_per_transaction": 333.33,
      "income_by_source": {
        "course_enrollments": 35000.00,
        "appointments": 15000.00
      },
      "income_by_period": {
        "current_month": 8000.00,
        "previous_month": 7500.00,
        "current_quarter": 22000.00,
        "current_year": 50000.00
      },
      "growth_rate": 6.67,
      "top_courses": [
        {
          "course_id": "uuid",
          "title": "Go Web开发实战",
          "income": 15000.00,
          "enrollment_count": 50
        }
      ]
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 10.2 获取收入明细
**GET** `/income/transactions`

**查询参数**:
- `type`: 收入类型 (course_enrollment, appointment, refund)
- `status`: 交易状态 (completed, pending, failed)
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
    "transactions": [
      {
        "id": "uuid",
        "type": "course_enrollment",
        "amount": 299.00,
        "status": "completed",
        "description": "Go Web开发实战课程报名",
        "student_name": "王同学",
        "course_title": "Go Web开发实战",
        "created_at": "2024-12-01T10:00:00Z",
        "completed_at": "2024-12-01T10:05:00Z",
        "platform_fee": 29.90,
        "net_income": 269.10
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 150,
      "total_pages": 8
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 10.3 获取收入趋势
**GET** `/income/trends`

**查询参数**:
- `period`: 趋势周期 (daily, weekly, monthly)
- `start_date`: 开始日期
- `end_date`: 结束日期

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "trends": [
      {
        "date": "2024-12-01",
        "income": 800.00,
        "transactions": 3,
        "course_enrollments": 2,
        "appointments": 1
      },
      {
        "date": "2024-12-02",
        "income": 1200.00,
        "transactions": 4,
        "course_enrollments": 3,
        "appointments": 1
      }
    ]
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 10.4 导出收入报告
**GET** `/income/export`

**查询参数**:
- `format`: 导出格式 (csv, excel, pdf)
- `start_date`: 开始日期
- `end_date`: 结束日期
- `type`: 收入类型 (all, course_enrollment, appointment)

**响应**:
```json
{
  "code": 0,
  "message": "报告生成成功",
  "data": {
    "download_url": "https://example.com/reports/income_report_20241201.pdf",
    "expires_at": "2024-12-08T10:00:00Z"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 10.5 获取提现记录
**GET** `/income/withdrawals`

**查询参数**:
- `status`: 提现状态 (pending, completed, failed)
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
    "withdrawals": [
      {
        "id": "uuid",
        "amount": 5000.00,
        "status": "completed",
        "bank_account": "****1234",
        "created_at": "2024-12-01T10:00:00Z",
        "completed_at": "2024-12-01T14:00:00Z",
        "fee": 10.00,
        "net_amount": 4990.00
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 25,
      "total_pages": 2
    }
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 10.6 申请提现
**POST** `/income/withdrawals`

**请求参数**:
```json
{
  "amount": 5000.00,
  "bank_account": "1234567890",
  "bank_name": "中国银行"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "提现申请提交成功",
  "data": {
    "withdrawal_id": "uuid",
    "estimated_completion_time": "2024-12-03T10:00:00Z"
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```

### 10.7 获取可提现金额
**GET** `/income/available`

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "available_amount": 15000.00,
    "pending_amount": 2000.00,
    "total_earned": 50000.00,
    "total_withdrawn": 33000.00,
    "min_withdrawal": 100.00,
    "max_withdrawal": 15000.00
  },
  "timestamp": "2024-12-01T10:00:00Z"
}
```
