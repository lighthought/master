<template>
  <div class="course-enroll-page">
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="10" animated />
    </div>
    
    <div v-else-if="course" class="enroll-content">
      <!-- 页面头部 -->
      <div class="page-header">
        <el-button @click="$router.go(-1)" type="text" class="back-button">
          <el-icon><ArrowLeft /></el-icon>
          返回
        </el-button>
        <h1 class="page-title">课程报名</h1>
      </div>
      
      <div class="enroll-layout">
        <!-- 左侧：课程信息 -->
        <div class="course-info-section">
          <div class="course-card">
            <img :src="course.cover" :alt="course.title" class="course-cover" />
            <div class="course-details">
              <h2 class="course-title">{{ course.title }}</h2>
              <p class="course-description">{{ course.description }}</p>
              
              <div class="course-meta">
                <div class="meta-item">
                  <el-icon><User /></el-icon>
                  <span>{{ course.mentorName }}</span>
                </div>
                <div class="meta-item">
                  <el-icon><Clock /></el-icon>
                  <span>{{ course.duration }}小时</span>
                </div>
                <div class="meta-item">
                  <el-icon><UserFilled /></el-icon>
                  <span>{{ course.studentCount }}名学员</span>
                </div>
                <div class="meta-item">
                  <el-icon><Star /></el-icon>
                  <span>{{ course.rating.toFixed(1) }}分</span>
                </div>
              </div>
              
              <div class="course-outline">
                <h4>课程大纲</h4>
                <div class="outline-summary">
                  <span>{{ course.outline.length }}个章节</span>
                  <span>{{ getTotalLessons() }}个课时</span>
                  <span>{{ getTotalDuration() }}小时</span>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- 右侧：报名表单 -->
        <div class="enroll-form-section">
          <div class="enroll-card">
            <h3 class="form-title">报名信息</h3>
            
            <!-- 价格信息 -->
            <div class="price-section">
              <div class="price-info">
                <span class="current-price">¥{{ course.price }}</span>
                <span class="original-price" v-if="course.originalPrice">¥{{ course.originalPrice }}</span>
                <span class="discount" v-if="course.originalPrice">
                  省¥{{ course.originalPrice - course.price }}
                </span>
              </div>
              <div class="price-details">
                <div class="price-item">
                  <span>课程费用</span>
                  <span>¥{{ course.price }}</span>
                </div>
                <div class="price-item">
                  <span>优惠折扣</span>
                  <span class="discount-text">-¥{{ course.originalPrice ? course.originalPrice - course.price : 0 }}</span>
                </div>
                <div class="price-item total">
                  <span>实付金额</span>
                  <span class="total-price">¥{{ course.price }}</span>
                </div>
              </div>
            </div>
            
            <!-- 报名表单 -->
            <el-form 
              ref="enrollFormRef"
              :model="enrollForm" 
              :rules="enrollRules"
              label-width="80px"
              class="enroll-form"
            >
                             <el-form-item label="姓名" prop="name">
                 <el-input 
                   v-model="enrollForm.name" 
                   placeholder="请输入真实姓名"
                   :disabled="!!authStore.user?.identities?.[0]?.name"
                 />
               </el-form-item>
               
               <el-form-item label="手机号" prop="phone">
                 <el-input 
                   v-model="enrollForm.phone" 
                   placeholder="请输入手机号码"
                 />
               </el-form-item>
               
               <el-form-item label="邮箱" prop="email">
                 <el-input 
                   v-model="enrollForm.email" 
                   placeholder="请输入邮箱地址"
                   :disabled="!!authStore.user?.email"
                 />
               </el-form-item>
              
              <el-form-item label="学习目标" prop="goal">
                <el-input 
                  v-model="enrollForm.goal" 
                  type="textarea"
                  :rows="3"
                  placeholder="请简要描述你的学习目标..."
                />
              </el-form-item>
            </el-form>
            
            <!-- 支付方式 -->
            <div class="payment-section">
              <h4>选择支付方式</h4>
              <el-radio-group v-model="paymentMethod" class="payment-methods">
                <el-radio label="alipay" class="payment-method">
                  <div class="payment-option">
                    <img src="https://via.placeholder.com/24x24/1677FF/FFFFFF?text=支" alt="支付宝" />
                    <span>支付宝</span>
                  </div>
                </el-radio>
                <el-radio label="wechat" class="payment-method">
                  <div class="payment-option">
                    <img src="https://via.placeholder.com/24x24/07C160/FFFFFF?text=微" alt="微信支付" />
                    <span>微信支付</span>
                  </div>
                </el-radio>
                <el-radio label="card" class="payment-method">
                  <div class="payment-option">
                    <img src="https://via.placeholder.com/24x24/FF6B35/FFFFFF?text=卡" alt="银行卡" />
                    <span>银行卡</span>
                  </div>
                </el-radio>
              </el-radio-group>
            </div>
            
            <!-- 协议确认 -->
            <div class="agreement-section">
              <el-checkbox v-model="agreedToTerms">
                我已阅读并同意
                <el-link type="primary" @click="showTerms = true">《课程报名协议》</el-link>
                和
                <el-link type="primary" @click="showPrivacy = true">《隐私政策》</el-link>
              </el-checkbox>
            </div>
            
            <!-- 报名按钮 -->
            <div class="enroll-actions">
              <el-button 
                type="primary" 
                size="large" 
                @click="submitEnrollment"
                :loading="submitting"
                :disabled="!agreedToTerms"
                class="enroll-button"
              >
                <el-icon><ShoppingCart /></el-icon>
                立即报名 ¥{{ course.price }}
              </el-button>
              
              <p class="enroll-tips">
                <el-icon><InfoFilled /></el-icon>
                报名成功后，您将立即获得课程访问权限
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 空状态 -->
    <div v-else class="empty-state">
      <el-empty description="课程不存在或已被删除">
        <el-button type="primary" @click="$router.push('/courses')">
          返回课程列表
        </el-button>
      </el-empty>
    </div>
    
    <!-- 协议对话框 -->
    <el-dialog v-model="showTerms" title="课程报名协议" width="80%" max-width="600px">
      <div class="terms-content">
        <h4>课程报名协议</h4>
        <p>1. 课程内容：本课程包含视频教学、实践项目、答疑服务等内容。</p>
        <p>2. 学习期限：报名后永久有效，可随时学习。</p>
        <p>3. 退款政策：报名后7天内可申请退款，超过7天不予退款。</p>
        <p>4. 知识产权：课程内容受知识产权保护，禁止传播和分享。</p>
        <p>5. 学习责任：学员应按时完成学习任务，积极参与讨论。</p>
      </div>
    </el-dialog>
    
    <el-dialog v-model="showPrivacy" title="隐私政策" width="80%" max-width="600px">
      <div class="privacy-content">
        <h4>隐私政策</h4>
        <p>1. 信息收集：我们收集您的基本信息用于课程服务。</p>
        <p>2. 信息使用：您的信息仅用于课程相关服务。</p>
        <p>3. 信息保护：我们采用安全措施保护您的个人信息。</p>
        <p>4. 信息共享：未经您同意，我们不会与第三方分享您的信息。</p>
        <p>5. 信息删除：您有权要求删除您的个人信息。</p>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  ArrowLeft, 
  User, 
  Clock, 
  UserFilled, 
  Star, 
  ShoppingCart, 
  InfoFilled 
} from '@element-plus/icons-vue'
import { ApiService } from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const course = ref<any>(null)
const loading = ref(true)
const submitting = ref(false)
const showTerms = ref(false)
const showPrivacy = ref(false)
const agreedToTerms = ref(false)
const paymentMethod = ref('alipay')

