<template>
  <div class="enroll-success-page">
    <div class="success-container">
      <!-- 成功图标 -->
      <div class="success-icon">
        <el-icon size="80" color="#67c23a">
          <CircleCheckFilled />
        </el-icon>
      </div>
      
      <!-- 成功标题 -->
      <h1 class="success-title">报名成功！</h1>
      <p class="success-subtitle">恭喜您成功报名课程，现在可以开始学习了</p>
      
      <!-- 报名信息卡片 -->
      <div class="enrollment-info" v-if="enrollmentInfo">
        <div class="info-card">
          <div class="course-info">
            <img :src="enrollmentInfo.courseCover" :alt="enrollmentInfo.courseTitle" class="course-cover" />
            <div class="course-details">
              <h3>{{ enrollmentInfo.courseTitle }}</h3>
              <p>{{ enrollmentInfo.courseDescription }}</p>
              <div class="course-meta">
                <span>大师：{{ enrollmentInfo.mentorName }}</span>
                <span>时长：{{ enrollmentInfo.duration }}小时</span>
              </div>
            </div>
          </div>
          
          <div class="enrollment-details">
            <div class="detail-item">
              <span class="label">报名时间：</span>
              <span class="value">{{ formatDate(enrollmentInfo.enrolledAt) }}</span>
            </div>
            <div class="detail-item">
              <span class="label">订单号：</span>
              <span class="value">{{ enrollmentInfo.orderId }}</span>
            </div>
            <div class="detail-item">
              <span class="label">支付金额：</span>
              <span class="value price">¥{{ enrollmentInfo.price }}</span>
            </div>
            <div class="detail-item">
              <span class="label">支付方式：</span>
              <span class="value">{{ getPaymentMethodText(enrollmentInfo.paymentMethod) }}</span>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 学习指南 -->
      <div class="learning-guide">
        <h3>开始学习</h3>
        <div class="guide-steps">
          <div class="step">
            <div class="step-number">1</div>
            <div class="step-content">
              <h4>进入学习页面</h4>
              <p>点击下方按钮进入课程学习页面</p>
            </div>
          </div>
          <div class="step">
            <div class="step-number">2</div>
            <div class="step-content">
              <h4>观看课程视频</h4>
              <p>按照课程大纲顺序学习，记录学习笔记</p>
            </div>
          </div>
          <div class="step">
            <div class="step-number">3</div>
            <div class="step-content">
              <h4>完成实践项目</h4>
              <p>动手完成课程中的实践项目，巩固所学知识</p>
            </div>
          </div>
          <div class="step">
            <div class="step-number">4</div>
            <div class="step-content">
              <h4>获得结业证书</h4>
              <p>完成所有课程内容后，获得学习证书</p>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 学习权益 -->
      <div class="learning-benefits">
        <h3>学习权益</h3>
        <div class="benefits-grid">
          <div class="benefit-item">
            <el-icon size="24" color="#409eff">
              <VideoPlay />
            </el-icon>
            <div class="benefit-content">
              <h4>高清视频</h4>
              <p>1080P高清视频，支持倍速播放</p>
            </div>
          </div>
          <div class="benefit-item">
            <el-icon size="24" color="#67c23a">
              <ChatDotRound />
            </el-icon>
            <div class="benefit-content">
              <h4>答疑服务</h4>
              <p>大师一对一答疑，解决学习疑问</p>
            </div>
          </div>
          <div class="benefit-item">
            <el-icon size="24" color="#e6a23c">
              <Document />
            </el-icon>
            <div class="benefit-content">
              <h4>学习资料</h4>
              <p>配套学习资料，代码示例下载</p>
            </div>
          </div>
          <div class="benefit-item">
            <el-icon size="24" color="#f56c6c">
              <Trophy />
            </el-icon>
            <div class="benefit-content">
              <h4>结业证书</h4>
              <p>完成学习后获得权威结业证书</p>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 操作按钮 -->
      <div class="action-buttons">
        <el-button 
          type="primary" 
          size="large" 
          @click="startLearning"
          class="start-learning-btn"
        >
          <el-icon><VideoPlay /></el-icon>
          开始学习
        </el-button>
        
        <el-button 
          type="default" 
          size="large" 
          @click="viewMyCourses"
        >
          <el-icon><Collection /></el-icon>
          我的课程
        </el-button>
        
        <el-button 
          type="info" 
          size="large" 
          @click="goHome"
        >
          <el-icon><House /></el-icon>
          返回首页
        </el-button>
      </div>
      
      <!-- 温馨提示 -->
      <div class="tips">
        <el-alert
          title="温馨提示"
          type="info"
          :closable="false"
          show-icon
        >
          <template #default>
            <p>• 课程内容永久有效，可随时学习</p>
            <p>• 学习过程中遇到问题，可在课程讨论区提问</p>
            <p>• 建议按照课程大纲顺序学习，效果更佳</p>
            <p>• 完成课程后可申请结业证书</p>
          </template>
        </el-alert>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  CircleCheckFilled, 
  VideoPlay, 
  ChatDotRound, 
  Document, 
  Trophy, 
  Collection, 
  House 
} from '@element-plus/icons-vue'
import { ApiService } from '@/services/api'

