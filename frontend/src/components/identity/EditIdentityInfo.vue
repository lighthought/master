<template>
  <div class="edit-identity-info">
    <div class="form-header">
      <h2 class="form-title">编辑身份信息</h2>
      <p class="form-subtitle">更新你的身份信息，让其他用户更好地了解你</p>
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
        
        <el-form-item label="身份名称" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入身份名称"
            size="large"
          />
        </el-form-item>
        
        <el-form-item label="专业领域" prop="domain">
          <el-select
            v-model="formData.domain"
            placeholder="请选择专业领域"
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
        
        <el-form-item label="头像" prop="avatar">
          <div class="avatar-upload">
            <el-upload
              class="avatar-uploader"
              :show-file-list="false"
              :before-upload="beforeAvatarUpload"
              :on-success="handleAvatarSuccess"
              :on-error="handleAvatarError"
            >
              <div class="avatar-preview">
                <el-avatar 
                  :size="80" 
                  :src="formData.avatar || identity?.avatar"
                  :icon="getIdentityIcon(identity?.type)"
                />
                <div class="upload-overlay">
                  <el-icon><Camera /></el-icon>
                  <span>更换头像</span>
                </div>
              </div>
            </el-upload>
            <div class="avatar-tip">
              <p>支持 JPG、PNG 格式，文件大小不超过 2MB</p>
            </div>
          </div>
        </el-form-item>
      </div>
      
      <!-- 个人介绍 -->
      <div class="form-section">
        <h3 class="section-title">个人介绍</h3>
        
        <el-form-item label="个人简介" prop="bio">
          <el-input
            v-model="formData.bio"
            type="textarea"
            :rows="4"
            placeholder="请介绍你的专业背景、技能特长、服务理念等..."
            maxlength="500"
            show-word-limit
          />
        </el-form-item>
        
        <el-form-item label="技能标签" prop="skills">
          <el-select
            v-model="formData.skills"
            multiple
            filterable
            allow-create
            default-first-option
            placeholder="请选择或输入技能标签"
            size="large"
            class="skills-select"
          >
            <el-option
              v-for="skill in skillOptions"
              :key="skill"
              :label="skill"
              :value="skill"
            />
          </el-select>
        </el-form-item>
      </div>
      
      <!-- 联系方式 -->
      <div class="form-section">
        <h3 class="section-title">联系方式</h3>
        
        <el-form-item label="邮箱" prop="email">
          <el-input
            v-model="formData.email"
            placeholder="请输入联系邮箱"
            size="large"
          />
        </el-form-item>
        
        <el-form-item label="微信" prop="wechat">
          <el-input
            v-model="formData.wechat"
            placeholder="请输入微信号（可选）"
            size="large"
          />
        </el-form-item>
        
        <el-form-item label="电话" prop="phone">
          <el-input
            v-model="formData.phone"
            placeholder="请输入联系电话（可选）"
            size="large"
          />
        </el-form-item>
      </div>
      
      <!-- 大师专属信息 -->
      <div v-if="identity?.type === 'master'" class="form-section">
        <h3 class="section-title">大师专属信息</h3>
        
        <el-form-item label="指导价格" prop="price">
          <el-input-number
            v-model="formData.price"
            :min="0"
            :max="10000"
            :step="50"
            size="large"
            class="price-input"
          >
            <template #prefix>¥</template>
          </el-input-number>
          <span class="price-unit">/ 小时</span>
        </el-form-item>
        
        <el-form-item label="服务类型" prop="serviceTypes">
          <el-checkbox-group v-model="formData.serviceTypes" class="service-types-group">
            <el-checkbox value="one-on-one">1对1指导</el-checkbox>
            <el-checkbox value="group-class">小组课程</el-checkbox>
            <el-checkbox value="project-review">项目评审</el-checkbox>
            <el-checkbox value="career-consulting">职业咨询</el-checkbox>
            <el-checkbox value="code-review">代码审查</el-checkbox>
            <el-checkbox value="mentorship">长期指导</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        
        <el-form-item label="工作经验" prop="experience">
          <el-input
            v-model="formData.experience"
            type="textarea"
            :rows="3"
            placeholder="请描述你的工作经验和专业成就..."
            maxlength="300"
            show-word-limit
          />
        </el-form-item>
      </div>
      
      <!-- 学徒专属信息 -->
      <div v-if="identity?.type === 'apprentice'" class="form-section">
        <h3 class="section-title">学习信息</h3>
        
        <el-form-item label="学习背景" prop="background">
          <el-input
            v-model="formData.background"
            type="textarea"
            :rows="3"
            placeholder="请描述你的学习背景和相关经验..."
            maxlength="300"
            show-word-limit
          />
        </el-form-item>
        
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
        
        <el-form-item label="学习偏好" prop="learningPreferences">
          <el-checkbox-group v-model="formData.learningPreferences" class="preferences-group">
            <el-checkbox value="one-on-one">1对1指导</el-checkbox>
            <el-checkbox value="group-class">小组课程</el-checkbox>
            <el-checkbox value="self-study">自主学习</el-checkbox>
            <el-checkbox value="project-based">项目实践</el-checkbox>
            <el-checkbox value="mentorship">导师制</el-checkbox>
          </el-checkbox-group>
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
          {{ loading ? '保存中...' : '保存修改' }}
        </el-button>
      </div>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Camera } from '@element-plus/icons-vue'
