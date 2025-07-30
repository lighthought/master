<template>
  <div class="post-card">
    <!-- 头部：用户信息 + 操作菜单 -->
    <div class="post-header">
      <div class="user-info">
        <el-avatar :size="40" :src="post.userAvatar" />
        <div class="user-details">
          <div class="user-name">{{ post.userName }}</div>
          <div class="user-meta">
            <el-tag 
              :type="post.userRole === 'master' ? 'success' : 'warning'" 
              size="small"
              class="role-tag"
            >
              {{ post.userRole === 'master' ? '大师' : '学徒' }}
            </el-tag>
            <span class="post-time">{{ formatTime(post.createdAt) }}</span>
          </div>
        </div>
      </div>
      
      <div class="post-actions" v-if="canManagePost">
        <el-dropdown @command="handleAction">
          <el-button type="text" class="more-btn">
            <el-icon><MoreFilled /></el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="edit">编辑</el-dropdown-item>
              <el-dropdown-item command="delete" divided>删除</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>

    <!-- 内容：文字 + 图片 -->
    <div class="post-content">
      <p class="post-text">{{ post.content }}</p>
      
      <!-- 图片网格 -->
      <div v-if="post.images && post.images.length > 0" class="image-grid">
        <div 
          v-for="(image, index) in post.images" 
          :key="index"
          class="image-item"
          :class="getImageGridClass(post.images.length)"
          @click="previewImage(image, post.images)"
        >
          <img :src="image" :alt="`图片${index + 1}`" />
          <div v-if="post.images.length > 1 && index === 2" class="more-images">
            +{{ post.images.length - 3 }}
          </div>
        </div>
      </div>
    </div>

    <!-- 底部：互动按钮 + 评论区域 -->
    <div class="post-footer">
      <!-- 互动按钮 -->
      <div class="interaction-buttons">
        <el-button 
          type="text" 
          class="action-btn"
          :class="{ 'liked': post.isLiked }"
          @click="handleLike"
          :loading="likeLoading"
        >
          <el-icon><Star /></el-icon>
          <span>{{ post.likes }}</span>
        </el-button>
        
        <el-button 
          type="text" 
          class="action-btn"
          @click="toggleComments"
        >
          <el-icon><ChatDotRound /></el-icon>
          <span>{{ post.comments }}</span>
        </el-button>
        
        <el-button 
          type="text" 
          class="action-btn"
          @click="handleShare"
        >
          <el-icon><Share /></el-icon>
          <span>{{ post.shares }}</span>
        </el-button>
      </div>

      <!-- 评论区域 -->
      <div v-if="showComments" class="comments-section">
        <!-- 评论输入框 -->
        <div class="comment-input">
          <el-avatar :size="32" :src="currentUserAvatar" />
          <el-input
            v-model="commentText"
            placeholder="写下你的评论..."
            @keyup.enter="submitComment"
            :disabled="commentLoading"
          >
            <template #append>
              <el-button 
                type="primary" 
                size="small"
                @click="submitComment"
                :loading="commentLoading"
              >
                发送
              </el-button>
            </template>
          </el-input>
        </div>

        <!-- 评论列表 -->
        <div class="comments-list">
          <div 
            v-for="comment in post.commentsList" 
            :key="comment.id"
            class="comment-item"
          >
            <el-avatar :size="32" :src="comment.userAvatar" />
            <div class="comment-content">
              <div class="comment-header">
                <span class="comment-user">{{ comment.userName }}</span>
                <span class="comment-time">{{ formatTime(comment.createdAt) }}</span>
              </div>
              <p class="comment-text">{{ comment.content }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  MoreFilled, Star, ChatDotRound, Share 
} from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'

// Props
const props = defineProps<{
  post: any
}>()

// Emits
const emit = defineEmits<{
  like: [postId: string, isLiked: boolean]
  comment: [postId: string, comment: any]
  delete: [postId: string]
  edit: [post: any]
}>()

// Store
const authStore = useAuthStore()

// 状态
const showComments = ref(false)
const commentText = ref('')
const likeLoading = ref(false)
const commentLoading = ref(false)

// 计算属性
const currentUserAvatar = computed(() => {
  return authStore.user?.identities?.[0]?.avatar || 'https://via.placeholder.com/32x32/999/FFF?text=U'
})

const canManagePost = computed(() => {
  return authStore.user?.id === props.post.userId
})

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

// 获取图片网格样式类
const getImageGridClass = (count: number) => {
  if (count === 1) return 'single'
  if (count === 2) return 'double'
  if (count === 3) return 'triple'
  return 'multiple'
}

// 预览图片
const previewImage = (image: string, images: string[]) => {
  // TODO: 实现图片预览功能
  console.log('预览图片:', image, images)
}

// 处理点赞
const handleLike = async () => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录')
    return
  }
  
  likeLoading.value = true
  try {
    const result = await ApiService.posts.toggleLike(props.post.id, authStore.user!.id)
    emit('like', props.post.id, result.data.isLiked)
  } catch (error) {
    ElMessage.error('操作失败')
  } finally {
    likeLoading.value = false
  }
}

// 切换评论显示
const toggleComments = () => {
  showComments.value = !showComments.value
}

