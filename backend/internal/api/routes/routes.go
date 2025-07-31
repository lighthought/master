package routes

import (
	"master-guide-backend/internal/api/handlers"
	"master-guide-backend/pkg/config"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes(engine *gin.Engine, cfg *config.Config, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, mentorHandler *handlers.MentorHandler, courseHandler *handlers.CourseHandler, appointmentHandler *handlers.AppointmentHandler, circleHandler *handlers.CircleHandler, postHandler *handlers.PostHandler, commentHandler *handlers.CommentHandler, reviewHandler *handlers.ReviewHandler) {
	// API v1 路由组
	v1 := engine.Group("/api/v1")
	{
		// 健康检查
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Master Guide API is running",
			})
		})

		// 认证相关路由
		auth := v1.Group("/auth")
		{
			if authHandler != nil {
				auth.POST("/register", authHandler.Register)
				auth.POST("/login", authHandler.Login)
				auth.POST("/refresh", authHandler.RefreshToken)
				auth.POST("/switch-identity", authHandler.SwitchIdentity)
				auth.POST("/change-password", authHandler.ChangePassword)
			} else {
				auth.POST("/register", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Register endpoint - TODO"})
				})
				auth.POST("/login", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Login endpoint - TODO"})
				})
				auth.POST("/refresh", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Refresh token endpoint - TODO"})
				})
			}
		}

		// 用户相关路由
		users := v1.Group("/users")
		{
			if userHandler != nil {
				users.GET("/profile", userHandler.GetUserProfile)
				users.PUT("/profile", userHandler.UpdateUserProfile)
				users.GET("/identities", userHandler.GetUserIdentities)
				users.POST("/identities", userHandler.CreateUserIdentity)
				users.PUT("/identities/:identity_id", userHandler.UpdateUserIdentity)
				users.GET("/stats/learning", userHandler.GetLearningStats)
				users.GET("/stats/teaching", userHandler.GetTeachingStats)
				users.GET("/stats/general", userHandler.GetGeneralStats)
				users.GET("/achievements", userHandler.GetUserAchievements)
				users.GET("/preferences", userHandler.GetUserPreferences)
				users.PUT("/preferences", userHandler.SaveUserPreferences)
				users.GET("/recommended-learning-path", userHandler.GetRecommendedLearningPath)
				users.GET("/learning-path-stats", userHandler.GetLearningPathStats)
			} else {
				users.GET("/profile", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get user profile - TODO"})
				})
				users.PUT("/profile", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Update user profile - TODO"})
				})
				users.GET("/identities", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get user identities - TODO"})
				})
				users.POST("/identities", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Create user identity - TODO"})
				})
			}
		}

		// 大师相关路由
		mentors := v1.Group("/mentors")
		{
			if mentorHandler != nil {
				mentors.GET("", mentorHandler.GetMentors)
				mentors.GET("/:mentor_id", mentorHandler.GetMentorDetail)
				mentors.GET("/search", mentorHandler.SearchMentors)
				mentors.GET("/recommended", mentorHandler.GetRecommendedMentors)
				mentors.GET("/:mentor_id/reviews", mentorHandler.GetMentorReviews)
			} else {
				mentors.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get mentors list - TODO"})
				})
				mentors.GET("/:id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get mentor detail - TODO"})
				})
				mentors.GET("/search", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Search mentors - TODO"})
				})
				mentors.GET("/recommended", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get recommended mentors - TODO"})
				})
			}
		}

		// 课程相关路由
		courses := v1.Group("/courses")
		{
			if courseHandler != nil {
				courses.GET("", courseHandler.GetCourses)
				courses.GET("/:course_id", courseHandler.GetCourseDetail)
				courses.POST("", courseHandler.CreateCourse)
				courses.POST("/:course_id/enroll", courseHandler.EnrollCourse)
				courses.GET("/:course_id/progress", courseHandler.GetCourseProgress)
				courses.GET("/search", courseHandler.SearchCourses)
				courses.GET("/recommended", courseHandler.GetRecommendedCourses)
				courses.GET("/enrolled", courseHandler.GetEnrolledCourses)
			} else {
				courses.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get courses list - TODO"})
				})
				courses.GET("/:id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get course detail - TODO"})
				})
				courses.POST("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Create course - TODO"})
				})
				courses.POST("/:id/enroll", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Enroll course - TODO"})
				})
				courses.GET("/:id/progress", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get course progress - TODO"})
				})
				courses.GET("/search", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Search courses - TODO"})
				})
				courses.GET("/recommended", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get recommended courses - TODO"})
				})
				courses.GET("/enrolled", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get enrolled courses - TODO"})
				})
			}
		}

		// 预约相关路由
		appointments := v1.Group("/appointments")
		{
			if appointmentHandler != nil {
				appointments.GET("", appointmentHandler.GetAppointments)
				appointments.POST("", appointmentHandler.CreateAppointment)
				appointments.GET("/:appointment_id", appointmentHandler.GetAppointmentDetail)
				appointments.PUT("/:appointment_id/status", appointmentHandler.UpdateAppointmentStatus)
				appointments.DELETE("/:appointment_id", appointmentHandler.CancelAppointment)
				appointments.GET("/mentor-stats", appointmentHandler.GetMentorAppointmentStats)
			} else {
				appointments.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get appointments list - TODO"})
				})
				appointments.POST("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Create appointment - TODO"})
				})
				appointments.GET("/:id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get appointment detail - TODO"})
				})
				appointments.PUT("/:id/status", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Update appointment status - TODO"})
				})
				appointments.DELETE("/:id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Cancel appointment - TODO"})
				})
				appointments.GET("/mentor-stats", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get mentor appointment stats - TODO"})
				})
			}
		}

		// 圈子相关路由
		circles := v1.Group("/circles")
		{
			if circleHandler != nil {
				circles.GET("", circleHandler.GetCircles)
				circles.GET("/recommended", circleHandler.GetRecommendedCircles)
				circles.POST("/:circle_id/join", circleHandler.JoinCircle)
				circles.DELETE("/:circle_id/join", circleHandler.LeaveCircle)
			} else {
				circles.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get circles list - TODO"})
				})
				circles.GET("/:id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get circle detail - TODO"})
				})
				circles.POST("/:id/join", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Join circle - TODO"})
				})
				circles.DELETE("/:id/join", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Leave circle - TODO"})
				})
			}
		}

		// 圈子动态相关路由
		circlePosts := v1.Group("/circles/:circle_id/posts")
		{
			if postHandler != nil {
				circlePosts.GET("", postHandler.GetPosts)
				circlePosts.POST("", postHandler.CreatePost)
			} else {
				circlePosts.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get circle posts - TODO"})
				})
				circlePosts.POST("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Create circle post - TODO"})
				})
			}
		}

		// 动态相关路由
		posts := v1.Group("/posts")
		{
			if postHandler != nil {
				posts.POST("/:post_id/like", postHandler.LikePost)
				posts.DELETE("/:post_id/like", postHandler.UnlikePost)
			} else {
				posts.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get posts list - TODO"})
				})
				posts.POST("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Create post - TODO"})
				})
				posts.GET("/:id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get post detail - TODO"})
				})
				posts.POST("/:id/like", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Like post - TODO"})
				})
				posts.DELETE("/:id/like", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Unlike post - TODO"})
				})
			}
		}

		// 动态评论相关路由
		postComments := v1.Group("/posts/:post_id/comments")
		{
			if commentHandler != nil {
				postComments.GET("", commentHandler.GetComments)
				postComments.POST("", commentHandler.CreateComment)
			} else {
				postComments.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get post comments - TODO"})
				})
				postComments.POST("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Create post comment - TODO"})
				})
			}
		}

		// 评论相关路由
		comments := v1.Group("/comments")
		{
			if commentHandler != nil {
				comments.POST("/:comment_id/replies", commentHandler.CreateReply)
				comments.POST("/:comment_id/like", commentHandler.LikeComment)
				comments.DELETE("/:comment_id/like", commentHandler.UnlikeComment)
				comments.DELETE("/:comment_id", commentHandler.DeleteComment)
			} else {
				comments.POST("/:comment_id/replies", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Create comment reply - TODO"})
				})
				comments.POST("/:comment_id/like", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Like comment - TODO"})
				})
				comments.DELETE("/:comment_id/like", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Unlike comment - TODO"})
				})
				comments.DELETE("/:comment_id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Delete comment - TODO"})
				})
			}
		}

		// 评价相关路由
		reviews := v1.Group("/reviews")
		{
			if reviewHandler != nil {
				reviews.GET("", reviewHandler.GetReviews)
				reviews.GET("/:review_id", reviewHandler.GetReviewByID)
				reviews.POST("", reviewHandler.CreateReview)
				reviews.PUT("/:review_id", reviewHandler.UpdateReview)
				reviews.DELETE("/:review_id", reviewHandler.DeleteReview)
				reviews.GET("/stats", reviewHandler.GetReviewStats)
			} else {
				reviews.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get reviews list - TODO"})
				})
				reviews.GET("/:id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get review detail - TODO"})
				})
				reviews.POST("", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Create review - TODO"})
				})
				reviews.PUT("/:id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Update review - TODO"})
				})
				reviews.DELETE("/:id", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Delete review - TODO"})
				})
				reviews.GET("/stats", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Get review stats - TODO"})
				})
			}
		}

		// 学习记录相关路由
		learning := v1.Group("/learning-records")
		{
			learning.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get learning records - TODO"})
			})
			learning.PUT("/:id/progress", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Update learning progress - TODO"})
			})
		}

		// 收入相关路由
		income := v1.Group("/income")
		{
			income.GET("/stats", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get income stats - TODO"})
			})
			income.GET("/transactions", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get income transactions - TODO"})
			})
		}

		// 支付相关路由
		payments := v1.Group("/payments")
		{
			payments.POST("/orders", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Create payment order - TODO"})
			})
			payments.GET("/orders/:id/status", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get payment status - TODO"})
			})
		}

		// 学生相关路由
		students := v1.Group("/students")
		{
			students.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get students list - TODO"})
			})
			students.GET("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get student detail - TODO"})
			})
		}

		// 文件上传路由
		upload := v1.Group("/upload")
		{
			upload.POST("/file", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Upload file - TODO"})
			})
		}

		// 搜索路由
		search := v1.Group("/search")
		{
			search.GET("/global", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Global search - TODO"})
			})
		}

		// 统计路由
		stats := v1.Group("/stats")
		{
			stats.GET("/system", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get system stats - TODO"})
			})
		}
	}

	// WebSocket路由
	engine.GET("/ws", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "WebSocket endpoint - TODO"})
	})
}
