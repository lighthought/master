<template>
  <div class="create-apprentice-identity">
    <div class="form-header">
      <h2 class="form-title">创建学徒身份</h2>
      <p class="form-subtitle">完善你的学习信息，开始专业技能学习之旅</p>
    </div>
    
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-position="top"
      class="identity-form"
      @submit.prevent="handleSubmit"
    >
      <!-- 基本信息 -->
      <div class="form-section">
        <h3 class="section-title">基本信息</h3>
        
        <el-form-item label="学徒姓名" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入你的学徒姓名"
            size="large"
          />
        </el-form-item>
        
        <el-form-item label="学习领域" prop="domain">
          <el-select
            v-model="formData.domain"
            placeholder="请选择学习领域"
            size="large"
            class="domain-select"
          >
            <el-option
              v-for="domain in domainOptions"
              :key="domain.value"
              :label="domain.label"
              :value="domain.value"
            />
          </el-select>
        </el-form-item>
      </div>
      
      <!-- 学习背景 -->
      <div class="form-section">
        <h3 class="section-title">学习背景</h3>
        
        <el-form-item label="学习背景" prop="background">
          <el-input
            v-model="formData.background"
            type="textarea"
            :rows="4"
            placeholder="请描述你的学习背景、工作经验或相关技能..."
            maxlength="500"
            show-word-limit
          />
        </el-form-item>
        
        <el-form-item label="当前水平" prop="currentLevel">
          <el-radio-group v-model="formData.currentLevel" class="level-group">
            <el-radio value="beginner" class="level-radio">
              <div class="level-option">
                <div class="level-icon">🌱</div>
                <div class="level-content">
                  <span class="level-title">初学者</span>
                  <span class="level-desc">完全新手，需要从基础开始</span>
                </div>
              </div>
            </el-radio>
            <el-radio value="intermediate" class="level-radio">
              <div class="level-option">
                <div class="level-icon">🌿</div>
                <div class="level-content">
                  <span class="level-title">进阶者</span>
                  <span class="level-desc">有一定基础，需要提升技能</span>
                </div>
              </div>
            </el-radio>
            <el-radio value="advanced" class="level-radio">
              <div class="level-option">
                <div class="level-icon">🌳</div>
                <div class="level-content">
                  <span class="level-title">高级者</span>
                  <span class="level-desc">经验丰富，需要精进专业</span>
                </div>
              </div>
            </el-radio>
          </el-radio-group>
        </el-form-item>
      </div>
      
      <!-- 学习目标 -->
      <div class="form-section">
        <h3 class="section-title">学习目标</h3>
        
        <el-form-item label="学习目标" prop="learningGoals">
          <el-select
            v-model="formData.learningGoals"
            multiple
            filterable
            allow-create
            default-first-option
            placeholder="请选择或输入学习目标"
            size="large"
            class="goals-select"
          >
            <el-option
              v-for="goal in learningGoalOptions"
              :key="goal"
              :label="goal"
              :value="goal"
            />
          </el-select>
        </el-form-item>
        
        <el-form-item label="期望学习时间" prop="expectedDuration">
          <el-radio-group v-model="formData.expectedDuration" class="duration-group">
            <el-radio value="1-3months">1-3个月</el-radio>
            <el-radio value="3-6months">3-6个月</el-radio>
            <el-radio value="6-12months">6-12个月</el-radio>
            <el-radio value="1year+">1年以上</el-radio>
          </el-radio-group>
        </el-form-item>
      </div>
      
      <!-- 学习偏好 -->
      <div class="form-section">
        <h3 class="section-title">学习偏好</h3>
        
        <el-form-item label="学习方式偏好" prop="learningPreferences">
          <el-checkbox-group v-model="formData.learningPreferences" class="preferences-group">
            <el-checkbox value="one-on-one">1对1指导</el-checkbox>
            <el-checkbox value="group-class">小组课程</el-checkbox>
            <el-checkbox value="self-study">自主学习</el-checkbox>
            <el-checkbox value="project-based">项目实践</el-checkbox>
            <el-checkbox value="mentorship">导师制</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        
        <el-form-item label="时间安排偏好" prop="timePreferences">
          <el-checkbox-group v-model="formData.timePreferences" class="time-group">
            <el-checkbox value="weekday-morning">工作日早上</el-checkbox>
            <el-checkbox value="weekday-evening">工作日晚上</el-checkbox>
            <el-checkbox value="weekend">周末</el-checkbox>
            <el-checkbox value="flexible">时间灵活</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        
        <el-form-item label="预算范围" prop="budgetRange">
          <el-select
            v-model="formData.budgetRange"
            placeholder="请选择预算范围"
            size="large"
            class="budget-select"
          >
            <el-option label="1000元以下" value="under-1000" />
            <el-option label="1000-3000元" value="1000-3000" />
            <el-option label="3000-5000元" value="3000-5000" />
            <el-option label="5000-10000元" value="5000-10000" />
            <el-option label="10000元以上" value="over-10000" />
          </el-select>
        </el-form-item>
      </div>
      
      <!-- 提交按钮 -->
      <div class="form-actions">
        <el-button @click="$emit('cancel')" size="large">取消</el-button>
        <el-button
          type="primary"
          size="large"
          :loading="loading"
          @click="handleSubmit"
        >
          {{ loading ? '创建中...' : '立即创建' }}
        </el-button>
      </div>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { useAuthStore } from '@/stores/auth'

// 定义事件
const emit = defineEmits<{
  'cancel': []
  'submit-success': [identityData: any]
}>()

// 表单引用
const formRef = ref<FormInstance>()

// 加载状态
const loading = ref(false)

// 认证store
const authStore = useAuthStore()

