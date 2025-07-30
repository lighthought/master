<template>
  <div class="mentors-page">
    <!-- 搜索和筛选区域 -->
    <div class="search-section">
      <div class="search-container">
        <!-- 搜索框 -->
        <div class="search-box">
          <el-input
            v-model="searchKeyword"
            placeholder="搜索大师姓名、技能、领域"
            class="search-input"
            clearable
            @input="handleSearch"
            @clear="handleSearch"
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>
        </div>
        
        <!-- 筛选按钮 -->
        <el-button 
          type="primary" 
          @click="showFilterDrawer = true"
          class="filter-button"
        >
          <el-icon><Filter /></el-icon>
          筛选
          <el-badge v-if="activeFilterCount > 0" :value="activeFilterCount" class="filter-badge" />
        </el-button>
      </div>
      
      <!-- 快速筛选标签 -->
      <div v-if="quickFilters.length > 0" class="quick-filters">
        <el-tag
          v-for="filter in quickFilters"
          :key="filter.key"
          closable
          @close="removeQuickFilter(filter.key)"
          class="quick-filter-tag"
        >
          {{ filter.label }}
        </el-tag>
      </div>
    </div>
    
    <!-- 搜索结果 -->
    <div class="results-section">
      <!-- 结果统计 -->
      <div class="results-header">
        <div class="results-info">
          <span class="results-count">找到 {{ totalCount }} 位大师</span>
          <span v-if="searchKeyword" class="search-keyword">关键词："{{ searchKeyword }}"</span>
        </div>
        <div class="sort-options">
          <el-select v-model="sortBy" placeholder="排序方式" @change="handleSearch">
            <el-option label="推荐排序" value="recommended" />
            <el-option label="评分最高" value="rating" />
            <el-option label="价格最低" value="price-asc" />
            <el-option label="价格最高" value="price-desc" />
            <el-option label="学生最多" value="students" />
          </el-select>
        </div>
      </div>
      
      <!-- 加载状态 -->
      <div v-if="loading" class="loading-container">
        <el-skeleton :rows="3" animated />
        <div class="skeleton-cards">
          <el-skeleton-item v-for="i in 6" :key="i" variant="card" style="width: 100%; height: 200px; margin-bottom: 16px;" />
        </div>
      </div>
      
      <!-- 大师列表 -->
      <div v-else-if="mentors.length > 0" class="mentors-grid">
        <div 
          v-for="mentor in mentors" 
          :key="mentor.id"
          class="mentor-card"
          @click="viewMentorDetail(mentor)"
        >
          <!-- 大师头像和状态 -->
          <div class="mentor-header">
            <div class="avatar-container">
              <el-avatar 
                :size="80" 
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
              v-for="skill in mentor.skills.slice(0, 4)" 
              :key="skill"
              size="small"
              class="skill-tag"
            >
              {{ skill }}
            </el-tag>
            <span v-if="mentor.skills.length > 4" class="more-skills">
              +{{ mentor.skills.length - 4 }}
            </span>
          </div>
          
          <!-- 个人介绍 -->
          <div class="mentor-bio">
            <p>{{ mentor.bio }}</p>
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
      
      <!-- 空状态 -->
      <div v-else class="empty-state">
        <el-empty description="没有找到符合条件的大师">
          <el-button type="primary" @click="clearFilters">
            清除筛选条件
          </el-button>
        </el-empty>
      </div>
      
      <!-- 分页 -->
      <div v-if="totalPages > 1" class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50]"
          :total="totalCount"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handlePageSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </div>
    
    <!-- 筛选抽屉 -->
    <el-drawer
      v-model="showFilterDrawer"
      title="筛选条件"
      direction="rtl"
      size="300px"
    >
      <div class="filter-content">
        <!-- 领域筛选 -->
        <div class="filter-group">
          <h4 class="filter-title">专业领域</h4>
          <el-select
            v-model="filters.domain"
            placeholder="选择领域"
            clearable
            style="width: 100%"
            @change="handleSearch"
          >
            <el-option
              v-for="domain in domainOptions"
              :key="domain.value"
              :label="domain.label"
              :value="domain.value"
            />
          </el-select>
        </div>
        
        <!-- 价格范围 -->
        <div class="filter-group">
          <h4 class="filter-title">价格范围</h4>
          <div class="price-range">
            <el-input-number
              v-model="filters.minPrice"
              placeholder="最低价格"
              :min="0"
              :max="filters.maxPrice || 1000"
              style="width: 100%"
              @change="handleSearch"
            />
            <span class="price-separator">-</span>
            <el-input-number
              v-model="filters.maxPrice"
              placeholder="最高价格"
              :min="filters.minPrice || 0"
              :max="1000"
              style="width: 100%"
              @change="handleSearch"
            />
          </div>
        </div>
        
        <!-- 评分筛选 -->
        <div class="filter-group">
          <h4 class="filter-title">最低评分</h4>
          <el-rate
            v-model="filters.minRating"
            :max="5"
            :low-threshold="3"
            :high-threshold="4"
            @change="handleSearch"
          />
          <span class="rating-text">{{ filters.minRating }}分以上</span>
        </div>
        
        <!-- 在线状态 -->
        <div class="filter-group">
          <h4 class="filter-title">在线状态</h4>
          <el-switch
            v-model="filters.isOnline"
            active-text="仅显示在线大师"
            @change="handleSearch"
          />
        </div>
        
        <!-- 认证状态 -->
        <div class="filter-group">
          <h4 class="filter-title">认证状态</h4>
          <el-checkbox
            v-model="filters.isVerified"
            @change="handleSearch"
          >
            仅显示已认证大师
          </el-checkbox>
        </div>
        
        <!-- 技能筛选 -->
        <div class="filter-group">
          <h4 class="filter-title">技能标签</h4>
          <el-select
            v-model="filters.skills"
            multiple
            placeholder="选择技能"
            style="width: 100%"
            @change="handleSearch"
          >
            <el-option
              v-for="skill in skillOptions"
              :key="skill.value"
              :label="skill.label"
              :value="skill.value"
            />
          </el-select>
        </div>
        
        <!-- 操作按钮 -->
        <div class="filter-actions">
          <el-button @click="clearFilters" style="width: 100%; margin-bottom: 8px;">
            清除所有筛选
          </el-button>
          <el-button type="primary" @click="showFilterDrawer = false" style="width: 100%;">
            应用筛选
          </el-button>
        </div>
      </div>
    </el-drawer>
    
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
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  Search, Filter, User, Check 
} from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'

