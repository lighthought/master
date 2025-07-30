<template>
  <div class="auth-page">
    <div class="auth-container">
      <div class="auth-content">
        <div class="auth-header">
          <div class="logo">
            <h1>Master Guide</h1>
            <p>大师指导平台</p>
          </div>
        </div>
        
        <div class="auth-form-container">
          <RegisterForm
            v-if="currentView === 'register'"
            @switch-to-login="switchToLogin"
            @register-success="handleRegisterSuccess"
          />
          
          <LoginForm
            v-else-if="currentView === 'login'"
            @switch-to-register="switchToRegister"
            @login-success="handleLoginSuccess"
          />
        </div>
      </div>
      
      <div class="auth-background">
        <div class="background-content">
          <h2>技艺传承的新时代</h2>
          <p>连接大师与学习者，让专业技能得到更好的传承</p>
          <div class="feature-list">
            <div class="feature-item">
              <el-icon><User /></el-icon>
              <span>双重身份支持</span>
            </div>
            <div class="feature-item">
              <el-icon><VideoPlay /></el-icon>
              <span>实时指导服务</span>
            </div>
            <div class="feature-item">
              <el-icon><ChatDotRound /></el-icon>
              <span>专业社群交流</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import RegisterForm from '@/components/auth/RegisterForm.vue'
import LoginForm from '@/components/auth/LoginForm.vue'

const router = useRouter()
const authStore = useAuthStore()

// 当前显示的视图
const currentView = ref<'login' | 'register'>('register')

// 切换到登录视图
const switchToLogin = () => {
  currentView.value = 'login'
}

// 切换到注册视图
const switchToRegister = () => {
  currentView.value = 'register'
}

// 处理注册成功
const handleRegisterSuccess = (userData: any) => {
  console.log('注册成功:', userData)
  // 注册成功后自动跳转到首页
  router.push('/')
}

// 处理登录成功
const handleLoginSuccess = (userData: any) => {
  console.log('登录成功:', userData)
  // 登录成功后自动跳转到首页
  router.push('/')
}

// 组件挂载时检查认证状态
onMounted(() => {
  authStore.initializeAuth()
  
  // 如果已经登录，直接跳转到首页
  if (authStore.isAuthenticated) {
    router.push('/')
  }
})
</script>

<style scoped lang="scss">
.auth-page {
  min-height: 100vh;
  background: linear-gradient(135deg, var(--bg-secondary) 0%, var(--bg-primary) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--spacing-lg);
}

.auth-container {
  display: flex;
  max-width: 1200px;
  width: 100%;
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  overflow: hidden;
  box-shadow: var(--shadow-heavy);
}

.auth-content {
  flex: 1;
  padding: var(--spacing-xxl);
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.auth-header {
  text-align: center;
  margin-bottom: var(--spacing-xl);
}

.logo h1 {
  font-size: var(--font-size-h1);
  font-weight: var(--font-weight-bold);
  color: var(--primary-color);
  margin-bottom: var(--spacing-xs);
}

.logo p {
  font-size: var(--font-size-large);
  color: var(--text-secondary);
}

.auth-form-container {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}

.auth-background {
  flex: 1;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-dark) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--spacing-xxl);
  color: var(--text-primary);
}

.background-content {
  text-align: center;
  max-width: 400px;
}

.background-content h2 {
  font-size: var(--font-size-h2);
  font-weight: var(--font-weight-bold);
  margin-bottom: var(--spacing-lg);
  line-height: 1.3;
}

.background-content p {
  font-size: var(--font-size-large);
  margin-bottom: var(--spacing-xl);
  line-height: 1.6;
  opacity: 0.9;
}

.feature-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.feature-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  font-size: var(--font-size-medium);
  
  .el-icon {
    font-size: var(--icon-size-lg);
    color: var(--secondary-color);
  }
}

// 响应式设计
@media (max-width: 1024px) {
  .auth-background {
    display: none;
  }
  
  .auth-content {
    padding: var(--spacing-xl);
  }
}

@media (max-width: 768px) {
  .auth-page {
    padding: var(--spacing-md);
  }
  
  .auth-container {
    border-radius: var(--border-radius-medium);
  }
  
  .auth-content {
    padding: var(--spacing-lg);
  }
  
  .logo h1 {
    font-size: var(--font-size-h2);
  }
  
  .logo p {
    font-size: var(--font-size-medium);
  }
}
</style>