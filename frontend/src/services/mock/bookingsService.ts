import type { User } from '@/types/user'

// 模拟延迟
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟预约数据
const mockBookings = [
  {
    id: '1',
    mentorId: '1',
    userId: '1',
    mentorName: '张大师',
    userName: '测试用户',
    date: '2024-01-20',
    timeSlot: '14:00-15:00',
    method: 'video',
    requirements: '希望学习Vue.js组件化开发的最佳实践',
    status: 'pending',
    price: 200,
    createdAt: '2024-01-15T10:30:00Z',
    updatedAt: '2024-01-15T10:30:00Z'
  },
  {
    id: '2',
    mentorId: '2',
    userId: '1',
    mentorName: '李导师',
    userName: '测试用户',
    date: '2024-01-22',
    timeSlot: '19:00-20:00',
    method: 'voice',
    requirements: '需要了解Spring Boot微服务架构设计',
    status: 'confirmed',
    price: 180,
    createdAt: '2024-01-14T15:20:00Z',
    updatedAt: '2024-01-15T09:00:00Z'
  }
]

export const mockBookingsService = {
  // 创建预约
  async createBooking(bookingData: {
    mentorId: string
    userId: string
    date: string
    timeSlot: string
    method: string
    requirements: string
  }) {
    await delay(1000)
    
    // 模拟预约创建
    const newBooking = {
      id: Date.now().toString(),
      ...bookingData,
      mentorName: '大师', // 这里应该从大师数据中获取
      userName: '用户', // 这里应该从用户数据中获取
      status: 'pending',
      price: 200, // 这里应该从大师数据中获取
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    }
    
    mockBookings.push(newBooking)
    
    return {
      success: true,
      data: newBooking,
      message: '预约创建成功'
    }
  },

  // 获取用户预约列表
  async getUserBookings(userId: string, status?: string) {
    await delay(500)
    
    let filteredBookings = mockBookings.filter(booking => booking.userId === userId)
    
    if (status) {
      filteredBookings = filteredBookings.filter(booking => booking.status === status)
    }
    
    return {
      success: true,
      data: filteredBookings
    }
  },

  // 获取大师预约列表
  async getMentorBookings(mentorId: string, status?: string) {
    await delay(500)
    
    let filteredBookings = mockBookings.filter(booking => booking.mentorId === mentorId)
    
    if (status) {
      filteredBookings = filteredBookings.filter(booking => booking.status === status)
    }
    
    return {
      success: true,
      data: filteredBookings
    }
  },

  // 更新预约状态
  async updateBookingStatus(bookingId: string, status: 'pending' | 'confirmed' | 'rejected' | 'completed' | 'cancelled') {
    await delay(600)
    
    const booking = mockBookings.find(b => b.id === bookingId)
    if (!booking) {
      throw new Error('预约不存在')
    }
    
    booking.status = status
    booking.updatedAt = new Date().toISOString()
    
    return {
      success: true,
      data: booking,
      message: '预约状态更新成功'
    }
  },

  // 取消预约
  async cancelBooking(bookingId: string, reason?: string) {
    await delay(500)
    
    const booking = mockBookings.find(b => b.id === bookingId)
    if (!booking) {
      throw new Error('预约不存在')
    }
    
    if (booking.status === 'completed') {
      throw new Error('已完成的预约无法取消')
    }
    
    booking.status = 'cancelled'
    booking.updatedAt = new Date().toISOString()
    
    return {
      success: true,
      data: booking,
      message: '预约取消成功'
    }
  },

  // 获取预约详情
  async getBookingDetail(bookingId: string) {
    await delay(300)
    
    const booking = mockBookings.find(b => b.id === bookingId)
    if (!booking) {
      throw new Error('预约不存在')
    }
    
    return {
      success: true,
      data: booking
    }
  }
}