// 路由
const router = useRouter()

// 认证store
const authStore = useAuthStore()

// 状态
const loading = ref(false)
const showFilterDrawer = ref(false)
const showBookingDialog = ref(false)
const bookingLoading = ref(false)
const selectedMentor = ref<any>(null)

// 搜索和筛选
const searchKeyword = ref('')
const currentPage = ref(1)
const pageSize = ref(10)
const sortBy = ref('recommended')

// 筛选条件
const filters = ref({
  domain: '',
  minPrice: null as number | null,
  maxPrice: null as number | null,
  minRating: 0,
  isOnline: false,
  isVerified: false,
  skills: [] as string[]
})

// 搜索结果
const mentors = ref<any[]>([])
const totalCount = ref(0)
const totalPages = ref(0)

// 选项数据
const domainOptions = [
  { label: '前端开发', value: '前端开发' },
  { label: '后端开发', value: '后端开发' },
  { label: '移动开发', value: '移动开发' },
  { label: '人工智能', value: '人工智能' },
  { label: '系统架构', value: '系统架构' },
  { label: 'UI/UX设计', value: 'UI/UX设计' }
]

const skillOptions = [
  { label: 'Vue.js', value: 'Vue.js' },
  { label: 'React', value: 'React' },
  { label: 'TypeScript', value: 'TypeScript' },
  { label: 'Java', value: 'Java' },
  { label: 'Spring Boot', value: 'Spring Boot' },
  { label: 'Flutter', value: 'Flutter' },
  { label: 'Python', value: 'Python' },
  { label: 'TensorFlow', value: 'TensorFlow' },
  { label: '微服务', value: '微服务' },
  { label: 'Figma', value: 'Figma' }
]

// 计算属性
const activeFilterCount = computed(() => {
  let count = 0
  if (filters.value.domain) count++
  if (filters.value.minPrice !== null) count++
  if (filters.value.maxPrice !== null) count++
  if (filters.value.minRating > 0) count++
  if (filters.value.isOnline) count++
  if (filters.value.isVerified) count++
  if (filters.value.skills.length > 0) count++
  return count
})

const quickFilters = computed(() => {
  const quick: any[] = []
  
  if (filters.value.domain) {
    quick.push({ key: 'domain', label: `领域：${filters.value.domain}` })
  }
  
  if (filters.value.minPrice !== null || filters.value.maxPrice !== null) {
    const min = filters.value.minPrice || 0
    const max = filters.value.maxPrice || '不限'
    quick.push({ key: 'price', label: `价格：¥${min}-${max}` })
  }
  
  if (filters.value.minRating > 0) {
    quick.push({ key: 'rating', label: `评分：${filters.value.minRating}分以上` })
  }
  
  if (filters.value.isOnline) {
    quick.push({ key: 'online', label: '仅在线' })
  }
  
  if (filters.value.isVerified) {
    quick.push({ key: 'verified', label: '已认证' })
  }
  
  return quick
})

// 预约表单
const bookingForm = ref({
  date: '',
  timeSlot: '',
  method: 'video',
  requirements: ''
})

