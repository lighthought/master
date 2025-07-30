<template>
  <div class="course-detail-page">
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="10" animated />
    </div>
    
    <div v-else-if="course" class="course-content">
      <!-- 课程头部信息 -->
      <div class="course-header">
        <div class="course-cover-section">
          <img :src="course.cover" :alt="course.title" class="course-cover" />
          <div class="course-badges">
            <el-tag v-if="course.isNew" type="success" effect="dark">新课程</el-tag>
            <el-tag v-if="course.isHot" type="danger" effect="dark">热门</el-tag>
            <el-tag v-if="course.isRecommended" type="warning" effect="dark">推荐</el-tag>
          </div>
        </div>
        
        <div class="course-info">
          <h1 class="course-title">{{ course.title }}</h1>
          <p class="course-description">{{ course.description }}</p>
          
          <div class="course-meta">
            <div class="mentor-info">
              <el-avatar :size="50" :src="course.mentorAvatar" />
              <div class="mentor-details">
                <h3 class="mentor-name">{{ course.mentorName }}</h3>
                <p class="mentor-title">资深{{ getCategoryText(course.category) }}专家</p>
              </div>
            </div>
            
            <div class="course-stats">
              <div class="stat-item">
                <span class="stat-label">学员数</span>
                <span class="stat-value">{{ course.studentCount }}</span>
              </div>
              <div class="stat-item">
                <span class="stat-label">课程时长</span>
                <span class="stat-value">{{ course.duration }}小时</span>
              </div>
              <div class="stat-item">
                <span class="stat-label">难度等级</span>
                <span class="stat-value">{{ getLevelText(course.level) }}</span>
              </div>
              <div class="stat-item">
                <span class="stat-label">评分</span>
                <span class="stat-value">
                  <el-rate v-model="course.rating" disabled />
                  {{ course.rating.toFixed(1) }}
                </span>
              </div>
            </div>
          </div>
          
          <div class="course-tags">
            <el-tag v-for="tag in course.tags" :key="tag" effect="light" class="tag">
              {{ tag }}
            </el-tag>
          </div>
          
          <!-- 学习进度指示器 -->
          <div class="learning-progress" v-if="course.enrollmentStatus">
            <div class="progress-header">
              <h4>学习进度</h4>
              <span class="progress-percentage">{{ course.progress || 0 }}%</span>
            </div>
            <el-progress 
              :percentage="course.progress || 0" 
              :stroke-width="8"
              :show-text="false"
              color="var(--primary-color)"
            />
            <div class="progress-stats">
              <span>已完成 {{ course.completedLessons || 0 }}/{{ course.totalLessons || 0 }} 课时</span>
              <span>学习时长 {{ course.studyTime || 0 }} 小时</span>
            </div>
          </div>
          
          <div class="course-price-section">
            <div class="price-info">
              <span class="current-price">¥{{ course.price }}</span>
              <span class="original-price" v-if="course.originalPrice">¥{{ course.originalPrice }}</span>
              <span class="discount" v-if="course.originalPrice">
                省¥{{ course.originalPrice - course.price }}
              </span>
            </div>
            
            <div class="action-buttons">
              <el-button 
                type="primary" 
                size="large" 
                @click="enrollCourse"
                :loading="enrolling"
                class="enroll-button"
              >
                <el-icon><ShoppingCart /></el-icon>
                立即报名
              </el-button>
              
              <el-button 
                type="default" 
                size="large" 
                @click="addToWishlist"
                class="wishlist-button"
              >
                <el-icon><Star /></el-icon>
                收藏课程
              </el-button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 课程详情标签页 -->
      <div class="course-tabs">
        <el-tabs v-model="activeTab" type="border-card">
          <el-tab-pane label="课程大纲" name="outline">
            <div class="outline-content">
              <div v-for="(section, index) in course.outline" :key="index" class="outline-section">
                <div class="section-header" @click="toggleSection(index)">
                  <h3 class="section-title">
                    <span class="section-number">{{ index + 1 }}</span>
                    {{ section.title }}
                  </h3>
                  <div class="section-meta">
                    <span class="lesson-count">{{ section.lessons.length }} 课时</span>
                    <span class="section-duration">{{ getSectionDuration(section) }}</span>
                    <el-icon class="toggle-icon" :class="{ 'is-expanded': expandedSections.includes(index) }">
                      <ArrowDown />
                    </el-icon>
                  </div>
                </div>
                <div class="lessons-list" v-show="expandedSections.includes(index)">
                  <div v-for="(lesson, lessonIndex) in section.lessons" :key="lessonIndex" class="lesson-item">
                    <div class="lesson-info">
                      <el-icon class="lesson-icon">
                        <VideoPlay v-if="lesson.type === 'video'" />
                        <Document v-else />
                      </el-icon>
                      <span class="lesson-title">{{ lesson.title }}</span>
                      <el-tag v-if="lesson.isFree" type="success" size="small" effect="light">免费</el-tag>
                    </div>
                    <div class="lesson-meta">
                      <span class="lesson-duration">{{ lesson.duration }}</span>
                      <el-button v-if="lesson.preview" type="text" size="small" @click="previewLesson(lesson)">
                        <el-icon><VideoPlay /></el-icon>
                        预览
                      </el-button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </el-tab-pane>
          
          <el-tab-pane label="课程介绍" name="introduction">
            <div class="introduction-content">
              <div class="intro-section">
                <h3>课程特色</h3>
                <ul class="feature-list">
                  <li>系统化的学习路径，从基础到进阶</li>
                  <li>实战项目驱动，学以致用</li>
                  <li>大师一对一答疑服务</li>
                  <li>终身学习，课程内容持续更新</li>
                  <li>结业证书，提升职业竞争力</li>
                </ul>
              </div>
              
              <div class="intro-section">
                <h3>学习目标</h3>
                <p>通过本课程的学习，你将掌握：</p>
                <ul class="goal-list">
                  <li>核心概念和基础理论</li>
                  <li>实际项目开发技能</li>
                  <li>最佳实践和设计模式</li>
                  <li>性能优化和调试技巧</li>
                  <li>团队协作和项目管理</li>
                </ul>
              </div>
              
              <div class="intro-section">
                <h3>适合人群</h3>
                <div class="target-audience">
                  <el-tag v-for="audience in course.targetAudience" :key="audience" effect="light" class="audience-tag">
                    {{ audience }}
                  </el-tag>
                </div>
              </div>
              
              <div class="intro-section">
                <h3>学习要求</h3>
                <ul class="requirement-list">
                  <li v-for="requirement in course.requirements" :key="requirement">{{ requirement }}</li>
                </ul>
              </div>
            </div>
          </el-tab-pane>
          
          <el-tab-pane label="学员评价" name="reviews">
            <div class="reviews-content">
              <div class="reviews-header">
                <h3>学员评价 ({{ course.reviews.length }})</h3>
                <div class="overall-rating">
                  <span class="rating-score">{{ course.rating.toFixed(1) }}</span>
                  <el-rate v-model="course.rating" disabled />
                  <span class="rating-count">{{ course.studentCount }} 名学员</span>
                </div>
              </div>
              
              <!-- 评分统计图表 -->
              <div class="rating-stats">
                <div class="rating-distribution">
                  <div v-for="i in 5" :key="i" class="rating-bar">
                    <span class="star-label">{{ 6 - i }}星</span>
                    <div class="bar-container">
                      <div class="bar-fill" :style="{ width: getRatingPercentage(6 - i) + '%' }"></div>
                    </div>
                    <span class="rating-count">{{ getRatingCount(6 - i) }}</span>
                  </div>
                </div>
              </div>
              
              <div class="reviews-list">
                <div v-for="review in course.reviews" :key="review.id" class="review-item">
                  <div class="review-header">
                    <div class="reviewer-info">
                      <el-avatar :size="40" :src="review.userAvatar" />
                      <div class="reviewer-details">
                        <span class="reviewer-name">{{ review.userName }}</span>
                        <el-rate v-model="review.rating" disabled size="small" />
                      </div>
                    </div>
                    <span class="review-date">{{ formatDate(review.createdAt) }}</span>
                  </div>
                  <p class="review-content">{{ review.content }}</p>
                </div>
              </div>
              
              <div class="write-review">
                <el-button type="primary" @click="showReviewDialog = true">
                  写评价
                </el-button>
              </div>
            </div>
          </el-tab-pane>
          
          <el-tab-pane label="相关课程" name="related">
            <div class="related-courses">
              <div v-for="relatedCourse in relatedCourses" :key="relatedCourse.id" class="related-course-card" @click="viewCourse(relatedCourse.id)">
                <img :src="relatedCourse.cover" :alt="relatedCourse.title" class="related-course-cover" />
                <div class="related-course-info">
                  <h4>{{ relatedCourse.title }}</h4>
                  <p>{{ relatedCourse.description.substring(0, 50) }}...</p>
                  <div class="related-course-meta">
                    <span class="mentor">{{ relatedCourse.mentorName }}</span>
                    <span class="price">¥{{ relatedCourse.price }}</span>
                  </div>
                </div>
              </div>
            </div>
          </el-tab-pane>
        </el-tabs>
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
    
    <!-- 报名对话框 -->
    <el-dialog
      v-model="showEnrollDialog"
      title="课程报名"
      width="90%"
      max-width="500px"
      :close-on-click-modal="false"
    >
      <div v-if="course" class="enroll-content">
        <div class="course-summary">
          <img :src="course.cover" :alt="course.title" class="course-cover-small" />
          <div class="course-info">
            <h3>{{ course.title }}</h3>
            <p>{{ course.description }}</p>
            <div class="course-details">
              <span>大师：{{ course.mentorName }}</span>
              <span>时长：{{ course.duration }}小时</span>
              <span>学员：{{ course.studentCount }}人</span>
            </div>
          </div>
        </div>
        
        <div class="price-section">
          <div class="price-info">
            <span class="current-price">¥{{ course.price }}</span>
            <span class="original-price" v-if="course.originalPrice">¥{{ course.originalPrice }}</span>
            <span class="discount" v-if="course.originalPrice">
              省¥{{ course.originalPrice - course.price }}
            </span>
          </div>
        </div>
        
        <div class="enroll-benefits">
          <h4>报名后你将获得：</h4>
          <ul>
            <li>完整的课程视频和资料</li>
            <li>大师一对一答疑服务</li>
            <li>学习进度跟踪</li>
            <li>结业证书</li>
          </ul>
        </div>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showEnrollDialog = false">取消</el-button>
          <el-button type="primary" @click="confirmEnroll" :loading="enrolling">
            确认报名
          </el-button>
        </div>
      </template>
    </el-dialog>
    
    <!-- 评价对话框 -->
    <el-dialog
      v-model="showReviewDialog"
      title="写评价"
      width="90%"
      max-width="500px"
    >
      <el-form :model="reviewForm" label-width="80px">
        <el-form-item label="评分">
          <el-rate v-model="reviewForm.rating" />
        </el-form-item>
        <el-form-item label="评价内容">
          <el-input
            v-model="reviewForm.content"
            type="textarea"
            :rows="4"
            placeholder="请分享你的学习体验..."
          />
        </el-form-item>
      </el-form>
      
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
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ShoppingCart, Star, VideoPlay, Document, ArrowDown } from '@element-plus/icons-vue'
import { ApiService } from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const course = ref<any>(null)
const relatedCourses = ref<any[]>([])
const loading = ref(true)
const enrolling = ref(false)
const submittingReview = ref(false)
const activeTab = ref('outline')
const showEnrollDialog = ref(false)
const showReviewDialog = ref(false)
const expandedSections = ref<number[]>([0]) // 默认展开第一个章节

