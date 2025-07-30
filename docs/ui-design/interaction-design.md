# Master Guide 交互设计规范

## 1. 用户交互流程

### 1.1 身份切换流程
```
用户点击身份切换器
    ↓
显示身份列表抽屉
    ↓
用户选择目标身份
    ↓
验证身份权限
    ↓
切换身份状态
    ↓
更新UI和权限
    ↓
显示切换成功提示
```

### 1.2 大师预约流程
```
用户浏览大师列表
    ↓
点击大师卡片
    ↓
查看大师详情
    ↓
点击预约按钮
    ↓
选择预约时间
    ↓
填写预约需求
    ↓
确认预约信息
    ↓
支付预约费用
    ↓
预约成功通知
```

### 1.3 课程学习流程
```
用户浏览课程列表
    ↓
点击课程卡片
    ↓
查看课程详情
    ↓
点击报名按钮
    ↓
确认报名信息
    ↓
支付课程费用
    ↓
进入学习页面
    ↓
观看课程内容
    ↓
完成学习任务
    ↓
获得学习证书
```

## 2. 微交互设计

### 2.1 按钮交互
```scss
// 按钮点击反馈
.button {
  transition: all 0.15s ease-out;
  
  &:active {
    transform: scale(0.98);
  }
  
  &:hover {
    box-shadow: 0 4px 12px rgba(255, 107, 53, 0.3);
  }
}

// 加载状态
.button--loading {
  position: relative;
  pointer-events: none;
  
  &::after {
    content: '';
    position: absolute;
    top: 50%;
    left: 50%;
    width: 16px;
    height: 16px;
    margin: -8px 0 0 -8px;
    border: 2px solid transparent;
    border-top: 2px solid currentColor;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }
}

@keyframes spin {
  to { transform: rotate(360deg); }
}
```

### 2.2 卡片交互
```scss
// 卡片悬停效果
.card {
  transition: all 0.25s ease-out;
  
  &:hover {
    transform: translateY(-4px);
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  }
}

// 卡片点击反馈
.card--clickable {
  cursor: pointer;
  
  &:active {
    transform: translateY(-2px);
  }
}
```

### 2.3 输入框交互
```scss
// 输入框聚焦效果
.input {
  transition: all 0.2s ease-out;
  
  &:focus {
    border-color: var(--primary-color);
    box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.1);
  }
}

// 输入框验证状态
.input--error {
  border-color: var(--danger-color);
  animation: shake 0.5s ease-in-out;
}

@keyframes shake {
  0%, 100% { transform: translateX(0); }
  25% { transform: translateX(-4px); }
  75% { transform: translateX(4px); }
}
```

## 3. 页面转场动画

### 3.1 路由转场
```vue
<template>
  <router-view v-slot="{ Component }">
    <transition name="page" mode="out-in">
      <component :is="Component" />
    </transition>
  </router-view>
</template>

<style>
.page-enter-active,
.page-leave-active {
  transition: all 0.3s ease-out;
}

.page-enter-from {
  opacity: 0;
  transform: translateX(20px);
}

.page-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}
</style>
```

### 3.2 模态框动画
```scss
// 模态框显示动画
.modal-enter-active,
.modal-leave-active {
  transition: all 0.3s ease-out;
}

.modal-enter-from {
  opacity: 0;
  transform: scale(0.9);
}

.modal-leave-to {
  opacity: 0;
  transform: scale(1.1);
}

// 遮罩层动画
.overlay-enter-active,
.overlay-leave-active {
  transition: opacity 0.3s ease-out;
}

.overlay-enter-from,
.overlay-leave-to {
  opacity: 0;
}
```

## 4. 加载状态设计

