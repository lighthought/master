<template>
  <div class="bookings-page">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1 class="page-title">我的预约</h1>
      <p class="page-subtitle">管理你的大师指导预约</p>
    </div>
    
    <!-- 状态筛选 -->
    <div class="status-filter">
      <el-radio-group v-model="currentStatus" @change="handleStatusChange">
        <el-radio-button label="all">全部</el-radio-button>
        <el-radio-button label="pending">待确认</el-radio-button>
        <el-radio-button label="confirmed">已确认</el-radio-button>
        <el-radio-button label="completed">已完成</el-radio-button>
        <el-radio-button label="cancelled">已取消</el-radio-button>
      </el-radio-group>
    </div>
    
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="3" animated />
      <div class="skeleton-cards">
        <el-skeleton-item v-for="i in 3" :key="i" variant="rect" style="width: 100%; height: 200px; margin-bottom: 16px;" />
      </div>
    </div>
    
    <!-- 预约列表 -->
    <div v-else-if="bookings.length > 0" class="bookings-list">
      <div 
        v-for="booking in bookings" 
        :key="booking.id"
        class="booking-card"
      >
        <!-- 预约头部 -->
        <div class="booking-header">
          <div class="mentor-info">
            <el-avatar :size="50" :src="booking.mentorAvatar" />
            <div class="mentor-details">
              <h3 class="mentor-name">{{ booking.mentorName }}</h3>
              <div class="booking-meta">
                <span class="booking-id">预约号：{{ booking.id }}</span>
                <span class="booking-date">{{ formatDate(booking.createdAt) }}</span>
              </div>
            </div>
          </div>
          
          <div class="booking-status">
            <el-tag 
              :type="getStatusType(booking.status)"
              size="large"
            >
              {{ getStatusText(booking.status) }}
            </el-tag>
          </div>
        </div>
        
        <!-- 预约详情 -->
        <div class="booking-details">
          <div class="detail-row">
            <div class="detail-item">
              <span class="label">指导时间：</span>
              <span class="value">{{ formatDate(booking.date) }} {{ booking.timeSlot }}</span>
            </div>
            <div class="detail-item">
              <span class="label">指导方式：</span>
              <span class="value">{{ getMethodText(booking.method) }}</span>
            </div>
          </div>
          
          <div class="detail-row">
            <div class="detail-item">
              <span class="label">指导时长：</span>
              <span class="value">{{ booking.duration || 1 }}小时</span>
            </div>
            <div class="detail-item">
              <span class="label">指导价格：</span>
              <span class="value price">¥{{ booking.price }}</span>
            </div>
          </div>
          
          <div class="detail-row" v-if="booking.requirements">
            <div class="detail-item full-width">
              <span class="label">指导需求：</span>
              <span class="value">{{ booking.requirements }}</span>
            </div>
          </div>
        </div>
        
        <!-- 预约操作 -->
        <div class="booking-actions">
          <div class="action-buttons">
            <el-button 
              v-if="booking.status === 'pending'"
              type="danger" 
              size="small"
              @click="cancelBooking(booking)"
            >
              取消预约
            </el-button>
            
            <el-button 
              v-if="booking.status === 'confirmed'"
              type="primary" 
              size="small"
              @click="joinMeeting(booking)"
            >
              加入会议
            </el-button>
            
            <el-button 
              v-if="booking.status === 'completed'"
              type="success" 
              size="small"
              @click="writeReview(booking)"
            >
              写评价
            </el-button>
            
            <el-button 
              type="default" 
              size="small"
              @click="viewDetail(booking)"
            >
              查看详情
            </el-button>
            
            <el-button 
              type="default" 
              size="small"
              @click="contactMentor(booking)"
            >
              联系大师
            </el-button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 空状态 -->
    <div v-else class="empty-state">
      <el-empty :description="getEmptyText()">
        <el-button type="primary" @click="$router.push('/mentors')">
          寻找大师
        </el-button>
      </el-empty>
    </div>
    
    <!-- 预约详情对话框 -->
    <el-dialog
      v-model="showDetailDialog"
      title="预约详情"
      width="90%"
      max-width="600px"
    >
      <div v-if="selectedBooking" class="booking-detail-content">
        <div class="detail-section">
          <h4>大师信息</h4>
          <div class="mentor-summary">
            <el-avatar :size="60" :src="selectedBooking.mentorAvatar" />
            <div class="summary-info">
              <h3>{{ selectedBooking.mentorName }}</h3>
              <p>专业领域：前端开发</p>
            </div>
          </div>
        </div>
        
        <div class="detail-section">
          <h4>预约信息</h4>
          <div class="info-grid">
            <div class="info-item">
              <span class="label">预约号：</span>
              <span class="value">{{ selectedBooking.id }}</span>
            </div>
            <div class="info-item">
              <span class="label">预约时间：</span>
              <span class="value">{{ formatDate(selectedBooking.date) }} {{ selectedBooking.timeSlot }}</span>
            </div>
            <div class="info-item">
              <span class="label">指导方式：</span>
              <span class="value">{{ getMethodText(selectedBooking.method) }}</span>
            </div>
            <div class="info-item">
              <span class="label">指导时长：</span>
              <span class="value">{{ selectedBooking.duration || 1 }}小时</span>
            </div>
            <div class="info-item">
              <span class="label">指导价格：</span>
              <span class="value price">¥{{ selectedBooking.price }}</span>
            </div>
            <div class="info-item">
              <span class="label">预约状态：</span>
              <span class="value">
                <el-tag :type="getStatusType(selectedBooking.status)">
                  {{ getStatusText(selectedBooking.status) }}
                </el-tag>
              </span>
            </div>
          </div>
        </div>
        
        <div class="detail-section" v-if="selectedBooking.requirements">
          <h4>指导需求</h4>
          <p class="requirements-text">{{ selectedBooking.requirements }}</p>
        </div>
        
        <div class="detail-section" v-if="selectedBooking.mentorMessage">
          <h4>大师留言</h4>
          <p class="message-text">{{ selectedBooking.mentorMessage }}</p>
        </div>
      </div>
    </el-dialog>
    
    <!-- 取消预约确认对话框 -->
    <el-dialog
      v-model="showCancelDialog"
      title="取消预约"
      width="90%"
      max-width="400px"
    >
      <div class="cancel-content">
        <p>确定要取消这个预约吗？</p>
        <p class="warning-text">取消后可能无法恢复，请谨慎操作。</p>
        
        <el-form :model="cancelForm" label-width="80px">
          <el-form-item label="取消原因">
            <el-select v-model="cancelForm.reason" placeholder="选择取消原因" style="width: 100%">
              <el-option label="时间冲突" value="time_conflict" />
              <el-option label="需求变更" value="requirement_change" />
              <el-option label="个人原因" value="personal_reason" />
              <el-option label="其他原因" value="other" />
            </el-select>
          </el-form-item>
          
          <el-form-item label="备注" v-if="cancelForm.reason === 'other'">
            <el-input
              v-model="cancelForm.note"
              type="textarea"
              :rows="3"
              placeholder="请说明具体原因..."
            />
          </el-form-item>
        </el-form>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showCancelDialog = false">取消</el-button>
          <el-button type="danger" @click="confirmCancel" :loading="cancelling">
            确认取消
          </el-button>
        </div>
      </template>
    </el-dialog>
    
    <!-- 写评价对话框 -->
    <el-dialog
      v-model="showReviewDialog"
      title="写评价"
      width="90%"
      max-width="500px"
    >
      <div v-if="selectedBooking" class="review-content">
        <div class="mentor-summary">
          <el-avatar :size="50" :src="selectedBooking.mentorAvatar" />
          <div class="summary-info">
            <h3>{{ selectedBooking.mentorName }}</h3>
            <p>指导时间：{{ formatDate(selectedBooking.date) }} {{ selectedBooking.timeSlot }}</p>
          </div>
        </div>
        
        <el-form :model="reviewForm" label-width="80px">
          <el-form-item label="评分">
            <el-rate
              v-model="reviewForm.rating"
              :max="5"
              show-score
              text-color="#ff9900"
            />
          </el-form-item>
          
          <el-form-item label="评价内容">
            <el-input
              v-model="reviewForm.content"
              type="textarea"
              :rows="4"
              placeholder="请分享你的学习体验和感受..."
            />
          </el-form-item>
        </el-form>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showReviewDialog = false">取消</el-button>
          <el-button type="primary" @click="submitReview" :loading="submittingReview">
            提交评价
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'

