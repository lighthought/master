import { mockAuthService } from './mock/authService'
import { mockUserStatsService } from './mock/userStatsService'
import { mockUserPreferencesService } from './mock/userPreferencesService'
import { mockMentorsService } from './mock/mentorsService'
import { mockBookingsService } from './mock/bookingsService'
import { mockCoursesService } from './mock/coursesService'
import { mockLearningService } from './mock/learningService'
import { mockCirclesService } from './mock/circlesService'
import { mockPostsService } from './mock/postsService'
import { mockLearningRecordsService } from './mock/learningRecordsService'
import { mockMasterService } from './mock/masterService'
import { mockIncomeService } from './mock/incomeService'

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
    },

    // 修改密码
    async changePassword(currentPassword: string, newPassword: string) {
      if (isDevelopment) {
        return await mockAuthService.changePassword(currentPassword, newPassword)
      }
      
      // TODO: 真实API调用
      const response = await fetch('/api/auth/change-password', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
        },
        body: JSON.stringify({ currentPassword, newPassword })
      })
      
      if (!response.ok) {
        throw new Error('修改密码失败')
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

  // 大师服务
  static mentors = {
    // 获取推荐大师
    async getRecommendedMentors(userId: string) {
      if (isDevelopment) {
        return mockMentorsService.getRecommendedMentors(userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取大师详情
    async getMentorDetail(mentorId: string) {
      if (isDevelopment) {
        return mockMentorsService.getMentorDetail(mentorId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 搜索大师
    async searchMentors(params: any) {
      if (isDevelopment) {
        return mockMentorsService.searchMentors(params)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取大师评价
    async getMentorReviews(mentorId: string, page = 1, pageSize = 10) {
      if (isDevelopment) {
        return mockMentorsService.getMentorReviews(mentorId, page, pageSize)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    }
  };

  // 预约服务
  static bookings = {
    // 创建预约
    async createBooking(bookingData: any) {
      if (isDevelopment) {
        return mockBookingsService.createBooking(bookingData)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取用户预约列表
    async getUserBookings(userId: string, status?: string) {
      if (isDevelopment) {
        return mockBookingsService.getUserBookings(userId, status)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取大师预约列表
    async getMentorBookings(mentorId: string, status?: string) {
      if (isDevelopment) {
        return mockBookingsService.getMentorBookings(mentorId, status)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 更新预约状态
    async updateBookingStatus(bookingId: string, status: string) {
      if (isDevelopment) {
        return mockBookingsService.updateBookingStatus(bookingId, status as any)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 取消预约
    async cancelBooking(bookingId: string, reason?: string) {
      if (isDevelopment) {
        return mockBookingsService.cancelBooking(bookingId, reason)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取预约详情
    async getBookingDetail(bookingId: string) {
      if (isDevelopment) {
        return mockBookingsService.getBookingDetail(bookingId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    }
  };

  // 课程服务
  static courses = {
    // 搜索课程
    async searchCourses(params: any) {
      if (isDevelopment) {
        return mockCoursesService.searchCourses(params)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取推荐课程
    async getRecommendedCourses() {
      if (isDevelopment) {
        return mockCoursesService.getRecommendedCourses()
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取课程详情
    async getCourseDetail(courseId: string) {
      if (isDevelopment) {
        return mockCoursesService.getCourseDetail(courseId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 报名课程
    async enrollCourse(enrollData: any) {
      if (isDevelopment) {
        return mockCoursesService.enrollCourse(enrollData)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },

    // 获取用户已报名课程
    async getUserEnrolledCourses(userId: string) {
      if (isDevelopment) {
        return mockCoursesService.getUserEnrolledCourses(userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    }
  };

  // 学习服务
  static learning = {
    // 获取用户学习记录
    async getUserLearningRecords(userId: string, courseId: string) {
      if (isDevelopment) {
        return mockLearningService.getUserLearningRecords(userId, courseId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 更新学习进度
    async updateLearningProgress(data: any) {
      if (isDevelopment) {
        return mockLearningService.updateLearningProgress(data)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 标记课程完成
    async markLessonCompleted(data: any) {
      if (isDevelopment) {
        return mockLearningService.markLessonCompleted(data)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 获取课程笔记
    async getCourseNotes(userId: string, courseId: string, lessonId?: string) {
      if (isDevelopment) {
        return mockLearningService.getCourseNotes(userId, courseId, lessonId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 添加笔记
    async addNote(data: any) {
      if (isDevelopment) {
        return mockLearningService.addNote(data)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 更新笔记
    async updateNote(noteId: string, content: string) {
      if (isDevelopment) {
        return mockLearningService.updateNote(noteId, content)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 删除笔记
    async deleteNote(noteId: string) {
      if (isDevelopment) {
        return mockLearningService.deleteNote(noteId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 获取课程讨论
    async getCourseDiscussions(courseId: string, lessonId?: string) {
      if (isDevelopment) {
        return mockLearningService.getCourseDiscussions(courseId, lessonId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 添加讨论
    async addDiscussion(data: any) {
      if (isDevelopment) {
        return mockLearningService.addDiscussion(data)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 点赞讨论
    async likeDiscussion(discussionId: string) {
      if (isDevelopment) {
        return mockLearningService.likeDiscussion(discussionId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 获取学习统计
    async getLearningStats(userId: string, courseId: string) {
      if (isDevelopment) {
        return mockLearningService.getLearningStats(userId, courseId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    }
  };

  // 圈子服务
  static circles = {
    // 获取圈子列表
    async getCircles(params: any) {
      if (isDevelopment) {
        return mockCirclesService.getCircles(params)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 获取圈子详情
    async getCircleDetail(circleId: string) {
      if (isDevelopment) {
        return mockCirclesService.getCircleDetail(circleId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 加入圈子
    async joinCircle(circleId: string, userId: string) {
      if (isDevelopment) {
        return mockCirclesService.joinCircle(circleId, userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 退出圈子
    async leaveCircle(circleId: string, userId: string) {
      if (isDevelopment) {
        return mockCirclesService.leaveCircle(circleId, userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 获取用户加入的圈子
    async getUserJoinedCircles(userId: string) {
      if (isDevelopment) {
        return mockCirclesService.getUserJoinedCircles(userId)
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 获取圈子分类
    async getCircleCategories() {
      if (isDevelopment) {
        return mockCirclesService.getCircleCategories()
      }
      // TODO: 实现真实API调用
      throw new Error('API not implemented')
    },
    // 搜索圈子
    async searchCircles(query: string) {
      if (isDevelopment) {
        return mockCirclesService.searchCircles(query)
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

  // 动态服务
  static posts = {
    // 获取动态列表
    async getPosts(params: any = {}) {
      if (isDevelopment) {
        return await mockPostsService.getPosts(params)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.getPosts(params)
    },
    // 发布动态
    async createPost(postData: any) {
      if (isDevelopment) {
        return await mockPostsService.createPost(postData)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.createPost(postData)
    },
    // 点赞/取消点赞
    async toggleLike(postId: string, userId: string) {
      if (isDevelopment) {
        return await mockPostsService.toggleLike(postId, userId)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.toggleLike(postId, userId)
    },
    // 添加评论
    async addComment(postId: string, commentData: any) {
      if (isDevelopment) {
        return await mockPostsService.addComment(postId, commentData)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.addComment(postId, commentData)
    },
    // 获取动态详情
    async getPostDetail(postId: string) {
      if (isDevelopment) {
        return await mockPostsService.getPostDetail(postId)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.getPostDetail(postId)
    },
    // 删除动态
    async deletePost(postId: string, userId: string) {
      if (isDevelopment) {
        return await mockPostsService.deletePost(postId, userId)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.deletePost(postId, userId)
    },
    // 获取用户的动态列表
    async getUserPosts(userId: string, params: any = {}) {
      if (isDevelopment) {
        return await mockPostsService.getUserPosts(userId, params)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.getUserPosts(userId, params)
    },
    // 评论点赞/取消点赞
    async toggleCommentLike(postId: string, commentId: string, userId: string) {
      if (isDevelopment) {
        return await mockPostsService.toggleCommentLike(postId, commentId, userId)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.toggleCommentLike(postId, commentId, userId)
    },
    // 添加回复
    async addReply(postId: string, replyData: any) {
      if (isDevelopment) {
        return await mockPostsService.addReply(postId, replyData)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.addReply(postId, replyData)
    },
    // 删除评论
    async deleteComment(postId: string, commentId: string, userId: string) {
      if (isDevelopment) {
        return await mockPostsService.deleteComment(postId, commentId, userId)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.deleteComment(postId, commentId, userId)
    },
    // 删除回复
    async deleteReply(postId: string, replyId: string, userId: string) {
      if (isDevelopment) {
        return await mockPostsService.deleteReply(postId, replyId, userId)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.deleteReply(postId, replyId, userId)
    },
    // 回复点赞/取消点赞
    async toggleReplyLike(postId: string, replyId: string, userId: string) {
      if (isDevelopment) {
        return await mockPostsService.toggleReplyLike(postId, replyId, userId)
      }
      // TODO: 实现真实API调用
      return await mockPostsService.toggleReplyLike(postId, replyId, userId)
    }
  };

  // 学习记录服务
  static learningRecords = {
    // 获取用户学习记录
    async getUserLearningRecords(userId: string, params: any = {}) {
      if (isDevelopment) {
        return await mockLearningRecordsService.getUserLearningRecords(userId, params)
      }
      // TODO: 实现真实API调用
      return await mockLearningRecordsService.getUserLearningRecords(userId, params)
    },

    // 获取学习记录详情
    async getLearningRecordDetail(recordId: string) {
      if (isDevelopment) {
        return await mockLearningRecordsService.getLearningRecordDetail(recordId)
      }
      // TODO: 实现真实API调用
      return await mockLearningRecordsService.getLearningRecordDetail(recordId)
    }
  };

  // 大师服务
  static master = {
    // 获取学生统计数据
    async getStudentStats(masterId: string) {
      if (isDevelopment) {
        return await mockMasterService.getStudentStats(masterId)
      }
      // TODO: 实现真实API调用
      return await mockMasterService.getStudentStats(masterId)
    },

    // 获取学生列表
    async getStudents(masterId: string, params: any = {}) {
      if (isDevelopment) {
        return await mockMasterService.getStudents(masterId, params)
      }
      // TODO: 实现真实API调用
      return await mockMasterService.getStudents(masterId, params)
    },

    // 发送消息
    async sendMessage(masterId: string, messageData: any) {
      if (isDevelopment) {
        return await mockMasterService.sendMessage(masterId, messageData)
      }
      // TODO: 实现真实API调用
      return await mockMasterService.sendMessage(masterId, messageData)
    }
  };

  // 收入服务
  static income = {
    // 获取收入统计数据
    async getIncomeStats(masterId: string, params: any = {}) {
      if (isDevelopment) {
        return await mockIncomeService.getIncomeStats(masterId, params)
      }
      // TODO: 实现真实API调用
      return await mockIncomeService.getIncomeStats(masterId, params)
    },

    // 获取收入明细
    async getIncomeDetails(masterId: string, params: any = {}) {
      if (isDevelopment) {
        return await mockIncomeService.getIncomeDetails(masterId, params)
      }
      // TODO: 实现真实API调用
      return await mockIncomeService.getIncomeDetails(masterId, params)
    },

    // 导出收入报表
    async exportIncomeReport(masterId: string, params: any = {}) {
      if (isDevelopment) {
        return await mockIncomeService.exportIncomeReport(masterId, params)
      }
      // TODO: 实现真实API调用
      return await mockIncomeService.exportIncomeReport(masterId, params)
    }
  };
}