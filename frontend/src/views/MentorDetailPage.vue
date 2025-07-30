<template>
  <div class="mentor-detail-page">
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="10" animated />
    </div>
    
    <!-- 大师详情内容 -->
    <div v-else-if="mentor" class="mentor-content">
      <!-- 基本信息区域 -->
      <div class="mentor-header">
        <div class="mentor-profile">
          <div class="avatar-section">
            <el-avatar 
              :size="120" 
              :src="mentor.avatar"
              :icon="User"
            />
            <div 
              class="online-indicator"
              :class="{ 'online': mentor.isOnline }"
            >
              <span class="indicator-text">{{ mentor.isOnline ? '在线' : '离线' }}</span>
            </div>
          </div>
          
          <div class="profile-info">
            <div class="name-section">
              <h1 class="mentor-name">{{ mentor.name }}</h1>
              <div class="verification-badge" v-if="mentor.isVerified">
                <el-icon><Check /></el-icon>
                <span>已认证</span>
              </div>
            </div>
            
            <div class="domain-section">
              <h2 class="mentor-domain">{{ mentor.domain }}</h2>
              <div class="mentor-tags">
                <el-tag 
                  v-if="mentor.isOnline" 
                  type="success" 
                  size="large"
                  class="online-tag"
                >
                  在线
                </el-tag>
                <el-tag 
                  v-if="mentor.isVerified" 
                  type="success" 
                  size="large"
                >
                  已认证
                </el-tag>
              </div>
            </div>
            
            <div class="stats-section">
              <div class="stat-item">
                <div class="stat-value">{{ mentor.rating }}</div>
                <div class="stat-label">评分</div>
                <el-rate v-model="mentor.rating" disabled size="small" />
              </div>
              <div class="stat-item">
                <div class="stat-value">{{ mentor.studentCount }}</div>
                <div class="stat-label">学生数</div>
              </div>
              <div class="stat-item">
                <div class="stat-value">{{ mentor.completionRate }}%</div>
                <div class="stat-label">完成率</div>
              </div>
              <div class="stat-item">
                <div class="stat-value">{{ mentor.responseTime }}</div>
                <div class="stat-label">响应时间</div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="action-section">
          <div class="price-info">
            <div class="price-label">指导价格</div>
            <div class="price-value">¥{{ mentor.price }}/小时</div>
          </div>
          
          <div class="action-buttons">
            <el-button 
              type="primary" 
              size="large"
              class="book-button"
              @click="showBookingDialog = true"
            >
              <el-icon><Calendar /></el-icon>
              立即预约
            </el-button>
            
            <el-button 
              type="default" 
              size="large"
              class="message-button"
              @click="showMessageDialog = true"
            >
              <el-icon><ChatDotRound /></el-icon>
              发送消息
            </el-button>
          </div>
        </div>
      </div>
      
      <!-- 详细信息区域 -->
      <div class="mentor-details">
        <el-tabs v-model="activeTab" class="detail-tabs">
          <!-- 个人介绍 -->
          <el-tab-pane label="个人介绍" name="profile">
            <div class="tab-content">
              <div class="section">
                <h3 class="section-title">个人简介</h3>
                <p class="bio-text">{{ mentor.bio }}</p>
              </div>
              
              <div class="section">
                <h3 class="section-title">专业经历</h3>
                <p class="experience-text">{{ mentor.experience }}</p>
              </div>
              
              <div class="section">
                <h3 class="section-title">技能标签</h3>
                <div class="skills-grid">
                  <el-tag 
                    v-for="skill in mentor.skills" 
                    :key="skill"
                    size="large"
                    class="skill-tag"
                  >
                    {{ skill }}
                  </el-tag>
                </div>
              </div>
              
              <div class="section">
                <h3 class="section-title">服务类型</h3>
                <div class="service-types">
                  <div 
                    v-for="service in mentor.serviceTypes" 
                    :key="service"
                    class="service-item"
                  >
                    <el-icon><Check /></el-icon>
                    <span>{{ service }}</span>
                  </div>
                </div>
              </div>
              
              <div class="section">
                <h3 class="section-title">成就荣誉</h3>
                <div class="achievements">
                  <div 
                    v-for="achievement in mentor.achievements" 
                    :key="achievement"
                    class="achievement-item"
                  >
                    <el-icon><Trophy /></el-icon>
                    <span>{{ achievement }}</span>
                  </div>
                </div>
              </div>
            </div>
          </el-tab-pane>
          
          <!-- 评价 -->
          <el-tab-pane label="学生评价" name="reviews">
            <div class="tab-content">
              <div class="reviews-header">
                <div class="reviews-summary">
                  <div class="overall-rating">
                    <div class="rating-score">{{ mentor.rating }}</div>
                    <div class="rating-stars">
                      <el-rate v-model="mentor.rating" disabled size="large" />
                    </div>
                    <div class="rating-count">{{ reviews.length }}条评价</div>
                  </div>
                </div>
              </div>
              
              <div class="reviews-list">
                <div 
                  v-for="review in reviews" 
                  :key="review.id"
                  class="review-item"
                >
                  <div class="review-header">
                    <div class="reviewer-info">
                      <el-avatar :size="40" :src="review.avatar" />
                      <div class="reviewer-details">
                        <div class="reviewer-name">{{ review.userName }}</div>
                        <div class="review-date">{{ formatDate(review.createdAt) }}</div>
                      </div>
                    </div>
                    <div class="review-rating">
                      <el-rate v-model="review.rating" disabled size="small" />
                    </div>
                  </div>
                  
                  <div class="review-content">
                    <p>{{ review.content }}</p>
                  </div>
                </div>
              </div>
              
              <!-- 加载更多评价 -->
              <div v-if="hasMoreReviews" class="load-more">
                <el-button @click="loadMoreReviews" :loading="loadingReviews">
                  加载更多评价
                </el-button>
              </div>
            </div>
          </el-tab-pane>
          
          <!-- 课程 -->
          <el-tab-pane label="课程" name="courses">
            <div class="tab-content">
              <div class="courses-header">
                <h3 class="section-title">大师课程</h3>
                <p class="section-subtitle">由{{ mentor.name }}开设的专业课程</p>
              </div>
              
              <div class="courses-grid">
                <div 
                  v-for="course in courses" 
                  :key="course.id"
                  class="course-card"
                  @click="viewCourse(course)"
                >
                  <div class="course-image">
                    <img :src="course.cover" :alt="course.title" />
                    <div class="course-overlay">
                      <el-button type="primary" size="small">查看详情</el-button>
                    </div>
                  </div>
                  
                  <div class="course-info">
                    <h4 class="course-title">{{ course.title }}</h4>
                    <p class="course-description">{{ course.description }}</p>
                    
                    <div class="course-meta">
                      <div class="course-stats">
                        <span class="stat">
                          <el-icon><Clock /></el-icon>
                          {{ course.duration }}
                        </span>
                        <span class="stat">
                          <el-icon><User /></el-icon>
                          {{ course.studentCount }}人学习
                        </span>
                        <span class="stat">
                          <el-icon><Star /></el-icon>
                          {{ course.rating }}分
                        </span>
                      </div>
                      
                      <div class="course-price">
                        <span class="price">¥{{ course.price }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- 空状态 -->
              <div v-if="courses.length === 0" class="empty-courses">
                <el-empty description="暂无课程">
                  <el-button type="primary" @click="showBookingDialog = true">
                    预约1对1指导
                  </el-button>
                </el-empty>
              </div>
            </div>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
    
    <!-- 错误状态 -->
    <div v-else class="error-container">
      <el-empty description="大师信息不存在">
        <el-button type="primary" @click="$router.push('/mentors')">
          返回大师列表
        </el-button>
      </el-empty>
    </div>
    
    <!-- 预约对话框 -->
    <el-dialog
      v-model="showBookingDialog"
      title="预约大师指导"
      width="90%"
      max-width="600px"
      :close-on-click-modal="false"
    >
      <div v-if="mentor" class="booking-content">
        <div class="mentor-summary">
          <el-avatar :size="60" :src="mentor.avatar" />
          <div class="summary-info">
            <h3>{{ mentor.name }}</h3>
            <p>{{ mentor.domain }}</p>
            <div class="summary-rating">
              <el-rate v-model="mentor.rating" disabled size="small" />
              <span>{{ mentor.rating }}分</span>
            </div>
          </div>
        </div>
        
        <el-form :model="bookingForm" label-width="100px">
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
          
          <el-form-item label="指导时长">
            <el-radio-group v-model="bookingForm.duration">
              <el-radio :label="1">1小时</el-radio>
              <el-radio :label="2">2小时</el-radio>
              <el-radio :label="3">3小时</el-radio>
            </el-radio-group>
          </el-form-item>
          
          <el-form-item label="指导需求">
            <el-input
              v-model="bookingForm.requirements"
              type="textarea"
              :rows="4"
              placeholder="请详细描述你的学习需求和问题..."
            />
          </el-form-item>
        </el-form>
        
        <div class="booking-summary">
          <div class="summary-item">
            <span>指导时长：</span>
            <span>{{ bookingForm.duration }}小时</span>
          </div>
          <div class="summary-item">
            <span>单价：</span>
            <span>¥{{ mentor.price }}/小时</span>
          </div>
          <div class="summary-item total">
            <span>总价：</span>
            <span class="total-price">¥{{ mentor.price * bookingForm.duration }}</span>
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
    
    <!-- 消息对话框 -->
    <el-dialog
      v-model="showMessageDialog"
      title="发送消息"
      width="90%"
      max-width="500px"
      :close-on-click-modal="false"
    >
      <div class="message-content">
        <el-form :model="messageForm" label-width="80px">
          <el-form-item label="主题">
            <el-input v-model="messageForm.subject" placeholder="消息主题" />
          </el-form-item>
          
          <el-form-item label="内容">
            <el-input
              v-model="messageForm.content"
              type="textarea"
              :rows="6"
              placeholder="请输入消息内容..."
            />
          </el-form-item>
        </el-form>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showMessageDialog = false">取消</el-button>
          <el-button type="primary" @click="sendMessage" :loading="messageLoading">
            发送消息
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  User, Check, Calendar, ChatDotRound, Trophy, Clock, Star 
} from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'