// 路由
const router = useRouter()

// 认证store
const authStore = useAuthStore()

// 状态
const loading = ref(false)
const showDetailDialog = ref(false)
const showCancelDialog = ref(false)
const showReviewDialog = ref(false)
const cancelling = ref(false)
const submittingReview = ref(false)
const currentStatus = ref('all')

// 数据
const bookings = ref<any[]>([])
const selectedBooking = ref<any>(null)

// 表单
const cancelForm = ref({
  reason: '',
  note: ''
})

const reviewForm = ref({
  rating: 5,
  content: ''
})

// 加载预约列表
const loadBookings = async () => {
  if (!authStore.user) return
  
  loading.value = true
  try {
    const status = currentStatus.value === 'all' ? undefined : currentStatus.value
    const result = await ApiService.bookings.getUserBookings(authStore.user.id, status)
    bookings.value = result.data
  } catch (error) {
    console.error('加载预约列表失败:', error)
    ElMessage.error('加载预约列表失败')
  } finally {
    loading.value = false
  }
}

// 处理状态筛选
const handleStatusChange = () => {
  loadBookings()
}

// 获取状态类型
const getStatusType = (status: string): 'warning' | 'primary' | 'success' | 'info' => {
  const types: Record<string, 'warning' | 'primary' | 'success' | 'info'> = {
    pending: 'warning',
    confirmed: 'primary',
    completed: 'success',
    cancelled: 'info'
  }
  return types[status] || 'info'
}

