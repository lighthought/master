# Master Guide 后端架构设计

## 1. 后端技术栈详解

### 1.1 核心技术
- **Go 1.21**：高性能编程语言，支持并发编程
- **Gin**：轻量级Web框架，高性能HTTP路由
- **PostgreSQL**：关系型数据库，支持ACID事务
- **Redis**：内存数据库，用于缓存和会话存储
- **GORM**：Go语言的ORM框架
- **JWT**：无状态身份认证

### 1.2 辅助技术
- **WebSocket**：实时通信支持
- **Validator**：数据验证库
- **Zap**：高性能日志库
- **Viper**：配置管理库
- **Testify**：测试框架
- **Swagger**：API文档生成

## 2. 项目结构设计

```
backend/
├── cmd/                     # 应用入口
│   ├── server/             # 主服务器
│   │   └── main.go
│   ├── migrate/            # 数据库迁移
│   │   └── main.go
│   └── seed/               # 数据种子
│       └── main.go
├── internal/               # 内部包
│   ├── config/            # 配置管理
│   │   ├── config.go
│   │   └── database.go
│   ├── models/            # 数据模型
│   │   ├── user.go
│   │   ├── identity.go
│   │   ├── course.go
│   │   ├── mentor.go
│   │   └── community.go
│   ├── handlers/          # HTTP处理器
│   │   ├── auth.go
│   │   ├── user.go
│   │   ├── course.go
│   │   ├── mentor.go
│   │   └── community.go
│   ├── services/          # 业务服务
│   │   ├── auth_service.go
│   │   ├── user_service.go
│   │   ├── course_service.go
│   │   ├── mentor_service.go
│   │   └── community_service.go
│   ├── middleware/        # 中间件
│   │   ├── auth.go
│   │   ├── cors.go
│   │   ├── logger.go
│   │   └── rate_limit.go
│   ├── repository/        # 数据访问层
│   │   ├── user_repo.go
│   │   ├── course_repo.go
│   │   ├── mentor_repo.go
│   │   └── community_repo.go
│   └── utils/             # 工具函数
│       ├── jwt.go
│       ├── password.go
│       ├── response.go
│       └── validator.go
├── pkg/                   # 公共包
│   ├── database/          # 数据库连接
│   │   └── postgres.go
│   ├── cache/             # 缓存管理
│   │   └── redis.go
│   ├── websocket/         # WebSocket管理
│   │   └── manager.go
│   └── logger/            # 日志管理
│       └── logger.go
├── migrations/            # 数据库迁移文件
│   ├── 001_create_users.sql
│   ├── 002_create_identities.sql
│   ├── 003_create_courses.sql
│   └── 004_create_community.sql
├── docs/                  # 文档
│   ├── api/              # API文档
│   │   └── swagger.json
│   └── architecture/     # 架构文档
├── scripts/              # 脚本文件
│   ├── build.sh
│   ├── deploy.sh
│   └── test.sh
├── go.mod
├── go.sum
├── .env.example
├── docker-compose.yml
└── README.md
```

## 3. 数据库设计

### 3.1 核心数据模型

#### 用户相关表
```sql
-- 用户基础表
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 用户身份表
CREATE TABLE user_identities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_type VARCHAR(20) NOT NULL CHECK (identity_type IN ('master', 'apprentice')),
    domain VARCHAR(100) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    verification_status VARCHAR(20) DEFAULT 'unverified',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, domain, identity_type)
);

-- 用户档案表
CREATE TABLE user_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    avatar VARCHAR(500),
    bio TEXT,
    skills TEXT[],
    experience_years INTEGER DEFAULT 0,
    hourly_rate DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### 课程相关表
```sql
-- 课程表
CREATE TABLE courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mentor_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    cover_image VARCHAR(500),
    price DECIMAL(10,2) NOT NULL,
    duration_hours INTEGER NOT NULL,
    difficulty VARCHAR(20) CHECK (difficulty IN ('beginner', 'intermediate', 'advanced')),
    status VARCHAR(20) DEFAULT 'draft',
    max_students INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 课程内容表
