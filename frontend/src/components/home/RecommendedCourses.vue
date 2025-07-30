<template>
  <div class="recommended-courses">
    <div class="section-header">
      <h2 class="section-title">推荐课程</h2>
      <div class="header-actions">
        <el-button type="text" @click="loadRecommendedCourses">
          <el-icon><Refresh /></el-icon>
          刷新推荐
        </el-button>
        <el-button type="text" @click="$router.push('/courses')">
          <el-icon><MoreFilled /></el-icon>
          查看更多
        </el-button>
      </div>
    </div>
    
    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="1" animated />
      <el-skeleton :rows="1" animated />
      <el-skeleton :rows="1" animated />
    </div>
    
    <div v-else-if="recommendedCourses.length > 0" class="courses-carousel">
      <div class="course-card-wrapper" v-for="course in recommendedCourses" :key="course.id">
        <div class="course-card" @click="viewCourseDetail(course.id)">
          <div class="course-cover">
            <img :src="course.cover" :alt="course.title" />
            <div class="course-badge" v-if="course.isNew">新课程</div>
            <div class="course-badge hot" v-if="course.isHot">热门</div>
          </div>
          <div class="course-content">
            <h3 class="course-title">{{ course.title }}</h3>
            <p class="course-description">{{ course.description.substring(0, 60) }}...</p>
            <div class="course-meta">
              <div class="mentor-info">
                <el-avatar :size="24" :src="course.mentorAvatar" />
                <span class="mentor-name">{{ course.mentorName }}</span>
              </div>
              <div class="course-stats">
                <span class="students">{{ course.studentCount }} 学员</span>
                <span class="rating">
                  <el-rate v-model="course.rating" disabled size="small" />
                  {{ course.rating.toFixed(1) }}
                </span>
              </div>
            </div>
            <div class="course-footer">
              <div class="course-price">
                <span class="price-value">¥{{ course.price }}</span>
                <span class="price-original" v-if="course.originalPrice">¥{{ course.originalPrice }}</span>
              </div>
              <el-button type="primary" size="small" class="enroll-button" @click.stop="enrollCourse(course)">
                立即报名
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <div v-else class="no-courses">
      <el-empty description="暂无推荐课程，请稍后再试" />
    </div>
    
    <!-- 报名对话框 -->
    <el-dialog
      v-model="showEnrollDialog"
      title="课程报名"
      width="90%"
      max-width="500px"
      :close-on-click-modal="false"
    >
      <div v-if="selectedCourse" class="enroll-content">
        <div class="course-summary">
          <img :src="selectedCourse.cover" :alt="selectedCourse.title" class="course-cover-small" />
          <div class="course-info">
            <h3>{{ selectedCourse.title }}</h3>
            <p>{{ selectedCourse.description }}</p>
            <div class="course-details">
              <span>大师：{{ selectedCourse.mentorName }}</span>
              <span>时长：{{ selectedCourse.duration }}小时</span>
              <span>学员：{{ selectedCourse.studentCount }}人</span>
            </div>
          </div>
        </div>
        
        <div class="price-section">
          <div class="price-info">
            <span class="current-price">¥{{ selectedCourse.price }}</span>
            <span class="original-price" v-if="selectedCourse.originalPrice">¥{{ selectedCourse.originalPrice }}</span>
            <span class="discount" v-if="selectedCourse.originalPrice">
              省¥{{ selectedCourse.originalPrice - selectedCourse.price }}
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
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Refresh, MoreFilled } from '@element-plus/icons-vue'
import { ApiService } from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const recommendedCourses = ref<any[]>([])
const loading = ref(true)
const showEnrollDialog = ref(false)
const selectedCourse = ref<any>(null)
const enrolling = ref(false)

const loadRecommendedCourses = async () => {
  loading.value = true
  try {
    const response = await ApiService.courses.getRecommendedCourses()
    recommendedCourses.value = response.data.slice(0, 4) // 只显示4个推荐课程
  } catch (error) {
    ElMessage.error('加载推荐课程失败')
    console.error('Failed to load recommended courses:', error)
  } finally {
    loading.value = false
  }
}

const viewCourseDetail = (courseId: string) => {
  router.push(`/courses/${courseId}`)
}

const enrollCourse = (course: any) => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录才能报名课程')
    router.push('/auth')
    return
  }
  
  selectedCourse.value = course
  showEnrollDialog.value = true
}

