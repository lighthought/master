// 延迟函数
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟学生统计数据
const mockStudentStats = {
  totalStudents: 24,
  activeStudents: 18,
  totalTeachingHours: 156,
  averageRating: 4.8
}

// 模拟学生列表数据
const mockStudents = [
  {
    id: '1',
    name: '张学徒',
    email: 'zhang@example.com',
    avatar: 'https://via.placeholder.com/60x60/4CAF50/FFFFFF?text=张',
    status: 'learning',
    enrollTime: '2024-01-15T10:00:00Z',
    lastStudyTime: '2024-03-20T14:30:00Z',
    totalStudyTime: 45,
    enrolledCourses: [
      {
        id: '1',
        title: 'Vue.js 基础入门',
        progress: 75,
        completedLessons: 15,
        totalLessons: 20,
        studyTime: 25
      },
      {
        id: '2',
        title: 'Vue.js 进阶实战',
        progress: 30,
        completedLessons: 6,
        totalLessons: 20,
        studyTime: 20
      }
    ],
    feedback: [
      {
        id: '1',
        courseTitle: 'Vue.js 基础入门',
        rating: 5,
        content: '老师讲解得很清楚，课程内容很实用！',
        createdAt: '2024-03-15T10:00:00Z'
      }
    ],
    messages: [
      {
        id: '1',
        content: '老师，我在第5章遇到了一些问题，能帮我看看吗？',
        senderId: 'student',
        createdAt: '2024-03-20T14:30:00Z'
      },
      {
        id: '2',
        content: '好的，请把具体问题发给我，我来帮你解决。',
        senderId: 'master',
        createdAt: '2024-03-20T15:00:00Z'
      }
    ]
  },
  {
    id: '2',
    name: '李同学',
    email: 'li@example.com',
    avatar: 'https://via.placeholder.com/60x60/2196F3/FFFFFF?text=李',
    status: 'completed',
    enrollTime: '2023-12-01T09:00:00Z',
    lastStudyTime: '2024-02-28T16:00:00Z',
    totalStudyTime: 80,
    enrolledCourses: [
      {
        id: '1',
        title: 'Vue.js 基础入门',
        progress: 100,
        completedLessons: 20,
        totalLessons: 20,
        studyTime: 40
      },
      {
        id: '3',
        title: 'Vue.js 高级特性',
        progress: 100,
        completedLessons: 15,
        totalLessons: 15,
        studyTime: 40
      }
    ],
    feedback: [
      {
        id: '2',
        courseTitle: 'Vue.js 基础入门',
        rating: 5,
        content: '课程设计得很好，从基础到进阶，循序渐进。',
        createdAt: '2024-02-25T10:00:00Z'
      },
      {
        id: '3',
        courseTitle: 'Vue.js 高级特性',
        rating: 4,
        content: '高级特性讲解得很深入，收获很大！',
        createdAt: '2024-02-28T16:00:00Z'
      }
    ],
    messages: []
  },
  {
    id: '3',
    name: '王学徒',
    email: 'wang@example.com',
    avatar: 'https://via.placeholder.com/60x60/FF9800/FFFFFF?text=王',
    status: 'paused',
    enrollTime: '2024-02-01T11:00:00Z',
    lastStudyTime: '2024-03-10T13:00:00Z',
    totalStudyTime: 15,
    enrolledCourses: [
      {
        id: '1',
        title: 'Vue.js 基础入门',
        progress: 25,
        completedLessons: 5,
        totalLessons: 20,
        studyTime: 15
      }
    ],
    feedback: [],
    messages: [
      {
        id: '3',
        content: '老师，我最近工作比较忙，可能需要暂停一段时间的学习。',
        senderId: 'student',
        createdAt: '2024-03-10T13:00:00Z'
      }
    ]
  }
]

export const mockMasterService = {
  // 获取学生统计数据
  async getStudentStats(masterId: string) {
    await delay(800)
    
    return {
      success: true,
      data: mockStudentStats
    }
  },

  // 获取学生列表
  async getStudents(masterId: string, params: any = {}) {
    await delay(1000)
    
    let filteredStudents = [...mockStudents]
    
    // 关键词搜索
    if (params.keyword) {
      const keyword = params.keyword.toLowerCase()
      filteredStudents = filteredStudents.filter(student => 
        student.name.toLowerCase().includes(keyword) ||
        student.email.toLowerCase().includes(keyword) ||
        student.enrolledCourses.some(course => 
          course.title.toLowerCase().includes(keyword)
        )
      )
    }
    
    // 状态筛选
    if (params.status) {
      filteredStudents = filteredStudents.filter(student => 
        student.status === params.status
      )
    }
    
    // 排序
    if (params.sortBy) {
      switch (params.sortBy) {
        case 'recent':
          filteredStudents.sort((a, b) => 
            new Date(b.lastStudyTime).getTime() - new Date(a.lastStudyTime).getTime()
          )
          break
        case 'progress':
          filteredStudents.sort((a, b) => {
            const aAvgProgress = a.enrolledCourses.reduce((sum, course) => sum + course.progress, 0) / a.enrolledCourses.length
            const bAvgProgress = b.enrolledCourses.reduce((sum, course) => sum + course.progress, 0) / b.enrolledCourses.length
            return bAvgProgress - aAvgProgress
          })
          break
        case 'enrollTime':
          filteredStudents.sort((a, b) => 
            new Date(b.enrollTime).getTime() - new Date(a.enrollTime).getTime()
          )
          break
      }
    }
    
    // 分页
    const page = params.page || 1
    const pageSize = params.pageSize || 10
    const start = (page - 1) * pageSize
    const end = start + pageSize
    const paginatedStudents = filteredStudents.slice(start, end)
    
    return {
      success: true,
      data: {
        students: paginatedStudents,
        total: filteredStudents.length,
        page,
        pageSize,
        totalPages: Math.ceil(filteredStudents.length / pageSize)
      }
    }
  },

  // 发送消息
  async sendMessage(masterId: string, messageData: any) {
    await delay(500)
    
    return {
      success: true,
      data: {
        id: Date.now().toString(),
        content: messageData.content,
        senderId: 'master',
        receiverId: messageData.receiverId,
        type: messageData.type,
        createdAt: new Date().toISOString()
      },
      message: '消息发送成功'
    }
  }
} 