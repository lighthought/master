 # Master Guide 前端架构设计

## 1. 前端技术栈详解

### 1.1 核心技术
- **Vue 3**：渐进式JavaScript框架，使用Composition API
- **Vite**：下一代前端构建工具，提供极速的开发体验
- **TypeScript**：类型安全的JavaScript超集
- **Vue Router 4**：Vue.js官方路由管理器
- **Pinia**：Vue 3状态管理库
- **Element Plus**：基于Vue 3的组件库

### 1.2 辅助技术
- **Socket.io-client**：实时通信客户端
- **Axios**：HTTP客户端
- **Day.js**：轻量级日期处理库
- **Lodash**：实用工具库
- **VueUse**：Vue组合式函数工具集

## 2. 项目结构设计

```
frontend/
├── public/                    # 静态资源
│   ├── favicon.ico
│   └── index.html
├── src/
│   ├── assets/               # 资源文件
│   │   ├── images/          # 图片资源
│   │   ├── icons/           # 图标资源
│   │   └── styles/          # 样式文件
│   ├── components/          # 通用组件
│   │   ├── common/          # 基础组件
│   │   ├── layout/          # 布局组件
│   │   └── business/        # 业务组件
│   ├── views/               # 页面组件
│   │   ├── home/            # 首页
│   │   ├── mentors/         # 大师页面
│   │   ├── courses/         # 课程页面
│   │   ├── community/       # 社群页面
│   │   └── profile/         # 个人中心
│   ├── router/              # 路由配置
│   │   ├── index.ts
│   │   └── routes.ts
│   ├── stores/              # 状态管理
│   │   ├── auth.ts          # 认证状态
│   │   ├── user.ts          # 用户状态
│   │   ├── identity.ts      # 身份状态
│   │   └── app.ts           # 应用状态
│   ├── services/            # API服务
│   │   ├── api.ts           # API配置
│   │   ├── auth.ts          # 认证服务
│   │   ├── user.ts          # 用户服务
│   │   ├── course.ts        # 课程服务
│   │   └── community.ts     # 社群服务
│   ├── utils/               # 工具函数
│   │   ├── request.ts       # HTTP请求封装
│   │   ├── storage.ts       # 本地存储
│   │   ├── socket.ts        # WebSocket封装
│   │   └── helpers.ts       # 辅助函数
│   ├── types/               # TypeScript类型定义
│   │   ├── user.ts
│   │   ├── course.ts
│   │   └── common.ts
│   ├── App.vue              # 根组件
│   └── main.ts              # 入口文件
├── .env                     # 环境变量
├── .env.development         # 开发环境变量
├── .env.production          # 生产环境变量
├── package.json
├── tsconfig.json            # TypeScript配置
├── vite.config.ts           # Vite配置
└── README.md
```

## 3. 核心组件设计

### 3.1 布局组件

#### AppLayout.vue
```vue
<template>
  <div class="app-layout">
    <!-- 顶部导航栏 -->
    <header class="header">
      <div class="logo">Master Guide</div>
      <div class="identity-switcher">
        <IdentitySwitcher />
      </div>
      <div class="user-menu">
        <UserMenu />
      </div>
    </header>
    
    <!-- 主要内容区域 -->
    <main class="main-content">
      <router-view />
    </main>
    
    <!-- 底部导航栏 -->
    <nav class="bottom-nav">
      <BottomNavigation />
    </nav>
  </div>
</template>
```

#### IdentitySwitcher.vue
```vue
<template>
  <div class="identity-switcher">
    <el-dropdown @command="handleIdentitySwitch">
      <span class="current-identity">
        {{ currentIdentity.name }} ({{ currentIdentity.type }})
        <el-icon><ArrowDown /></el-icon>
      </span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item 
            v-for="identity in userIdentities" 
            :key="identity.id"
            :command="identity.id"
          >
            {{ identity.name }} ({{ identity.type }})
          </el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>
</template>
```

