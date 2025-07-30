<template>
  <div class="home-page">
    <!-- 个性化欢迎区域 -->
    <PersonalizedWelcome />
    
    <!-- 原有内容保留作为备用 -->
    <div class="welcome-section">
      <h1 class="welcome-title">{{ greeting }}</h1>
      <div class="user-info">
        <el-avatar :size="60" :src="currentIdentity?.avatar" />
        <div class="user-details">
          <h2 class="user-name">{{ currentIdentity?.name || '用户' }}</h2>
          <el-tag :type="getIdentityType(currentIdentity?.type)" size="large">
            {{ getIdentityLabel(currentIdentity?.type) }}
          </el-tag>
        </div>
      </div>
    </div>
    
    <div class="quick-actions">
      <h3>快速操作</h3>
      <div class="action-buttons">
        <el-button type="primary" @click="$router.push('/identity')">
          <el-icon><User /></el-icon>
          身份管理
        </el-button>
        <el-button 
          v-if="authStore.isApprentice"
          type="success" 
          @click="$router.push('/mentors')"
        >
          <el-icon><User /></el-icon>
          浏览大师
        </el-button>
        <el-button 
          v-if="authStore.isApprentice"
          type="warning" 
          @click="$router.push('/courses')"
        >
          <el-icon><Reading /></el-icon>
          课程学习
        </el-button>
        <el-button 
          v-if="authStore.isMaster"
          type="success" 
          @click="$router.push('/mentors')"
        >
          <el-icon><User /></el-icon>
          管理指导
        </el-button>
      </div>
    </div>
    
    <!-- 身份切换提示 -->
    <div v-if="userIdentities.length > 1" class="identity-tip">
      <div class="tip-content">
        <el-icon class="tip-icon"><InfoFilled /></el-icon>
        <div class="tip-text">
          <h4>多身份用户</h4>
          <p>你拥有 {{ userIdentities.length }} 个身份，可以在顶部导航栏快速切换身份，体验不同的功能。</p>
        </div>
      </div>
    </div>
    
    <div class="status-info">
      <h3>状态信息</h3>
      <div class="status-cards">
        <div class="status-card">
          <div class="status-icon">
            <el-icon><User /></el-icon>
          </div>
          <div class="status-content">
            <h4>身份数量</h4>
            <p>{{ userIdentities.length }} 个</p>
          </div>
        </div>
        <div class="status-card">
          <div class="status-icon">
            <el-icon><Check /></el-icon>
          </div>
          <div class="status-content">
            <h4>已认证</h4>
            <p>{{ verifiedIdentities.length }} 个</p>
          </div>
        </div>
        <div class="status-card">
          <div class="status-icon">
            <el-icon><Clock /></el-icon>
          </div>
          <div class="status-content">
            <h4>审核中</h4>
            <p>{{ pendingIdentities.length }} 个</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { InfoFilled } from '@element-plus/icons-vue'
import PersonalizedWelcome from '@/components/home/PersonalizedWelcome.vue'

const authStore = useAuthStore()

// 计算属性
const currentIdentity = computed(() => authStore.currentIdentity)
const userIdentities = computed(() => authStore.userIdentities)

const verifiedIdentities = computed(() => 
  userIdentities.value.filter(id => id.isVerified)
)

const pendingIdentities = computed(() => 
  userIdentities.value.filter(id => id.status === 'pending')
)

// 问候语
const greeting = computed(() => {
  const hour = new Date().getHours()
  if (hour < 6) return '夜深了，注意休息'
  if (hour < 12) return '早上好'
  if (hour < 18) return '下午好'
  return '晚上好'
})

// 获取身份类型标签
const getIdentityType = (type?: string) => {
  return type === 'master' ? 'warning' : 'success'
}

// 获取身份标签文本
const getIdentityLabel = (type?: string) => {
  return type === 'master' ? '大师' : '学徒'
}
</script>

<style scoped lang="scss">
.home-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: var(--spacing-xl);
}

.welcome-section {
  text-align: center;
  margin-bottom: var(--spacing-xxl);
  padding: var(--spacing-xl);
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
}

.welcome-title {
  font-size: var(--font-size-h1);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-lg);
}

.user-info {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--spacing-lg);
}

.user-details {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-sm);
}

.user-name {
  font-size: var(--font-size-h3);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0;
}

.quick-actions {
  margin-bottom: var(--spacing-xxl);
  padding: var(--spacing-xl);
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
}

.quick-actions h3 {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-lg);
}

.action-buttons {
  display: flex;
  gap: var(--spacing-md);
  flex-wrap: wrap;
}

.status-info {
  padding: var(--spacing-xl);
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
}

.status-info h3 {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-lg);
}

.status-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-lg);
}

.status-card {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  transition: all var(--transition-normal);
  
  &:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-light);
  }
}

.status-icon {
  width: 50px;
  height: 50px;
  background: var(--primary-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: var(--icon-size-lg);
  color: var(--text-primary);
}

.status-content h4 {
  font-size: var(--font-size-medium);
  font-weight: var(--font-weight-medium);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-xs) 0;
}

.status-content p {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-bold);
  color: var(--primary-color);
  margin: 0;
}

.identity-tip {
  margin-top: var(--spacing-xl);
  padding: var(--spacing-lg);
  background: linear-gradient(135deg, var(--primary-color), var(--apprentice-color));
  border-radius: var(--border-radius-large);
  color: white;
}

.tip-content {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.tip-icon {
  font-size: var(--icon-size-lg);
  flex-shrink: 0;
}

.tip-text h4 {
  font-size: var(--font-size-h5);
  font-weight: var(--font-weight-semibold);
  margin: 0 0 var(--spacing-xs) 0;
}

.tip-text p {
  font-size: var(--font-size-medium);
  margin: 0;
  opacity: 0.9;
}

// 响应式设计
@media (max-width: 768px) {
  .home-page {
    padding: var(--spacing-lg);
  }
  
  .user-info {
    flex-direction: column;
    gap: var(--spacing-md);
  }
  
  .action-buttons {
    flex-direction: column;
  }
  
  .status-cards {
    grid-template-columns: 1fr;
  }
}
</style>