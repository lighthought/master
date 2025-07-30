<template>
  <div class="learning-records-page">
    <div class="container">
      <!-- 页面头部 -->
      <div class="page-header">
        <h1 class="page-title">学习记录</h1>
        <p class="page-description">查看您的学习进度和成果</p>
      </div>

      <!-- 加载状态 -->
      <div v-if="loading" class="loading-container">
        <el-skeleton :rows="10" animated />
      </div>

      <div v-else class="learning-content">
        <!-- 学习统计概览 -->
        <div class="stats-overview">
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><VideoPlay /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-number">{{ learningStats.totalCourses }}</div>
              <div class="stat-label">已报名课程</div>
            </div>
          </div>
          
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><Clock /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-number">{{ learningStats.totalStudyTime }}</div>
              <div class="stat-label">总学习时长(小时)</div>
            </div>
          </div>
          
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><Trophy /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-number">{{ learningStats.completedCourses }}</div>
              <div class="stat-label">已完成课程</div>
            </div>
          </div>
          
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><Star /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-number">{{ learningStats.averageRating }}</div>
              <div class="stat-label">平均评分</div>
            </div>
          </div>
        </div>

        <!-- 筛选和搜索 -->
        <div class="filter-section">
          <div class="search-box">
            <el-input
              v-model="searchKeyword"
              placeholder="搜索课程名称或导师..."
              clearable
              @input="handleSearch"
            >
              <template #prefix>
                <el-icon><Search /></el-icon>
              </template>
            </el-input>
          </div>
          <div class="filter-options">
            <el-select v-model="filterStatus" placeholder="学习状态" clearable @change="loadRecords">
              <el-option label="全部" value="" />
              <el-option label="学习中" value="learning" />
              <el-option label="已完成" value="completed" />
              <el-option label="已暂停" value="paused" />
            </el-select>
            <el-select v-model="sortBy" placeholder="排序方式" @change="loadRecords">
              <el-option label="最近学习" value="recent" />
              <el-option label="学习进度" value="progress" />
              <el-option label="报名时间" value="enrollTime" />
            </el-select>
          </div>
        </div>

        <!-- 学习记录列表 -->
        <div class="records-container">
          <div v-if="loadingRecords" class="loading-container">
            <el-skeleton :rows="3" animated />
            <el-skeleton :rows="3" animated />
            <el-skeleton :rows="3" animated />
          </div>
          
          <div v-else-if="learningRecords.length > 0" class="records-list">
            <div 
              v-for="record in learningRecords" 
              :key="record.id"
              class="record-card"
            >
              <div class="course-info">
                <div class="course-cover">
                  <img :src="record.courseCover" :alt="record.courseTitle" />
                  <div class="course-status">
                    <el-tag 
                      :type="getStatusType(record.status)"
                      size="small"
                    >
                      {{ getStatusText(record.status) }}
                    </el-tag>
                  </div>
                </div>
                
                <div class="course-details">
                  <h3 class="course-title">{{ record.courseTitle }}</h3>
                  <p class="course-description">{{ record.courseDescription }}</p>
                  
                  <div class="mentor-info">
                    <el-avatar :size="24" :src="record.mentorAvatar" />
                    <span class="mentor-name">{{ record.mentorName }}</span>
                  </div>
                  
                  <div class="progress-info">
                    <div class="progress-header">
                      <span class="progress-text">学习进度</span>
                      <span class="progress-percentage">{{ record.progress }}%</span>
                    </div>
                    <el-progress 
                      :percentage="record.progress" 
                      :stroke-width="6"
                      :show-text="false"
                      :color="getProgressColor(record.progress)"
                    />
                    <div class="progress-details">
                      <span>已完成 {{ record.completedLessons }}/{{ record.totalLessons }} 课时</span>
                      <span>学习时长 {{ record.studyTime }} 小时</span>
                    </div>
                  </div>
                  
                  <div class="record-meta">
                    <div class="meta-item">
                      <el-icon><Calendar /></el-icon>
                      <span>报名时间：{{ formatDate(record.enrollTime) }}</span>
                    </div>
                    <div class="meta-item">
                      <el-icon><Clock /></el-icon>
                      <span>最近学习：{{ formatTime(record.lastStudyTime) }}</span>
                    </div>
                    <div v-if="record.score" class="meta-item">
                      <el-icon><Star /></el-icon>
                      <span>课程评分：{{ record.score }}/100</span>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="record-actions">
                <el-button 
                  type="primary" 
                  @click="continueLearning(record)"
                  :disabled="record.status === 'completed'"
                >
                  <el-icon><VideoPlay /></el-icon>
                  {{ record.status === 'completed' ? '查看课程' : '继续学习' }}
                </el-button>
                
                <el-button 
                  type="default" 
                  @click="viewCourseDetail(record)"
                >
                  <el-icon><InfoFilled /></el-icon>
                  课程详情
                </el-button>
                
                <el-button 
                  type="text" 
                  @click="viewLearningDetail(record)"
                >
                  <el-icon><Document /></el-icon>
                  学习详情
                </el-button>
              </div>
            </div>
            
            <!-- 加载更多 -->
            <div v-if="hasMore" class="load-more">
              <el-button
                type="text"
                @click="loadMore"
                :loading="loadingMore"
              >
                加载更多
              </el-button>
            </div>
          </div>
          
          <div v-else class="empty-state">
            <el-empty description="暂无学习记录">
              <el-button type="primary" @click="$router.push('/courses')">
                去报名课程
              </el-button>
            </el-empty>
          </div>
        </div>
      </div>
    </div>

    <!-- 学习详情对话框 -->
    <el-dialog
      v-model="showLearningDetailDialog"
      title="学习详情"
      width="90%"
      max-width="800px"
      :close-on-click-modal="false"
    >
      <div v-if="selectedRecord" class="learning-detail">
        <div class="detail-header">
          <h3>{{ selectedRecord.courseTitle }}</h3>
          <div class="detail-stats">
            <div class="stat-item">
              <span class="stat-label">学习进度</span>
              <span class="stat-value">{{ selectedRecord.progress }}%</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">学习时长</span>
              <span class="stat-value">{{ selectedRecord.studyTime }} 小时</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">完成课时</span>
              <span class="stat-value">{{ selectedRecord.completedLessons }}/{{ selectedRecord.totalLessons }}</span>
            </div>
          </div>
        </div>
        
        <div class="detail-content">
          <h4>学习记录</h4>
          <div class="learning-timeline">
            <div 
              v-for="log in selectedRecord.learningLogs" 
              :key="log.id"
              class="timeline-item"
            >
              <div class="timeline-marker"></div>
              <div class="timeline-content">
                <div class="log-header">
                  <span class="log-action">{{ log.action }}</span>
                  <span class="log-time">{{ formatTime(log.createdAt) }}</span>
                </div>
                <div class="log-details">
                  <span>{{ log.details }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import {
  VideoPlay, Clock, Trophy, Star, Search, Calendar, InfoFilled, Document
} from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'

// Router
const router = useRouter()

// Store
const authStore = useAuthStore()

// 状态
const loading = ref(false)
const loadingRecords = ref(false)
const loadingMore = ref(false)
const showLearningDetailDialog = ref(false)

// 数据
const learningStats = ref({
  totalCourses: 0,
  totalStudyTime: 0,
  completedCourses: 0,
  averageRating: 0
})
const learningRecords = ref<any[]>([])
const searchKeyword = ref('')
const filterStatus = ref('')
const sortBy = ref('recent')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const selectedRecord = ref<any>(null)

// 计算属性
const hasMore = computed(() => {
  return learningRecords.value.length < total.value
})

// 方法
const loadStats = async () => {
  try {
    const result = await ApiService.userStats.getLearningStats(authStore.user!.id)
    learningStats.value = result.data
  } catch (error) {
    ElMessage.error('加载学习统计失败')
  }
}

const loadRecords = async (reset = true) => {
  if (reset) {
    currentPage.value = 1
    learningRecords.value = []
  }
  
  loadingRecords.value = true
  try {
    const params = {
      page: currentPage.value,
      pageSize: pageSize.value,
      keyword: searchKeyword.value,
      status: filterStatus.value,
      sortBy: sortBy.value
    }
    
    const result = await ApiService.learningRecords.getUserLearningRecords(authStore.user!.id, params)
    
    if (reset) {
      learningRecords.value = result.data.records
    } else {
      learningRecords.value.push(...result.data.records)
    }
    
    total.value = result.data.total
  } catch (error) {
    ElMessage.error('加载学习记录失败')
  } finally {
    loadingRecords.value = false
  }
}

const loadMore = async () => {
  currentPage.value++
  await loadRecords(false)
}

const handleSearch = () => {
  loadRecords()
}

const continueLearning = (record: any) => {
  router.push(`/learning/${record.courseId}`)
}

const viewCourseDetail = (record: any) => {
  router.push(`/courses/${record.courseId}`)
}

const viewLearningDetail = (record: any) => {
  selectedRecord.value = record
  showLearningDetailDialog.value = true
}

const getStatusType = (status: string) => {
  const statusMap: Record<string, 'primary' | 'success' | 'warning' | 'info'> = {
    learning: 'primary',
    completed: 'success',
    paused: 'warning'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status: string) => {
  const statusMap: Record<string, string> = {
    learning: '学习中',
    completed: '已完成',
    paused: '已暂停'
  }
  return statusMap[status] || '未知'
}

const getProgressColor = (progress: number) => {
  if (progress >= 80) return '#67C23A'
  if (progress >= 50) return '#E6A23C'
  return '#F56C6C'
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString()
}

const formatTime = (time: string) => {
  const now = new Date()
  const targetTime = new Date(time)
  const diff = now.getTime() - targetTime.getTime()
  
  const minutes = Math.floor(diff / (1000 * 60))
  const hours = Math.floor(diff / (1000 * 60 * 60))
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  
  if (minutes < 60) {
    return `${minutes}分钟前`
  } else if (hours < 24) {
    return `${hours}小时前`
  } else if (days < 7) {
    return `${days}天前`
  } else {
    return targetTime.toLocaleDateString()
  }
}

onMounted(async () => {
  loading.value = true
  try {
    await Promise.all([
      loadStats(),
      loadRecords()
    ])
  } finally {
    loading.value = false
  }
})
</script>

<style scoped lang="scss">
.learning-records-page {
  padding: var(--spacing-xl) 0;
  background: var(--bg-page);
  min-height: 100vh;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 var(--spacing-lg);
}

.page-header {
  margin-bottom: var(--spacing-xl);
  
  .page-title {
    font-size: var(--font-size-h1);
    font-weight: var(--font-weight-bold);
    color: var(--text-primary);
    margin-bottom: var(--spacing-sm);
  }
  
  .page-description {
    font-size: var(--font-size-large);
    color: var(--text-secondary);
  }
}

.loading-container {
  padding: var(--spacing-xl) 0;
}

.learning-content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xl);
}

