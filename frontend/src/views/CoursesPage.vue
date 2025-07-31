<template>
  <div class="courses-page">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1 class="page-title">课程中心</h1>
      <p class="page-subtitle">发现优质课程，提升专业技能</p>
    </div>
    
    <!-- 搜索和筛选区域 -->
    <div class="search-filter-area">
      <div class="search-section">
        <el-input
          v-model="searchQuery"
          placeholder="搜索课程名称、大师姓名、技能关键词"
          clearable
          class="search-input"
          @keyup.enter="handleSearch"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-button type="primary" :icon="Filter" @click="showFilterDrawer = true" class="filter-button">
          筛选
        </el-button>
      </div>
      
      <!-- 课程分类标签 -->
      <div class="category-tags">
        <el-tag
          v-for="category in categories"
          :key="category.value"
          :type="selectedCategory === category.value ? 'primary' : 'info'"
          effect="light"
          @click="selectCategory(category.value)"
          class="category-tag"
        >
          {{ category.label }}
        </el-tag>
      </div>
    </div>
    
    <!-- 活跃筛选条件 -->
    <div v-if="activeFilters.length > 0" class="active-filters">
      <el-tag
        v-for="(filter, index) in activeFilters"
        :key="index"
        closable
        @close="removeFilter(filter.type, filter.value)"
        type="info"
        effect="plain"
        class="filter-tag"
      >
        {{ filter.label }}: {{ filter.displayValue }}
      </el-tag>
      <el-button type="text" size="small" @click="clearAllFilters">清除所有</el-button>
    </div>
    
    <!-- 排序选项 -->
    <div class="sort-options">
      <el-radio-group v-model="sortOption" size="small" @change="handleSearch">
        <el-radio-button label="recommended">推荐排序</el-radio-button>
        <el-radio-button label="latest">最新发布</el-radio-button>
        <el-radio-button label="popular">最受欢迎</el-radio-button>
        <el-radio-button label="rating">评分最高</el-radio-button>
        <el-radio-button label="price_asc">价格从低到高</el-radio-button>
        <el-radio-button label="price_desc">价格从高到低</el-radio-button>
      </el-radio-group>
    </div>
    
    <!-- 推荐课程区域 -->
    <div v-if="!searchQuery && !hasActiveFilters" class="recommended-section">
      <h2 class="section-title">推荐课程</h2>
      <div class="recommended-courses">
        <div 
          v-for="course in recommendedCourses" 
          :key="course.id"
          class="course-card featured"
          @click="viewCourseDetail(course.id)"
        >
          <div class="course-cover">
            <img :src="course.cover" :alt="course.title" />
            <div class="course-badge">推荐</div>
          </div>
          <div class="course-content">
            <h3 class="course-title">{{ course.title }}</h3>
            <p class="course-description">{{ course.description }}</p>
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
              <el-button type="primary" size="small" @click.stop="enrollCourse(course)">
                立即报名
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 课程列表 -->
    <div class="courses-section">
      <div class="section-header">
        <h2 class="section-title">
          {{ searchQuery ? '搜索结果' : '全部课程' }}
          <span class="course-count">({{ totalCourses }} 门课程)</span>
        </h2>
      </div>
      
      <!-- 加载状态 -->
      <div v-if="loading" class="loading-container">
        <el-skeleton :rows="3" animated />
        <div class="skeleton-cards">
          <el-skeleton-item v-for="i in 6" :key="i" variant="image" style="width: 100%; height: 300px; margin-bottom: 16px;" />
        </div>
      </div>
      
      <!-- 课程网格 -->
      <div v-else-if="courses.length > 0" class="courses-grid">
        <div 
          v-for="course in courses" 
          :key="course.id"
          class="course-card"
          @click="viewCourseDetail(course.id)"
        >
          <div class="course-cover">
            <img :src="course.cover" :alt="course.title" />
            <div class="course-badge" v-if="course.isNew">新课程</div>
            <div class="course-badge hot" v-if="course.isHot">热门</div>
          </div>
          <div class="course-content">
            <h3 class="course-title">{{ course.title }}</h3>
            <p class="course-description">{{ course.description }}</p>
            <div class="course-tags">
              <el-tag v-for="tag in course.tags.slice(0, 3)" :key="tag" size="small" effect="light">
                {{ tag }}
              </el-tag>
            </div>
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
              <el-button type="primary" size="small" @click.stop="enrollCourse(course)">
                立即报名
              </el-button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 空状态 -->
      <div v-else class="empty-state">
        <el-empty :description="getEmptyText()">
          <el-button type="primary" @click="clearAllFilters">
            清除筛选条件
          </el-button>
        </el-empty>
      </div>
      
      <!-- 分页 -->
      <div v-if="totalCourses > pageSize" class="pagination-container">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="totalCourses"
          :page-size="pageSize"
          v-model:current-page="currentPage"
          @current-change="handlePageChange"
        />
      </div>
    </div>
    
    <!-- 筛选抽屉 -->
    <el-drawer
      v-model="showFilterDrawer"
      title="筛选课程"
      direction="rtl"
      size="80%"
      :close-on-click-modal="false"
    >
      <el-form :model="filterForm" label-position="top">
        <el-form-item label="课程分类">
          <el-select v-model="filterForm.category" placeholder="选择分类" clearable style="width: 100%;">
            <el-option 
              v-for="category in categories" 
              :key="category.value" 
              :label="category.label" 
              :value="category.value" 
            />
          </el-select>
        </el-form-item>
        
        <el-form-item label="价格范围 (¥)">
          <el-input-number v-model="filterForm.minPrice" :min="0" :max="10000" placeholder="最低" style="width: 48%;" />
          <span style="margin: 0 5px;">-</span>
          <el-input-number v-model="filterForm.maxPrice" :min="0" :max="10000" placeholder="最高" style="width: 48%;" />
        </el-form-item>
        
        <el-form-item label="最低评分">
          <el-rate v-model="filterForm.minRating" :max="5" allow-half />
        </el-form-item>
        
        <el-form-item label="课程时长">
          <el-radio-group v-model="filterForm.duration">
            <el-radio label="">不限</el-radio>
            <el-radio label="short">1-5小时</el-radio>
            <el-radio label="medium">5-20小时</el-radio>
            <el-radio label="long">20小时以上</el-radio>
          </el-radio-group>
        </el-form-item>
        
        <el-form-item label="难度等级">
          <el-radio-group v-model="filterForm.level">
            <el-radio label="">不限</el-radio>
            <el-radio label="beginner">初级</el-radio>
            <el-radio label="intermediate">中级</el-radio>
            <el-radio label="advanced">高级</el-radio>
          </el-radio-group>
        </el-form-item>
        
        <el-form-item label="技能标签">
          <el-select
            v-model="filterForm.skills"
            multiple
            filterable
            allow-create
            default-first-option
            placeholder="输入或选择技能"
            style="width: 100%;"
          >
            <el-option label="Vue.js" value="Vue.js" />
            <el-option label="React" value="React" />
            <el-option label="Angular" value="Angular" />
            <el-option label="Java" value="Java" />
            <el-option label="Python" value="Python" />
            <el-option label="Node.js" value="Node.js" />
            <el-option label="Spring Boot" value="Spring Boot" />
            <el-option label="数据分析" value="数据分析" />
            <el-option label="机器学习" value="机器学习" />
            <el-option label="UI/UX设计" value="UI/UX设计" />
          </el-select>
        </el-form-item>
      </el-form>
      
      <template #footer>
        <div style="flex: auto">
          <el-button @click="clearFilterForm">重置</el-button>
          <el-button type="primary" @click="applyFilters">应用筛选</el-button>
        </div>
      </template>
    </el-drawer>
    
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
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Search, Filter } from '@element-plus/icons-vue'
import { ApiService } from '@/services/api'
import { useAuthStore } from '@/stores/auth'