### 4.1 骨架屏
```vue
<template>
  <div class="skeleton-card">
    <div class="skeleton-header">
      <div class="skeleton-avatar"></div>
      <div class="skeleton-info">
        <div class="skeleton-title"></div>
        <div class="skeleton-subtitle"></div>
      </div>
    </div>
    <div class="skeleton-content">
      <div class="skeleton-line"></div>
      <div class="skeleton-line"></div>
      <div class="skeleton-line short"></div>
    </div>
  </div>
</template>

<style scoped>
.skeleton-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  padding: var(--card-padding);
}

.skeleton-header {
  display: flex;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-md);
}

.skeleton-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: linear-gradient(90deg, #404040 25%, #666666 50%, #404040 75%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
}

.skeleton-info {
  flex: 1;
}

.skeleton-title {
  height: 20px;
  background: linear-gradient(90deg, #404040 25%, #666666 50%, #404040 75%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
  margin-bottom: var(--spacing-xs);
  border-radius: var(--border-radius-small);
}

.skeleton-subtitle {
  height: 16px;
  width: 60%;
  background: linear-gradient(90deg, #404040 25%, #666666 50%, #404040 75%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
  border-radius: var(--border-radius-small);
}

.skeleton-content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.skeleton-line {
  height: 16px;
  background: linear-gradient(90deg, #404040 25%, #666666 50%, #404040 75%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
  border-radius: var(--border-radius-small);
}

.skeleton-line.short {
  width: 70%;
}

@keyframes shimmer {
  0% { background-position: -200% 0; }
  100% { background-position: 200% 0; }
}
</style>
```

### 4.2 进度指示器
```vue
<template>
  <div class="progress-indicator">
    <div class="progress-bar">
      <div 
        class="progress-fill" 
        :style="{ width: `${percentage}%` }"
      ></div>
    </div>
    <div class="progress-text">{{ percentage }}%</div>
  </div>
</template>

<script setup>
defineProps({
  percentage: {
    type: Number,
    default: 0
  }
})
</script>

<style scoped>
.progress-indicator {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
}

.progress-bar {
  flex: 1;
  height: 8px;
  background: var(--bg-tertiary);
  border-radius: var(--border-radius-full);
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--primary-color), var(--primary-light));
  border-radius: var(--border-radius-full);
  transition: width 0.3s ease-out;
}

.progress-text {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  min-width: 40px;
  text-align: right;
}
</style>
```

## 5. 手势交互

### 5.1 滑动操作
```vue
<template>
  <div 
    class="swipeable-item"
    @touchstart="handleTouchStart"
    @touchmove="handleTouchMove"
    @touchend="handleTouchEnd"
  >
    <div 
      class="item-content"
      :style="{ transform: `translateX(${translateX}px)` }"
    >
      <slot />
    </div>
    <div class="swipe-actions">
      <button class="action-btn edit" @click="handleEdit">编辑</button>
      <button class="action-btn delete" @click="handleDelete">删除</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const translateX = ref(0)
let startX = 0
let currentX = 0

const handleTouchStart = (event) => {
  startX = event.touches[0].clientX
}

const handleTouchMove = (event) => {
  currentX = event.touches[0].clientX
  const diff = currentX - startX
  
  if (diff < 0) {
    translateX.value = Math.max(diff, -120)
  }
}

const handleTouchEnd = () => {
  if (translateX.value < -60) {
    translateX.value = -120
  } else {
    translateX.value = 0
  }
}

const handleEdit = () => {
  // 处理编辑操作
}

const handleDelete = () => {
  // 处理删除操作
}
</script>

<style scoped>
.swipeable-item {
  position: relative;
  overflow: hidden;
}

.item-content {
  background: var(--bg-card);
  transition: transform 0.3s ease-out;
}

.swipe-actions {
  position: absolute;
  top: 0;
  right: 0;
  height: 100%;
  display: flex;
}

.action-btn {
  width: 60px;
  border: none;
  color: white;
  font-size: var(--font-size-small);
  cursor: pointer;
  transition: background-color 0.2s ease-out;
}

.action-btn.edit {
  background: var(--info-color);
}

.action-btn.edit:hover {
  background: var(--info-color-dark);
}

.action-btn.delete {
  background: var(--danger-color);
}

.action-btn.delete:hover {
  background: var(--danger-color-dark);
}
</style>
```