const route = useRoute()
const router = useRouter()

const enrollmentInfo = ref<any>(null)

// 加载报名信息
const loadEnrollmentInfo = async () => {
  const enrollmentId = route.params.id as string
  
  try {
    // 模拟获取报名信息
    enrollmentInfo.value = {
      id: enrollmentId,
      courseId: '1',
      courseTitle: 'Vue.js 3.0 从入门到精通',
      courseDescription: '系统学习Vue.js 3.0的核心概念和实战应用，掌握现代前端开发技能',
      courseCover: 'https://via.placeholder.com/300x200/4CAF50/FFFFFF?text=Vue.js',
      mentorName: '张大师',
      duration: 20,
      price: 299,
      paymentMethod: 'alipay',
      orderId: `ORDER${enrollmentId}`,
      enrolledAt: new Date().toISOString()
    }
  } catch (error) {
    console.error('加载报名信息失败:', error)
    ElMessage.error('加载报名信息失败')
  }
}

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('zh-CN')
}

// 获取支付方式文本
const getPaymentMethodText = (method: string) => {
  const methods: Record<string, string> = {
    alipay: '支付宝',
    wechat: '微信支付',
    card: '银行卡'
  }
  return methods[method] || method
}

// 开始学习
const startLearning = () => {
  router.push(`/learning/${enrollmentInfo.value.courseId}`)
}

// 查看我的课程
const viewMyCourses = () => {
  router.push('/my-courses')
}

// 返回首页
const goHome = () => {
  router.push('/')
}

onMounted(() => {
  loadEnrollmentInfo()
})
</script>

<style scoped lang="scss">
.enroll-success-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--spacing-xl);
}

.success-container {
  max-width: 800px;
  width: 100%;
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  padding: var(--spacing-xxl);
  box-shadow: var(--shadow-heavy);
  text-align: center;
}

.success-icon {
  margin-bottom: var(--spacing-lg);
}

.success-title {
  font-size: var(--font-size-h1);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-sm) 0;
}

.success-subtitle {
  font-size: var(--font-size-h5);
  color: var(--text-secondary);
  margin: 0 0 var(--spacing-xl) 0;
}

