<template>
  <div class="circles-page">
    <!-- 页面头部 -->
    <div class="page-header">
      <h1 class="page-title">圈子</h1>
      <p class="page-subtitle">发现志同道合的伙伴，加入感兴趣的圈子</p>
    </div>

    <!-- 搜索和筛选区域 -->
    <div class="search-filter-section">
      <div class="search-box">
        <el-input
          v-model="searchQuery"
          placeholder="搜索圈子..."
          clearable
          @input="handleSearch"
          @clear="handleSearch"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
      </div>

      <div class="filter-options">
        <el-select v-model="selectedCategory" placeholder="选择分类" clearable @change="handleFilter">
          <el-option
            v-for="category in categories"
            :key="category.id"
            :label="category.name"
            :value="category.id"
          />
        </el-select>

        <el-select v-model="sortBy" placeholder="排序方式" @change="handleSort">
          <el-option label="成员数量" value="memberCount" />
          <el-option label="动态数量" value="postCount" />
          <el-option label="创建时间" value="createdAt" />
          <el-option label="名称排序" value="name" />
        </el-select>
      </div>
    </div>

    <!-- 分类标签 -->
    <div class="category-tabs">
      <el-tabs v-model="activeCategory" @tab-click="handleCategoryTab">
        <el-tab-pane label="全部" name="all" />
        <el-tab-pane
          v-for="category in categories"
          :key="category.id"
          :label="category.name"
          :name="category.id"
        />
      </el-tabs>
    </div>

    <!-- 圈子列表 -->
    <div class="circles-content">
      <div v-if="loading" class="loading-container">
        <el-skeleton :rows="6" animated />
      </div>

      <div v-else-if="circles.length > 0" class="circles-grid">
        <div
          v-for="circle in circles"
          :key="circle.id"
          class="circle-card"
          @click="viewCircleDetail(circle)"
        >
          <!-- 圈子封面 -->
          <div class="circle-cover">
            <img :src="circle.cover" :alt="circle.name" />
            <div class="circle-status">
              <el-tag v-if="circle.isActive" type="success" size="small">活跃</el-tag>
              <el-tag v-if="circle.isJoined" type="warning" size="small">已加入</el-tag>
            </div>
          </div>

          <!-- 圈子信息 -->
          <div class="circle-info">
            <h3 class="circle-name">{{ circle.name }}</h3>
            <p class="circle-description">{{ circle.description }}</p>

            <!-- 圈子标签 -->
            <div class="circle-tags">
              <el-tag
                v-for="tag in circle.tags.slice(0, 3)"
                :key="tag"
                size="small"
                effect="light"
                class="tag-item"
              >
                {{ tag }}
              </el-tag>
              <span v-if="circle.tags.length > 3" class="more-tags">
                +{{ circle.tags.length - 3 }}
              </span>
            </div>

            <!-- 圈子统计 -->
            <div class="circle-stats">
              <div class="stat-item">
                <el-icon><User /></el-icon>
                <span>{{ formatNumber(circle.memberCount) }} 成员</span>
              </div>
              <div class="stat-item">
                <el-icon><ChatDotRound /></el-icon>
                <span>{{ formatNumber(circle.postCount) }} 动态</span>
              </div>
            </div>

            <!-- 操作按钮 -->
            <div class="circle-actions">
              <el-button
                v-if="!circle.isJoined"
                type="primary"
                size="small"
                @click.stop="joinCircle(circle)"
                :loading="joiningCircleId === circle.id"
              >
                加入圈子
              </el-button>
              <el-button
                v-else
                type="default"
                size="small"
                @click.stop="leaveCircle(circle)"
                :loading="leavingCircleId === circle.id"
              >
                退出圈子
              </el-button>
              <el-button
                type="text"
                size="small"
                @click.stop="viewCircleDetail(circle)"
              >
                查看详情
              </el-button>
            </div>
          </div>
        </div>
      </div>

      <!-- 空状态 -->
      <div v-else class="empty-state">
        <el-empty description="暂无相关圈子">
          <el-button type="primary" @click="resetFilters">
            查看全部圈子
          </el-button>
        </el-empty>
      </div>

      <!-- 分页 -->
      <div v-if="totalPages > 1" class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="total"
          :page-sizes="[10, 20, 50]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>

    <!-- 圈子详情对话框 -->
    <el-dialog
      v-model="showCircleDetail"
      :title="selectedCircle?.name"
      width="80%"
      max-width="800px"
    >
      <div v-if="selectedCircle" class="circle-detail">
        <div class="detail-header">
          <img :src="selectedCircle.cover" :alt="selectedCircle.name" class="detail-cover" />
          <div class="detail-info">
            <h2>{{ selectedCircle.name }}</h2>
            <p>{{ selectedCircle.description }}</p>
            <div class="detail-stats">
              <span>{{ formatNumber(selectedCircle.memberCount) }} 成员</span>
              <span>{{ formatNumber(selectedCircle.postCount) }} 动态</span>
              <span>创建于 {{ formatDate(selectedCircle.createdAt) }}</span>
            </div>
          </div>
        </div>

        <div class="detail-tags">
          <h4>标签</h4>
          <div class="tags-list">
            <el-tag
              v-for="tag in selectedCircle.tags"
              :key="tag"
              size="medium"
              effect="light"
            >
              {{ tag }}
            </el-tag>
          </div>
        </div>

        <div class="detail-rules">
          <h4>圈子规则</h4>
          <ul class="rules-list">
            <li v-for="rule in selectedCircle.rules" :key="rule">{{ rule }}</li>
          </ul>
        </div>

        <div class="detail-actions">
          <el-button
            v-if="!selectedCircle.isJoined"
            type="primary"
            size="large"
            @click="joinCircle(selectedCircle)"
            :loading="joiningCircleId === selectedCircle.id"
          >
            加入圈子
          </el-button>
          <el-button
            v-else
            type="default"
            size="large"
            @click="leaveCircle(selectedCircle)"
            :loading="leavingCircleId === selectedCircle.id"
          >
            退出圈子
          </el-button>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Search, User, ChatDotRound } from '@element-plus/icons-vue'
