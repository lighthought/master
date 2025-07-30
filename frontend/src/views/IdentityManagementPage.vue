<template>
  <div class="identity-management-page">
    <div class="page-header">
      <h1 class="page-title">身份管理</h1>
      <p class="page-subtitle">管理你的多重身份，在不同角色间自由切换</p>
    </div>
    
    <!-- 当前身份信息 -->
    <div class="current-identity-section">
      <h2 class="section-title">当前身份</h2>
      <div class="current-identity-card">
        <div class="identity-info">
          <div class="identity-avatar">
            <el-avatar :size="80" :src="currentIdentity?.avatar" />
            <div class="identity-badge" :class="currentIdentity?.type">
              {{ getIdentityIcon(currentIdentity?.type) }}
            </div>
          </div>
          <div class="identity-details">
            <h3 class="identity-name">{{ currentIdentity?.name || '未设置' }}</h3>
            <p class="identity-domain">{{ currentIdentity?.domain || '未选择领域' }}</p>
            <el-tag :type="getIdentityType(currentIdentity?.type)" size="large">
              {{ getIdentityLabel(currentIdentity?.type) }}
            </el-tag>
            <div class="identity-status">
              <el-tag 
                v-if="currentIdentity?.isVerified" 
                type="success" 
                size="small"
              >
                已认证
              </el-tag>
              <el-tag 
                v-else-if="currentIdentity?.status === 'pending'" 
                type="warning" 
                size="small"
              >
                审核中
              </el-tag>
              <el-tag 
                v-else 
                type="info" 
                size="small"
              >
                未认证
              </el-tag>
            </div>
          </div>
        </div>
        <div class="identity-actions">
          <el-button type="primary" @click="editCurrentIdentity">
            编辑身份
          </el-button>
        </div>
      </div>
    </div>
    
    <!-- 所有身份列表 -->
    <div class="identities-section">
      <div class="section-header">
        <h2 class="section-title">所有身份</h2>
        <el-button type="primary" @click="showCreateIdentity">
          <el-icon><Plus /></el-icon>
          创建新身份
        </el-button>
      </div>
      
      <div class="identities-grid">
        <div 
          v-for="identity in userIdentities" 
          :key="identity.id"
          class="identity-card"
          :class="{ active: identity.id === currentIdentity?.id }"
        >
          <div class="identity-header">
            <div class="identity-avatar">
              <el-avatar :size="60" :src="identity.avatar" />
              <div class="identity-badge" :class="identity.type">
                {{ getIdentityIcon(identity.type) }}
              </div>
            </div>
            <div class="identity-status">
              <el-tag 
                v-if="identity.isVerified" 
                type="success" 
                size="small"
              >
                已认证
              </el-tag>
              <el-tag 
                v-else-if="identity.status === 'pending'" 
                type="warning" 
                size="small"
              >
                审核中
              </el-tag>
              <el-tag 
                v-else 
                type="info" 
                size="small"
              >
                未认证
              </el-tag>
            </div>
          </div>
          
          <div class="identity-content">
            <h4 class="identity-name">{{ identity.name }}</h4>
            <p class="identity-domain">{{ identity.domain }}</p>
            <el-tag :type="getIdentityType(identity.type)" size="small">
              {{ getIdentityLabel(identity.type) }}
            </el-tag>
          </div>
          
          <div class="identity-actions">
            <el-button 
              v-if="identity.id !== currentIdentity?.id"
              type="primary" 
              size="small"
              @click="switchToIdentity(identity)"
            >
              切换到此身份
            </el-button>
            <el-button 
              v-else
              type="info" 
              size="small"
              disabled
            >
              当前身份
            </el-button>
            <el-button 
              type="text" 
              size="small"
              @click="editIdentity(identity)"
            >
              编辑
            </el-button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 创建身份对话框 -->
    <el-dialog
      v-model="showCreateDialog"
      title="创建新身份"
      width="90%"
      max-width="900px"
      :close-on-click-modal="false"
    >
      <div class="identity-type-selection">
        <h3>选择身份类型</h3>
        <div class="identity-types">
          <div 
            class="identity-type-card"
            :class="{ selected: selectedIdentityType === 'master' }"
            @click="selectIdentityType('master')"
          >
            <div class="type-icon">
              <el-icon><User /></el-icon>
            </div>
            <h4>大师身份</h4>
            <p>提供专业指导服务，分享你的技能和经验</p>
            <ul>
              <li>设置指导价格</li>
              <li>上传资质证明</li>
              <li>接受学员预约</li>
              <li>获得收入分成</li>
            </ul>
          </div>
          
          <div 
            class="identity-type-card"
            :class="{ selected: selectedIdentityType === 'apprentice' }"
            @click="selectIdentityType('apprentice')"
          >
            <div class="type-icon">
              <el-icon><Reading /></el-icon>
            </div>
            <h4>学徒身份</h4>
            <p>学习专业技能，寻找合适的导师指导</p>
            <ul>
              <li>选择学习领域</li>
              <li>设置学习目标</li>
              <li>预约大师指导</li>
              <li>参与课程学习</li>
            </ul>
          </div>
        </div>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showCreateDialog = false">取消</el-button>
          <el-button 
            type="primary" 
            :disabled="!selectedIdentityType"
            @click="createIdentity"
          >
            继续创建
          </el-button>
        </div>
      </template>
    </el-dialog>
    
    <!-- 创建大师身份对话框 -->
    <el-dialog
      v-model="showMasterDialog"
      title="创建大师身份"
      width="90%"
      max-width="900px"
      :close-on-click-modal="false"
      :show-close="false"
    >
      <CreateMasterIdentity
        @cancel="showMasterDialog = false"
        @submit-success="handleMasterCreated"
      />
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ElMessage } from 'element-plus'
import CreateMasterIdentity from '@/components/identity/CreateMasterIdentity.vue'
import type { Identity } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

