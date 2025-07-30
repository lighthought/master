<template>
  <div id="app">
    <!-- 认证页面不需要导航栏 -->
    <template v-if="$route.path === '/auth'">
      <router-view />
    </template>
    
    <!-- 主应用布局 -->
    <template v-else>
      <!-- 顶部导航栏 -->
      <header class="app-header">
        <div class="header-content">
          <div class="header-left">
            <div class="logo">Master Guide</div>
          </div>
          
          <div class="header-center">
            <div class="search-container">
              <el-input
                v-model="searchQuery"
                placeholder="搜索大师、课程、内容..."
                prefix-icon="Search"
                class="search-input"
              />
            </div>
          </div>
          
          <div class="header-right">
            <div class="notification-container">
              <el-badge :value="3" class="notification-badge">
                <el-button type="text" class="notification-btn">
                  <el-icon><Bell /></el-icon>
                </el-button>
              </el-badge>
            </div>
            
            <!-- 身份切换器 -->
            <div v-if="authStore.isAuthenticated && authStore.userIdentities.length > 1" class="identity-switcher-container">
              <IdentitySwitcher @identity-changed="handleIdentityChanged" />
            </div>
            
            <div class="user-menu-container">
              <el-dropdown>
                <el-avatar :size="40" src="" />
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item @click="$router.push('/profile')">个人中心</el-dropdown-item>
                    <el-dropdown-item @click="$router.push('/identity')">身份管理</el-dropdown-item>
                    <el-dropdown-item>设置</el-dropdown-item>
                    <el-dropdown-item divided @click="handleLogout">退出登录</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
          </div>
        </div>
      </header>
      
      <!-- 主要内容区域 -->
      <main class="app-main">
        <router-view />
      </main>
      
      <!-- 底部导航栏 -->
      <nav class="app-bottom-nav">
        <div class="nav-items">
          <router-link 
            v-for="item in navItems" 
            :key="item.path"
            :to="item.path"
            class="nav-item"
            :class="{ active: $route.path === item.path }"
          >
            <el-icon class="nav-icon">
              <component :is="item.icon" />
            </el-icon>
            <span class="nav-text">{{ item.text }}</span>
          </router-link>
        </div>
      </nav>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ElMessage } from 'element-plus'
import IdentitySwitcher from '@/components/identity/IdentitySwitcher.vue'

const router = useRouter()
const authStore = useAuthStore()

const searchQuery = ref('')

const navItems = [
  { path: '/', text: '首页', icon: 'House' },
  { path: '/mentors', text: '大师', icon: 'User' },
  { path: '/bookings', text: '预约', icon: 'Calendar' },
  { path: '/courses', text: '课程', icon: 'Reading' },
  { path: '/circles', text: '圈子', icon: 'ChatDotRound' },
  { path: '/profile', text: '我的', icon: 'UserFilled' }
]

// 处理退出登录
const handleLogout = () => {
  authStore.logout()
  ElMessage.success('已退出登录')
  router.push('/auth')
}

// 处理身份切换
const handleIdentityChanged = (identity: any) => {
  console.log('身份已切换:', identity)
  // 可以在这里添加身份切换后的逻辑，比如更新页面内容等
}

// 组件挂载时初始化认证状态
onMounted(() => {
  authStore.initializeAuth()
})
</script>

<style scoped lang="scss">
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: var(--bg-primary);
}

// 顶部导航栏
.app-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  background: var(--bg-secondary);
  border-bottom: 1px solid var(--border-color);
  height: 64px;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 var(--spacing-lg);
  max-width: 1200px;
  margin: 0 auto;
}

.header-left {
  flex-shrink: 0;
}

.logo {
  font-size: var(--font-size-h3);
  font-weight: var(--font-weight-bold);
  color: var(--primary-color);
}

.header-center {
  flex: 1;
  max-width: 400px;
  margin: 0 var(--spacing-xl);
}

.search-input {
  width: 100%;
}

.header-right {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  flex-shrink: 0;
}

.identity-switcher-container {
  margin-right: var(--spacing-md);
}

.notification-btn {
  color: var(--text-secondary);
  font-size: var(--icon-size-lg);
  
  &:hover {
    color: var(--primary-color);
  }
}

// 主要内容区域
.app-main {
  flex: 1;
  margin-top: 64px;
  margin-bottom: 60px;
}

// 底部导航栏
.app-bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  background: var(--bg-secondary);
  border-top: 1px solid var(--border-color);
  height: 60px;
}

.nav-items {
  display: flex;
  height: 100%;
}

.nav-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-decoration: none;
  color: var(--text-secondary);
  transition: all var(--transition-normal);
  gap: var(--spacing-xs);
  
  &:hover {
    color: var(--primary-color);
  }
  
  &.active {
    color: var(--primary-color);
  }
}

.nav-icon {
  font-size: var(--icon-size-lg);
}

.nav-text {
  font-size: var(--font-size-small);
  font-weight: var(--font-weight-medium);
}

// 响应式设计
@media (min-width: 768px) {
  .app-bottom-nav {
    display: none;
  }
  
  .app-main {
    margin-bottom: 0;
  }
}
</style>