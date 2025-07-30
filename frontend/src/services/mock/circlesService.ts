const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 圈子分类
export const circleCategories = [
  { id: 'software', name: '软件开发', icon: 'Code', color: '#409eff' },
  { id: 'design', name: 'UI设计', icon: 'Brush', color: '#67c23a' },
  { id: 'marketing', name: '数字营销', icon: 'TrendCharts', color: '#e6a23c' },
  { id: 'craft', name: '传统工艺', icon: 'Tools', color: '#f56c6c' }
]

// 模拟圈子数据
const mockCircles = [
  {
    id: '1',
    name: 'Vue.js 开发者社区',
    description: '专注于Vue.js技术交流和学习，分享开发经验和最佳实践',
    category: 'software',
    cover: 'https://via.placeholder.com/400x200/409eff/FFFFFF?text=Vue.js',
    memberCount: 1250,
    postCount: 3420,
    isActive: true,
    isJoined: false,
    createdAt: '2024-01-01T00:00:00Z',
    tags: ['Vue.js', '前端开发', 'JavaScript'],
    rules: [
      '禁止发布与Vue.js无关的内容',
      '保持友善的交流氛围',
      '分享有价值的开发经验'
    ]
  },
  {
    id: '2',
    name: 'UI设计师交流群',
    description: 'UI设计师的交流平台，分享设计作品、设计理念和行业动态',
    category: 'design',
    cover: 'https://via.placeholder.com/400x200/67c23a/FFFFFF?text=UI+Design',
    memberCount: 890,
    postCount: 2156,
    isActive: true,
    isJoined: true,
    createdAt: '2024-01-05T00:00:00Z',
    tags: ['UI设计', '用户体验', '设计工具'],
    rules: [
      '分享原创设计作品',
      '尊重他人设计成果',
      '提供建设性意见'
    ]
  },
  {
    id: '3',
    name: '数字营销实战',
    description: '数字营销从业者的学习交流平台，分享营销策略和实战经验',
    category: 'marketing',
    cover: 'https://via.placeholder.com/400x200/e6a23c/FFFFFF?text=Digital+Marketing',
    memberCount: 1560,
    postCount: 4280,
    isActive: true,
    isJoined: false,
    createdAt: '2024-01-10T00:00:00Z',
    tags: ['数字营销', 'SEO', '社交媒体'],
    rules: [
      '分享真实的营销案例',
      '禁止虚假宣传',
      '保持专业讨论'
    ]
  },
  {
    id: '4',
    name: '传统木工技艺',
    description: '传统木工技艺传承与交流，分享木工技巧和作品展示',
    category: 'craft',
    cover: 'https://via.placeholder.com/400x200/f56c6c/FFFFFF?text=Woodworking',
    memberCount: 320,
    postCount: 890,
    isActive: true,
    isJoined: false,
    createdAt: '2024-01-15T00:00:00Z',
    tags: ['木工', '传统工艺', '手工制作'],
    rules: [
      '分享木工制作过程',
      '展示传统工艺作品',
      '传承技艺文化'
    ]
  },
  {
    id: '5',
    name: 'React 技术社区',
    description: 'React开发者技术交流社区，探讨React生态和开发技巧',
    category: 'software',
    cover: 'https://via.placeholder.com/400x200/409eff/FFFFFF?text=React',
    memberCount: 2100,
    postCount: 5680,
    isActive: true,
    isJoined: false,
    createdAt: '2024-01-20T00:00:00Z',
    tags: ['React', '前端开发', 'JavaScript'],
    rules: [
      '专注React技术讨论',
      '分享开发经验',
      '保持技术交流'
    ]
  },
  {
    id: '6',
    name: '插画设计师联盟',
    description: '插画设计师的创作交流平台，分享插画作品和创作心得',
    category: 'design',
    cover: 'https://via.placeholder.com/400x200/67c23a/FFFFFF?text=Illustration',
    memberCount: 680,
    postCount: 1890,
    isActive: true,
    isJoined: false,
    createdAt: '2024-01-25T00:00:00Z',
    tags: ['插画', '艺术设计', '创作分享'],
    rules: [
      '分享原创插画作品',
      '尊重版权',
      '鼓励创作交流'
    ]
  }
]

