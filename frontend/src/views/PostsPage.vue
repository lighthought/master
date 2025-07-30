<template>
  <div class="posts-page">
    <!-- 页面头部 -->
    <div class="page-header">
      <h1 class="page-title">动态广场</h1>
      <el-button type="primary" @click="showPublishDialog = true">
        <el-icon><Plus /></el-icon>
        发布动态
      </el-button>
    </div>

    <!-- 搜索和筛选 -->
    <div class="filter-section">
      <div class="search-box">
        <el-input
          v-model="searchKeyword"
          placeholder="搜索动态内容或用户..."
          clearable
          @input="handleSearch"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
      </div>
      
      <div class="filter-options">
        <el-select v-model="filterRole" placeholder="选择身份" clearable @change="loadPosts">
          <el-option label="全部" value="" />
          <el-option label="大师" value="master" />
          <el-option label="学徒" value="apprentice" />
        </el-select>
        
        <el-select v-model="sortBy" placeholder="排序方式" @change="loadPosts">
          <el-option label="最新发布" value="latest" />
          <el-option label="最受欢迎" value="popular" />
        </el-select>
      </div>
    </div>

    <!-- 动态列表 -->
    <div class="posts-container">
      <div v-if="loading" class="loading-container">
        <el-skeleton :rows="3" animated />
        <el-skeleton :rows="3" animated />
        <el-skeleton :rows="3" animated />
      </div>
      
      <div v-else-if="posts.length > 0" class="posts-list">
        <PostCard
          v-for="post in posts"
          :key="post.id"
          :post="post"
          @like="handleLike"
          @comment="handleComment"
          @delete="handleDelete"
          @edit="handleEdit"
        />
        
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
        <el-empty description="暂无动态">
          <el-button type="primary" @click="showPublishDialog = true">
            发布第一条动态
          </el-button>
        </el-empty>
      </div>
    </div>

    <!-- 发布动态对话框 -->
    <el-dialog
      v-model="showPublishDialog"
      title="发布动态"
      width="90%"
      max-width="600px"
      :close-on-click-modal="false"
    >
      <div class="publish-form">
        <el-form :model="publishForm" label-width="80px">
          <el-form-item label="动态内容">
            <el-input
              v-model="publishForm.content"
              type="textarea"
              :rows="4"
              placeholder="分享你的想法、经验或学习心得..."
              maxlength="500"
              show-word-limit
            />
          </el-form-item>
          
          <el-form-item label="添加图片">
            <div class="image-upload">
              <div 
                v-for="(image, index) in publishForm.images" 
                :key="index"
                class="image-preview"
              >
                <img :src="image" :alt="`图片${index + 1}`" />
                <el-button 
                  type="danger" 
                  size="small" 
                  circle
                  class="remove-btn"
                  @click="removeImage(index)"
                >
                  <el-icon><Close /></el-icon>
                </el-button>
              </div>
              
              <el-upload
                v-if="publishForm.images.length < 9"
                action="#"
                :auto-upload="false"
                :show-file-list="false"
                accept="image/*"
                @change="handleImageUpload"
              >
                <div class="upload-placeholder">
                  <el-icon><Plus /></el-icon>
                  <span>添加图片</span>
                </div>
              </el-upload>
            </div>
          </el-form-item>
        </el-form>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showPublishDialog = false">取消</el-button>
          <el-button 
            type="primary" 
            @click="publishPost"
            :loading="publishing"
            :disabled="!publishForm.content.trim()"
          >
            发布动态
          </el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 编辑动态对话框 -->
    <el-dialog
      v-model="showEditDialog"
      title="编辑动态"
      width="90%"
      max-width="600px"
      :close-on-click-modal="false"
    >
      <div class="publish-form">
        <el-form :model="editForm" label-width="80px">
          <el-form-item label="动态内容">
            <el-input
              v-model="editForm.content"
              type="textarea"
              :rows="4"
              placeholder="分享你的想法、经验或学习心得..."
              maxlength="500"
              show-word-limit
            />
          </el-form-item>
          
          <el-form-item label="图片">
            <div class="image-upload">
              <div 
                v-for="(image, index) in editForm.images" 
                :key="index"
                class="image-preview"
              >
                <img :src="image" :alt="`图片${index + 1}`" />
                <el-button 
                  type="danger" 
                  size="small" 
                  circle
                  class="remove-btn"
                  @click="removeEditImage(index)"
                >
                  <el-icon><Close /></el-icon>
                </el-button>
              </div>
              
              <el-upload
                v-if="editForm.images.length < 9"
                action="#"
                :auto-upload="false"
                :show-file-list="false"
                accept="image/*"
                @change="handleEditImageUpload"
              >
                <div class="upload-placeholder">
                  <el-icon><Plus /></el-icon>
                  <span>添加图片</span>
                </div>
              </el-upload>
            </div>
          </el-form-item>
        </el-form>
      </div>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showEditDialog = false">取消</el-button>
          <el-button 
            type="primary" 
            @click="updatePost"
            :loading="updating"
            :disabled="!editForm.content.trim()"
          >
            更新动态
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { 
  Plus, Search, Close 
} from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'
import PostCard from '@/components/posts/PostCard.vue'