.enrollment-info {
  margin-bottom: var(--spacing-xxl);
  
  .info-card {
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
    padding: var(--spacing-lg);
    text-align: left;
  }
  
  .course-info {
    display: flex;
    gap: var(--spacing-md);
    margin-bottom: var(--spacing-lg);
    padding-bottom: var(--spacing-lg);
    border-bottom: 1px solid var(--border-color-light);
    
    .course-cover {
      width: 80px;
      height: 60px;
      object-fit: cover;
      border-radius: var(--border-radius-small);
    }
    
    .course-details {
      flex: 1;
      
      h3 {
        font-size: var(--font-size-h5);
        font-weight: var(--font-weight-semibold);
        color: var(--text-primary);
        margin: 0 0 var(--spacing-xs) 0;
      }
      
      p {
        font-size: var(--font-size-small);
        color: var(--text-secondary);
        margin: 0 0 var(--spacing-sm) 0;
        line-height: 1.4;
      }
      
      .course-meta {
        display: flex;
        gap: var(--spacing-md);
        font-size: var(--font-size-small);
        color: var(--text-tertiary);
      }
    }
  }
  
  .enrollment-details {
    .detail-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: var(--spacing-sm) 0;
      font-size: var(--font-size-medium);
      
      .label {
        color: var(--text-secondary);
      }
      
      .value {
        color: var(--text-primary);
        font-weight: var(--font-weight-medium);
        
        &.price {
          color: #f56c6c;
          font-weight: var(--font-weight-bold);
        }
      }
    }
  }
}

.learning-guide {
  margin-bottom: var(--spacing-xxl);
  text-align: left;
  
  h3 {
    font-size: var(--font-size-h4);
    font-weight: var(--font-weight-bold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-lg) 0;
    text-align: center;
  }
  
  .guide-steps {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: var(--spacing-lg);
    
    .step {
      display: flex;
      align-items: flex-start;
      gap: var(--spacing-md);
      
      .step-number {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 40px;
        height: 40px;
        background: var(--primary-color);
        color: white;
        border-radius: 50%;
        font-weight: var(--font-weight-bold);
        font-size: var(--font-size-h5);
        flex-shrink: 0;
      }
      
      .step-content {
        h4 {
          font-size: var(--font-size-h5);
          font-weight: var(--font-weight-semibold);
          color: var(--text-primary);
          margin: 0 0 var(--spacing-xs) 0;
        }
        
        p {
          font-size: var(--font-size-small);
          color: var(--text-secondary);
          margin: 0;
          line-height: 1.4;
        }
      }
    }
  }
}

.learning-benefits {
  margin-bottom: var(--spacing-xxl);
  text-align: left;
  
  h3 {
    font-size: var(--font-size-h4);
    font-weight: var(--font-weight-bold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-lg) 0;
    text-align: center;
  }
  
  .benefits-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: var(--spacing-lg);
    
    .benefit-item {
      display: flex;
      align-items: flex-start;
      gap: var(--spacing-md);
      padding: var(--spacing-md);
      background: var(--bg-secondary);
      border-radius: var(--border-radius-medium);
      
      .benefit-content {
        h4 {
          font-size: var(--font-size-h5);
          font-weight: var(--font-weight-semibold);
          color: var(--text-primary);
          margin: 0 0 var(--spacing-xs) 0;
        }
        
        p {
          font-size: var(--font-size-small);
          color: var(--text-secondary);
          margin: 0;
          line-height: 1.4;
        }
      }
    }
  }
}

.action-buttons {
  display: flex;
  justify-content: center;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-xl);
  flex-wrap: wrap;
  
  .start-learning-btn {
    background: linear-gradient(135deg, var(--primary-color), var(--master-color));
    border: none;
    color: white;
    font-weight: var(--font-weight-medium);
  }
}

.tips {
  .el-alert {
    text-align: left;
    
    p {
      margin: var(--spacing-xs) 0;
      font-size: var(--font-size-small);
    }
  }
}

// 响应式设计
@media (max-width: 768px) {
  .enroll-success-page {
    padding: var(--spacing-lg);
  }
  
  .success-container {
    padding: var(--spacing-lg);
  }
  
  .course-info {
    flex-direction: column;
    text-align: center;
  }
  
  .guide-steps {
    grid-template-columns: 1fr;
  }
  
  .benefits-grid {
    grid-template-columns: 1fr;
  }
  
  .action-buttons {
    flex-direction: column;
    align-items: center;
    
    .el-button {
      width: 100%;
      max-width: 300px;
    }
  }
}
</style> 