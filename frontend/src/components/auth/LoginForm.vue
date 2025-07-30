<template>
  <div class="login-form">
    <div class="form-header">
      <h2 class="form-title">欢迎回来</h2>
      <p class="form-subtitle">登录你的Master Guide账号</p>
    </div>
    
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-position="top"
      class="login-form-content"
      @submit.prevent="handleSubmit"
    >
      <el-form-item label="邮箱地址" prop="email">
        <el-input
          v-model="formData.email"
          type="email"
          placeholder="请输入邮箱地址"
          prefix-icon="Message"
          size="large"
        />
      </el-form-item>
      
      <el-form-item label="密码" prop="password">
        <el-input
          v-model="formData.password"
          type="password"
          placeholder="请输入密码"
          prefix-icon="Lock"
          size="large"
          show-password
        />
      </el-form-item>
      
      <div class="form-options">
        <el-checkbox v-model="formData.rememberMe">记住我</el-checkbox>
        <el-link type="primary" @click="forgotPassword">忘记密码？</el-link>
      </div>
      
      <el-form-item>
        <el-button
          type="primary"
          size="large"
          class="submit-btn"
          :loading="loading"
          @click="handleSubmit"
        >
          {{ loading ? '登录中...' : '立即登录' }}
        </el-button>
      </el-form-item>
    </el-form>
    
    <div class="form-footer">
      <span class="register-link">
        还没有账号？
        <el-link type="primary" @click="$emit('switch-to-register')">立即注册</el-link>
      </span>
      <div class="test-account">
        <p class="test-tip">测试账号：test@example.com / 123456</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { useAuthStore } from '@/stores/auth'

// 定义事件
const emit = defineEmits<{
  'switch-to-register': []
  'login-success': [userData: any]
}>()

// 表单引用
const formRef = ref<FormInstance>()

// 加载状态
const loading = ref(false)

// 认证store
const authStore = useAuthStore()

// 表单数据
const formData = reactive({
  email: '',
  password: '',
  rememberMe: false
})

// 表单验证规则
const formRules: FormRules = {
  email: [
    { required: true, message: '请输入邮箱地址', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ]
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value?.validate()
    
    loading.value = true
    
    // 调用认证store的登录方法
    const result = await authStore.login(formData.email, formData.password)
    
    // 登录成功
    ElMessage.success('登录成功！欢迎回来')
    
    // 触发成功事件
    emit('login-success', {
      email: formData.email,
      rememberMe: formData.rememberMe,
      user: result.user
    })
    
  } catch (error) {
    console.error('登录失败:', error)
    ElMessage.error('登录失败，请检查邮箱和密码')
  } finally {
    loading.value = false
  }
}

// 忘记密码
const forgotPassword = () => {
  ElMessage.info('忘记密码功能开发中...')
}
</script>

<style scoped lang="scss">
.login-form {
  max-width: 480px;
  margin: 0 auto;
  padding: var(--spacing-xl);
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
}

.form-header {
  text-align: center;
  margin-bottom: var(--spacing-xl);
}

.form-title {
  font-size: var(--font-size-h2);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-sm);
}

.form-subtitle {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  line-height: 1.5;
}

.login-form-content {
  margin-bottom: var(--spacing-lg);
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
}

.submit-btn {
  width: 100%;
  height: 48px;
  font-size: var(--font-size-large);
  font-weight: var(--font-weight-semibold);
}

.form-footer {
  text-align: center;
  padding-top: var(--spacing-lg);
  border-top: 1px solid var(--bg-tertiary);
}

.register-link {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
}

.test-account {
  margin-top: var(--spacing-md);
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--bg-tertiary);
}

.test-tip {
  font-size: var(--font-size-small);
  color: var(--text-tertiary);
  text-align: center;
  margin: 0;
}

// 响应式设计
@media (max-width: 768px) {
  .login-form {
    margin: var(--spacing-md);
    padding: var(--spacing-lg);
  }
}
</style>