<template>
    <div class="student-management-page">
      <div class="container">
        <!-- 页面头部 -->
        <div class="page-header">
          <h1 class="page-title">学生管理</h1>
          <p class="page-description">管理您的学生，提供更好的指导服务</p>
        </div>
  
        <!-- 加载状态 -->
        <div v-if="loading" class="loading-container">
          <el-skeleton :rows="10" animated />
        </div>
  
        <div v-else class="management-content">
          <!-- 学生统计概览 -->
          <div class="stats-overview">
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><User /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ studentStats.totalStudents }}</div>
                <div class="stat-label">总学生数</div>
              </div>
            </div>
            
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><VideoPlay /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ studentStats.activeStudents }}</div>
                <div class="stat-label">活跃学生</div>
              </div>
            </div>
            
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><Clock /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ studentStats.totalTeachingHours }}</div>
                <div class="stat-label">总教学时长(小时)</div>
              </div>
            </div>
            
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><Star /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ studentStats.averageRating }}</div>
                <div class="stat-label">平均评分</div>
              </div>
            </div>
          </div>
  
          <!-- 筛选和搜索 -->
          <div class="filter-section">
            <div class="search-box">
              <el-input
                v-model="searchKeyword"
                placeholder="搜索学生姓名或课程..."
                clearable
                @input="handleSearch"
              >
                <template #prefix>
                  <el-icon><Search /></el-icon>
                </template>
              </el-input>
            </div>
            <div class="filter-options">
              <el-select v-model="filterStatus" placeholder="学生状态" clearable @change="loadStudents">
                <el-option label="全部" value="" />
                <el-option label="学习中" value="learning" />
                <el-option label="已完成" value="completed" />
                <el-option label="已暂停" value="paused" />
              </el-select>
              <el-select v-model="sortBy" placeholder="排序方式" @change="loadStudents">
                <el-option label="最近学习" value="recent" />
                <el-option label="学习进度" value="progress" />
                <el-option label="报名时间" value="enrollTime" />
              </el-select>
            </div>
          </div>
  
          <!-- 学生列表 -->
          <div class="students-container">
            <div v-if="loadingStudents" class="loading-container">
              <el-skeleton :rows="3" animated />
              <el-skeleton :rows="3" animated />
              <el-skeleton :rows="3" animated />
            </div>
            
            <div v-else-if="students.length > 0" class="students-list">
              <div 
                v-for="student in students" 
                :key="student.id"
                class="student-card"
              >
                <div class="student-info">
                  <div class="student-avatar">
                    <el-avatar :size="60" :src="student.avatar" />
                    <div class="student-status">
                      <el-tag 
                        :type="getStatusType(student.status)"
                        size="small"
                      >
                        {{ getStatusText(student.status) }}
                      </el-tag>
                    </div>
                  </div>
                  
                  <div class="student-details">
                    <h3 class="student-name">{{ student.name }}</h3>
                    <p class="student-email">{{ student.email }}</p>
                    
                    <div class="enrolled-courses">
                      <h4>已报名课程</h4>
                      <div class="courses-list">
                        <div 
                          v-for="course in student.enrolledCourses" 
                          :key="course.id"
                          class="course-item"
                        >
                          <span class="course-name">{{ course.title }}</span>
                          <div class="course-progress">
                            <el-progress 
                              :percentage="course.progress" 
                              :stroke-width="4"
                              :show-text="false"
                              :color="getProgressColor(course.progress)"
                            />
                            <span class="progress-text">{{ course.progress }}%</span>
                          </div>
                        </div>
                      </div>
                    </div>
                    
                    <div class="student-meta">
                      <div class="meta-item">
                        <el-icon><Calendar /></el-icon>
                        <span>报名时间：{{ formatDate(student.enrollTime) }}</span>
                      </div>
                      <div class="meta-item">
                        <el-icon><Clock /></el-icon>
                        <span>最近学习：{{ formatTime(student.lastStudyTime) }}</span>
                      </div>
                      <div class="meta-item">
                        <el-icon><VideoPlay /></el-icon>
                        <span>学习时长：{{ student.totalStudyTime }} 小时</span>
                      </div>
                    </div>
                  </div>
                </div>
                
                <div class="student-actions">
                  <el-button 
                    type="primary" 
                    @click="viewStudentDetail(student)"
                  >
                    <el-icon><InfoFilled /></el-icon>
                    查看详情
                  </el-button>
                  
                  <el-button 
                    type="default" 
                    @click="startCommunication(student)"
                  >
                    <el-icon><ChatDotRound /></el-icon>
                    开始沟通
                  </el-button>
                  
                  <el-button 
                    type="text" 
                    @click="viewStudentProgress(student)"
                  >
                    <el-icon><TrendCharts /></el-icon>
                    学习进度
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
              <el-empty description="暂无学生">
                <el-button type="primary" @click="$router.push('/courses')">
                  发布课程
                </el-button>
              </el-empty>
            </div>
          </div>
        </div>
      </div>
  
      <!-- 学生详情对话框 -->
      <el-dialog
        v-model="showStudentDetailDialog"
        title="学生详情"
        width="90%"
        max-width="800px"
        :close-on-click-modal="false"
      >
        <div v-if="selectedStudent" class="student-detail">
          <div class="detail-header">
            <div class="student-profile">
              <el-avatar :size="80" :src="selectedStudent.avatar" />
              <div class="profile-info">
                <h3>{{ selectedStudent.name }}</h3>
                <p>{{ selectedStudent.email }}</p>
                <el-tag :type="getStatusType(selectedStudent.status)">
                  {{ getStatusText(selectedStudent.status) }}
                </el-tag>
              </div>
            </div>
          </div>
          
          <div class="detail-content">
            <el-tabs v-model="activeTab">
              <el-tab-pane label="学习进度" name="progress">
                <div class="progress-section">
                  <h4>课程学习进度</h4>
                  <div class="courses-progress">
                    <div 
                      v-for="course in selectedStudent.enrolledCourses" 
                      :key="course.id"
                      class="course-progress-item"
                    >
                      <div class="course-header">
                        <span class="course-title">{{ course.title }}</span>
                        <span class="course-progress-percent">{{ course.progress }}%</span>
                      </div>
                      <el-progress 
                        :percentage="course.progress" 
                        :stroke-width="8"
                        :show-text="false"
                        :color="getProgressColor(course.progress)"
                      />
                      <div class="course-stats">
                        <span>已完成 {{ course.completedLessons }}/{{ course.totalLessons }} 课时</span>
                        <span>学习时长 {{ course.studyTime }} 小时</span>
                      </div>
                    </div>
                  </div>
                </div>
              </el-tab-pane>
              
              <el-tab-pane label="评价反馈" name="feedback">
                <div class="feedback-section">
                  <h4>学生评价</h4>
                  <div v-if="selectedStudent.feedback && selectedStudent.feedback.length > 0" class="feedback-list">
                    <div 
                      v-for="feedback in selectedStudent.feedback" 
                      :key="feedback.id"
                      class="feedback-item"
                    >
                      <div class="feedback-header">
                        <span class="feedback-course">{{ feedback.courseTitle }}</span>
                        <el-rate v-model="feedback.rating" disabled />
                        <span class="feedback-time">{{ formatTime(feedback.createdAt) }}</span>
                      </div>
                      <p class="feedback-content">{{ feedback.content }}</p>
                    </div>
                  </div>
                  <div v-else class="no-feedback">
                    <el-empty description="暂无评价反馈" />
                  </div>
                </div>
              </el-tab-pane>
              
              <el-tab-pane label="沟通记录" name="communication">
                <div class="communication-section">
                  <h4>沟通记录</h4>
                  <div class="communication-input">
                    <el-input
                      v-model="newMessage"
                      type="textarea"
                      :rows="3"
                      placeholder="输入消息内容..."
                      @keyup.enter="sendMessage"
                    />
                    <el-button 
                      type="primary" 
                      @click="sendMessage"
                      :loading="sendingMessage"
                      style="margin-top: 10px;"
                    >
                      发送消息
                    </el-button>
                  </div>
                  <div class="message-list">
                    <div 
                      v-for="message in selectedStudent.messages" 
                      :key="message.id"
                      class="message-item"
                      :class="{ 'sent': message.senderId === 'master' }"
                    >
                      <div class="message-content">
                        <p>{{ message.content }}</p>
                        <span class="message-time">{{ formatTime(message.createdAt) }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </el-tab-pane>
            </el-tabs>
          </div>
        </div>
      </el-dialog>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, onMounted, computed } from 'vue'
  import { ElMessage } from 'element-plus'
  import {
    User, VideoPlay, Clock, Star, Search, Calendar, InfoFilled, ChatDotRound, TrendCharts
  } from '@element-plus/icons-vue'
  import { useAuthStore } from '@/stores/auth'
  import { ApiService } from '@/services/api'
  
  // Store
  const authStore = useAuthStore()
  
  // 状态
  const loading = ref(false)
  const loadingStudents = ref(false)
  const loadingMore = ref(false)
  const sendingMessage = ref(false)
  const showStudentDetailDialog = ref(false)
  const activeTab = ref('progress')
  
  // 数据
  const studentStats = ref({
    totalStudents: 0,
    activeStudents: 0,
    totalTeachingHours: 0,
    averageRating: 0
  })
  const students = ref<any[]>([])
  const searchKeyword = ref('')
  const filterStatus = ref('')
  const sortBy = ref('recent')
  const currentPage = ref(1)
  const pageSize = ref(10)
  const total = ref(0)
  const selectedStudent = ref<any>(null)
  const newMessage = ref('')
  
  // 计算属性
  const hasMore = computed(() => {
    return students.value.length < total.value
  })
  
  // 方法
  const loadStats = async () => {
    try {
      const result = await ApiService.master.getStudentStats(authStore.user!.id)
      studentStats.value = result.data
    } catch (error) {
      ElMessage.error('加载学生统计失败')
    }
  }
  
  const loadStudents = async (reset = true) => {
    if (reset) {
      currentPage.value = 1
      students.value = []
    }
    
    loadingStudents.value = true
    try {
      const params = {
        page: currentPage.value,
        pageSize: pageSize.value,
        keyword: searchKeyword.value,
        status: filterStatus.value,
        sortBy: sortBy.value
      }
      
      const result = await ApiService.master.getStudents(authStore.user!.id, params)
      
      if (reset) {
        students.value = result.data.students
      } else {
        students.value.push(...result.data.students)
      }
      
      total.value = result.data.total
    } catch (error) {
      ElMessage.error('加载学生列表失败')
    } finally {
      loadingStudents.value = false
    }
  }
  
  const loadMore = async () => {
    currentPage.value++
    await loadStudents(false)
  }
  
  const handleSearch = () => {
    loadStudents()
  }
  
  const viewStudentDetail = (student: any) => {
    selectedStudent.value = student
    showStudentDetailDialog.value = true
  }
  
  const startCommunication = (student: any) => {
    selectedStudent.value = student
    activeTab.value = 'communication'
    showStudentDetailDialog.value = true
  }
  
  const viewStudentProgress = (student: any) => {
    selectedStudent.value = student
    activeTab.value = 'progress'
    showStudentDetailDialog.value = true
  }
  
  const sendMessage = async () => {
    if (!newMessage.value.trim()) {
      ElMessage.warning('请输入消息内容')
      return
    }
    
    sendingMessage.value = true
    try {
      const messageData = {
        receiverId: selectedStudent.value.id,
        content: newMessage.value.trim(),
        type: 'text'
      }
      
      await ApiService.master.sendMessage(authStore.user!.id, messageData)
      
      // 添加到消息列表
      if (!selectedStudent.value.messages) {
        selectedStudent.value.messages = []
      }
      
      selectedStudent.value.messages.push({
        id: Date.now().toString(),
        content: newMessage.value.trim(),
        senderId: 'master',
        createdAt: new Date().toISOString()
      })
      
      newMessage.value = ''
      ElMessage.success('消息发送成功')
    } catch (error) {
      ElMessage.error('消息发送失败')
    } finally {
      sendingMessage.value = false
    }
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
        loadStudents()
      ])
    } finally {
      loading.value = false
    }
  })
  </script>
  
  <style scoped lang="scss">
  .student-management-page {
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
  
  .management-content {
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
  
  .students-container {
    .students-list {
      display: flex;
      flex-direction: column;
      gap: var(--spacing-lg);
    }
    
    .student-card {
      background: var(--bg-card);
      border-radius: var(--border-radius-large);
      box-shadow: var(--shadow-card);
      padding: var(--spacing-lg);
      
      .student-info {
        display: flex;
        gap: var(--spacing-lg);
        margin-bottom: var(--spacing-lg);
        
        .student-avatar {
          position: relative;
          flex-shrink: 0;
          
          .student-status {
            position: absolute;
            bottom: -5px;
            right: -5px;
          }
        }
        
        .student-details {
          flex: 1;
          
          .student-name {
            font-size: var(--font-size-h3);
            font-weight: var(--font-weight-medium);
            color: var(--text-primary);
            margin: 0 0 var(--spacing-xs) 0;
          }
          
          .student-email {
            color: var(--text-secondary);
            margin: 0 0 var(--spacing-md) 0;
          }
          
          .enrolled-courses {
            margin-bottom: var(--spacing-md);
            
            h4 {
              font-size: var(--font-size-medium);
              font-weight: var(--font-weight-medium);
              color: var(--text-primary);
              margin: 0 0 var(--spacing-sm) 0;
            }
            
            .courses-list {
              display: flex;
              flex-direction: column;
              gap: var(--spacing-sm);
              
              .course-item {
                display: flex;
                align-items: center;
                gap: var(--spacing-md);
                
                .course-name {
                  font-size: var(--font-size-small);
                  color: var(--text-primary);
                  min-width: 120px;
                }
                
                .course-progress {
                  flex: 1;
                  display: flex;
                  align-items: center;
                  gap: var(--spacing-sm);
                  
                  .progress-text {
                    font-size: var(--font-size-small);
                    color: var(--text-secondary);
                    min-width: 40px;
                  }
                }
              }
            }
          }
          
          .student-meta {
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
      
      .student-actions {
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
  
  .student-detail {
    .detail-header {
      margin-bottom: var(--spacing-lg);
      
      .student-profile {
        display: flex;
        align-items: center;
        gap: var(--spacing-lg);
        
        .profile-info {
          h3 {
            font-size: var(--font-size-h3);
            font-weight: var(--font-weight-medium);
            color: var(--text-primary);
            margin: 0 0 var(--spacing-xs) 0;
          }
          
          p {
            color: var(--text-secondary);
            margin: 0 0 var(--spacing-sm) 0;
          }
        }
      }
    }
    
    .detail-content {
      .progress-section {
        h4 {
          font-size: var(--font-size-medium);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
          margin: 0 0 var(--spacing-md) 0;
        }
        
        .courses-progress {
          display: flex;
          flex-direction: column;
          gap: var(--spacing-lg);
          
          .course-progress-item {
            padding: var(--spacing-md);
            border: 1px solid var(--border-color);
            border-radius: var(--border-radius-medium);
            
            .course-header {
              display: flex;
              justify-content: space-between;
              align-items: center;
              margin-bottom: var(--spacing-sm);
              
              .course-title {
                font-size: var(--font-size-medium);
                font-weight: var(--font-weight-medium);
                color: var(--text-primary);
              }
              
              .course-progress-percent {
                font-size: var(--font-size-small);
                color: var(--text-secondary);
              }
            }
            
            .course-stats {
              display: flex;
              justify-content: space-between;
              margin-top: var(--spacing-sm);
              font-size: var(--font-size-small);
              color: var(--text-secondary);
            }
          }
        }
      }
      
      .feedback-section {
        h4 {
          font-size: var(--font-size-medium);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
          margin: 0 0 var(--spacing-md) 0;
        }
        
        .feedback-list {
          display: flex;
          flex-direction: column;
          gap: var(--spacing-md);
          
          .feedback-item {
            padding: var(--spacing-md);
            border: 1px solid var(--border-color);
            border-radius: var(--border-radius-medium);
            
            .feedback-header {
              display: flex;
              align-items: center;
              gap: var(--spacing-md);
              margin-bottom: var(--spacing-sm);
              
              .feedback-course {
                font-size: var(--font-size-medium);
                font-weight: var(--font-weight-medium);
                color: var(--text-primary);
              }
              
              .feedback-time {
                font-size: var(--font-size-small);
                color: var(--text-secondary);
              }
            }
            
            .feedback-content {
              color: var(--text-secondary);
              margin: 0;
              line-height: 1.5;
            }
          }
        }
        
        .no-feedback {
          padding: var(--spacing-xl) 0;
        }
      }
      
      .communication-section {
        h4 {
          font-size: var(--font-size-medium);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
          margin: 0 0 var(--spacing-md) 0;
        }
        
        .communication-input {
          margin-bottom: var(--spacing-lg);
        }
        
        .message-list {
          display: flex;
          flex-direction: column;
          gap: var(--spacing-md);
          max-height: 300px;
          overflow-y: auto;
          
          .message-item {
            display: flex;
            
            &.sent {
              justify-content: flex-end;
              
              .message-content {
                background: var(--primary-color);
                color: white;
                
                .message-time {
                  color: rgba(255, 255, 255, 0.8);
                }
              }
            }
            
            .message-content {
              max-width: 70%;
              padding: var(--spacing-sm) var(--spacing-md);
              background: var(--bg-secondary);
              border-radius: var(--border-radius-medium);
              
              p {
                margin: 0 0 var(--spacing-xs) 0;
                line-height: 1.4;
              }
              
              .message-time {
                font-size: var(--font-size-small);
                color: var(--text-secondary);
              }
            }
          }
        }
      }
    }
  }
  
  // 响应式设计
  @media (max-width: 768px) {
    .stats-overview {
      grid-template-columns: repeat(2, 1fr);
    }
    
    .filter-section {
      flex-direction: column;
      align-items: stretch;
      
      .search-box {
        max-width: none;
      }
      
      .filter-options {
        justify-content: space-between;
      }
    }
    
    .student-card {
      .student-info {
        flex-direction: column;
        text-align: center;
        
        .student-avatar {
          align-self: center;
        }
      }
      
      .student-actions {
        flex-direction: column;
        align-items: stretch;
      }
    }
    
    .student-detail {
      .detail-header {
        .student-profile {
          flex-direction: column;
          text-align: center;
        }
      }
    }
  }
  
  @media (max-width: 480px) {
    .stats-overview {
      grid-template-columns: 1fr;
    }
    
    .filter-options {
      flex-direction: column;
    }
  }
  </style>