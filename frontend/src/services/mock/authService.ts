import type { User, Identity } from '@/stores/auth'

// 模拟用户数据
const mockUsers: User[] = [
  {
    id: '1',
    email: 'test@example.com',
    primaryIdentity: 'apprentice',
    identities: [
      {
        id: '1',
        type: 'apprentice',
        domain: '软件开发',
        name: '学习中的开发者',
        avatar: 'https://via.placeholder.com/80x80/4CAF50/FFFFFF?text=学徒',
        isActive: true,
        isVerified: true,
        status: 'approved',
        createdAt: '2024-01-01T00:00:00Z'
      }
    ],
    currentIdentityId: '1',
    createdAt: '2024-01-01T00:00:00Z',
    updatedAt: '2024-01-01T00:00:00Z'
  }
]

// 模拟延迟
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

export const mockAuthService = {
  // 用户注册
  async register(userData: {
    email: string
    password: string
    primaryIdentity: 'master' | 'apprentice'
  }) {
    await delay(1500)
    
    // 检查邮箱是否已存在
    const existingUser = mockUsers.find(user => user.email === userData.email)
    if (existingUser) {
      throw new Error('邮箱已被注册')
    }
    
    // 创建新用户
    const newUser: User = {
      id: Date.now().toString(),
      email: userData.email,
      primaryIdentity: userData.primaryIdentity,
      identities: [
        {
          id: Date.now().toString(),
          type: userData.primaryIdentity,
          domain: userData.primaryIdentity === 'master' ? '待选择' : '待选择',
          name: userData.primaryIdentity === 'master' ? '新晋大师' : '学习中的学徒',
          avatar: userData.primaryIdentity === 'master' 
            ? 'https://via.placeholder.com/80x80/FF6B35/FFFFFF?text=大师'
            : 'https://via.placeholder.com/80x80/4CAF50/FFFFFF?text=学徒',
          isActive: true,
          isVerified: userData.primaryIdentity === 'apprentice',
          status: userData.primaryIdentity === 'apprentice' ? 'approved' : 'pending',
          createdAt: new Date().toISOString()
        }
      ],
      currentIdentityId: Date.now().toString(),
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    }
    
    mockUsers.push(newUser)
    
    return {
      success: true,
      user: newUser,
      token: `mock-jwt-token-${Date.now()}`
    }
  },

  // 用户登录
  async login(email: string, password: string) {
    await delay(1000)
    
    const user = mockUsers.find(u => u.email === email)
    if (!user) {
      throw new Error('用户不存在')
    }
    
    // 模拟密码验证
    if (password !== '123456') {
      throw new Error('密码错误')
    }
    
    return {
      success: true,
      user,
      token: `mock-jwt-token-${Date.now()}`
    }
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
    await delay(2000)
    
    const user = mockUsers.find(u => u.id === userId)
    if (!user) {
      throw new Error('用户不存在')
    }
    
    const newIdentity: Identity = {
      id: Date.now().toString(),
      type: 'master',
      domain: identityData.domain,
      name: identityData.name,
      avatar: 'https://via.placeholder.com/80x80/FF6B35/FFFFFF?text=大师',
      isActive: false, // 审核中
      isVerified: false,
      status: 'pending',
      createdAt: new Date().toISOString()
    }
    
    user.identities.push(newIdentity)
    user.updatedAt = new Date().toISOString()
    
    return {
      success: true,
      identity: newIdentity
    }
  },

  // 创建学徒身份
  async createApprenticeIdentity(userId: string, identityData: {
    name: string
    domain: string
    learningGoals: string[]
    background: string
    preferences: string[]
  }) {
    await delay(1500)
    
    const user = mockUsers.find(u => u.id === userId)
    if (!user) {
      throw new Error('用户不存在')
    }
    
    const newIdentity: Identity = {
      id: Date.now().toString(),
      type: 'apprentice',
      domain: identityData.domain,
      name: identityData.name,
      avatar: 'https://via.placeholder.com/80x80/4CAF50/FFFFFF?text=学徒',
      isActive: true,
      isVerified: true,
      status: 'approved',
      createdAt: new Date().toISOString()
    }
    
    user.identities.push(newIdentity)
    user.updatedAt = new Date().toISOString()
    
    return {
      success: true,
      identity: newIdentity
    }
  },

  // 切换身份
  async switchIdentity(userId: string, identityId: string) {
    await delay(500)
    
    const user = mockUsers.find(u => u.id === userId)
    if (!user) {
      throw new Error('用户不存在')
    }
    
    const identity = user.identities.find(id => id.id === identityId)
    if (!identity) {
      throw new Error('身份不存在')
    }
    
    if (!identity.isActive) {
      throw new Error('身份未激活')
    }
    
    user.currentIdentityId = identityId
    user.updatedAt = new Date().toISOString()
    
    return {
      success: true,
      currentIdentity: identity
    }
  },

  // 获取用户信息
  async getUserInfo(userId: string) {
    await delay(300)
    
    const user = mockUsers.find(u => u.id === userId)
    if (!user) {
      throw new Error('用户不存在')
    }
    
    return {
      success: true,
      user
    }
  }
}