// Store
const authStore = useAuthStore()

// 状态
const loading = ref(false)
const loadingMore = ref(false)
const publishing = ref(false)
const updating = ref(false)
const showPublishDialog = ref(false)
const showEditDialog = ref(false)

// 数据
const posts = ref<any[]>([])
const searchKeyword = ref('')
const filterRole = ref('')
const sortBy = ref('latest')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)

// 表单
const publishForm = ref({
  content: '',
  images: [] as string[]
})

const editForm = ref({
  id: '',
  content: '',
  images: [] as string[]
})

// 计算属性
const hasMore = computed(() => {
  return posts.value.length < total.value
})

// 加载动态列表
const loadPosts = async (reset = true) => {
  if (reset) {
    currentPage.value = 1
    posts.value = []
  }
  
  loading.value = true
  try {
    const params = {
      page: currentPage.value,
      pageSize: pageSize.value,
      sort: sortBy.value,
      keyword: searchKeyword.value,
      userRole: filterRole.value
    }
    
    const result = await ApiService.posts.getPosts(params)
    const newPosts = result.data.posts
    
    if (reset) {
      posts.value = newPosts
    } else {
      posts.value.push(...newPosts)
    }
    
    total.value = result.data.total
  } catch (error) {
    ElMessage.error('加载动态失败')
  } finally {
    loading.value = false
  }
}

// 加载更多
const loadMore = async () => {
  if (loadingMore.value) return
  
  loadingMore.value = true
  currentPage.value++
  await loadPosts(false)
  loadingMore.value = false
}

// 搜索
const handleSearch = () => {
  loadPosts()
}

// 发布动态
const publishPost = async () => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('请先登录')
    return
  }
  
  if (!publishForm.value.content.trim()) {
    ElMessage.warning('请输入动态内容')
    return
  }
  
  publishing.value = true
  try {
    const postData = {
      userId: authStore.user!.id,
      userAvatar: authStore.user!.identities?.[0]?.avatar || 'https://via.placeholder.com/40x40/999/FFF?text=U',
      userName: authStore.user!.identities?.[0]?.name || '用户',
      userRole: authStore.user!.identities?.[0]?.type || 'apprentice',
      content: publishForm.value.content.trim(),
      images: publishForm.value.images
    }
    
    await ApiService.posts.createPost(postData)
    ElMessage.success('动态发布成功')
    
    // 重置表单并关闭对话框
    publishForm.value = { content: '', images: [] }
    showPublishDialog.value = false
    
    // 重新加载动态列表
    loadPosts()
  } catch (error) {
    ElMessage.error('发布失败')
  } finally {
    publishing.value = false
  }
}

// 处理图片上传
const handleImageUpload = (file: any) => {
  // 模拟图片上传，实际项目中需要调用真实的上传API
  const reader = new FileReader()
  reader.onload = (e) => {
    publishForm.value.images.push(e.target?.result as string)
  }
  reader.readAsDataURL(file.raw)
}