const enrollFormRef = ref()
const enrollForm = ref({
  name: authStore.user?.identities?.[0]?.name || '',
  phone: '',
  email: authStore.user?.email || '',
  goal: ''
})

const enrollRules = {
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' },
    { min: 2, max: 20, message: '姓名长度在2到20个字符', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email' as const, message: '请输入正确的邮箱地址', trigger: 'blur' }
  ],
  goal: [
    { required: true, message: '请输入学习目标', trigger: 'blur' },
    { min: 10, max: 200, message: '学习目标长度在10到200个字符', trigger: 'blur' }
  ]
}

// 加载课程详情
const loadCourseDetail = async () => {
  const courseId = route.params.id as string
  loading.value = true
  
  try {
    const response = await ApiService.courses.getCourseDetail(courseId)
    course.value = response.data
    
    // 检查是否已报名
    if (course.value.enrollmentStatus === 'enrolled') {
      ElMessage.warning('您已经报名了这门课程')
      router.push(`/courses/${courseId}`)
      return
    }
  } catch (error) {
    console.error('加载课程详情失败:', error)
    ElMessage.error('加载课程详情失败')
  } finally {
    loading.value = false
  }
}

// 计算总课时数
const getTotalLessons = () => {
  if (!course.value?.outline) return 0
  return course.value.outline.reduce((total: number, section: any) => {
    return total + section.lessons.length
  }, 0)
}