import { ApiService } from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

// 响应式数据
const circles = ref<any[]>([])
const categories = ref<any[]>([])
const loading = ref(true)
const searchQuery = ref('')
const selectedCategory = ref('')
const sortBy = ref('')
const activeCategory = ref('all')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const totalPages = ref(0)
const joiningCircleId = ref('')
const leavingCircleId = ref('')
const showCircleDetail = ref(false)
const selectedCircle = ref<any>(null)

// 加载圈子列表
const loadCircles = async () => {
  loading.value = true
  
  try {
    const params = {
      category: selectedCategory.value || undefined,
      query: searchQuery.value || undefined,
      sort: sortBy.value || undefined,
      page: currentPage.value,
      pageSize: pageSize.value
    }
    
    const response = await ApiService.circles.getCircles(params)
    circles.value = response.data.circles
    total.value = response.data.total
    totalPages.value = response.data.totalPages
  } catch (error) {
    console.error('加载圈子列表失败:', error)
    ElMessage.error('加载圈子列表失败')
  } finally {
    loading.value = false
  }
}

// 加载圈子分类
const loadCategories = async () => {
  try {
    const response = await ApiService.circles.getCircleCategories()
    categories.value = response.data
  } catch (error) {
    console.error('加载圈子分类失败:', error)
  }
}

// 搜索处理
const handleSearch = () => {
  currentPage.value = 1
  loadCircles()
}

// 筛选处理
const handleFilter = () => {
  currentPage.value = 1
  loadCircles()
}

// 排序处理
const handleSort = () => {
  currentPage.value = 1
  loadCircles()
}

