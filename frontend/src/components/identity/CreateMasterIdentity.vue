<template>
  <div class="create-master-identity">
    <div class="form-header">
      <h2 class="form-title">创建大师身份</h2>
      <p class="form-subtitle">完善你的大师信息，开始提供专业指导服务</p>
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
        
        <el-form-item label="大师姓名" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入你的大师姓名"
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
      
      <!-- 个人介绍 -->
      <div class="form-section">
        <h3 class="section-title">个人介绍</h3>
        
        <el-form-item label="个人简介" prop="bio">
          <el-input
            v-model="formData.bio"
            type="textarea"
            :rows="4"
            placeholder="请介绍你的专业背景、经验和专长..."
            maxlength="500"
            show-word-limit
          />
        </el-form-item>
        
        <el-form-item label="工作经历" prop="experience">
          <el-input
            v-model="formData.experience"
            type="textarea"
            :rows="3"
            placeholder="请描述你的工作经历和项目经验..."
            maxlength="300"
            show-word-limit
          />
        </el-form-item>
      </div>
      
      <!-- 服务设置 -->
      <div class="form-section">
        <h3 class="section-title">服务设置</h3>
        
        <el-form-item label="指导价格" prop="price">
          <el-input-number
            v-model="formData.price"
            :min="50"
            :max="2000"
            :step="10"
            size="large"
            class="price-input"
          >
            <template #prefix>¥</template>
          </el-input-number>
          <span class="price-unit">/小时</span>
        </el-form-item>
        
        <el-form-item label="服务方式" prop="serviceTypes">
          <el-checkbox-group v-model="formData.serviceTypes" class="service-types">
            <el-checkbox value="video">视频指导</el-checkbox>
            <el-checkbox value="voice">语音指导</el-checkbox>
            <el-checkbox value="text">文字指导</el-checkbox>
            <el-checkbox value="onsite">现场指导</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </div>
      
      <!-- 资质证明 -->
      <div class="form-section">
        <h3 class="section-title">资质证明</h3>
        <p class="section-desc">请上传相关资质证明文件，这将有助于提高你的可信度</p>
        
        <el-form-item label="资质证书" prop="certificates">
          <el-upload
            v-model:file-list="certificateFiles"
            action="#"
            :auto-upload="false"
            :on-change="handleCertificateChange"
            :on-remove="handleCertificateRemove"
            multiple
            accept=".pdf,.jpg,.jpeg,.png"
            class="certificate-upload"
          >
            <el-button type="primary" plain>
              <el-icon><Upload /></el-icon>
              上传证书
            </el-button>
            <template #tip>
              <div class="upload-tip">
                支持 PDF、JPG、PNG 格式，单个文件不超过 10MB
              </div>
            </template>
          </el-upload>
        </el-form-item>
        
        <el-form-item label="作品展示" prop="portfolio">
          <el-upload
            v-model:file-list="portfolioFiles"
            action="#"
            :auto-upload="false"
            :on-change="handlePortfolioChange"
            :on-remove="handlePortfolioRemove"
            multiple
            accept=".pdf,.jpg,.jpeg,.png,.mp4"
            class="portfolio-upload"
          >
            <el-button type="primary" plain>
              <el-icon><Picture /></el-icon>
              上传作品
            </el-button>
            <template #tip>
              <div class="upload-tip">
                支持 PDF、图片、视频格式，展示你的专业能力
              </div>
            </template>
          </el-upload>
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
          {{ loading ? '提交中...' : '提交审核' }}
        </el-button>
      </div>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules, UploadFile } from 'element-plus'

// 定义事件
const emit = defineEmits<{
  'cancel': []
  'submit-success': [identityData: any]
}>()

// 表单引用
const formRef = ref<FormInstance>()

// 加载状态
const loading = ref(false)

// 文件列表
const certificateFiles = ref<UploadFile[]>([])
const portfolioFiles = ref<UploadFile[]>([])

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

