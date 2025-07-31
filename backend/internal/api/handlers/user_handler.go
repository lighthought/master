package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUserProfile 获取用户档案
// @Summary 获取用户档案
// @Description 获取当前用户的档案信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.UserProfileResponse}
// @Failure 401 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /users/profile [get]
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	// 从JWT中获取用户ID（这里简化处理，实际应该从JWT中解析）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.userService.GetUserProfile(c.Request.Context(), userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "用户不存在" || err.Error() == "用户没有活跃身份" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, model.Response{
			Code:      statusCode,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "success",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// UpdateUserProfile 更新用户档案
// @Summary 更新用户档案
// @Description 更新当前用户的档案信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body model.UpdateProfileRequest true "更新档案请求"
// @Security Bearer
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /users/profile [put]
func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	var req model.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从JWT中获取用户ID和身份ID（这里简化处理）
	userID := c.GetString("user_id")
	identityID := c.GetString("identity_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	err := h.userService.UpdateUserProfile(c.Request.Context(), userID, identityID, &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "身份不存在" || err.Error() == "无权访问此身份" {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, model.Response{
			Code:      statusCode,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "用户档案更新成功",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetUserIdentities 获取用户身份列表
// @Summary 获取用户身份列表
// @Description 获取当前用户的所有身份信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.IdentityListResponse}
// @Failure 401 {object} model.ErrorResponse
// @Router /users/identities [get]
func (h *UserHandler) GetUserIdentities(c *gin.Context) {
	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.userService.GetUserIdentities(c.Request.Context(), userID)
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
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// CreateUserIdentity 创建用户身份
// @Summary 创建用户身份
// @Description 为当前用户创建新的身份
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body model.CreateIdentityRequest true "创建身份请求"
// @Security Bearer
// @Success 200 {object} model.Response{data=model.CreateIdentityResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /users/identities [post]
func (h *UserHandler) CreateUserIdentity(c *gin.Context) {
	var req model.CreateIdentityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.userService.CreateUserIdentity(c.Request.Context(), userID, &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "已存在相同类型和领域的身份" {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, model.Response{
			Code:      statusCode,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "身份创建成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// UpdateUserIdentity 更新用户身份
// @Summary 更新用户身份
// @Description 更新指定身份的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param identity_id path string true "身份ID"
// @Param request body model.UpdateIdentityRequest true "更新身份请求"
// @Security Bearer
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /users/identities/{identity_id} [put]
func (h *UserHandler) UpdateUserIdentity(c *gin.Context) {
	identityID := c.Param("identity_id")
	if identityID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "身份ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	var req model.UpdateIdentityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	err := h.userService.UpdateUserIdentity(c.Request.Context(), userID, identityID, &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "身份不存在" || err.Error() == "无权访问此身份" || err.Error() == "档案不存在" {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, model.Response{
			Code:      statusCode,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "身份信息更新成功",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetLearningStats 获取学习统计
// @Summary 获取学习统计
// @Description 获取当前用户的学习统计数据
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.LearningStatsResponse}
// @Failure 401 {object} model.ErrorResponse
// @Router /users/stats/learning [get]
func (h *UserHandler) GetLearningStats(c *gin.Context) {
	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.userService.GetLearningStats(c.Request.Context(), userID)
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
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetTeachingStats 获取教学统计
// @Summary 获取教学统计
// @Description 获取当前用户的教学统计数据
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.TeachingStatsResponse}
// @Failure 401 {object} model.ErrorResponse
// @Router /users/stats/teaching [get]
func (h *UserHandler) GetTeachingStats(c *gin.Context) {
	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.userService.GetTeachingStats(c.Request.Context(), userID)
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
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetGeneralStats 获取通用统计
// @Summary 获取通用统计
// @Description 获取当前用户的通用统计数据
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.GeneralStatsResponse}
// @Failure 401 {object} model.ErrorResponse
// @Router /users/stats/general [get]
func (h *UserHandler) GetGeneralStats(c *gin.Context) {
	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.userService.GetGeneralStats(c.Request.Context(), userID)
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
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetUserAchievements 获取用户成就
// @Summary 获取用户成就
// @Description 获取当前用户的成就列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param identity_type query string false "身份类型"
// @Security Bearer
// @Success 200 {object} model.Response{data=model.AchievementsResponse}
// @Failure 401 {object} model.ErrorResponse
// @Router /users/achievements [get]
func (h *UserHandler) GetUserAchievements(c *gin.Context) {
	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	identityType := c.Query("identity_type")

	response, err := h.userService.GetUserAchievements(c.Request.Context(), userID, identityType)
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
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetUserPreferences 获取用户偏好
// @Summary 获取用户偏好
// @Description 获取当前用户的偏好设置
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.UserPreferencesResponse}
// @Failure 401 {object} model.ErrorResponse
// @Router /users/preferences [get]
func (h *UserHandler) GetUserPreferences(c *gin.Context) {
	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.userService.GetUserPreferences(c.Request.Context(), userID)
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
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// SaveUserPreferences 保存用户偏好
// @Summary 保存用户偏好
// @Description 保存当前用户的偏好设置
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body model.UserPreferencesRequest true "用户偏好请求"
// @Security Bearer
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /users/preferences [put]
func (h *UserHandler) SaveUserPreferences(c *gin.Context) {
	var req model.UserPreferencesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	err := h.userService.SaveUserPreferences(c.Request.Context(), userID, &req)
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
		Message:   "用户偏好保存成功",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetRecommendedLearningPath 获取推荐学习路径
// @Summary 获取推荐学习路径
// @Description 获取基于用户偏好的推荐学习路径
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.RecommendedLearningPathResponse}
// @Failure 401 {object} model.ErrorResponse
// @Router /users/recommended-learning-path [get]
func (h *UserHandler) GetRecommendedLearningPath(c *gin.Context) {
	// 从JWT中获取用户ID（这里简化处理）
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.userService.GetRecommendedLearningPath(c.Request.Context(), userID)
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
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetLearningPathStats 获取学习路径统计
// @Summary 获取学习路径统计
// @Description 获取学习路径的统计数据
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.LearningPathStatsResponse}
// @Failure 401 {object} model.ErrorResponse
// @Router /users/learning-path-stats [get]
func (h *UserHandler) GetLearningPathStats(c *gin.Context) {
	response, err := h.userService.GetLearningPathStats(c.Request.Context())
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
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
