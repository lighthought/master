import { mockCirclesService } from './mock/circlesService'
import { mockPostsService } from './mock/postsService'
import { mockLearningRecordsService } from './mock/learningRecordsService'
import { mockMasterService } from './mock/masterService'
import { mockIncomeService } from './mock/incomeService'

// 环境配置
const isDevelopment = import.meta.env.DEV

// API服务类
export class ApiService {
  // 认证相关API - 从 auth.ts 模块导入
  static auth = {
    // 用户注册
    async register(userData: {
      email: string
      password: string
      phone?: string
      primaryIdentity: {
        identity_type: 'master' | 'apprentice'
        domain: string
        name: string
      }
    }) {
      const { authApi } = await import('./api/auth')
      return await authApi.register(userData)
    },

    // 用户登录
    async login(email: string, password: string) {
      const { authApi } = await import('./api/auth')
      return await authApi.login(email, password)
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
      const { authApi } = await import('./api/auth')
      return await authApi.createMasterIdentity(userId, identityData)
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
      const { authApi } = await import('./api/auth')
      return await authApi.createApprenticeIdentity(userId, identityData)
    },

    // 更新身份信息
    async updateIdentityInfo(userId: string, identityId: string, identityData: any) {
      const { authApi } = await import('./api/auth')
      return await authApi.updateIdentityInfo(userId, identityId, identityData)
    },

    // 切换身份
    async switchIdentity(userId: string, identityId: string) {
      const { authApi } = await import('./api/auth')
      return await authApi.switchIdentity(identityId)
    },

    // 获取用户信息
    async getUserInfo(userId: string) {
      const { authApi } = await import('./api/auth')
      return await authApi.getUserInfo(userId)
    },

    // 修改密码
    async changePassword(currentPassword: string, newPassword: string) {
      const { authApi } = await import('./api/auth')
      return await authApi.changePassword(currentPassword, newPassword)
    }
  };

  // 用户统计服务 - 从 user.ts 模块导入
  static userStats = {
    // 获取学习统计
    async getLearningStats(userId: string) {
      const { userApi } = await import('./api/user')
      return await userApi.getLearningStats(userId)
    },

    // 获取教学统计
    async getTeachingStats(userId: string) {
      const { userApi } = await import('./api/user')
      return await userApi.getTeachingStats(userId)
    },

    // 获取通用统计
    async getGeneralStats(userId: string) {
      const { userApi } = await import('./api/user')
      return await userApi.getGeneralStats(userId)
    },

    // 获取用户成就
    async getUserAchievements(userId: string, identityType: 'master' | 'apprentice') {
      const { userApi } = await import('./api/user')
      return await userApi.getUserAchievements(userId, identityType)
    }
  };

  // 大师服务 - 从 mentor.ts 模块导入
  static mentors = {
    // 获取推荐大师
    async getRecommendedMentors(userId: string) {
      const { mentorApi } = await import('./api/mentor')
      return await mentorApi.getRecommendedMentors(userId)
    },

    // 获取大师详情
    async getMentorDetail(mentorId: string) {
      const { mentorApi } = await import('./api/mentor')
      return await mentorApi.getMentorDetail(mentorId)
    },

    // 搜索大师
    async searchMentors(params: any) {
      const { mentorApi } = await import('./api/mentor')
      return await mentorApi.searchMentors(params)
    },

    // 获取大师评价
    async getMentorReviews(mentorId: string, page = 1, pageSize = 10) {
      const { mentorApi } = await import('./api/mentor')
      return await mentorApi.getMentorReviews(mentorId, page, pageSize)
    }
  };