// 路由
const route = useRoute()
const router = useRouter()

// 认证store
const authStore = useAuthStore()

// 状态
const loading = ref(false)
const showBookingDialog = ref(false)
const showMessageDialog = ref(false)
const bookingLoading = ref(false)
const messageLoading = ref(false)
const loadingReviews = ref(false)
const activeTab = ref('profile')

// 数据
const mentor = ref<any>(null)
const reviews = ref<any[]>([])
const courses = ref<any[]>([])
const hasMoreReviews = ref(true)
const currentReviewPage = ref(1)

// 表单
const bookingForm = ref({
  date: '',
  timeSlot: '',
  method: 'video',
  duration: 1,
  requirements: ''
})

const messageForm = ref({
  subject: '',
  content: ''
})

// 计算属性
const mentorId = computed(() => route.params.id as string)

// 加载大师详情
const loadMentorDetail = async () => {
  if (!mentorId.value) return
  
  loading.value = true
  try {
    const result = await ApiService.mentors.getMentorDetail(mentorId.value)
    mentor.value = result.data
  } catch (error) {
    console.error('加载大师详情失败:', error)
    ElMessage.error('加载大师详情失败')
  } finally {
    loading.value = false
  }
}

// 加载评价
const loadReviews = async (page = 1) => {
  if (!mentorId.value) return
  
  try {
    const result = await ApiService.mentors.getMentorReviews(mentorId.value, page)
    if (page === 1) {
      reviews.value = result.data.reviews
    } else {
      reviews.value.push(...result.data.reviews)
    }
    hasMoreReviews.value = result.data.reviews.length > 0
    currentReviewPage.value = page
  } catch (error) {
    console.error('加载评价失败:', error)
  }
}