// 处理分享
const handleShare = () => {
  // TODO: 实现分享功能
  ElMessage.info('分享功能开发中')
}

// 提交评论
const submitComment = async () => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录')
    return
  }
  
  if (!commentText.value.trim()) {
    ElMessage.warning('请输入评论内容')
    return
  }
  
  commentLoading.value = true
  try {
    const commentData = {
      userId: authStore.user!.id,
      userAvatar: currentUserAvatar.value,
      userName: authStore.user!.identities?.[0]?.name || '用户',
      content: commentText.value.trim()
    }
    
    const result = await ApiService.posts.addComment(props.post.id, commentData)
    emit('comment', props.post.id, result.data)
    commentText.value = ''
  } catch (error) {
    ElMessage.error('评论发布失败')
  } finally {
    commentLoading.value = false
  }
}

// 处理操作菜单
const handleAction = async (command: string) => {
  if (command === 'edit') {
    emit('edit', props.post)
  } else if (command === 'delete') {
    try {
      await ElMessageBox.confirm(
        '确定要删除这条动态吗？',
        '确认删除',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
      
      await ApiService.posts.deletePost(props.post.id, authStore.user!.id)
      emit('delete', props.post.id)
      ElMessage.success('删除成功')
    } catch (error) {
      if (error !== 'cancel') {
        ElMessage.error('删除失败')
      }
    }
  }
}
</script>

<style scoped lang="scss">
.post-card {
  background: #333333;
  border: 1px solid #404040;
  border-radius: 12px;
  padding: var(--spacing-lg);
  margin-bottom: var(--spacing-lg);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.post-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--spacing-md);
}

.user-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
}

.user-details {
  .user-name {
    font-size: 14px;
    font-weight: var(--font-weight-bold);
    color: white;
    margin-bottom: var(--spacing-xs);
  }
  
  .user-meta {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
    
    .role-tag {
      font-size: 10px;
    }
    
    .post-time {
      font-size: 12px;
      color: #999;
    }
  }
}

.post-actions {
  .more-btn {
    color: #999;
    padding: 4px;
    
    &:hover {
      color: white;
    }
  }
}

.post-content {
  margin-bottom: var(--spacing-md);
  
  .post-text {
    font-size: 14px;
    color: white;
    line-height: 1.6;
    margin-bottom: var(--spacing-md);
    white-space: pre-wrap;
  }
}

.image-grid {
  display: grid;
  gap: var(--spacing-xs);
  
  &.single {
    grid-template-columns: 1fr;
    
    .image-item {
      max-height: 300px;
    }
  }
  
  &.double {
    grid-template-columns: 1fr 1fr;
    
    .image-item {
      max-height: 200px;
    }
  }
  
  &.triple {
    grid-template-columns: 1fr 1fr 1fr;
    
    .image-item {
      max-height: 150px;
    }
  }
  
  &.multiple {
    grid-template-columns: 1fr 1fr 1fr;
    
    .image-item {
      max-height: 120px;
      
      &:nth-child(3) {
        position: relative;
      }
    }
  }
}

.image-item {
  position: relative;
  border-radius: 8px;
  overflow: hidden;
  cursor: pointer;
  
  img {
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
    font-size: 18px;
    font-weight: var(--font-weight-bold);
  }
}

.post-footer {
  .interaction-buttons {
    display: flex;
    gap: var(--spacing-lg);
    margin-bottom: var(--spacing-md);
    
    .action-btn {
      color: #999;
      font-size: 14px;
      
      &:hover {
        color: var(--primary-color);
      }
      
      &.liked {
        color: var(--primary-color);
      }
      
      .el-icon {
        margin-right: var(--spacing-xs);
      }
    }
  }
}

.comments-section {
  border-top: 1px solid #404040;
  padding-top: var(--spacing-md);
  
  .comment-input {
    display: flex;
    gap: var(--spacing-sm);
    margin-bottom: var(--spacing-md);
    
    .el-input {
      flex: 1;
    }
  }
  
  .comments-list {
    .comment-item {
      display: flex;
      gap: var(--spacing-sm);
      margin-bottom: var(--spacing-md);
      
      &:last-child {
        margin-bottom: 0;
      }
      
      .comment-content {
        flex: 1;
        
        .comment-header {
          display: flex;
          align-items: center;
          gap: var(--spacing-sm);
          margin-bottom: var(--spacing-xs);
          
          .comment-user {
            font-size: 12px;
            font-weight: var(--font-weight-semibold);
            color: white;
          }
          
          .comment-time {
            font-size: 11px;
            color: #999;
          }
        }
        
        .comment-text {
          font-size: 13px;
          color: #ccc;
          line-height: 1.4;
          margin: 0;
        }
      }
    }
  }
}

// 响应式设计
@media (max-width: 768px) {
  .post-card {
    padding: var(--spacing-md);
  }
  
  .image-grid {
    &.double,
    &.triple,
    &.multiple {
      grid-template-columns: 1fr 1fr;
      
      .image-item {
        max-height: 120px;
      }
    }
  }
  
  .interaction-buttons {
    gap: var(--spacing-md);
  }
}
</style> 