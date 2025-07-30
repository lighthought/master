<template>
  <div class="recommended-mentors">
    <div class="section-header">
      <div class="header-content">
        <h2 class="section-title">推荐大师</h2>
        <p class="section-subtitle">为你精选的优质导师，助你快速提升技能</p>
      </div>
      <div class="header-actions">
        <el-button type="text" @click="refreshRecommendations">
          <el-icon><Refresh /></el-icon>
          刷新推荐
        </el-button>
        <el-button type="primary" @click="$router.push('/mentors')">
          查看更多
          <el-icon><ArrowRight /></el-icon>
        </el-button>
      </div>
    </div>
    
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="1" animated />
      <div class="skeleton-cards">
        <el-skeleton-item v-for="i in 4" :key="i" variant="card" style="width: 280px; height: 200px; margin-right: 16px;" />
      </div>
    </div>
    
    <!-- 推荐大师列表 -->
    <div v-else class="mentors-container">
      <div class="mentors-scroll">
        <div 
          v-for="mentor in recommendedMentors" 
          :key="mentor.id"
          class="mentor-card"
          @click="viewMentorDetail(mentor)"
        >
          <!-- 大师头像和状态 -->
          <div class="mentor-header">
            <div class="avatar-container">
              <el-avatar 
                :size="60" 
                :src="mentor.avatar"
                :icon="User"
              />
              <div 
                class="online-indicator"
                :class="{ 'online': mentor.isOnline }"
              ></div>
            </div>
            <div class="mentor-info">
              <div class="mentor-name">
                {{ mentor.name }}
                <el-icon v-if="mentor.isVerified" class="verified-badge">
                  <Check />
                </el-icon>
              </div>
              <div class="mentor-domain">{{ mentor.domain }}</div>
              <div class="mentor-tags">
                <el-tag 
                  v-if="mentor.isOnline" 
                  type="success" 
                  size="small"
                  class="online-tag"
                >
                  在线
                </el-tag>
                <el-tag 
                  v-if="mentor.isVerified" 
                  type="success" 
                  size="small"
                >
                  已认证
                </el-tag>
              </div>
            </div>
          </div>
          
          <!-- 评分和学生数量 -->
          <div class="mentor-stats">
            <div class="rating-section">
              <el-rate 
                v-model="mentor.rating" 
                disabled 
                show-score 
                text-color="#ff9900"
                score-template="{value}"
                class="mentor-rating"
              />
              <span class="student-count">{{ mentor.studentCount }}名学生</span>
            </div>
          </div>
          
          <!-- 技能标签 -->
          <div class="mentor-skills">
            <el-tag 
              v-for="skill in mentor.skills.slice(0, 3)" 
              :key="skill"
              size="small"
              class="skill-tag"
            >
              {{ skill }}
            </el-tag>
            <span v-if="mentor.skills.length > 3" class="more-skills">
              +{{ mentor.skills.length - 3 }}
            </span>
          </div>
          
          <!-- 价格和预约 -->
          <div class="mentor-footer">
            <div class="price-info">
              <span class="price-label">指导价格</span>
              <span class="price-value">¥{{ mentor.price }}/小时</span>
            </div>
            <el-button 
              type="primary" 
              size="small"
              class="book-button"
              @click.stop="bookMentor(mentor)"
            >
              立即预约
            </el-button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 空状态 -->
    <div v-if="!loading && recommendedMentors.length === 0" class="empty-state">
      <el-empty description="暂无推荐大师">
        <el-button type="primary" @click="$router.push('/mentors')">
          浏览所有大师
        </el-button>
      </el-empty>
    </div>
    
    <!-- 预约对话框 -->
    <el-dialog
      v-model="showBookingDialog"
      title="预约大师指导"
      width="90%"
      max-width="500px"
      :close-on-click-modal="false"
    >
      <div v-if="selectedMentor" class="booking-content">
        <div class="mentor-summary">
          <el-avatar :size="50" :src="selectedMentor.avatar" />
          <div class="summary-info">
            <h3>{{ selectedMentor.name }}</h3>
            <p>{{ selectedMentor.domain }}</p>
            <div class="summary-rating">
              <el-rate v-model="selectedMentor.rating" disabled size="small" />
              <span>{{ selectedMentor.rating }}分</span>
            </div>
          </div>
        </div>
        
        <el-form :model="bookingForm" label-width="80px">
          <el-form-item label="指导时间">
            <el-date-picker
              v-model="bookingForm.date"
              type="date"
              placeholder="选择日期"
              :disabled-date="disabledDate"
              style="width: 100%"
            />
          </el-form-item>
          
          <el-form-item label="时间段">
            <el-select v-model="bookingForm.timeSlot" placeholder="选择时间段" style="width: 100%">
              <el-option label="09:00-10:00" value="09:00-10:00" />
              <el-option label="10:00-11:00" value="10:00-11:00" />
              <el-option label="14:00-15:00" value="14:00-15:00" />
              <el-option label="15:00-16:00" value="15:00-16:00" />
              <el-option label="19:00-20:00" value="19:00-20:00" />
              <el-option label="20:00-21:00" value="20:00-21:00" />
            </el-select>
          </el-form-item>
          
          <el-form-item label="指导方式">
            <el-radio-group v-model="bookingForm.method">
              <el-radio label="video">视频通话</el-radio>
              <el-radio label="voice">语音通话</el-radio>
              <el-radio label="text">文字指导</el-radio>
            </el-radio-group>
          </el-form-item>
          
          <el-form-item label="指导需求">
            <el-input
              v-model="bookingForm.requirements"
              type="textarea"
              :rows="3"
              placeholder="请描述你的学习需求和问题..."
            />
          </el-form-item>
        </el-form>
        
        <div class="booking-summary">
          <div class="summary-item">
            <span>指导时长：</span>
            <span>1小时</span>
          </div>
          <div class="summary-item">
            <span>指导价格：</span>
            <span class="price">¥{{ selectedMentor.price }}</span>
          </div>
        </div>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showBookingDialog = false">取消</el-button>
          <el-button type="primary" @click="submitBooking" :loading="bookingLoading">
            确认预约
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  User, Check, Refresh, ArrowRight 
} from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'