const reviewForm = ref({
  rating: 5,
  content: ''
})

// 加载课程详情
const loadCourseDetail = async () => {
  const courseId = route.params.id as string
  loading.value = true
  
  try {
    const response = await ApiService.courses.getCourseDetail(courseId)
    course.value = response.data
    
    // 加载相关课程
    await loadRelatedCourses()
  } catch (error) {
    console.error('加载课程详情失败:', error)
    ElMessage.error('加载课程详情失败')
  } finally {
    loading.value = false
  }
}

// 加载相关课程
const loadRelatedCourses = async () => {
  try {
    const params = {
      category: course.value.category,
      pageSize: 4
    }
    const response = await ApiService.courses.searchCourses(params)
    relatedCourses.value = response.data.courses.filter((c: any) => c.id !== course.value.id)
  } catch (error) {
    console.error('加载相关课程失败:', error)
  }
}

// 获取分类文本
const getCategoryText = (category: string) => {
  const categories: Record<string, string> = {
    frontend: '前端开发',
    backend: '后端开发',
    mobile: '移动开发',
    data: '数据科学',
    ai: '人工智能',
    design: '产品设计',
    management: '项目管理'
  }
  return categories[category] || category
}

// 获取难度等级文本
const getLevelText = (level: string) => {
  const levels: Record<string, string> = {
    beginner: '初级',
    intermediate: '中级',
    advanced: '高级'
  }
  return levels[level] || level
}

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('zh-CN')
}