// 加载更多评价
const loadMoreReviews = async () => {
  loadingReviews.value = true
  await loadReviews(currentReviewPage.value + 1)
  loadingReviews.value = false
}

// 加载课程
const loadCourses = async () => {
  if (!mentorId.value) return
  
  try {
    // 模拟课程数据
    courses.value = [
      {
        id: '1',
        title: 'Vue.js 从入门到精通',
        description: '系统学习Vue.js框架，掌握组件化开发技能',
        cover: 'https://via.placeholder.com/300x200/4CAF50/FFFFFF?text=Vue',
        duration: '20小时',
        studentCount: 156,
        rating: 4.8,
        price: 299
      },
      {
        id: '2',
        title: '前端性能优化实战',
        description: '学习前端性能优化技巧，提升用户体验',
        cover: 'https://via.placeholder.com/300x200/2196F3/FFFFFF?text=Performance',
        duration: '15小时',
        studentCount: 89,
        rating: 4.9,
        price: 199
      }
    ]
  } catch (error) {
    console.error('加载课程失败:', error)
  }
}

// 提交预约
const submitBooking = async () => {
  if (!mentor.value || !authStore.user) return
  
  if (!bookingForm.value.date || !bookingForm.value.timeSlot) {
    ElMessage.warning('请选择指导时间和时间段')
    return
  }
  
  bookingLoading.value = true
  try {
    const result = await ApiService.bookings.createBooking({
      mentorId: mentor.value.id,
      userId: authStore.user.id,
      date: bookingForm.value.date,
      timeSlot: bookingForm.value.timeSlot,
      method: bookingForm.value.method,
      duration: bookingForm.value.duration,
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

// 发送消息
const sendMessage = async () => {
  if (!messageForm.value.subject || !messageForm.value.content) {
    ElMessage.warning('请填写消息主题和内容')
    return
  }
  
  messageLoading.value = true
  try {
    // 模拟发送消息
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    ElMessage.success('消息发送成功')
    showMessageDialog.value = false
    
    // 重置表单
    messageForm.value = {
      subject: '',
      content: ''
    }
  } catch (error) {
    ElMessage.error('发送失败，请重试')
  } finally {
    messageLoading.value = false
  }
}

// 查看课程
const viewCourse = (course: any) => {
  ElMessage.info(`查看课程：${course.title}`)
  // 实际应用中会跳转到课程详情页
  // router.push(`/courses/${course.id}`)
}

// 禁用过去的日期
const disabledDate = (time: Date) => {
  return time.getTime() < Date.now() - 8.64e7
}

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('zh-CN')
}

// 监听标签页切换
const handleTabChange = (tabName: string) => {
  if (tabName === 'reviews' && reviews.value.length === 0) {
    loadReviews()
  } else if (tabName === 'courses' && courses.value.length === 0) {
    loadCourses()
  }
}

// 组件挂载时加载数据
onMounted(() => {
  loadMentorDetail()
  loadReviews()
  loadCourses()
})
</script>

<style scoped lang="scss">
.mentor-detail-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: var(--spacing-xl);
}

.loading-container {
  padding: var(--spacing-xxl) 0;
}

.mentor-content {
  .mentor-header {
    display: flex;
    gap: var(--spacing-xl);
    margin-bottom: var(--spacing-xxl);
    padding: var(--spacing-xl);
    background: var(--bg-card);
    border-radius: var(--border-radius-large);
    box-shadow: var(--shadow-card);
  }
  
  .mentor-profile {
    flex: 1;
    display: flex;
    gap: var(--spacing-lg);
  }
  
  .avatar-section {
    position: relative;
    flex-shrink: 0;
  }
  
  .online-indicator {
    position: absolute;
    bottom: 8px;
    right: 8px;
    background: #ccc;
    color: white;
    padding: 4px 8px;
    border-radius: 12px;
    font-size: 12px;
    font-weight: var(--font-weight-medium);
    
    &.online {
      background: #67c23a;
    }
    
    .indicator-text {
      font-size: 10px;
    }
  }
  
  .profile-info {
    flex: 1;
  }
  
  .name-section {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
    margin-bottom: var(--spacing-md);
  }
  
  .mentor-name {
    font-size: var(--font-size-h1);
    font-weight: var(--font-weight-bold);
    color: var(--text-primary);
    margin: 0;
  }
  
  .verification-badge {
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
    background: rgba(103, 194, 58, 0.1);
    color: #67c23a;
    padding: 4px 8px;
    border-radius: var(--border-radius-small);
    font-size: var(--font-size-small);
    font-weight: var(--font-weight-medium);
  }
  
  .domain-section {
    margin-bottom: var(--spacing-lg);
  }
  
  .mentor-domain {
    font-size: var(--font-size-h3);
    font-weight: var(--font-weight-semibold);
    color: var(--text-secondary);
    margin: 0 0 var(--spacing-sm) 0;
  }
  
  .mentor-tags {
    display: flex;
    gap: var(--spacing-sm);
  }
  
  .online-tag {
    background: rgba(103, 194, 58, 0.1);
    border-color: #67c23a;
    color: #67c23a;
  }
  
  .stats-section {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: var(--spacing-md);
  }
  
  .stat-item {
    text-align: center;
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
  }
  
  .stat-value {
    font-size: var(--font-size-h3);
    font-weight: var(--font-weight-bold);
    color: var(--primary-color);
    margin-bottom: var(--spacing-xs);
  }
  
  .stat-label {
    font-size: var(--font-size-small);
    color: var(--text-secondary);
    margin-bottom: var(--spacing-xs);
  }
  
  .action-section {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-lg);
    min-width: 200px;
  }
  
  .price-info {
    text-align: center;
    padding: var(--spacing-lg);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
  }
  
  .price-label {
    font-size: var(--font-size-small);
    color: var(--text-secondary);
    margin-bottom: var(--spacing-xs);
  }
  
  .price-value {
    font-size: var(--font-size-h2);
    font-weight: var(--font-weight-bold);
    color: #ff9900;
  }
  
  .action-buttons {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-md);
  }
  
  .book-button {
    background: linear-gradient(135deg, var(--primary-color), var(--apprentice-color));
    border: none;
    height: 50px;
    font-size: var(--font-size-medium);
    font-weight: var(--font-weight-medium);
    
    &:hover {
      background: linear-gradient(135deg, var(--apprentice-color), var(--primary-color));
    }
  }
  
  .message-button {
    height: 50px;
    font-size: var(--font-size-medium);
    font-weight: var(--font-weight-medium);
  }
}