// 搜索大师
const searchMentors = async () => {
  if (!authStore.user) return
  
  loading.value = true
  try {
    const params = {
      keyword: searchKeyword.value,
      domain: filters.value.domain,
      minPrice: filters.value.minPrice,
      maxPrice: filters.value.maxPrice,
      minRating: filters.value.minRating,
      isOnline: filters.value.isOnline,
      isVerified: filters.value.isVerified,
      skills: filters.value.skills,
      sortBy: sortBy.value,
      page: currentPage.value,
      pageSize: pageSize.value
    }
    
    const result = await ApiService.mentors.searchMentors(params)
    mentors.value = result.data.mentors
    totalCount.value = result.data.total
    totalPages.value = result.data.totalPages
  } catch (error) {
    console.error('搜索大师失败:', error)
    ElMessage.error('搜索失败，请重试')
  } finally {
    loading.value = false
  }
}

// 处理搜索
const handleSearch = () => {
  currentPage.value = 1
  searchMentors()
}

// 处理分页
const handlePageChange = (page: number) => {
  currentPage.value = page
  searchMentors()
}

const handlePageSizeChange = (size: number) => {
  pageSize.value = size
  currentPage.value = 1
  searchMentors()
}

// 移除快速筛选
const removeQuickFilter = (key: string) => {
  switch (key) {
    case 'domain':
      filters.value.domain = ''
      break
    case 'price':
      filters.value.minPrice = null
      filters.value.maxPrice = null
      break
    case 'rating':
      filters.value.minRating = 0
      break
    case 'online':
      filters.value.isOnline = false
      break
    case 'verified':
      filters.value.isVerified = false
      break
  }
  handleSearch()
}

// 清除所有筛选
const clearFilters = () => {
  filters.value = {
    domain: '',
    minPrice: null,
    maxPrice: null,
    minRating: 0,
    isOnline: false,
    isVerified: false,
    skills: []
  }
  searchKeyword.value = ''
  sortBy.value = 'recommended'
  currentPage.value = 1
  searchMentors()
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

// 监听筛选条件变化
watch(filters, () => {
  handleSearch()
}, { deep: true })

// 组件挂载时加载数据
onMounted(() => {
  searchMentors()
})
</script>

<style scoped lang="scss">
.mentors-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: var(--spacing-xl);
}

.search-section {
  margin-bottom: var(--spacing-xl);
}

.search-container {
  display: flex;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-md);
}

.search-box {
  flex: 1;
}

.search-input {
  :deep(.el-input__wrapper) {
    background: #2D2D2D;
    border-radius: var(--border-radius-medium);
    border: 1px solid var(--border-color-light);
    
    &:hover {
      border-color: var(--primary-color);
    }
    
    &.is-focus {
      border-color: var(--primary-color);
      box-shadow: 0 0 0 1px var(--primary-color);
    }
  }
  
  :deep(.el-input__prefix) {
    color: var(--text-secondary);
  }
}

.filter-button {
  position: relative;
  min-width: 80px;
}

.filter-badge {
  position: absolute;
  top: -8px;
  right: -8px;
}

.quick-filters {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
}

.quick-filter-tag {
  background: rgba(64, 158, 255, 0.1);
  border-color: var(--primary-color);
  color: var(--primary-color);
}

.results-section {
  .results-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: var(--spacing-lg);
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
  }
  
  .results-info {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-xs);
  }
  
  .results-count {
    font-size: var(--font-size-medium);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
  }
  
  .search-keyword {
    font-size: var(--font-size-small);
    color: var(--text-secondary);
  }
}

.loading-container {
  .skeleton-cards {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-md);
  }
}

.mentors-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-xl);
}

.mentor-card {
  background: var(--bg-secondary);
  border: 1px solid var(--border-color-light);
  border-radius: var(--border-radius-medium);
  padding: var(--spacing-lg);
  cursor: pointer;
  transition: all var(--transition-normal);
  
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
  width: 16px;
  height: 16px;
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
  font-size: 18px;
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
}

.verified-badge {
  color: #67c23a;
  font-size: 16px;
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

.mentor-bio {
  margin-bottom: var(--spacing-md);
  
  p {
    font-size: 14px;
    color: var(--text-secondary);
    line-height: 1.5;
    margin: 0;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }
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
  font-size: 18px;
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

.pagination-container {
  display: flex;
  justify-content: center;
  margin-top: var(--spacing-xl);
}

.filter-content {
  padding: var(--spacing-md);
}

.filter-group {
  margin-bottom: var(--spacing-lg);
  
  .filter-title {
    font-size: var(--font-size-medium);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-sm) 0;
  }
}

.price-range {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  
  .price-separator {
    color: var(--text-secondary);
    font-weight: var(--font-weight-medium);
  }
}

.rating-text {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  margin-left: var(--spacing-sm);
}

.filter-actions {
  margin-top: var(--spacing-xl);
  padding-top: var(--spacing-lg);
  border-top: 1px solid var(--border-color-light);
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
  .mentors-page {
    padding: var(--spacing-lg);
  }
  
  .search-container {
    flex-direction: column;
  }
  
  .results-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-md);
  }
  
  .mentors-grid {
    grid-template-columns: 1fr;
    gap: var(--spacing-md);
  }
  
  .mentor-card {
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