// 技能标签选项
const skillOptions = [
  'JavaScript', 'Python', 'Java', 'React', 'Vue.js', 'Node.js',
  'UI设计', 'UX设计', 'Photoshop', 'Figma', 'Sketch',
  'SEO', 'SEM', '社交媒体营销', '内容营销', '品牌营销',
  '数据分析', '机器学习', '深度学习', '统计学',
  '产品设计', '用户研究', '项目管理', '敏捷开发',
  '投资理财', '股票分析', '基金投资', '风险管理',
  '英语', '日语', '韩语', '法语', '德语',
  '钢琴', '吉他', '小提琴', '声乐', '作曲',
  '力量训练', '有氧运动', '瑜伽', '普拉提', '营养学',
  '中餐', '西餐', '烘焙', '调酒', '咖啡制作'
]

// 表单数据
const formData = reactive({
  name: '',
  domain: '',
  skills: [] as string[],
  bio: '',
  experience: '',
  price: 200,
  serviceTypes: ['video', 'voice'] as string[]
})

// 表单验证规则
const formRules: FormRules = {
  name: [
    { required: true, message: '请输入大师姓名', trigger: 'blur' },
    { min: 2, max: 20, message: '姓名长度在2-20个字符', trigger: 'blur' }
  ],
  domain: [
    { required: true, message: '请选择专业领域', trigger: 'change' }
  ],
  skills: [
    { required: true, message: '请选择技能标签', trigger: 'change' },
    { type: 'array', min: 1, max: 10, message: '请选择1-10个技能标签', trigger: 'change' }
  ],
  bio: [
    { required: true, message: '请填写个人简介', trigger: 'blur' },
    { min: 50, max: 500, message: '个人简介长度在50-500个字符', trigger: 'blur' }
  ],
  experience: [
    { required: true, message: '请填写工作经历', trigger: 'blur' },
    { min: 30, max: 300, message: '工作经历长度在30-300个字符', trigger: 'blur' }
  ],
  price: [
    { required: true, message: '请设置指导价格', trigger: 'blur' },
    { type: 'number', min: 50, max: 2000, message: '价格范围在50-2000元/小时', trigger: 'blur' }
  ],
  serviceTypes: [
    { required: true, message: '请选择服务方式', trigger: 'change' },
    { type: 'array', min: 1, message: '至少选择一种服务方式', trigger: 'change' }
  ]
}

// 处理证书文件变化
const handleCertificateChange = (file: UploadFile) => {
  // 验证文件大小
  if (file.size && file.size > 10 * 1024 * 1024) {
    ElMessage.error('文件大小不能超过10MB')
    return false
  }
  return true
}

// 处理证书文件移除
const handleCertificateRemove = (file: UploadFile) => {
  console.log('移除证书文件:', file)
}

// 处理作品文件变化
const handlePortfolioChange = (file: UploadFile) => {
  // 验证文件大小
  if (file.size && file.size > 50 * 1024 * 1024) {
    ElMessage.error('文件大小不能超过50MB')
    return false
  }
  return true
}

// 处理作品文件移除
const handlePortfolioRemove = (file: UploadFile) => {
  console.log('移除作品文件:', file)
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value?.validate()
    
    loading.value = true
    
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // 提交成功
    ElMessage.success('大师身份创建成功！正在审核中，请耐心等待')
    
    // 触发成功事件
    emit('submit-success', {
      ...formData,
      certificates: certificateFiles.value,
      portfolio: portfolioFiles.value,
      status: 'pending'
    })
    
  } catch (error) {
    console.error('提交失败:', error)
    ElMessage.error('提交失败，请检查表单信息')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped lang="scss">
.create-master-identity {
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

.section-desc {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-lg);
  line-height: 1.5;
}

.domain-select,
.skills-select {
  width: 100%;
}

.price-input {
  width: 200px;
}

.price-unit {
  margin-left: var(--spacing-sm);
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
}

.service-types {
  display: flex;
  flex-wrap: wrap;
  gap: var(--spacing-md);
}

.certificate-upload,
.portfolio-upload {
  width: 100%;
}

.upload-tip {
  font-size: var(--font-size-small);
  color: var(--text-tertiary);
  margin-top: var(--spacing-xs);
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
  .create-master-identity {
    margin: var(--spacing-md);
    padding: var(--spacing-lg);
  }
  
  .service-types {
    flex-direction: column;
    gap: var(--spacing-sm);
  }
  
  .form-actions {
    flex-direction: column;
    gap: var(--spacing-md);
  }
  
  .price-input {
    width: 100%;
  }
}
</style>