CREATE TABLE course_contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    content_type VARCHAR(20) NOT NULL CHECK (content_type IN ('video', 'text', 'quiz')),
    content_url VARCHAR(500),
    content_text TEXT,
    order_index INTEGER NOT NULL,
    duration_minutes INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 学习记录表
CREATE TABLE learning_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    progress_percentage DECIMAL(5,2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'enrolled',
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    last_accessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### 社群相关表
```sql
-- 圈子表
CREATE TABLE circles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    domain VARCHAR(100) NOT NULL,
    created_by UUID NOT NULL REFERENCES user_identities(id),
    status VARCHAR(20) DEFAULT 'active',
    member_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 动态表
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    circle_id UUID NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    media_urls TEXT[],
    post_type VARCHAR(20) DEFAULT 'text' CHECK (post_type IN ('text', 'image', 'video', 'link')),
    status VARCHAR(20) DEFAULT 'active',
    like_count INTEGER DEFAULT 0,
    comment_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 评论表
CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    parent_id UUID REFERENCES comments(id) ON DELETE CASCADE,
    like_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### 预约和评价表
```sql
-- 预约表
CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    mentor_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    appointment_time TIMESTAMP NOT NULL,
    duration_minutes INTEGER NOT NULL,
    meeting_type VARCHAR(20) DEFAULT 'video' CHECK (meeting_type IN ('video', 'voice', 'text')),
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'confirmed', 'completed', 'cancelled')),
    price DECIMAL(10,2) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 评价表
CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reviewer_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    reviewed_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
    appointment_id UUID REFERENCES appointments(id) ON DELETE CASCADE,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    content TEXT,
    review_type VARCHAR(20) NOT NULL CHECK (review_type IN ('course', 'mentor', 'appointment')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 3.2 索引设计
```sql
-- 用户表索引
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_status ON users(status);

-- 身份表索引
CREATE INDEX idx_user_identities_user_id ON user_identities(user_id);
CREATE INDEX idx_user_identities_type_domain ON user_identities(identity_type, domain);
CREATE INDEX idx_user_identities_status ON user_identities(status);

-- 课程表索引
CREATE INDEX idx_courses_mentor_id ON courses(mentor_id);
CREATE INDEX idx_courses_status ON courses(status);
CREATE INDEX idx_courses_domain ON courses(domain);

-- 学习记录索引
CREATE INDEX idx_learning_records_user_course ON learning_records(user_id, course_id);
CREATE INDEX idx_learning_records_status ON learning_records(status);

-- 社群表索引
CREATE INDEX idx_posts_circle_id ON posts(circle_id);
CREATE INDEX idx_posts_user_id ON posts(user_id);
CREATE INDEX idx_posts_created_at ON posts(created_at DESC);

-- 预约表索引
CREATE INDEX idx_appointments_mentor_time ON appointments(mentor_id, appointment_time);
CREATE INDEX idx_appointments_student_time ON appointments(student_id, appointment_time);
CREATE INDEX idx_appointments_status ON appointments(status);
```

## 4. 核心模块设计

### 4.1 身份管理模块

#### 模块职责
- 用户身份创建和管理
- 身份切换逻辑处理
- 权限验证和访问控制
- 身份认证状态管理

#### 核心接口
```go
// IdentityService 身份管理服务接口
type IdentityService interface {
    CreateIdentity(ctx context.Context, req *CreateIdentityRequest) (*Identity, error)
    GetUserIdentities(ctx context.Context, userID string) ([]*Identity, error)
    SwitchIdentity(ctx context.Context, userID, identityID string) (*Identity, error)
    UpdateIdentityStatus(ctx context.Context, identityID, status string) error
    VerifyIdentity(ctx context.Context, identityID string, documents []string) error
}

// IdentityRepository 身份数据访问接口
type IdentityRepository interface {
    Create(ctx context.Context, identity *Identity) error
    GetByUserID(ctx context.Context, userID string) ([]*Identity, error)
    GetByID(ctx context.Context, id string) (*Identity, error)
    Update(ctx context.Context, identity *Identity) error
    Delete(ctx context.Context, id string) error
}
```

#### 数据模型
```go
// Identity 身份模型
type Identity struct {
    ID               string    `json:"id" gorm:"primaryKey;type:uuid"`
    UserID           string    `json:"user_id" gorm:"type:uuid;not null"`
    IdentityType     string    `json:"identity_type" gorm:"not null"` // master/apprentice
    Domain           string    `json:"domain" gorm:"not null"`
    Status           string    `json:"status" gorm:"default:'pending'"`
    VerificationStatus string  `json:"verification_status" gorm:"default:'unverified'"`
    CreatedAt        time.Time `json:"created_at"`
    UpdatedAt        time.Time `json:"updated_at"`
    
    // 关联关系
    User             *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Profile          *UserProfile `json:"profile,omitempty" gorm:"foreignKey:IdentityID"`
}
```

### 4.2 课程管理模块

#### 模块职责
- 课程创建和编辑
- 课程内容管理
- 学习进度跟踪
- 课程评价系统

#### 核心接口
```go
// CourseService 课程管理服务接口
type CourseService interface {
    CreateCourse(ctx context.Context, req *CreateCourseRequest) (*Course, error)
    UpdateCourse(ctx context.Context, courseID string, req *UpdateCourseRequest) (*Course, error)
    GetCourse(ctx context.Context, courseID string) (*CourseDetail, error)
    ListCourses(ctx context.Context, filter *CourseFilter) (*CourseList, error)
    EnrollCourse(ctx context.Context, userID, courseID string) error
    UpdateProgress(ctx context.Context, userID, courseID string, progress float64) error
    AddReview(ctx context.Context, req *AddReviewRequest) (*Review, error)
}

// CourseRepository 课程数据访问接口
type CourseRepository interface {
    Create(ctx context.Context, course *Course) error
    GetByID(ctx context.Context, id string) (*Course, error)
    Update(ctx context.Context, course *Course) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context, filter *CourseFilter) ([]*Course, error)
    Count(ctx context.Context, filter *CourseFilter) (int64, error)
}
```

#### 数据模型
```go
// Course 课程模型
type Course struct {
    ID            string    `json:"id" gorm:"primaryKey;type:uuid"`
    MentorID      string    `json:"mentor_id" gorm:"type:uuid;not null"`
    Title         string    `json:"title" gorm:"not null"`
    Description   string    `json:"description"`
    CoverImage    string    `json:"cover_image"`
    Price         float64   `json:"price" gorm:"not null"`
    DurationHours int       `json:"duration_hours" gorm:"not null"`
    Difficulty    string    `json:"difficulty" gorm:"check:difficulty IN ('beginner','intermediate','advanced')"`
    Status        string    `json:"status" gorm:"default:'draft'"`
    MaxStudents   int       `json:"max_students"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
    
    // 关联关系
    Mentor        *UserIdentity `json:"mentor,omitempty" gorm:"foreignKey:MentorID"`
    Contents      []CourseContent `json:"contents,omitempty" gorm:"foreignKey:CourseID"`
    Reviews       []Review `json:"reviews,omitempty" gorm:"foreignKey:CourseID"`
}
```

### 4.3 社群管理模块

#### 模块职责
- 动态发布和管理
- 评论和点赞功能
- 圈子管理
- 内容审核

#### 核心接口
```go
// CommunityService 社群管理服务接口
type CommunityService interface {
    CreatePost(ctx context.Context, req *CreatePostRequest) (*Post, error)
    GetPosts(ctx context.Context, circleID string, filter *PostFilter) (*PostList, error)
    AddComment(ctx context.Context, req *AddCommentRequest) (*Comment, error)
    LikePost(ctx context.Context, userID, postID string) error
    UnlikePost(ctx context.Context, userID, postID string) error
    CreateCircle(ctx context.Context, req *CreateCircleRequest) (*Circle, error)
    JoinCircle(ctx context.Context, userID, circleID string) error
    ModerateContent(ctx context.Context, contentID, action string) error
}