// 路由
const router = useRouter()

// 认证store
const authStore = useAuthStore()

// 状态
const loading = ref(false)
const enrolling = ref(false)
const showFilterDrawer = ref(false)
const showEnrollDialog = ref(false)
const currentPage = ref(1)
const pageSize = 12
const searchQuery = ref('')
const selectedCategory = ref('')
const sortOption = ref('recommended')
const selectedCourse = ref<any>(null)

// 数据
const courses = ref<any[]>([])
const recommendedCourses = ref<any[]>([])
const totalCourses = ref(0)

// 筛选表单
const filterForm = ref({
  category: '',
  minPrice: undefined,
  maxPrice: undefined,
  minRating: 0,
  duration: '',
  level: '',
  skills: [] as string[]
})

// 课程分类
const categories = [
  { label: '全部', value: '' },
  { label: '前端开发', value: 'frontend' },
  { label: '后端开发', value: 'backend' },
  { label: '移动开发', value: 'mobile' },
  { label: '数据科学', value: 'data' },
  { label: '人工智能', value: 'ai' },
  { label: '产品设计', value: 'design' },
  { label: '项目管理', value: 'management' }
]

// 计算属性
const hasActiveFilters = computed(() => {
  return selectedCategory.value || 
         filterForm.value.category || 
         filterForm.value.minPrice !== undefined || 
         filterForm.value.maxPrice !== undefined ||
         filterForm.value.minRating > 0 ||
         filterForm.value.duration ||
         filterForm.value.level ||
         filterForm.value.skills.length > 0
})

