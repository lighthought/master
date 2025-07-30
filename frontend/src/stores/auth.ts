import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { ApiService } from '@/services/api'
import type { User, Identity, RegisterData } from '@/types/user'



export const useAuthStore = defineStore('auth', () => {
  // 状态
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)
  const loading = ref(false)

  // 计算属性
  const isAuthenticated = computed(() => !!user.value && !!token.value)
  
  const currentIdentity = computed(() => {
    if (!user.value) return null
    
    if (user.value.currentIdentityId) {
      return user.value.identities.find(id => id.id === user.value!.currentIdentityId)
    }
    
    // 如果没有设置当前身份，返回第一个活跃身份
    return user.value.identities.find(id => id.isActive) || null
  })

  const isMaster = computed(() => currentIdentity.value?.type === 'master')
  const isApprentice = computed(() => currentIdentity.value?.type === 'apprentice')

  // 动作
  const login = async (email: string, password: string) => {
    loading.value = true
    try {
      const result = await ApiService.auth.login(email, password)
      
      user.value = result.user
      token.value = result.token
      
      // 保存到本地存储
      localStorage.setItem('auth_token', result.token)
      localStorage.setItem('user_data', JSON.stringify(result.user))
      
      return result
    } catch (error) {
      console.error('登录失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const register = async (userData: {
    email: string
    password: string
    primaryIdentity: 'master' | 'apprentice'
  }) => {
    loading.value = true
    try {
      const result = await ApiService.auth.register(userData)
      
      user.value = result.user
      token.value = result.token
      
      // 保存到本地存储
      localStorage.setItem('auth_token', result.token)
      localStorage.setItem('user_data', JSON.stringify(result.user))
      
      return result
    } catch (error) {
      console.error('注册失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const logout = () => {
    user.value = null
    token.value = null
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user_data')
  }

  const switchIdentity = async (identityId: string) => {
    if (!user.value) return
    
    try {
      const result = await ApiService.auth.switchIdentity(user.value.id, identityId)
      
      user.value.currentIdentityId = identityId
      localStorage.setItem('user_data', JSON.stringify(user.value))
      
      return result
    } catch (error) {
      console.error('切换身份失败:', error)
      throw error
    }
  }

  const createMasterIdentity = async (identityData: {
    name: string
    domain: string
    skills: string[]
    bio: string
    experience: string
    price: number
    serviceTypes: string[]
  }) => {
    if (!user.value) return
    
    try {
      const result = await ApiService.auth.createMasterIdentity(user.value.id, identityData)
      
      // 更新用户身份列表
      user.value.identities.push(result.identity)
      localStorage.setItem('user_data', JSON.stringify(user.value))
      
      return result.identity
    } catch (error) {
      console.error('创建大师身份失败:', error)
      throw error
    }
  }

  const createApprenticeIdentity = async (identityData: {
    name: string
    domain: string
    background: string
    currentLevel: string
    learningGoals: string[]
    expectedDuration: string
    learningPreferences: string[]
    timePreferences: string[]
    budgetRange: string
  }) => {
    if (!user.value) return
    
    try {
      const result = await ApiService.auth.createApprenticeIdentity(user.value.id, identityData)
      
      // 更新用户身份列表
      user.value.identities.push(result.identity)
      localStorage.setItem('user_data', JSON.stringify(user.value))
      
      return result.identity
    } catch (error) {
      console.error('创建学徒身份失败:', error)
      throw error
    }
  }

  const updateIdentityInfo = async (identityId: string, identityData: any) => {
    if (!user.value) return
    
    try {
      const result = await ApiService.auth.updateIdentityInfo(user.value.id, identityId, identityData)
      
      // 更新用户身份列表中的对应身份
      const identityIndex = user.value.identities.findIndex(id => id.id === identityId)
      if (identityIndex !== -1) {
        user.value.identities[identityIndex] = { ...user.value.identities[identityIndex], ...result.identity }
        localStorage.setItem('user_data', JSON.stringify(user.value))
      }
      
      return result.identity
    } catch (error) {
      console.error('更新身份信息失败:', error)
      throw error
    }
  }

  const initializeAuth = () => {
    const savedToken = localStorage.getItem('auth_token')
    const savedUser = localStorage.getItem('user_data')
    
    if (savedToken && savedUser) {
      try {
        token.value = savedToken
        user.value = JSON.parse(savedUser)
      } catch (error) {
        console.error('初始化认证状态失败:', error)
        logout()
      }
    }
  }

  return {
    // 状态
    user,
    token,
    loading,
    
    // 计算属性
    isAuthenticated,
    currentIdentity,
    isMaster,
    isApprentice,
    userIdentities: computed(() => user.value?.identities || []),
    
    // 动作
    login,
    register,
    logout,
    switchIdentity,
    createMasterIdentity,
    createApprenticeIdentity,
    updateIdentityInfo,
    initializeAuth
  }
})