# Master Guide UI 组件库

## 1. 基础组件

### 1.1 按钮组件 (Button)
```vue
<template>
  <button 
    :class="buttonClasses" 
    :disabled="disabled"
    @click="handleClick"
  >
    <el-icon v-if="icon" class="button-icon">
      <component :is="icon" />
    </el-icon>
    <span class="button-text">{{ text }}</span>
  </button>
</template>

<script setup>
const props = defineProps({
  type: { type: String, default: 'primary' }, // primary, secondary, text, danger
  size: { type: String, default: 'medium' },  // small, medium, large
  disabled: { type: Boolean, default: false },
  loading: { type: Boolean, default: false },
  icon: { type: String, default: '' },
  text: { type: String, required: true }
})

const buttonClasses = computed(() => [
  'mg-button',
  `mg-button--${props.type}`,
  `mg-button--${props.size}`,
  { 'mg-button--disabled': props.disabled },
  { 'mg-button--loading': props.loading }
])
</script>

<style scoped>
.mg-button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border: none;
  border-radius: var(--border-radius-medium);
  font-weight: var(--font-weight-medium);
  cursor: pointer;
  transition: all var(--transition-normal);
  gap: var(--spacing-xs);
}

.mg-button--primary {
  background: var(--primary-color);
  color: var(--text-primary);
}

.mg-button--primary:hover {
  background: var(--primary-dark);
  box-shadow: var(--shadow-medium);
}

.mg-button--secondary {
  background: var(--bg-secondary);
  color: var(--text-primary);
  border: 1px solid var(--border-color);
}

.mg-button--text {
  background: transparent;
  color: var(--text-primary);
}

.mg-button--danger {
  background: var(--danger-color);
  color: var(--text-primary);
}

.mg-button--small {
  padding: 6px 12px;
  font-size: var(--font-size-small);
}

.mg-button--medium {
  padding: 8px 16px;
  font-size: var(--font-size-medium);
}

.mg-button--large {
  padding: 12px 24px;
  font-size: var(--font-size-large);
}

.mg-button--disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.mg-button--loading {
  pointer-events: none;
}
</style>
```

### 1.2 输入框组件 (Input)
```vue
<template>
  <div class="mg-input-wrapper">
    <div class="mg-input-container" :class="inputClasses">
      <el-icon v-if="prefixIcon" class="input-prefix-icon">
        <component :is="prefixIcon" />
      </el-icon>
      <input
        :value="modelValue"
        :type="type"
        :placeholder="placeholder"
        :disabled="disabled"
        class="mg-input"
        @input="handleInput"
        @focus="handleFocus"
        @blur="handleBlur"
      />
      <el-icon v-if="suffixIcon" class="input-suffix-icon">
        <component :is="suffixIcon" />
      </el-icon>
    </div>
    <div v-if="error" class="input-error">{{ error }}</div>
  </div>
</template>

<script setup>
const props = defineProps({
  modelValue: { type: String, default: '' },
  type: { type: String, default: 'text' },
  placeholder: { type: String, default: '' },
  disabled: { type: Boolean, default: false },
  error: { type: String, default: '' },
  prefixIcon: { type: String, default: '' },
  suffixIcon: { type: String, default: '' }
})

const emit = defineEmits(['update:modelValue', 'focus', 'blur'])

const inputClasses = computed(() => [
  'mg-input-container',
  { 'mg-input--error': props.error },
  { 'mg-input--disabled': props.disabled }
])

const handleInput = (event) => {
  emit('update:modelValue', event.target.value)
}

const handleFocus = (event) => {
  emit('focus', event)
}

const handleBlur = (event) => {
  emit('blur', event)
}
</script>

<style scoped>
.mg-input-wrapper {
  width: 100%;
}

.mg-input-container {
  display: flex;
  align-items: center;
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius-medium);
  padding: 12px 16px;
  transition: all var(--transition-normal);
}

.mg-input-container:focus-within {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 2px rgba(255, 107, 53, 0.2);
}

.mg-input--error {
  border-color: var(--danger-color);
  box-shadow: 0 0 0 2px rgba(244, 67, 54, 0.2);
}

.mg-input--disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.mg-input {
  flex: 1;
  background: transparent;
  border: none;
  outline: none;
  color: var(--text-primary);
  font-size: var(--font-size-medium);
}

.mg-input::placeholder {
  color: var(--text-tertiary);
}

.input-prefix-icon,
.input-suffix-icon {
  color: var(--text-tertiary);
  font-size: var(--icon-size-md);
}

.input-error {
  color: var(--danger-color);
  font-size: var(--font-size-small);
  margin-top: var(--spacing-xs);
}
</style>
```