// CommunityRepository 社群数据访问接口
type CommunityRepository interface {
    CreatePost(ctx context.Context, post *Post) error
    GetPosts(ctx context.Context, filter *PostFilter) ([]*Post, error)
    GetPostByID(ctx context.Context, id string) (*Post, error)
    UpdatePost(ctx context.Context, post *Post) error
    DeletePost(ctx context.Context, id string) error
}
```

#### 数据模型
```go
// Post 动态模型
type Post struct {
    ID           string    `json:"id" gorm:"primaryKey;type:uuid"`
    UserID       string    `json:"user_id" gorm:"type:uuid;not null"`
    IdentityID   string    `json:"identity_id" gorm:"type:uuid;not null"`
    CircleID     string    `json:"circle_id" gorm:"type:uuid;not null"`
    Content      string    `json:"content" gorm:"not null"`
    MediaURLs    []string  `json:"media_urls" gorm:"type:text[]"`
    PostType     string    `json:"post_type" gorm:"default:'text'"`
    Status       string    `json:"status" gorm:"default:'active'"`
    LikeCount    int       `json:"like_count" gorm:"default:0"`
    CommentCount int       `json:"comment_count" gorm:"default:0"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
    
    // 关联关系
    User         *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Identity     *UserIdentity `json:"identity,omitempty" gorm:"foreignKey:IdentityID"`
    Circle       *Circle   `json:"circle,omitempty" gorm:"foreignKey:CircleID"`
    Comments     []Comment `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}
```

### 4.4 实时通信模块

#### 模块职责
- WebSocket连接管理
- 实时消息推送
- 在线状态管理
- 视频通话信令

#### 核心接口
```go
// WebSocketManager WebSocket管理接口
type WebSocketManager interface {
    Connect(ctx context.Context, userID string, conn *websocket.Conn) error
    Disconnect(userID string) error
    SendMessage(userID string, message *Message) error
    BroadcastToCircle(circleID string, message *Message) error
    GetOnlineUsers() []string
    UpdateUserStatus(userID, status string) error
}

// MessageService 消息服务接口
type MessageService interface {
    SendMessage(ctx context.Context, req *SendMessageRequest) error
    GetMessages(ctx context.Context, filter *MessageFilter) (*MessageList, error)
    MarkAsRead(ctx context.Context, userID, messageID string) error
    GetUnreadCount(ctx context.Context, userID string) (int, error)
}
```

#### 数据模型
```go
// Message 消息模型
type Message struct {
    ID        string    `json:"id" gorm:"primaryKey;type:uuid"`
    FromID    string    `json:"from_id" gorm:"type:uuid;not null"`
    ToID      string    `json:"to_id" gorm:"type:uuid"`
    CircleID  string    `json:"circle_id" gorm:"type:uuid"`
    Type      string    `json:"type" gorm:"not null"` // text, image, system
    Content   string    `json:"content" gorm:"not null"`
    IsRead    bool      `json:"is_read" gorm:"default:false"`
    CreatedAt time.Time `json:"created_at"`
    
    // 关联关系
    FromUser  *User     `json:"from_user,omitempty" gorm:"foreignKey:FromID"`
    ToUser    *User     `json:"to_user,omitempty" gorm:"foreignKey:ToID"`
}
```

## 5. 中间件设计

### 5.1 认证中间件
```go
// AuthMiddleware 认证中间件
type AuthMiddleware struct {
    jwtService JWTService
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
    return func(c *gin.Context) {
        // JWT Token验证
        // 用户身份验证
        // 权限检查
    }
}