const activeFilters = computed(() => {
  const filters = []
  if (searchQuery.value) {
    filters.push({ type: 'search', label: '关键词', value: searchQuery.value, displayValue: searchQuery.value })
  }
  if (selectedCategory.value) {
    const category = categories.find(c => c.value === selectedCategory.value)
    filters.push({ type: 'category', label: '分类', value: selectedCategory.value, displayValue: category?.label })
  }
  if (filterForm.value.category) {
    const category = categories.find(c => c.value === filterForm.value.category)
    filters.push({ type: 'filterCategory', label: '分类', value: filterForm.value.category, displayValue: category?.label })
  }
  if (filterForm.value.minPrice !== undefined || filterForm.value.maxPrice !== undefined) {
    const min = filterForm.value.minPrice !== undefined ? filterForm.value.minPrice : '不限'
    const max = filterForm.value.maxPrice !== undefined ? filterForm.value.maxPrice : '不限'
    filters.push({ type: 'price', label: '价格', value: `${min}-${max}`, displayValue: `¥${min}-${max}` })
  }
  if (filterForm.value.minRating > 0) {
    filters.push({ type: 'rating', label: '评分', value: filterForm.value.minRating, displayValue: `${filterForm.value.minRating}星及以上` })
  }
  if (filterForm.value.duration) {
    const durationTexts = { short: '1-5小时', medium: '5-20小时', long: '20小时以上' }
    filters.push({ type: 'duration', label: '时长', value: filterForm.value.duration, displayValue: durationTexts[filterForm.value.duration as keyof typeof durationTexts] })
  }
  if (filterForm.value.level) {
    const levelTexts = { beginner: '初级', intermediate: '中级', advanced: '高级' }
    filters.push({ type: 'level', label: '难度', value: filterForm.value.level, displayValue: levelTexts[filterForm.value.level as keyof typeof levelTexts] })
  }
  if (filterForm.value.skills.length > 0) {
    filters.push({ type: 'skills', label: '技能', value: filterForm.value.skills.join(','), displayValue: filterForm.value.skills.join(', ') })
  }
  return filters
})

// 加载课程列表
const loadCourses = async () => {
  loading.value = true
  try {
    const params = {
      query: searchQuery.value,
      ...filterForm.value,
      category: selectedCategory.value || filterForm.value.category,
      sort: sortOption.value,
      page: currentPage.value,
      pageSize: pageSize
    }
    const response = await ApiService.courses.searchCourses(params)
    courses.value = response.data.courses
    totalCourses.value = response.data.total
  } catch (error) {
    console.error('加载课程列表失败:', error)
    ElMessage.error('加载课程列表失败')
  } finally {
    loading.value = false
  }
}

