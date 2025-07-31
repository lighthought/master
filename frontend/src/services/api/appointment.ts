import { mockBookingsService } from '../mock/bookingsService'

// 环境配置
const isDevelopment = import.meta.env.DEV
const API_BASE_URL = 'https://api.masterguide.com/v1'

// 预约管理API
export const appointmentApi = {
  // 创建预约
  async createAppointment(appointmentData: {
    mentor_id: string
    appointment_time: string
    duration_minutes: number
    meeting_type: string
    notes?: string
  }) {
    if (isDevelopment) {
      // 适配mock服务的类型
      const mockData = {
        mentorId: appointmentData.mentor_id,
        userId: '1', // 从当前用户获取
        date: appointmentData.appointment_time.split('T')[0],
        timeSlot: appointmentData.appointment_time.split('T')[1].substring(0, 5),
        method: appointmentData.meeting_type,
        duration: appointmentData.duration_minutes / 60,
        requirements: appointmentData.notes || ''
      }
      return await mockBookingsService.createBooking(mockData)
    }
    
    const response = await fetch(`${API_BASE_URL}/appointments`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(appointmentData)
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
    if (isDevelopment) {
      // 根据type参数决定调用哪个mock方法
      if (params.type === 'mentor') {
        return await mockBookingsService.getMentorBookings('', params.status)
      } else {
        return await mockBookingsService.getUserBookings('', params.status)
      }
    }
    
    const searchParams = new URLSearchParams()
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        searchParams.append(key, value.toString())
      }
    })
    
    const response = await fetch(`${API_BASE_URL}/appointments?${searchParams}`, {
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
  async updateAppointmentStatus(appointmentId: string, status: string) {
    if (isDevelopment) {
      return await mockBookingsService.updateBookingStatus(appointmentId, status as any)
    }
    
    const response = await fetch(`${API_BASE_URL}/appointments/${appointmentId}/status`, {
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

  // 获取预约详情
  async getAppointmentDetail(appointmentId: string) {
    if (isDevelopment) {
      return await mockBookingsService.getBookingDetail(appointmentId)
    }
    
    const response = await fetch(`${API_BASE_URL}/appointments/${appointmentId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取预约详情失败')
    }
    
    return await response.json()
  },

  // 取消预约
  async cancelAppointment(appointmentId: string) {
    if (isDevelopment) {
      return await mockBookingsService.cancelBooking(appointmentId)
    }
    
    const response = await fetch(`${API_BASE_URL}/appointments/${appointmentId}`, {
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
  },

  // 获取大师预约统计
  async getMentorStats() {
    if (isDevelopment) {
      return {
        code: 0,
        message: 'success',
        data: {
          stats: {
            total_appointments: 150,
            pending_appointments: 5,
            confirmed_appointments: 120,
            completed_appointments: 100,
            cancelled_appointments: 10,
            total_earnings: 30000.00,
            average_rating: 4.8
          }
        },
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/appointments/mentor-stats`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取大师预约统计失败')
    }
    
    return await response.json()
  },

  // 兼容旧接口的方法
  // 创建预约 (兼容旧接口)
  async createBooking(bookingData: any) {
    return await this.createAppointment(bookingData)
  },

  // 获取用户预约列表 (兼容旧接口)
  async getUserBookings(userId: string, status?: string) {
    return await this.getAppointments({ status, type: 'student' })
  },

  // 获取大师预约列表 (兼容旧接口)
  async getMentorBookings(mentorId: string, status?: string) {
    return await this.getAppointments({ status, type: 'mentor' })
  },

  // 更新预约状态 (兼容旧接口)
  async updateBookingStatus(bookingId: string, status: string) {
    return await this.updateAppointmentStatus(bookingId, status)
  },

  // 取消预约 (兼容旧接口)
  async cancelBooking(bookingId: string, reason?: string) {
    return await this.cancelAppointment(bookingId)
  },

  // 获取预约详情 (兼容旧接口)
  async getBookingDetail(bookingId: string) {
    return await this.getAppointmentDetail(bookingId)
  }
} 