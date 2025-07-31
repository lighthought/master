package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// SearchHandler 搜索处理器
type SearchHandler struct {
	searchService service.SearchService
}

// NewSearchHandler 创建搜索处理器
func NewSearchHandler(searchService service.SearchService) *SearchHandler {
	return &SearchHandler{
		searchService: searchService,
	}
}

// GlobalSearch 全局搜索
// @Summary 全局搜索
// @Description 搜索导师、课程、帖子等内容
// @Tags 搜索
// @Accept json
// @Produce json
// @Param q query string true "搜索关键词"
// @Param type query string false "搜索类型" Enums(mentors, courses, posts)
// @Param domain query string false "专业领域"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.SearchResult}
// @Failure 400 {object} model.ErrorResponse
// @Router /search [get]
func (h *SearchHandler) GlobalSearch(c *gin.Context) {
	var req model.SearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 验证搜索关键词
	if req.Query == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "搜索关键词不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 执行搜索
	result, err := h.searchService.GlobalSearch(c.Request.Context(), &req)
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
