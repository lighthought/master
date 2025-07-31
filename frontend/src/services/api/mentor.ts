import { mockMentorsService } from '../mock/mentorsService'

// 环境配置
const isDevelopment = import.meta.env.DEV
const API_BASE_URL = 'https://api.masterguide.com/v1'

// 大师管理API
export const mentorApi = {
  // 获取大师列表
  async getMentors(params: {
    domain?: string
    min_rating?: number
    max_price?: number
    is_online?: boolean
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockMentorsService.searchMentors(params)
    }
    
    const searchParams = new URLSearchParams()
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        searchParams.append(key, value.toString())
      }
    })
    
    const response = await fetch(`${API_BASE_URL}/mentors?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取大师列表失败')
    }
    
    return await response.json()
  },

  // 获取大师详情
  async getMentorDetail(mentorId: string) {
    if (isDevelopment) {
      return await mockMentorsService.getMentorDetail(mentorId)
    }
    
    const response = await fetch(`${API_BASE_URL}/mentors/${mentorId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取大师详情失败')
    }
    
    return await response.json()
  },

  // 搜索大师
  async searchMentors(params: {
    q?: string
    domain?: string
    min_rating?: number
    max_price?: number
    is_online?: boolean
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockMentorsService.searchMentors(params)
    }
    
    const searchParams = new URLSearchParams()
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        searchParams.append(key, value.toString())
      }
    })
    
    const response = await fetch(`${API_BASE_URL}/mentors/search?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '搜索大师失败')
    }
    
    return await response.json()
  },

  // 获取推荐大师
  async getRecommendedMentors(userId?: string) {
    if (isDevelopment) {
      return await mockMentorsService.getRecommendedMentors(userId || '')
    }
    
    const params = userId ? `?user_id=${userId}` : ''
    const response = await fetch(`${API_BASE_URL}/mentors/recommended${params}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取推荐大师失败')
    }
    
    return await response.json()
  },

  // 获取大师评价
  async getMentorReviews(mentorId: string, page = 1, pageSize = 10) {
    if (isDevelopment) {
      return await mockMentorsService.getMentorReviews(mentorId, page, pageSize)
    }
    
    const response = await fetch(`${API_BASE_URL}/mentors/${mentorId}/reviews?page=${page}&page_size=${pageSize}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取大师评价失败')
    }
    
    return await response.json()
  }
} 