.mentor-details {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
  overflow: hidden;
}

.detail-tabs {
  :deep(.el-tabs__header) {
    margin: 0;
    padding: 0 var(--spacing-xl);
    background: var(--bg-secondary);
    border-bottom: 1px solid var(--border-color-light);
  }
  
  :deep(.el-tabs__nav-wrap) {
    padding: 0;
  }
  
  :deep(.el-tabs__item) {
    font-size: var(--font-size-medium);
    font-weight: var(--font-weight-medium);
    padding: var(--spacing-lg) var(--spacing-xl);
  }
}

.tab-content {
  padding: var(--spacing-xl);
}

.section {
  margin-bottom: var(--spacing-xl);
  
  &:last-child {
    margin-bottom: 0;
  }
}

.section-title {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-md) 0;
}

.section-subtitle {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin: 0 0 var(--spacing-lg) 0;
}

.bio-text,
.experience-text {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  line-height: 1.6;
  margin: 0;
}

.skills-grid {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
}

.skill-tag {
  background: rgba(64, 158, 255, 0.1);
  border-color: var(--primary-color);
  color: var(--primary-color);
  font-weight: var(--font-weight-medium);
}

.service-types,
.achievements {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.service-item,
.achievement-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-small);
  
  .el-icon {
    color: var(--primary-color);
    font-size: 16px;
  }
}

