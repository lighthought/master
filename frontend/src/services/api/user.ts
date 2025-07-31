import { mockUserStatsService } from '../mock/userStatsService'
import { mockUserPreferencesService } from '../mock/userPreferencesService'

// 环境配置
const isDevelopment = import.meta.env.DEV
const API_BASE_URL = 'https://api.masterguide.com/v1'

// 用户管理API
export const userApi = {
  // 获取用户信息
  async getProfile() {
    if (isDevelopment) {
      return {
        code: 0,
        message: 'success',
        data: {
          user: {
            id: 'uuid',
            email: 'user@example.com',
            phone: '13800138000',
            status: 'active',
            created_at: '2024-12-01T10:00:00Z'
          },
          current_identity: {
            id: 'uuid',
            identity_type: 'apprentice',
            domain: 'software_development',
            status: 'active',
            profile: {
              name: '张三',
              avatar: 'https://example.com/avatar.jpg',
              bio: '热爱学习的新手',
              skills: ['JavaScript', 'Vue.js'],
              experience_years: 1
            }
          }
        },
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/users/profile`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取用户信息失败')
    }
    
    return await response.json()
  },

  // 更新用户档案
  async updateProfile(profileData: {
    name: string
    avatar?: string
    bio?: string
    skills?: string[]
    experience_years?: number
    hourly_rate?: number
  }) {
    if (isDevelopment) {
      return {
        code: 0,
        message: '用户档案更新成功',
        data: {},
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/users/profile`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(profileData)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '更新用户档案失败')
    }
    
    return await response.json()
  },

  // 获取用户身份列表
  async getIdentities() {
    if (isDevelopment) {
      return {
        code: 0,
        message: 'success',
        data: {
          identities: [
            {
              id: 'uuid',
              identity_type: 'apprentice',
              domain: 'software_development',
              status: 'active',
              profile: {
                name: '张三',
                avatar: 'https://example.com/avatar.jpg'
              }
            }
          ]
        },
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/users/identities`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取身份列表失败')
    }
    
    return await response.json()
  },

  // 添加新身份
  async addIdentity(identityData: {
    identity_type: 'master' | 'apprentice'
    domain: string
    name: string
    bio?: string
    skills?: string[]
    experience_years?: number
    hourly_rate?: number
  }) {
    if (isDevelopment) {
      return {
        code: 0,
        message: '身份创建成功',
        data: {
          identity_id: 'uuid',
          status: 'pending'
        },
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/users/identities`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(identityData)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '创建身份失败')
    }
    
    return await response.json()
  },

  // 更新身份信息
  async updateIdentity(identityId: string, identityData: {
    name?: string
    bio?: string
    skills?: string[]
    experience_years?: number
    hourly_rate?: number
  }) {
    if (isDevelopment) {
      return {
        code: 0,
        message: '身份信息更新成功',
        data: {},
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/users/identities/${identityId}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(identityData)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '更新身份信息失败')
    }
    
    return await response.json()
  },

  // 获取用户学习统计
  async getLearningStats(userId?: string) {
    if (isDevelopment) {
      return await mockUserStatsService.getLearningStats(userId || '1')
    }
    
    const response = await fetch(`${API_BASE_URL}/users/stats/learning`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取学习统计失败')
    }
    
    return await response.json()
  },

  // 获取用户教学统计
  async getTeachingStats(userId?: string) {
    if (isDevelopment) {
      return await mockUserStatsService.getTeachingStats(userId || '1')
    }
    
    const response = await fetch(`${API_BASE_URL}/users/stats/teaching`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取教学统计失败')
    }
    
    return await response.json()
  },

  // 获取用户通用统计
  async getGeneralStats(userId ?: string) {
    if (isDevelopment) {
      return await mockUserStatsService.getGeneralStats(userId || '1')
    }
    
    const response = await fetch(`${API_BASE_URL}/users/stats/general`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取通用统计失败')
    }
    
    return await response.json()
  },

  // 获取用户成就列表
  async getAchievements(identityType?: 'master' | 'apprentice') {
    if (isDevelopment) {
      return await mockUserStatsService.getUserAchievements('', identityType || 'apprentice')
    }
    
    const params = identityType ? `?identity_type=${identityType}` : ''
    const response = await fetch(`${API_BASE_URL}/users/achievements${params}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取成就列表失败')
    }
    
    return await response.json()
  },

  // 获取用户偏好
  async getPreferences() {
    if (isDevelopment) {
      return await mockUserPreferencesService.getUserPreferences('')
    }
    
    const response = await fetch(`${API_BASE_URL}/users/preferences`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取用户偏好失败')
    }
    
    return await response.json()
  },

  // 保存用户偏好
  async savePreferences(preferences: {
    learning_style?: string
    time_preference?: string
    budget_range?: string
    learning_goals?: string[]
    preferred_domains?: string[]
    experience_level?: string
  }) {
    if (isDevelopment) {
      return await mockUserPreferencesService.saveUserPreferences('', preferences)
    }
    
    const response = await fetch(`${API_BASE_URL}/users/preferences`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(preferences)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '保存用户偏好失败')
    }
    
    return await response.json()
  },

  // 获取推荐学习路径
  async getRecommendedLearningPath() {
    if (isDevelopment) {
      return await mockUserPreferencesService.getRecommendedLearningPath('1')
    }
    
    const response = await fetch(`${API_BASE_URL}/users/recommended-learning-path`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取推荐学习路径失败')
    }
    
    return await response.json()
  },

  // 获取学习路径统计
  async getLearningPathStats() {
    if (isDevelopment) {
      return await mockUserPreferencesService.getLearningPathStats()
    }
    
    const response = await fetch(`${API_BASE_URL}/users/learning-path-stats`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取学习路径统计失败')
    }
    
    return await response.json()
  },

  // 获取用户成就列表 (兼容旧接口)
  async getUserAchievements(userId?: string, identityType?: 'master' | 'apprentice') {
    if (isDevelopment) {
      return await mockUserStatsService.getUserAchievements(userId || '', identityType || 'apprentice')
    }
    
    const params = identityType ? `?identity_type=${identityType}` : ''
    const response = await fetch(`${API_BASE_URL}/users/achievements${params}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取成就列表失败')
    }
    
    return await response.json()
  },

  // 获取用户偏好 (兼容旧接口)
  async getUserPreferences(userId?: string) {
    if (isDevelopment) {
      return await mockUserPreferencesService.getUserPreferences(userId || '')
    }
    
    const response = await fetch(`${API_BASE_URL}/users/preferences`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取用户偏好失败')
    }
    
    return await response.json()
  },

  // 保存用户偏好 (兼容旧接口)
  async saveUserPreferences(userId?: string, preferences?: any) {
    if (isDevelopment) {
      return await mockUserPreferencesService.saveUserPreferences(userId || '', preferences || {})
    }
    
    const response = await fetch(`${API_BASE_URL}/users/preferences`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(preferences || {})
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '保存用户偏好失败')
    }
    
    return await response.json()
  }
} 