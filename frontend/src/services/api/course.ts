import { mockCoursesService } from '../mock/coursesService'

// 环境配置
const isDevelopment = import.meta.env.DEV
const API_BASE_URL = 'https://api.masterguide.com/v1'

// 课程管理API
export const courseApi = {
  // 获取课程列表
  async getCourses(params: {
    domain?: string
    difficulty?: string
    min_price?: number
    max_price?: number
    sort_by?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockCoursesService.searchCourses(params)
    }
    
    const searchParams = new URLSearchParams()
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        searchParams.append(key, value.toString())
      }
    })
    
    const response = await fetch(`${API_BASE_URL}/courses?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取课程列表失败')
    }
    
    return await response.json()
  },

  // 获取课程详情
  async getCourseDetail(courseId: string) {
    if (isDevelopment) {
      return await mockCoursesService.getCourseDetail(courseId)
    }
    
    const response = await fetch(`${API_BASE_URL}/courses/${courseId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取课程详情失败')
    }
    
    return await response.json()
  },

  // 创建课程（大师身份）
  async createCourse(courseData: {
    title: string
    description: string
    cover_image: string
    price: number
    duration_hours: number
    difficulty: string
    max_students: number
    contents: Array<{
      title: string
      content_type: string
      content_url: string
      duration_minutes: number
      order_index: number
    }>
  }) {
    if (isDevelopment) {
      return {
        code: 0,
        message: '课程创建成功',
        data: {
          course_id: 'uuid'
        },
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/courses`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(courseData)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '创建课程失败')
    }
    
    return await response.json()
  },

  // 报名课程
  async enrollCourse(courseId: string, enrollData: {
    payment_method: string
    user_info: {
      name: string
      phone: string
    }
  }) {
    if (isDevelopment) {
      // 适配mock服务的类型
      const mockData = {
        courseId,
        userId: '1', // 从当前用户获取
        price: 299, // 从课程获取
        paymentMethod: enrollData.payment_method,
        userInfo: enrollData.user_info
      }
      return await mockCoursesService.enrollCourse(mockData)
    }
    
    const response = await fetch(`${API_BASE_URL}/courses/${courseId}/enroll`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(enrollData)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '报名课程失败')
    }
    
    return await response.json()
  },

  // 获取学习进度
  async getCourseProgress(courseId: string) {
    if (isDevelopment) {
      return {
        code: 0,
        message: 'success',
        data: {
          course_id: courseId,
          progress_percentage: 65.5,
          status: 'learning',
          enrolled_at: '2024-12-01T10:00:00Z',
          last_accessed_at: '2024-12-01T15:30:00Z',
          completed_contents: ['uuid1', 'uuid2']
        },
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/courses/${courseId}/progress`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取学习进度失败')
    }
    
    return await response.json()
  },

  // 搜索课程
  async searchCourses(params: {
    q?: string
    domain?: string
    difficulty?: string
    min_price?: number
    max_price?: number
    sort_by?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockCoursesService.searchCourses(params)
    }
    
    const searchParams = new URLSearchParams()
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        searchParams.append(key, value.toString())
      }
    })
    
    const response = await fetch(`${API_BASE_URL}/courses/search?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '搜索课程失败')
    }
    
    return await response.json()
  },

  // 获取推荐课程
  async getRecommendedCourses(userId?: string) {
    if (isDevelopment) {
      return await mockCoursesService.getRecommendedCourses()
    }
    
    const params = userId ? `?user_id=${userId}` : ''
    const response = await fetch(`${API_BASE_URL}/courses/recommended${params}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取推荐课程失败')
    }
    
    return await response.json()
  },

  // 获取用户已报名课程
  async getEnrolledCourses(params: {
    status?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockCoursesService.getUserEnrolledCourses('')
    }
    
    const searchParams = new URLSearchParams()
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        searchParams.append(key, value.toString())
      }
    })
    
    const response = await fetch(`${API_BASE_URL}/courses/enrolled?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取已报名课程失败')
    }
    
    return await response.json()
  }
} 