### 1.3 卡片组件 (Card)
```vue
<template>
  <div class="mg-card" :class="cardClasses">
    <div v-if="$slots.header" class="card-header">
      <slot name="header" />
    </div>
    <div class="card-content">
      <slot />
    </div>
    <div v-if="$slots.footer" class="card-footer">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  shadow: { type: String, default: 'light' }, // light, medium, heavy
  hoverable: { type: Boolean, default: false },
  clickable: { type: Boolean, default: false }
})

const cardClasses = computed(() => [
  'mg-card',
  `mg-card--shadow-${props.shadow}`,
  { 'mg-card--hoverable': props.hoverable },
  { 'mg-card--clickable': props.clickable }
])
</script>

<style scoped>
.mg-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  overflow: hidden;
  transition: all var(--transition-normal);
}

.mg-card--shadow-light {
  box-shadow: var(--shadow-light);
}

.mg-card--shadow-medium {
  box-shadow: var(--shadow-medium);
}

.mg-card--shadow-heavy {
  box-shadow: var(--shadow-heavy);
}

.mg-card--hoverable:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-heavy);
}

.mg-card--clickable {
  cursor: pointer;
}

.mg-card--clickable:hover {
  transform: translateY(-1px);
  box-shadow: var(--shadow-medium);
}

.card-header {
  padding: var(--card-padding);
  border-bottom: 1px solid var(--border-color);
}

.card-content {
  padding: var(--card-padding);
}

.card-footer {
  padding: var(--card-padding);
  border-top: 1px solid var(--border-color);
  background: var(--bg-secondary);
}
</style>
```

## 2. 业务组件

