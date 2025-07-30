// 模拟延迟
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟课程数据
const mockCourses = [
  {
    id: '1',
    title: 'Vue.js 3.0 从入门到精通',
    description: '全面学习Vue.js 3.0的新特性和最佳实践，包括Composition API、响应式系统、组件开发等核心内容。',
    cover: 'https://via.placeholder.com/400x250/4CAF50/FFFFFF?text=Vue.js+3.0',
    mentorId: '1',
    mentorName: '张大师',
    mentorAvatar: 'https://via.placeholder.com/50x50/4CAF50/FFFFFF?text=张',
    category: 'frontend',
    tags: ['Vue.js', '前端开发', 'JavaScript'],
    price: 299,
    originalPrice: 399,
    duration: 15,
    level: 'intermediate',
    rating: 4.8,
    studentCount: 1250,
    isNew: true,
    isHot: true,
    isRecommended: true,
    createdAt: '2024-01-10T10:00:00Z',
    updatedAt: '2024-01-15T14:30:00Z'
  },
  {
    id: '2',
    title: 'Spring Boot 微服务架构实战',
    description: '深入学习Spring Boot微服务架构设计，包括服务拆分、API网关、配置中心、服务治理等核心概念。',
    cover: 'https://via.placeholder.com/400x250/2196F3/FFFFFF?text=Spring+Boot',
    mentorId: '2',
    mentorName: '李导师',
    mentorAvatar: 'https://via.placeholder.com/50x50/2196F3/FFFFFF?text=李',
    category: 'backend',
    tags: ['Spring Boot', 'Java', '微服务'],
    price: 399,
    originalPrice: 499,
    duration: 20,
    level: 'advanced',
    rating: 4.9,
    studentCount: 890,
    isNew: false,
    isHot: true,
    isRecommended: true,
    createdAt: '2024-01-05T09:00:00Z',
    updatedAt: '2024-01-12T16:20:00Z'
  },
  {
    id: '3',
    title: 'Flutter 跨平台开发实战',
    description: '掌握Flutter跨平台开发技能，学习Widget体系、状态管理、网络请求、本地存储等核心知识。',
    cover: 'https://via.placeholder.com/400x250/FF9800/FFFFFF?text=Flutter',
    mentorId: '3',
    mentorName: '王老师',
    mentorAvatar: 'https://via.placeholder.com/50x50/FF9800/FFFFFF?text=王',
    category: 'mobile',
    tags: ['Flutter', '移动开发', 'Dart'],
    price: 349,
    originalPrice: 449,
    duration: 18,
    level: 'intermediate',
    rating: 4.7,
    studentCount: 756,
    isNew: true,
    isHot: false,
    isRecommended: true,
    createdAt: '2024-01-08T11:30:00Z',
    updatedAt: '2024-01-14T10:15:00Z'
  },
  {
    id: '4',
    title: 'Python 数据分析与可视化',
    description: '学习Python数据分析的核心技能，包括pandas、numpy、matplotlib等库的使用，掌握数据清洗、分析和可视化。',
    cover: 'https://via.placeholder.com/400x250/9C27B0/FFFFFF?text=Python+Data',
    mentorId: '4',
    mentorName: '陈工程师',
    mentorAvatar: 'https://via.placeholder.com/50x50/9C27B0/FFFFFF?text=陈',
    category: 'data',
    tags: ['Python', '数据分析', '可视化'],
    price: 279,
    originalPrice: 379,
    duration: 12,
    level: 'beginner',
    rating: 4.6,
    studentCount: 1120,
    isNew: false,
    isHot: true,
    isRecommended: true,
    createdAt: '2024-01-03T14:20:00Z',
    updatedAt: '2024-01-10T09:45:00Z'
  },
  {
    id: '5',
    title: 'React 18 新特性与最佳实践',
    description: '深入学习React 18的新特性，包括Concurrent Features、Suspense、自动批处理等，掌握现代React开发模式。',
    cover: 'https://via.placeholder.com/400x250/00BCD4/FFFFFF?text=React+18',
    mentorId: '1',
    mentorName: '张大师',
    mentorAvatar: 'https://via.placeholder.com/50x50/4CAF50/FFFFFF?text=张',
    category: 'frontend',
    tags: ['React', '前端开发', 'JavaScript'],
    price: 329,
    originalPrice: 429,
    duration: 16,
    level: 'advanced',
    rating: 4.8,
    studentCount: 680,
    isNew: true,
    isHot: false,
    isRecommended: false,
    createdAt: '2024-01-12T08:00:00Z',
    updatedAt: '2024-01-16T15:30:00Z'
  },
  {
    id: '6',
    title: 'Node.js 后端开发实战',
    description: '学习Node.js后端开发，包括Express框架、数据库操作、API设计、性能优化等核心技能。',
    cover: 'https://via.placeholder.com/400x250/4CAF50/FFFFFF?text=Node.js',
    mentorId: '2',
    mentorName: '李导师',
    mentorAvatar: 'https://via.placeholder.com/50x50/2196F3/FFFFFF?text=李',
    category: 'backend',
    tags: ['Node.js', '后端开发', 'JavaScript'],
    price: 289,
    originalPrice: 389,
    duration: 14,
    level: 'intermediate',
    rating: 4.7,
    studentCount: 945,
    isNew: false,
    isHot: true,
    isRecommended: false,
    createdAt: '2024-01-06T13:15:00Z',
    updatedAt: '2024-01-13T11:20:00Z'
  },
  {
    id: '7',
    title: '机器学习基础与实践',
    description: '入门机器学习，学习基本算法、模型训练、特征工程等核心概念，通过实际项目掌握机器学习技能。',
    cover: 'https://via.placeholder.com/400x250/E91E63/FFFFFF?text=Machine+Learning',
    mentorId: '4',
    mentorName: '陈工程师',
    mentorAvatar: 'https://via.placeholder.com/50x50/9C27B0/FFFFFF?text=陈',
    category: 'ai',
    tags: ['机器学习', 'Python', '算法'],
    price: 459,
    originalPrice: 559,
    duration: 25,
    level: 'advanced',
    rating: 4.9,
    studentCount: 520,
    isNew: false,
    isHot: true,
    isRecommended: false,
    createdAt: '2024-01-02T10:30:00Z',
    updatedAt: '2024-01-09T14:10:00Z'
  },
  {
    id: '8',
    title: 'UI/UX 设计基础与实战',
    description: '学习UI/UX设计的基本原则和工具使用，包括用户研究、原型设计、视觉设计等核心技能。',
    cover: 'https://via.placeholder.com/400x250/FF5722/FFFFFF?text=UI+UX+Design',
    mentorId: '5',
    mentorName: '赵设计师',
    mentorAvatar: 'https://via.placeholder.com/50x50/FF5722/FFFFFF?text=赵',
    category: 'design',
    tags: ['UI设计', 'UX设计', 'Figma'],
    price: 259,
    originalPrice: 359,
    duration: 10,
    level: 'beginner',
    rating: 4.5,
    studentCount: 680,
    isNew: true,
    isHot: false,
    isRecommended: false,
    createdAt: '2024-01-15T09:00:00Z',
    updatedAt: '2024-01-17T16:45:00Z'
  }
]