// 对话框状态
const showCreateDialog = ref(false)
const showMasterDialog = ref(false)
const selectedIdentityType = ref<'master' | 'apprentice' | null>(null)

// 计算属性
const currentIdentity = computed(() => authStore.currentIdentity)
const userIdentities = computed(() => authStore.userIdentities)

// 获取身份图标
const getIdentityIcon = (type?: string) => {
  return type === 'master' ? '👨‍🏫' : '👨‍🎓'
}

// 获取身份类型标签
const getIdentityType = (type?: string) => {
  return type === 'master' ? 'warning' : 'success'
}

// 获取身份标签文本
const getIdentityLabel = (type?: string) => {
  return type === 'master' ? '大师' : '学徒'
}

// 显示创建身份对话框
const showCreateIdentity = () => {
  selectedIdentityType.value = null
  showCreateDialog.value = true
}

// 选择身份类型
const selectIdentityType = (type: 'master' | 'apprentice') => {
  selectedIdentityType.value = type
}

// 创建身份
const createIdentity = () => {
  if (selectedIdentityType.value === 'master') {
    showCreateDialog.value = false
    showMasterDialog.value = true
  } else if (selectedIdentityType.value === 'apprentice') {
    // TODO: 实现创建学徒身份
    ElMessage.info('创建学徒身份功能开发中...')
  }
}

// 处理大师身份创建成功
const handleMasterCreated = async (identityData: any) => {
  try {
    await authStore.createMasterIdentity(identityData)
    showMasterDialog.value = false
    ElMessage.success('大师身份创建成功！正在审核中，请耐心等待')
  } catch (error) {
    ElMessage.error('创建失败，请重试')
  }
}

// 编辑当前身份
const editCurrentIdentity = () => {
  if (currentIdentity.value) {
    editIdentity(currentIdentity.value)
  }
}

// 编辑身份
const editIdentity = (identity: Identity) => {
  ElMessage.info('编辑身份功能开发中...')
}

// 切换到指定身份
const switchToIdentity = (identity: Identity) => {
  try {
    authStore.switchIdentity(identity.id)
    ElMessage.success(`已切换到${getIdentityLabel(identity.type)}身份`)
  } catch (error) {
    ElMessage.error('切换身份失败')
  }
}
</script>