### 3.2 业务组件

#### MentorCard.vue
```vue
<template>
  <div class="mentor-card" @click="handleMentorClick">
    <div class="mentor-avatar">
      <el-avatar :src="mentor.avatar" :size="80" />
      <div class="online-status" :class="{ online: mentor.isOnline }"></div>
    </div>
    
    <div class="mentor-info">
      <h3 class="mentor-name">{{ mentor.name }}</h3>
      <p class="mentor-domain">{{ mentor.domain }}</p>
      <div class="mentor-rating">
        <el-rate v-model="mentor.rating" disabled />
        <span class="student-count">{{ mentor.studentCount }} 学生</span>
      </div>
      <p class="mentor-price">¥{{ mentor.price }}/小时</p>
    </div>
  </div>
</template>
```

#### CourseCard.vue
```vue
<template>
  <div class="course-card">
    <div class="course-cover">
      <img :src="course.cover" :alt="course.title" />
      <div class="course-status" :class="course.status">
        {{ getStatusText(course.status) }}
      </div>
    </div>
    
    <div class="course-info">
      <h3 class="course-title">{{ course.title }}</h3>
      <p class="course-description">{{ course.description }}</p>
      
      <div class="course-meta">
        <span class="duration">{{ course.duration }}小时</span>
        <span class="difficulty">{{ course.difficulty }}</span>
        <span class="students">{{ course.studentCount }}人学习</span>
      </div>
      
      <div class="course-footer">
        <span class="price">¥{{ course.price }}</span>
        <el-button type="primary" @click="handleEnroll">立即报名</el-button>
      </div>
    </div>
  </div>
</template>
```

## 4. 状态管理设计

### 4.1 Pinia Store结构

#### auth.ts - 认证状态管理
```typescript
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User, Identity } from '@/types/user'

export const useAuthStore = defineStore('auth', () => {
  // 状态
  const user = ref<User | null>(null)
  const currentIdentity = ref<Identity | null>(null)
  const token = ref<string | null>(null)
  const isAuthenticated = ref(false)

  // 计算属性
  const isMaster = computed(() => 
    currentIdentity.value?.type === 'master'
  )
  
  const isApprentice = computed(() => 
    currentIdentity.value?.type === 'apprentice'
  )

  // 方法
  const login = async (credentials: LoginCredentials) => {
    try {
      const response = await authService.login(credentials)
      user.value = response.user
      token.value = response.token
      isAuthenticated.value = true
      
      // 设置默认身份
      if (response.identities.length > 0) {
        currentIdentity.value = response.identities[0]
      }
      
      // 存储到本地
      localStorage.setItem('token', response.token)
      localStorage.setItem('currentIdentity', JSON.stringify(currentIdentity.value))
    } catch (error) {
      throw error
    }
  }

  const logout = () => {
    user.value = null
    currentIdentity.value = null
    token.value = null
    isAuthenticated.value = false
    
    localStorage.removeItem('token')
    localStorage.removeItem('currentIdentity')
  }

  const switchIdentity = async (identityId: string) => {
    try {
      const response = await authService.switchIdentity(identityId)
      currentIdentity.value = response.identity
      localStorage.setItem('currentIdentity', JSON.stringify(currentIdentity.value))
    } catch (error) {
      throw error
    }
  }

  return {
    user,
    currentIdentity,
    token,
    isAuthenticated,
    isMaster,
    isApprentice,
    login,
    logout,
    switchIdentity
  }
})
```

