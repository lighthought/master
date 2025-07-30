import type { User } from '@/types/user'

// 模拟延迟
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟大师数据
const mockMentors = [
  {
    id: '1',
    name: '张大师',
    avatar: 'https://via.placeholder.com/80x80/4CAF50/FFFFFF?text=张',
    domain: '前端开发',
    skills: ['Vue.js', 'React', 'TypeScript', 'Node.js', 'Webpack'],
    rating: 4.9,
    studentCount: 128,
    price: 200,
    isOnline: true,
    isVerified: true,
    bio: '资深前端开发工程师，5年Vue.js开发经验，擅长组件化开发和性能优化。专注于现代前端技术栈，对Vue.js生态系统有深入研究，曾参与多个大型项目的架构设计和开发。',
    experience: '5年前端开发经验，曾就职于BAT等知名互联网公司。主导过多个千万级用户的前端项目，对前端工程化、性能优化、用户体验设计有丰富经验。',
    serviceTypes: ['1对1指导', '代码审查', '项目实战', '架构设计'],
    achievements: ['Vue.js官方文档贡献者', '开源项目维护者', '技术大会演讲嘉宾'],
    languages: ['中文', '英文'],
    responseTime: '2小时内',
    completionRate: 98
  },
  {
    id: '2',
    name: '李导师',
    avatar: 'https://via.placeholder.com/80x80/2196F3/FFFFFF?text=李',
    domain: '后端开发',
    skills: ['Java', 'Spring Boot', 'MySQL', 'Redis', 'Docker'],
    rating: 4.8,
    studentCount: 95,
    price: 180,
    isOnline: false,
    isVerified: true,
    bio: '全栈开发工程师，专注于Java后端开发，有丰富的微服务架构经验。对分布式系统、高并发处理、数据库优化有深入研究，曾设计并实现多个高可用系统。',
    experience: '8年Java开发经验，精通Spring生态，有大型项目架构经验。曾主导过多个亿级用户的后端系统设计，对微服务架构、容器化部署、DevOps有丰富实践。',
    serviceTypes: ['1对1指导', '架构设计', '代码优化', '系统设计'],
    achievements: ['Spring社区贡献者', '技术博客作者', '架构师认证'],
    languages: ['中文'],
    responseTime: '4小时内',
    completionRate: 96
  },
  {
    id: '3',
    name: '王老师',
    avatar: 'https://via.placeholder.com/80x80/FF9800/FFFFFF?text=王',
    domain: '移动开发',
    skills: ['Flutter', 'React Native', 'iOS', 'Android', 'Dart'],
    rating: 4.7,
    studentCount: 76,
    price: 220,
    isOnline: true,
    isVerified: true,
    bio: '移动端开发专家，跨平台开发经验丰富，擅长性能优化和用户体验设计。',
    experience: '6年移动开发经验，曾开发多个百万级用户应用。',
    serviceTypes: ['1对1指导', 'UI/UX设计', '性能优化'],
    achievements: ['Flutter社区贡献者', '移动应用设计奖'],
    languages: ['中文', '英文'],
    responseTime: '1小时内',
    completionRate: 99
  },
  {
    id: '4',
    name: '陈工程师',
    avatar: 'https://via.placeholder.com/80x80/9C27B0/FFFFFF?text=陈',
    domain: '人工智能',
    skills: ['Python', 'TensorFlow', 'PyTorch', '机器学习', '深度学习'],
    rating: 4.9,
    studentCount: 64,
    price: 300,
    isOnline: true,
    isVerified: true,
    bio: 'AI算法工程师，专注于机器学习和深度学习，有丰富的实战项目经验。',
    experience: '7年AI开发经验，曾参与多个国家级AI项目。',
    serviceTypes: ['1对1指导', '算法设计', '模型优化'],
    achievements: ['AI论文发表', '算法竞赛获奖'],
    languages: ['中文', '英文'],
    responseTime: '3小时内',
    completionRate: 97
  },
  {
    id: '5',
    name: '刘架构师',
    avatar: 'https://via.placeholder.com/80x80/E91E63/FFFFFF?text=刘',
    domain: '系统架构',
    skills: ['微服务', '分布式系统', 'Kubernetes', '云原生', 'DevOps'],
    rating: 4.8,
    studentCount: 52,
    price: 350,
    isOnline: false,
    isVerified: true,
    bio: '系统架构师，专注于大规模分布式系统设计和云原生架构。',
    experience: '10年系统架构经验，曾设计多个亿级用户系统。',
    serviceTypes: ['1对1指导', '架构设计', '技术咨询'],
    achievements: ['架构设计专利', '技术大会演讲'],
    languages: ['中文', '英文'],
    responseTime: '6小时内',
    completionRate: 95
  },
  {
    id: '6',
    name: '赵设计师',
    avatar: 'https://via.placeholder.com/80x80/00BCD4/FFFFFF?text=赵',
    domain: 'UI/UX设计',
    skills: ['Figma', 'Sketch', 'Adobe XD', '用户研究', '交互设计'],
    rating: 4.6,
    studentCount: 88,
    price: 150,
    isOnline: true,
    isVerified: true,
    bio: '资深UI/UX设计师，专注于用户体验设计和产品设计。',
    experience: '5年设计经验，曾为多个知名产品设计界面。',
    serviceTypes: ['1对1指导', '设计评审', '原型设计'],
    achievements: ['设计大赛获奖', '设计作品展示'],
    languages: ['中文'],
    responseTime: '2小时内',
    completionRate: 94
  }
]

