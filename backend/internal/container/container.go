package container

import (
	"master-guide-backend/internal/api/handlers"
	"master-guide-backend/internal/repository"
	"master-guide-backend/internal/service"
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

	// Services
	AuthService        service.AuthService
	UserService        service.UserService
	MentorService      service.MentorService
	CourseService      service.CourseService
	AppointmentService service.AppointmentService

	// Handlers
	AuthHandler        *handlers.AuthHandler
	UserHandler        *handlers.UserHandler
	MentorHandler      *handlers.MentorHandler
	CourseHandler      *handlers.CourseHandler
	AppointmentHandler *handlers.AppointmentHandler
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

	// 初始化Services
	authService := service.NewAuthService(userRepo, identityRepo, cfg.JWT.Secret, cfg.JWT.ExpireHours)
	userService := service.NewUserService(userRepo, identityRepo, profileRepo, preferencesRepo)
	mentorService := service.NewMentorService(mentorRepo)
	courseService := service.NewCourseService(courseRepo, courseContentRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo, mentorRepo)

	// 初始化Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	mentorHandler := handlers.NewMentorHandler(mentorService)
	courseHandler := handlers.NewCourseHandler(courseService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)

	return &Container{
		UserRepository:          userRepo,
		IdentityRepository:      identityRepo,
		ProfileRepository:       profileRepo,
		PreferencesRepository:   preferencesRepo,
		MentorRepository:        mentorRepo,
		CourseRepository:        courseRepo,
		CourseContentRepository: courseContentRepo,
		AppointmentRepository:   appointmentRepo,
		AuthService:             authService,
		UserService:             userService,
		MentorService:           mentorService,
		CourseService:           courseService,
		AppointmentService:      appointmentService,
		AuthHandler:             authHandler,
		UserHandler:             userHandler,
		MentorHandler:           mentorHandler,
		CourseHandler:           courseHandler,
		AppointmentHandler:      appointmentHandler,
	}
}
