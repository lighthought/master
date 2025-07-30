const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟学习记录数据
const mockLearningRecords = [
  {
    id: '1',
    userId: '1',
    courseId: '1',
    lessonId: '0-0',
    progress: 100,
    completed: true,
    completedAt: '2024-01-15T10:30:00Z',
    studyTime: 300, // 秒
    lastAccessedAt: '2024-01-15T10:30:00Z'
  },
  {
    id: '2',
    userId: '1',
    courseId: '1',
    lessonId: '0-1',
    progress: 75,
    completed: false,
    studyTime: 225,
    lastAccessedAt: '2024-01-15T11:00:00Z'
  },
  {
    id: '3',
    userId: '1',
    courseId: '1',
    lessonId: '0-2',
    progress: 0,
    completed: false,
    studyTime: 0,
    lastAccessedAt: null
  }
]

// 模拟笔记数据
const mockNotes = [
  {
    id: '1',
    userId: '1',
    courseId: '1',
    lessonId: '0-0',
    content: 'Vue.js 3.0的Composition API相比Options API更加灵活，可以更好地组织代码逻辑。',
    createdAt: '2024-01-15T10:30:00Z',
    updatedAt: '2024-01-15T10:30:00Z'
  },
  {
    id: '2',
    userId: '1',
    courseId: '1',
    lessonId: '0-1',
    content: '响应式系统的原理是通过Proxy对象来实现数据的响应式监听。',
    createdAt: '2024-01-15T11:00:00Z',
    updatedAt: '2024-01-15T11:00:00Z'
  }
]

// 模拟讨论数据
const mockDiscussions = [
  {
    id: '1',
    courseId: '1',
    lessonId: '0-0',
    userId: '1',
    userName: '学员A',
    userAvatar: 'https://via.placeholder.com/32x32/4CAF50/FFFFFF?text=A',
    title: 'Composition API的使用场景',
    content: '在什么情况下使用Composition API会比Options API更好？',
    status: 'answered',
    likeCount: 5,
    answerCount: 3,
    createdAt: '2024-01-15T10:30:00Z'
  },
  {
    id: '2',
    courseId: '1',
    lessonId: '0-1',
    userId: '2',
    userName: '学员B',
    userAvatar: 'https://via.placeholder.com/32x32/2196F3/FFFFFF?text=B',
    title: '响应式系统的性能问题',
    content: '当数据量很大时，响应式系统会不会有性能问题？',
    status: 'pending',
    likeCount: 2,
    answerCount: 0,
    createdAt: '2024-01-15T11:00:00Z'
  }
]