<style scoped lang="scss">
.identity-management-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: var(--spacing-xl);
}

.page-header {
  text-align: center;
  margin-bottom: var(--spacing-xxl);
}

.page-title {
  font-size: var(--font-size-h1);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-sm);
}

.page-subtitle {
  font-size: var(--font-size-large);
  color: var(--text-secondary);
  line-height: 1.5;
}

.current-identity-section {
  margin-bottom: var(--spacing-xxl);
}

.section-title {
  font-size: var(--font-size-h3);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-lg);
}

.current-identity-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  padding: var(--spacing-xl);
  box-shadow: var(--shadow-card);
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: var(--spacing-xl);
}

.identity-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-lg);
}

.identity-avatar {
  position: relative;
}

.identity-badge {
  position: absolute;
  bottom: -4px;
  right: -4px;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  border: 3px solid var(--bg-card);
}

.identity-badge.master {
  background: var(--master-color);
}

.identity-badge.apprentice {
  background: var(--apprentice-color);
}

.identity-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.identity-name {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0;
}

.identity-domain {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin: 0;
}

.identity-status {
  display: flex;
  gap: var(--spacing-xs);
}

.identities-section {
  margin-bottom: var(--spacing-xxl);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
}

.identities-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: var(--spacing-lg);
}

.identity-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  padding: var(--spacing-lg);
  box-shadow: var(--shadow-card);
  transition: all var(--transition-normal);
  border: 2px solid transparent;
  
  &:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-medium);
  }
  
  &.active {
    border-color: var(--primary-color);
    background: rgba(255, 107, 53, 0.05);
  }
}

.identity-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--spacing-md);
}

.identity-content {
  margin-bottom: var(--spacing-md);
}

.identity-content .identity-name {
  font-size: var(--font-size-h5);
  margin-bottom: var(--spacing-xs);
}

.identity-content .identity-domain {
  font-size: var(--font-size-small);
  margin-bottom: var(--spacing-sm);
}

.identity-actions {
  display: flex;
  gap: var(--spacing-sm);
}

.identity-type-selection {
  padding: var(--spacing-lg);
}

.identity-type-selection h3 {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-lg);
  text-align: center;
}

.identity-types {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--spacing-xl);
}

.identity-type-card {
  background: var(--bg-secondary);
  border: 2px solid var(--bg-tertiary);
  border-radius: var(--border-radius-large);
  padding: var(--spacing-xl);
  text-align: center;
  cursor: pointer;
  transition: all var(--transition-normal);
  
  &:hover {
    border-color: var(--primary-color);
    background: var(--bg-card);
  }
  
  &.selected {
    border-color: var(--primary-color);
    background: rgba(255, 107, 53, 0.1);
  }
}

.type-icon {
  width: 80px;
  height: 80px;
  margin: 0 auto var(--spacing-lg);
  background: var(--primary-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 40px;
  color: var(--text-primary);
}

.identity-type-card h4 {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-sm);
}

.identity-type-card p {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-lg);
  line-height: 1.5;
}

.identity-type-card ul {
  text-align: left;
  list-style: none;
  padding: 0;
  margin: 0;
}

.identity-type-card li {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-xs);
  padding-left: var(--spacing-md);
  position: relative;
  
  &::before {
    content: '✓';
    position: absolute;
    left: 0;
    color: var(--primary-color);
    font-weight: var(--font-weight-bold);
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
}

// 响应式设计
@media (max-width: 768px) {
  .identity-management-page {
    padding: var(--spacing-lg);
  }
  
  .current-identity-card {
    flex-direction: column;
    text-align: center;
    gap: var(--spacing-lg);
  }
  
  .identities-grid {
    grid-template-columns: 1fr;
  }
  
  .identity-types {
    grid-template-columns: 1fr;
    gap: var(--spacing-lg);
  }
  
  .section-header {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: stretch;
  }
}
</style>