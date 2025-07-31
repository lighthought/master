import { mockCirclesService } from '../mock/circlesService'

const isDevelopment = import.meta.env.DEV
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:3000/api'

// 圈子管理API
export const circlesApi = {
  // 获取圈子列表
  async getCircles(params: {
    category?: string
    query?: string
    sort?: string
    page?: number
    pageSize?: number
  } = {}) {
    if (isDevelopment) {
      return await mockCirclesService.getCircles(params)
    }

    const searchParams = new URLSearchParams()
    if (params.category) searchParams.append('category', params.category)
    if (params.query) searchParams.append('query', params.query)
    if (params.sort) searchParams.append('sort', params.sort)
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.pageSize) searchParams.append('page_size', params.pageSize.toString())

    const response = await fetch(`${API_BASE_URL}/circles?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取圈子列表失败')
    }

    return await response.json()
  },

  // 获取圈子详情
  async getCircleDetail(circleId: string) {
    if (isDevelopment) {
      return await mockCirclesService.getCircleDetail(circleId)
    }

    const response = await fetch(`${API_BASE_URL}/circles/${circleId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取圈子详情失败')
    }

    return await response.json()
  },

  // 获取圈子动态
  async getCirclePosts(circleId: string, params: {
    post_type?: string
    page?: number
    page_size?: number
  } = {}) {
    if (isDevelopment) {
      // 模拟获取圈子动态
      return {
        code: 0,
        message: 'success',
        data: {
          posts: [],
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
    if (params.post_type) searchParams.append('post_type', params.post_type)
    if (params.page) searchParams.append('page', params.page.toString())
    if (params.page_size) searchParams.append('page_size', params.page_size.toString())

    const response = await fetch(`${API_BASE_URL}/circles/${circleId}/posts?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取圈子动态失败')
    }

    return await response.json()
  },

  // 发布动态
  async createPost(circleId: string, postData: {
    content: string
    media_urls?: string[]
    post_type: string
  }) {
    if (isDevelopment) {
      // 模拟发布动态
      return {
        code: 0,
        message: '动态发布成功',
        data: {
          post_id: Date.now().toString()
        },
        timestamp: new Date().toISOString()
      }
    }

    const response = await fetch(`${API_BASE_URL}/circles/${circleId}/posts`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(postData)
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
      // 模拟点赞
      return {
        code: 0,
        message: '点赞成功',
        data: {
          post_id: postId,
          like_count: Math.floor(Math.random() * 100) + 1
        },
        timestamp: new Date().toISOString()
      }
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

  // 取消点赞
  async unlikePost(postId: string) {
    if (isDevelopment) {
      // 模拟取消点赞
      return {
        code: 0,
        message: '取消点赞成功',
        data: {
          post_id: postId,
          like_count: Math.floor(Math.random() * 100)
        },
        timestamp: new Date().toISOString()
      }
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
      // 模拟发表评论
      return {
        code: 0,
        message: '评论发表成功',
        data: {
          comment_id: Date.now().toString()
        },
        timestamp: new Date().toISOString()
      }
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
      // 模拟回复评论
      return {
        code: 0,
        message: '回复发表成功',
        data: {
          reply_id: Date.now().toString()
        },
        timestamp: new Date().toISOString()
      }
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
      // 模拟点赞评论
      return {
        code: 0,
        message: '点赞成功',
        data: {
          comment_id: commentId,
          like_count: Math.floor(Math.random() * 50) + 1
        },
        timestamp: new Date().toISOString()
      }
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
      // 模拟取消点赞评论
      return {
        code: 0,
        message: '取消点赞成功',
        data: {
          comment_id: commentId,
          like_count: Math.floor(Math.random() * 50)
        },
        timestamp: new Date().toISOString()
      }
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
      // 模拟点赞回复
      return {
        code: 0,
        message: '点赞成功',
        data: {
          reply_id: replyId,
          like_count: Math.floor(Math.random() * 20) + 1
        },
        timestamp: new Date().toISOString()
      }
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
      // 模拟取消点赞回复
      return {
        code: 0,
        message: '取消点赞成功',
        data: {
          reply_id: replyId,
          like_count: Math.floor(Math.random() * 20)
        },
        timestamp: new Date().toISOString()
      }
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
      // 模拟删除评论
      return {
        code: 0,
        message: '评论删除成功',
        data: {
          comment_id: commentId
        },
        timestamp: new Date().toISOString()
      }
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
      // 模拟删除回复
      return {
        code: 0,
        message: '回复删除成功',
        data: {
          reply_id: replyId
        },
        timestamp: new Date().toISOString()
      }
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

  // 加入圈子
  async joinCircle(circleId: string) {
    if (isDevelopment) {
      return await mockCirclesService.joinCircle(circleId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/circles/${circleId}/join`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '加入圈子失败')
    }

    return await response.json()
  },

  // 退出圈子
  async leaveCircle(circleId: string) {
    if (isDevelopment) {
      return await mockCirclesService.leaveCircle(circleId, '1')
    }

    const response = await fetch(`${API_BASE_URL}/circles/${circleId}/join`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '退出圈子失败')
    }

    return await response.json()
  },

  // 获取推荐圈子
  async getRecommendedCircles(userId?: string) {
    if (isDevelopment) {
      // 模拟获取推荐圈子
      return {
        code: 0,
        message: 'success',
        data: {
          circles: []
        },
        timestamp: new Date().toISOString()
      }
    }

    const searchParams = new URLSearchParams()
    if (userId) searchParams.append('user_id', userId)

    const response = await fetch(`${API_BASE_URL}/circles/recommended?${searchParams}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`,
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || '获取推荐圈子失败')
    }

    return await response.json()
  },

  // 兼容旧接口的方法
  async getUserJoinedCircles(userId: string) {
    if (isDevelopment) {
      return await mockCirclesService.getUserJoinedCircles(userId)
    }
    // 在生产环境中，可以通过获取用户信息来获取加入的圈子
    return await this.getCircles({ query: 'joined' })
  },

  async getCircleCategories() {
    if (isDevelopment) {
      return await mockCirclesService.getCircleCategories()
    }
    // 在生产环境中，可以通过获取圈子列表来获取分类信息
    return {
      code: 0,
      message: 'success',
      data: {
        categories: [
          { id: 'software', name: '软件开发' },
          { id: 'design', name: 'UI设计' },
          { id: 'marketing', name: '数字营销' },
          { id: 'craft', name: '传统工艺' }
        ]
      },
      timestamp: new Date().toISOString()
    }
  },

  async searchCircles(query: string) {
    if (isDevelopment) {
      return await mockCirclesService.searchCircles(query)
    }
    return await this.getCircles({ query })
  }
} 