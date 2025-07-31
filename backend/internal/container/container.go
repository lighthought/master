package container

import (
	"master-guide-backend/internal/api/handlers"
	"master-guide-backend/internal/repository"
	"master-guide-backend/internal/service"
	"master-guide-backend/internal/utils"
	"master-guide-backend/pkg/config"

	"gorm.io/gorm"
)

// Container 依赖注入容器
type Container struct {
	// Repositories
	UserRepository          repository.UserRepository
	IdentityRepository      repository.IdentityRepository
	ProfileRepository       repository.ProfileRepository
	PreferencesRepository   repository.PreferencesRepository
	MentorRepository        repository.MentorRepository
	CourseRepository        repository.CourseRepository
	CourseContentRepository repository.CourseContentRepository
	AppointmentRepository   repository.AppointmentRepository
	CircleRepository        repository.CircleRepository
	PostRepository          repository.PostRepository
	CommentRepository       repository.CommentRepository
	ReviewRepository        repository.ReviewRepository
	NotificationRepository  repository.NotificationRepository
	LearningRepository      repository.LearningRepository
	StudentRepository       repository.StudentRepository
	MessageRepository       repository.MessageRepository
	IncomeRepository        repository.IncomeRepository
	PaymentRepository       repository.PaymentRepository
	UploadRepository        repository.UploadRepository
	SearchRepository        repository.SearchRepository
	StatsRepository         repository.StatsRepository
	ChatRepository          repository.ChatRepository

	// Services
	AuthService         service.AuthService
	UserService         service.UserService
	MentorService       service.MentorService
	CourseService       service.CourseService
	AppointmentService  service.AppointmentService
	CircleService       service.CircleService
	PostService         service.PostService
	CommentService      service.CommentService
	ReviewService       service.ReviewService
	NotificationService service.NotificationService
	LearningService     service.LearningService
	StudentService      service.StudentService
	IncomeService       service.IncomeService
	PaymentService      service.PaymentService
	UploadService       service.UploadService
	SearchService       service.SearchService
	StatsService        service.StatsService
	ChatService         service.ChatService

	// Handlers
	AuthHandler         *handlers.AuthHandler
	UserHandler         *handlers.UserHandler
	MentorHandler       *handlers.MentorHandler
	CourseHandler       *handlers.CourseHandler
	AppointmentHandler  *handlers.AppointmentHandler
	CircleHandler       *handlers.CircleHandler
	PostHandler         *handlers.PostHandler
	CommentHandler      *handlers.CommentHandler
	ReviewHandler       *handlers.ReviewHandler
	NotificationHandler *handlers.NotificationHandler
	LearningHandler     *handlers.LearningHandler
	StudentHandler      *handlers.StudentHandler
	IncomeHandler       *handlers.IncomeHandler
	PaymentHandler      *handlers.PaymentHandler
	UploadHandler       *handlers.UploadHandler
	SearchHandler       *handlers.SearchHandler
	StatsHandler        *handlers.StatsHandler
	ChatHandler         *handlers.ChatHandler
	WebSocketHandler    *handlers.WebSocketHandler
}