.stats-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: var(--spacing-lg);
  
  .stat-card {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
    padding: var(--spacing-lg);
    background: var(--bg-card);
    border-radius: var(--border-radius-large);
    box-shadow: var(--shadow-card);
    
    .stat-icon {
      width: 50px;
      height: 50px;
      border-radius: 50%;
      background: var(--primary-color);
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      font-size: 20px;
    }
    
    .stat-content {
      .stat-number {
        font-size: var(--font-size-h2);
        font-weight: var(--font-weight-bold);
        color: var(--text-primary);
        margin-bottom: var(--spacing-xs);
      }
      
      .stat-label {
        font-size: var(--font-size-small);
        color: var(--text-secondary);
      }
    }
  }
}

.filter-section {
  display: flex;
  gap: var(--spacing-lg);
  align-items: center;
  background: var(--bg-card);
  padding: var(--spacing-lg);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
  
  .search-box {
    flex: 1;
    max-width: 400px;
  }
  
  .filter-options {
    display: flex;
    gap: var(--spacing-md);
  }
}

.records-container {
  .records-list {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-lg);
  }
  
  .record-card {
    background: var(--bg-card);
    border-radius: var(--border-radius-large);
    box-shadow: var(--shadow-card);
    padding: var(--spacing-lg);
    
    .course-info {
      display: flex;
      gap: var(--spacing-lg);
      margin-bottom: var(--spacing-lg);
      
      .course-cover {
        position: relative;
        flex-shrink: 0;
        
        img {
          width: 120px;
          height: 80px;
          object-fit: cover;
          border-radius: var(--border-radius-medium);
        }
        
        .course-status {
          position: absolute;
          top: var(--spacing-xs);
          right: var(--spacing-xs);
        }
      }
      
      .course-details {
        flex: 1;
        
        .course-title {
          font-size: var(--font-size-h3);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
          margin: 0 0 var(--spacing-sm) 0;
        }
        
        .course-description {
          color: var(--text-secondary);
          margin: 0 0 var(--spacing-md) 0;
          line-height: 1.5;
        }
        
        .mentor-info {
          display: flex;
          align-items: center;
          gap: var(--spacing-xs);
          margin-bottom: var(--spacing-md);
          
          .mentor-name {
            font-size: var(--font-size-small);
            color: var(--text-secondary);
          }
        }
        
        .progress-info {
          margin-bottom: var(--spacing-md);
          
          .progress-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: var(--spacing-xs);
            
            .progress-text {
              font-size: var(--font-size-small);
              color: var(--text-secondary);
            }
            
            .progress-percentage {
              font-size: var(--font-size-small);
              font-weight: var(--font-weight-medium);
              color: var(--text-primary);
            }
          }
          
          .progress-details {
            display: flex;
            gap: var(--spacing-lg);
            margin-top: var(--spacing-xs);
            
            span {
              font-size: var(--font-size-small);
              color: var(--text-secondary);
            }
          }
        }
        
        .record-meta {
          display: flex;
          flex-wrap: wrap;
          gap: var(--spacing-lg);
          
          .meta-item {
            display: flex;
            align-items: center;
            gap: var(--spacing-xs);
            font-size: var(--font-size-small);
            color: var(--text-secondary);
          }
        }
      }
    }
    
    .record-actions {
      display: flex;
      gap: var(--spacing-md);
      justify-content: flex-end;
    }
  }
}

