<template>
  <div id="app">
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
          
          <div class="user-menu-container">
            <el-dropdown>
              <el-avatar :size="40" src="" />
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>个人中心</el-dropdown-item>
                  <el-dropdown-item>设置</el-dropdown-item>
                  <el-dropdown-item divided>退出登录</el-dropdown-item>
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
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const searchQuery = ref('')

const navItems = [
  { path: '/', text: '首页', icon: 'House' },
  { path: '/mentors', text: '大师', icon: 'User' },
  { path: '/courses', text: '课程', icon: 'Reading' },
  { path: '/community', text: '社群', icon: 'ChatDotRound' },
  { path: '/profile', text: '我的', icon: 'UserFilled' }
]
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