// 学习记录服务
export const mockLearningService = {
  // 获取用户的学习记录
  async getUserLearningRecords(userId: string, courseId: string) {
    await delay(500)
    
    const records = mockLearningRecords.filter(
      record => record.userId === userId && record.courseId === courseId
    )
    
    return {
      success: true,
      data: records,
      message: '获取学习记录成功'
    }
  },

  // 更新学习进度
  async updateLearningProgress(data: {
    userId: string
    courseId: string
    lessonId: string
    progress: number
    studyTime: number
  }) {
    await delay(300)
    
    const existingRecord = mockLearningRecords.find(
      record => record.userId === data.userId && 
                record.courseId === data.courseId && 
                record.lessonId === data.lessonId
    )
    
    if (existingRecord) {
      existingRecord.progress = data.progress
      existingRecord.studyTime = data.studyTime
      existingRecord.lastAccessedAt = new Date().toISOString()
      
      if (data.progress >= 100 && !existingRecord.completed) {
        existingRecord.completed = true
        existingRecord.completedAt = new Date().toISOString()
      }
    } else {
      mockLearningRecords.push({
        id: Date.now().toString(),
        userId: data.userId,
        courseId: data.courseId,
        lessonId: data.lessonId,
        progress: data.progress,
        completed: data.progress >= 100,
        completedAt: data.progress >= 100 ? new Date().toISOString() : undefined,
        studyTime: data.studyTime,
        lastAccessedAt: new Date().toISOString()
      })
    }
    
    return {
      success: true,
      data: { message: '学习进度更新成功' },
      message: '学习进度更新成功'
    }
  },

  // 标记课程完成
  async markLessonCompleted(data: {
    userId: string
    courseId: string
    lessonId: string
  }) {
    await delay(300)
    
    const record = mockLearningRecords.find(
      record => record.userId === data.userId && 
                record.courseId === data.courseId && 
                record.lessonId === data.lessonId
    )
    
    if (record) {
      record.completed = true
      record.progress = 100
      record.completedAt = new Date().toISOString()
    }
    
    return {
      success: true,
      data: { message: '课程标记完成成功' },
      message: '课程标记完成成功'
    }
  },

  // 获取课程笔记
  async getCourseNotes(userId: string, courseId: string, lessonId?: string) {
    await delay(400)
    
    let notes = mockNotes.filter(
      note => note.userId === userId && note.courseId === courseId
    )
    
    if (lessonId) {
      notes = notes.filter(note => note.lessonId === lessonId)
    }
    
    return {
      success: true,
      data: notes,
      message: '获取笔记成功'
    }
  },

  // 添加笔记
  async addNote(data: {
    userId: string
    courseId: string
    lessonId: string
    content: string
  }) {
    await delay(300)
    
    const newNote = {
      id: Date.now().toString(),
      userId: data.userId,
      courseId: data.courseId,
      lessonId: data.lessonId,
      content: data.content,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    }
    
    mockNotes.push(newNote)
    
    return {
      success: true,
      data: newNote,
      message: '笔记添加成功'
    }
  },

  // 更新笔记
  async updateNote(noteId: string, content: string) {
    await delay(300)
    
    const note = mockNotes.find(note => note.id === noteId)
    if (note) {
      note.content = content
      note.updatedAt = new Date().toISOString()
    }
    
    return {
      success: true,
      data: note,
      message: '笔记更新成功'
    }
  },

  // 删除笔记
  async deleteNote(noteId: string) {
    await delay(300)
    
    const index = mockNotes.findIndex(note => note.id === noteId)
    if (index > -1) {
      mockNotes.splice(index, 1)
    }
    
    return {
      success: true,
      data: { message: '笔记删除成功' },
      message: '笔记删除成功'
    }
  },

  // 获取课程讨论
  async getCourseDiscussions(courseId: string, lessonId?: string) {
    await delay(500)
    
    let discussions = mockDiscussions.filter(
      discussion => discussion.courseId === courseId
    )
    
    if (lessonId) {
      discussions = discussions.filter(discussion => discussion.lessonId === lessonId)
    }
    
    return {
      success: true,
      data: discussions,
      message: '获取讨论成功'
    }
  },

  // 添加讨论
  async addDiscussion(data: {
    courseId: string
    lessonId: string
    userId: string
    userName: string
    userAvatar: string
    title: string
    content: string
  }) {
    await delay(400)
    
    const newDiscussion = {
      id: Date.now().toString(),
      courseId: data.courseId,
      lessonId: data.lessonId,
      userId: data.userId,
      userName: data.userName,
      userAvatar: data.userAvatar,
      title: data.title,
      content: data.content,
      status: 'pending',
      likeCount: 0,
      answerCount: 0,
      createdAt: new Date().toISOString()
    }
    
    mockDiscussions.push(newDiscussion)
    
    return {
      success: true,
      data: newDiscussion,
      message: '讨论发布成功'
    }
  },

  // 点赞讨论
  async likeDiscussion(discussionId: string) {
    await delay(200)
    
    const discussion = mockDiscussions.find(d => d.id === discussionId)
    if (discussion) {
      discussion.likeCount++
    }
    
    return {
      success: true,
      data: { likeCount: discussion?.likeCount },
      message: '点赞成功'
    }
  },

  // 获取学习统计
  async getLearningStats(userId: string, courseId: string) {
    await delay(400)
    
    const records = mockLearningRecords.filter(
      record => record.userId === userId && record.courseId === courseId
    )
    
    const totalLessons = 20 // 假设总课时数
    const completedLessons = records.filter(record => record.completed).length
    const totalStudyTime = records.reduce((total, record) => total + record.studyTime, 0)
    const progress = Math.round((completedLessons / totalLessons) * 100)
    
    return {
      success: true,
      data: {
        totalLessons,
        completedLessons,
        totalStudyTime,
        progress,
                 lastStudyDate: records.length > 0 ? 
           new Date(Math.max(...records.map(r => new Date(r.lastAccessedAt || '').getTime()))).toISOString() : undefined
      },
      message: '获取学习统计成功'
    }
  }
} 