// 分类标签处理
const handleCategoryTab = (tab: any) => {
  selectedCategory.value = tab.name === 'all' ? '' : tab.name
  currentPage.value = 1
  loadCircles()
}

// 分页处理
const handleSizeChange = (size: number) => {
  pageSize.value = size
  currentPage.value = 1
  loadCircles()
}

const handleCurrentChange = (page: number) => {
  currentPage.value = page
  loadCircles()
}

// 加入圈子
const joinCircle = async (circle: any) => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录')
    router.push('/auth')
    return
  }
  
  joiningCircleId.value = circle.id
  
  try {
    await ApiService.circles.joinCircle(circle.id, authStore.user?.id || '1')
    ElMessage.success('加入圈子成功')
    
    // 更新圈子状态
    circle.isJoined = true
    circle.memberCount++
    
    // 重新加载列表
    loadCircles()
  } catch (error) {
    console.error('加入圈子失败:', error)
    ElMessage.error('加入圈子失败')
  } finally {
    joiningCircleId.value = ''
  }
}

// 退出圈子
const leaveCircle = async (circle: any) => {
  try {
    await ElMessageBox.confirm(
      '确定要退出该圈子吗？退出后将无法查看圈子内容。',
      '确认退出',
      {
        confirmButtonText: '确定退出',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    leavingCircleId.value = circle.id
    
    await ApiService.circles.leaveCircle(circle.id, authStore.user?.id || '1')
    ElMessage.success('退出圈子成功')
    
    // 更新圈子状态
    circle.isJoined = false
    circle.memberCount--
    
    // 重新加载列表
    loadCircles()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('退出圈子失败:', error)
      ElMessage.error('退出圈子失败')
    }
  } finally {
    leavingCircleId.value = ''
  }
}

// 查看圈子详情
const viewCircleDetail = (circle: any) => {
  selectedCircle.value = circle
  showCircleDetail.value = true
}

// 重置筛选
const resetFilters = () => {
  searchQuery.value = ''
  selectedCategory.value = ''
  sortBy.value = ''
  activeCategory.value = 'all'
  currentPage.value = 1
  loadCircles()
}

// 格式化数字
const formatNumber = (num: number) => {
  if (num >= 10000) {
    return (num / 10000).toFixed(1) + '万'
  }
  return num.toString()
}

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('zh-CN')
}

onMounted(() => {
  loadCategories()
  loadCircles()
})
</script>

<style scoped lang="scss">
.circles-page {
  padding: var(--spacing-xl);
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  text-align: center;
  margin-bottom: var(--spacing-xxl);
  
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
}

.search-filter-section {
  display: flex;
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-xl);
  align-items: center;
  
  .search-box {
    flex: 1;
    max-width: 400px;
  }
  
  .filter-options {
    display: flex;
    gap: var(--spacing-md);
  }
}

.category-tabs {
  margin-bottom: var(--spacing-xl);
  
  :deep(.el-tabs__header) {
    margin-bottom: var(--spacing-lg);
  }
}

.circles-content {
  .loading-container {
    padding: var(--spacing-xl) 0;
  }
  
  .circles-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    gap: var(--spacing-lg);
    margin-bottom: var(--spacing-xl);
  }
  
  .empty-state {
    padding: var(--spacing-xxl) 0;
    text-align: center;
  }
  
  .pagination-container {
    display: flex;
    justify-content: center;
    margin-top: var(--spacing-xl);
  }
}