export const mockMentorsService = {
  // 获取推荐大师列表
  async getRecommendedMentors(userId: string) {
    await delay(800)
    
    // 模拟推荐算法，根据用户偏好返回推荐大师
    const recommendedMentors = mockMentors
      .sort(() => Math.random() - 0.5) // 随机排序
      .slice(0, 6) // 返回6个推荐大师
    
    return {
      success: true,
      data: recommendedMentors
    }
  },

  // 获取大师详情
  async getMentorDetail(mentorId: string) {
    await delay(500)
    
    const mentor = mockMentors.find(m => m.id === mentorId)
    if (!mentor) {
      throw new Error('大师不存在')
    }
    
    return {
      success: true,
      data: mentor
    }
  },

  // 搜索大师
  async searchMentors(params: {
    keyword?: string
    domain?: string
    minPrice?: number
    maxPrice?: number
    minRating?: number
    isOnline?: boolean
    page?: number
    pageSize?: number
  }) {
    await delay(600)
    
    let filteredMentors = [...mockMentors]
    
    // 关键词搜索
    if (params.keyword) {
      filteredMentors = filteredMentors.filter(mentor => 
        mentor.name.includes(params.keyword!) ||
        mentor.domain.includes(params.keyword!) ||
        mentor.skills.some(skill => skill.includes(params.keyword!))
      )
    }
    
    // 领域筛选
    if (params.domain) {
      filteredMentors = filteredMentors.filter(mentor => 
        mentor.domain === params.domain
      )
    }
    
    // 价格筛选
    if (params.minPrice !== undefined) {
      filteredMentors = filteredMentors.filter(mentor => 
        mentor.price >= params.minPrice!
      )
    }
    
    if (params.maxPrice !== undefined) {
      filteredMentors = filteredMentors.filter(mentor => 
        mentor.price <= params.maxPrice!
      )
    }
    
    // 评分筛选
    if (params.minRating !== undefined) {
      filteredMentors = filteredMentors.filter(mentor => 
        mentor.rating >= params.minRating!
      )
    }
    
    // 在线状态筛选
    if (params.isOnline !== undefined) {
      filteredMentors = filteredMentors.filter(mentor => 
        mentor.isOnline === params.isOnline
      )
    }
    
    // 分页
    const page = params.page || 1
    const pageSize = params.pageSize || 10
    const start = (page - 1) * pageSize
    const end = start + pageSize
    
    return {
      success: true,
      data: {
        mentors: filteredMentors.slice(start, end),
        total: filteredMentors.length,
        page,
        pageSize,
        totalPages: Math.ceil(filteredMentors.length / pageSize)
      }
    }
  },

  // 获取大师评价
  async getMentorReviews(mentorId: string, page = 1, pageSize = 10) {
    await delay(400)
    
    // 模拟评价数据
    const mockReviews = [
      {
        id: '1',
        userId: 'user1',
        userName: '学员A',
        avatar: 'https://via.placeholder.com/40x40/4CAF50/FFFFFF?text=A',
        rating: 5,
        content: '张大师的指导非常专业，帮我解决了很多技术难题。他的Vue.js知识非常扎实，不仅教会了我基础概念，还分享了很多实战经验和最佳实践。强烈推荐给想要学习前端开发的同学！',
        createdAt: '2024-01-15T10:30:00Z'
      },
      {
        id: '2',
        userId: 'user2',
        userName: '学员B',
        avatar: 'https://via.placeholder.com/40x40/2196F3/FFFFFF?text=B',
        rating: 4,
        content: '老师很有耐心，讲解得很清楚，学到了很多实用的知识。特别是在性能优化方面给了我很多启发，让我对前端开发有了更深的理解。',
        createdAt: '2024-01-14T15:20:00Z'
      },
      {
        id: '3',
        userId: 'user3',
        userName: '学员C',
        avatar: 'https://via.placeholder.com/40x40/FF9800/FFFFFF?text=C',
        rating: 5,
        content: '技术实力很强，实战经验丰富，指导效果很好。张大师不仅技术过硬，而且很会教学，能够把复杂的概念讲得很简单易懂。',
        createdAt: '2024-01-13T09:15:00Z'
      },
      {
        id: '4',
        userId: 'user4',
        userName: '学员D',
        avatar: 'https://via.placeholder.com/40x40/9C27B0/FFFFFF?text=D',
        rating: 5,
        content: '非常专业的大师！在项目实战指导中，他帮我梳理了整个开发流程，从需求分析到代码实现，再到测试部署，每个环节都讲解得很详细。',
        createdAt: '2024-01-12T14:45:00Z'
      },
      {
        id: '5',
        userId: 'user5',
        userName: '学员E',
        avatar: 'https://via.placeholder.com/40x40/E91E63/FFFFFF?text=E',
        rating: 4,
        content: '张大师的代码审查很仔细，指出了很多我之前没有注意到的问题。通过他的指导，我的代码质量有了很大提升。',
        createdAt: '2024-01-11T16:20:00Z'
      }
    ]
    
    return {
      success: true,
      data: {
        reviews: mockReviews,
        total: mockReviews.length,
        page,
        pageSize,
        totalPages: 1
      }
    }
  }
}