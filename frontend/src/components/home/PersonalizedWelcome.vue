<template>
  <div class="personalized-welcome">
    <!-- 问候语和身份信息 -->
    <div class="welcome-header">
      <div class="greeting-section">
        <h1 class="greeting">{{ greeting }}</h1>
        <p class="user-name">{{ authStore.user?.email }}</p>
      </div>
      
      <div class="identity-section">
        <div class="identity-info">
          <el-avatar 
            :size="60" 
            :src="currentIdentity?.avatar"
            :icon="getIdentityIcon(currentIdentity?.type)"
          />
          <div class="identity-details">
            <h2 class="identity-name">{{ currentIdentity?.name || '用户' }}</h2>
            <div class="identity-meta">
              <el-tag 
                :type="getIdentityType(currentIdentity?.type)" 
                size="large"
                class="identity-tag"
              >
                {{ getIdentityLabel(currentIdentity?.type) }}
              </el-tag>
              <span class="identity-domain">{{ currentIdentity?.domain }}</span>
            </div>
          </div>
        </div>
        
        <!-- 快速身份切换 -->
        <div v-if="authStore.userIdentities.length > 1" class="identity-switcher">
          <el-button 
            type="text" 
            size="small"
            @click="showIdentitySwitcher = true"
          >
            <el-icon><Switch /></el-icon>
            切换身份
          </el-button>
        </div>
      </div>
    </div>
    
    <!-- 统计信息卡片 -->
    <div class="stats-section">
      <div class="stats-grid">
        <!-- 学徒统计 -->
        <div v-if="authStore.isApprentice" class="stat-card">
          <div class="stat-icon apprentice">
            <el-icon><Reading /></el-icon>
          </div>
          <div class="stat-content">
            <h3 class="stat-title">学习进度</h3>
            <div class="stat-value">{{ learningStats.totalCourses }}</div>
            <div class="stat-label">已完成课程</div>
            <el-progress 
              :percentage="learningStats.progress" 
              :color="progressColors"
              :stroke-width="8"
              class="stat-progress"
            />
          </div>
        </div>
        
        <!-- 大师统计 -->
        <div v-if="authStore.isMaster" class="stat-card">
          <div class="stat-icon master">
            <el-icon><Star /></el-icon>
          </div>
          <div class="stat-content">
            <h3 class="stat-title">指导统计</h3>
            <div class="stat-value">{{ teachingStats.totalStudents }}</div>
            <div class="stat-label">指导学员</div>
            <div class="stat-secondary">
              <span class="secondary-item">
                <el-icon><Clock /></el-icon>
                {{ teachingStats.totalHours }}小时
              </span>
              <span class="secondary-item">
                <el-icon><Money /></el-icon>
                ¥{{ teachingStats.totalEarnings }}
              </span>
            </div>
          </div>
        </div>
        
        <!-- 通用统计 -->
        <div class="stat-card">
          <div class="stat-icon general">
            <el-icon><Calendar /></el-icon>
          </div>
          <div class="stat-content">
            <h3 class="stat-title">活跃天数</h3>
            <div class="stat-value">{{ generalStats.activeDays }}</div>
            <div class="stat-label">连续登录</div>
            <div class="stat-secondary">
              <span class="secondary-item">
                <el-icon><Trophy /></el-icon>
                {{ generalStats.achievements }}个成就
              </span>
            </div>
          </div>
        </div>
        
        <!-- 身份数量统计 -->
        <div class="stat-card">
          <div class="stat-icon identity">
            <el-icon><User /></el-icon>
          </div>
          <div class="stat-content">
            <h3 class="stat-title">身份管理</h3>
            <div class="stat-value">{{ authStore.userIdentities.length }}</div>
            <div class="stat-label">拥有身份</div>
            <div class="stat-secondary">
              <span class="secondary-item">
                <el-icon><Check /></el-icon>
                {{ verifiedIdentities.length }}个已认证
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 快速操作 -->
    <div class="quick-actions">
      <h3 class="section-title">快速操作</h3>
      <div class="action-buttons">
        <el-button 
          v-if="authStore.isApprentice"
          type="primary" 
          size="large"
          @click="$router.push('/courses')"
          class="action-btn"
        >
          <el-icon><Reading /></el-icon>
          开始学习
        </el-button>
        
                       <el-button 
                 v-if="authStore.isApprentice"
                 type="success" 
                 size="large"
                 @click="$router.push('/mentors')"
                 class="action-btn"
               >
                 <el-icon><User /></el-icon>
                 寻找大师
               </el-button>
        
        <el-button 
          v-if="authStore.isMaster"
          type="primary" 
          size="large"
          @click="$router.push('/mentors')"
          class="action-btn"
        >
          <el-icon><Star /></el-icon>
          管理指导
        </el-button>
        
                       <el-button 
                 type="warning" 
                 size="large"
                 @click="$router.push('/identity')"
                 class="action-btn"
               >
                 <el-icon><Setting /></el-icon>
                 身份管理
               </el-button>
               
               <el-button 
                 type="info" 
                 size="large"
                 @click="$router.push('/bookings')"
                 class="action-btn"
               >
                 <el-icon><Calendar /></el-icon>
                 我的预约
               </el-button>
               
               <el-button 
                 type="success" 
                 size="large"
                 @click="$router.push('/courses')"
                 class="action-btn"
               >
                 <el-icon><Reading /></el-icon>
                 浏览课程
               </el-button>
               
               <el-button 
                 v-if="authStore.isMaster"
                 type="warning" 
                 size="large"
                 @click="$router.push('/mentor-bookings')"
                 class="action-btn"
               >
                 <el-icon><Setting /></el-icon>
                 预约管理
               </el-button>
      </div>
    </div>
    
    <!-- 身份切换器弹窗 -->
    <el-dialog
      v-model="showIdentitySwitcher"
      title="切换身份"
      width="90%"
      max-width="400px"
      :close-on-click-modal="true"
    >
      <div class="identity-list">
        <div
          v-for="identity in authStore.userIdentities"
          :key="identity.id"
          class="identity-item"
          :class="{ 'is-active': identity.id === currentIdentity?.id }"
          @click="switchToIdentity(identity)"
        >
          <el-avatar 
            :size="40" 
            :src="identity.avatar"
            :icon="getIdentityIcon(identity.type)"
          />
          <div class="item-info">
            <div class="item-name">{{ identity.name }}</div>
            <div class="item-meta">
              <el-tag 
                :type="getIdentityType(identity.type)" 
                size="small"
              >
                {{ getIdentityLabel(identity.type) }}
              </el-tag>
              <span class="item-domain">{{ identity.domain }}</span>
            </div>
          </div>
          <div class="item-action">
            <el-icon v-if="identity.id === currentIdentity?.id" class="check-icon">
              <Check />
            </el-icon>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { 
  Switch, Reading, Star, Calendar, User, Check, 
  Clock, Money, Trophy, Setting 
} from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'
import type { Identity } from '@/types/user'