import type { FormInstance, FormRules, UploadProps } from 'element-plus'
import { useAuthStore } from '@/stores/auth'
import type { Identity } from '@/types/user'

// 定义props和事件
interface Props {
  identity: Identity
}

const props = defineProps<Props>()
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

// 计算属性
const identity = computed(() => props.identity)

// 专业领域选项
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

// 技能选项
const skillOptions = [
  'JavaScript', 'Python', 'Java', 'React', 'Vue', 'Node.js', 'TypeScript',
  'UI设计', 'UX设计', 'Photoshop', 'Figma', 'Sketch', 'Illustrator',
  '数据分析', '机器学习', '深度学习', 'SQL', 'Excel', 'Tableau',
  '产品管理', '项目管理', '敏捷开发', 'Scrum', '用户研究',
  '市场营销', 'SEO', 'SEM', '社交媒体', '内容营销', '品牌策划',
  '投资理财', '股票分析', '基金投资', '风险管理', '财务规划',
  '英语', '日语', '韩语', '法语', '德语', '西班牙语',
  '钢琴', '吉他', '小提琴', '声乐', '作曲', '编曲',
  '健身训练', '瑜伽', '跑步', '游泳', '力量训练', '营养搭配',
  '中餐', '西餐', '烘焙', '甜点', '咖啡', '茶艺'
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
  avatar: '',
  bio: '',
  skills: [] as string[],
  email: '',
  wechat: '',
  phone: '',
  price: 0,
  serviceTypes: [] as string[],
  experience: '',
  background: '',
  learningGoals: [] as string[],
  learningPreferences: [] as string[]
})

