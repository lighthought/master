package handlers

import (
	"net/http"
	"strconv"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// StudentHandler 学生处理器
type StudentHandler struct {
	studentService service.StudentService
}

// NewStudentHandler 创建学生处理器实例
func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
	}
}

// GetStudents 获取学生列表
// @Summary 获取学生列表
// @Description 获取导师的学生列表，支持分页、搜索和筛选
// @Tags 学生管理
// @Accept json
// @Produce json
// @Param status query string false "学生状态 (active, inactive, graduated)"
// @Param course_id query string false "课程ID"
// @Param search query string false "搜索关键词（姓名、邮箱）"
// @Param sort_by query string false "排序方式 (name, enrollment_date, progress)"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.StudentListResponse}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /students [get]
func (h *StudentHandler) GetStudents(c *gin.Context) {
	// 从JWT中获取导师ID
	mentorID := c.GetString("user_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:    401,
			Message: "未授权访问",
		})
		return
	}

	// 解析请求参数
	var req model.StudentListRequest
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

	// 调用服务
	response, err := h.studentService.GetStudents(c.Request.Context(), mentorID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    response,
	})
}

// GetStudentByID 获取学生详情
// @Summary 获取学生详情
// @Description 根据学生ID获取学生详细信息
// @Tags 学生管理
// @Accept json
// @Produce json
// @Param student_id path string true "学生ID"
// @Success 200 {object} model.Response{data=model.StudentDetailResponse}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /students/{student_id} [get]
func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	studentID := c.Param("student_id")
	if studentID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "学生ID不能为空",
		})
		return
	}

	response, err := h.studentService.GetStudentByID(c.Request.Context(), studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    response,
	})
}

// GetStudentStats 获取学生统计
// @Summary 获取学生统计
// @Description 获取导师的学生统计数据
// @Tags 学生管理
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.StudentStatsResponse}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /students/stats [get]
func (h *StudentHandler) GetStudentStats(c *gin.Context) {
	// 从JWT中获取导师ID
	mentorID := c.GetString("user_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:    401,
			Message: "未授权访问",
		})
		return
	}

	response, err := h.studentService.GetStudentStats(c.Request.Context(), mentorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    response,
	})
}

// SendMessage 发送消息给学生
// @Summary 发送消息给学生
// @Description 导师向学生发送消息
// @Tags 学生管理
// @Accept json
// @Produce json
// @Param student_id path string true "学生ID"
// @Param request body model.SendMessageRequest true "发送消息请求"
// @Success 200 {object} model.Response{data=model.SendMessageResponse}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /students/{student_id}/messages [post]
func (h *StudentHandler) SendMessage(c *gin.Context) {
	// 从JWT中获取导师ID
	mentorID := c.GetString("user_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:    401,
			Message: "未授权访问",
		})
		return
	}

	studentID := c.Param("student_id")
	if studentID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "学生ID不能为空",
		})
		return
	}

	var req model.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	response, err := h.studentService.SendMessage(c.Request.Context(), mentorID, studentID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "消息发送成功",
		Data:    response,
	})
}

// GetMessages 获取与学生聊天记录
// @Summary 获取与学生聊天记录
// @Description 获取导师与学生的聊天记录
// @Tags 学生管理
// @Accept json
// @Produce json
// @Param student_id path string true "学生ID"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(50)
// @Success 200 {object} model.Response{data=model.MessageListResponse}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /students/{student_id}/messages [get]
func (h *StudentHandler) GetMessages(c *gin.Context) {
	// 从JWT中获取导师ID
	mentorID := c.GetString("user_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:    401,
			Message: "未授权访问",
		})
		return
	}

	studentID := c.Param("student_id")
	if studentID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "学生ID不能为空",
		})
		return
	}

	// 解析分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))

	response, err := h.studentService.GetMessages(c.Request.Context(), mentorID, studentID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    response,
	})
}

// UpdateStudentProgress 更新学生学习进度
// @Summary 更新学生学习进度
// @Description 导师更新学生的学习进度
// @Tags 学生管理
// @Accept json
// @Produce json
// @Param student_id path string true "学生ID"
// @Param course_id path string true "课程ID"
// @Param request body model.StudentProgressRequest true "更新进度请求"
// @Success 200 {object} model.Response{data=model.StudentProgressResponse}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /students/{student_id}/courses/{course_id}/progress [put]
func (h *StudentHandler) UpdateStudentProgress(c *gin.Context) {
	// 从JWT中获取导师ID
	mentorID := c.GetString("user_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:    401,
			Message: "未授权访问",
		})
		return
	}

	studentID := c.Param("student_id")
	courseID := c.Param("course_id")
	if studentID == "" || courseID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "学生ID和课程ID不能为空",
		})
		return
	}

	var req model.StudentProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	response, err := h.studentService.UpdateStudentProgress(c.Request.Context(), mentorID, studentID, courseID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "学习进度更新成功",
		Data:    response,
	})
}

// GradeAssignment 评价学生作业
// @Summary 评价学生作业
// @Description 导师评价学生的作业
// @Tags 学生管理
// @Accept json
// @Produce json
// @Param student_id path string true "学生ID"
// @Param assignment_id path string true "作业ID"
// @Param request body model.GradeAssignmentRequest true "评价作业请求"
// @Success 200 {object} model.Response{data=model.GradeAssignmentResponse}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /students/{student_id}/assignments/{assignment_id}/grade [post]
func (h *StudentHandler) GradeAssignment(c *gin.Context) {
	// 从JWT中获取导师ID
	mentorID := c.GetString("user_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:    401,
			Message: "未授权访问",
		})
		return
	}

	studentID := c.Param("student_id")
	assignmentID := c.Param("assignment_id")
	if studentID == "" || assignmentID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "学生ID和作业ID不能为空",
		})
		return
	}

	var req model.GradeAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	response, err := h.studentService.GradeAssignment(c.Request.Context(), mentorID, studentID, assignmentID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "作业评价成功",
		Data:    response,
	})
}

// GetStudentReport 获取学生学习报告
// @Summary 获取学生学习报告
// @Description 获取学生的学习报告
// @Tags 学生管理
// @Accept json
// @Produce json
// @Param student_id path string true "学生ID"
// @Param period query string true "报告周期 (week, month, quarter, year)"
// @Success 200 {object} model.Response{data=model.StudentReportResponse}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /students/{student_id}/report [get]
func (h *StudentHandler) GetStudentReport(c *gin.Context) {
	// 从JWT中获取导师ID
	mentorID := c.GetString("user_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:    401,
			Message: "未授权访问",
		})
		return
	}

	studentID := c.Param("student_id")
	if studentID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "学生ID不能为空",
		})
		return
	}

	var req model.StudentReportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	response, err := h.studentService.GetStudentReport(c.Request.Context(), mentorID, studentID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    response,
	})
}
