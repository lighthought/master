<template>
  <div class="recommended-posts">
    <div class="section-header">
      <h2 class="section-title">最新动态</h2>
      <el-button type="text" @click="$router.push('/posts')" class="view-all-btn">
        查看全部
        <el-icon><ArrowRight /></el-icon>
      </el-button>
    </div>

    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="2" animated />
      <el-skeleton :rows="2" animated />
    </div>

    <div v-else-if="recommendedPosts.length > 0" class="posts-preview">
      <div 
        v-for="post in recommendedPosts" 
        :key="post.id"
        class="post-preview-card"
        @click="viewPost(post)"
      >
        <div class="post-header">
          <el-avatar :size="32" :src="post.userAvatar" />
          <div class="post-info">
            <div class="user-name">{{ post.userName }}</div>
            <div class="post-meta">
              <el-tag 
                :type="post.userRole === 'master' ? 'success' : 'warning'" 
                size="small"
              >
                {{ post.userRole === 'master' ? '大师' : '学徒' }}
              </el-tag>
              <span class="post-time">{{ formatTime(post.createdAt) }}</span>
            </div>
          </div>
        </div>

        <div class="post-content">
          <p class="post-text">{{ truncateText(post.content, 80) }}</p>
          
          <div v-if="post.images && post.images.length > 0" class="post-images">
            <img 
              :src="post.images[0]" 
              :alt="post.content"
              class="preview-image"
            />
            <div v-if="post.images.length > 1" class="more-images">
              +{{ post.images.length - 1 }}
            </div>
          </div>
        </div>

        <div class="post-stats">
          <span class="stat-item">
            <el-icon><Star /></el-icon>
            {{ post.likes }}
          </span>
          <span class="stat-item">
            <el-icon><ChatDotRound /></el-icon>
            {{ post.comments }}
          </span>
        </div>
      </div>
    </div>

    <div v-else class="empty-state">
      <el-empty description="暂无动态">
        <el-button type="primary" @click="$router.push('/posts')">
          发布第一条动态
        </el-button>
      </el-empty>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ArrowRight, Star, ChatDotRound } from '@element-plus/icons-vue'
import { ApiService } from '@/services/api'

const router = useRouter()

const recommendedPosts = ref<any[]>([])
const loading = ref(true)

// 加载推荐动态
const loadRecommendedPosts = async () => {
  loading.value = true
  try {
    const response = await ApiService.posts.getPosts({
      page: 1,
      pageSize: 3,
      sort: 'latest'
    })
    recommendedPosts.value = response.data.posts
  } catch (error) {
    console.error('加载推荐动态失败:', error)
  } finally {
    loading.value = false
  }
}

// 格式化时间
const formatTime = (time: string) => {
  const now = new Date()
  const postTime = new Date(time)
  const diff = now.getTime() - postTime.getTime()
  
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
    return postTime.toLocaleDateString()
  }
}

// 截断文本
const truncateText = (text: string, maxLength: number) => {
  if (text.length <= maxLength) return text
  return text.substring(0, maxLength) + '...'
}

// 查看动态
const viewPost = (post: any) => {
  router.push('/posts')
}

onMounted(() => {
  loadRecommendedPosts()
})
</script>

<style scoped lang="scss">
.recommended-posts {
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

.loading-container {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.posts-preview {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.post-preview-card {
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  padding: var(--spacing-md);
  cursor: pointer;
  transition: all var(--transition-normal);
  border: 1px solid transparent;
  
  &:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-light);
    border-color: var(--primary-color);
  }
}

.post-header {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-sm);
}

.post-info {
  flex: 1;
  
  .user-name {
    font-size: var(--font-size-small);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin-bottom: var(--spacing-xs);
  }
  
  .post-meta {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
    
    .post-time {
      font-size: var(--font-size-small);
      color: var(--text-tertiary);
    }
  }
}

.post-content {
  margin-bottom: var(--spacing-sm);
  
  .post-text {
    font-size: var(--font-size-small);
    color: var(--text-secondary);
    line-height: 1.4;
    margin: 0 0 var(--spacing-sm) 0;
  }
  
  .post-images {
    position: relative;
    width: 100%;
    height: 120px;
    border-radius: var(--border-radius-small);
    overflow: hidden;
    
    .preview-image {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
    
    .more-images {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: rgba(0, 0, 0, 0.7);
      color: white;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: var(--font-size-small);
      font-weight: var(--font-weight-semibold);
    }
  }
}

.post-stats {
  display: flex;
  gap: var(--spacing-md);
  
  .stat-item {
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
    font-size: var(--font-size-small);
    color: var(--text-tertiary);
    
    .el-icon {
      font-size: var(--font-size-medium);
    }
  }
}

.empty-state {
  text-align: center;
  padding: var(--spacing-xl) 0;
}

// 响应式设计
@media (max-width: 768px) {
  .recommended-posts {
    padding: var(--spacing-lg);
  }
  
  .section-header {
    flex-direction: column;
    gap: var(--spacing-sm);
    align-items: flex-start;
  }
}
</style> 