.load-more {
  text-align: center;
  padding: var(--spacing-lg) 0;
}

.empty-state {
  padding: var(--spacing-xl) 0;
}

.learning-detail {
  .detail-header {
    margin-bottom: var(--spacing-lg);
    
    h3 {
      font-size: var(--font-size-h3);
      font-weight: var(--font-weight-medium);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-md) 0;
    }
    
    .detail-stats {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
      gap: var(--spacing-md);
      
      .stat-item {
        text-align: center;
        padding: var(--spacing-md);
        background: var(--bg-light);
        border-radius: var(--border-radius-medium);
        
        .stat-label {
          display: block;
          font-size: var(--font-size-small);
          color: var(--text-secondary);
          margin-bottom: var(--spacing-xs);
        }
        
        .stat-value {
          font-size: var(--font-size-large);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
        }
      }
    }
  }
  
  .detail-content {
    h4 {
      font-size: var(--font-size-medium);
      font-weight: var(--font-weight-medium);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-md) 0;
    }
    
    .learning-timeline {
      .timeline-item {
        display: flex;
        gap: var(--spacing-md);
        margin-bottom: var(--spacing-md);
        
        .timeline-marker {
          width: 8px;
          height: 8px;
          border-radius: 50%;
          background: var(--primary-color);
          margin-top: 6px;
          flex-shrink: 0;
        }
        
        .timeline-content {
          flex: 1;
          
          .log-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: var(--spacing-xs);
            
            .log-action {
              font-size: var(--font-size-small);
              font-weight: var(--font-weight-medium);
              color: var(--text-primary);
            }
            
            .log-time {
              font-size: var(--font-size-small);
              color: var(--text-secondary);
            }
          }
          
          .log-details {
            font-size: var(--font-size-small);
            color: var(--text-secondary);
          }
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .filter-section {
    flex-direction: column;
    gap: var(--spacing-md);
    
    .search-box {
      max-width: none;
    }
  }
  
  .record-card {
    .course-info {
      flex-direction: column;
      gap: var(--spacing-md);
      
      .course-cover {
        align-self: center;
      }
    }
    
    .record-actions {
      flex-direction: column;
    }
  }
  
  .stats-overview {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style> 