#### user.ts - 用户状态管理
```typescript
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User, Identity, UserProfile } from '@/types/user'

export const useUserStore = defineStore('user', () => {
  // 状态
  const userIdentities = ref<Identity[]>([])
  const userProfile = ref<UserProfile | null>(null)
  const learningProgress = ref<any[]>([])
  const teachingStats = ref<any>(null)

  // 计算属性
  const hasMultipleIdentities = computed(() => 
    userIdentities.value.length > 1
  )

  const masterIdentities = computed(() => 
    userIdentities.value.filter(id => id.type === 'master')
  )

  const apprenticeIdentities = computed(() => 
    userIdentities.value.filter(id => id.type === 'apprentice')
  )

  // 方法
  const fetchUserIdentities = async () => {
    try {
      const response = await userService.getIdentities()
      userIdentities.value = response.identities
    } catch (error) {
      throw error
    }
  }

  const updateProfile = async (profile: Partial<UserProfile>) => {
    try {
      const response = await userService.updateProfile(profile)
      userProfile.value = response.profile
    } catch (error) {
      throw error
    }
  }

  return {
    userIdentities,
    userProfile,
    learningProgress,
    teachingStats,
    hasMultipleIdentities,
    masterIdentities,
    apprenticeIdentities,
    fetchUserIdentities,
    updateProfile
  }
})
```

## 5. 路由设计

### 5.1 路由配置
```typescript
// router/routes.ts
import type { RouteRecordRaw } from 'vue-router'

export const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/home/HomePage.vue'),
    meta: { title: '首页', requiresAuth: false }
  },
  {
    path: '/mentors',
    name: 'Mentors',
    component: () => import('@/views/mentors/MentorsPage.vue'),
    meta: { title: '大师', requiresAuth: false }
  },
  {
    path: '/mentors/:id',
    name: 'MentorDetail',
    component: () => import('@/views/mentors/MentorDetailPage.vue'),
    meta: { title: '大师详情', requiresAuth: false }
  },
  {
    path: '/courses',
    name: 'Courses',
    component: () => import('@/views/courses/CoursesPage.vue'),
    meta: { title: '课程', requiresAuth: false }
  },
  {
    path: '/courses/:id',
    name: 'CourseDetail',
    component: () => import('@/views/courses/CourseDetailPage.vue'),
    meta: { title: '课程详情', requiresAuth: false }
  },
  {
    path: '/community',
    name: 'Community',
    component: () => import('@/views/community/CommunityPage.vue'),
    meta: { title: '社群', requiresAuth: true }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/profile/ProfilePage.vue'),
    meta: { title: '个人中心', requiresAuth: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/LoginPage.vue'),
    meta: { title: '登录', requiresAuth: false }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/auth/RegisterPage.vue'),
    meta: { title: '注册', requiresAuth: false }
  }
]
```

### 5.2 路由守卫
```typescript
// router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import { routes } from './routes'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  // 设置页面标题
  document.title = to.meta.title ? `${to.meta.title} - Master Guide` : 'Master Guide'
  
  // 检查是否需要认证
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
    return
  }
  
  // 已登录用户访问登录页面，重定向到首页
  if (authStore.isAuthenticated && (to.name === 'Login' || to.name === 'Register')) {
    next('/')
    return
  }
  
  next()
})

export default router
```

## 6. API服务设计

### 6.1 HTTP请求封装
```typescript
// utils/request.ts
import axios from 'axios'
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '@/stores/auth'

class RequestService {
  private instance: AxiosInstance

  constructor() {
    this.instance = axios.create({
      baseURL: import.meta.env.VITE_API_BASE_URL,
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json'
      }
    })

    this.setupInterceptors()
  }

  private setupInterceptors() {
    // 请求拦截器
    this.instance.interceptors.request.use(
      (config) => {
        const authStore = useAuthStore()
        if (authStore.token) {
          config.headers.Authorization = `Bearer ${authStore.token}`
        }
        return config
      },
      (error) => {
        return Promise.reject(error)
      }
    )

    // 响应拦截器
    this.instance.interceptors.response.use(
      (response: AxiosResponse) => {
        return response.data
      },
      (error) => {
        if (error.response?.status === 401) {
          const authStore = useAuthStore()
          authStore.logout()
          window.location.href = '/login'
        }
        
        ElMessage.error(error.response?.data?.message || '请求失败')
        return Promise.reject(error)
      }
    )
  }

  public get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.get(url, config)
  }

  public post<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.post(url, data, config)
  }

  public put<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.put(url, data, config)
  }

  public delete<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.delete(url, config)
  }
}

export const request = new RequestService()
```

