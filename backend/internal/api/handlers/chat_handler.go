package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// ChatHandler 聊天处理器
type ChatHandler struct {
	chatService service.ChatService
}

// NewChatHandler 创建聊天处理器
func NewChatHandler(chatService service.ChatService) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
	}
}

// GetOnlineUsers 获取在线用户
// @Summary 获取在线用户
// @Description 获取当前在线用户列表
// @Tags 聊天
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.OnlineUsersResponse}
// @Failure 500 {object} model.ErrorResponse
// @Router /chat/online-users [get]
func (h *ChatHandler) GetOnlineUsers(c *gin.Context) {
	// 获取在线用户
	result, err := h.chatService.GetOnlineUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:      500,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "success",
		Data:      result,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetChatMessages 获取聊天记录
// @Summary 获取聊天记录
// @Description 获取私聊或圈子聊天记录
// @Tags 聊天
// @Accept json
// @Produce json
// @Param target_id query string false "目标用户ID（私聊时使用）"
// @Param circle_id query string false "圈子ID（圈子聊天时使用）"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(50)
// @Success 200 {object} model.Response{data=model.ChatMessagesResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /chat/messages [get]
func (h *ChatHandler) GetChatMessages(c *gin.Context) {
	var req model.ChatMessagesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 验证参数：必须指定 target_id 或 circle_id 之一
	if req.TargetID == "" && req.CircleID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "必须指定 target_id 或 circle_id",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 获取聊天记录
	result, err := h.chatService.GetChatMessages(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:      500,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "success",
		Data:      result,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
