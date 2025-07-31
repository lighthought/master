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

	// Services
	AuthService service.AuthService
	UserService service.UserService

	// Handlers
	AuthHandler *handlers.AuthHandler
	UserHandler *handlers.UserHandler
}

// NewContainer 创建依赖注入容器
func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
	// 初始化Repositories
	userRepo := repository.NewUserRepository(db)
	identityRepo := repository.NewIdentityRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	preferencesRepo := repository.NewPreferencesRepository(db)

	// 初始化Services
	authService := service.NewAuthService(userRepo, identityRepo, cfg.JWT.Secret, cfg.JWT.ExpireHours)
	userService := service.NewUserService(userRepo, identityRepo, profileRepo, preferencesRepo)

	// 初始化Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	return &Container{
		UserRepository:        userRepo,
		IdentityRepository:    identityRepo,
		ProfileRepository:     profileRepo,
		PreferencesRepository: preferencesRepo,
		AuthService:           authService,
		UserService:           userService,
		AuthHandler:           authHandler,
		UserHandler:           userHandler,
	}
}
