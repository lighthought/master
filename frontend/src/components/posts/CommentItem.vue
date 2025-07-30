<template>
  <div class="comment-item">
    <el-avatar :size="32" :src="comment.userAvatar" />
    <div class="comment-content">
      <div class="comment-header">
        <div class="comment-user-info">
          <span class="comment-user">{{ comment.userName }}</span>
          <el-tag 
            :type="comment.userRole === 'master' ? 'success' : 'warning'" 
            size="small"
            class="role-tag"
          >
            {{ comment.userRole === 'master' ? '大师' : '学徒' }}
          </el-tag>
        </div>
        <span class="comment-time">{{ formatTime(comment.createdAt) }}</span>
      </div>
      
      <p class="comment-text">{{ comment.content }}</p>
      
      <!-- 评论操作 -->
      <div class="comment-actions">
        <el-button 
          type="text" 
          size="small"
          @click="toggleLike"
          :class="{ 'liked': comment.isLiked }"
        >
          <el-icon><Star /></el-icon>
          <span>{{ comment.likes || 0 }}</span>
        </el-button>
        
        <el-button 
          type="text" 
          size="small"
          @click="toggleReply"
        >
          回复
        </el-button>
        
        <el-button 
          v-if="canDeleteComment"
          type="text" 
          size="small"
          @click="deleteComment"
        >
          删除
        </el-button>
      </div>
      
      <!-- 回复输入框 -->
      <div v-if="showReplyInput" class="reply-input">
        <el-input
          v-model="replyText"
          placeholder="回复评论..."
          @keyup.enter="submitReply"
          :disabled="replyLoading"
        >
          <template #append>
            <el-button 
              type="primary" 
              size="small"
              @click="submitReply"
              :loading="replyLoading"
            >
              发送
            </el-button>
          </template>
        </el-input>
      </div>
      
      <!-- 回复列表 -->
      <div v-if="comment.replies && comment.replies.length > 0" class="replies-list">
        <div 
          v-for="reply in comment.replies" 
          :key="reply.id"
          class="reply-item"
        >
          <el-avatar :size="24" :src="reply.userAvatar" />
          <div class="reply-content">
            <div class="reply-header">
              <div class="reply-user-info">
                <span class="reply-user">{{ reply.userName }}</span>
                <el-tag 
                  :type="reply.userRole === 'master' ? 'success' : 'warning'" 
                  size="small"
                  class="role-tag"
                >
                  {{ reply.userRole === 'master' ? '大师' : '学徒' }}
                </el-tag>
              </div>
              <span class="reply-time">{{ formatTime(reply.createdAt) }}</span>
            </div>
            <p class="reply-text">{{ reply.content }}</p>
            
            <!-- 回复操作 -->
            <div class="reply-actions">
              <el-button 
                type="text" 
                size="small"
                @click="toggleReplyLike(reply)"
                :class="{ 'liked': reply.isLiked }"
              >
                <el-icon><Star /></el-icon>
                <span>{{ reply.likes || 0 }}</span>
              </el-button>
              
              <el-button 
                v-if="canDeleteReply(reply)"
                type="text" 
                size="small"
                @click="deleteReply(reply)"
              >
                删除
              </el-button>
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
import { Star } from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'

// Props
const props = defineProps<{
  comment: any
  postId: string
}>()

// Emits
const emit = defineEmits<{
  'update-comment': [comment: any]
  'delete-comment': [commentId: string]
}>()

// Store
const authStore = useAuthStore()

// 状态
const showReplyInput = ref(false)
const replyText = ref('')
const replyLoading = ref(false)

// 计算属性
const canDeleteComment = computed(() => {
  return authStore.user?.id === props.comment.userId
})

const canDeleteReply = (reply: any) => {
  return authStore.user?.id === reply.userId
}

// 格式化时间
const formatTime = (time: string) => {
  const now = new Date()
  const commentTime = new Date(time)
  const diff = now.getTime() - commentTime.getTime()
  
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
    return commentTime.toLocaleDateString()
  }
}

// 切换点赞
const toggleLike = async () => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录')
    return
  }
  
  try {
    const result = await ApiService.posts.toggleCommentLike(
      props.postId, 
      props.comment.id, 
      authStore.user!.id
    )
    
    // 更新评论的点赞状态
    const updatedComment = { ...props.comment }
    updatedComment.isLiked = result.data.isLiked
    updatedComment.likes = result.data.likes
    emit('update-comment', updatedComment)
  } catch (error) {
    ElMessage.error('操作失败')
  }
}

// 切换回复输入框
const toggleReply = () => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录')
    return
  }
  showReplyInput.value = !showReplyInput.value
}

