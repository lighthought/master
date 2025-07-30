<template>
  <div class="register-form">
    <div class="form-header">
      <h2 class="form-title">创建账号</h2>
      <p class="form-subtitle">加入Master Guide，开启你的技艺传承之旅</p>
    </div>
    
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-position="top"
      class="register-form-content"
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
          placeholder="请输入密码（至少8位）"
          prefix-icon="Lock"
          size="large"
          show-password
        />
      </el-form-item>
      
      <el-form-item label="确认密码" prop="confirmPassword">
        <el-input
          v-model="formData.confirmPassword"
          type="password"
          placeholder="请再次输入密码"
          prefix-icon="Lock"
          size="large"
          show-password
        />
      </el-form-item>
      
      <el-form-item label="主要身份" prop="primaryIdentity">
        <el-radio-group v-model="formData.primaryIdentity" class="identity-group">
          <el-radio value="master" class="identity-radio">
            <div class="identity-option">
              <el-icon class="identity-icon"><User /></el-icon>
              <div class="identity-content">
                <span class="identity-title">大师</span>
                <span class="identity-desc">提供专业指导服务</span>
              </div>
            </div>
          </el-radio>
          <el-radio value="apprentice" class="identity-radio">
            <div class="identity-option">
              <el-icon class="identity-icon"><Reading /></el-icon>
              <div class="identity-content">
                <span class="identity-title">学徒</span>
                <span class="identity-desc">学习专业技能</span>
              </div>
            </div>
          </el-radio>
        </el-radio-group>
      </el-form-item>
      
      <el-form-item>
        <el-checkbox v-model="formData.agreeTerms" class="terms-checkbox">
          我已阅读并同意
          <el-link type="primary" @click="showTerms">《用户协议》</el-link>
          和
          <el-link type="primary" @click="showPrivacy">《隐私政策》</el-link>
        </el-checkbox>
      </el-form-item>
      
      <el-form-item>
        <el-button
          type="primary"
          size="large"
          class="submit-btn"
          :loading="loading"
          @click="handleSubmit"
        >
          {{ loading ? '注册中...' : '立即注册' }}
        </el-button>
      </el-form-item>
    </el-form>
    
    <div class="form-footer">
      <span class="login-link">
        已有账号？
        <el-link type="primary" @click="$emit('switch-to-login')">立即登录</el-link>
      </span>
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
  'switch-to-login': []
  'register-success': [userData: any]
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
  confirmPassword: '',
  primaryIdentity: 'apprentice',
  agreeTerms: false
})

// 表单验证规则
const formRules: FormRules = {
  email: [
    { required: true, message: '请输入邮箱地址', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 8, message: '密码长度至少8位', trigger: 'blur' },
    { 
      pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/,
      message: '密码必须包含大小写字母和数字',
      trigger: 'blur'
    }
  ],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== formData.password) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
  primaryIdentity: [
    { required: true, message: '请选择主要身份', trigger: 'change' }
  ]
}

// 提交表单
const handleSubmit = async () => {
  if (!formData.agreeTerms) {
    ElMessage.warning('请先同意用户协议和隐私政策')
    return
  }
  
  try {
    await formRef.value?.validate()
    
    loading.value = true
    
    // 调用认证store的注册方法
    const result = await authStore.register({
      email: formData.email,
      password: formData.password,
      primaryIdentity: formData.primaryIdentity as 'master' | 'apprentice'
    })
    
    // 注册成功
    ElMessage.success('注册成功！欢迎加入Master Guide')
    
    // 触发成功事件
    emit('register-success', {
      email: formData.email,
      primaryIdentity: formData.primaryIdentity,
      user: result.user
    })
    
  } catch (error) {
    console.error('注册失败:', error)
    ElMessage.error('注册失败，请检查输入信息')
  } finally {
    loading.value = false
  }
}

// 显示用户协议
const showTerms = () => {
  ElMessage.info('用户协议功能开发中...')
}

// 显示隐私政策
const showPrivacy = () => {
  ElMessage.info('隐私政策功能开发中...')
}
</script>

<style scoped lang="scss">
.register-form {
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

.register-form-content {
  margin-bottom: var(--spacing-lg);
}

.identity-group {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
  width: 100%;
}

.identity-radio {
  width: 100%;
  margin-right: 0;
  margin-bottom: var(--spacing-sm);
  
  :deep(.el-radio__label) {
    width: 100%;
    padding-left: var(--spacing-sm);
  }
  
  :deep(.el-radio__input) {
    margin-right: var(--spacing-sm);
  }
}

.identity-option {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  border: 2px solid var(--bg-tertiary);
  border-radius: var(--border-radius-medium);
  transition: all var(--transition-normal);
  cursor: pointer;
  
  &:hover {
    border-color: var(--primary-color);
    background: var(--bg-secondary);
  }
}

.identity-radio.is-checked .identity-option {
  border-color: var(--primary-color);
  background: rgba(255, 107, 53, 0.1);
}

.identity-icon {
  font-size: var(--icon-size-lg);
  color: var(--primary-color);
  flex-shrink: 0;
}

.identity-content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.identity-title {
  font-size: var(--font-size-h5);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
}

.identity-desc {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
}

.terms-checkbox {
  color: var(--text-secondary);
  font-size: var(--font-size-medium);
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

.login-link {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
}

// 响应式设计
@media (max-width: 768px) {
  .register-form {
    margin: var(--spacing-md);
    padding: var(--spacing-lg);
  }
  
  .identity-option {
    padding: var(--spacing-sm);
  }
  
  .identity-icon {
    font-size: var(--icon-size-md);
  }
}
</style>