### 2.1 身份切换器 (IdentitySwitcher)
```vue
<template>
  <div class="identity-switcher">
    <div class="current-identity" @click="showIdentityList">
      <div class="identity-avatar">
        <el-avatar :src="currentIdentity.avatar" :size="40" />
        <div class="identity-badge" :class="currentIdentity.type">
          {{ getIdentityIcon(currentIdentity.type) }}
        </div>
      </div>
      <div class="identity-info">
        <h4 class="identity-name">{{ currentIdentity.name }}</h4>
        <el-tag :type="getIdentityType(currentIdentity.type)" size="small">
          {{ getIdentityLabel(currentIdentity.type) }}
        </el-tag>
      </div>
      <el-icon class="switch-icon"><ArrowDown /></el-icon>
    </div>
    
    <el-drawer 
      v-model="showIdentityDrawer" 
      title="切换身份" 
      direction="rtl"
      size="300px"
    >
      <div class="identity-list">
        <div 
          v-for="identity in userIdentities" 
          :key="identity.id"
          class="identity-item"
          :class="{ active: identity.id === currentIdentity.id }"
          @click="switchToIdentity(identity)"
        >
          <div class="identity-avatar">
            <el-avatar :src="identity.avatar" :size="40" />
            <div class="identity-badge" :class="identity.type">
              {{ getIdentityIcon(identity.type) }}
            </div>
          </div>
          <div class="identity-info">
            <h4 class="identity-name">{{ identity.name }}</h4>
            <p class="identity-domain">{{ identity.domain }}</p>
            <el-tag :type="getIdentityType(identity.type)" size="small">
              {{ getIdentityLabel(identity.type) }}
            </el-tag>
          </div>
          <el-icon v-if="identity.id === currentIdentity.id" class="check-icon">
            <Check />
          </el-icon>
        </div>
        
        <el-button 
          type="dashed" 
          class="add-identity-btn"
          @click="addNewIdentity"
        >
          <el-icon><Plus /></el-icon>
          添加新身份
        </el-button>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const showIdentityDrawer = ref(false)

const currentIdentity = computed(() => authStore.currentIdentity)
const userIdentities = computed(() => authStore.userIdentities)

const getIdentityIcon = (type) => {
  return type === 'master' ? '👨‍🏫' : '👨‍🎓'
}

const getIdentityType = (type) => {
  return type === 'master' ? 'warning' : 'success'
}

const getIdentityLabel = (type) => {
  return type === 'master' ? '大师' : '学徒'
}

const showIdentityList = () => {
  showIdentityDrawer.value = true
}

const switchToIdentity = async (identity) => {
  try {
    await authStore.switchIdentity(identity.id)
    showIdentityDrawer.value = false
  } catch (error) {
    console.error('切换身份失败:', error)
  }
}

const addNewIdentity = () => {
  // 跳转到添加身份页面
  router.push('/profile/add-identity')
}
</script>

<style scoped>
.identity-switcher {
  position: relative;
}

.current-identity {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm);
  border-radius: var(--border-radius-medium);
  cursor: pointer;
  transition: all var(--transition-normal);
}

.current-identity:hover {
  background: var(--bg-secondary);
}

.identity-avatar {
  position: relative;
}

.identity-badge {
  position: absolute;
  bottom: -2px;
  right: -2px;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  border: 2px solid var(--bg-primary);
}

.identity-badge.master {
  background: var(--master-color);
}

.identity-badge.apprentice {
  background: var(--apprentice-color);
}

.identity-info {
  flex: 1;
  min-width: 0;
}

.identity-name {
  margin: 0 0 var(--spacing-xs) 0;
  font-size: var(--font-size-medium);
  font-weight: var(--font-weight-medium);
  color: var(--text-primary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.identity-domain {
  margin: 0 0 var(--spacing-xs) 0;
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.switch-icon {
  color: var(--text-tertiary);
  transition: transform var(--transition-normal);
}

.current-identity:hover .switch-icon {
  transform: rotate(180deg);
}

.identity-list {
  padding: var(--spacing-md);
}

.identity-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-md);
  border-radius: var(--border-radius-medium);
  cursor: pointer;
  transition: all var(--transition-normal);
  margin-bottom: var(--spacing-sm);
}

.identity-item:hover {
  background: var(--bg-secondary);
}

.identity-item.active {
  background: var(--primary-color);
  color: var(--text-primary);
}

.identity-item.active .identity-name,
.identity-item.active .identity-domain {
  color: var(--text-primary);
}

.check-icon {
  color: var(--text-primary);
  margin-left: auto;
}

.add-identity-btn {
  width: 100%;
  margin-top: var(--spacing-md);
}
</style>
```

