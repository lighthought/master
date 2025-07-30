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
    mentorAvatar: 'https://via.placeholder.com/50x50/4CAF50/FFFFFF?text=张',
    studentName: '李学徒',
    studentAvatar: 'https://via.placeholder.com/50x50/2196F3/FFFFFF?text=李',
    studentDomain: '前端开发',
    studentGoal: '掌握Vue.js开发技能',
    userName: '测试用户',
    date: '2024-01-20',
    timeSlot: '14:00-15:00',
    method: 'video',
    duration: 1,
    requirements: '希望学习Vue.js组件化开发的最佳实践，包括组件设计原则、状态管理、性能优化等方面。',
    status: 'pending',
    price: 200,
    mentorMessage: '',
    createdAt: '2024-01-15T10:30:00Z',
    updatedAt: '2024-01-15T10:30:00Z'
  },
  {
    id: '2',
    mentorId: '2',
    userId: '1',
    mentorName: '李导师',
    mentorAvatar: 'https://via.placeholder.com/50x50/2196F3/FFFFFF?text=李',
    studentName: '王学员',
    studentAvatar: 'https://via.placeholder.com/50x50/FF9800/FFFFFF?text=王',
    studentDomain: '后端开发',
    studentGoal: '掌握Spring Boot开发',
    userName: '测试用户',
    date: '2024-01-22',
    timeSlot: '19:00-20:00',
    method: 'voice',
    duration: 2,
    requirements: '需要了解Spring Boot微服务架构设计，包括服务拆分、API网关、配置中心等核心概念。',
    status: 'confirmed',
    price: 360,
    mentorMessage: '已确认预约，请准时参加。如有问题请提前联系。',
    createdAt: '2024-01-14T15:20:00Z',
    updatedAt: '2024-01-15T09:00:00Z'
  },
  {
    id: '3',
    mentorId: '3',
    userId: '1',
    mentorName: '王老师',
    mentorAvatar: 'https://via.placeholder.com/50x50/FF9800/FFFFFF?text=王',
    studentName: '陈同学',
    studentAvatar: 'https://via.placeholder.com/50x50/9C27B0/FFFFFF?text=陈',
    studentDomain: '移动开发',
    studentGoal: '掌握Flutter开发',
    userName: '测试用户',
    date: '2024-01-18',
    timeSlot: '10:00-11:00',
    method: 'video',
    duration: 1,
    requirements: '学习Flutter跨平台开发，了解Widget体系结构和状态管理。',
    status: 'completed',
    price: 220,
    mentorMessage: '指导已完成，请及时提交学习反馈。',
    createdAt: '2024-01-13T14:30:00Z',
    updatedAt: '2024-01-18T11:00:00Z'
  },
  {
    id: '4',
    mentorId: '4',
    userId: '1',
    mentorName: '陈工程师',
    mentorAvatar: 'https://via.placeholder.com/50x50/9C27B0/FFFFFF?text=陈',
    studentName: '赵学生',
    studentAvatar: 'https://via.placeholder.com/50x50/E91E63/FFFFFF?text=赵',
    studentDomain: '数据科学',
    studentGoal: '掌握机器学习基础',
    userName: '测试用户',
    date: '2024-01-25',
    timeSlot: '15:00-16:00',
    method: 'text',
    duration: 1,
    requirements: '了解机器学习基础算法，包括监督学习和无监督学习的应用场景。',
    status: 'cancelled',
    price: 300,
    mentorMessage: '预约已取消',
    createdAt: '2024-01-16T11:20:00Z',
    updatedAt: '2024-01-17T16:45:00Z'
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
    duration?: number
    requirements: string
  }) {
    await delay(1000)
    
    // 模拟预约创建
    const newBooking = {
      id: Date.now().toString(),
      ...bookingData,
      duration: bookingData.duration || 1,
      mentorName: '大师', // 这里应该从大师数据中获取
      mentorAvatar: 'https://via.placeholder.com/50x50/4CAF50/FFFFFF?text=大',
      studentName: '学生', // 这里应该从用户数据中获取
      studentAvatar: 'https://via.placeholder.com/50x50/2196F3/FFFFFF?text=学',
      studentDomain: '前端开发',
      studentGoal: '掌握开发技能',
      userName: '用户', // 这里应该从用户数据中获取
      status: 'pending',
      price: 200, // 这里应该从大师数据中获取
      mentorMessage: '',
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
    
    // 确保返回的数据包含学生信息
    const bookingsWithStudentInfo = filteredBookings.map(booking => ({
      ...booking,
      studentName: booking.studentName || '未知学生',
      studentAvatar: booking.studentAvatar || 'https://via.placeholder.com/50x50/2196F3/FFFFFF?text=学',
      studentDomain: booking.studentDomain || '未指定',
      studentGoal: booking.studentGoal || '未指定'
    }))
    
    return {
      success: true,
      data: bookingsWithStudentInfo
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