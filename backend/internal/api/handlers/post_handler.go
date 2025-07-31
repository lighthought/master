package handlers

import (
	"net/http"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// PostHandler 动态处理器
type PostHandler struct {
	postService service.PostService
}

// NewPostHandler 创建动态处理器
func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

// GetPosts 获取圈子动态
// @Summary 获取圈子动态
// @Description 获取指定圈子的动态列表
// @Tags 动态管理
// @Accept json
// @Produce json
// @Param circle_id path string true "圈子ID"
// @Param post_type query string false "动态类型"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.PostListResponse}
// @Router /circles/{circle_id}/posts [get]
func (h *PostHandler) GetPosts(c *gin.Context) {
	circleID := c.Param("circle_id")
	var req model.PostListRequest
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

	resp, err := h.postService.GetPosts(c.Request.Context(), circleID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取动态列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// CreatePost 发布动态
// @Summary 发布动态
// @Description 在指定圈子发布动态
// @Tags 动态管理
// @Accept json
// @Produce json
// @Param circle_id path string true "圈子ID"
// @Param request body model.CreatePostRequest true "动态内容"
// @Success 200 {object} model.Response{data=model.CreatePostResponse}
// @Router /circles/{circle_id}/posts [post]
func (h *PostHandler) CreatePost(c *gin.Context) {
	circleID := c.Param("circle_id")
	userID := c.GetString("user_id")         // 从JWT中获取
	identityID := c.GetString("identity_id") // 从JWT中获取

	var req model.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.postService.CreatePost(c.Request.Context(), userID, identityID, circleID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "发布动态失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "动态发布成功",
		Data:    resp,
	})
}

// LikePost 点赞动态
// @Summary 点赞动态
// @Description 点赞指定动态
// @Tags 动态管理
// @Accept json
// @Produce json
// @Param post_id path string true "动态ID"
// @Success 200 {object} model.Response{data=model.LikePostResponse}
// @Router /posts/{post_id}/like [post]
func (h *PostHandler) LikePost(c *gin.Context) {
	postID := c.Param("post_id")
	userID := c.GetString("user_id") // 从JWT中获取

	resp, err := h.postService.LikePost(c.Request.Context(), userID, postID)
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

// UnlikePost 取消点赞动态
// @Summary 取消点赞动态
// @Description 取消点赞指定动态
// @Tags 动态管理
// @Accept json
// @Produce json
// @Param post_id path string true "动态ID"
// @Success 200 {object} model.Response{data=model.LikePostResponse}
// @Router /posts/{post_id}/like [delete]
func (h *PostHandler) UnlikePost(c *gin.Context) {
	postID := c.Param("post_id")
	userID := c.GetString("user_id") // 从JWT中获取

	resp, err := h.postService.UnlikePost(c.Request.Context(), userID, postID)
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