### 2.2 大师卡片 (MentorCard)
```vue
<template>
  <div class="mentor-card" @click="viewMentorDetail">
    <div class="mentor-header">
      <div class="mentor-avatar">
        <el-avatar :src="mentor.avatar" :size="60" />
        <div class="online-status" :class="{ online: mentor.isOnline }"></div>
      </div>
      <div class="mentor-info">
        <h3 class="mentor-name">{{ mentor.name }}</h3>
        <div class="mentor-badges">
          <el-tag v-if="mentor.isVerified" type="success" size="small">
            已认证
          </el-tag>
          <el-tag v-if="mentor.isOnline" type="primary" size="small">
            在线
          </el-tag>
        </div>
      </div>
      <div class="mentor-rating">
        <el-rate v-model="mentor.rating" disabled />
        <span class="rating-text">{{ mentor.rating }}分</span>
      </div>
    </div>
    
    <div class="mentor-details">
      <p class="mentor-domain">{{ mentor.domain }}</p>
      <p class="mentor-description">{{ mentor.description }}</p>
      
      <div class="mentor-stats">
        <span class="stat-item">
          <el-icon><User /></el-icon>
          {{ mentor.studentCount }} 学生
        </span>
        <span class="stat-item">
          <el-icon><Clock /></el-icon>
          {{ mentor.experience }} 年经验
        </span>
      </div>
    </div>
    
    <div class="mentor-footer">
      <div class="price-info">
        <span class="price">¥{{ mentor.price }}/小时</span>
      </div>
      <el-button type="primary" size="small" @click.stop="bookMentor">
        立即预约
      </el-button>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  mentor: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['view-detail', 'book'])

const viewMentorDetail = () => {
  emit('view-detail', props.mentor)
}

const bookMentor = () => {
  emit('book', props.mentor)
}
</script>

<style scoped>
.mentor-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  padding: var(--card-padding);
  cursor: pointer;
  transition: all var(--transition-normal);
  border: 1px solid var(--border-color);
}

.mentor-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-medium);
}

.mentor-header {
  display: flex;
  align-items: flex-start;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-md);
}

.mentor-avatar {
  position: relative;
  flex-shrink: 0;
}

.online-status {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: var(--text-tertiary);
  border: 2px solid var(--bg-card);
}

.online-status.online {
  background: var(--success-color);
}

.mentor-info {
  flex: 1;
  min-width: 0;
}

.mentor-name {
  margin: 0 0 var(--spacing-xs) 0;
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
}

.mentor-badges {
  display: flex;
  gap: var(--spacing-xs);
  margin-bottom: var(--spacing-xs);
}

.mentor-rating {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: var(--spacing-xs);
}

.rating-text {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
}

.mentor-details {
  margin-bottom: var(--spacing-md);
}

.mentor-domain {
  margin: 0 0 var(--spacing-xs) 0;
  font-size: var(--font-size-medium);
  color: var(--primary-color);
  font-weight: var(--font-weight-medium);
}

.mentor-description {
  margin: 0 0 var(--spacing-md) 0;
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.mentor-stats {
  display: flex;
  gap: var(--spacing-md);
}

.stat-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: var(--font-size-small);
  color: var(--text-tertiary);
}

.mentor-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: var(--spacing-md);
}

.price-info {
  flex: 1;
}

.price {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-bold);
  color: var(--primary-color);
}
</style>
```

## 3. 布局组件

### 3.1 应用布局 (AppLayout)
```vue
<template>
  <div class="app-layout">
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
          <div class="identity-switcher-container">
            <IdentitySwitcher />
          </div>
          
          <div class="notification-container">
            <NotificationBadge />
          </div>
          
          <div class="user-menu-container">
            <UserMenu />
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

<script setup>
import { ref } from 'vue'
import IdentitySwitcher from '@/components/IdentitySwitcher.vue'
import NotificationBadge from '@/components/NotificationBadge.vue'
import UserMenu from '@/components/UserMenu.vue'

const searchQuery = ref('')

const navItems = [
  { path: '/', text: '首页', icon: 'House' },
  { path: '/mentors', text: '大师', icon: 'User' },
  { path: '/courses', text: '课程', icon: 'Reading' },
  { path: '/community', text: '社群', icon: 'ChatDotRound' },
  { path: '/profile', text: '我的', icon: 'UserFilled' }
]
</script>

<style scoped>
.app-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: var(--bg-primary);
}

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

.app-main {
  flex: 1;
  margin-top: 64px;
  margin-bottom: 60px;
  padding: var(--spacing-lg);
  max-width: 1200px;
  margin-left: auto;
  margin-right: auto;
  width: 100%;
}

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
}

.nav-item:hover {
  color: var(--primary-color);
}

.nav-item.active {
  color: var(--primary-color);
}

.nav-icon {
  font-size: var(--icon-size-lg);
}

.nav-text {
  font-size: var(--font-size-small);
  font-weight: var(--font-weight-medium);
}

@media (min-width: 768px) {
  .app-bottom-nav {
    display: none;
  }
  
  .app-main {
    margin-bottom: 0;
  }
}
</style>
```

---

**文档版本**：v1.0.0  
**创建日期**：2024年12月  
**设计负责人**：Sally (UX Expert) 