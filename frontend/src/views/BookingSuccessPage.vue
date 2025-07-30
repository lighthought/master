<template>
  <div class="booking-success-page">
    <div class="success-container">
      <!-- 成功图标 -->
      <div class="success-icon">
        <el-icon class="check-icon"><Check /></el-icon>
      </div>
      
      <!-- 成功标题 -->
      <h1 class="success-title">预约成功！</h1>
      <p class="success-subtitle">你的预约已提交，大师会尽快确认</p>
      
      <!-- 预约信息卡片 -->
      <div v-if="bookingInfo" class="booking-info-card">
        <div class="card-header">
          <h3>预约详情</h3>
        </div>
        
        <div class="booking-details">
          <div class="detail-row">
            <div class="detail-item">
              <span class="label">预约号：</span>
              <span class="value">{{ bookingInfo.id }}</span>
            </div>
            <div class="detail-item">
              <span class="label">大师：</span>
              <span class="value">{{ bookingInfo.mentorName }}</span>
            </div>
          </div>
          
          <div class="detail-row">
            <div class="detail-item">
              <span class="label">指导时间：</span>
              <span class="value">{{ formatDate(bookingInfo.date) }} {{ bookingInfo.timeSlot }}</span>
            </div>
            <div class="detail-item">
              <span class="label">指导方式：</span>
              <span class="value">{{ getMethodText(bookingInfo.method) }}</span>
            </div>
          </div>
          
          <div class="detail-row">
            <div class="detail-item">
              <span class="label">指导时长：</span>
              <span class="value">{{ bookingInfo.duration || 1 }}小时</span>
            </div>
            <div class="detail-item">
              <span class="label">指导价格：</span>
              <span class="value price">¥{{ bookingInfo.price }}</span>
            </div>
          </div>
          
          <div class="detail-row" v-if="bookingInfo.requirements">
            <div class="detail-item full-width">
              <span class="label">指导需求：</span>
              <span class="value">{{ bookingInfo.requirements }}</span>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 下一步操作 -->
      <div class="next-steps">
        <h3>接下来你需要：</h3>
        <div class="steps-list">
          <div class="step-item">
            <div class="step-number">1</div>
            <div class="step-content">
              <h4>等待大师确认</h4>
              <p>大师会在24小时内确认你的预约，请保持关注</p>
            </div>
          </div>
          
          <div class="step-item">
            <div class="step-number">2</div>
            <div class="step-content">
              <h4>准备学习内容</h4>
              <p>提前准备好你的问题和学习目标，让指导更有效果</p>
            </div>
          </div>
          
          <div class="step-item">
            <div class="step-number">3</div>
            <div class="step-content">
              <h4>准时参加指导</h4>
              <p>在预约时间前5分钟进入会议，确保网络和设备正常</p>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 操作按钮 -->
      <div class="action-buttons">
        <el-button 
          type="primary" 
          size="large"
          @click="$router.push('/bookings')"
        >
          <el-icon><Calendar /></el-icon>
          查看我的预约
        </el-button>
        
        <el-button 
          type="default" 
          size="large"
          @click="$router.push('/mentors')"
        >
          <el-icon><User /></el-icon>
          继续寻找大师
        </el-button>
        
        <el-button 
          type="default" 
          size="large"
          @click="$router.push('/')"
        >
          <el-icon><House /></el-icon>
          返回首页
        </el-button>
      </div>
      
      <!-- 温馨提示 -->
      <div class="tips-section">
        <h4>温馨提示</h4>
        <ul class="tips-list">
          <li>预约确认后，系统会发送通知到你的邮箱和手机</li>
          <li>如需修改预约时间，请提前24小时联系大师</li>
          <li>指导结束后，记得给大师写评价，帮助其他学员选择</li>
          <li>如有任何问题，请联系客服：400-123-4567</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Check, Calendar, User, House } from '@element-plus/icons-vue'

// 路由
const route = useRoute()

// 数据
const bookingInfo = ref<any>(null)

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('zh-CN')
}

// 获取指导方式文本
const getMethodText = (method: string) => {
  const texts: Record<string, string> = {
    video: '视频通话',
    voice: '语音通话',
    text: '文字指导'
  }
  return texts[method] || method
}

