import { mockAuthService } from '../mock/authService'

// 环境配置
const isDevelopment = import.meta.env.DEV
const API_BASE_URL = 'https://api.masterguide.com/v1'

// 认证相关API
export const authApi = {
  // 用户注册
  async register(userData: {
    email: string
    password: string
    phone?: string
    primaryIdentity: {
      identity_type: 'master' | 'apprentice'
      domain: string
      name: string
    }
  }) {
    if (isDevelopment) {
      // 适配mock服务的类型
      const mockData = {
        email: userData.email,
        password: userData.password,
        primaryIdentity: userData.primaryIdentity.identity_type
      }
      return await mockAuthService.register(mockData)
    }
    
    const response = await fetch(`${API_BASE_URL}/auth/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        email: userData.email,
        password: userData.password,
        phone: userData.phone,
        primary_identity: userData.primaryIdentity
      })
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '注册失败')
    }
    
    return await response.json()
  },

  // 用户登录
  async login(email: string, password: string) {
    if (isDevelopment) {
      return await mockAuthService.login(email, password)
    }
    
    const response = await fetch(`${API_BASE_URL}/auth/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ email, password })
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '登录失败')
    }
    
    return await response.json()
  },

  // 身份切换
  async switchIdentity(identityId: string) {
    if (isDevelopment) {
      return await mockAuthService.switchIdentity('', identityId)
    }
    
    const response = await fetch(`${API_BASE_URL}/auth/switch-identity`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify({ identity_id: identityId })
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '身份切换失败')
    }
    
    return await response.json()
  },

  // 刷新Token
  async refreshToken() {
    if (isDevelopment) {
      // Mock服务中没有refreshToken方法，返回模拟数据
      return {
        code: 0,
        message: 'Token刷新成功',
        data: {
          token: 'new_jwt_token'
        },
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/auth/refresh`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || 'Token刷新失败')
    }
    
    return await response.json()
  },

  // 修改密码
  async changePassword(currentPassword: string, newPassword: string) {
    if (isDevelopment) {
      return await mockAuthService.changePassword(currentPassword, newPassword)
    }
    
    const response = await fetch(`${API_BASE_URL}/auth/change-password`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify({ 
        current_password: currentPassword, 
        new_password: newPassword 
      })
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '修改密码失败')
    }
    
    return await response.json()
  },

  // 创建大师身份
  async createMasterIdentity(userId: string, identityData: {
    name: string
    domain: string
    skills: string[]
    bio: string
    experience: string
    price: number
    serviceTypes: string[]
  }) {
    if (isDevelopment) {
      return await mockAuthService.createMasterIdentity(userId, identityData)
    }
    
    const response = await fetch(`${API_BASE_URL}/users/identities`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify({
        identity_type: 'master',
        domain: identityData.domain,
        name: identityData.name,
        bio: identityData.bio,
        skills: identityData.skills,
        experience_years: parseInt(identityData.experience),
        hourly_rate: identityData.price
      })
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '创建大师身份失败')
    }
    
    return await response.json()
  },

  // 创建学徒身份
  async createApprenticeIdentity(userId: string, identityData: {
    name: string
    domain: string
    background: string
    currentLevel: string
    learningGoals: string[]
    expectedDuration: string
    learningPreferences: string[]
    timePreferences: string[]
    budgetRange: string
  }) {
    if (isDevelopment) {
      return await mockAuthService.createApprenticeIdentity(userId, identityData)
    }
    
    const response = await fetch(`${API_BASE_URL}/users/identities`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify({
        identity_type: 'apprentice',
        domain: identityData.domain,
        name: identityData.name,
        bio: identityData.background,
        skills: identityData.learningGoals,
        experience_years: 0,
        hourly_rate: 0
      })
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '创建学徒身份失败')
    }
    
    return await response.json()
  },

  // 更新身份信息
  async updateIdentityInfo(userId: string, identityId: string, identityData: any) {
    if (isDevelopment) {
      return mockAuthService.updateIdentityInfo(userId, identityId, identityData)
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

  // 获取用户信息
  async getUserInfo(userId: string) {
    if (isDevelopment) {
      return await mockAuthService.getUserInfo(userId)
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
  }
} 