// 提交回复
const submitReply = async () => {
  if (!replyText.value.trim()) {
    ElMessage.warning('请输入回复内容')
    return
  }
  
  replyLoading.value = true
  try {
    const replyData = {
      userId: authStore.user!.id,
      userAvatar: authStore.user!.identities?.[0]?.avatar || 'https://via.placeholder.com/24x24/999/FFF?text=U',
      userName: authStore.user!.identities?.[0]?.name || '用户',
      userRole: authStore.user!.identities?.[0]?.type || 'apprentice',
      content: replyText.value.trim(),
      parentId: props.comment.id
    }
    
    const result = await ApiService.posts.addReply(props.postId, replyData)
    
    // 更新评论的回复列表
    const updatedComment = { ...props.comment }
    if (!updatedComment.replies) {
      updatedComment.replies = []
    }
    updatedComment.replies.unshift(result.data)
    emit('update-comment', updatedComment)
    
    replyText.value = ''
    showReplyInput.value = false
  } catch (error) {
    ElMessage.error('回复发布失败')
  } finally {
    replyLoading.value = false
  }
}

// 删除评论
const deleteComment = async () => {
  try {
    await ElMessageBox.confirm(
      '确定要删除这条评论吗？',
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    await ApiService.posts.deleteComment(props.postId, props.comment.id, authStore.user!.id)
    emit('delete-comment', props.comment.id)
    ElMessage.success('删除成功')
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

// 删除回复
const deleteReply = async (reply: any) => {
  try {
    await ElMessageBox.confirm(
      '确定要删除这条回复吗？',
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    await ApiService.posts.deleteReply(props.postId, reply.id, authStore.user!.id)
    
    // 从评论的回复列表中移除
    const updatedComment = { ...props.comment }
    updatedComment.replies = updatedComment.replies.filter((r: any) => r.id !== reply.id)
    emit('update-comment', updatedComment)
    
    ElMessage.success('删除成功')
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

// 切换回复点赞
const toggleReplyLike = async (reply: any) => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录')
    return
  }
  
  try {
    const result = await ApiService.posts.toggleReplyLike(
      props.postId, 
      reply.id, 
      authStore.user!.id
    )
    
    // 更新回复的点赞状态
    const updatedComment = { ...props.comment }
    const replyIndex = updatedComment.replies.findIndex((r: any) => r.id === reply.id)
    if (replyIndex > -1) {
      updatedComment.replies[replyIndex].isLiked = result.data.isLiked
      updatedComment.replies[replyIndex].likes = result.data.likes
      emit('update-comment', updatedComment)
    }
  } catch (error) {
    ElMessage.error('操作失败')
  }
}
</script>

<style scoped lang="scss">
.comment-item {
  display: flex;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-md);
  
  .comment-content {
    flex: 1;
    
    .comment-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: var(--spacing-xs);
      
      .comment-user-info {
        display: flex;
        align-items: center;
        gap: var(--spacing-xs);
        
        .comment-user {
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
        }
        
        .role-tag {
          font-size: 10px;
        }
      }
      
      .comment-time {
        font-size: var(--font-size-small);
        color: var(--text-secondary);
      }
    }
    
    .comment-text {
      margin: 0 0 var(--spacing-sm) 0;
      color: var(--text-primary);
      line-height: 1.5;
    }
    
    .comment-actions {
      display: flex;
      gap: var(--spacing-md);
      margin-bottom: var(--spacing-sm);
      
      .el-button {
        color: var(--text-secondary);
        font-size: var(--font-size-small);
        
        &.liked {
          color: var(--primary-color);
        }
        
        .el-icon {
          margin-right: 4px;
        }
      }
    }
    
    .reply-input {
      margin-bottom: var(--spacing-sm);
    }
    
    .replies-list {
      margin-top: var(--spacing-sm);
      padding-left: var(--spacing-md);
      border-left: 2px solid var(--border-color);
      
      .reply-item {
        display: flex;
        gap: var(--spacing-xs);
        margin-bottom: var(--spacing-sm);
        
        .reply-content {
          flex: 1;
          
          .reply-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: var(--spacing-xs);
            
            .reply-user-info {
              display: flex;
              align-items: center;
              gap: var(--spacing-xs);
              
              .reply-user {
                font-weight: var(--font-weight-medium);
                color: var(--text-primary);
                font-size: var(--font-size-small);
              }
              
              .role-tag {
                font-size: 8px;
              }
            }
            
            .reply-time {
              font-size: 10px;
              color: var(--text-secondary);
            }
          }
          
          .reply-text {
            margin: 0 0 var(--spacing-xs) 0;
            color: var(--text-primary);
            line-height: 1.4;
            font-size: var(--font-size-small);
          }
          
          .reply-actions {
            display: flex;
            gap: var(--spacing-sm);
            
            .el-button {
              color: var(--text-secondary);
              font-size: 10px;
              
              &.liked {
                color: var(--primary-color);
              }
              
              .el-icon {
                margin-right: 2px;
              }
            }
          }
        }
      }
    }
  }
}
</style> 