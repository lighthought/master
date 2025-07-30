<template>
    <div class="income-stats-page">
      <div class="container">
        <!-- 页面头部 -->
        <div class="page-header">
          <h1 class="page-title">收入统计</h1>
          <p class="page-description">查看您的收入情况和财务数据</p>
        </div>
  
        <!-- 加载状态 -->
        <div v-if="loading" class="loading-container">
          <el-skeleton :rows="10" animated />
        </div>
  
        <div v-else class="stats-content">
          <!-- 收入概览 -->
          <div class="income-overview">
            <div class="overview-card total-income">
              <div class="card-icon">
                <el-icon><Money /></el-icon>
              </div>
              <div class="card-content">
                <div class="card-title">总收入</div>
                <div class="card-amount">¥{{ formatMoney(incomeStats.totalIncome) }}</div>
                <div class="card-growth">
                  <el-icon><TrendCharts /></el-icon>
                  <span>+{{ incomeStats.incomeGrowth }}%</span>
                </div>
              </div>
            </div>
  
            <div class="overview-card monthly-income">
              <div class="card-icon">
                <el-icon><Calendar /></el-icon>
              </div>
              <div class="card-content">
                <div class="card-title">本月收入</div>
                <div class="card-amount">¥{{ formatMoney(incomeStats.monthlyIncome) }}</div>
                <div class="card-subtitle">较上月增长 12.5%</div>
              </div>
            </div>
  
            <div class="overview-card total-students">
              <div class="card-icon">
                <el-icon><User /></el-icon>
              </div>
              <div class="card-content">
                <div class="card-title">总学生数</div>
                <div class="card-amount">{{ incomeStats.totalStudents }}</div>
                <div class="card-subtitle">人均收入 ¥{{ formatMoney(incomeStats.averagePerStudent) }}</div>
              </div>
            </div>
  
            <div class="overview-card avg-rating">
              <div class="card-icon">
                <el-icon><Star /></el-icon>
              </div>
              <div class="card-content">
                <div class="card-title">平均评分</div>
                <div class="card-amount">4.8</div>
                <div class="card-subtitle">来自学生评价</div>
              </div>
            </div>
          </div>
  
          <!-- 图表区域 -->
          <div class="charts-section">
            <div class="chart-card">
              <div class="chart-header">
                <h3>收入趋势</h3>
                <el-select v-model="chartPeriod" size="small">
                  <el-option label="最近6个月" value="6months" />
                  <el-option label="最近12个月" value="12months" />
                  <el-option label="今年" value="thisYear" />
                </el-select>
              </div>
              <div class="chart-container">
                <div class="chart-placeholder">
                  <el-icon><TrendCharts /></el-icon>
                  <p>收入趋势图表</p>
                  <small>显示月度收入变化趋势</small>
                </div>
              </div>
            </div>
  
            <div class="chart-card">
              <div class="chart-header">
                <h3>热门课程收入</h3>
              </div>
              <div class="top-courses">
                <div 
                  v-for="course in incomeStats.topCourses" 
                  :key="course.id"
                  class="course-item"
                >
                  <div class="course-info">
                    <h4>{{ course.title }}</h4>
                    <div class="course-stats">
                      <span>{{ course.students }} 名学生</span>
                      <el-rate v-model="course.averageRating" disabled />
                    </div>
                  </div>
                  <div class="course-income">
                    <div class="income-amount">¥{{ formatMoney(course.income) }}</div>
                    <div class="income-percent">{{ getIncomePercent(course.income) }}%</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
  
          <!-- 收入明细 -->
          <div class="income-details">
            <div class="details-header">
              <h3>收入明细</h3>
              <div class="header-actions">
                <el-button type="primary" @click="exportReport" :loading="exporting">
                  <el-icon><Download /></el-icon>
                  导出报表
                </el-button>
              </div>
            </div>
  
            <!-- 筛选条件 -->
            <div class="filter-section">
              <div class="filter-row">
                <el-select v-model="filterStatus" placeholder="交易状态" clearable @change="loadTransactions">
                  <el-option label="全部" value="" />
                  <el-option label="已完成" value="completed" />
                  <el-option label="待处理" value="pending" />
                  <el-option label="已取消" value="cancelled" />
                </el-select>
  
                <el-select v-model="filterType" placeholder="交易类型" clearable @change="loadTransactions">
                  <el-option label="全部" value="" />
                  <el-option label="课程报名" value="course_enrollment" />
                  <el-option label="一对一指导" value="mentoring" />
                  <el-option label="其他" value="other" />
                </el-select>
  
                <el-date-picker
                  v-model="dateRange"
                  type="daterange"
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                  @change="loadTransactions"
                />
  
                <el-input
                  v-model="searchKeyword"
                  placeholder="搜索学生姓名或课程..."
                  clearable
                  @input="handleSearch"
                  style="width: 200px;"
                >
                  <template #prefix>
                    <el-icon><Search /></el-icon>
                  </template>
                </el-input>
              </div>
            </div>
  
            <!-- 交易列表 -->
            <div class="transactions-container">
              <div v-if="loadingTransactions" class="loading-container">
                <el-skeleton :rows="5" animated />
              </div>
              
              <div v-else-if="transactions.length > 0" class="transactions-list">
                <div 
                  v-for="transaction in transactions" 
                  :key="transaction.id"
                  class="transaction-item"
                >
                  <div class="transaction-info">
                    <div class="transaction-main">
                      <h4>{{ transaction.studentName }}</h4>
                      <p>{{ transaction.courseTitle }}</p>
                    </div>
                    <div class="transaction-meta">
                      <el-tag :type="getStatusType(transaction.status)" size="small">
                        {{ getStatusText(transaction.status) }}
                      </el-tag>
                      <el-tag :type="getTypeType(transaction.type)" size="small">
                        {{ getTypeText(transaction.type) }}
                      </el-tag>
                    </div>
                  </div>
                  
                  <div class="transaction-amount">
                    <div class="amount">¥{{ formatMoney(transaction.amount) }}</div>
                    <div class="time">{{ formatTime(transaction.createdAt) }}</div>
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
                <el-empty description="暂无收入记录" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, onMounted, computed } from 'vue'
  import { ElMessage } from 'element-plus'
  import {
    Money, Calendar, User, Star, TrendCharts, Download, Search
  } from '@element-plus/icons-vue'
  import { useAuthStore } from '@/stores/auth'
  import { ApiService } from '@/services/api'
  
  // Store
  const authStore = useAuthStore()
  
  // 状态
  const loading = ref(false)
  const loadingTransactions = ref(false)
  const loadingMore = ref(false)
  const exporting = ref(false)
  
  // 数据
  const incomeStats = ref({
    totalIncome: 0,
    monthlyIncome: 0,
    totalStudents: 0,
    averagePerStudent: 0,
    incomeGrowth: 0,
    topCourses: [] as any[],
    monthlyData: [] as any[],
    recentTransactions: [] as any[]
  })
  const transactions = ref<any[]>([])
  const searchKeyword = ref('')
  const filterStatus = ref('')
  const filterType = ref('')
  const dateRange = ref<any>(null)
  const chartPeriod = ref('6months')
  const currentPage = ref(1)
  const pageSize = ref(10)
  const total = ref(0)
  
  // 计算属性
  const hasMore = computed(() => {
    return transactions.value.length < total.value
  })
  
  // 方法
  const loadIncomeStats = async () => {
    try {
      const result = await ApiService.income.getIncomeStats(authStore.user!.id)
      incomeStats.value = result.data
    } catch (error) {
      ElMessage.error('加载收入统计失败')
    }
  }
  
  const loadTransactions = async (reset = true) => {
    if (reset) {
      currentPage.value = 1
      transactions.value = []
    }
    
    loadingTransactions.value = true
    try {
      const params = {
        page: currentPage.value,
        pageSize: pageSize.value,
        keyword: searchKeyword.value,
        status: filterStatus.value,
        type: filterType.value,
        startDate: dateRange.value?.[0]?.toISOString(),
        endDate: dateRange.value?.[1]?.toISOString(),
        sortBy: 'date'
      }
      
      const result = await ApiService.income.getIncomeDetails(authStore.user!.id, params)
      
      if (reset) {
        transactions.value = result.data.transactions
      } else {
        transactions.value.push(...result.data.transactions)
      }
      
      total.value = result.data.total
    } catch (error) {
      ElMessage.error('加载收入明细失败')
    } finally {
      loadingTransactions.value = false
    }
  }
  
  const loadMore = async () => {
    currentPage.value++
    await loadTransactions(false)
  }
  
  const handleSearch = () => {
    loadTransactions()
  }
  
  const exportReport = async () => {
    exporting.value = true
    try {
      const result = await ApiService.income.exportIncomeReport(authStore.user!.id, {
        startDate: dateRange.value?.[0]?.toISOString(),
        endDate: dateRange.value?.[1]?.toISOString()
      })
      
      // 模拟下载
      const link = document.createElement('a')
      link.href = result.data.downloadUrl
      link.download = `收入报表_${new Date().toISOString().split('T')[0]}.pdf`
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
      
      ElMessage.success('报表导出成功')
    } catch (error) {
      ElMessage.error('报表导出失败')
    } finally {
      exporting.value = false
    }
  }
  
  const getStatusType = (status: string) => {
    const statusMap: Record<string, 'primary' | 'success' | 'warning' | 'info'> = {
      completed: 'success',
      pending: 'warning',
      cancelled: 'info'
    }
    return statusMap[status] || 'info'
  }
  
  const getStatusText = (status: string) => {
    const statusMap: Record<string, string> = {
      completed: '已完成',
      pending: '待处理',
      cancelled: '已取消'
    }
    return statusMap[status] || '未知'
  }
  
  const getTypeType = (type: string) => {
    const typeMap: Record<string, 'primary' | 'success' | 'warning' | 'info'> = {
      course_enrollment: 'primary',
      mentoring: 'success',
      other: 'info'
    }
    return typeMap[type] || 'info'
  }
  
  const getTypeText = (type: string) => {
    const typeMap: Record<string, string> = {
      course_enrollment: '课程报名',
      mentoring: '一对一指导',
      other: '其他'
    }
    return typeMap[type] || '未知'
  }
  
  const getIncomePercent = (income: number) => {
    const total = incomeStats.value.topCourses.reduce((sum, course) => sum + course.income, 0)
    return total > 0 ? Math.round((income / total) * 100) : 0
  }
  
  const formatMoney = (amount: number) => {
    return amount.toLocaleString()
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
        loadIncomeStats(),
        loadTransactions()
      ])
    } finally {
      loading.value = false
    }
  })
  </script>
  
  <style scoped lang="scss">
  .income-stats-page {
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
  
  .stats-content {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-xl);
  }
  
  .income-overview {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: var(--spacing-lg);
    
    .overview-card {
      display: flex;
      align-items: center;
      gap: var(--spacing-md);
      padding: var(--spacing-lg);
      background: var(--bg-card);
      border-radius: var(--border-radius-large);
      box-shadow: var(--shadow-card);
      
      .card-icon {
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
      
      .card-content {
        flex: 1;
        
        .card-title {
          font-size: var(--font-size-small);
          color: var(--text-secondary);
          margin-bottom: var(--spacing-xs);
        }
        
        .card-amount {
          font-size: var(--font-size-h2);
          font-weight: var(--font-weight-bold);
          color: var(--text-primary);
          margin-bottom: var(--spacing-xs);
        }
        
        .card-subtitle {
          font-size: var(--font-size-small);
          color: var(--text-secondary);
        }
        
        .card-growth {
          display: flex;
          align-items: center;
          gap: var(--spacing-xs);
          font-size: var(--font-size-small);
          color: var(--success-color);
        }
      }
    }
  }
  
  .charts-section {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: var(--spacing-lg);
    
    .chart-card {
      background: var(--bg-card);
      border-radius: var(--border-radius-large);
      box-shadow: var(--shadow-card);
      padding: var(--spacing-lg);
      
      .chart-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: var(--spacing-lg);
        
        h3 {
          font-size: var(--font-size-h3);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
          margin: 0;
        }
      }
      
      .chart-container {
        height: 200px;
        display: flex;
        align-items: center;
        justify-content: center;
        
        .chart-placeholder {
          text-align: center;
          color: var(--text-secondary);
          
          .el-icon {
            font-size: 48px;
            margin-bottom: var(--spacing-sm);
          }
          
          p {
            font-size: var(--font-size-medium);
            margin: 0 0 var(--spacing-xs) 0;
          }
          
          small {
            font-size: var(--font-size-small);
          }
        }
      }
      
      .top-courses {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-md);
        
        .course-item {
          display: flex;
          justify-content: space-between;
          align-items: center;
          padding: var(--spacing-md);
          border: 1px solid var(--border-color);
          border-radius: var(--border-radius-medium);
          
          .course-info {
            flex: 1;
            
            h4 {
              font-size: var(--font-size-medium);
              font-weight: var(--font-weight-medium);
              color: var(--text-primary);
              margin: 0 0 var(--spacing-xs) 0;
            }
            
            .course-stats {
              display: flex;
              align-items: center;
              gap: var(--spacing-md);
              font-size: var(--font-size-small);
              color: var(--text-secondary);
            }
          }
          
          .course-income {
            text-align: right;
            
            .income-amount {
              font-size: var(--font-size-medium);
              font-weight: var(--font-weight-medium);
              color: var(--text-primary);
              margin-bottom: var(--spacing-xs);
            }
            
            .income-percent {
              font-size: var(--font-size-small);
              color: var(--text-secondary);
            }
          }
        }
      }
    }
  }
  
  .income-details {
    background: var(--bg-card);
    border-radius: var(--border-radius-large);
    box-shadow: var(--shadow-card);
    padding: var(--spacing-lg);
    
    .details-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: var(--spacing-lg);
      
      h3 {
        font-size: var(--font-size-h3);
        font-weight: var(--font-weight-medium);
        color: var(--text-primary);
        margin: 0;
      }
    }
    
    .filter-section {
      margin-bottom: var(--spacing-lg);
      
      .filter-row {
        display: flex;
        gap: var(--spacing-md);
        align-items: center;
        flex-wrap: wrap;
      }
    }
    
    .transactions-container {
      .transactions-list {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-md);
        
        .transaction-item {
          display: flex;
          justify-content: space-between;
          align-items: center;
          padding: var(--spacing-md);
          border: 1px solid var(--border-color);
          border-radius: var(--border-radius-medium);
          
          .transaction-info {
            flex: 1;
            
            .transaction-main {
              margin-bottom: var(--spacing-sm);
              
              h4 {
                font-size: var(--font-size-medium);
                font-weight: var(--font-weight-medium);
                color: var(--text-primary);
                margin: 0 0 var(--spacing-xs) 0;
              }
              
              p {
                font-size: var(--font-size-small);
                color: var(--text-secondary);
                margin: 0;
              }
            }
            
            .transaction-meta {
              display: flex;
              gap: var(--spacing-sm);
            }
          }
          
          .transaction-amount {
            text-align: right;
            
            .amount {
              font-size: var(--font-size-medium);
              font-weight: var(--font-weight-medium);
              color: var(--text-primary);
              margin-bottom: var(--spacing-xs);
            }
            
            .time {
              font-size: var(--font-size-small);
              color: var(--text-secondary);
            }
          }
        }
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


// 响应式设计
@media (max-width: 768px) {
  .income-overview {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .charts-section {
    grid-template-columns: 1fr;
  }
  
  .filter-section {
    .filter-row {
      flex-direction: column;
      align-items: stretch;
    }
  }
  
  .transaction-item {
    flex-direction: column;
    align-items: stretch;
    text-align: center;
    
    .transaction-amount {
      text-align: center;
      margin-top: var(--spacing-sm);
    }
  }
}

@media (max-width: 480px) {
  .income-overview {
    grid-template-columns: 1fr;
  }
}
</style>