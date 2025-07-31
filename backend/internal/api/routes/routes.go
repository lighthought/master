package routes

import (
	"master-guide-backend/internal/api/handlers"
	"master-guide-backend/pkg/config"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes(engine *gin.Engine, cfg *config.Config, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, mentorHandler *handlers.MentorHandler) {
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
			courses.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get courses list - TODO"})
			})
			courses.GET("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get course detail - TODO"})
			})
			courses.POST("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Create course - TODO"})
			})
			courses.PUT("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Update course - TODO"})
			})
			courses.POST("/:id/enroll", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Enroll course - TODO"})
			})
		}

		// 预约相关路由
		appointments := v1.Group("/appointments")
		{
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
		}

		// 圈子相关路由
		circles := v1.Group("/circles")
		{
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

		// 动态相关路由
		posts := v1.Group("/posts")
		{
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
