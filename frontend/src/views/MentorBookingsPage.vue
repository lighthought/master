<template>
  <div class="mentor-bookings-page">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1 class="page-title">预约管理</h1>
      <p class="page-subtitle">管理你的学生预约请求</p>
    </div>
    
    <!-- 预约统计 -->
    <BookingStats :mentor-id="authStore.user?.id" />
    
    <!-- 筛选和操作栏 -->
    <div class="filter-actions-bar">
      <div class="filter-section">
        <el-radio-group v-model="currentStatus" @change="handleStatusChange">
          <el-radio-button label="all">全部</el-radio-button>
          <el-radio-button label="pending">待确认</el-radio-button>
          <el-radio-button label="confirmed">已确认</el-radio-button>
          <el-radio-button label="completed">已完成</el-radio-button>
        </el-radio-group>
        
        <el-date-picker
          v-model="dateRange"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          @change="handleDateChange"
          style="margin-left: 16px;"
        />
      </div>
      
      <div class="actions-section">
        <el-button 
          type="primary" 
          @click="handleBatchConfirm"
          :disabled="selectedBookings.length === 0"
        >
          <el-icon><Check /></el-icon>
          批量确认 ({{ selectedBookings.length }})
        </el-button>
        
        <el-button 
          type="danger" 
          @click="handleBatchReject"
          :disabled="selectedBookings.length === 0"
        >
          <el-icon><Close /></el-icon>
          批量拒绝 ({{ selectedBookings.length }})
        </el-button>
      </div>
    </div>
    
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="3" animated />
      <div class="skeleton-cards">
        <el-skeleton-item v-for="i in 3" :key="i" variant="image" style="width: 100%; height: 200px; margin-bottom: 16px;" />
      </div>
    </div>
    
    <!-- 预约列表 -->
    <div v-else-if="bookings.length > 0" class="bookings-list">
      <div 
        v-for="booking in bookings" 
        :key="booking.id"
        class="booking-card"
        :class="{ 'selected': selectedBookings.includes(booking.id) }"
      >
        <!-- 选择框 -->
        <div class="booking-select">
          <el-checkbox 
            v-model="selectedBookings" 
            :value="booking.id"
            @change="handleSelectionChange"
          />
        </div>
        
        <!-- 预约头部 -->
        <div class="booking-header">
          <div class="student-info">
            <el-avatar :size="50" :src="booking.studentAvatar" />
            <div class="student-details">
              <h3 class="student-name">{{ booking.studentName }}</h3>
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
              type="success" 
              size="small"
              @click="confirmBooking(booking)"
            >
              确认预约
            </el-button>
            
            <el-button 
              v-if="booking.status === 'pending'"
              type="danger" 
              size="small"
              @click="rejectBooking(booking)"
            >
              拒绝预约
            </el-button>
            
            <el-button 
              v-if="booking.status === 'confirmed'"
              type="primary" 
              size="small"
              @click="startMeeting(booking)"
            >
              开始指导
            </el-button>
            
            <el-button 
              v-if="booking.status === 'confirmed'"
              type="warning" 
              size="small"
              @click="rescheduleBooking(booking)"
            >
              修改时间
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
              @click="contactStudent(booking)"
            >
              联系学生
            </el-button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 空状态 -->
    <div v-else class="empty-state">
      <el-empty :description="getEmptyText()">
        <el-button type="primary" @click="refreshBookings">
          刷新列表
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
          <h4>学生信息</h4>
          <div class="student-summary">
            <el-avatar :size="60" :src="selectedBooking.studentAvatar" />
            <div class="summary-info">
              <h3>{{ selectedBooking.studentName }}</h3>
              <p>学习领域：{{ selectedBooking.studentDomain || '未指定' }}</p>
              <p>学习目标：{{ selectedBooking.studentGoal || '未指定' }}</p>
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
    
    <!-- 拒绝预约对话框 -->
    <el-dialog
      v-model="showRejectDialog"
      title="拒绝预约"
      width="90%"
      max-width="400px"
    >
      <div class="reject-content">
        <p>确定要拒绝这个预约吗？</p>
        
        <el-form :model="rejectForm" label-width="80px">
          <el-form-item label="拒绝原因">
            <el-select v-model="rejectForm.reason" placeholder="选择拒绝原因" style="width: 100%">
              <el-option label="时间冲突" value="time_conflict" />
              <el-option label="专业不匹配" value="skill_mismatch" />
              <el-option label="个人原因" value="personal_reason" />
              <el-option label="其他原因" value="other" />
            </el-select>
          </el-form-item>
          
          <el-form-item label="备注" v-if="rejectForm.reason === 'other'">
            <el-input
              v-model="rejectForm.note"
              type="textarea"
              :rows="3"
              placeholder="请说明具体原因..."
            />
          </el-form-item>
        </el-form>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showRejectDialog = false">取消</el-button>
          <el-button type="danger" @click="confirmReject" :loading="rejecting">
            确认拒绝
          </el-button>
        </div>
      </template>
    </el-dialog>
    
    <!-- 修改时间对话框 -->
    <el-dialog
      v-model="showRescheduleDialog"
      title="修改预约时间"
      width="90%"
      max-width="500px"
    >
      <div v-if="selectedBooking" class="reschedule-content">
        <div class="current-time">
          <h4>当前预约时间</h4>
          <p>{{ formatDate(selectedBooking.date) }} {{ selectedBooking.timeSlot }}</p>
        </div>
        
        <el-form :model="rescheduleForm" label-width="100px">
          <el-form-item label="新预约日期">
            <el-date-picker
              v-model="rescheduleForm.date"
              type="date"
              placeholder="选择新日期"
              style="width: 100%;"
            />
          </el-form-item>
          
          <el-form-item label="新时间段">
            <el-time-select
              v-model="rescheduleForm.timeSlot"
              start="09:00"
              step="00:30"
              end="22:00"
              placeholder="选择新时间"
              style="width: 100%;"
            />
          </el-form-item>
          
          <el-form-item label="修改原因">
            <el-input
              v-model="rescheduleForm.reason"
              type="textarea"
              :rows="3"
              placeholder="请说明修改时间的原因..."
            />
          </el-form-item>
        </el-form>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showRescheduleDialog = false">取消</el-button>
          <el-button type="primary" @click="confirmReschedule" :loading="rescheduling">
            确认修改
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
import { Clock, Check, Trophy, Calendar, Close } from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'
import BookingStats from '@/components/mentor/BookingStats.vue'

