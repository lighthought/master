// 导入所有API模块
import { authApi } from './auth'
import { userApi } from './user'
import { mentorApi } from './mentor'
import { appointmentApi } from './appointment'
import { courseApi } from './course'
import { learningApi } from './learning'

// 统一的API服务类
export class ApiService {
  // 认证相关API
  static auth = authApi

  // 用户管理API
  static user = userApi

  // 大师管理API
  static mentors = mentorApi

  // 预约管理API
  static appointments = appointmentApi

  // 课程管理API
  static courses = courseApi

  // 学习记录管理API
  static learning = learningApi

  // 预约管理API (待实现)
  static bookings = {
    // 创建预约
    async createBooking(bookingData: {
      mentor_id: string
      appointment_time: string
      duration_minutes: number
      meeting_type: string
      notes?: string
    }) {
      const response = await fetch('https://api.masterguide.com/v1/appointments', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        },
        body: JSON.stringify(bookingData)
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '创建预约失败')
      }
      
      return await response.json()
    },

    // 获取预约列表
    async getAppointments(params: {
      status?: string
      type?: string
      page?: number
      page_size?: number
    } = {}) {
      const searchParams = new URLSearchParams()
      Object.entries(params).forEach(([key, value]) => {
        if (value !== undefined) {
          searchParams.append(key, value.toString())
        }
      })
      
      const response = await fetch(`https://api.masterguide.com/v1/appointments?${searchParams}`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '获取预约列表失败')
      }
      
      return await response.json()
    },

    // 更新预约状态
    async updateBookingStatus(appointmentId: string, status: string) {
      const response = await fetch(`https://api.masterguide.com/v1/appointments/${appointmentId}/status`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        },
        body: JSON.stringify({ status })
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '更新预约状态失败')
      }
      
      return await response.json()
    },

    // 取消预约
    async cancelBooking(appointmentId: string) {
      const response = await fetch(`https://api.masterguide.com/v1/appointments/${appointmentId}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '取消预约失败')
      }
      
      return await response.json()
    }
  }

  // 圈子管理API (待实现)
  static circles = {
    // 获取圈子列表
    async getCircles(params: {
      domain?: string
      page?: number
      page_size?: number
    } = {}) {
      const searchParams = new URLSearchParams()
      Object.entries(params).forEach(([key, value]) => {
        if (value !== undefined) {
          searchParams.append(key, value.toString())
        }
      })
      
      const response = await fetch(`https://api.masterguide.com/v1/circles?${searchParams}`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '获取圈子列表失败')
      }
      
      return await response.json()
    },

    // 加入圈子
    async joinCircle(circleId: string) {
      const response = await fetch(`https://api.masterguide.com/v1/circles/${circleId}/join`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '加入圈子失败')
      }
      
      return await response.json()
    },

    // 退出圈子
    async leaveCircle(circleId: string) {
      const response = await fetch(`https://api.masterguide.com/v1/circles/${circleId}/join`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '退出圈子失败')
      }
      
      return await response.json()
    },

    // 获取推荐圈子
    async getRecommendedCircles(userId?: string) {
      const params = userId ? `?user_id=${userId}` : ''
      const response = await fetch(`https://api.masterguide.com/v1/circles/recommended${params}`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '获取推荐圈子失败')
      }
      
      return await response.json()
    }
  }

  // 动态管理API (待实现)
  static posts = {
    // 获取动态列表
    async getPosts(params: {
      circle_id?: string
      post_type?: string
      page?: number
      page_size?: number
    } = {}) {
      const searchParams = new URLSearchParams()
      Object.entries(params).forEach(([key, value]) => {
        if (value !== undefined) {
          searchParams.append(key, value.toString())
        }
      })
      
      const response = await fetch(`https://api.masterguide.com/v1/posts?${searchParams}`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '获取动态列表失败')
      }
      
      return await response.json()
    },

    // 发布动态
    async createPost(circleId: string, postData: {
      content: string
      media_urls?: string[]
      post_type: string
    }) {
      const response = await fetch(`https://api.masterguide.com/v1/circles/${circleId}/posts`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        },
        body: JSON.stringify(postData)
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '发布动态失败')
      }
      
      return await response.json()
    },

    // 点赞/取消点赞动态
    async toggleLike(postId: string) {
      const response = await fetch(`https://api.masterguide.com/v1/posts/${postId}/like`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '操作失败')
      }
      
      return await response.json()
    },

    // 获取评论列表
    async getComments(postId: string, params: {
      page?: number
      page_size?: number
    } = {}) {
      const searchParams = new URLSearchParams()
      Object.entries(params).forEach(([key, value]) => {
        if (value !== undefined) {
          searchParams.append(key, value.toString())
        }
      })
      
      const response = await fetch(`https://api.masterguide.com/v1/posts/${postId}/comments?${searchParams}`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '获取评论列表失败')
      }
      
      return await response.json()
    },

    // 发表评论
    async addComment(postId: string, commentData: {
      content: string
    }) {
      const response = await fetch(`https://api.masterguide.com/v1/posts/${postId}/comments`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        },
        body: JSON.stringify(commentData)
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '发表评论失败')
      }
      
      return await response.json()
    }
  }

  // 支付管理API (待实现)
  static payments = {
    // 创建支付订单
    async createOrder(orderData: {
      order_type: string
      order_id: string
      amount: number
      currency: string
      payment_method: string
      description: string
      metadata: any
    }) {
      const response = await fetch('https://api.masterguide.com/v1/payments/orders', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        },
        body: JSON.stringify(orderData)
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '创建支付订单失败')
      }
      
      return await response.json()
    },

    // 查询支付状态
    async getOrderStatus(orderId: string) {
      const response = await fetch(`https://api.masterguide.com/v1/payments/orders/${orderId}/status`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '查询支付状态失败')
      }
      
      return await response.json()
    },

    // 获取支付方式列表
    async getPaymentMethods() {
      const response = await fetch('https://api.masterguide.com/v1/payments/methods', {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || '获取支付方式失败')
      }
      
      return await response.json()
    }
  }
}

// 导出默认实例
export default ApiService 