### 5.2 下拉刷新
```vue
<template>
  <div 
    class="pull-refresh"
    @touchstart="handleTouchStart"
    @touchmove="handleTouchMove"
    @touchend="handleTouchEnd"
  >
    <div 
      class="refresh-indicator"
      :class="{ active: isRefreshing }"
      :style="{ transform: `translateY(${translateY}px)` }"
    >
      <el-icon v-if="!isRefreshing" class="refresh-icon">
        <ArrowDown />
      </el-icon>
      <el-icon v-else class="refresh-icon loading">
        <Loading />
      </el-icon>
      <span class="refresh-text">
        {{ isRefreshing ? '正在刷新...' : '下拉刷新' }}
      </span>
    </div>
    
    <div class="content" :style="{ transform: `translateY(${translateY}px)` }">
      <slot />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const translateY = ref(0)
const isRefreshing = ref(false)
let startY = 0
let currentY = 0

const handleTouchStart = (event) => {
  startY = event.touches[0].clientY
}

const handleTouchMove = (event) => {
  currentY = event.touches[0].clientY
  const diff = currentY - startY
  
  if (diff > 0 && window.scrollY === 0) {
    translateY.value = Math.min(diff * 0.5, 80)
  }
}

const handleTouchEnd = async () => {
  if (translateY.value > 40) {
    isRefreshing.value = true
    translateY.value = 60
    
    try {
      await refresh()
    } finally {
      isRefreshing.value = false
      translateY.value = 0
    }
  } else {
    translateY.value = 0
  }
}

const refresh = async () => {
  // 执行刷新逻辑
  await new Promise(resolve => setTimeout(resolve, 1000))
}
</script>

<style scoped>
.pull-refresh {
  position: relative;
  overflow: hidden;
}

.refresh-indicator {
  position: absolute;
  top: -60px;
  left: 0;
  right: 0;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--spacing-sm);
  background: var(--bg-secondary);
  transition: transform 0.3s ease-out;
}

.refresh-icon {
  font-size: var(--icon-size-md);
  color: var(--text-secondary);
  transition: transform 0.3s ease-out;
}

.refresh-icon.loading {
  animation: spin 1s linear infinite;
}

.refresh-text {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
}

.content {
  transition: transform 0.3s ease-out;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}
</style>
```

## 6. 反馈设计

### 6.1 成功反馈
```vue
<template>
  <div class="feedback-toast" :class="type">
    <el-icon class="feedback-icon">
      <component :is="icon" />
    </el-icon>
    <span class="feedback-text">{{ message }}</span>
  </div>
</template>

<script setup>
const props = defineProps({
  type: {
    type: String,
    default: 'success'
  },
  message: {
    type: String,
    required: true
  }
})

const icon = computed(() => {
  const icons = {
    success: 'Check',
    error: 'Close',
    warning: 'Warning',
    info: 'InfoFilled'
  }
  return icons[props.type]
})
</script>

<style scoped>
.feedback-toast {
  position: fixed;
  top: 20px;
  right: 20px;
  z-index: 9999;
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-md) var(--spacing-lg);
  border-radius: var(--border-radius-medium);
  color: white;
  font-size: var(--font-size-medium);
  animation: slideIn 0.3s ease-out;
}

.feedback-toast.success {
  background: var(--success-color);
}

.feedback-toast.error {
  background: var(--danger-color);
}

.feedback-toast.warning {
  background: var(--warning-color);
}

.feedback-toast.info {
  background: var(--info-color);
}

.feedback-icon {
  font-size: var(--icon-size-md);
}

@keyframes slideIn {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}
</style>
```

### 6.2 错误处理
```vue
<template>
  <div class="error-boundary" v-if="hasError">
    <div class="error-content">
      <el-icon class="error-icon">
        <Warning />
      </el-icon>
      <h3 class="error-title">出错了</h3>
      <p class="error-message">{{ errorMessage }}</p>
      <div class="error-actions">
        <el-button @click="retry">重试</el-button>
        <el-button @click="goHome">返回首页</el-button>
      </div>
    </div>
  </div>
  <slot v-else />
</template>

<script setup>
import { ref, onErrorCaptured } from 'vue'

const hasError = ref(false)
const errorMessage = ref('')

onErrorCaptured((error) => {
  hasError.value = true
  errorMessage.value = error.message || '发生未知错误'
  return false
})

const retry = () => {
  hasError.value = false
  errorMessage.value = ''
  window.location.reload()
}

const goHome = () => {
  router.push('/')
}
</script>

<style scoped>
.error-boundary {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  padding: var(--spacing-xl);
}

.error-content {
  text-align: center;
  max-width: 400px;
}

.error-icon {
  font-size: 64px;
  color: var(--warning-color);
  margin-bottom: var(--spacing-lg);
}

.error-title {
  font-size: var(--font-size-h3);
  color: var(--text-primary);
  margin-bottom: var(--spacing-md);
}

.error-message {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-lg);
  line-height: 1.5;
}

.error-actions {
  display: flex;
  gap: var(--spacing-md);
  justify-content: center;
}
</style>
```

---

**文档版本**：v1.0.0  
**创建日期**：2024年12月  
**设计负责人**：Sally (UX Expert) 