// NewContainer 创建依赖注入容器
func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
	// 初始化Repositories
	userRepo := repository.NewUserRepository(db)
	identityRepo := repository.NewIdentityRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	preferencesRepo := repository.NewPreferencesRepository(db)
	mentorRepo := repository.NewMentorRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	courseContentRepo := repository.NewCourseContentRepository(db)
	appointmentRepo := repository.NewAppointmentRepository(db)
	circleRepo := repository.NewCircleRepository(db)
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	notificationRepo := repository.NewNotificationRepository(db)
	learningRepo := repository.NewLearningRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	incomeRepo := repository.NewIncomeRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	uploadRepo := repository.NewUploadRepository(db)
	searchRepo := repository.NewSearchRepository(db)
	statsRepo := repository.NewStatsRepository(db)
	chatRepo := repository.NewChatRepository(db)

	// 初始化Services
	authService := service.NewAuthService(userRepo, identityRepo, cfg.JWT.Secret, cfg.JWT.ExpireHours)
	userService := service.NewUserService(userRepo, identityRepo, profileRepo, preferencesRepo, learningRepo, mentorRepo, appointmentRepo)
	mentorService := service.NewMentorService(mentorRepo)
	courseService := service.NewCourseService(courseRepo, courseContentRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo, mentorRepo)
	circleService := service.NewCircleService(circleRepo)
	postService := service.NewPostService(postRepo)
	commentService := service.NewCommentService(commentRepo)
	reviewService := service.NewReviewService(reviewRepo)
	notificationService := service.NewNotificationService(notificationRepo)
	learningService := service.NewLearningService(learningRepo)
	studentService := service.NewStudentService(studentRepo, userRepo, identityRepo, appointmentRepo, messageRepo)
	incomeService := service.NewIncomeService(incomeRepo)
	paymentService := service.NewPaymentService(paymentRepo)
	uploadService := service.NewUploadService(uploadRepo)
	searchService := service.NewSearchService(searchRepo)
	statsService := service.NewStatsService(statsRepo)

	// 创建 WebSocket 管理器
	websocketMgr := utils.NewWebSocketManager()
	go websocketMgr.Start()

	chatService := service.NewChatService(chatRepo, websocketMgr)

	// 初始化Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	mentorHandler := handlers.NewMentorHandler(mentorService)
	courseHandler := handlers.NewCourseHandler(courseService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)
	circleHandler := handlers.NewCircleHandler(circleService)
	postHandler := handlers.NewPostHandler(postService)
	commentHandler := handlers.NewCommentHandler(commentService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)
	learningHandler := handlers.NewLearningHandler(learningService)
	studentHandler := handlers.NewStudentHandler(studentService)
	incomeHandler := handlers.NewIncomeHandler(incomeService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	uploadHandler := handlers.NewUploadHandler(uploadService)
	searchHandler := handlers.NewSearchHandler(searchService)
	statsHandler := handlers.NewStatsHandler(statsService)
	chatHandler := handlers.NewChatHandler(chatService)
	websocketHandler := handlers.NewWebSocketHandler(websocketMgr)

	return &Container{
		UserRepository:          userRepo,
		IdentityRepository:      identityRepo,
		ProfileRepository:       profileRepo,
		PreferencesRepository:   preferencesRepo,
		MentorRepository:        mentorRepo,
		CourseRepository:        courseRepo,
		CourseContentRepository: courseContentRepo,
		AppointmentRepository:   appointmentRepo,
		CircleRepository:        circleRepo,
		PostRepository:          postRepo,
		CommentRepository:       commentRepo,
		ReviewRepository:        reviewRepo,
		NotificationRepository:  notificationRepo,
		LearningRepository:      learningRepo,
		StudentRepository:       studentRepo,
		MessageRepository:       messageRepo,
		IncomeRepository:        incomeRepo,
		PaymentRepository:       paymentRepo,
		UploadRepository:        uploadRepo,
		SearchRepository:        searchRepo,
		StatsRepository:         statsRepo,
		ChatRepository:          chatRepo,
		AuthService:             authService,
		UserService:             userService,
		MentorService:           mentorService,
		CourseService:           courseService,
		AppointmentService:      appointmentService,
		CircleService:           circleService,
		PostService:             postService,
		CommentService:          commentService,
		ReviewService:           reviewService,
		NotificationService:     notificationService,
		LearningService:         learningService,
		StudentService:          studentService,
		IncomeService:           incomeService,
		PaymentService:          paymentService,
		UploadService:           uploadService,
		SearchService:           searchService,
		StatsService:            statsService,
		ChatService:             chatService,
		AuthHandler:             authHandler,
		UserHandler:             userHandler,
		MentorHandler:           mentorHandler,
		CourseHandler:           courseHandler,
		AppointmentHandler:      appointmentHandler,
		CircleHandler:           circleHandler,
		PostHandler:             postHandler,
		CommentHandler:          commentHandler,
		ReviewHandler:           reviewHandler,
		NotificationHandler:     notificationHandler,
		LearningHandler:         learningHandler,
		StudentHandler:          studentHandler,
		IncomeHandler:           incomeHandler,
		PaymentHandler:          paymentHandler,
		UploadHandler:           uploadHandler,
		SearchHandler:           searchHandler,
		StatsHandler:            statsHandler,
		ChatHandler:             chatHandler,
		WebSocketHandler:        websocketHandler,
	}
}