// 认证store
const authStore = useAuthStore()

// 状态
const showIdentitySwitcher = ref(false)
const loading = ref(false)

// 统计数据
const learningStats = ref({
  totalCourses: 0,
  progress: 0,
  completedLessons: 0,
  totalLessons: 0
})

const teachingStats = ref({
  totalStudents: 0,
  totalHours: 0,
  totalEarnings: 0,
  averageRating: 0
})

const generalStats = ref({
  activeDays: 0,
  achievements: 0,
  totalLoginDays: 0
})

// 计算属性
const currentIdentity = computed(() => authStore.currentIdentity)

const verifiedIdentities = computed(() => 
  authStore.userIdentities.filter(id => id.isVerified)
)

// 问候语
const greeting = computed(() => {
  const hour = new Date().getHours()
  if (hour < 6) return '夜深了，注意休息'
  if (hour < 12) return '早上好'
  if (hour < 18) return '下午好'
  return '晚上好'
})

// 加载用户统计数据
const loadUserStats = async () => {
  if (!authStore.user) return
  
  loading.value = true
  try {
    // 加载通用统计
    const generalResult = await ApiService.userStats.getGeneralStats(authStore.user.id)
    generalStats.value = generalResult.data
    
    // 根据身份类型加载相应统计
    if (authStore.isApprentice) {
      const learningResult = await ApiService.userStats.getLearningStats(authStore.user.id)
      learningStats.value = learningResult.data
    } else if (authStore.isMaster) {
      const teachingResult = await ApiService.userStats.getTeachingStats(authStore.user.id)
      teachingStats.value = teachingResult.data
    }
  } catch (error) {
    console.error('加载用户统计数据失败:', error)
  } finally {
    loading.value = false
  }
}

// 进度条颜色
const progressColors = [
  { color: '#f56c6c', percentage: 20 },
  { color: '#e6a23c', percentage: 40 },
  { color: '#5cb87a', percentage: 60 },
  { color: '#1989fa', percentage: 80 },
  { color: '#6f7ad3', percentage: 100 }
]

// 获取身份图标
const getIdentityIcon = (type?: string) => {
  if (type === 'master') {
    return 'Star'
  } else if (type === 'apprentice') {
    return 'User'
  }
  return 'User'
}

// 获取身份类型标签
const getIdentityType = (type?: string) => {
  return type === 'master' ? 'warning' : 'success'
}

// 获取身份标签文本
const getIdentityLabel = (type?: string) => {
  return type === 'master' ? '大师' : '学徒'
}