// 获取状态文本
const getStatusText = (status: string) => {
  const texts: Record<string, string> = {
    pending: '待确认',
    confirmed: '已确认',
    completed: '已完成',
    cancelled: '已取消'
  }
  return texts[status] || status
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

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('zh-CN')
}

// 获取空状态文本
const getEmptyText = () => {
  const texts: Record<string, string> = {
    all: '暂无预约记录',
    pending: '暂无待确认的预约',
    confirmed: '暂无已确认的预约',
    completed: '暂无已完成的预约',
    cancelled: '暂无已取消的预约'
  }
  return texts[currentStatus.value] || '暂无预约记录'
}

// 查看详情
const viewDetail = (booking: any) => {
  selectedBooking.value = booking
  showDetailDialog.value = true
}

// 取消预约
const cancelBooking = (booking: any) => {
  selectedBooking.value = booking
  showCancelDialog.value = true
}

// 确认取消
const confirmCancel = async () => {
  if (!selectedBooking.value) return
  
  if (!cancelForm.value.reason) {
    ElMessage.warning('请选择取消原因')
    return
  }
  
  cancelling.value = true
  try {
    const reason = cancelForm.value.reason === 'other' 
      ? cancelForm.value.note 
      : cancelForm.value.reason
    
    await ApiService.bookings.cancelBooking(selectedBooking.value.id, reason)
    
    ElMessage.success('预约已取消')
    showCancelDialog.value = false
    
    // 重置表单
    cancelForm.value = {
      reason: '',
      note: ''
    }
    
    // 重新加载列表
    loadBookings()
  } catch (error) {
    ElMessage.error('取消预约失败，请重试')
  } finally {
    cancelling.value = false
  }
}

// 加入会议
const joinMeeting = (booking: any) => {
  ElMessage.info(`加入会议：${booking.mentorName}的指导会议`)
  // 实际应用中会跳转到会议页面或打开会议链接
}

