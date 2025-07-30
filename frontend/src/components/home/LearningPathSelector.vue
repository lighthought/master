<template>
  <div class="learning-path-selector">
    <div class="section-header">
      <h2 class="section-title">选择你的学习路径</h2>
      <p class="section-subtitle">根据你的学习偏好和目标，选择最适合的学习方式</p>
    </div>
    
    <div class="learning-paths-grid">
      <!-- 1对1指导 -->
      <div 
        class="path-card"
        :class="{ 'recommended': recommendedPath === 'one-on-one' }"
        @click="selectPath('one-on-one')"
      >
        <div class="path-icon">
          <el-icon><User /></el-icon>
        </div>
        <div class="path-content">
          <h3 class="path-title">1对1指导</h3>
          <p class="path-description">直接预约大师进行一对一指导</p>
          <div class="path-features">
            <span class="feature-tag">个性化</span>
            <span class="feature-tag">实时反馈</span>
            <span class="feature-tag">灵活时间</span>
          </div>
        </div>
        <div class="path-action">
          <el-button type="primary" size="small" @click.stop="navigateToPath('one-on-one')">
            开始指导
          </el-button>
        </div>
        <div v-if="recommendedPath === 'one-on-one'" class="recommended-badge">
          <el-icon><Star /></el-icon>
          推荐
        </div>
      </div>
      
      <!-- 结构化学习 -->
      <div 
        class="path-card"
        :class="{ 'recommended': recommendedPath === 'structured' }"
        @click="selectPath('structured')"
      >
        <div class="path-icon">
          <el-icon><Reading /></el-icon>
        </div>
        <div class="path-content">
          <h3 class="path-title">结构化学习</h3>
          <p class="path-description">报名大师设计的系统化课程</p>
          <div class="path-features">
            <span class="feature-tag">系统化</span>
            <span class="feature-tag">循序渐进</span>
            <span class="feature-tag">完整体系</span>
          </div>
        </div>
        <div class="path-action">
          <el-button type="success" size="small" @click.stop="navigateToPath('structured')">
            查看课程
          </el-button>
        </div>
        <div v-if="recommendedPath === 'structured'" class="recommended-badge">
          <el-icon><Star /></el-icon>
          推荐
        </div>
      </div>
      
      <!-- 浏览大师 -->
      <div 
        class="path-card"
        :class="{ 'recommended': recommendedPath === 'browse' }"
        @click="selectPath('browse')"
      >
        <div class="path-icon">
          <el-icon><Search /></el-icon>
        </div>
        <div class="path-content">
          <h3 class="path-title">浏览大师</h3>
          <p class="path-description">浏览和筛选平台上的大师</p>
          <div class="path-features">
            <span class="feature-tag">丰富选择</span>
            <span class="feature-tag">对比筛选</span>
            <span class="feature-tag">评价参考</span>
          </div>
        </div>
        <div class="path-action">
          <el-button type="warning" size="small" @click.stop="navigateToPath('browse')">
            浏览大师
          </el-button>
        </div>
        <div v-if="recommendedPath === 'browse'" class="recommended-badge">
          <el-icon><Star /></el-icon>
          推荐
        </div>
      </div>
      
      <!-- 其他方式 -->
      <div 
        class="path-card"
        :class="{ 'recommended': recommendedPath === 'other' }"
        @click="selectPath('other')"
      >
        <div class="path-icon">
          <el-icon><More /></el-icon>
        </div>
        <div class="path-content">
          <h3 class="path-title">其他方式</h3>
          <p class="path-description">探索更多学习途径</p>
          <div class="path-features">
            <span class="feature-tag">多样化</span>
            <span class="feature-tag">创新方式</span>
            <span class="feature-tag">灵活选择</span>
          </div>
        </div>
        <div class="path-action">
          <el-button type="info" size="small" @click.stop="navigateToPath('other')">
            探索更多
          </el-button>
        </div>
        <div v-if="recommendedPath === 'other'" class="recommended-badge">
          <el-icon><Star /></el-icon>
          推荐
        </div>
      </div>
    </div>
    
    <!-- 学习偏好设置 -->
    <div class="preferences-section">
      <div class="preferences-header">
        <h3 class="preferences-title">学习偏好设置</h3>
        <el-button type="text" size="small" @click="showPreferencesDialog = true">
          <el-icon><Setting /></el-icon>
          调整偏好
        </el-button>
      </div>
      <div class="preferences-summary">
        <div class="preference-item">
          <span class="preference-label">学习方式：</span>
          <span class="preference-value">{{ getPreferenceLabel(userPreferences.learningStyle) }}</span>
        </div>
        <div class="preference-item">
          <span class="preference-label">时间安排：</span>
          <span class="preference-value">{{ getPreferenceLabel(userPreferences.timePreference) }}</span>
        </div>
        <div class="preference-item">
          <span class="preference-label">预算范围：</span>
          <span class="preference-value">{{ getPreferenceLabel(userPreferences.budgetRange) }}</span>
        </div>
      </div>
    </div>
    
    <!-- 学习偏好设置对话框 -->
    <el-dialog
      v-model="showPreferencesDialog"
      title="学习偏好设置"
      width="90%"
      max-width="600px"
      :close-on-click-modal="false"
    >
      <div class="preferences-form">
        <el-form :model="userPreferences" label-width="100px">
          <el-form-item label="学习方式">
            <el-radio-group v-model="userPreferences.learningStyle">
              <el-radio label="one-on-one">1对1指导</el-radio>
              <el-radio label="structured">结构化学习</el-radio>
              <el-radio label="browse">浏览大师</el-radio>
              <el-radio label="other">其他方式</el-radio>
            </el-radio-group>
          </el-form-item>
          
          <el-form-item label="时间安排">
            <el-radio-group v-model="userPreferences.timePreference">
              <el-radio label="flexible">灵活时间</el-radio>
              <el-radio label="fixed">固定时间</el-radio>
              <el-radio label="weekend">周末学习</el-radio>
              <el-radio label="evening">晚上学习</el-radio>
            </el-radio-group>
          </el-form-item>
          
          <el-form-item label="预算范围">
            <el-radio-group v-model="userPreferences.budgetRange">
              <el-radio label="low">经济型 (¥100-500)</el-radio>
              <el-radio label="medium">标准型 (¥500-2000)</el-radio>
              <el-radio label="high">高端型 (¥2000+)</el-radio>
              <el-radio label="flexible">灵活预算</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-form>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showPreferencesDialog = false">取消</el-button>
          <el-button type="primary" @click="savePreferences">保存偏好</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  User, Reading, Search, More, Star, Setting 
} from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'

