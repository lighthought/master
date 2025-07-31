import type { User } from '@/types/user'

// 模拟延迟
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟用户偏好数据
const mockUserPreferences: { [key: string]: {
  learningStyle: string
  timePreference: string
  budgetRange: string
  learningGoals: string[]
  preferredDomains: string[]
  experienceLevel: string
  updatedAt: string
} } = {
  '1': {
    learningStyle: 'one-on-one',
    timePreference: 'flexible',
    budgetRange: 'medium',
    learningGoals: ['掌握前端开发', '提升编程技能'],
    preferredDomains: ['软件开发', '前端开发'],
    experienceLevel: 'beginner',
    updatedAt: '2024-01-15T10:30:00Z'
  }
}

// 推荐算法配置
const recommendationWeights = {
  learningStyle: 0.4,
  timePreference: 0.2,
  budgetRange: 0.2,
  experienceLevel: 0.1,
  learningGoals: 0.1
}

export const mockUserPreferencesService = {
  // 获取用户偏好
  async getUserPreferences(userId: string) {
    await delay(500)
    
    const preferences = mockUserPreferences[userId]
    if (!preferences) {
      // 返回默认偏好
      return {
        success: true,
        data: {
          learningStyle: 'one-on-one',
          timePreference: 'flexible',
          budgetRange: 'medium',
          learningGoals: [],
          preferredDomains: [],
          experienceLevel: 'beginner',
          updatedAt: new Date().toISOString()
        }
      }
    }
    
    return {
      success: true,
      data: preferences
    }
  },

  // 保存用户偏好
  async saveUserPreferences(userId: string, preferences: any) {
    await delay(800)
    
    // 更新用户偏好
    mockUserPreferences[userId] = {
      ...preferences,
      updatedAt: new Date().toISOString()
    }
    
    return {
      success: true,
      message: '用户偏好保存成功'
    }
  },

  // 获取推荐学习路径
  async getRecommendedLearningPath(userId: string) {
    await delay(600)
    
    const preferences = mockUserPreferences[userId]
    if (!preferences) {
      return {
        success: true,
        data: {
          recommendedPath: 'one-on-one',
          confidence: 0.7,
          reasons: ['基于默认偏好推荐', '适合初学者']
        }
      }
    }
    
    // 简单的推荐算法
    let recommendedPath = 'one-on-one'
    let confidence = 0.7
    const reasons = []
    
    // 根据学习方式偏好
    if (preferences.learningStyle === 'structured') {
      recommendedPath = 'structured'
      confidence = 0.8
      reasons.push('用户偏好结构化学习')
    } else if (preferences.learningStyle === 'browse') {
      recommendedPath = 'browse'
      confidence = 0.75
      reasons.push('用户希望浏览更多选择')
    } else if (preferences.learningStyle === 'other') {
      recommendedPath = 'other'
      confidence = 0.6
      reasons.push('用户希望探索其他方式')
    }
    
    // 根据经验水平调整
    if (preferences.experienceLevel === 'beginner') {
      reasons.push('适合初学者')
    } else if (preferences.experienceLevel === 'intermediate') {
      reasons.push('适合中级学习者')
    } else if (preferences.experienceLevel === 'advanced') {
      reasons.push('适合高级学习者')
    }
    
    // 根据预算调整
    if (preferences.budgetRange === 'low') {
      reasons.push('经济型预算推荐')
    } else if (preferences.budgetRange === 'high') {
      reasons.push('高端服务推荐')
    }
    
    return {
      success: true,
      data: {
        recommendedPath,
        confidence,
        reasons
      }
    }
  },

  // 获取学习路径统计
  async getLearningPathStats() {
    await delay(400)
    
    return {
      success: true,
      data: {
        totalUsers: 1250,
        pathDistribution: {
          'one-on-one': 45,
          'structured': 30,
          'browse': 20,
          'other': 5
        },
        satisfactionRates: {
          'one-on-one': 4.8,
          'structured': 4.6,
          'browse': 4.4,
          'other': 4.2
        }
      }
    }
  }
}