func (m *AuthMiddleware) RequireIdentity() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 检查当前身份
        // 验证身份权限
    }
}
```

### 5.2 限流中间件
```go
// RateLimitMiddleware 限流中间件
type RateLimitMiddleware struct {
    redisClient *redis.Client
}

func (m *RateLimitMiddleware) RateLimit(limit int, window time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Redis计数器
        // 滑动窗口限流
        // 返回限流响应
    }
}
```

### 5.3 日志中间件
```go
// LogMiddleware 日志中间件
type LogMiddleware struct {
    logger *zap.Logger
}

func (m *LogMiddleware) LogRequest() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 请求开始时间
        // 请求信息记录
        // 响应时间统计
        // 错误日志记录
    }
}
```

## 6. 缓存策略设计

### 6.1 Redis缓存结构
```go
// 用户会话缓存
// Key: session:{user_id}
// Value: {token, identity_id, permissions, expires_at}

// 用户信息缓存
// Key: user:{user_id}
// Value: {user_data, expires_at}

// 课程信息缓存
// Key: course:{course_id}
// Value: {course_data, expires_at}

// 热门内容缓存
// Key: hot:posts:{circle_id}
// Value: {post_ids, expires_at}

// 在线用户缓存
// Key: online:users
// Value: Set of user_ids
```

### 6.2 缓存服务接口
```go
// CacheService 缓存服务接口
type CacheService interface {
    Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
    Get(ctx context.Context, key string, dest interface{}) error
    Delete(ctx context.Context, key string) error
    Exists(ctx context.Context, key string) (bool, error)
    Incr(ctx context.Context, key string) (int64, error)
    SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
}
```

## 7. 错误处理设计

### 7.1 错误类型定义
```go
// AppError 应用错误
type AppError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}

// 预定义错误码
const (
    ErrCodeSuccess        = 0
    ErrCodeInvalidInput   = 400
    ErrCodeUnauthorized   = 401
    ErrCodeForbidden      = 403
    ErrCodeNotFound       = 404
    ErrCodeConflict       = 409
    ErrCodeInternalError  = 500
    ErrCodeServiceUnavailable = 503
)
```

### 7.2 错误处理中间件
```go
// ErrorHandler 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        
        // 检查是否有错误
        if len(c.Errors) > 0 {
            err := c.Errors.Last()
            
            // 根据错误类型返回相应响应
            switch e := err.Err.(type) {
            case *AppError:
                c.JSON(e.Code, e)
            default:
                c.JSON(500, &AppError{
                    Code:    ErrCodeInternalError,
                    Message: "Internal server error",
                })
            }
        }
    }
}
```

## 8. 配置管理设计

### 8.1 配置结构
```go
// Config 应用配置
type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    JWT      JWTConfig      `mapstructure:"jwt"`
    Log      LogConfig      `mapstructure:"log"`
    Upload   UploadConfig   `mapstructure:"upload"`
}