// 组件挂载时获取预约信息
onMounted(() => {
  // 从路由参数或本地存储获取预约信息
  const bookingId = route.params.bookingId
  if (bookingId) {
    // 实际应用中会从API获取预约详情
    bookingInfo.value = {
      id: bookingId,
      mentorName: '张大师',
      date: '2024-01-20',
      timeSlot: '14:00-15:00',
      method: 'video',
      duration: 1,
      price: 200,
      requirements: '希望学习Vue.js组件化开发的最佳实践'
    }
  }
})
</script>

<style scoped lang="scss">
.booking-success-page {
  min-height: 100vh;
  background: linear-gradient(135deg, var(--bg-primary), var(--bg-secondary));
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
  box-shadow: var(--shadow-card);
  padding: var(--spacing-xxl);
  text-align: center;
}

.success-icon {
  margin-bottom: var(--spacing-xl);
  
  .check-icon {
    font-size: 80px;
    color: #67c23a;
    background: rgba(103, 194, 58, 0.1);
    border-radius: 50%;
    padding: var(--spacing-lg);
  }
}

.success-title {
  font-size: var(--font-size-h1);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-sm) 0;
}

.success-subtitle {
  font-size: var(--font-size-large);
  color: var(--text-secondary);
  margin: 0 0 var(--spacing-xl) 0;
}

.booking-info-card {
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  padding: var(--spacing-xl);
  margin-bottom: var(--spacing-xl);
  text-align: left;
}

.card-header {
  margin-bottom: var(--spacing-lg);
  
  h3 {
    font-size: var(--font-size-h4);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0;
  }
}

.booking-details {
  .detail-row {
    display: flex;
    gap: var(--spacing-xl);
    margin-bottom: var(--spacing-md);
    
    &:last-child {
      margin-bottom: 0;
    }
  }
  
  .detail-item {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
    flex: 1;
    
    &.full-width {
      width: 100%;
    }
  }
  
  .label {
    font-size: var(--font-size-small);
    color: var(--text-secondary);
    min-width: 80px;
  }
  
  .value {
    font-size: var(--font-size-medium);
    color: var(--text-primary);
    font-weight: var(--font-weight-medium);
    
    &.price {
      color: #ff9900;
      font-weight: var(--font-weight-bold);
    }
  }
}

.next-steps {
  margin-bottom: var(--spacing-xl);
  text-align: left;
  
  h3 {
    font-size: var(--font-size-h4);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-lg) 0;
    text-align: center;
  }
}

.steps-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.step-item {
  display: flex;
  align-items: flex-start;
  gap: var(--spacing-md);
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
}

.step-number {
  width: 40px;
  height: 40px;
  background: var(--primary-color);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: var(--font-size-h5);
  font-weight: var(--font-weight-bold);
  flex-shrink: 0;
}

.step-content {
  flex: 1;
  
  h4 {
    font-size: var(--font-size-h5);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-xs) 0;
  }
  
  p {
    font-size: var(--font-size-medium);
    color: var(--text-secondary);
    margin: 0;
    line-height: 1.5;
  }
}

.action-buttons {
  display: flex;
  gap: var(--spacing-md);
  justify-content: center;
  margin-bottom: var(--spacing-xl);
  flex-wrap: wrap;
}

.tips-section {
  text-align: left;
  
  h4 {
    font-size: var(--font-size-h5);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-md) 0;
  }
}

.tips-list {
  list-style: none;
  padding: 0;
  margin: 0;
  
  li {
    font-size: var(--font-size-medium);
    color: var(--text-secondary);
    margin-bottom: var(--spacing-sm);
    padding-left: var(--spacing-lg);
    position: relative;
    
    &:before {
      content: '•';
      color: var(--primary-color);
      font-weight: var(--font-weight-bold);
      position: absolute;
      left: 0;
    }
    
    &:last-child {
      margin-bottom: 0;
    }
  }
}

// 响应式设计
@media (max-width: 768px) {
  .booking-success-page {
    padding: var(--spacing-lg);
  }
  
  .success-container {
    padding: var(--spacing-lg);
  }
  
  .booking-details .detail-row {
    flex-direction: column;
    gap: var(--spacing-sm);
  }
  
  .action-buttons {
    flex-direction: column;
  }
  
  .action-buttons .el-button {
    width: 100%;
  }
}
</style>