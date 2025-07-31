import { mockMasterService } from '../mock/masterService'

const isDevelopment = import.meta.env.DEV
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:3000/api'

// 学生管理API
export const studentApi = {
  // 获取学生列表
  async getStudents(params: {
    status?: string
    course_id?: string
    search?: string
    sort_by?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockMasterService.getStudents('1', params)
    }

    const searchParams = new URLSearchParams()
    if (params.status) searchParams.append('status', params.status)
    if (params.course_id) searchParams.append('course_id', params.course_id)
    if (params.search) searchParams.append('search', params.search)
    if (params.sort_by) searchParams.append('sort_by', params.sort_by)
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.page_size) searchParams.append('page_size', params.page_size.toString())

    const response = await fetch(`${API_BASE_URL}/students?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取学生列表失败')
    }

    return await response.json()
  },

  // 获取学生详情
  async getStudentDetail(studentId: string) {
    if (isDevelopment) {
      // 模拟学生详情
      const students = await mockMasterService.getStudents('1', {})
      const student = students.data?.students?.find((s: any) => s.id === studentId)
      
      return {
        code: 0,
        message: 'success',
        data: {
          student: {
            id: studentId,
            name: student?.name || '王同学',
            avatar: student?.avatar || 'https://example.com/avatar.jpg',
            email: student?.email || 'wang@example.com',
            phone: '138****1234',
            bio: '热爱编程的在校学生',
            enrollment_date: student?.enrollTime || '2024-12-01T10:00:00Z',
            status: student?.status || 'active',
            learning_goals: ['掌握Go开发', '提升编程技能'],
            preferred_learning_style: 'hands-on',
            courses: student?.enrolledCourses?.map((course: any) => ({
              course_id: course.id,
              title: course.title,
              enrollment_date: student.enrollTime,
              progress_percentage: course.progress,
              status: course.progress === 100 ? 'completed' : 'learning',
              last_study_date: student.lastStudyTime,
              total_study_time: course.studyTime,
              assignments: [
                {
                  id: 'uuid',
                  title: 'Web API开发作业',
                  status: 'submitted',
                  score: 85,
                  submitted_at: '2024-12-01T16:00:00Z'
                }
              ]
            })) || [],
            appointments: [
              {
                id: 'uuid',
                appointment_time: '2024-12-02T14:00:00Z',
                status: 'confirmed',
                topic: 'Go并发编程问题'
              }
            ],
            reviews: student?.feedback?.map((feedback: any) => ({
              id: feedback.id,
              rating: feedback.rating,
              content: feedback.content,
              created_at: feedback.createdAt
            })) || []
          }
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/students/${studentId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取学生详情失败')
    }

    return await response.json()
  },

  // 获取学生统计
  async getStudentStats() {
    if (isDevelopment) {
      return await mockMasterService.getStudentStats('1')
    }

    const response = await fetch(`${API_BASE_URL}/students/stats`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取学生统计失败')
    }

    return await response.json()
  },

  // 发送消息给学生
  async sendMessage(studentId: string, messageData: {
    content: string
    type: string
  }) {
    if (isDevelopment) {
      return await mockMasterService.sendMessage('1', { studentId, ...messageData })
    }

    const response = await fetch(`${API_BASE_URL}/students/${studentId}/messages`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(messageData)
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '发送消息失败')
    }

    return await response.json()
  },

  // 获取与学生聊天记录
  async getMessages(studentId: string, params: {
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      // 模拟聊天记录
      const students = await mockMasterService.getStudents('1', {})
      const student = students.data?.students?.find((s: any) => s.id === studentId)
      
      return {
        code: 0,
        message: 'success',
        data: {
          messages: student?.messages?.map((msg: any) => ({
            id: msg.id,
            from_user: {
              id: msg.senderId === 'master' ? '1' : studentId,
              name: msg.senderId === 'master' ? '李大师' : student?.name,
              avatar: msg.senderId === 'master' ? 'https://example.com/master-avatar.jpg' : student?.avatar
            },
            to_user: {
              id: msg.senderId === 'master' ? studentId : '1',
              name: msg.senderId === 'master' ? student?.name : '李大师',
              avatar: msg.senderId === 'master' ? student?.avatar : 'https://example.com/master-avatar.jpg'
            },
            content: msg.content,
            type: 'text',
            created_at: msg.createdAt
          })) || [],
          pagination: {
            page: params.page || 1,
            page_size: params.page_size || 50,
            total: student?.messages?.length || 0,
            total_pages: 1
          }
        },
        timestamp: new Date().toISOString()
      }
    }

    const searchParams = new URLSearchParams()
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.page_size) searchParams.append('page_size', params.page_size.toString())

    const response = await fetch(`${API_BASE_URL}/students/${studentId}/messages?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取聊天记录失败')
    }

    return await response.json()
  },

  // 更新学生学习进度
  async updateStudentProgress(studentId: string, courseId: string, progressData: {
    progress_percentage: number
    notes: string
  }) {
    if (isDevelopment) {
      // 模拟更新学习进度
      return {
        code: 0,
        message: '学习进度更新成功',
        data: {
          student_id: studentId,
          course_id: courseId,
          progress_percentage: progressData.progress_percentage
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/students/${studentId}/courses/${courseId}/progress`, {
      method: 'PUT',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(progressData)
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '更新学习进度失败')
    }

    return await response.json()
  },

  // 评价学生作业
  async gradeAssignment(studentId: string, assignmentId: string, gradeData: {
    score: number
    feedback: string
    comments: string
  }) {
    if (isDevelopment) {
      // 模拟评价作业
      return {
        code: 0,
        message: '作业评价成功',
        data: {
          assignment_id: assignmentId,
          score: gradeData.score
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/students/${studentId}/assignments/${assignmentId}/grade`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(gradeData)
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '评价作业失败')
    }

    return await response.json()
  },

  // 获取学生学习报告
  async getStudentReport(studentId: string, params: {
    period?: string
  } = {}) {
    if (isDevelopment) {
      // 模拟学习报告
      const students = await mockMasterService.getStudents('1', {})
      const student = students.data?.students?.find((s: any) => s.id === studentId)
      
      return {
        code: 0,
        message: 'success',
        data: {
          report: {
            student_id: studentId,
            student_name: student?.name || '王同学',
            period: params.period || 'month',
            study_time: student?.totalStudyTime || 45.5,
            courses_progress: student?.enrolledCourses?.map((course: any) => ({
              course_id: course.id,
              title: course.title,
              progress_percentage: course.progress,
              study_time: course.studyTime,
              assignments_completed: 3,
              average_score: 85.6
            })) || [],
            strengths: ['编程逻辑清晰', '学习态度积极'],
            areas_for_improvement: ['需要加强算法思维', '可以多练习项目实战'],
            recommendations: ['建议学习数据结构', '可以尝试开源项目']
          }
        },
        timestamp: new Date().toISOString()
      }
    }

    const searchParams = new URLSearchParams()
    if (params.period) searchParams.append('period', params.period)

    const response = await fetch(`${API_BASE_URL}/students/${studentId}/report?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取学习报告失败')
    }

    return await response.json()
  },

  // 兼容旧接口的方法
  async getStudentStatsLegacy(masterId: string) {
    return await this.getStudentStats()
  },

  async getStudentsLegacy(masterId: string, params: any = {}) {
    return await this.getStudents(params)
  },

  async sendMessageLegacy(masterId: string, messageData: any) {
    // 从messageData中提取studentId
    const studentId = messageData.studentId || messageData.to_user_id
    const content = messageData.content || messageData.message
    const type = messageData.type || 'text'
    
    return await this.sendMessage(studentId, { content, type })
  }
} 