// 报名课程
const enrollCourse = () => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录才能报名课程')
    router.push('/auth')
    return
  }
  
  showEnrollDialog.value = true
}

// 确认报名
const confirmEnroll = async () => {
  if (!course.value || !authStore.user) return
  
  enrolling.value = true
  try {
    const enrollData = {
      courseId: course.value.id,
      userId: authStore.user.id,
      price: course.value.price
    }
    
    await ApiService.courses.enrollCourse(enrollData)
    
    ElMessage.success('报名成功！')
    showEnrollDialog.value = false
    
    // 跳转到学习页面
    router.push(`/learning/${course.value.id}`)
  } catch (error) {
    ElMessage.error('报名失败，请重试')
  } finally {
    enrolling.value = false
  }
}

// 添加到收藏
const addToWishlist = () => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录才能收藏课程')
    router.push('/auth')
    return
  }
  
  ElMessage.success('课程已添加到收藏')
}

// 查看课程
const viewCourse = (courseId: string) => {
  router.push(`/courses/${courseId}`)
}

// 切换章节展开状态
const toggleSection = (index: number) => {
  const sectionIndex = expandedSections.value.indexOf(index)
  if (sectionIndex > -1) {
    expandedSections.value.splice(sectionIndex, 1)
  } else {
    expandedSections.value.push(index)
  }
}