### 6.2 API服务模块
```typescript
// services/auth.ts
import { request } from '@/utils/request'
import type { LoginCredentials, RegisterData, AuthResponse } from '@/types/auth'

export const authService = {
  async login(credentials: LoginCredentials): Promise<AuthResponse> {
    return request.post('/auth/login', credentials)
  },

  async register(data: RegisterData): Promise<AuthResponse> {
    return request.post('/auth/register', data)
  },

  async switchIdentity(identityId: string): Promise<{ identity: any }> {
    return request.post('/auth/switch-identity', { identityId })
  },

  async refreshToken(): Promise<{ token: string }> {
    return request.post('/auth/refresh')
  }
}
```

## 7. 实时通信设计

### 7.1 WebSocket封装
```typescript
// utils/socket.ts
import { io, Socket } from 'socket.io-client'
import { useAuthStore } from '@/stores/auth'

class SocketService {
  private socket: Socket | null = null
  private reconnectAttempts = 0
  private maxReconnectAttempts = 5

  connect() {
    const authStore = useAuthStore()
    
    if (!authStore.token) {
      console.warn('No token available for socket connection')
      return
    }

    this.socket = io(import.meta.env.VITE_WS_URL, {
      auth: {
        token: authStore.token
      },
      transports: ['websocket', 'polling']
    })

    this.setupEventListeners()
  }

  private setupEventListeners() {
    if (!this.socket) return

    this.socket.on('connect', () => {
      console.log('Socket connected')
      this.reconnectAttempts = 0
    })

    this.socket.on('disconnect', () => {
      console.log('Socket disconnected')
    })

    this.socket.on('connect_error', (error) => {
      console.error('Socket connection error:', error)
      this.handleReconnect()
    })

    // 消息事件
    this.socket.on('message', (data) => {
      this.handleMessage(data)
    })

    // 在线状态更新
    this.socket.on('user_status_update', (data) => {
      this.handleUserStatusUpdate(data)
    })
  }

  private handleReconnect() {
    if (this.reconnectAttempts < this.maxReconnectAttempts) {
      this.reconnectAttempts++
      setTimeout(() => {
        this.connect()
      }, 1000 * this.reconnectAttempts)
    }
  }

  private handleMessage(data: any) {
    // 处理接收到的消息
    console.log('Received message:', data)
  }

  private handleUserStatusUpdate(data: any) {
    // 处理用户状态更新
    console.log('User status update:', data)
  }

  emit(event: string, data: any) {
    if (this.socket) {
      this.socket.emit(event, data)
    }
  }

  disconnect() {
    if (this.socket) {
      this.socket.disconnect()
      this.socket = null
    }
  }
}

export const socketService = new SocketService()
```

## 8. 样式设计

### 8.1 主题配置
```scss
// assets/styles/variables.scss
:root {
  // 主色调
  --primary-color: #FF6B35;
  --secondary-color: #FFD93D;
  --success-color: #4CAF50;
  --warning-color: #FF9800;
  --danger-color: #F44336;
  
  // 背景色
  --bg-primary: #1A1A1A;
  --bg-secondary: #2D2D2D;
  --bg-tertiary: #404040;
  
  // 文字色
  --text-primary: #FFFFFF;
  --text-secondary: #CCCCCC;
  --text-tertiary: #999999;
  
  // 边框色
  --border-color: #404040;
  --border-light: #666666;
  
  // 阴影
  --shadow-light: 0 2px 8px rgba(0, 0, 0, 0.1);
  --shadow-medium: 0 4px 16px rgba(0, 0, 0, 0.15);
  --shadow-heavy: 0 8px 32px rgba(0, 0, 0, 0.2);
  
  // 圆角
  --border-radius-small: 4px;
  --border-radius-medium: 8px;
  --border-radius-large: 16px;
  
  // 间距
  --spacing-xs: 4px;
  --spacing-sm: 8px;
  --spacing-md: 16px;
  --spacing-lg: 24px;
  --spacing-xl: 32px;
}
```