// 学习领域选项
const domainOptions = [
  { value: 'software-development', label: '软件开发' },
  { value: 'ui-design', label: 'UI设计' },
  { value: 'digital-marketing', label: '数字营销' },
  { value: 'data-analysis', label: '数据分析' },
  { value: 'product-management', label: '产品管理' },
  { value: 'business-strategy', label: '商业策略' },
  { value: 'finance', label: '金融投资' },
  { value: 'language-teaching', label: '语言教学' },
  { value: 'music', label: '音乐艺术' },
  { value: 'fitness', label: '健身指导' },
  { value: 'cooking', label: '烹饪技艺' },
  { value: 'other', label: '其他领域' }
]

// 学习目标选项
const learningGoalOptions = [
  '掌握基础技能', '提升专业技能', '转行换岗', '职业发展',
  '兴趣爱好', '创业准备', '项目实战', '技术认证',
  '团队管理', '领导力提升', '沟通技巧', '时间管理',
  '创新思维', '问题解决', '数据分析', '市场洞察'
]

// 表单数据
const formData = reactive({
  name: '',
  domain: '',
  background: '',
  currentLevel: 'beginner',
  learningGoals: [] as string[],
  expectedDuration: '3-6months',
  learningPreferences: ['one-on-one', 'self-study'] as string[],
  timePreferences: ['weekday-evening', 'weekend'] as string[],
  budgetRange: '1000-3000'
})

// 表单验证规则
const formRules: FormRules = {
  name: [
    { required: true, message: '请输入学徒姓名', trigger: 'blur' },
    { min: 2, max: 20, message: '姓名长度在2-20个字符', trigger: 'blur' }
  ],
  domain: [
    { required: true, message: '请选择学习领域', trigger: 'change' }
  ],
  background: [
    { required: true, message: '请填写学习背景', trigger: 'blur' },
    { min: 30, max: 500, message: '学习背景长度在30-500个字符', trigger: 'blur' }
  ],
  currentLevel: [
    { required: true, message: '请选择当前水平', trigger: 'change' }
  ],
  learningGoals: [
    { required: true, message: '请选择学习目标', trigger: 'change' },
    { type: 'array', min: 1, max: 8, message: '请选择1-8个学习目标', trigger: 'change' }
  ],
  expectedDuration: [
    { required: true, message: '请选择期望学习时间', trigger: 'change' }
  ],
  learningPreferences: [
    { required: true, message: '请选择学习方式偏好', trigger: 'change' },
    { type: 'array', min: 1, max: 5, message: '请选择1-5种学习方式', trigger: 'change' }
  ],
  timePreferences: [
    { required: true, message: '请选择时间安排偏好', trigger: 'change' },
    { type: 'array', min: 1, max: 4, message: '请选择1-4种时间安排', trigger: 'change' }
  ],
  budgetRange: [
    { required: true, message: '请选择预算范围', trigger: 'change' }
  ]
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value?.validate()
    
    loading.value = true
    
    // 调用认证store的创建学徒身份方法
    const result = await authStore.createApprenticeIdentity(formData)
    
    // 创建成功
    ElMessage.success('学徒身份创建成功！')
    
    // 触发成功事件
    emit('submit-success', {
      ...formData,
      status: 'active',
      identity: result
    })
    
  } catch (error) {
    console.error('创建失败:', error)
    ElMessage.error('创建失败，请检查表单信息')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped lang="scss">
.create-apprentice-identity {
  max-width: 800px;
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

.identity-form {
  margin-bottom: var(--spacing-lg);
}

.form-section {
  margin-bottom: var(--spacing-xl);
  padding-bottom: var(--spacing-xl);
  border-bottom: 1px solid var(--bg-tertiary);
  
  &:last-child {
    border-bottom: none;
    margin-bottom: 0;
    padding-bottom: 0;
  }
}

.section-title {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-lg);
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  
  &::before {
    content: '';
    width: 4px;
    height: 20px;
    background: var(--apprentice-color);
    border-radius: 2px;
  }
}

.domain-select,
.goals-select,
.budget-select {
  width: 100%;
}

.level-group {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
  width: 100%;
}

.level-radio {
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

.level-option {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  border: 2px solid var(--bg-tertiary);
  border-radius: var(--border-radius-medium);
  transition: all var(--transition-normal);
  cursor: pointer;
  
  &:hover {
    border-color: var(--apprentice-color);
    background: var(--bg-secondary);
  }
}

.level-radio.is-checked .level-option {
  border-color: var(--apprentice-color);
  background: rgba(76, 175, 80, 0.1);
}

.level-icon {
  font-size: 24px;
  flex-shrink: 0;
}

.level-content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.level-title {
  font-size: var(--font-size-h5);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
}

.level-desc {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
}

.duration-group {
  display: flex;
  gap: var(--spacing-lg);
  flex-wrap: wrap;
}

.preferences-group,
.time-group {
  display: flex;
  gap: var(--spacing-md);
  flex-wrap: wrap;
}

.form-actions {
  display: flex;
  justify-content: center;
  gap: var(--spacing-lg);
  padding-top: var(--spacing-xl);
  border-top: 1px solid var(--bg-tertiary);
}

// 响应式设计
@media (max-width: 768px) {
  .create-apprentice-identity {
    margin: var(--spacing-md);
    padding: var(--spacing-lg);
  }
  
  .level-option {
    padding: var(--spacing-sm);
  }
  
  .level-icon {
    font-size: 20px;
  }
  
  .duration-group,
  .preferences-group,
  .time-group {
    flex-direction: column;
    gap: var(--spacing-sm);
  }
  
  .form-actions {
    flex-direction: column;
    gap: var(--spacing-md);
  }
}
</style>