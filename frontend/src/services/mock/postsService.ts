const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟动态数据
const mockPosts = [
  {
    id: '1',
    userId: '1',
    userAvatar: 'https://via.placeholder.com/40x40/4CAF50/FFFFFF?text=张',
    userName: '张大师',
    userRole: 'master',
    content: '今天分享一个Vue.js 3.0的实战技巧，使用Composition API可以让代码更加清晰和可维护。大家有什么问题欢迎在评论区讨论！',
    images: [
      'https://via.placeholder.com/300x200/4CAF50/FFFFFF?text=Vue3',
      'https://via.placeholder.com/300x200/2196F3/FFFFFF?text=Code'
    ],
    likes: 24,
    comments: 8,
    shares: 3,
    isLiked: false,
    createdAt: new Date(Date.now() - 2 * 60 * 60 * 1000).toISOString(),
    commentsList: [
      {
        id: '1',
        userId: '2',
        userAvatar: 'https://via.placeholder.com/32x32/FF9800/FFFFFF?text=李',
        userName: '李学徒',
        userRole: 'apprentice',
        content: '感谢分享！这个技巧确实很实用',
        likes: 3,
        isLiked: false,
        createdAt: new Date(Date.now() - 1 * 60 * 60 * 1000).toISOString(),
        replies: []
      },
      {
        id: '2',
        userId: '3',
        userAvatar: 'https://via.placeholder.com/32x32/9C27B0/FFFFFF?text=王',
        userName: '王同学',
        userRole: 'apprentice',
        content: '请问这个和Options API相比有什么优势？',
        likes: 1,
        isLiked: false,
        createdAt: new Date(Date.now() - 30 * 60 * 1000).toISOString(),
        replies: []
      }
    ]
  },
  {
    id: '2',
    userId: '4',
    userAvatar: 'https://via.placeholder.com/40x40/FF5722/FFFFFF?text=陈',
    userName: '陈工程师',
    userRole: 'master',
    content: '刚完成了一个微服务架构的项目，使用Spring Boot + Docker + Kubernetes。整个部署过程非常顺畅，推荐大家尝试！',
    images: [
      'https://via.placeholder.com/300x200/FF5722/FFFFFF?text=Spring',
      'https://via.placeholder.com/300x200/607D8B/FFFFFF?text=Docker',
      'https://via.placeholder.com/300x200/795548/FFFFFF?text=K8s'
    ],
    likes: 18,
    comments: 5,
    shares: 2,
    isLiked: true,
    createdAt: new Date(Date.now() - 4 * 60 * 60 * 1000).toISOString(),
    commentsList: [
      {
        id: '3',
        userId: '5',
        userAvatar: 'https://via.placeholder.com/32x32/3F51B5/FFFFFF?text=赵',
        userName: '赵开发者',
        userRole: 'apprentice',
        content: '能分享一下具体的部署配置吗？',
        likes: 2,
        isLiked: false,
        createdAt: new Date(Date.now() - 2 * 60 * 60 * 1000).toISOString(),
        replies: []
      }
    ]
  },
  {
    id: '3',
    userId: '6',
    userAvatar: 'https://via.placeholder.com/40x40/00BCD4/FFFFFF?text=刘',
    userName: '刘学徒',
    userRole: 'apprentice',
    content: '今天学习了React Hooks，感觉比class组件好用多了！感谢导师的耐心指导。',
    images: [
      'https://via.placeholder.com/300x200/00BCD4/FFFFFF?text=React'
    ],
    likes: 12,
    comments: 3,
    shares: 1,
    isLiked: false,
    createdAt: new Date(Date.now() - 6 * 60 * 60 * 1000).toISOString(),
    commentsList: []
  }
]

