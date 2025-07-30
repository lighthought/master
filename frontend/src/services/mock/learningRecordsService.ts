const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟学习记录数据
const mockLearningRecords = [
  {
    id: '1',
    userId: '1',
    courseId: '1',
    courseTitle: 'Vue.js 3.0 实战开发',
    courseDescription: '从零开始学习Vue.js 3.0，掌握现代前端开发技术',
    courseCover: 'https://via.placeholder.com/300x200/4CAF50/FFFFFF?text=Vue3',
    mentorId: '2',
    mentorName: '张大师',
    mentorAvatar: 'https://via.placeholder.com/40x40/4CAF50/FFFFFF?text=张',
    status: 'learning',
    progress: 65,
    completedLessons: 13,
    totalLessons: 20,
    studyTime: 24.5,
    score: 85,
    enrollTime: new Date(Date.now() - 30 * 24 * 60 * 60 * 1000).toISOString(),
    lastStudyTime: new Date(Date.now() - 2 * 60 * 60 * 1000).toISOString(),
    learningLogs: [
      {
        id: '1',
        action: '开始学习',
        details: '开始学习第1课时：Vue.js 3.0 介绍',
        createdAt: new Date(Date.now() - 30 * 24 * 60 * 60 * 1000).toISOString()
      },
      {
        id: '2',
        action: '完成课时',
        details: '完成第1课时：Vue.js 3.0 介绍',
        createdAt: new Date(Date.now() - 29 * 24 * 60 * 60 * 1000).toISOString()
      },
      {
        id: '3',
        action: '完成课时',
        details: '完成第13课时：Composition API 实战',
        createdAt: new Date(Date.now() - 2 * 60 * 60 * 1000).toISOString()
      }
    ]
  },
  {
    id: '2',
    userId: '1',
    courseId: '2',
    courseTitle: 'React Hooks 深度解析',
    courseDescription: '深入学习React Hooks，掌握函数式组件的精髓',
    courseCover: 'https://via.placeholder.com/300x200/2196F3/FFFFFF?text=React',
    mentorId: '3',
    mentorName: '李大师',
    mentorAvatar: 'https://via.placeholder.com/40x40/2196F3/FFFFFF?text=李',
    status: 'completed',
    progress: 100,
    completedLessons: 15,
    totalLessons: 15,
    studyTime: 18.5,
    score: 92,
    enrollTime: new Date(Date.now() - 60 * 24 * 60 * 60 * 1000).toISOString(),
    lastStudyTime: new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toISOString(),
    learningLogs: [
      {
        id: '4',
        action: '开始学习',
        details: '开始学习第1课时：React Hooks 基础',
        createdAt: new Date(Date.now() - 60 * 24 * 60 * 60 * 1000).toISOString()
      },
      {
        id: '5',
        action: '完成课时',
        details: '完成第15课时：自定义Hooks实战',
        createdAt: new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toISOString()
      },
      {
        id: '6',
        action: '课程完成',
        details: '恭喜完成React Hooks 深度解析课程',
        createdAt: new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toISOString()
      }
    ]
  },
  {
    id: '3',
    userId: '1',
    courseId: '3',
    courseTitle: 'Node.js 后端开发',
    courseDescription: '学习Node.js后端开发，构建完整的Web应用',
    courseCover: 'https://via.placeholder.com/300x200/795548/FFFFFF?text=Node',
    mentorId: '4',
    mentorName: '王大师',
    mentorAvatar: 'https://via.placeholder.com/40x40/795548/FFFFFF?text=王',
    status: 'paused',
    progress: 30,
    completedLessons: 6,
    totalLessons: 20,
    studyTime: 8.5,
    score: null,
    enrollTime: new Date(Date.now() - 45 * 24 * 60 * 60 * 1000).toISOString(),
    lastStudyTime: new Date(Date.now() - 15 * 24 * 60 * 60 * 1000).toISOString(),
    learningLogs: [
      {
        id: '7',
        action: '开始学习',
        details: '开始学习第1课时：Node.js 环境搭建',
        createdAt: new Date(Date.now() - 45 * 24 * 60 * 60 * 1000).toISOString()
      },
      {
        id: '8',
        action: '完成课时',
        details: '完成第6课时：Express框架基础',
        createdAt: new Date(Date.now() - 15 * 24 * 60 * 60 * 1000).toISOString()
      }
    ]
  }
]

export const mockLearningRecordsService = {
  // 获取用户学习记录
  async getUserLearningRecords(userId: string, params: any = {}) {
    await delay(800)
    
    let records = [...mockLearningRecords].filter(record => record.userId === userId)
    
    // 根据状态过滤
    if (params.status) {
      records = records.filter(record => record.status === params.status)
    }
    
    // 根据关键词搜索
    if (params.keyword) {
      const keyword = params.keyword.toLowerCase()
      records = records.filter(record => 
        record.courseTitle.toLowerCase().includes(keyword) ||
        record.mentorName.toLowerCase().includes(keyword) ||
        record.courseDescription.toLowerCase().includes(keyword)
      )
    }
    
    // 排序
    if (params.sortBy === 'recent') {
      records.sort((a, b) => new Date(b.lastStudyTime).getTime() - new Date(a.lastStudyTime).getTime())
    } else if (params.sortBy === 'progress') {
      records.sort((a, b) => b.progress - a.progress)
    } else if (params.sortBy === 'enrollTime') {
      records.sort((a, b) => new Date(b.enrollTime).getTime() - new Date(a.enrollTime).getTime())
    }
    
    // 分页
    const page = params.page || 1
    const pageSize = params.pageSize || 10
    const start = (page - 1) * pageSize
    const end = start + pageSize
    
    return {
      success: true,
      data: {
        records: records.slice(start, end),
        total: records.length,
        page,
        pageSize
      }
    }
  },

  // 获取学习记录详情
  async getLearningRecordDetail(recordId: string) {
    await delay(500)
    
    const record = mockLearningRecords.find(r => r.id === recordId)
    if (!record) {
      throw new Error('学习记录不存在')
    }
    
    return {
      success: true,
      data: record
    }
  }
} 