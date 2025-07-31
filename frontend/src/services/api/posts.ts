import { mockPostsService } from '../mock/postsService'

const isDevelopment = import.meta.env.DEV
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:3000/api'

// 帖子管理API
export const postsApi = {
  // 获取动态列表
  async getPosts(params: {
    circle_id?: string
    post_type?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockPostsService.getPosts(params)
    }

    const searchParams = new URLSearchParams()
    if (params.circle_id) searchParams.append('circle_id', params.circle_id)
    if (params.post_type) searchParams.append('post_type', params.post_type)
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.page_size) searchParams.append('page_size', params.page_size.toString())

    const response = await fetch(`${API_BASE_URL}/posts?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取动态列表失败')
    }

    return await response.json()
  },

  // 发布动态
  async createPost(postData: {
    circle_id: string
    content: string
    media_urls?: string[]
    post_type: string
  }) {
    if (isDevelopment) {
      return await mockPostsService.createPost(postData)
    }

    const response = await fetch(`${API_BASE_URL}/circles/${postData.circle_id}/posts`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        content: postData.content,
        media_urls: postData.media_urls || [],
        post_type: postData.post_type
      })
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '发布动态失败')
    }

    return await response.json()
  },

  // 点赞动态
  async likePost(postId: string) {
    if (isDevelopment) {
      return await mockPostsService.toggleLike(postId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/posts/${postId}/like`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '点赞失败')
    }

    return await response.json()
  },

  // 取消点赞动态
  async unlikePost(postId: string) {
    if (isDevelopment) {
      return await mockPostsService.toggleLike(postId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/posts/${postId}/like`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '取消点赞失败')
    }

    return await response.json()
  },

  // 获取动态详情
  async getPostDetail(postId: string) {
    if (isDevelopment) {
      return await mockPostsService.getPostDetail(postId)
    }

    const response = await fetch(`${API_BASE_URL}/posts/${postId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取动态详情失败')
    }

    return await response.json()
  },

  // 删除动态
  async deletePost(postId: string) {
    if (isDevelopment) {
      return await mockPostsService.deletePost(postId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/posts/${postId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '删除动态失败')
    }

    return await response.json()
  },

  // 获取用户的动态列表
  async getUserPosts(userId: string, params: {
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      return await mockPostsService.getUserPosts(userId, params)
    }

    const searchParams = new URLSearchParams()
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.page_size) searchParams.append('page_size', params.page_size.toString())

    const response = await fetch(`${API_BASE_URL}/users/${userId}/posts?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取用户动态失败')
    }

    return await response.json()
  },

  // 获取评论列表
  async getComments(postId: string, params: {
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      // 模拟获取评论
      return {
        code: 0,
        message: 'success',
        data: {
          comments: [],
          pagination: {
            page: 1,
            page_size: 20,
            total: 0,
            total_pages: 0
          }
        },
        timestamp: new Date().toISOString()
      }
    }

    const searchParams = new URLSearchParams()
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.page_size) searchParams.append('page_size', params.page_size.toString())

    const response = await fetch(`${API_BASE_URL}/posts/${postId}/comments?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取评论失败')
    }

    return await response.json()
  },

  // 发表评论
  async createComment(postId: string, commentData: {
    content: string
  }) {
    if (isDevelopment) {
      return await mockPostsService.addComment(postId, commentData)
    }

    const response = await fetch(`${API_BASE_URL}/posts/${postId}/comments`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(commentData)
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '发表评论失败')
    }

    return await response.json()
  },

  // 回复评论
  async createReply(commentId: string, replyData: {
    content: string
  }) {
    if (isDevelopment) {
      return await mockPostsService.addReply('1', replyData)
    }

    const response = await fetch(`${API_BASE_URL}/comments/${commentId}/replies`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(replyData)
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '回复评论失败')
    }

    return await response.json()
  },

  // 点赞评论
  async likeComment(commentId: string) {
    if (isDevelopment) {
      return await mockPostsService.toggleCommentLike('1', commentId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/comments/${commentId}/like`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '点赞评论失败')
    }

    return await response.json()
  },

  // 取消点赞评论
  async unlikeComment(commentId: string) {
    if (isDevelopment) {
      return await mockPostsService.toggleCommentLike('1', commentId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/comments/${commentId}/like`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '取消点赞评论失败')
    }

    return await response.json()
  },

  // 点赞回复
  async likeReply(replyId: string) {
    if (isDevelopment) {
      return await mockPostsService.toggleReplyLike('1', replyId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/replies/${replyId}/like`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '点赞回复失败')
    }

    return await response.json()
  },

  // 取消点赞回复
  async unlikeReply(replyId: string) {
    if (isDevelopment) {
      return await mockPostsService.toggleReplyLike('1', replyId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/replies/${replyId}/like`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '取消点赞回复失败')
    }

    return await response.json()
  },

  // 删除评论
  async deleteComment(commentId: string) {
    if (isDevelopment) {
      return await mockPostsService.deleteComment('1', commentId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/comments/${commentId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '删除评论失败')
    }

    return await response.json()
  },

  // 删除回复
  async deleteReply(replyId: string) {
    if (isDevelopment) {
      return await mockPostsService.deleteReply('1', replyId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/replies/${replyId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '删除回复失败')
    }

    return await response.json()
  },

  // 兼容旧接口的方法
  async toggleLike(postId: string, userId: string) {
    if (isDevelopment) {
      return await mockPostsService.toggleLike(postId, userId)
    }
    // 在生产环境中，先检查是否已点赞，然后决定是点赞还是取消点赞
    const postDetail = await this.getPostDetail(postId)
    if (postDetail.data.is_liked) {
      return await this.unlikePost(postId)
    } else {
      return await this.likePost(postId)
    }
  },

  async toggleCommentLike(postId: string, commentId: string, userId: string) {
    if (isDevelopment) {
      return await mockPostsService.toggleCommentLike(postId, commentId, userId)
    }
    // 在生产环境中，先检查是否已点赞，然后决定是点赞还是取消点赞
    const comments = await this.getComments(postId)
    const comment = comments.data.comments.find((c: any) => c.id === commentId)
    if (comment && comment.is_liked) {
      return await this.unlikeComment(commentId)
    } else {
      return await this.likeComment(commentId)
    }
  },

  async toggleReplyLike(postId: string, replyId: string, userId: string) {
    if (isDevelopment) {
      return await mockPostsService.toggleReplyLike(postId, replyId, userId)
    }
    // 在生产环境中，先检查是否已点赞，然后决定是点赞还是取消点赞
    const comments = await this.getComments(postId)
    let reply = null
    for (const comment of comments.data.comments) {
      reply = comment.replies?.find((r: any) => r.id === replyId)
      if (reply) break
    }
    if (reply && reply.is_liked) {
      return await this.unlikeReply(replyId)
    } else {
      return await this.likeReply(replyId)
    }
  },

  async addComment(postId: string, commentData: any) {
    return await this.createComment(postId, commentData)
  },

  async addReply(postId: string, replyData: any) {
    // 注意：这里需要commentId，但旧接口没有提供，需要从replyData中获取
    const commentId = replyData.commentId || replyData.parent_id
    if (!commentId) {
      throw new Error('缺少commentId参数')
    }
    return await this.createReply(commentId, { content: replyData.content })
  },

  async deleteCommentLegacy(postId: string, commentId: string, userId: string) {
    return await this.deleteComment(commentId)
  },

  async deleteReplyLegacy(postId: string, replyId: string, userId: string) {
    return await this.deleteReply(replyId)
  }
} 