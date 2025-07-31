package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// UploadHandler 文件上传处理器
type UploadHandler struct {
	uploadService service.UploadService
}

// NewUploadHandler 创建文件上传处理器
func NewUploadHandler(uploadService service.UploadService) *UploadHandler {
	return &UploadHandler{
		uploadService: uploadService,
	}
}

// UploadFile 上传文件
// @Summary 上传文件
// @Description 上传文件到服务器，支持头像、课程封面、帖子图片等类型
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "要上传的文件"
// @Param type formData string true "文件类型" Enums(avatar, course_cover, post_image)
// @Success 200 {object} model.Response{data=model.UploadFileResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /upload/file [post]
func (h *UploadHandler) UploadFile(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "获取文件失败",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 获取文件类型
	fileType := c.PostForm("type")
	if fileType == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "文件类型不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 获取用户ID（这里简化处理，实际应该从JWT token中获取）
	userID := c.GetString("user_id")
	if userID == "" {
		userID = "default_user" // 临时默认用户ID
	}

	// 上传文件
	response, err := h.uploadService.UploadFile(c.Request.Context(), file, fileType, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "上传成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