// 路由
const router = useRouter()

// 认证store
const authStore = useAuthStore()

// 状态
const loading = ref(false)
const showBookingDialog = ref(false)
const bookingLoading = ref(false)
const selectedMentor = ref<any>(null)

// 推荐大师列表
const recommendedMentors = ref<any[]>([])

// 预约表单
const bookingForm = ref({
  date: '',
  timeSlot: '',
  method: 'video',
  requirements: ''
})

// 加载推荐大师
const loadRecommendedMentors = async () => {
  if (!authStore.user) return
  
  loading.value = true
  try {
    const result = await ApiService.mentors.getRecommendedMentors(authStore.user.id)
    recommendedMentors.value = result.data
  } catch (error) {
    console.error('加载推荐大师失败:', error)
    ElMessage.error('加载推荐大师失败')
  } finally {
    loading.value = false
  }
}

// 刷新推荐
const refreshRecommendations = async () => {
  await loadRecommendedMentors()
  ElMessage.success('推荐已刷新')
}

// 查看大师详情
const viewMentorDetail = (mentor: any) => {
  router.push(`/mentors/${mentor.id}`)
}

// 预约大师
const bookMentor = (mentor: any) => {
  selectedMentor.value = mentor
  showBookingDialog.value = true
}

// 提交预约
const submitBooking = async () => {
  if (!selectedMentor.value || !authStore.user) return
  
  if (!bookingForm.value.date || !bookingForm.value.timeSlot) {
    ElMessage.warning('请选择指导时间和时间段')
    return
  }
  
  bookingLoading.value = true
  try {
    const result = await ApiService.bookings.createBooking({
      mentorId: selectedMentor.value.id,
      userId: authStore.user.id,
      date: bookingForm.value.date,
      timeSlot: bookingForm.value.timeSlot,
      method: bookingForm.value.method,
      requirements: bookingForm.value.requirements
    })
    
    ElMessage.success('预约成功！大师会尽快确认你的预约')
    showBookingDialog.value = false
    
    // 跳转到预约成功页面
    router.push(`/booking-success/${result.data.id}`)
  } catch (error) {
    ElMessage.error('预约失败，请重试')
  } finally {
    bookingLoading.value = false
  }
}

// 禁用过去的日期
const disabledDate = (time: Date) => {
  return time.getTime() < Date.now() - 8.64e7
}

// 组件挂载时加载数据
onMounted(() => {
  loadRecommendedMentors()
})
</script>