const confirmEnroll = async () => {
  if (!selectedCourse.value || !authStore.user) return
  
  enrolling.value = true
  try {
    const enrollData = {
      courseId: selectedCourse.value.id,
      userId: authStore.user.id,
      price: selectedCourse.value.price
    }
    
    await ApiService.courses.enrollCourse(enrollData)
    
    ElMessage.success('报名成功！')
    showEnrollDialog.value = false
    
    // 跳转到课程详情页
    router.push(`/courses/${selectedCourse.value.id}`)
  } catch (error) {
    ElMessage.error('报名失败，请重试')
  } finally {
    enrolling.value = false
  }
}

onMounted(() => {
  loadRecommendedCourses()
})
</script>

<style scoped lang="scss">
.recommended-courses {
  padding: var(--spacing-xl);
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
  margin-top: var(--spacing-xl);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
}

.section-title {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0;
}

.header-actions {
  .el-button {
    color: var(--primary-color);
    font-size: var(--font-size-small);
    margin-left: var(--spacing-md);
  }
}

.loading-container {
  display: flex;
  gap: var(--spacing-lg);
  overflow-x: auto;
  padding-bottom: var(--spacing-md);
}

.courses-carousel {
  display: flex;
  gap: var(--spacing-lg);
  overflow-x: auto;
  padding-bottom: var(--spacing-md);
  -webkit-overflow-scrolling: touch;
  
  &::-webkit-scrollbar {
    height: 8px;
  }
  
  &::-webkit-scrollbar-thumb {
    background-color: var(--border-color-light);
    border-radius: 4px;
  }
  
  &::-webkit-scrollbar-track {
    background-color: var(--bg-secondary);
  }
}

.course-card-wrapper {
  flex-shrink: 0;
  width: 300px;
}

.course-card {
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  overflow: hidden;
  box-shadow: var(--shadow-light);
  transition: all var(--transition-normal);
  cursor: pointer;
  height: 100%;
  
  &:hover {
    transform: translateY(-5px);
    box-shadow: var(--shadow-medium);
  }
}

.course-cover {
  position: relative;
  height: 160px;
  overflow: hidden;
  
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  
  .course-badge {
    position: absolute;
    top: var(--spacing-sm);
    right: var(--spacing-sm);
    background: var(--primary-color);
    color: white;
    padding: 4px 8px;
    border-radius: var(--border-radius-small);
    font-size: var(--font-size-small);
    font-weight: var(--font-weight-medium);
    
    &.hot {
      background: #f56c6c;
    }
  }
}

.course-content {
  padding: var(--spacing-lg);
  display: flex;
  flex-direction: column;
  height: calc(100% - 160px);
}

.course-title {
  font-size: var(--font-size-h5);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-sm) 0;
  line-height: 1.4;
}

.course-description {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-md);
  line-height: 1.5;
  flex-grow: 1;
}

.course-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-md);
}

.mentor-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  
  .mentor-name {
    font-size: var(--font-size-small);
    color: var(--text-secondary);
  }
}

.course-stats {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: var(--spacing-xs);
  
  .students {
    font-size: var(--font-size-small);
    color: var(--text-tertiary);
  }
  
  .rating {
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
    font-size: var(--font-size-small);
    color: var(--text-secondary);
  }
}

.course-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--border-color-light);
  margin-top: auto;
}

.course-price {
  display: flex;
  align-items: baseline;
  gap: var(--spacing-xs);
  
  .price-value {
    font-size: var(--font-size-h5);
    font-weight: var(--font-weight-bold);
    color: #f56c6c;
  }
  
  .price-original {
    font-size: var(--font-size-small);
    color: var(--text-tertiary);
    text-decoration: line-through;
  }
}

.enroll-button {
  background: linear-gradient(135deg, var(--primary-color), var(--master-color));
  border: none;
  color: white;
  font-weight: var(--font-weight-medium);
}

.no-courses {
  text-align: center;
  padding: var(--spacing-xl);
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

// 响应式设计
@media (max-width: 768px) {
  .recommended-courses {
    padding: var(--spacing-lg);
  }
  
  .courses-carousel {
    padding-bottom: var(--spacing-lg);
  }
  
  .course-card-wrapper {
    width: 280px;
  }
  
  .course-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-sm);
  }
  
  .course-footer {
    flex-direction: column;
    gap: var(--spacing-sm);
    align-items: stretch;
  }
}
</style>