.circle-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  overflow: hidden;
  box-shadow: var(--shadow-light);
  transition: all var(--transition-normal);
  cursor: pointer;
  
  &:hover {
    transform: translateY(-4px);
    box-shadow: var(--shadow-medium);
  }
  
  .circle-cover {
    position: relative;
    height: 160px;
    overflow: hidden;
    
    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
    
    .circle-status {
      position: absolute;
      top: var(--spacing-sm);
      right: var(--spacing-sm);
      display: flex;
      gap: var(--spacing-xs);
    }
  }
  
  .circle-info {
    padding: var(--spacing-lg);
    
    .circle-name {
      font-size: var(--font-size-h4);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-sm) 0;
      line-height: 1.3;
    }
    
    .circle-description {
      font-size: var(--font-size-medium);
      color: var(--text-secondary);
      margin: 0 0 var(--spacing-md) 0;
      line-height: 1.5;
      display: -webkit-box;
      -webkit-line-clamp: 2;
      -webkit-box-orient: vertical;
      overflow: hidden;
    }
    
    .circle-tags {
      display: flex;
      flex-wrap: wrap;
      gap: var(--spacing-xs);
      margin-bottom: var(--spacing-md);
      
      .tag-item {
        font-size: var(--font-size-small);
      }
      
      .more-tags {
        font-size: var(--font-size-small);
        color: var(--text-tertiary);
        align-self: center;
      }
    }
    
    .circle-stats {
      display: flex;
      gap: var(--spacing-lg);
      margin-bottom: var(--spacing-md);
      
      .stat-item {
        display: flex;
        align-items: center;
        gap: var(--spacing-xs);
        font-size: var(--font-size-small);
        color: var(--text-secondary);
        
        .el-icon {
          font-size: var(--font-size-medium);
        }
      }
    }
    
    .circle-actions {
      display: flex;
      gap: var(--spacing-sm);
      align-items: center;
    }
  }
}

.circle-detail {
  .detail-header {
    display: flex;
    gap: var(--spacing-lg);
    margin-bottom: var(--spacing-xl);
    
    .detail-cover {
      width: 200px;
      height: 120px;
      object-fit: cover;
      border-radius: var(--border-radius-medium);
    }
    
    .detail-info {
      flex: 1;
      
      h2 {
        font-size: var(--font-size-h3);
        font-weight: var(--font-weight-bold);
        color: var(--text-primary);
        margin: 0 0 var(--spacing-sm) 0;
      }
      
      p {
        font-size: var(--font-size-medium);
        color: var(--text-secondary);
        margin: 0 0 var(--spacing-md) 0;
        line-height: 1.6;
      }
      
      .detail-stats {
        display: flex;
        gap: var(--spacing-lg);
        font-size: var(--font-size-small);
        color: var(--text-tertiary);
      }
    }
  }
  
  .detail-tags,
  .detail-rules {
    margin-bottom: var(--spacing-lg);
    
    h4 {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-md) 0;
    }
    
    .tags-list {
      display: flex;
      flex-wrap: wrap;
      gap: var(--spacing-sm);
    }
    
    .rules-list {
      list-style: none;
      padding: 0;
      margin: 0;
      
      li {
        padding: var(--spacing-sm) 0;
        border-bottom: 1px solid var(--border-color-light);
        font-size: var(--font-size-medium);
        color: var(--text-secondary);
        
        &:last-child {
          border-bottom: none;
        }
        
        &::before {
          content: '•';
          color: var(--primary-color);
          font-weight: bold;
          margin-right: var(--spacing-sm);
        }
      }
    }
  }
  
  .detail-actions {
    display: flex;
    gap: var(--spacing-md);
    justify-content: center;
    margin-top: var(--spacing-xl);
  }
}

// 响应式设计
@media (max-width: 768px) {
  .circles-page {
    padding: var(--spacing-lg);
  }
  
  .search-filter-section {
    flex-direction: column;
    align-items: stretch;
    
    .search-box {
      max-width: none;
    }
    
    .filter-options {
      justify-content: space-between;
    }
  }
  
  .circles-grid {
    grid-template-columns: 1fr;
  }
  
  .circle-detail {
    .detail-header {
      flex-direction: column;
      text-align: center;
      
      .detail-cover {
        width: 100%;
        height: 200px;
      }
    }
  }
}
</style> 