<style scoped lang="scss">
.recommended-mentors {
  padding: var(--spacing-xl);
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
  margin-bottom: var(--spacing-xl);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--spacing-xl);
  gap: var(--spacing-lg);
}

.header-content {
  flex: 1;
}

.section-title {
  font-size: var(--font-size-h3);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-sm) 0;
}

.section-subtitle {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin: 0;
}

.header-actions {
  display: flex;
  gap: var(--spacing-md);
  align-items: center;
}

.loading-container {
  .skeleton-cards {
    display: flex;
    gap: var(--spacing-md);
    overflow-x: auto;
    padding: var(--spacing-md) 0;
  }
}

.mentors-container {
  overflow: hidden;
}

.mentors-scroll {
  display: flex;
  gap: var(--spacing-md);
  overflow-x: auto;
  padding: var(--spacing-sm) 0;
  
  &::-webkit-scrollbar {
    height: 6px;
  }
  
  &::-webkit-scrollbar-track {
    background: var(--bg-tertiary);
    border-radius: 3px;
  }
  
  &::-webkit-scrollbar-thumb {
    background: var(--primary-color);
    border-radius: 3px;
  }
}

.mentor-card {
  flex-shrink: 0;
  width: 280px;
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  cursor: pointer;
  transition: all var(--transition-normal);
  border: 1px solid transparent;
  
  &:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-light);
    border-color: var(--primary-color);
  }
}

.mentor-header {
  display: flex;
  align-items: flex-start;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-md);
}

.avatar-container {
  position: relative;
  flex-shrink: 0;
}

.online-indicator {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: #ccc;
  border: 2px solid var(--bg-secondary);
  
  &.online {
    background: #67c23a;
  }
}

.mentor-info {
  flex: 1;
  min-width: 0;
}

.mentor-name {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: 16px;
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
}

.verified-badge {
  color: #67c23a;
  font-size: 14px;
}

.mentor-domain {
  font-size: 14px;
  color: var(--text-secondary);
  margin-bottom: var(--spacing-xs);
}

.mentor-tags {
  display: flex;
  gap: var(--spacing-xs);
}

.online-tag {
  background: rgba(103, 194, 58, 0.1);
  border-color: #67c23a;
  color: #67c23a;
}

.mentor-stats {
  margin-bottom: var(--spacing-md);
}

.rating-section {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
}

.mentor-rating {
  :deep(.el-rate__icon) {
    font-size: 14px;
  }
}

.student-count {
  font-size: 12px;
  color: var(--text-secondary);
}

.mentor-skills {
  display: flex;
  gap: var(--spacing-xs);
  margin-bottom: var(--spacing-md);
  flex-wrap: wrap;
}

.skill-tag {
  background: rgba(64, 158, 255, 0.1);
  border-color: var(--primary-color);
  color: var(--primary-color);
}

.more-skills {
  font-size: 12px;
  color: var(--text-secondary);
  align-self: center;
}

.mentor-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price-info {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.price-label {
  font-size: 12px;
  color: var(--text-secondary);
}

.price-value {
  font-size: 16px;
  font-weight: var(--font-weight-bold);
  color: #ff9900;
}

.book-button {
  background: linear-gradient(135deg, var(--primary-color), var(--apprentice-color));
  border: none;
  
  &:hover {
    background: linear-gradient(135deg, var(--apprentice-color), var(--primary-color));
  }
}

.empty-state {
  text-align: center;
  padding: var(--spacing-xxl) 0;
}

.booking-content {
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
    margin: 0 0 var(--spacing-xs) 0;
  }
  
  .summary-rating {
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
    font-size: var(--font-size-small);
    color: var(--text-secondary);
  }
}

.booking-summary {
  margin-top: var(--spacing-lg);
  padding: var(--spacing-md);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  
  .summary-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: var(--spacing-sm);
    
    &:last-child {
      margin-bottom: 0;
    }
    
    .price {
      font-weight: var(--font-weight-bold);
      color: #ff9900;
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
  .recommended-mentors {
    padding: var(--spacing-lg);
    margin: var(--spacing-md);
  }
  
  .section-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-md);
  }
  
  .header-actions {
    width: 100%;
    justify-content: space-between;
  }
  
  .mentor-card {
    width: 260px;
    padding: var(--spacing-md);
  }
  
  .mentor-header {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
  
  .mentor-footer {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: stretch;
  }
  
  .book-button {
    width: 100%;
  }
}
</style>