export interface User {
  id: string
  email: string
  primaryIdentity: 'master' | 'apprentice'
  identities: Identity[]
  currentIdentityId: string
  avatar?: string
  createdAt: string
  updatedAt: string
}

export interface Identity {
  id: string
  type: 'master' | 'apprentice'
  domain: string
  name: string
  avatar?: string
  bio?: string
  skills?: string[]
  email?: string
  wechat?: string
  phone?: string
  price?: number
  serviceTypes?: string[]
  experience?: string
  background?: string
  learningGoals?: string[]
  learningPreferences?: string[]
  isActive: boolean
  isVerified?: boolean
  status?: 'pending' | 'approved' | 'rejected'
  createdAt: string
  updatedAt?: string
}

export interface RegisterData {
  email: string
  password: string
  primaryIdentity: 'master' | 'apprentice'
}