package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// AuthHandler 认证处理器
type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler 创建认证处理器
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册接口
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body model.RegisterRequest true "注册请求"
// @Success 200 {object} model.Response{data=model.AuthResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.authService.Register(c.Request.Context(), &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "邮箱已存在" {
			statusCode = http.StatusConflict
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
		Message:   "注册成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "登录请求"
// @Success 200 {object} model.Response{data=model.AuthResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "用户名或密码错误" || err.Error() == "账户已被禁用" || err.Error() == "用户没有可用身份" || err.Error() == "用户没有活跃身份" {
			statusCode = http.StatusUnauthorized
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
		Message:   "登录成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// SwitchIdentity 身份切换
// @Summary 身份切换
// @Description 用户身份切换接口
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body model.SwitchIdentityRequest true "身份切换请求"
// @Security Bearer
// @Success 200 {object} model.Response{data=model.AuthResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /auth/switch-identity [post]
func (h *AuthHandler) SwitchIdentity(c *gin.Context) {
	var req model.SwitchIdentityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

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

	response, err := h.authService.SwitchIdentity(c.Request.Context(), userID, req.IdentityID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "身份不存在" || err.Error() == "无权访问此身份" || err.Error() == "身份未激活" {
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
		Message:   "身份切换成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// RefreshToken 刷新Token
// @Summary 刷新Token
// @Description 刷新JWT Token接口
// @Tags 认证
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.TokenResponse}
// @Failure 401 {object} model.ErrorResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
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

	response, err := h.authService.RefreshToken(c.Request.Context(), userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "用户不存在" || err.Error() == "账户已被禁用" || err.Error() == "用户没有活跃身份" {
			statusCode = http.StatusUnauthorized
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
		Message:   "Token刷新成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Description 用户修改密码接口
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body model.ChangePasswordRequest true "修改密码请求"
// @Security Bearer
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /auth/change-password [post]
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var req model.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

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

	err := h.authService.ChangePassword(c.Request.Context(), userID, req.CurrentPassword, req.NewPassword)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "用户不存在" || err.Error() == "当前密码错误" {
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
		Message:   "密码修改成功",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