// 课程服务
export const mockCoursesService = {
  // 搜索课程
  async searchCourses(params: {
    query?: string
    category?: string
    minPrice?: number
    maxPrice?: number
    minRating?: number
    duration?: string
    level?: string
    skills?: string[]
    sort?: string
    page?: number
    pageSize?: number
  }) {
    await delay(800)
    
    let filteredCourses = [...mockCourses]
    
    // 关键词搜索
    if (params.query) {
      const query = params.query.toLowerCase()
      filteredCourses = filteredCourses.filter(course => 
        course.title.toLowerCase().includes(query) ||
        course.description.toLowerCase().includes(query) ||
        course.mentorName.toLowerCase().includes(query) ||
        course.tags.some(tag => tag.toLowerCase().includes(query))
      )
    }
    
    // 分类筛选
    if (params.category) {
      filteredCourses = filteredCourses.filter(course => course.category === params.category)
    }
    
    // 价格筛选
    if (params.minPrice !== undefined) {
      filteredCourses = filteredCourses.filter(course => course.price >= params.minPrice!)
    }
    if (params.maxPrice !== undefined) {
      filteredCourses = filteredCourses.filter(course => course.price <= params.maxPrice!)
    }
    
    // 评分筛选
    if (params.minRating && params.minRating > 0) {
      filteredCourses = filteredCourses.filter(course => course.rating >= params.minRating!)
    }
    
    // 时长筛选
    if (params.duration) {
      const durationRanges = {
        short: (duration: number) => duration >= 1 && duration <= 5,
        medium: (duration: number) => duration > 5 && duration <= 20,
        long: (duration: number) => duration > 20
      }
      const range = durationRanges[params.duration as keyof typeof durationRanges]
      if (range) {
        filteredCourses = filteredCourses.filter(course => range(course.duration))
      }
    }
    
    // 难度筛选
    if (params.level) {
      filteredCourses = filteredCourses.filter(course => course.level === params.level)
    }
    
    // 技能标签筛选
    if (params.skills && params.skills.length > 0) {
      filteredCourses = filteredCourses.filter(course => 
        params.skills!.some(skill => course.tags.some(tag => tag.toLowerCase().includes(skill.toLowerCase())))
      )
    }
    
    // 排序
    if (params.sort) {
      const sortFunctions = {
        latest: (a: any, b: any) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime(),
        popular: (a: any, b: any) => b.studentCount - a.studentCount,
        rating: (a: any, b: any) => b.rating - a.rating,
        price_asc: (a: any, b: any) => a.price - b.price,
        price_desc: (a: any, b: any) => b.price - a.price,
        recommended: (a: any, b: any) => {
          // 推荐排序：推荐课程优先，然后按评分和学员数排序
          if (a.isRecommended && !b.isRecommended) return -1
          if (!a.isRecommended && b.isRecommended) return 1
          return (b.rating * 0.6 + b.studentCount / 1000 * 0.4) - (a.rating * 0.6 + a.studentCount / 1000 * 0.4)
        }
      }
      const sortFunction = sortFunctions[params.sort as keyof typeof sortFunctions]
      if (sortFunction) {
        filteredCourses.sort(sortFunction)
      }
    }
    
    // 分页
    const page = params.page || 1
    const pageSize = params.pageSize || 12
    const start = (page - 1) * pageSize
    const end = start + pageSize
    const paginatedCourses = filteredCourses.slice(start, end)
    
    return {
      success: true,
      data: {
        courses: paginatedCourses,
        total: filteredCourses.length,
        page,
        pageSize,
        totalPages: Math.ceil(filteredCourses.length / pageSize)
      },
      message: '搜索课程成功'
    }
  },
  
  // 获取推荐课程
  async getRecommendedCourses() {
    await delay(500)
    
    const recommendedCourses = mockCourses.filter(course => course.isRecommended)
    
    return {
      success: true,
      data: recommendedCourses,
      message: '获取推荐课程成功'
    }
  },
  
  // 获取课程详情
  async getCourseDetail(courseId: string) {
    await delay(600)
    
    const course = mockCourses.find(c => c.id === courseId)
    if (!course) {
      throw new Error('课程不存在')
    }
    
    // 添加课程大纲
    const courseDetail = {
      ...course,
      outline: [
        {
          title: '课程介绍',
          lessons: [
            { title: '课程概述', duration: '10分钟', type: 'video', isFree: true, preview: true },
            { title: '学习目标', duration: '5分钟', type: 'video', isFree: true },
            { title: '环境准备', duration: '15分钟', type: 'video', isFree: true }
          ]
        },
        {
          title: '基础知识',
          lessons: [
            { title: 'Vue.js 3.0 新特性', duration: '30分钟', type: 'video', isFree: false },
            { title: 'Composition API 入门', duration: '45分钟', type: 'video', isFree: false },
            { title: '响应式系统', duration: '40分钟', type: 'video', isFree: false },
            { title: '生命周期钩子', duration: '25分钟', type: 'video', isFree: false }
          ]
        },
        {
          title: '组件开发',
          lessons: [
            { title: '组件基础', duration: '35分钟', type: 'video', isFree: false },
            { title: '组件通信', duration: '50分钟', type: 'video', isFree: false },
            { title: '插槽和指令', duration: '40分钟', type: 'video', isFree: false },
            { title: '组件复用', duration: '30分钟', type: 'video', isFree: false }
          ]
        },
        {
          title: '状态管理',
          lessons: [
            { title: 'Pinia 状态管理', duration: '60分钟', type: 'video', isFree: false },
            { title: '路由管理', duration: '45分钟', type: 'video', isFree: false },
            { title: 'API 集成', duration: '50分钟', type: 'video', isFree: false }
          ]
        },
        {
          title: '项目实战',
          lessons: [
            { title: '项目搭建', duration: '30分钟', type: 'video', isFree: false },
            { title: '功能开发', duration: '90分钟', type: 'video', isFree: false },
            { title: '性能优化', duration: '40分钟', type: 'video', isFree: false },
            { title: '部署上线', duration: '25分钟', type: 'video', isFree: false }
          ]
        }
      ],
      requirements: [
        '基本的编程基础',
        '熟悉HTML、CSS、JavaScript',
        '具备一定的学习能力'
      ],
      targetAudience: [
        '前端开发初学者',
        '想要学习Vue.js的开发者',
        '希望提升前端技能的工程师'
      ],
      reviews: [
        {
          id: '1',
          userName: '学员A',
          userAvatar: 'https://via.placeholder.com/40x40/4CAF50/FFFFFF?text=A',
          rating: 5,
          content: '课程内容非常实用，老师讲解得很清楚，学到了很多实用的技能。',
          createdAt: '2024-01-15T10:30:00Z'
        },
        {
          id: '2',
          userName: '学员B',
          userAvatar: 'https://via.placeholder.com/40x40/2196F3/FFFFFF?text=B',
          rating: 4,
          content: '课程质量很高，项目实战部分很有帮助，推荐给想要学习Vue.js的朋友。',
          createdAt: '2024-01-14T14:20:00Z'
        },
        {
          id: '3',
          userName: '学员C',
          userAvatar: 'https://via.placeholder.com/40x40/FF9800/FFFFFF?text=C',
          rating: 5,
          content: '大师的讲解非常深入浅出，从基础概念到实战应用都有详细说明。',
          createdAt: '2024-01-13T16:45:00Z'
        },
        {
          id: '4',
          userName: '学员D',
          userAvatar: 'https://via.placeholder.com/40x40/9C27B0/FFFFFF?text=D',
          rating: 3,
          content: '课程内容不错，但有些地方讲解得比较快，需要反复观看。',
          createdAt: '2024-01-12T09:15:00Z'
        },
        {
          id: '5',
          userName: '学员E',
          userAvatar: 'https://via.placeholder.com/40x40/E91E63/FFFFFF?text=E',
          rating: 5,
          content: '学完这个课程后，我对Vue.js的理解更加深入了，强烈推荐！',
          createdAt: '2024-01-11T11:30:00Z'
        }
      ],
      // 学习进度相关字段（如果用户已报名）
      enrollmentStatus: 'enrolled', // enrolled, not_enrolled
      progress: 65, // 学习进度百分比
      completedLessons: 13, // 已完成课时数
      totalLessons: 20, // 总课时数
      studyTime: 8.5 // 学习时长（小时）
    }
    
    return {
      success: true,
      data: courseDetail,
      message: '获取课程详情成功'
    }
  },
  
  // 报名课程
  async enrollCourse(enrollData: {
    courseId: string
    userId: string
    price: number
    paymentMethod?: string
    userInfo?: any
  }) {
    await delay(1000)
    
    const course = mockCourses.find(c => c.id === enrollData.courseId)
    if (!course) {
      throw new Error('课程不存在')
    }
    
    // 模拟报名成功
    const enrollment = {
      id: Date.now().toString(),
      courseId: enrollData.courseId,
      userId: enrollData.userId,
      courseTitle: course.title,
      courseDescription: course.description,
      courseCover: course.cover,
      mentorName: course.mentorName,
      duration: course.duration,
      price: enrollData.price,
      paymentMethod: enrollData.paymentMethod || 'alipay',
      userInfo: enrollData.userInfo,
      status: 'enrolled',
      enrolledAt: new Date().toISOString(),
      expiresAt: new Date(Date.now() + 365 * 24 * 60 * 60 * 1000).toISOString() // 1年后过期
    }
    
    return {
      success: true,
      data: enrollment,
      message: '课程报名成功'
    }
  },
  
  // 获取用户已报名课程
  async getUserEnrolledCourses(userId: string) {
    await delay(600)
    
    // 模拟用户已报名的课程
    const enrolledCourses = mockCourses.slice(0, 3).map(course => ({
      ...course,
      enrollmentId: `enrollment_${course.id}`,
      enrolledAt: new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000).toISOString(),
      progress: Math.floor(Math.random() * 100),
      lastAccessedAt: new Date(Date.now() - Math.random() * 7 * 24 * 60 * 60 * 1000).toISOString()
    }))
    
    return {
      success: true,
      data: enrolledCourses,
      message: '获取已报名课程成功'
    }
  }
}