// 路由
const router = useRouter()

// 认证store
const authStore = useAuthStore()

// 状态
const loading = ref(false)
const showDetailDialog = ref(false)
const showRejectDialog = ref(false)
const showRescheduleDialog = ref(false)
const rejecting = ref(false)
const rescheduling = ref(false)
const currentStatus = ref('all')
const dateRange = ref<[Date, Date] | null>(null)
const selectedBookings = ref<string[]>([])

// 数据
const bookings = ref<any[]>([])
const selectedBooking = ref<any>(null)

// 表单
const rejectForm = ref({
  reason: '',
  note: ''
})

const rescheduleForm = ref({
  date: '',
  timeSlot: '',
  reason: ''
})



// 加载预约列表
const loadBookings = async () => {
  if (!authStore.user) return
  
  loading.value = true
  try {
    const status = currentStatus.value === 'all' ? undefined : currentStatus.value
    const result = await ApiService.bookings.getMentorBookings(authStore.user.id, status)
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
  selectedBookings.value = [] // 清空选择
  loadBookings()
}

// 处理日期筛选
const handleDateChange = () => {
  selectedBookings.value = [] // 清空选择
  loadBookings()
}

// 处理选择变化
const handleSelectionChange = () => {
  // 选择变化时的处理逻辑
}

// 获取状态类型
const getStatusType = (status: string) => {
  const types: Record<string, string> = {
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
    completed: '暂无已完成的预约'
  }
  return texts[currentStatus.value] || '暂无预约记录'
}

// 刷新预约列表
const refreshBookings = () => {
  loadBookings()
}

// 查看详情
const viewDetail = (booking: any) => {
  selectedBooking.value = booking
  showDetailDialog.value = true
}

// 确认预约
const confirmBooking = async (booking: any) => {
  try {
    await ElMessageBox.confirm(
      `确定要确认学生 ${booking.studentName} 的预约吗？`,
      '确认预约',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    await ApiService.bookings.updateBookingStatus(booking.id, 'confirmed')
    ElMessage.success('预约已确认')
    loadBookings()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('确认预约失败，请重试')
    }
  }
}

// 拒绝预约
const rejectBooking = (booking: any) => {
  selectedBooking.value = booking
  showRejectDialog.value = true
}

// 确认拒绝
const confirmReject = async () => {
  if (!selectedBooking.value) return
  
  if (!rejectForm.value.reason) {
    ElMessage.warning('请选择拒绝原因')
    return
  }
  
  rejecting.value = true
  try {
    const reason = rejectForm.value.reason === 'other' 
      ? rejectForm.value.note 
      : rejectForm.value.reason
    
    await ApiService.bookings.updateBookingStatus(selectedBooking.value.id, 'cancelled')
    
    ElMessage.success('预约已拒绝')
    showRejectDialog.value = false
    
    // 重置表单
    rejectForm.value = {
      reason: '',
      note: ''
    }
    
    // 重新加载列表
    loadBookings()
  } catch (error) {
    ElMessage.error('拒绝预约失败，请重试')
  } finally {
    rejecting.value = false
  }
}

// 开始指导
const startMeeting = (booking: any) => {
  ElMessage.info(`开始指导：${booking.studentName}的指导会议`)
  // 实际应用中会跳转到会议页面或打开会议链接
}

// 修改时间
const rescheduleBooking = (booking: any) => {
  selectedBooking.value = booking
  rescheduleForm.value = {
    date: booking.date,
    timeSlot: booking.timeSlot,
    reason: ''
  }
  showRescheduleDialog.value = true
}

// 确认修改时间
const confirmReschedule = async () => {
  if (!selectedBooking.value) return
  
  if (!rescheduleForm.value.date || !rescheduleForm.value.timeSlot) {
    ElMessage.warning('请选择新的预约时间')
    return
  }
  
  rescheduling.value = true
  try {
    // 模拟修改预约时间
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    ElMessage.success('预约时间已修改')
    showRescheduleDialog.value = false
    
    // 重置表单
    rescheduleForm.value = {
      date: '',
      timeSlot: '',
      reason: ''
    }
    
    // 重新加载列表
    loadBookings()
  } catch (error) {
    ElMessage.error('修改时间失败，请重试')
  } finally {
    rescheduling.value = false
  }
}

// 联系学生
const contactStudent = (booking: any) => {
  ElMessage.info(`联系学生：${booking.studentName}`)
  // 实际应用中会打开聊天窗口或跳转到消息页面
}

// 批量确认
const handleBatchConfirm = async () => {
  if (selectedBookings.value.length === 0) return
  
  try {
    await ElMessageBox.confirm(
      `确定要确认选中的 ${selectedBookings.value.length} 个预约吗？`,
      '批量确认',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    // 批量更新状态
    for (const bookingId of selectedBookings.value) {
      await ApiService.bookings.updateBookingStatus(bookingId, 'confirmed')
    }
    
    ElMessage.success(`已确认 ${selectedBookings.value.length} 个预约`)
    selectedBookings.value = []
    loadBookings()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('批量确认失败，请重试')
    }
  }
}

// 批量拒绝
const handleBatchReject = async () => {
  if (selectedBookings.value.length === 0) return
  
  try {
    await ElMessageBox.confirm(
      `确定要拒绝选中的 ${selectedBookings.value.length} 个预约吗？`,
      '批量拒绝',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    // 批量更新状态
    for (const bookingId of selectedBookings.value) {
      await ApiService.bookings.updateBookingStatus(bookingId, 'cancelled')
    }
    
    ElMessage.success(`已拒绝 ${selectedBookings.value.length} 个预约`)
    selectedBookings.value = []
    loadBookings()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('批量拒绝失败，请重试')
    }
  }
}

// 组件挂载时加载数据
onMounted(() => {
  loadBookings()
})
</script>

<style scoped lang="scss">
.mentor-bookings-page {
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



.filter-actions-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-xl);
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
}

