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
	UserRepository        repository.UserRepository
	IdentityRepository    repository.IdentityRepository
	ProfileRepository     repository.ProfileRepository
	PreferencesRepository repository.PreferencesRepository
	MentorRepository      repository.MentorRepository

	// Services
	AuthService   service.AuthService
	UserService   service.UserService
	MentorService service.MentorService

	// Handlers
	AuthHandler   *handlers.AuthHandler
	UserHandler   *handlers.UserHandler
	MentorHandler *handlers.MentorHandler
}

// NewContainer 创建依赖注入容器
func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
	// 初始化Repositories
	userRepo := repository.NewUserRepository(db)
	identityRepo := repository.NewIdentityRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	preferencesRepo := repository.NewPreferencesRepository(db)
	mentorRepo := repository.NewMentorRepository(db)

	// 初始化Services
	authService := service.NewAuthService(userRepo, identityRepo, cfg.JWT.Secret, cfg.JWT.ExpireHours)
	userService := service.NewUserService(userRepo, identityRepo, profileRepo, preferencesRepo)
	mentorService := service.NewMentorService(mentorRepo)

	// 初始化Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	mentorHandler := handlers.NewMentorHandler(mentorService)

	return &Container{
		UserRepository:        userRepo,
		IdentityRepository:    identityRepo,
		ProfileRepository:     profileRepo,
		PreferencesRepository: preferencesRepo,
		MentorRepository:      mentorRepo,
		AuthService:           authService,
		UserService:           userService,
		MentorService:         mentorService,
		AuthHandler:           authHandler,
		UserHandler:           userHandler,
		MentorHandler:         mentorHandler,
	}
}