// 路由
const router = useRouter()

// 认证store
const authStore = useAuthStore()

// 状态
const showPreferencesDialog = ref(false)
const selectedPath = ref('')

// 用户偏好
const userPreferences = ref({
  learningStyle: 'one-on-one',
  timePreference: 'flexible',
  budgetRange: 'medium'
})

// 计算推荐路径
const recommendedPath = computed(() => {
  return userPreferences.value.learningStyle
})

// 选择学习路径
const selectPath = (path: string) => {
  selectedPath.value = path
  ElMessage.success(`已选择${getPathTitle(path)}学习路径`)
}

// 导航到指定路径
const navigateToPath = (path: string) => {
  switch (path) {
    case 'one-on-one':
      router.push('/mentors?type=one-on-one')
      break
    case 'structured':
      router.push('/courses')
      break
    case 'browse':
      router.push('/mentors')
      break
    case 'other':
      router.push('/explore')
      break
    default:
      ElMessage.warning('该功能正在开发中')
  }
}

// 获取路径标题
const getPathTitle = (path: string) => {
  const titles: Record<string, string> = {
    'one-on-one': '1对1指导',
    'structured': '结构化学习',
    'browse': '浏览大师',
    'other': '其他方式'
  }
  return titles[path] || '未知路径'
}

// 获取偏好标签
const getPreferenceLabel = (value: string) => {
  const labels: Record<string, string> = {
    // 学习方式
    'one-on-one': '1对1指导',
    'structured': '结构化学习',
    'browse': '浏览大师',
    'other': '其他方式',
    // 时间安排
    'flexible': '灵活时间',
    'fixed': '固定时间',
    'weekend': '周末学习',
    'evening': '晚上学习',
    // 预算范围
    'low': '经济型 (¥100-500)',
    'medium': '标准型 (¥500-2000)',
    'high': '高端型 (¥2000+)',
    'flexible-budget': '灵活预算'
  }
  return labels[value] || value
}

// 保存偏好设置
const savePreferences = async () => {
  if (!authStore.user) return
  
  try {
    await ApiService.userPreferences.saveUserPreferences(authStore.user.id, userPreferences.value)
    
    ElMessage.success('学习偏好设置已保存')
    showPreferencesDialog.value = false
    
    // 重新计算推荐路径
    await loadRecommendedPath()
  } catch (error) {
    ElMessage.error('保存失败，请重试')
  }
}