.filter-section {
  display: flex;
  align-items: center;
}

.actions-section {
  display: flex;
  gap: var(--spacing-md);
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
  display: flex;
  gap: var(--spacing-md);
  
  &:hover {
    box-shadow: var(--shadow-light);
  }
  
  &.selected {
    border-color: var(--primary-color);
    background: rgba(64, 158, 255, 0.05);
  }
}

.booking-select {
  display: flex;
  align-items: flex-start;
  padding-top: var(--spacing-sm);
}

.booking-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--spacing-lg);
  flex: 1;
}

.student-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.student-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.student-name {
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
  flex: 1;
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
  flex: 1;
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
  
  .student-summary {
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
    margin: 0 0 var(--spacing-xs) 0;
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

.reject-content {
  p {
    font-size: var(--font-size-medium);
    color: var(--text-primary);
    margin-bottom: var(--spacing-md);
  }
}

.reschedule-content {
  .current-time {
    margin-bottom: var(--spacing-lg);
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
    
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
    }
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
}

// 响应式设计
@media (max-width: 768px) {
  .mentor-bookings-page {
    padding: var(--spacing-lg);
  }
  
  .stats-cards {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .filter-actions-bar {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: stretch;
  }
  
  .filter-section {
    flex-direction: column;
    gap: var(--spacing-md);
  }
  
  .actions-section {
    flex-direction: column;
  }
  
  .booking-card {
    flex-direction: column;
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