type ServerConfig struct {
    Port         string `mapstructure:"port"`
    Mode         string `mapstructure:"mode"`
    ReadTimeout  int    `mapstructure:"read_timeout"`
    WriteTimeout int    `mapstructure:"write_timeout"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DBName   string `mapstructure:"dbname"`
    SSLMode  string `mapstructure:"sslmode"`
}
```

### 8.2 配置加载
```go
// LoadConfig 加载配置
func LoadConfig(configPath string) (*Config, error) {
    viper.SetConfigFile(configPath)
    viper.AutomaticEnv()
    
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }
    
    return &config, nil
}
```

## 9. 监控和日志设计

### 9.1 日志配置
```go
// Logger 日志配置
type Logger struct {
    *zap.Logger
}

func NewLogger(config LogConfig) (*Logger, error) {
    var cfg zap.Config
    
    if config.Environment == "production" {
        cfg = zap.NewProductionConfig()
    } else {
        cfg = zap.NewDevelopmentConfig()
    }
    
    logger, err := cfg.Build()
    if err != nil {
        return nil, err
    }
    
    return &Logger{logger}, nil
}
```

### 9.2 监控指标
```go
// Metrics 监控指标
type Metrics struct {
    RequestCounter   *prometheus.CounterVec
    RequestDuration  *prometheus.HistogramVec
    ActiveConnections prometheus.Gauge
    ErrorCounter     *prometheus.CounterVec
}

func NewMetrics() *Metrics {
    return &Metrics{
        RequestCounter: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Name: "http_requests_total",
                Help: "Total number of HTTP requests",
            },
            []string{"method", "endpoint", "status"},
        ),
        RequestDuration: prometheus.NewHistogramVec(
            prometheus.HistogramOpts{
                Name: "http_request_duration_seconds",
                Help: "HTTP request duration in seconds",
            },
            []string{"method", "endpoint"},
        ),
        ActiveConnections: prometheus.NewGauge(
            prometheus.GaugeOpts{
                Name: "websocket_active_connections",
                Help: "Number of active WebSocket connections",
            },
        ),
        ErrorCounter: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Name: "app_errors_total",
                Help: "Total number of application errors",
            },
            []string{"type", "service"},
        ),
    }
}
```

## 10. 部署和运维设计

### 10.1 Docker配置
```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080
CMD ["./main"]
```

### 10.2 健康检查
```go
// HealthCheck 健康检查
type HealthCheck struct {
    db    *gorm.DB
    redis *redis.Client
}

func (h *HealthCheck) Check() map[string]interface{} {
    status := map[string]interface{}{
        "status": "healthy",
        "timestamp": time.Now(),
        "services": map[string]interface{}{},
    }
    
    // 数据库健康检查
    if err := h.db.Raw("SELECT 1").Error; err != nil {
        status["status"] = "unhealthy"
        status["services"].(map[string]interface{})["database"] = "unhealthy"
    } else {
        status["services"].(map[string]interface{})["database"] = "healthy"
    }
    
    // Redis健康检查
    if err := h.redis.Ping(context.Background()).Err(); err != nil {
        status["status"] = "unhealthy"
        status["services"].(map[string]interface{})["redis"] = "unhealthy"
    } else {
        status["services"].(map[string]interface{})["redis"] = "healthy"
    }
    
    return status
}
```

## 11. 安全设计

### 11.1 数据加密
```go
// EncryptionService 加密服务
type EncryptionService interface {
    Encrypt(data []byte) ([]byte, error)
    Decrypt(data []byte) ([]byte, error)
    HashPassword(password string) (string, error)
    ComparePassword(hashedPassword, password string) error
}
```

### 11.2 输入验证
```go
// Validator 输入验证
type Validator struct {
    validate *validator.Validate
}

func (v *Validator) ValidateStruct(s interface{}) error {
    return v.validate.Struct(s)
}

func (v *Validator) ValidateEmail(email string) error {
    return v.validate.Var(email, "required,email")
}
```

## 12. 总结

Master Guide后端采用Go + Gin + PostgreSQL + Redis技术栈，具备以下优势：

### 技术优势
- **高性能**：Go的并发模型和Gin的轻量级设计
- **数据一致性**：PostgreSQL的ACID事务支持
- **缓存效率**：Redis的高性能缓存和会话管理
- **实时通信**：WebSocket支持实时消息推送
- **安全性**：JWT认证、密码加密、限流保护

### 架构特点
- **模块化设计**：清晰的模块划分和职责分离
- **数据驱动**：基于数据模型的设计
- **缓存策略**：多级缓存提升性能
- **监控运维**：完整的监控和日志体系
- **安全防护**：多层次安全保护机制

该后端架构设计确保了系统的稳定性、可扩展性和可维护性，为Master Guide平台提供了坚实的后端基础。 