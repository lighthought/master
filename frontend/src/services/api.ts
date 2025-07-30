import { mockAuthService } from './mock/authService'

// 环境配置
const isDevelopment = import.meta.env.DEV

// API服务类
export class ApiService {
  // 认证相关API
  static auth = {
    // 用户注册
    async register(userData: {
      email: string
      password: string
      primaryIdentity: 'master' | 'apprentice'
    }) {
      if (isDevelopment) {
        return await mockAuthService.register(userData)
      }
      
      // TODO: 真实API调用
      const response = await fetch('/api/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(userData)
      })
      
      if (!response.ok) {
        throw new Error('注册失败')
      }
      
      return await response.json()
    },

    // 用户登录
    async login(email: string, password: string) {
      if (isDevelopment) {
        return await mockAuthService.login(email, password)
      }
      
      // TODO: 真实API调用
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
      })
      
      if (!response.ok) {
        throw new Error('登录失败')
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
      
      // TODO: 真实API调用
      const response = await fetch('/api/identity/master', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        },
        body: JSON.stringify({ userId, ...identityData })
      })
      
      if (!response.ok) {
        throw new Error('创建大师身份失败')
      }
      
      return await response.json()
    },

    // 创建学徒身份
    async createApprenticeIdentity(userId: string, identityData: {
      name: string
      domain: string
      learningGoals: string[]
      background: string
      preferences: string[]
    }) {
      if (isDevelopment) {
        return await mockAuthService.createApprenticeIdentity(userId, identityData)
      }
      
      // TODO: 真实API调用
      const response = await fetch('/api/identity/apprentice', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        },
        body: JSON.stringify({ userId, ...identityData })
      })
      
      if (!response.ok) {
        throw new Error('创建学徒身份失败')
      }
      
      return await response.json()
    },

    // 切换身份
    async switchIdentity(userId: string, identityId: string) {
      if (isDevelopment) {
        return await mockAuthService.switchIdentity(userId, identityId)
      }
      
      // TODO: 真实API调用
      const response = await fetch('/api/identity/switch', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        },
        body: JSON.stringify({ userId, identityId })
      })
      
      if (!response.ok) {
        throw new Error('切换身份失败')
      }
      
      return await response.json()
    },

    // 获取用户信息
    async getUserInfo(userId: string) {
      if (isDevelopment) {
        return await mockAuthService.getUserInfo(userId)
      }
      
      // TODO: 真实API调用
      const response = await fetch(`/api/user/${userId}`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        }
      })
      
      if (!response.ok) {
        throw new Error('获取用户信息失败')
      }
      
      return await response.json()
    }
  }
}