// 计算章节总时长
const getSectionDuration = (section: any) => {
  const totalMinutes = section.lessons.reduce((total: number, lesson: any) => {
    const duration = lesson.duration
    const minutes = parseInt(duration.match(/(\d+)/)?.[1] || '0')
    return total + minutes
  }, 0)
  
  if (totalMinutes >= 60) {
    const hours = Math.floor(totalMinutes / 60)
    const minutes = totalMinutes % 60
    return `${hours}小时${minutes > 0 ? minutes + '分钟' : ''}`
  }
  return `${totalMinutes}分钟`
}

// 预览课程
const previewLesson = (lesson: any) => {
  ElMessage.info(`预览课程：${lesson.title}`)
  // 实际应用中会打开视频预览窗口
}

// 获取评分百分比
const getRatingPercentage = (rating: number) => {
  if (!course.value?.reviews) return 0
  const count = course.value.reviews.filter((r: any) => r.rating === rating).length
  return Math.round((count / course.value.reviews.length) * 100)
}

// 获取评分数量
const getRatingCount = (rating: number) => {
  if (!course.value?.reviews) return 0
  return course.value.reviews.filter((r: any) => r.rating === rating).length
}

// 提交评价
const submitReview = async () => {
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
    
    // 重新加载课程详情
    await loadCourseDetail()
  } catch (error) {
    ElMessage.error('评价提交失败，请重试')
  } finally {
    submittingReview.value = false
  }
}