.reviews-header {
  margin-bottom: var(--spacing-xl);
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
}

.reviews-summary {
  display: flex;
  justify-content: center;
}

.overall-rating {
  text-align: center;
}

.rating-score {
  font-size: var(--font-size-h1);
  font-weight: var(--font-weight-bold);
  color: var(--primary-color);
  margin-bottom: var(--spacing-sm);
}

.rating-stars {
  margin-bottom: var(--spacing-sm);
}

.rating-count {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
}

.reviews-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.review-item {
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
}

.review-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--spacing-md);
}

.reviewer-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.reviewer-name {
  font-size: var(--font-size-medium);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
}

.review-date {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
}

.review-content {
  p {
    font-size: var(--font-size-medium);
    color: var(--text-secondary);
    line-height: 1.6;
    margin: 0;
  }
}

.load-more {
  text-align: center;
  margin-top: var(--spacing-xl);
}

.courses-header {
  margin-bottom: var(--spacing-xl);
}

.courses-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: var(--spacing-lg);
}

.course-card {
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  overflow: hidden;
  cursor: pointer;
  transition: all var(--transition-normal);
  
  &:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-light);
  }
}

.course-image {
  position: relative;
  height: 200px;
  overflow: hidden;
  
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  
  .course-overlay {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    opacity: 0;
    transition: opacity var(--transition-normal);
  }
  
  &:hover .course-overlay {
    opacity: 1;
  }
}

.course-info {
  padding: var(--spacing-lg);
}

.course-title {
  font-size: var(--font-size-h5);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-sm) 0;
}

.course-description {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  line-height: 1.5;
  margin: 0 0 var(--spacing-md) 0;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.course-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.course-stats {
  display: flex;
  gap: var(--spacing-md);
}

.stat {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  
  .el-icon {
    font-size: 14px;
  }
}

.course-price {
  .price {
    font-size: var(--font-size-h5);
    font-weight: var(--font-weight-bold);
    color: #ff9900;
  }
}

.empty-courses {
  text-align: center;
  padding: var(--spacing-xxl) 0;
}

.error-container {
  text-align: center;
  padding: var(--spacing-xxl) 0;
}

.booking-content,
.message-content {
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
    
    &.total {
      border-top: 1px solid var(--border-color-light);
      padding-top: var(--spacing-sm);
      font-weight: var(--font-weight-semibold);
      
      .total-price {
        font-size: var(--font-size-h4);
        color: #ff9900;
      }
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
  .mentor-detail-page {
    padding: var(--spacing-lg);
  }
  
  .mentor-header {
    flex-direction: column;
    gap: var(--spacing-lg);
  }
  
  .mentor-profile {
    flex-direction: column;
    text-align: center;
  }
  
  .stats-section {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .action-section {
    min-width: auto;
  }
  
  .courses-grid {
    grid-template-columns: 1fr;
  }
  
  .course-stats {
    flex-direction: column;
    gap: var(--spacing-xs);
  }
}
</style>