import { mockIncomeService } from '../mock/incomeService'

const isDevelopment = import.meta.env.DEV
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:3000/api'

// 收入管理API
export const incomeApi = {
  // 获取收入统计
  async getIncomeStats(params: {
    period?: string
    start_date?: string
    end_date?: string
  } = {}) {
    if (isDevelopment) {
      return await mockIncomeService.getIncomeStats('1', params)
    }

    const searchParams = new URLSearchParams()
    if (params.period) searchParams.append('period', params.period)
    if (params.start_date) searchParams.append('start_date', params.start_date)
    if (params.end_date) searchParams.append('end_date', params.end_date)

    const response = await fetch(`${API_BASE_URL}/income/stats?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取收入统计失败')
    }

    return await response.json()
  },

  // 获取收入明细
  async getIncomeTransactions(params: {
    type?: string
    status?: string
    start_date?: string
    end_date?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockIncomeService.getIncomeDetails('1', params)
    }

    const searchParams = new URLSearchParams()
    if (params.type) searchParams.append('type', params.type)
    if (params.status) searchParams.append('status', params.status)
    if (params.start_date) searchParams.append('start_date', params.start_date)
    if (params.end_date) searchParams.append('end_date', params.end_date)
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.page_size) searchParams.append('page_size', params.page_size.toString())

    const response = await fetch(`${API_BASE_URL}/income/transactions?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取收入明细失败')
    }

    return await response.json()
  },

  // 获取收入趋势
  async getIncomeTrends(params: {
    period?: string
    start_date?: string
    end_date?: string
  } = {}) {
    if (isDevelopment) {
      // 模拟收入趋势数据
      return {
        code: 0,
        message: 'success',
        data: {
          trends: [
            {
              date: '2024-12-01',
              income: 800.00,
              transactions: 3,
              course_enrollments: 2,
              appointments: 1
            },
            {
              date: '2024-12-02',
              income: 1200.00,
              transactions: 4,
              course_enrollments: 3,
              appointments: 1
            }
          ]
        },
        timestamp: new Date().toISOString()
      }
    }

    const searchParams = new URLSearchParams()
    if (params.period) searchParams.append('period', params.period)
    if (params.start_date) searchParams.append('start_date', params.start_date)
    if (params.end_date) searchParams.append('end_date', params.end_date)

    const response = await fetch(`${API_BASE_URL}/income/trends?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取收入趋势失败')
    }

    return await response.json()
  },

  // 导出收入报告
  async exportIncomeReport(params: {
    format?: string
    start_date?: string
    end_date?: string
    type?: string
  } = {}) {
    if (isDevelopment) {
      return await mockIncomeService.exportIncomeReport('1', params)
    }

    const searchParams = new URLSearchParams()
    if (params.format) searchParams.append('format', params.format)
    if (params.start_date) searchParams.append('start_date', params.start_date)
    if (params.end_date) searchParams.append('end_date', params.end_date)
    if (params.type) searchParams.append('type', params.type)

    const response = await fetch(`${API_BASE_URL}/income/export?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '导出收入报告失败')
    }

    return await response.json()
  },

  // 获取提现记录
  async getWithdrawals(params: {
    status?: string
    start_date?: string
    end_date?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      // 模拟提现记录
      return {
        code: 0,
        message: 'success',
        data: {
          withdrawals: [
            {
              id: 'uuid',
              amount: 5000.00,
              status: 'completed',
              bank_account: '****1234',
              created_at: '2024-12-01T10:00:00Z',
              completed_at: '2024-12-01T14:00:00Z',
              fee: 10.00,
              net_amount: 4990.00
            }
          ],
          pagination: {
            page: 1,
            page_size: 20,
            total: 25,
            total_pages: 2
          }
        },
        timestamp: new Date().toISOString()
      }
    }

    const searchParams = new URLSearchParams()
    if (params.status) searchParams.append('status', params.status)
    if (params.start_date) searchParams.append('start_date', params.start_date)
    if (params.end_date) searchParams.append('end_date', params.end_date)
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.page_size) searchParams.append('page_size', params.page_size.toString())

    const response = await fetch(`${API_BASE_URL}/income/withdrawals?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取提现记录失败')
    }

    return await response.json()
  },

  // 申请提现
  async requestWithdrawal(withdrawalData: {
    amount: number
    bank_account: string
    bank_name: string
  }) {
    if (isDevelopment) {
      // 模拟申请提现
      return {
        code: 0,
        message: '提现申请提交成功',
        data: {
          withdrawal_id: 'uuid',
          estimated_completion_time: '2024-12-03T10:00:00Z'
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/income/withdrawals`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(withdrawalData)
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '申请提现失败')
    }

    return await response.json()
  },

  // 获取可提现金额
  async getAvailableAmount() {
    if (isDevelopment) {
      // 模拟可提现金额
      return {
        code: 0,
        message: 'success',
        data: {
          available_amount: 15000.00,
          pending_amount: 2000.00,
          total_earned: 50000.00,
          total_withdrawn: 33000.00,
          min_withdrawal: 100.00,
          max_withdrawal: 15000.00
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/income/available`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取可提现金额失败')
    }

    return await response.json()
  },

  // 支付相关API
  // 创建支付订单
  async createPaymentOrder(orderData: {
    order_type: string
    order_id: string
    amount: number
    currency: string
    payment_method: string
    description: string
    metadata: any
  }) {
    if (isDevelopment) {
      // 模拟创建支付订单
      return {
        code: 0,
        message: '支付订单创建成功',
        data: {
          order_id: 'uuid',
          payment_id: 'uuid',
          payment_url: 'https://example.com/payment/gateway',
          qr_code: 'https://example.com/qr-code.png',
          expires_at: new Date(Date.now() + 30 * 60 * 1000).toISOString()
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/payments/orders`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(orderData)
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '创建支付订单失败')
    }

    return await response.json()
  },

  // 查询支付状态
  async getPaymentStatus(orderId: string) {
    if (isDevelopment) {
      // 模拟支付状态
      return {
        code: 0,
        message: 'success',
        data: {
          order_id: orderId,
          payment_id: 'uuid',
          status: 'completed',
          amount: 299.00,
          currency: 'CNY',
          payment_method: 'alipay',
          paid_at: new Date().toISOString(),
          transaction_id: '202412011234567890'
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/payments/orders/${orderId}/status`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '查询支付状态失败')
    }

    return await response.json()
  },

  // 获取支付历史
  async getPaymentHistory(params: {
    type?: string
    status?: string
    start_date?: string
    end_date?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      // 模拟支付历史
      return {
        code: 0,
        message: 'success',
        data: {
          payments: [
            {
              id: 'uuid',
              order_id: 'uuid',
              type: 'course_enrollment',
              amount: 299.00,
              currency: 'CNY',
              payment_method: 'alipay',
              status: 'completed',
              description: 'Go Web开发实战课程报名',
              created_at: '2024-12-01T10:00:00Z',
              paid_at: '2024-12-01T10:05:00Z',
              transaction_id: '202412011234567890',
              metadata: {
                course_id: 'uuid',
                course_title: 'Go Web开发实战'
              }
            }
          ],
          pagination: {
            page: 1,
            page_size: 20,
            total: 50,
            total_pages: 3
          }
        },
        timestamp: new Date().toISOString()
      }
    }

    const searchParams = new URLSearchParams()
    if (params.type) searchParams.append('type', params.type)
    if (params.status) searchParams.append('status', params.status)
    if (params.start_date) searchParams.append('start_date', params.start_date)
    if (params.end_date) searchParams.append('end_date', params.end_date)
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.page_size) searchParams.append('page_size', params.page_size.toString())

    const response = await fetch(`${API_BASE_URL}/payments/history?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取支付历史失败')
    }

    return await response.json()
  },

  // 申请退款
  async requestRefund(refundData: {
    payment_id: string
    amount: number
    reason: string
    description: string
  }) {
    if (isDevelopment) {
      // 模拟申请退款
      return {
        code: 0,
        message: '退款申请提交成功',
        data: {
          refund_id: 'uuid',
          status: 'pending',
          estimated_completion_time: '2024-12-03T10:00:00Z'
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/payments/refunds`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(refundData)
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '申请退款失败')
    }

    return await response.json()
  },

  // 查询退款状态
  async getRefundStatus(refundId: string) {
    if (isDevelopment) {
      // 模拟退款状态
      return {
        code: 0,
        message: 'success',
        data: {
          refund_id: refundId,
          payment_id: 'uuid',
          amount: 299.00,
          status: 'completed',
          reason: '课程内容不符合预期',
          created_at: '2024-12-01T10:00:00Z',
          completed_at: '2024-12-01T14:00:00Z',
          refund_transaction_id: '202412011234567890'
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/payments/refunds/${refundId}/status`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '查询退款状态失败')
    }

    return await response.json()
  },

  // 获取支付方式列表
  async getPaymentMethods() {
    if (isDevelopment) {
      // 模拟支付方式列表
      return {
        code: 0,
        message: 'success',
        data: {
          payment_methods: [
            {
              id: 'alipay',
              name: '支付宝',
              icon: 'https://example.com/alipay-icon.png',
              enabled: true,
              min_amount: 0.01,
              max_amount: 50000.00
            },
            {
              id: 'wechat',
              name: '微信支付',
              icon: 'https://example.com/wechat-icon.png',
              enabled: true,
              min_amount: 0.01,
              max_amount: 50000.00
            },
            {
              id: 'bank_card',
              name: '银行卡',
              icon: 'https://example.com/bank-icon.png',
              enabled: true,
              min_amount: 1.00,
              max_amount: 100000.00
            }
          ]
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/payments/methods`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取支付方式列表失败')
    }

    return await response.json()
  },

  // 获取支付统计
  async getPaymentStats(params: {
    period?: string
    start_date?: string
    end_date?: string
  } = {}) {
    if (isDevelopment) {
      // 模拟支付统计
      return {
        code: 0,
        message: 'success',
        data: {
          stats: {
            total_amount: 50000.00,
            total_transactions: 150,
            successful_transactions: 145,
            failed_transactions: 5,
            refund_amount: 2000.00,
            refund_count: 8,
            payment_methods_distribution: {
              alipay: 60,
              wechat: 30,
              bank_card: 10
            },
            daily_stats: [
              {
                date: '2024-12-01',
                amount: 800.00,
                transactions: 3
              }
            ]
          }
        },
        timestamp: new Date().toISOString()
      }
    }

    const searchParams = new URLSearchParams()
    if (params.period) searchParams.append('period', params.period)
    if (params.start_date) searchParams.append('start_date', params.start_date)
    if (params.end_date) searchParams.append('end_date', params.end_date)

    const response = await fetch(`${API_BASE_URL}/payments/stats?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取支付统计失败')
    }

    return await response.json()
  },

  // 兼容旧接口的方法
  async getIncomeStatsLegacy(masterId: string, params: any = {}) {
    return await this.getIncomeStats(params)
  },

  async getIncomeDetailsLegacy(masterId: string, params: any = {}) {
    return await this.getIncomeTransactions(params)
  },

  async exportIncomeReportLegacy(masterId: string, params: any = {}) {
    return await this.exportIncomeReport(params)
  }
} 