onMounted(() => {
  loadCourseDetail()
})
</script>

<style scoped lang="scss">
.course-detail-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: var(--spacing-xl);
}

.loading-container {
  padding: var(--spacing-xl);
}

.course-content {
  .course-header {
    display: grid;
    grid-template-columns: 400px 1fr;
    gap: var(--spacing-xl);
    margin-bottom: var(--spacing-xxl);
    
    @media (max-width: 768px) {
      grid-template-columns: 1fr;
      gap: var(--spacing-lg);
    }
  }
}

.course-cover-section {
  position: relative;
  
  .course-cover {
    width: 100%;
    height: 300px;
    object-fit: cover;
    border-radius: var(--border-radius-medium);
    box-shadow: var(--shadow-medium);
  }
  
  .course-badges {
    position: absolute;
    top: var(--spacing-md);
    right: var(--spacing-md);
    display: flex;
    flex-direction: column;
    gap: var(--spacing-xs);
  }
}

.course-info {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.course-title {
  font-size: var(--font-size-h2);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin: 0;
  line-height: 1.3;
}

.course-description {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  line-height: 1.6;
  margin: 0;
}

.course-meta {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: var(--spacing-lg);
  
  @media (max-width: 768px) {
    flex-direction: column;
    gap: var(--spacing-md);
  }
}

.mentor-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  
  .mentor-details {
    .mentor-name {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-xs) 0;
    }
    
    .mentor-title {
      font-size: var(--font-size-small);
      color: var(--text-secondary);
      margin: 0;
    }
  }
}

.course-stats {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-md);
  
  .stat-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    
    .stat-label {
      font-size: var(--font-size-small);
      color: var(--text-secondary);
      margin-bottom: var(--spacing-xs);
    }
    
    .stat-value {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
    }
  }
}

.course-tags {
  display: flex;
  flex-wrap: wrap;
  gap: var(--spacing-sm);
  
  .tag {
    font-size: var(--font-size-small);
  }
}

.learning-progress {
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  border: 1px solid var(--border-color-light);
  
  .progress-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: var(--spacing-md);
    
    h4 {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0;
    }
    
    .progress-percentage {
      font-size: var(--font-size-h4);
      font-weight: var(--font-weight-bold);
      color: var(--primary-color);
    }
  }
  
  .progress-stats {
    display: flex;
    justify-content: space-between;
    margin-top: var(--spacing-sm);
    font-size: var(--font-size-small);
    color: var(--text-secondary);
  }
}

.course-price-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  
  @media (max-width: 768px) {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: stretch;
  }
}

.price-info {
  display: flex;
  align-items: baseline;
  gap: var(--spacing-sm);
  
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

.action-buttons {
  display: flex;
  gap: var(--spacing-md);
  
  .enroll-button {
    background: linear-gradient(135deg, var(--primary-color), var(--master-color));
    border: none;
    color: white;
    font-weight: var(--font-weight-medium);
  }
  
  .wishlist-button {
    border: 1px solid var(--border-color-light);
    color: var(--text-secondary);
  }
}

.course-tabs {
  .el-tabs {
    :deep(.el-tabs__header) {
      margin-bottom: var(--spacing-lg);
    }
    
    :deep(.el-tabs__content) {
      padding: var(--spacing-lg);
      background: var(--bg-card);
      border-radius: var(--border-radius-medium);
    }
  }
}

// 课程大纲样式
.outline-content {
  .outline-section {
    margin-bottom: var(--spacing-lg);
    border: 1px solid var(--border-color-light);
    border-radius: var(--border-radius-medium);
    overflow: hidden;
    
    &:last-child {
      margin-bottom: 0;
    }
  }
  
  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: var(--spacing-lg);
    background: var(--bg-secondary);
    cursor: pointer;
    transition: all var(--transition-normal);
    
    &:hover {
      background: var(--bg-tertiary);
    }
    
    .section-title {
      display: flex;
      align-items: center;
      gap: var(--spacing-md);
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0;
      
      .section-number {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 30px;
        height: 30px;
        background: var(--primary-color);
        color: white;
        border-radius: 50%;
        font-size: var(--font-size-small);
        font-weight: var(--font-weight-bold);
      }
    }
    
    .section-meta {
      display: flex;
      align-items: center;
      gap: var(--spacing-md);
      font-size: var(--font-size-small);
      color: var(--text-secondary);
      
      .lesson-count {
        color: var(--primary-color);
        font-weight: var(--font-weight-medium);
      }
      
      .toggle-icon {
        transition: transform var(--transition-normal);
        
        &.is-expanded {
          transform: rotate(180deg);
        }
      }
    }
  }
  
  .lessons-list {
    display: flex;
    flex-direction: column;
    gap: 1px;
    background: var(--border-color-light);
  }
  
  .lesson-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: var(--spacing-md);
    background: var(--bg-card);
    transition: all var(--transition-normal);
    
    &:hover {
      background: var(--bg-secondary);
    }
    
    .lesson-info {
      display: flex;
      align-items: center;
      gap: var(--spacing-sm);
      flex: 1;
      
      .lesson-icon {
        color: var(--primary-color);
      }
      
      .lesson-title {
        font-size: var(--font-size-medium);
        color: var(--text-primary);
        flex: 1;
      }
    }
    
    .lesson-meta {
      display: flex;
      align-items: center;
      gap: var(--spacing-md);
      
      .lesson-duration {
        font-size: var(--font-size-small);
        color: var(--text-secondary);
        min-width: 60px;
        text-align: right;
      }
    }
  }
}

