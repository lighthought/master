package handlers

import (
	"net/http"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// CommentHandler 评论处理器
type CommentHandler struct {
	commentService service.CommentService
}

// NewCommentHandler 创建评论处理器
func NewCommentHandler(commentService service.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

// GetComments 获取评论列表
// @Summary 获取评论列表
// @Description 获取指定动态的评论列表
// @Tags 评论管理
// @Accept json
// @Produce json
// @Param post_id path string true "动态ID"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.CommentListResponse}
// @Router /posts/{post_id}/comments [get]
func (h *CommentHandler) GetComments(c *gin.Context) {
	postID := c.Param("post_id")
	var req model.CommentListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	resp, err := h.commentService.GetComments(c.Request.Context(), postID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取评论列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// CreateComment 发表评论
// @Summary 发表评论
// @Description 对指定动态发表评论
// @Tags 评论管理
// @Accept json
// @Produce json
// @Param post_id path string true "动态ID"
// @Param request body model.CreateCommentRequest true "评论内容"
// @Success 200 {object} model.Response{data=model.CreateCommentResponse}
// @Router /posts/{post_id}/comments [post]
func (h *CommentHandler) CreateComment(c *gin.Context) {
	postID := c.Param("post_id")
	userID := c.GetString("user_id")         // 从JWT中获取
	identityID := c.GetString("identity_id") // 从JWT中获取

	var req model.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.commentService.CreateComment(c.Request.Context(), userID, identityID, postID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "评论发表失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "评论发表成功",
		Data:    resp,
	})
}

// CreateReply 回复评论
// @Summary 回复评论
// @Description 回复指定评论
// @Tags 评论管理
// @Accept json
// @Produce json
// @Param comment_id path string true "评论ID"
// @Param request body model.CreateReplyRequest true "回复内容"
// @Success 200 {object} model.Response{data=model.CreateReplyResponse}
// @Router /comments/{comment_id}/replies [post]
func (h *CommentHandler) CreateReply(c *gin.Context) {
	commentID := c.Param("comment_id")
	userID := c.GetString("user_id")         // 从JWT中获取
	identityID := c.GetString("identity_id") // 从JWT中获取

	var req model.CreateReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.commentService.CreateReply(c.Request.Context(), userID, identityID, commentID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "回复发表失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "回复发表成功",
		Data:    resp,
	})
}

// LikeComment 点赞评论
// @Summary 点赞评论
// @Description 点赞指定评论
// @Tags 评论管理
// @Accept json
// @Produce json
// @Param comment_id path string true "评论ID"
// @Success 200 {object} model.Response{data=model.LikeCommentResponse}
// @Router /comments/{comment_id}/like [post]
func (h *CommentHandler) LikeComment(c *gin.Context) {
	commentID := c.Param("comment_id")
	userID := c.GetString("user_id") // 从JWT中获取

	resp, err := h.commentService.LikeComment(c.Request.Context(), userID, commentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "点赞成功",
		Data:    resp,
	})
}

// UnlikeComment 取消点赞评论
// @Summary 取消点赞评论
// @Description 取消点赞指定评论
// @Tags 评论管理
// @Accept json
// @Produce json
// @Param comment_id path string true "评论ID"
// @Success 200 {object} model.Response{data=model.LikeCommentResponse}
// @Router /comments/{comment_id}/like [delete]
func (h *CommentHandler) UnlikeComment(c *gin.Context) {
	commentID := c.Param("comment_id")
	userID := c.GetString("user_id") // 从JWT中获取

	resp, err := h.commentService.UnlikeComment(c.Request.Context(), userID, commentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "取消点赞成功",
		Data:    resp,
	})
}

// DeleteComment 删除评论
// @Summary 删除评论
// @Description 删除指定评论
// @Tags 评论管理
// @Accept json
// @Produce json
// @Param comment_id path string true "评论ID"
// @Success 200 {object} model.Response{data=model.DeleteCommentResponse}
// @Router /comments/{comment_id} [delete]
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentID := c.Param("comment_id")

	resp, err := h.commentService.DeleteComment(c.Request.Context(), commentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "评论删除成功",
		Data:    resp,
	})
}