// 表单验证规则
const formRules: FormRules = {
  name: [
    { required: true, message: '请输入身份名称', trigger: 'blur' },
    { min: 2, max: 20, message: '名称长度在2-20个字符', trigger: 'blur' }
  ],
  domain: [
    { required: true, message: '请选择专业领域', trigger: 'change' }
  ],
  bio: [
    { required: true, message: '请填写个人简介', trigger: 'blur' },
    { min: 30, max: 500, message: '个人简介长度在30-500个字符', trigger: 'blur' }
  ],
  skills: [
    { required: true, message: '请选择技能标签', trigger: 'change' },
    { type: 'array', min: 1, max: 10, message: '请选择1-10个技能标签', trigger: 'change' }
  ],
  email: [
    { required: true, message: '请输入联系邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  price: [
    { required: true, message: '请设置指导价格', trigger: 'blur' },
    { type: 'number', min: 0, message: '价格不能为负数', trigger: 'blur' }
  ],
  serviceTypes: [
    { required: true, message: '请选择服务类型', trigger: 'change' },
    { type: 'array', min: 1, max: 6, message: '请选择1-6种服务类型', trigger: 'change' }
  ],
  background: [
    { required: true, message: '请填写学习背景', trigger: 'blur' },
    { min: 30, max: 300, message: '学习背景长度在30-300个字符', trigger: 'blur' }
  ],
  learningGoals: [
    { required: true, message: '请选择学习目标', trigger: 'change' },
    { type: 'array', min: 1, max: 8, message: '请选择1-8个学习目标', trigger: 'change' }
  ],
  learningPreferences: [
    { required: true, message: '请选择学习偏好', trigger: 'change' },
    { type: 'array', min: 1, max: 5, message: '请选择1-5种学习偏好', trigger: 'change' }
  ]
}

// 获取身份图标
const getIdentityIcon = (type?: string) => {
  if (type === 'master') {
    return 'Star'
  } else if (type === 'apprentice') {
    return 'User'
  }
  return 'User'
}

// 头像上传前验证
const beforeAvatarUpload: UploadProps['beforeUpload'] = (file) => {
  const isJPG = file.type === 'image/jpeg'
  const isPNG = file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isJPG && !isPNG) {
    ElMessage.error('头像只能是 JPG 或 PNG 格式!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('头像大小不能超过 2MB!')
    return false
  }
  return true
}

// 头像上传成功
const handleAvatarSuccess: UploadProps['onSuccess'] = (response) => {
  formData.avatar = response.url
  ElMessage.success('头像上传成功')
}

// 头像上传失败
const handleAvatarError: UploadProps['onError'] = () => {
  ElMessage.error('头像上传失败')
}

// 初始化表单数据
const initializeFormData = () => {
  if (identity.value) {
    formData.name = identity.value.name || ''
    formData.domain = identity.value.domain || ''
    formData.avatar = identity.value.avatar || ''
    formData.bio = identity.value.bio || ''
    formData.skills = identity.value.skills || []
    formData.email = identity.value.email || ''
    formData.wechat = identity.value.wechat || ''
    formData.phone = identity.value.phone || ''
    formData.price = identity.value.price || 0
    formData.serviceTypes = identity.value.serviceTypes || []
    formData.experience = identity.value.experience || ''
    formData.background = identity.value.background || ''
    formData.learningGoals = identity.value.learningGoals || []
    formData.learningPreferences = identity.value.learningPreferences || []
  }
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value?.validate()
    
    loading.value = true
    
    // 调用认证store的更新身份信息方法
    const result = await authStore.updateIdentityInfo(identity.value.id, formData)
    
    // 更新成功
    ElMessage.success('身份信息更新成功！')
    
    // 触发成功事件
    emit('submit-success', {
      ...formData,
      identity: result
    })
    
  } catch (error) {
    console.error('更新失败:', error)
    ElMessage.error('更新失败，请检查表单信息')
  } finally {
    loading.value = false
  }
}

// 组件挂载时初始化数据
onMounted(() => {
  initializeFormData()
})
</script>

<style scoped lang="scss">
.edit-identity-info {
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
    background: var(--primary-color);
    border-radius: 2px;
  }
}

.domain-select,
.skills-select,
.goals-select {
  width: 100%;
}

.avatar-upload {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-md);
}

.avatar-uploader {
  text-align: center;
}

.avatar-preview {
  position: relative;
  display: inline-block;
  cursor: pointer;
  border-radius: 50%;
  overflow: hidden;
  transition: all var(--transition-normal);
  
  &:hover .upload-overlay {
    opacity: 1;
  }
}

.upload-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: white;
  opacity: 0;
  transition: opacity var(--transition-normal);
  
  .el-icon {
    font-size: var(--icon-size-lg);
    margin-bottom: var(--spacing-xs);
  }
  
  span {
    font-size: var(--font-size-small);
  }
}

.avatar-tip {
  text-align: center;
  
  p {
    font-size: var(--font-size-small);
    color: var(--text-tertiary);
    margin: 0;
  }
}

.price-input {
  width: 200px;
}

.price-unit {
  margin-left: var(--spacing-sm);
  color: var(--text-secondary);
}

.service-types-group,
.preferences-group {
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
  .edit-identity-info {
    margin: var(--spacing-md);
    padding: var(--spacing-lg);
  }
  
  .service-types-group,
  .preferences-group {
    flex-direction: column;
    gap: var(--spacing-sm);
  }
  
  .form-actions {
    flex-direction: column;
    gap: var(--spacing-md);
  }
}
</style>