// 课程介绍样式
.introduction-content {
  .intro-section {
    margin-bottom: var(--spacing-xl);
    
    &:last-child {
      margin-bottom: 0;
    }
    
    h3 {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-md) 0;
    }
    
    p {
      font-size: var(--font-size-medium);
      color: var(--text-secondary);
      line-height: 1.6;
      margin: 0 0 var(--spacing-md) 0;
    }
  }
  
  .feature-list,
  .goal-list,
  .requirement-list {
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
        content: '✓';
        color: #67c23a;
        font-weight: var(--font-weight-bold);
        position: absolute;
        left: 0;
      }
      
      &:last-child {
        margin-bottom: 0;
      }
    }
  }
  
  .target-audience {
    display: flex;
    flex-wrap: wrap;
    gap: var(--spacing-sm);
    
    .audience-tag {
      font-size: var(--font-size-small);
    }
  }
}

// 评价样式
.reviews-content {
  .reviews-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: var(--spacing-lg);
    padding-bottom: var(--spacing-md);
    border-bottom: 1px solid var(--border-color-light);
    
    h3 {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0;
    }
    
    .overall-rating {
      display: flex;
      align-items: center;
      gap: var(--spacing-sm);
      
      .rating-score {
        font-size: var(--font-size-h4);
        font-weight: var(--font-weight-bold);
        color: var(--text-primary);
      }
      
      .rating-count {
        font-size: var(--font-size-small);
        color: var(--text-secondary);
      }
    }
  }
  
  .rating-stats {
    margin-bottom: var(--spacing-xl);
    padding: var(--spacing-lg);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
    
    .rating-distribution {
      display: flex;
      flex-direction: column;
      gap: var(--spacing-sm);
      
      .rating-bar {
        display: flex;
        align-items: center;
        gap: var(--spacing-md);
        
        .star-label {
          font-size: var(--font-size-small);
          color: var(--text-secondary);
          min-width: 30px;
        }
        
        .bar-container {
          flex: 1;
          height: 8px;
          background: var(--border-color-light);
          border-radius: 4px;
          overflow: hidden;
          
          .bar-fill {
            height: 100%;
            background: linear-gradient(90deg, #ffd700, #ffed4e);
            border-radius: 4px;
            transition: width var(--transition-normal);
          }
        }
        
        .rating-count {
          font-size: var(--font-size-small);
          color: var(--text-secondary);
          min-width: 30px;
          text-align: right;
        }
      }
    }
  }
  
  .reviews-list {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-lg);
    margin-bottom: var(--spacing-lg);
  }
  
  .review-item {
    padding: var(--spacing-lg);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
    
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
      
      .reviewer-details {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-xs);
        
        .reviewer-name {
          font-size: var(--font-size-medium);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
        }
      }
    }
    
    .review-date {
      font-size: var(--font-size-small);
      color: var(--text-tertiary);
    }
    
    .review-content {
      font-size: var(--font-size-medium);
      color: var(--text-secondary);
      line-height: 1.6;
      margin: 0;
    }
  }
  
  .write-review {
    text-align: center;
    padding: var(--spacing-lg);
    border: 2px dashed var(--border-color-light);
    border-radius: var(--border-radius-medium);
  }
}

