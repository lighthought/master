<template>
  <div class="recommended-circles">
    <div class="section-header">
      <h2 class="section-title">推荐圈子</h2>
      <el-button type="primary" @click="$router.push('/circles')">
          <el-icon><ArrowRight /></el-icon>
          查看更多
        </el-button>
    </div>

    <div class="circles-grid">
    <div class="circles-scroll">
      <div
        v-for="circle in recommendedCircles"
        :key="circle.id"
        class="circle-card"
        @click="viewCircle(circle)"
      >
        <div class="circle-cover">
          <img :src="circle.cover" :alt="circle.name" />
          <div class="circle-status">
            <el-tag v-if="circle.isActive" type="success" size="small">活跃</el-tag>
            <el-tag v-if="circle.isJoined" type="warning" size="small">已加入</el-tag>
          </div>
        </div>

        <div class="circle-info">
          <h3 class="circle-name">{{ circle.name }}</h3>
          <p class="circle-description">{{ circle.description }}</p>

          <div class="circle-tags">
            <el-tag
              v-for="tag in circle.tags.slice(0, 2)"
              :key="tag"
              size="small"
              effect="light"
              class="tag-item"
            >
              {{ tag }}
            </el-tag>
            <span v-if="circle.tags.length > 2" class="more-tags">
              +{{ circle.tags.length - 2 }}
            </span>
          </div>

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
          </div>
        </div>
      </div>
    </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ArrowRight, User, ChatDotRound } from '@element-plus/icons-vue'
import { ApiService } from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const recommendedCircles = ref<any[]>([])
const joiningCircleId = ref('')
const leavingCircleId = ref('')

// 加载推荐圈子
const loadRecommendedCircles = async () => {
  try {
    const response = await ApiService.circles.getCircles({
      page: 1,
      pageSize: 4,
      sort: 'memberCount'
    })
    recommendedCircles.value = response.data.circles
  } catch (error) {
    console.error('加载推荐圈子失败:', error)
  }
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
      '确定要退出该圈子吗？',
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
  } catch (error) {
    if (error !== 'cancel') {
      console.error('退出圈子失败:', error)
      ElMessage.error('退出圈子失败')
    }
  } finally {
    leavingCircleId.value = ''
  }
}

// 查看圈子
const viewCircle = (circle: any) => {
  router.push(`/circles`)
}

// 格式化数字
const formatNumber = (num: number) => {
  if (num >= 10000) {
    return (num / 10000).toFixed(1) + '万'
  }
  return num.toString()
}

onMounted(() => {
  loadRecommendedCircles()
})
</script>

<style scoped lang="scss">
.recommended-circles {
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
  
  .section-title {
    font-size: var(--font-size-h3);
    font-weight: var(--font-weight-bold);
    color: var(--text-primary);
    margin: 0;
  }
  
  .view-all-btn {
    font-size: var(--font-size-medium);
    color: var(--primary-color);
    
    .el-icon {
      margin-left: var(--spacing-xs);
    }
  }
}

.circles-scroll {
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

.circles-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: var(--spacing-lg);
}

.circle-card {
  flex-shrink: 0;
  width: 280px;
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  padding: var(--spacing-lg);
  cursor: pointer;
  transition: all var(--transition-normal);
  border: 1px solid transparent;

  &:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-light);
    border-color: var(--primary-color);
  }
  
  .circle-cover {
    position: relative;
    height: 140px;
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
    padding: var(--spacing-md);
    
    .circle-name {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-sm) 0;
      line-height: 1.3;
    }
    
    .circle-description {
      font-size: var(--font-size-small);
      color: var(--text-secondary);
      margin: 0 0 var(--spacing-sm) 0;
      line-height: 1.4;
      display: -webkit-box;
      -webkit-line-clamp: 2;
      -webkit-box-orient: vertical;
      overflow: hidden;
    }
    
    .circle-tags {
      display: flex;
      flex-wrap: wrap;
      gap: var(--spacing-xs);
      margin-bottom: var(--spacing-sm);
      
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
      gap: var(--spacing-md);
      margin-bottom: var(--spacing-sm);
      
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
      justify-content: center;
    }
  }
}

// 响应式设计
@media (max-width: 768px) {
  .circles-grid {
    grid-template-columns: 1fr;
  }
  
  .section-header {
    flex-direction: column;
    gap: var(--spacing-sm);
    align-items: flex-start;
  }
}
</style>