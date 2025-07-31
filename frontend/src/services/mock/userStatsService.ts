import type { User } from '@/types/user'

// 模拟延迟
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟用户统计数据
const mockUserStats: { [key: string]: {
  apprentice: {
    learningStats: {
      totalCourses: number
      progress: number
      completedLessons: number
      totalLessons: number
      currentCourse: string
      nextLesson: string
    }
    achievements: Array<{
      id: string
      name: string
      description: string
      icon: string
    }>
  }
  master: {
    teachingStats: {
      totalStudents: number
      totalHours: number
      totalEarnings: number
      averageRating: number
      completedSessions: number
      upcomingSessions: number
    }
    achievements: Array<{
      id: string
      name: string
      description: string
      icon: string
    }>
  }
  general: {
    activeDays: number
    achievements: number
    totalLoginDays: number
    lastLoginDate: string
    streakDays: number
  }
} } = {
  '1': {
    // 学徒统计数据
    apprentice: {
      learningStats: {
        totalCourses: 12,
        progress: 65,
        completedLessons: 8,
        totalLessons: 15,
        currentCourse: 'Vue.js 进阶开发',
        nextLesson: '组件通信与状态管理'
      },
      achievements: [
        { id: '1', name: '学习新手', description: '完成第一门课程', icon: '🎓' },
        { id: '2', name: '坚持不懈', description: '连续学习7天', icon: '🔥' },
        { id: '3', name: '技能提升', description: '掌握5个技能标签', icon: '⭐' }
      ]
    },
    // 大师统计数据
    master: {
      teachingStats: {
        totalStudents: 8,
        totalHours: 24,
        totalEarnings: 2400,
        averageRating: 4.8,
        completedSessions: 12,
        upcomingSessions: 3
      },
      achievements: [
        { id: '1', name: '指导新手', description: '指导第一个学员', icon: '👨‍🏫' },
        { id: '2', name: '好评如潮', description: '获得5星评价', icon: '🌟' },
        { id: '3', name: '收入达人', description: '累计收入超过1000元', icon: '💰' }
      ]
    },
    // 通用统计数据
    general: {
      activeDays: 7,
      achievements: 3,
      totalLoginDays: 15,
      lastLoginDate: '2024-01-15',
      streakDays: 5
    }
  }
}

export const mockUserStatsService = {
  // 获取用户学习统计（学徒）
  async getLearningStats(userId: string) {
    await delay(500)
    
    const userStats = mockUserStats[userId]
    if (!userStats) {
      throw new Error('用户统计数据不存在')
    }
    
    return {
      success: true,
      data: userStats.apprentice.learningStats
    }
  },

  // 获取用户教学统计（大师）
  async getTeachingStats(userId: string) {
    await delay(500)
    
    const userStats = mockUserStats[userId]
    if (!userStats) {
      throw new Error('用户统计数据不存在')
    }
    
    return {
      success: true,
      data: userStats.master.teachingStats
    }
  },

  // 获取用户通用统计
  async getGeneralStats(userId: string) {
    await delay(500)
    
    const userStats = mockUserStats[userId]
    if (!userStats) {
      throw new Error('用户统计数据不存在')
    }
    
    return {
      success: true,
      data: userStats.general
    }
  },

  // 获取用户成就列表
  async getUserAchievements(userId: string, identityType: 'master' | 'apprentice') {
    await delay(500)
    
    const userStats = mockUserStats[userId]
    if (!userStats) {
      throw new Error('用户统计数据不存在')
    }
    
    return {
      success: true,
      data: userStats[identityType].achievements
    }
  },

  // 更新用户统计数据
  async updateUserStats(userId: string, statsType: string, data: any) {
    await delay(300)
    
    const userStats = mockUserStats[userId]
    if (!userStats) {
      throw new Error('用户统计数据不存在')
    }
    
    // 模拟更新统计数据
    if (statsType === 'learning') {
      Object.assign(userStats.apprentice.learningStats, data)
    } else if (statsType === 'teaching') {
      Object.assign(userStats.master.teachingStats, data)
    } else if (statsType === 'general') {
      Object.assign(userStats.general, data)
    }
    
    return {
      success: true,
      message: '统计数据更新成功'
    }
  }
}