// 计算总时长
const getTotalDuration = () => {
  if (!course.value?.outline) return 0
  let totalMinutes = 0
  
  course.value.outline.forEach((section: any) => {
    section.lessons.forEach((lesson: any) => {
      const duration = lesson.duration
      const minutes = parseInt(duration.match(/(\d+)/)?.[1] || '0')
      totalMinutes += minutes
    })
  })
  
  return Math.round(totalMinutes / 60)
}

// 提交报名
const submitEnrollment = async () => {
  if (!agreedToTerms.value) {
    ElMessage.warning('请先同意课程报名协议和隐私政策')
    return
  }
  
  try {
    await enrollFormRef.value.validate()
  } catch (error) {
    return
  }
  
  submitting.value = true
  
  try {
    // 确认报名信息
    await ElMessageBox.confirm(
      `确认报名课程"${course.value.title}"？\n支付金额：¥${course.value.price}`,
      '确认报名',
      {
        confirmButtonText: '确认支付',
        cancelButtonText: '取消',
        type: 'info'
      }
    )
    
    // 模拟支付过程
    ElMessage.info('正在跳转到支付页面...')
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // 创建报名记录
    const enrollData = {
      courseId: course.value.id,
      userId: authStore.user?.id,
      price: course.value.price,
      paymentMethod: paymentMethod.value,
      userInfo: enrollForm.value
    }
    
    const result = await ApiService.courses.enrollCourse(enrollData)
    
    ElMessage.success('报名成功！')
    
    // 跳转到报名成功页面
    router.push(`/enroll-success/${result.data.id}`)
    
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('报名失败，请重试')
    }
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  loadCourseDetail()
})
</script>

<style scoped lang="scss">
.course-enroll-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: var(--spacing-xl);
}

.loading-container {
  padding: var(--spacing-xl);
}

.page-header {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-xl);
  
  .back-button {
    font-size: var(--font-size-medium);
  }
  
  .page-title {
    font-size: var(--font-size-h2);
    font-weight: var(--font-weight-bold);
    color: var(--text-primary);
    margin: 0;
  }
}

.enroll-layout {
  display: grid;
  grid-template-columns: 1fr 400px;
  gap: var(--spacing-xl);
  
  @media (max-width: 1024px) {
    grid-template-columns: 1fr;
    gap: var(--spacing-lg);
  }
}

.course-info-section {
  .course-card {
    background: var(--bg-card);
    border-radius: var(--border-radius-medium);
    overflow: hidden;
    box-shadow: var(--shadow-light);
    
    .course-cover {
      width: 100%;
      height: 200px;
      object-fit: cover;
    }
    
    .course-details {
      padding: var(--spacing-lg);
    }
    
    .course-title {
      font-size: var(--font-size-h3);
      font-weight: var(--font-weight-bold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-md) 0;
      line-height: 1.3;
    }
    
    .course-description {
      font-size: var(--font-size-medium);
      color: var(--text-secondary);
      line-height: 1.6;
      margin: 0 0 var(--spacing-lg) 0;
    }
    
    .course-meta {
      display: grid;
      grid-template-columns: repeat(2, 1fr);
      gap: var(--spacing-md);
      margin-bottom: var(--spacing-lg);
      
      .meta-item {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        font-size: var(--font-size-medium);
        color: var(--text-secondary);
        
        .el-icon {
          color: var(--primary-color);
        }
      }
    }
    
    .course-outline {
      h4 {
        font-size: var(--font-size-h5);
        font-weight: var(--font-weight-semibold);
        color: var(--text-primary);
        margin: 0 0 var(--spacing-md) 0;
      }
      
      .outline-summary {
        display: flex;
        gap: var(--spacing-lg);
        font-size: var(--font-size-medium);
        color: var(--text-secondary);
        
        span {
          display: flex;
          align-items: center;
          gap: var(--spacing-xs);
          
          &:before {
            content: '•';
            color: var(--primary-color);
            font-weight: var(--font-weight-bold);
          }
        }
      }
    }
  }
}