// 圈子服务
export const mockCirclesService = {
  // 获取圈子列表
  async getCircles(params: {
    category?: string
    query?: string
    sort?: string
    page?: number
    pageSize?: number
  }) {
    await delay(600)
    
    let filteredCircles = [...mockCircles]
    
    // 分类筛选
    if (params.category) {
      filteredCircles = filteredCircles.filter(circle => circle.category === params.category)
    }
    
    // 关键词搜索
    if (params.query) {
      const query = params.query.toLowerCase()
      filteredCircles = filteredCircles.filter(circle => 
        circle.name.toLowerCase().includes(query) ||
        circle.description.toLowerCase().includes(query) ||
        circle.tags.some(tag => tag.toLowerCase().includes(query))
      )
    }
    
    // 排序
    if (params.sort) {
      switch (params.sort) {
        case 'memberCount':
          filteredCircles.sort((a, b) => b.memberCount - a.memberCount)
          break
        case 'postCount':
          filteredCircles.sort((a, b) => b.postCount - a.postCount)
          break
        case 'createdAt':
          filteredCircles.sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime())
          break
        case 'name':
          filteredCircles.sort((a, b) => a.name.localeCompare(b.name))
          break
      }
    }
    
    // 分页
    const page = params.page || 1
    const pageSize = params.pageSize || 10
    const startIndex = (page - 1) * pageSize
    const endIndex = startIndex + pageSize
    const paginatedCircles = filteredCircles.slice(startIndex, endIndex)
    
    return {
      success: true,
      data: {
        circles: paginatedCircles,
        total: filteredCircles.length,
        page,
        pageSize,
        totalPages: Math.ceil(filteredCircles.length / pageSize)
      },
      message: '获取圈子列表成功'
    }
  },

  // 获取圈子详情
  async getCircleDetail(circleId: string) {
    await delay(400)
    
    const circle = mockCircles.find(c => c.id === circleId)
    if (!circle) {
      throw new Error('圈子不存在')
    }
    
    return {
      success: true,
      data: circle,
      message: '获取圈子详情成功'
    }
  },

  // 加入圈子
  async joinCircle(circleId: string, userId: string) {
    await delay(300)
    
    const circle = mockCircles.find(c => c.id === circleId)
    if (!circle) {
      throw new Error('圈子不存在')
    }
    
    if (circle.isJoined) {
      throw new Error('已经加入该圈子')
    }
    
    circle.isJoined = true
    circle.memberCount++
    
    return {
      success: true,
      data: { message: '加入圈子成功' },
      message: '加入圈子成功'
    }
  },

  // 退出圈子
  async leaveCircle(circleId: string, userId: string) {
    await delay(300)
    
    const circle = mockCircles.find(c => c.id === circleId)
    if (!circle) {
      throw new Error('圈子不存在')
    }
    
    if (!circle.isJoined) {
      throw new Error('未加入该圈子')
    }
    
    circle.isJoined = false
    circle.memberCount--
    
    return {
      success: true,
      data: { message: '退出圈子成功' },
      message: '退出圈子成功'
    }
  },

  // 获取用户加入的圈子
  async getUserJoinedCircles(userId: string) {
    await delay(500)
    
    const joinedCircles = mockCircles.filter(circle => circle.isJoined)
    
    return {
      success: true,
      data: joinedCircles,
      message: '获取已加入圈子成功'
    }
  },

  // 获取圈子分类
  async getCircleCategories() {
    await delay(200)
    
    return {
      success: true,
      data: circleCategories,
      message: '获取圈子分类成功'
    }
  },

  // 搜索圈子
  async searchCircles(query: string) {
    await delay(400)
    
    const searchResults = mockCircles.filter(circle => 
      circle.name.toLowerCase().includes(query.toLowerCase()) ||
      circle.description.toLowerCase().includes(query.toLowerCase()) ||
      circle.tags.some(tag => tag.toLowerCase().includes(query.toLowerCase()))
    )
    
    return {
      success: true,
      data: searchResults,
      message: '搜索圈子成功'
    }
  }
} 