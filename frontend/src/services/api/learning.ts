import { mockLearningService } from '../mock/learningService'
import { mockLearningRecordsService } from '../mock/learningRecordsService'

// 环境配置
const isDevelopment = import.meta.env.DEV
const API_BASE_URL = 'https://api.masterguide.com/v1'

// 学习记录管理API
export const learningApi = {
  // 获取学习记录列表
  async getLearningRecords(params: {
    course_id?: string
    status?: string
    start_date?: string
    end_date?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockLearningRecordsService.getUserLearningRecords('1', params)
    }
    
    const searchParams = new URLSearchParams()
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        searchParams.append(key, value.toString())
      }
    })
    
    const response = await fetch(`${API_BASE_URL}/learning-records?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取学习记录列表失败')
    }
    
    return await response.json()
  },

  // 获取学习记录详情
  async getLearningRecordDetail(recordId: string) {
    if (isDevelopment) {
      return await mockLearningRecordsService.getLearningRecordDetail(recordId)
    }
    
    const response = await fetch(`${API_BASE_URL}/learning-records/${recordId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取学习记录详情失败')
    }
    
    return await response.json()
  },

  // 更新学习进度
  async updateLearningProgress(recordId: string, progressData: {
    progress_percentage: number
    current_chapter: string
    study_time_minutes: number
  }) {
    if (isDevelopment) {
      return await mockLearningService.updateLearningProgress({
        userId: '1',
        courseId: '1',
        lessonId: progressData.current_chapter,
        progress: progressData.progress_percentage,
        studyTime: progressData.study_time_minutes
      })
    }
    
    const response = await fetch(`${API_BASE_URL}/learning-records/${recordId}/progress`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(progressData)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '更新学习进度失败')
    }
    
    return await response.json()
  },

  // 提交作业
  async submitAssignment(recordId: string, assignmentData: {
    title: string
    content: string
    attachment_urls?: string[]
  }) {
    if (isDevelopment) {
      return {
        code: 0,
        message: '作业提交成功',
        data: {
          assignment_id: 'uuid'
        },
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/learning-records/${recordId}/assignments`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(assignmentData)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '提交作业失败')
    }
    
    return await response.json()
  },

  // 获取学习统计
  async getLearningStats(params: {
    period?: string
  } = {}) {
    if (isDevelopment) {
      return await mockLearningService.getLearningStats('', '')
    }
    
    const searchParams = new URLSearchParams()
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        searchParams.append(key, value.toString())
      }
    })
    
    const response = await fetch(`${API_BASE_URL}/learning-records/stats?${searchParams}`, {
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

  // 获取学习路径推荐
  async getRecommendedLearningPath() {
    if (isDevelopment) {
      return {
        code: 0,
        message: 'success',
        data: {
          path: {
            current_level: 'intermediate',
            next_courses: [
              {
                id: 'uuid',
                title: 'Go微服务架构',
                reason: '基于您当前的学习进度推荐',
                estimated_duration: 25
              }
            ],
            skills_to_develop: ['微服务设计', '服务治理', '容器化部署'],
            estimated_completion_time: '3个月'
          }
        },
        timestamp: new Date().toISOString()
      }
    }
    
    const response = await fetch(`${API_BASE_URL}/learning-records/recommended-path`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取学习路径推荐失败')
    }
    
    return await response.json()
  },

  // 兼容旧接口的方法
  // 获取用户学习记录 (兼容旧接口)
  async getUserLearningRecords(userId: string, courseId: string) {
    const params: any = {}
    if (courseId) params.course_id = courseId
    return await this.getLearningRecords(params)
  },

  // 更新学习进度 (兼容旧接口)
  async updateLearningProgressLegacy(data: any) {
    const { recordId, ...progressData } = data
    return await this.updateLearningProgress(recordId, progressData)
  },

  // 标记课程完成 (兼容旧接口)
  async markLessonCompleted(data: any) {
    // 标记课程完成可以视为更新进度到100%
    const { recordId, ...progressData } = data
    return await this.updateLearningProgress(recordId, {
      ...progressData,
      progress_percentage: 100
    })
  },

  // 获取课程笔记 (兼容旧接口)
  async getCourseNotes(userId: string, courseId: string, lessonId?: string) {
    if (isDevelopment) {
      return await mockLearningService.getCourseNotes(userId, courseId, lessonId)
    }
    
    const params = new URLSearchParams()
    if (courseId) params.append('course_id', courseId)
    if (lessonId) params.append('lesson_id', lessonId)
    
    const response = await fetch(`${API_BASE_URL}/learning-records/notes?${params}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取课程笔记失败')
    }
    
    return await response.json()
  },

  // 添加笔记 (兼容旧接口)
  async addNote(data: any) {
    if (isDevelopment) {
      return await mockLearningService.addNote(data)
    }
    
    const response = await fetch(`${API_BASE_URL}/learning-records/notes`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(data)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '添加笔记失败')
    }
    
    return await response.json()
  },

  // 更新笔记 (兼容旧接口)
  async updateNote(noteId: string, content: string) {
    if (isDevelopment) {
      return await mockLearningService.updateNote(noteId, content)
    }
    
    const response = await fetch(`${API_BASE_URL}/learning-records/notes/${noteId}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify({ content })
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '更新笔记失败')
    }
    
    return await response.json()
  },

  // 删除笔记 (兼容旧接口)
  async deleteNote(noteId: string) {
    if (isDevelopment) {
      return await mockLearningService.deleteNote(noteId)
    }
    
    const response = await fetch(`${API_BASE_URL}/learning-records/notes/${noteId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '删除笔记失败')
    }
    
    return await response.json()
  },

  // 获取课程讨论 (兼容旧接口)
  async getCourseDiscussions(courseId: string, lessonId?: string) {
    if (isDevelopment) {
      return await mockLearningService.getCourseDiscussions(courseId, lessonId)
    }
    
    const params = new URLSearchParams()
    if (courseId) params.append('course_id', courseId)
    if (lessonId) params.append('lesson_id', lessonId)
    
    const response = await fetch(`${API_BASE_URL}/learning-records/discussions?${params}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '获取课程讨论失败')
    }
    
    return await response.json()
  },

  // 添加讨论 (兼容旧接口)
  async addDiscussion(data: any) {
    if (isDevelopment) {
      return await mockLearningService.addDiscussion(data)
    }
    
    const response = await fetch(`${API_BASE_URL}/learning-records/discussions`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      },
      body: JSON.stringify(data)
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '添加讨论失败')
    }
    
    return await response.json()
  },

  // 点赞讨论 (兼容旧接口)
  async likeDiscussion(discussionId: string) {
    if (isDevelopment) {
      return await mockLearningService.likeDiscussion(discussionId)
    }
    
    const response = await fetch(`${API_BASE_URL}/learning-records/discussions/${discussionId}/like`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || '点赞讨论失败')
    }
    
    return await response.json()
  },

  // 获取学习统计 (兼容旧接口)
  async getLearningStatsLegacy(userId: string, courseId: string) {
    const params: any = {}
    if (courseId) params.course_id = courseId
    return await this.getLearningStats(params)
  }
} 