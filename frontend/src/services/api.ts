import { mockAuthService } from './mock/authService'
import { mockUserStatsService } from './mock/userStatsService'
import { mockUserPreferencesService } from './mock/userPreferencesService'

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

    // 更新身份信息
    async updateIdentityInfo(userId: string, identityId: string, identityData: any) {
      if (isDevelopment) {
        return mockAuthService.updateIdentityInfo(userId, identityId, identityData)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
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
  };

  // 用户统计服务
  static userStats = {
    // 获取学习统计
    async getLearningStats(userId: string) {
      if (isDevelopment) {
        return mockUserStatsService.getLearningStats(userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取教学统计
    async getTeachingStats(userId: string) {
      if (isDevelopment) {
        return mockUserStatsService.getTeachingStats(userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取通用统计
    async getGeneralStats(userId: string) {
      if (isDevelopment) {
        return mockUserStatsService.getGeneralStats(userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取用户成就
    async getUserAchievements(userId: string, identityType: 'master' | 'apprentice') {
      if (isDevelopment) {
        return mockUserStatsService.getUserAchievements(userId, identityType)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    }
  };

  // 用户偏好服务
  static userPreferences = {
    // 获取用户偏好
    async getUserPreferences(userId: string) {
      if (isDevelopment) {
        return mockUserPreferencesService.getUserPreferences(userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 保存用户偏好
    async saveUserPreferences(userId: string, preferences: any) {
      if (isDevelopment) {
        return mockUserPreferencesService.saveUserPreferences(userId, preferences)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取推荐学习路径
    async getRecommendedLearningPath(userId: string) {
      if (isDevelopment) {
        return mockUserPreferencesService.getRecommendedLearningPath(userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取学习路径统计
    async getLearningPathStats() {
      if (isDevelopment) {
        return mockUserPreferencesService.getLearningPathStats()
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    }
  };
}