// 加载推荐课程
const loadRecommendedCourses = async () => {
  try {
    const response = await ApiService.courses.getRecommendedCourses()
    recommendedCourses.value = response.data.slice(0, 4) // 只显示4个推荐课程
  } catch (error) {
    console.error('加载推荐课程失败:', error)
  }
}

// 处理搜索
const handleSearch = () => {
  currentPage.value = 1
  loadCourses()
}

// 选择分类
const selectCategory = (category: string) => {
  selectedCategory.value = category
  handleSearch()
}

// 应用筛选
const applyFilters = () => {
  showFilterDrawer.value = false
  handleSearch()
}

// 清除筛选表单
const clearFilterForm = () => {
  filterForm.value = {
    category: '',
    minPrice: undefined,
    maxPrice: undefined,
    minRating: 0,
    duration: '',
    level: '',
    skills: []
  }
}

// 移除筛选条件
const removeFilter = (type: string, value: any) => {
  if (type === 'search') {
    searchQuery.value = ''
  } else if (type === 'category') {
    selectedCategory.value = ''
  } else if (type === 'filterCategory') {
    filterForm.value.category = ''
  } else if (type === 'price') {
    filterForm.value.minPrice = undefined
    filterForm.value.maxPrice = undefined
  } else if (type === 'skills') {
    filterForm.value.skills = []
  } else {
    // 对于其他单选筛选条件
    (filterForm.value as any)[type] = (typeof (filterForm.value as any)[type] === 'boolean') ? false : (typeof (filterForm.value as any)[type] === 'number' ? 0 : '')
  }
  handleSearch()
}

// 清除所有筛选条件
const clearAllFilters = () => {
  searchQuery.value = ''
  selectedCategory.value = ''
  clearFilterForm()
  handleSearch()
}

// 处理分页
const handlePageChange = (page: number) => {
  currentPage.value = page
  loadCourses()
}

// 查看课程详情
const viewCourseDetail = (courseId: string) => {
  router.push(`/courses/${courseId}`)
}

// 报名课程
const enrollCourse = (course: any) => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录才能报名课程')
    router.push('/auth')
    return
  }
  
  selectedCourse.value = course
  showEnrollDialog.value = true
}

// 确认报名
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

// 获取空状态文本
const getEmptyText = () => {
  if (searchQuery.value) {
    return `没有找到包含"${searchQuery.value}"的课程`
  }
  if (hasActiveFilters.value) {
    return '没有找到符合筛选条件的课程'
  }
  return '暂无课程'
}

// 组件挂载时加载数据
onMounted(() => {
  loadRecommendedCourses()
  loadCourses()
})
</script>

<style scoped lang="scss">
.courses-page {
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

.search-filter-area {
  margin-bottom: var(--spacing-xl);
}

.search-section {
  display: flex;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-lg);
}

.search-input {
  flex-grow: 1;
  
  :deep(.el-input__wrapper) {
    background-color: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
    box-shadow: none;
    
    &.is-focus {
      box-shadow: 0 0 0 1px var(--primary-color) inset;
    }
  }
  
  :deep(.el-input__inner) {
    color: var(--text-primary);
  }
  
  :deep(.el-input__prefix) {
    color: var(--text-secondary);
  }
}

.filter-button {
  background-color: var(--primary-color);
  color: white;
  border-radius: var(--border-radius-medium);
}

.category-tags {
  display: flex;
  flex-wrap: wrap;
  gap: var(--spacing-sm);
}

.category-tag {
  cursor: pointer;
  transition: all var(--transition-normal);
  
  &:hover {
    transform: translateY(-1px);
  }
}

.active-filters {
  display: flex;
  flex-wrap: wrap;
  gap: var(--spacing-xs);
  margin-bottom: var(--spacing-lg);
}