  // 预约服务 - 从 appointment.ts 模块导入
  static bookings = {
    // 创建预约
    async createBooking(bookingData: any) {
      const { appointmentApi } = await import('./api/appointment')
      return await appointmentApi.createBooking(bookingData)
    },

    // 获取用户预约列表
    async getUserBookings(userId: string, status?: string) {
      const { appointmentApi } = await import('./api/appointment')
      return await appointmentApi.getUserBookings(userId, status)
    },

    // 获取大师预约列表
    async getMentorBookings(mentorId: string, status?: string) {
      const { appointmentApi } = await import('./api/appointment')
      return await appointmentApi.getMentorBookings(mentorId, status)
    },

    // 更新预约状态
    async updateBookingStatus(bookingId: string, status: string) {
      const { appointmentApi } = await import('./api/appointment')
      return await appointmentApi.updateBookingStatus(bookingId, status)
    },

    // 取消预约
    async cancelBooking(bookingId: string, reason?: string) {
      const { appointmentApi } = await import('./api/appointment')
      return await appointmentApi.cancelBooking(bookingId, reason)
    },

    // 获取预约详情
    async getBookingDetail(bookingId: string) {
      const { appointmentApi } = await import('./api/appointment')
      return await appointmentApi.getBookingDetail(bookingId)
    }
  };

  // 课程服务 - 从 course.ts 模块导入
  static courses = {
    // 搜索课程
    async searchCourses(params: any) {
      const { courseApi } = await import('./api/course')
      return await courseApi.searchCourses(params)
    },

    // 获取推荐课程
    async getRecommendedCourses() {
      const { courseApi } = await import('./api/course')
      return await courseApi.getRecommendedCourses()
    },

    // 获取课程详情
    async getCourseDetail(courseId: string) {
      const { courseApi } = await import('./api/course')
      return await courseApi.getCourseDetail(courseId)
    },

    // 报名课程
    async enrollCourse(enrollData: any) {
      const { courseApi } = await import('./api/course')
      return await courseApi.enrollCourse(enrollData.courseId, enrollData)
    },

    // 获取用户已报名课程
    async getUserEnrolledCourses(userId: string) {
      const { courseApi } = await import('./api/course')
      return await courseApi.getEnrolledCourses()
    }
  };

  // 学习服务 - 从 learning.ts 模块导入
  static learning = {
    // 获取用户学习记录
    async getUserLearningRecords(userId: string, courseId: string) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.getLearningRecords({ course_id: courseId })
    },
    // 更新学习进度
    async updateLearningProgress(data: any) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.updateLearningProgress(data.recordId, data)
    },
    // 标记课程完成
    async markLessonCompleted(data: any) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.markLessonCompleted(data)
    },
    // 获取课程笔记
    async getCourseNotes(userId: string, courseId: string, lessonId?: string) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.getCourseNotes(userId, courseId, lessonId)
    },
    // 添加笔记
    async addNote(data: any) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.addNote(data)
    },
    // 更新笔记
    async updateNote(noteId: string, content: string) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.updateNote(noteId, content)
    },
    // 删除笔记
    async deleteNote(noteId: string) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.deleteNote(noteId)
    },
    // 获取课程讨论
    async getCourseDiscussions(courseId: string, lessonId?: string) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.getCourseDiscussions(courseId, lessonId)
    },
    // 添加讨论
    async addDiscussion(data: any) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.addDiscussion(data)
    },
    // 点赞讨论
    async likeDiscussion(discussionId: string) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.likeDiscussion(discussionId)
    },
    // 获取学习统计
    async getLearningStats(userId: string, courseId: string) {
      const { learningApi } = await import('./api/learning')
      return await learningApi.getLearningStatsLegacy(userId, courseId)
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

  // 用户偏好服务 - 从 user.ts 模块导入
  static userPreferences = {
    // 获取用户偏好
    async getUserPreferences(userId: string) {
      const { userApi } = await import('./api/user')
      return await userApi.getUserPreferences(userId)
    },

    // 保存用户偏好
    async saveUserPreferences(userId: string, preferences: any) {
      const { userApi } = await import('./api/user')
      return await userApi.saveUserPreferences(userId, preferences)
    },

    // 获取推荐学习路径
    async getRecommendedLearningPath(userId: string) {
      const { userApi } = await import('./api/user')
      return await userApi.getRecommendedLearningPath()
    },

    // 获取学习路径统计
    async getLearningPathStats() {
      const { userApi } = await import('./api/user')
      return await userApi.getLearningPathStats()
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