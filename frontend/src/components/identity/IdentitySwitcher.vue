<template>
  <div class="identity-switcher">
    <!-- 当前身份显示 -->
    <div class="current-identity" @click="toggleDropdown">
      <div class="identity-info">
        <div class="identity-avatar">
          <el-avatar 
            :size="32" 
            :src="currentIdentity?.avatar"
            :icon="getIdentityIcon(currentIdentity?.type)"
          />
        </div>
        <div class="identity-details">
          <div class="identity-name">{{ currentIdentity?.name || '未选择身份' }}</div>
          <div class="identity-type">
            <el-tag 
              :type="currentIdentity?.type === 'master' ? 'success' : 'info'"
              size="small"
            >
              {{ currentIdentity?.type === 'master' ? '大师' : '学徒' }}
            </el-tag>
            <span class="identity-domain">{{ currentIdentity?.domain }}</span>
          </div>
        </div>
      </div>
      <div class="identity-arrow">
        <el-icon :class="{ 'is-open': showDropdown }">
          <ArrowDown />
        </el-icon>
      </div>
    </div>
    
    <!-- 身份列表下拉框 -->
    <transition name="dropdown">
      <div v-if="showDropdown" class="identity-dropdown">
        <div class="dropdown-header">
          <h4>选择身份</h4>
          <el-button 
            type="text" 
            size="small" 
            @click="$router.push('/identity')"
          >
            管理身份
          </el-button>
        </div>
        
        <div class="identity-list">
          <div
            v-for="identity in userIdentities"
            :key="identity.id"
            class="identity-item"
            :class="{ 'is-active': identity.id === currentIdentity?.id }"
            @click="switchToIdentity(identity)"
          >
            <div class="item-avatar">
              <el-avatar 
                :size="40" 
                :src="identity.avatar"
                :icon="getIdentityIcon(identity.type)"
              />
            </div>
            <div class="item-info">
              <div class="item-name">{{ identity.name }}</div>
              <div class="item-meta">
                <el-tag 
                  :type="identity.type === 'master' ? 'success' : 'info'"
                  size="small"
                >
                  {{ identity.type === 'master' ? '大师' : '学徒' }}
                </el-tag>
                <span class="item-domain">{{ identity.domain }}</span>
              </div>
              <div class="item-status">
                <el-tag 
                  v-if="identity.status === 'pending'"
                  type="warning" 
                  size="small"
                >
                  审核中
                </el-tag>
                <el-tag 
                  v-else-if="identity.isActive"
                  type="success" 
                  size="small"
                >
                  已激活
                </el-tag>
                <el-tag 
                  v-else
                  type="info" 
                  size="small"
                >
                  未激活
                </el-tag>
              </div>
            </div>
            <div class="item-action">
              <el-icon v-if="identity.id === currentIdentity?.id" class="check-icon">
                <Check />
              </el-icon>
            </div>
          </div>
        </div>
        
        <!-- 创建新身份 -->
        <div class="create-identity">
          <el-button 
            type="primary" 
            size="small" 
            @click="createNewIdentity"
            class="create-btn"
          >
            <el-icon><Plus /></el-icon>
            创建新身份
          </el-button>
        </div>
      </div>
    </transition>
    
    <!-- 遮罩层 -->
    <div 
      v-if="showDropdown" 
      class="dropdown-overlay"
      @click="closeDropdown"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ArrowDown, Check, Plus } from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import type { Identity } from '@/types/user'

// 定义事件
const emit = defineEmits<{
  'identity-changed': [identity: Identity]
}>()

// 路由
const router = useRouter()

// 认证store
const authStore = useAuthStore()

// 下拉框状态
const showDropdown = ref(false)

// 计算属性
const currentIdentity = computed(() => authStore.currentIdentity)
const userIdentities = computed(() => authStore.userIdentities)

// 获取身份图标
const getIdentityIcon = (type?: string) => {
  if (type === 'master') {
    return 'Star'
  } else if (type === 'apprentice') {
    return 'User'
  }
  return 'User'
}

// 切换下拉框
const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
}

// 关闭下拉框
const closeDropdown = () => {
  showDropdown.value = false
}

// 切换到指定身份
const switchToIdentity = async (identity: Identity) => {
  if (identity.id === currentIdentity.value?.id) {
    closeDropdown()
    return
  }
  
  if (!identity.isActive) {
    ElMessage.warning('该身份尚未激活，无法切换')
    return
  }
  
  try {
    await authStore.switchIdentity(identity.id)
    ElMessage.success(`已切换到${identity.name}身份`)
    emit('identity-changed', identity)
    closeDropdown()
  } catch (error) {
    ElMessage.error('身份切换失败，请重试')
  }
}

// 创建新身份
const createNewIdentity = () => {
  closeDropdown()
  router.push('/identity')
}

// 点击外部关闭下拉框
const handleClickOutside = (event: Event) => {
  const target = event.target as HTMLElement
  if (!target.closest('.identity-switcher')) {
    closeDropdown()
  }
}

// 键盘事件处理
const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape') {
    closeDropdown()
  }
}

// 生命周期
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  document.removeEventListener('keydown', handleKeydown)
})
</script>

<style scoped lang="scss">
.identity-switcher {
  position: relative;
  display: inline-block;
}

.current-identity {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm) var(--spacing-md);
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius-medium);
  cursor: pointer;
  transition: all var(--transition-normal);
  min-width: 200px;
  
  &:hover {
    border-color: var(--primary-color);
    box-shadow: var(--shadow-small);
  }
}

.identity-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  flex: 1;
}

.identity-avatar {
  flex-shrink: 0;
}

.identity-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
  min-width: 0;
}

.identity-name {
  font-size: var(--font-size-medium);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.identity-type {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
}

.identity-domain {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.identity-arrow {
  flex-shrink: 0;
  transition: transform var(--transition-normal);
  
  .is-open {
    transform: rotate(180deg);
  }
}

.identity-dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius-medium);
  box-shadow: var(--shadow-large);
  z-index: 1000;
  margin-top: var(--spacing-xs);
  min-width: 280px;
}

.dropdown-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md);
  border-bottom: 1px solid var(--bg-tertiary);
  
  h4 {
    margin: 0;
    font-size: var(--font-size-medium);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
  }
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
  border-bottom: 1px solid var(--bg-tertiary);
  
  &:hover {
    background: var(--bg-secondary);
  }
  
  &:last-child {
    border-bottom: none;
  }
  
  &.is-active {
    background: rgba(64, 158, 255, 0.1);
    border-left: 3px solid var(--primary-color);
  }
}

.item-avatar {
  flex-shrink: 0;
}

.item-info {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
  flex: 1;
  min-width: 0;
}

.item-name {
  font-size: var(--font-size-medium);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
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

.item-status {
  margin-top: var(--spacing-xs);
}

.item-action {
  flex-shrink: 0;
}

.check-icon {
  color: var(--primary-color);
  font-size: 16px;
}

.create-identity {
  padding: var(--spacing-md);
  border-top: 1px solid var(--bg-tertiary);
  text-align: center;
}

.create-btn {
  width: 100%;
}

.dropdown-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 999;
}

// 下拉动画
.dropdown-enter-active,
.dropdown-leave-active {
  transition: all var(--transition-normal);
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

// 响应式设计
@media (max-width: 768px) {
  .current-identity {
    min-width: 160px;
    padding: var(--spacing-xs) var(--spacing-sm);
  }
  
  .identity-name {
    font-size: var(--font-size-small);
  }
  
  .identity-dropdown {
    min-width: 240px;
  }
  
  .identity-item {
    padding: var(--spacing-sm);
  }
}
</style>