// 写评价
const writeReview = (booking: any) => {
  selectedBooking.value = booking
  showReviewDialog.value = true
}

// 提交评价
const submitReview = async () => {
  if (!selectedBooking.value || !authStore.user) return
  
  if (!reviewForm.value.content.trim()) {
    ElMessage.warning('请填写评价内容')
    return
  }
  
  submittingReview.value = true
  try {
    // 模拟提交评价
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    ElMessage.success('评价提交成功')
    showReviewDialog.value = false
    
    // 重置表单
    reviewForm.value = {
      rating: 5,
      content: ''
    }
    
    // 重新加载列表
    loadBookings()
  } catch (error) {
    ElMessage.error('提交评价失败，请重试')
  } finally {
    submittingReview.value = false
  }
}

// 联系大师
const contactMentor = (booking: any) => {
  ElMessage.info(`联系大师：${booking.mentorName}`)
  // 实际应用中会打开聊天窗口或跳转到消息页面
}

// 组件挂载时加载数据
onMounted(() => {
  loadBookings()
})
</script>

<style scoped lang="scss">
.bookings-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: var(--spacing-xl);
}

.page-header {
  text-align: center;
  margin-bottom: var(--spacing-xl);
}

.page-title {
  font-size: var(--font-size-h1);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-sm) 0;
}

.page-subtitle {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin: 0;
}

.status-filter {
  display: flex;
  justify-content: center;
  margin-bottom: var(--spacing-xl);
}

.loading-container {
  .skeleton-cards {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-md);
  }
}

.bookings-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.booking-card {
  background: var(--bg-card);
  border: 1px solid var(--border-color-light);
  border-radius: var(--border-radius-medium);
  padding: var(--spacing-lg);
  transition: all var(--transition-normal);
  
  &:hover {
    box-shadow: var(--shadow-light);
  }
}

.booking-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--spacing-lg);
}

.mentor-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.mentor-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.mentor-name {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0;
}

.booking-meta {
  display: flex;
  gap: var(--spacing-md);
  font-size: var(--font-size-small);
  color: var(--text-secondary);
}

.booking-details {
  margin-bottom: var(--spacing-lg);
}

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

.booking-actions {
  border-top: 1px solid var(--border-color-light);
  padding-top: var(--spacing-lg);
}

.action-buttons {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
}

.empty-state {
  text-align: center;
  padding: var(--spacing-xxl) 0;
}

.booking-detail-content {
  .detail-section {
    margin-bottom: var(--spacing-xl);
    
    &:last-child {
      margin-bottom: 0;
    }
    
    h4 {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-md) 0;
    }
  }
  
  .mentor-summary {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
  }
  
  .summary-info h3 {
    font-size: var(--font-size-h5);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-xs) 0;
  }
  
  .summary-info p {
    font-size: var(--font-size-small);
    color: var(--text-secondary);
    margin: 0;
  }
  
  .info-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: var(--spacing-md);
  }
  
  .info-item {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
  }
  
  .requirements-text,
  .message-text {
    font-size: var(--font-size-medium);
    color: var(--text-secondary);
    line-height: 1.6;
    margin: 0;
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
  }
}

.cancel-content {
  p {
    font-size: var(--font-size-medium);
    color: var(--text-primary);
    margin-bottom: var(--spacing-md);
  }
  
  .warning-text {
    color: #e6a23c;
    font-weight: var(--font-weight-medium);
  }
}

.review-content {
  .mentor-summary {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
    margin-bottom: var(--spacing-lg);
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
  }
  
  .summary-info h3 {
    font-size: var(--font-size-h5);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-xs) 0;
  }
  
  .summary-info p {
    font-size: var(--font-size-small);
    color: var(--text-secondary);
    margin: 0;
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
}

// 响应式设计
@media (max-width: 768px) {
  .bookings-page {
    padding: var(--spacing-lg);
  }
  
  .booking-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-md);
  }
  
  .detail-row {
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