export const mockPostsService = {
  // 获取动态列表
  async getPosts(params: any = {}) {
    await delay(500)
    
    let posts = [...mockPosts]
    
    // 根据用户角色过滤
    if (params.userRole) {
      posts = posts.filter(post => post.userRole === params.userRole)
    }
    
    // 根据关键词搜索
    if (params.keyword) {
      posts = posts.filter(post => 
        post.content.includes(params.keyword) || 
        post.userName.includes(params.keyword)
      )
    }
    
    // 排序
    if (params.sort === 'latest') {
      posts.sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime())
    } else if (params.sort === 'popular') {
      posts.sort((a, b) => (b.likes + b.comments) - (a.likes + a.comments))
    }
    
    // 分页
    const page = params.page || 1
    const pageSize = params.pageSize || 10
    const start = (page - 1) * pageSize
    const end = start + pageSize
    
    return {
      success: true,
      data: {
        posts: posts.slice(start, end),
        total: posts.length,
        page,
        pageSize
      }
    }
  },

  // 发布动态
  async createPost(postData: any) {
    await delay(800)
    
    const newPost = {
      id: Date.now().toString(),
      userId: postData.userId,
      userAvatar: postData.userAvatar,
      userName: postData.userName,
      userRole: postData.userRole,
      content: postData.content,
      images: postData.images || [],
      likes: 0,
      comments: 0,
      shares: 0,
      isLiked: false,
      createdAt: new Date().toISOString(),
      commentsList: []
    }
    
    mockPosts.unshift(newPost)
    
    return {
      success: true,
      data: newPost,
      message: '动态发布成功'
    }
  },

  // 点赞/取消点赞
  async toggleLike(postId: string, userId: string) {
    await delay(300)
    
    const post = mockPosts.find(p => p.id === postId)
    if (!post) {
      throw new Error('动态不存在')
    }
    
    if (post.isLiked) {
      post.likes--
      post.isLiked = false
    } else {
      post.likes++
      post.isLiked = true
    }
    
    return {
      success: true,
      data: {
        likes: post.likes,
        isLiked: post.isLiked
      }
    }
  },

  // 添加评论
  async addComment(postId: string, commentData: any) {
    await delay(500)
    
    const post = mockPosts.find(p => p.id === postId)
    if (!post) {
      throw new Error('动态不存在')
    }
    
    const newComment = {
      id: Date.now().toString(),
      userId: commentData.userId,
      userAvatar: commentData.userAvatar,
      userName: commentData.userName,
      userRole: commentData.userRole || 'apprentice',
      content: commentData.content,
      likes: 0,
      isLiked: false,
      createdAt: new Date().toISOString(),
      replies: []
    }
    
    post.commentsList.unshift(newComment)
    post.comments++
    
    return {
      success: true,
      data: newComment,
      message: '评论发布成功'
    }
  },

  // 获取动态详情
  async getPostDetail(postId: string) {
    await delay(300)
    
    const post = mockPosts.find(p => p.id === postId)
    if (!post) {
      throw new Error('动态不存在')
    }
    
    return {
      success: true,
      data: post
    }
  },

  // 删除动态
  async deletePost(postId: string, userId: string) {
    await delay(500)
    
    const index = mockPosts.findIndex(p => p.id === postId && p.userId === userId)
    if (index === -1) {
      throw new Error('动态不存在或无权限删除')
    }
    
    mockPosts.splice(index, 1)
    
    return {
      success: true,
      message: '动态删除成功'
    }
  },

  // 获取用户的动态列表
  async getUserPosts(userId: string, params: any = {}) {
    await delay(400)
    
    let posts = mockPosts.filter(post => post.userId === userId)
    
    // 排序
    if (params.sort === 'latest') {
      posts.sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime())
    }
    
    // 分页
    const page = params.page || 1
    const pageSize = params.pageSize || 10
    const start = (page - 1) * pageSize
    const end = start + pageSize
    
    return {
      success: true,
      data: {
        posts: posts.slice(start, end),
        total: posts.length,
        page,
        pageSize
      }
    }
  },

  // 评论点赞/取消点赞
  async toggleCommentLike(postId: string, commentId: string, userId: string) {
    await delay(300)
    
    const post = mockPosts.find(p => p.id === postId)
    if (!post) {
      throw new Error('动态不存在')
    }
    
    const comment = post.commentsList.find(c => c.id === commentId)
    if (!comment) {
      throw new Error('评论不存在')
    }
    
    // 初始化点赞相关字段
    if (!comment.likes) comment.likes = 0
    if (!comment.isLiked) comment.isLiked = false
    
    if (comment.isLiked) {
      comment.likes--
      comment.isLiked = false
    } else {
      comment.likes++
      comment.isLiked = true
    }
    
    return {
      success: true,
      data: {
        likes: comment.likes,
        isLiked: comment.isLiked
      }
    }
  },

  // 添加回复
  async addReply(postId: string, replyData: any) {
    await delay(500)
    
    const post = mockPosts.find(p => p.id === postId)
    if (!post) {
      throw new Error('动态不存在')
    }
    
    const parentComment = post.commentsList.find(c => c.id === replyData.parentId)
    if (!parentComment) {
      throw new Error('父评论不存在')
    }
    
    const newReply = {
      id: Date.now().toString(),
      userId: replyData.userId,
      userAvatar: replyData.userAvatar,
      userName: replyData.userName,
      userRole: replyData.userRole,
      content: replyData.content,
      likes: 0,
      isLiked: false,
      createdAt: new Date().toISOString()
    }
    
    // 初始化回复列表
    if (!parentComment.replies) {
      parentComment.replies = []
    }
    
    parentComment.replies.unshift(newReply)
    
    return {
      success: true,
      data: newReply,
      message: '回复发布成功'
    }
  },

  // 删除评论
  async deleteComment(postId: string, commentId: string, userId: string) {
    await delay(400)
    
    const post = mockPosts.find(p => p.id === postId)
    if (!post) {
      throw new Error('动态不存在')
    }
    
    const commentIndex = post.commentsList.findIndex(c => c.id === commentId && c.userId === userId)
    if (commentIndex === -1) {
      throw new Error('评论不存在或无权限删除')
    }
    
    post.commentsList.splice(commentIndex, 1)
    post.comments--
    
    return {
      success: true,
      message: '评论删除成功'
    }
  },

  // 删除回复
  async deleteReply(postId: string, replyId: string, userId: string) {
    await delay(400)
    
    const post = mockPosts.find(p => p.id === postId)
    if (!post) {
      throw new Error('动态不存在')
    }
    
    let replyFound = false
    for (const comment of post.commentsList) {
      if (comment.replies) {
        const replyIndex = comment.replies.findIndex(r => r.id === replyId && r.userId === userId)
        if (replyIndex > -1) {
          comment.replies.splice(replyIndex, 1)
          replyFound = true
          break
        }
      }
    }
    
    if (!replyFound) {
      throw new Error('回复不存在或无权限删除')
    }
    
    return {
      success: true,
      message: '回复删除成功'
    }
  },

  // 回复点赞/取消点赞
  async toggleReplyLike(postId: string, replyId: string, userId: string) {
    await delay(300)
    
    const post = mockPosts.find(p => p.id === postId)
    if (!post) {
      throw new Error('动态不存在')
    }
    
    let reply = null
    for (const comment of post.commentsList) {
      if (comment.replies) {
        reply = comment.replies.find(r => r.id === replyId)
        if (reply) break
      }
    }
    
    if (!reply) {
      throw new Error('回复不存在')
    }
    
    // 初始化点赞相关字段
    if (!reply.likes) reply.likes = 0
    if (!reply.isLiked) reply.isLiked = false
    
    if (reply.isLiked) {
      reply.likes--
      reply.isLiked = false
    } else {
      reply.likes++
      reply.isLiked = true
    }
    
    return {
      success: true,
      data: {
        likes: reply.likes,
        isLiked: reply.isLiked
      }
    }
  }
} 