.enroll-form-section {
  .enroll-card {
    background: var(--bg-card);
    border-radius: var(--border-radius-medium);
    padding: var(--spacing-xl);
    box-shadow: var(--shadow-light);
    position: sticky;
    top: var(--spacing-xl);
    
    .form-title {
      font-size: var(--font-size-h4);
      font-weight: var(--font-weight-bold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-lg) 0;
    }
  }
}

.price-section {
  margin-bottom: var(--spacing-xl);
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  
  .price-info {
    display: flex;
    align-items: baseline;
    gap: var(--spacing-sm);
    margin-bottom: var(--spacing-md);
    
    .current-price {
      font-size: var(--font-size-h2);
      font-weight: var(--font-weight-bold);
      color: #f56c6c;
    }
    
    .original-price {
      font-size: var(--font-size-h4);
      color: var(--text-tertiary);
      text-decoration: line-through;
    }
    
    .discount {
      font-size: var(--font-size-small);
      color: #67c23a;
      background: rgba(103, 194, 58, 0.1);
      padding: 4px 8px;
      border-radius: var(--border-radius-small);
    }
  }
  
  .price-details {
    .price-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: var(--spacing-sm) 0;
      font-size: var(--font-size-medium);
      
      &.total {
        border-top: 1px solid var(--border-color-light);
        margin-top: var(--spacing-sm);
        padding-top: var(--spacing-md);
        font-weight: var(--font-weight-semibold);
        
        .total-price {
          font-size: var(--font-size-h4);
          color: #f56c6c;
        }
      }
      
      .discount-text {
        color: #67c23a;
      }
    }
  }
}

.enroll-form {
  margin-bottom: var(--spacing-xl);
}

.payment-section {
  margin-bottom: var(--spacing-lg);
  
  h4 {
    font-size: var(--font-size-h5);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-md) 0;
  }
  
  .payment-methods {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-md);
    
    .payment-method {
      width: 100%;
      margin: 0;
      
      .payment-option {
        display: flex;
        align-items: center;
        gap: var(--spacing-md);
        padding: var(--spacing-md);
        border: 1px solid var(--border-color-light);
        border-radius: var(--border-radius-small);
        transition: all var(--transition-normal);
        
        &:hover {
          border-color: var(--primary-color);
        }
        
        img {
          width: 24px;
          height: 24px;
          border-radius: var(--border-radius-small);
        }
        
        span {
          font-size: var(--font-size-medium);
          color: var(--text-primary);
        }
      }
    }
  }
}

.agreement-section {
  margin-bottom: var(--spacing-lg);
  padding: var(--spacing-md);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-small);
  font-size: var(--font-size-small);
  color: var(--text-secondary);
}

.enroll-actions {
  .enroll-button {
    width: 100%;
    background: linear-gradient(135deg, var(--primary-color), var(--master-color));
    border: none;
    color: white;
    font-weight: var(--font-weight-medium);
    font-size: var(--font-size-medium);
    height: 48px;
  }
  
  .enroll-tips {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
    margin-top: var(--spacing-md);
    font-size: var(--font-size-small);
    color: var(--text-secondary);
    text-align: center;
    
    .el-icon {
      color: var(--primary-color);
    }
  }
}

.terms-content,
.privacy-content {
  h4 {
    font-size: var(--font-size-h5);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-md) 0;
  }
  
  p {
    font-size: var(--font-size-medium);
    color: var(--text-secondary);
    line-height: 1.6;
    margin: 0 0 var(--spacing-sm) 0;
  }
}

.empty-state {
  text-align: center;
  padding: var(--spacing-xxl) 0;
}

// 响应式设计
@media (max-width: 768px) {
  .course-enroll-page {
    padding: var(--spacing-lg);
  }
  
  .enroll-layout {
    grid-template-columns: 1fr;
  }
  
  .course-meta {
    grid-template-columns: 1fr;
  }
  
  .outline-summary {
    flex-direction: column;
    gap: var(--spacing-sm);
  }
}
</style> 