// 加载推荐路径
const loadRecommendedPath = async () => {
  if (!authStore.user) return
  
  try {
    const result = await ApiService.userPreferences.getRecommendedLearningPath(authStore.user.id)
    // 这里可以根据推荐结果更新UI
    console.log('推荐路径:', result.data)
  } catch (error) {
    console.error('加载推荐路径失败:', error)
  }
}

// 加载用户偏好
const loadUserPreferences = async () => {
  if (!authStore.user) return
  
  try {
    const result = await ApiService.userPreferences.getUserPreferences(authStore.user.id)
    userPreferences.value = result.data
  } catch (error) {
    console.error('加载用户偏好失败:', error)
    // 使用默认偏好
    userPreferences.value = {
      learningStyle: 'one-on-one',
      timePreference: 'flexible',
      budgetRange: 'medium'
    }
  }
}

// 组件挂载时加载数据
onMounted(() => {
  loadUserPreferences()
  loadRecommendedPath()
})
</script>

<style scoped lang="scss">
.learning-path-selector {
  padding: var(--spacing-xl);
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
  margin-bottom: var(--spacing-xl);
}

.section-header {
  text-align: center;
  margin-bottom: var(--spacing-xxl);
}

.section-title {
  font-size: var(--font-size-h2);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-sm) 0;
}

.section-subtitle {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin: 0;
}

.learning-paths-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-xxl);
}

.path-card {
  position: relative;
  padding: var(--spacing-xl);
  background: #2D2D2D;
  border: 1px solid #404040;
  border-radius: var(--border-radius-large);
  cursor: pointer;
  transition: all var(--transition-normal);
  
  &:hover {
    transform: translateY(-4px);
    box-shadow: var(--shadow-light);
    border-color: var(--primary-color);
  }
  
  &.recommended {
    border-color: var(--primary-color);
    background: linear-gradient(135deg, #2D2D2D, #1a1a1a);
    
    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      height: 3px;
      background: linear-gradient(90deg, var(--primary-color), var(--apprentice-color));
      border-radius: var(--border-radius-large) var(--border-radius-large) 0 0;
    }
  }
}

.path-icon {
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, var(--primary-color), var(--apprentice-color));
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: var(--icon-size-xl);
  color: white;
  margin-bottom: var(--spacing-lg);
}

.path-content {
  margin-bottom: var(--spacing-lg);
}

.path-title {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-sm) 0;
}

.path-description {
  font-size: var(--font-size-medium);
  color: #CCCCCC;
  margin: 0 0 var(--spacing-md) 0;
  line-height: 1.5;
}

.path-features {
  display: flex;
  gap: var(--spacing-xs);
  flex-wrap: wrap;
}

.feature-tag {
  padding: var(--spacing-xs) var(--spacing-sm);
  background: rgba(64, 158, 255, 0.1);
  color: var(--primary-color);
  border-radius: var(--border-radius-small);
  font-size: var(--font-size-small);
  font-weight: var(--font-weight-medium);
}

.path-action {
  display: flex;
  justify-content: flex-end;
}

.recommended-badge {
  position: absolute;
  top: var(--spacing-md);
  right: var(--spacing-md);
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  padding: var(--spacing-xs) var(--spacing-sm);
  background: linear-gradient(135deg, var(--primary-color), var(--apprentice-color));
  color: white;
  border-radius: var(--border-radius-small);
  font-size: var(--font-size-small);
  font-weight: var(--font-weight-medium);
  
  .el-icon {
    font-size: var(--icon-size-sm);
  }
}

.preferences-section {
  border-top: 1px solid var(--bg-tertiary);
  padding-top: var(--spacing-xl);
}

.preferences-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
}

.preferences-title {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0;
}

.preferences-summary {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-md);
}

.preference-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-md);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
}

.preference-label {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  font-weight: var(--font-weight-medium);
}

.preference-value {
  font-size: var(--font-size-small);
  color: var(--text-primary);
  font-weight: var(--font-weight-semibold);
}

.preferences-form {
  padding: var(--spacing-lg) 0;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
}

// 响应式设计
@media (max-width: 768px) {
  .learning-path-selector {
    padding: var(--spacing-lg);
    margin: var(--spacing-md);
  }
  
  .learning-paths-grid {
    grid-template-columns: 1fr;
    gap: var(--spacing-md);
  }
  
  .path-card {
    padding: var(--spacing-lg);
  }
  
  .preferences-summary {
    grid-template-columns: 1fr;
  }
  
  .preferences-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-sm);
  }
}
</style>