.filter-tag {
  background-color: var(--bg-tertiary);
  color: var(--text-tertiary);
  border: 1px solid var(--border-color-light);
}

.sort-options {
  margin-bottom: var(--spacing-xl);
  text-align: center;
  
  .el-radio-button {
    :deep(.el-radio-button__inner) {
      background-color: var(--bg-secondary);
      color: var(--text-secondary);
      border: 1px solid var(--border-color-light);
      border-left: none;
      
      &:hover {
        color: var(--primary-color);
      }
    }
    
    &.is-active {
      :deep(.el-radio-button__inner) {
        background-color: var(--primary-color);
        border-color: var(--primary-color);
        color: white;
        box-shadow: -1px 0 0 0 var(--primary-color);
      }
    }
    
    &:first-child {
      :deep(.el-radio-button__inner) {
        border-left: 1px solid var(--border-color-light);
      }
    }
  }
}

.recommended-section {
  margin-bottom: var(--spacing-xxl);
}

.section-title {
  font-size: var(--font-size-h3);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin: 0 0 var(--spacing-lg) 0;
}

.course-count {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  font-weight: var(--font-weight-normal);
}

.recommended-courses {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-lg);
}

.courses-section {
  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: var(--spacing-lg);
  }
}

.loading-container {
  .skeleton-cards {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: var(--spacing-lg);
  }
}

.courses-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-xl);
}

.course-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-medium);
  overflow: hidden;
  box-shadow: var(--shadow-light);
  transition: all var(--transition-normal);
  cursor: pointer;
  
  &:hover {
    transform: translateY(-5px);
    box-shadow: var(--shadow-medium);
  }
  
  &.featured {
    border: 2px solid var(--primary-color);
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
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.course-tags {
  display: flex;
  flex-wrap: wrap;
  gap: var(--spacing-xs);
  margin-bottom: var(--spacing-md);
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
}

.course-price {
  display: flex;
  align-items: baseline;
  gap: var(--spacing-xs);
  
  .price-value {
    font-size: var(--font-size-h4);
    font-weight: var(--font-weight-bold);
    color: #f56c6c;
  }
  
  .price-original {
    font-size: var(--font-size-small);
    color: var(--text-tertiary);
    text-decoration: line-through;
  }
}

.empty-state {
  text-align: center;
  padding: var(--spacing-xxl) 0;
}

.pagination-container {
  display: flex;
  justify-content: center;
  margin-top: var(--spacing-xl);
}

// 抽屉样式
.el-drawer {
  .el-drawer__header {
    margin-bottom: var(--spacing-lg);
  }
  .el-drawer__body {
    padding: var(--spacing-lg);
  }
  .el-form-item {
    margin-bottom: var(--spacing-lg);
  }
  .el-input-number {
    width: 100%;
  }
  .el-rate {
    --el-rate-fill-color: var(--primary-color);
  }
  .el-radio-group {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-sm);
  }
  .el-select {
    width: 100%;
  }
  .el-drawer__footer {
    padding: var(--spacing-lg);
    border-top: 1px solid var(--border-color-light);
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

// 响应式设计
@media (max-width: 768px) {
  .courses-page {
    padding: var(--spacing-lg);
  }
  
  .search-section {
    flex-direction: column;
  }
  
  .filter-button {
    width: 100%;
  }
  
  .sort-options {
    .el-radio-group {
      display: flex;
      flex-wrap: wrap;
      .el-radio-button {
        flex: 1 0 50%;
        :deep(.el-radio-button__inner) {
          width: 100%;
          border-left: 1px solid var(--border-color-light) !important;
          border-top: 1px solid var(--border-color-light);
          &:first-child {
            border-top-left-radius: var(--border-radius-medium);
            border-bottom-left-radius: 0;
          }
          &:last-child {
            border-top-right-radius: var(--border-radius-medium);
            border-bottom-right-radius: 0;
          }
        }
      }
    }
  }
  
  .recommended-courses,
  .courses-grid {
    grid-template-columns: 1fr;
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