// 切换到指定身份
const switchToIdentity = async (identity: Identity) => {
  if (identity.id === currentIdentity.value?.id) {
    showIdentitySwitcher.value = false
    return
  }
  
  if (!identity.isActive) {
    ElMessage.warning('该身份尚未激活，无法切换')
    return
  }
  
  try {
    await authStore.switchIdentity(identity.id)
    ElMessage.success(`已切换到${getIdentityLabel(identity.type)}身份`)
    showIdentitySwitcher.value = false
  } catch (error) {
    ElMessage.error('身份切换失败，请重试')
  }
}

// 组件挂载时加载数据
onMounted(() => {
  loadUserStats()
})
</script>

<style scoped lang="scss">
.personalized-welcome {
  padding: var(--spacing-xl);
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
  margin-bottom: var(--spacing-xl);
}

.welcome-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--spacing-xl);
  gap: var(--spacing-lg);
}

.greeting-section {
  flex: 1;
}

.greeting {
  font-size: var(--font-size-h1);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-sm) 0;
}

.user-name {
  font-size: var(--font-size-large);
  color: var(--text-secondary);
  margin: 0;
}

.identity-section {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: var(--spacing-md);
}

.identity-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.identity-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.identity-name {
  font-size: var(--font-size-h3);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0;
}

.identity-meta {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
}

.identity-tag {
  font-weight: var(--font-weight-medium);
}

.identity-domain {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
}

.identity-switcher {
  .el-button {
    color: var(--primary-color);
    font-size: var(--font-size-small);
  }
}

.stats-section {
  margin-bottom: var(--spacing-xl);
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: var(--spacing-lg);
}

.stat-card {
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

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: var(--icon-size-lg);
  color: white;
  flex-shrink: 0;
  
  &.apprentice {
    background: linear-gradient(135deg, var(--apprentice-color), #66bb6a);
  }
  
  &.master {
    background: linear-gradient(135deg, var(--master-color), #ff9800);
  }
  
  &.general {
    background: linear-gradient(135deg, var(--primary-color), #2196f3);
  }
  
  &.identity {
    background: linear-gradient(135deg, #9c27b0, #e91e63);
  }
}

.stat-content {
  flex: 1;
  min-width: 0;
}

.stat-title {
  font-size: var(--font-size-medium);
  font-weight: var(--font-weight-medium);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-xs) 0;
}

.stat-value {
  font-size: var(--font-size-h2);
  font-weight: var(--font-weight-bold);
  color: var(--primary-color);
  margin-bottom: var(--spacing-xs);
}

.stat-label {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-sm);
}

.stat-progress {
  margin-top: var(--spacing-sm);
}

.stat-secondary {
  display: flex;
  gap: var(--spacing-md);
  flex-wrap: wrap;
}

.secondary-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: var(--font-size-small);
  color: var(--text-tertiary);
  
  .el-icon {
    font-size: var(--icon-size-sm);
  }
}

.quick-actions {
  border-top: 1px solid var(--bg-tertiary);
  padding-top: var(--spacing-xl);
}

.section-title {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-lg) 0;
}

.action-buttons {
  display: flex;
  gap: var(--spacing-md);
  flex-wrap: wrap;
}

.action-btn {
  min-width: 140px;
  height: 50px;
  font-size: var(--font-size-medium);
  font-weight: var(--font-weight-medium);
}

.identity-list {
  max-height: 300px;
  overflow-y: auto;
}

.identity-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  cursor: pointer;
  transition: all var(--transition-normal);
  border-radius: var(--border-radius-medium);
  margin-bottom: var(--spacing-sm);
  
  &:hover {
    background: var(--bg-secondary);
  }
  
  &.is-active {
    background: rgba(64, 158, 255, 0.1);
    border-left: 3px solid var(--primary-color);
  }
  
  &:last-child {
    margin-bottom: 0;
  }
}

.item-info {
  flex: 1;
  min-width: 0;
}

.item-name {
  font-size: var(--font-size-medium);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
}

.item-meta {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
}

.item-domain {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
}

.item-action {
  flex-shrink: 0;
}

.check-icon {
  color: var(--primary-color);
  font-size: 16px;
}

// 响应式设计
@media (max-width: 768px) {
  .personalized-welcome {
    padding: var(--spacing-lg);
    margin: var(--spacing-md);
  }
  
  .welcome-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-md);
  }
  
  .identity-section {
    align-items: flex-start;
    width: 100%;
  }
  
  .identity-info {
    width: 100%;
    justify-content: space-between;
  }
  
  .stats-grid {
    grid-template-columns: 1fr;
    gap: var(--spacing-md);
  }
  
  .stat-card {
    padding: var(--spacing-md);
  }
  
  .action-buttons {
    flex-direction: column;
  }
  
  .action-btn {
    width: 100%;
  }
}
</style>