// 相关课程样式
.related-courses {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: var(--spacing-lg);
  
  .related-course-card {
    display: flex;
    gap: var(--spacing-md);
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
    cursor: pointer;
    transition: all var(--transition-normal);
    
    &:hover {
      transform: translateY(-2px);
      box-shadow: var(--shadow-light);
    }
    
    .related-course-cover {
      width: 80px;
      height: 60px;
      object-fit: cover;
      border-radius: var(--border-radius-small);
    }
    
    .related-course-info {
      flex: 1;
      
      h4 {
        font-size: var(--font-size-medium);
        font-weight: var(--font-weight-semibold);
        color: var(--text-primary);
        margin: 0 0 var(--spacing-xs) 0;
        line-height: 1.4;
      }
      
      p {
        font-size: var(--font-size-small);
        color: var(--text-secondary);
        margin: 0 0 var(--spacing-sm) 0;
        line-height: 1.4;
      }
      
      .related-course-meta {
        display: flex;
        justify-content: space-between;
        align-items: center;
        
        .mentor {
          font-size: var(--font-size-small);
          color: var(--text-tertiary);
        }
        
        .price {
          font-size: var(--font-size-small);
          font-weight: var(--font-weight-medium);
          color: #f56c6c;
        }
      }
    }
  }
}

// 报名对话框样式
.enroll-content {
  .course-summary {
    display: flex;
    gap: var(--spacing-md);
    margin-bottom: var(--spacing-lg);
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
  }
  
  .course-cover-small {
    width: 80px;
    height: 60px;
    object-fit: cover;
    border-radius: var(--border-radius-small);
  }
  
  .course-info {
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
    
    .course-details {
      display: flex;
      flex-direction: column;
      gap: var(--spacing-xs);
      font-size: var(--font-size-small);
      color: var(--text-tertiary);
    }
  }
  
  .price-section {
    margin-bottom: var(--spacing-lg);
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
  }
  
  .price-info {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
    
    .current-price {
      font-size: var(--font-size-h3);
      font-weight: var(--font-weight-bold);
      color: #f56c6c;
    }
    
    .original-price {
      font-size: var(--font-size-medium);
      color: var(--text-tertiary);
      text-decoration: line-through;
    }
    
    .discount {
      font-size: var(--font-size-small);
      color: #67c23a;
      background: rgba(103, 194, 58, 0.1);
      padding: 2px 6px;
      border-radius: var(--border-radius-small);
    }
  }
  
  .enroll-benefits {
    h4 {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-md) 0;
    }
    
    ul {
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
          content: '✓';
          color: #67c23a;
          font-weight: var(--font-weight-bold);
          position: absolute;
          left: 0;
        }
        
        &:last-child {
          margin-bottom: 0;
        }
      }
    }
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
}

.empty-state {
  text-align: center;
  padding: var(--spacing-xxl) 0;
}

// 响应式设计
@media (max-width: 768px) {
  .course-detail-page {
    padding: var(--spacing-lg);
  }
  
  .course-header {
    grid-template-columns: 1fr;
  }
  
  .course-meta {
    flex-direction: column;
  }
  
  .course-price-section {
    flex-direction: column;
    align-items: stretch;
  }
  
  .action-buttons {
    flex-direction: column;
  }
  
  .course-stats {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .related-courses {
    grid-template-columns: 1fr;
  }
}
</style>