### 8.2 响应式设计
```scss
// assets/styles/responsive.scss
// 移动端优先设计
.container {
  width: 100%;
  padding: 0 var(--spacing-md);
  margin: 0 auto;
}

// 平板
@media (min-width: 768px) {
  .container {
    max-width: 750px;
  }
}

// 桌面端
@media (min-width: 1024px) {
  .container {
    max-width: 960px;
  }
}

// 大屏
@media (min-width: 1200px) {
  .container {
    max-width: 1140px;
  }
}
```

## 9. 性能优化

### 9.1 代码分割
```typescript
// vite.config.ts
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src')
    }
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          'vendor': ['vue', 'vue-router', 'pinia'],
          'element-plus': ['element-plus'],
          'socket': ['socket.io-client']
        }
      }
    }
  }
})
```

### 9.2 图片优化
```vue
<!-- 组件中使用懒加载 -->
<template>
  <img 
    v-lazy="imageUrl" 
    :alt="alt"
    loading="lazy"
    @error="handleImageError"
  />
</template>
```

## 10. 测试策略

### 10.1 单元测试
```typescript
// tests/components/MentorCard.test.ts
import { mount } from '@vue/test-utils'
import MentorCard from '@/components/MentorCard.vue'

describe('MentorCard', () => {
  it('renders mentor information correctly', () => {
    const mentor = {
      id: '1',
      name: '张三',
      avatar: '/avatar.jpg',
      domain: '软件开发',
      rating: 4.5,
      studentCount: 100,
      price: 200,
      isOnline: true
    }

    const wrapper = mount(MentorCard, {
      props: { mentor }
    })

    expect(wrapper.find('.mentor-name').text()).toBe('张三')
    expect(wrapper.find('.mentor-domain').text()).toBe('软件开发')
    expect(wrapper.find('.mentor-price').text()).toBe('¥200/小时')
  })
})
```

### 10.2 E2E测试
```typescript
// tests/e2e/login.spec.ts
import { test, expect } from '@playwright/test'

test('user can login successfully', async ({ page }) => {
  await page.goto('/login')
  
  await page.fill('[data-testid="email"]', 'test@example.com')
  await page.fill('[data-testid="password"]', 'password123')
  await page.click('[data-testid="login-button"]')
  
  await expect(page).toHaveURL('/')
  await expect(page.locator('[data-testid="user-menu"]')).toBeVisible()
})
```

## 11. 部署配置

### 11.1 构建配置
```typescript
// vite.config.ts
export default defineConfig({
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    sourcemap: false,
    minify: 'terser',
    terserOptions: {
      compress: {
        drop_console: true,
        drop_debugger: true
      }
    }
  }
})
```

### 11.2 Docker配置
```dockerfile
# Dockerfile
FROM node:18-alpine as builder

WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

## 12. 总结

Master Guide前端采用Vue 3 + Vite技术栈，具备以下优势：

### 技术优势
- **开发体验**：Vite提供极速的开发体验
- **类型安全**：TypeScript确保代码质量
- **组件化**：Element Plus提供丰富的UI组件
- **状态管理**：Pinia提供简洁的状态管理
- **实时通信**：Socket.io支持实时功能

### 架构特点
- **模块化设计**：清晰的目录结构和组件划分
- **状态管理**：基于Pinia的响应式状态管理
- **路由管理**：Vue Router 4支持路由守卫
- **API封装**：统一的HTTP请求处理
- **实时通信**：WebSocket支持实时消息

该前端架构设计确保了良好的用户体验、开发效率和代码质量，为Master Guide平台提供了坚实的前端基础。