// 移除图片
const removeImage = (index: number) => {
  publishForm.value.images.splice(index, 1)
}

// 处理编辑图片上传
const handleEditImageUpload = (file: any) => {
  const reader = new FileReader()
  reader.onload = (e) => {
    editForm.value.images.push(e.target?.result as string)
  }
  reader.readAsDataURL(file.raw)
}

// 移除编辑图片
const removeEditImage = (index: number) => {
  editForm.value.images.splice(index, 1)
}

// 处理点赞
const handleLike = (postId: string, isLiked: boolean) => {
  const post = posts.value.find(p => p.id === postId)
  if (post) {
    post.isLiked = isLiked
    post.likes += isLiked ? 1 : -1
  }
}

// 处理评论
const handleComment = (postId: string, comment: any) => {
  const post = posts.value.find(p => p.id === postId)
  if (post) {
    post.commentsList.unshift(comment)
    post.comments++
  }
}

// 处理删除
const handleDelete = (postId: string) => {
  const index = posts.value.findIndex(p => p.id === postId)
  if (index > -1) {
    posts.value.splice(index, 1)
  }
}

// 处理编辑
const handleEdit = (post: any) => {
  editForm.value = {
    id: post.id,
    content: post.content,
    images: [...post.images]
  }
  showEditDialog.value = true
}

// 更新动态
const updatePost = async () => {
  if (!editForm.value.content.trim()) {
    ElMessage.warning('请输入动态内容')
    return
  }
  
  updating.value = true
  try {
    // 这里应该调用更新API，暂时模拟更新
    const post = posts.value.find(p => p.id === editForm.value.id)
    if (post) {
      post.content = editForm.value.content
      post.images = [...editForm.value.images]
    }
    
    ElMessage.success('动态更新成功')
    showEditDialog.value = false
  } catch (error) {
    ElMessage.error('更新失败')
  } finally {
    updating.value = false
  }
}

// 组件挂载时加载数据
onMounted(() => {
  loadPosts()
})
</script>

<style scoped lang="scss">
.posts-page {
  max-width: 800px;
  margin: 0 auto;
  padding: var(--spacing-xl);
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-xl);
  
  .page-title {
    font-size: var(--font-size-h2);
    font-weight: var(--font-weight-bold);
    color: var(--text-primary);
    margin: 0;
  }
}

.filter-section {
  display: flex;
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-xl);
  align-items: center;
  
  .search-box {
    flex: 1;
  }
  
  .filter-options {
    display: flex;
    gap: var(--spacing-md);
  }
}

.posts-container {
  .loading-container {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-lg);
  }
  
  .posts-list {
    .load-more {
      text-align: center;
      margin-top: var(--spacing-lg);
    }
  }
  
  .empty-state {
    text-align: center;
    padding: var(--spacing-xxl) 0;
  }
}

.publish-form {
  .image-upload {
    display: flex;
    flex-wrap: wrap;
    gap: var(--spacing-sm);
    
    .image-preview {
      position: relative;
      width: 100px;
      height: 100px;
      border-radius: var(--border-radius-small);
      overflow: hidden;
      
      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
      
      .remove-btn {
        position: absolute;
        top: 4px;
        right: 4px;
        width: 20px;
        height: 20px;
        font-size: 10px;
      }
    }
    
    .upload-placeholder {
      width: 100px;
      height: 100px;
      border: 2px dashed var(--border-color);
      border-radius: var(--border-radius-small);
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      color: var(--text-secondary);
      
      &:hover {
        border-color: var(--primary-color);
        color: var(--primary-color);
      }
      
      .el-icon {
        font-size: 24px;
        margin-bottom: var(--spacing-xs);
      }
      
      span {
        font-size: var(--font-size-small);
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
  .posts-page {
    padding: var(--spacing-lg);
  }
  
  .page-header {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: stretch;
  }
  
  .filter-section {
    flex-direction: column;
    gap: var(--spacing-md);
    
    .filter-options {
      justify-content: space-between;
    }
  }
}
</style> 