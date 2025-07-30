# Master Guide 数据库设计文档

## 1. 数据库概述

### 1.1 技术选型
- **主数据库**: PostgreSQL 15.0+
- **缓存数据库**: Redis 7.0+
- **对象存储**: 阿里云 OSS
- **搜索引擎**: Elasticsearch 8.0+

### 1.2 设计原则
- **数据一致性**: 使用ACID事务保证数据一致性
- **性能优化**: 合理的索引设计和查询优化
- **扩展性**: 支持水平扩展和分库分表
- **安全性**: 敏感数据加密存储
- **可维护性**: 清晰的表结构和命名规范

## 2. 数据库架构

### 2.1 整体架构
```
┌─────────────────────────────────────────────────────────────┐
│                    应用层 (Application)                      │
├─────────────────────────────────────────────────────────────┤
│  Go + Gin + GORM + Repository Pattern                      │
└─────────────────────────────────────────────────────────────┘
                                │
                                │ SQL/ORM
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                  数据库层 (Database Layer)                   │
├─────────────────────────────────────────────────────────────┤
│  PostgreSQL (主数据库) + Redis (缓存) + OSS (文件存储)        │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 数据库连接池配置
```go
// 连接池配置
MaxOpenConns: 100
MaxIdleConns: 10
ConnMaxLifetime: 1 * time.Hour
ConnMaxIdleTime: 15 * time.Minute
```

## 3. 核心数据表设计

### 3.1 用户相关表

#### 3.1.1 users (用户基础表)
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'banned')),
    email_verified BOOLEAN DEFAULT FALSE,
    phone_verified BOOLEAN DEFAULT FALSE,
    last_login_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_users_created_at ON users(created_at);
```

#### 3.1.2 user_identities (用户身份表)
```sql
CREATE TABLE user_identities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_type VARCHAR(20) NOT NULL CHECK (identity_type IN ('master', 'apprentice')),
    domain VARCHAR(100) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'active', 'suspended', 'rejected')),
    verification_status VARCHAR(20) DEFAULT 'unverified' CHECK (verification_status IN ('unverified', 'pending', 'verified', 'rejected')),
    verification_documents TEXT[], -- 认证文档URL数组
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, domain, identity_type)
);

-- 索引
CREATE INDEX idx_user_identities_user_id ON user_identities(user_id);
CREATE INDEX idx_user_identities_type_domain ON user_identities(identity_type, domain);
CREATE INDEX idx_user_identities_status ON user_identities(status);
CREATE INDEX idx_user_identities_verification_status ON user_identities(verification_status);
```

#### 3.1.3 user_profiles (用户档案表)
```sql
CREATE TABLE user_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    avatar VARCHAR(500),
    bio TEXT,
    skills TEXT[], -- 技能标签数组
    experience_years INTEGER DEFAULT 0,
    hourly_rate DECIMAL(10,2),
    location VARCHAR(200),
    website VARCHAR(500),
    social_links JSONB, -- 社交媒体链接
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(identity_id)
);

-- 索引
CREATE INDEX idx_user_profiles_user_id ON user_profiles(user_id);
CREATE INDEX idx_user_profiles_identity_id ON user_profiles(identity_id);
CREATE INDEX idx_user_profiles_name ON user_profiles(name);
CREATE INDEX idx_user_profiles_skills ON user_profiles USING GIN(skills);
```

### 3.2 课程相关表

#### 3.2.1 courses (课程表)
```sql
CREATE TABLE courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mentor_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    cover_image VARCHAR(500),
    price DECIMAL(10,2) NOT NULL,
    duration_hours INTEGER NOT NULL,
    difficulty VARCHAR(20) CHECK (difficulty IN ('beginner', 'intermediate', 'advanced')),
    status VARCHAR(20) DEFAULT 'draft' CHECK (status IN ('draft', 'published', 'archived')),
    max_students INTEGER,
    current_students INTEGER DEFAULT 0,
    rating DECIMAL(3,2) DEFAULT 0,
    review_count INTEGER DEFAULT 0,
    tags TEXT[], -- 课程标签
    prerequisites TEXT, -- 前置要求
    learning_objectives TEXT[], -- 学习目标
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_courses_mentor_id ON courses(mentor_id);
CREATE INDEX idx_courses_status ON courses(status);
CREATE INDEX idx_courses_difficulty ON courses(difficulty);
CREATE INDEX idx_courses_price ON courses(price);
CREATE INDEX idx_courses_rating ON courses(rating);
CREATE INDEX idx_courses_created_at ON courses(created_at);
CREATE INDEX idx_courses_tags ON courses USING GIN(tags);
```

#### 3.2.2 course_contents (课程内容表)
```sql
CREATE TABLE course_contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    content_type VARCHAR(20) NOT NULL CHECK (content_type IN ('video', 'text', 'quiz', 'assignment')),
    content_url VARCHAR(500),
    content_text TEXT,
    order_index INTEGER NOT NULL,
    duration_minutes INTEGER,
    is_free BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_course_contents_course_id ON course_contents(course_id);
CREATE INDEX idx_course_contents_order_index ON course_contents(course_id, order_index);
```

#### 3.2.3 learning_records (学习记录表)
```sql
CREATE TABLE learning_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    progress_percentage DECIMAL(5,2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'enrolled' CHECK (status IN ('enrolled', 'learning', 'completed', 'dropped')),
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    last_accessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_study_time INTEGER DEFAULT 0, -- 总学习时间（分钟）
    UNIQUE(user_id, course_id)
);

-- 索引
CREATE INDEX idx_learning_records_user_id ON learning_records(user_id);
CREATE INDEX idx_learning_records_course_id ON learning_records(course_id);
CREATE INDEX idx_learning_records_status ON learning_records(status);
CREATE INDEX idx_learning_records_enrolled_at ON learning_records(enrolled_at);
```

#### 3.2.4 content_progress (内容学习进度表)
```sql
CREATE TABLE content_progress (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content_id UUID NOT NULL REFERENCES course_contents(id) ON DELETE CASCADE,
    progress_percentage DECIMAL(5,2) DEFAULT 0,
    is_completed BOOLEAN DEFAULT FALSE,
    completed_at TIMESTAMP,
    last_accessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    study_time INTEGER DEFAULT 0, -- 学习时间（分钟）
    UNIQUE(user_id, content_id)
);

-- 索引
CREATE INDEX idx_content_progress_user_id ON content_progress(user_id);
CREATE INDEX idx_content_progress_content_id ON content_progress(content_id);
CREATE INDEX idx_content_progress_is_completed ON content_progress(is_completed);
```

### 3.3 社群相关表

#### 3.3.1 circles (圈子表)
```sql
CREATE TABLE circles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    domain VARCHAR(100) NOT NULL,
    created_by UUID NOT NULL REFERENCES user_identities(id),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'archived')),
    member_count INTEGER DEFAULT 0,
    post_count INTEGER DEFAULT 0,
    avatar VARCHAR(500),
    rules TEXT, -- 圈子规则
    tags TEXT[], -- 圈子标签
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_circles_domain ON circles(domain);
CREATE INDEX idx_circles_status ON circles(status);
CREATE INDEX idx_circles_created_by ON circles(created_by);
CREATE INDEX idx_circles_member_count ON circles(member_count);
CREATE INDEX idx_circles_tags ON circles USING GIN(tags);
```

#### 3.3.2 circle_members (圈子成员表)
```sql
CREATE TABLE circle_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    circle_id UUID NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    role VARCHAR(20) DEFAULT 'member' CHECK (role IN ('member', 'moderator', 'admin')),
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    UNIQUE(circle_id, user_id)
);

-- 索引
CREATE INDEX idx_circle_members_circle_id ON circle_members(circle_id);
CREATE INDEX idx_circle_members_user_id ON circle_members(user_id);
CREATE INDEX idx_circle_members_role ON circle_members(role);
```

#### 3.3.3 posts (动态表)
```sql
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    circle_id UUID NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    media_urls TEXT[], -- 媒体文件URL数组
    post_type VARCHAR(20) DEFAULT 'text' CHECK (post_type IN ('text', 'image', 'video', 'link', 'poll')),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'hidden', 'deleted')),
    like_count INTEGER DEFAULT 0,
    comment_count INTEGER DEFAULT 0,
    share_count INTEGER DEFAULT 0,
    view_count INTEGER DEFAULT 0,
    is_top BOOLEAN DEFAULT FALSE, -- 是否置顶
    is_essence BOOLEAN DEFAULT FALSE, -- 是否精华
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_posts_circle_id ON posts(circle_id);
CREATE INDEX idx_posts_user_id ON posts(user_id);
CREATE INDEX idx_posts_identity_id ON posts(identity_id);
CREATE INDEX idx_posts_status ON posts(status);
CREATE INDEX idx_posts_created_at ON posts(created_at DESC);
CREATE INDEX idx_posts_like_count ON posts(like_count DESC);
CREATE INDEX idx_posts_post_type ON posts(post_type);
```

#### 3.3.4 comments (评论表)
```sql
CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    parent_id UUID REFERENCES comments(id) ON DELETE CASCADE, -- 父评论ID，支持回复
    reply_to_user_id UUID REFERENCES users(id), -- 回复的用户ID
    like_count INTEGER DEFAULT 0,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'hidden', 'deleted')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_comments_post_id ON comments(post_id);
CREATE INDEX idx_comments_user_id ON comments(user_id);
CREATE INDEX idx_comments_parent_id ON comments(parent_id);
CREATE INDEX idx_comments_created_at ON comments(created_at);
```

#### 3.3.5 post_likes (动态点赞表)
```sql
CREATE TABLE post_likes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(post_id, user_id)
);

-- 索引
CREATE INDEX idx_post_likes_post_id ON post_likes(post_id);
CREATE INDEX idx_post_likes_user_id ON post_likes(user_id);
```

#### 3.3.6 comment_likes (评论点赞表)
```sql
CREATE TABLE comment_likes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    comment_id UUID NOT NULL REFERENCES comments(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(comment_id, user_id)
);

-- 索引
CREATE INDEX idx_comment_likes_comment_id ON comment_likes(comment_id);
CREATE INDEX idx_comment_likes_user_id ON comment_likes(user_id);
```

### 3.4 预约和评价表

#### 3.4.1 appointments (预约表)
```sql
CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    mentor_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    appointment_time TIMESTAMP NOT NULL,
    duration_minutes INTEGER NOT NULL,
    meeting_type VARCHAR(20) DEFAULT 'video' CHECK (meeting_type IN ('video', 'voice', 'text')),
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'confirmed', 'completed', 'cancelled', 'no_show')),
    price DECIMAL(10,2) NOT NULL,
    notes TEXT,
    meeting_url VARCHAR(500), -- 会议链接
    meeting_id VARCHAR(100), -- 第三方会议ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_appointments_student_id ON appointments(student_id);
CREATE INDEX idx_appointments_mentor_id ON appointments(mentor_id);
CREATE INDEX idx_appointments_appointment_time ON appointments(appointment_time);
CREATE INDEX idx_appointments_status ON appointments(status);
CREATE INDEX idx_appointments_meeting_type ON appointments(meeting_type);
```

#### 3.4.2 reviews (评价表)
```sql
CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reviewer_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    reviewed_id UUID NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
    appointment_id UUID REFERENCES appointments(id) ON DELETE CASCADE,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    content TEXT,
    review_type VARCHAR(20) NOT NULL CHECK (review_type IN ('course', 'mentor', 'appointment')),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'hidden', 'deleted')),
    like_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_reviews_reviewer_id ON reviews(reviewer_id);
CREATE INDEX idx_reviews_reviewed_id ON reviews(reviewed_id);
CREATE INDEX idx_reviews_course_id ON reviews(course_id);
CREATE INDEX idx_reviews_appointment_id ON reviews(appointment_id);
CREATE INDEX idx_reviews_rating ON reviews(rating);
CREATE INDEX idx_reviews_review_type ON reviews(review_type);
CREATE INDEX idx_reviews_created_at ON reviews(created_at);
```

### 3.5 消息和通知表

#### 3.5.1 messages (消息表)
```sql
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    to_user_id UUID REFERENCES users(id), -- 私聊消息
    circle_id UUID REFERENCES circles(id), -- 群聊消息
    content TEXT NOT NULL,
    message_type VARCHAR(20) DEFAULT 'text' CHECK (message_type IN ('text', 'image', 'file', 'system')),
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_messages_from_user_id ON messages(from_user_id);
CREATE INDEX idx_messages_to_user_id ON messages(to_user_id);
CREATE INDEX idx_messages_circle_id ON messages(circle_id);
CREATE INDEX idx_messages_is_read ON messages(is_read);
CREATE INDEX idx_messages_created_at ON messages(created_at);
```

#### 3.5.2 notifications (通知表)
```sql
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    content TEXT,
    notification_type VARCHAR(50) NOT NULL,
    related_id UUID, -- 相关对象ID
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_notifications_user_id ON notifications(user_id);
CREATE INDEX idx_notifications_is_read ON notifications(is_read);
CREATE INDEX idx_notifications_created_at ON notifications(created_at);
```

### 3.6 系统配置表

#### 3.6.1 domains (专业领域表)
```sql
CREATE TABLE domains (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    icon VARCHAR(500),
    sort_order INTEGER DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_domains_code ON domains(code);
CREATE INDEX idx_domains_is_active ON domains(is_active);
CREATE INDEX idx_domains_sort_order ON domains(sort_order);
```

#### 3.6.2 system_configs (系统配置表)
```sql
CREATE TABLE system_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    config_key VARCHAR(100) NOT NULL UNIQUE,
    config_value TEXT,
    config_type VARCHAR(20) DEFAULT 'string' CHECK (config_type IN ('string', 'number', 'boolean', 'json')),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_system_configs_key ON system_configs(config_key);
```

## 4. 数据库视图

### 4.1 大师统计视图
```sql
CREATE VIEW mentor_stats AS
SELECT 
    ui.id as identity_id,
    ui.user_id,
    up.name,
    up.avatar,
    ui.domain,
    up.hourly_rate,
    COUNT(DISTINCT c.id) as course_count,
    COUNT(DISTINCT lr.user_id) as student_count,
    AVG(r.rating) as average_rating,
    COUNT(r.id) as review_count
FROM user_identities ui
LEFT JOIN user_profiles up ON ui.id = up.identity_id
LEFT JOIN courses c ON ui.id = c.mentor_id
LEFT JOIN learning_records lr ON c.id = lr.course_id
LEFT JOIN reviews r ON ui.id = r.reviewed_id AND r.review_type = 'mentor'
WHERE ui.identity_type = 'master' AND ui.status = 'active'
GROUP BY ui.id, ui.user_id, up.name, up.avatar, ui.domain, up.hourly_rate;
```

### 4.2 课程统计视图
```sql
CREATE VIEW course_stats AS
SELECT 
    c.id,
    c.title,
    c.mentor_id,
    up.name as mentor_name,
    c.price,
    c.duration_hours,
    c.difficulty,
    COUNT(lr.user_id) as enrolled_students,
    AVG(lr.progress_percentage) as avg_progress,
    AVG(r.rating) as average_rating,
    COUNT(r.id) as review_count
FROM courses c
LEFT JOIN user_profiles up ON c.mentor_id = up.identity_id
LEFT JOIN learning_records lr ON c.id = lr.course_id
LEFT JOIN reviews r ON c.id = r.course_id AND r.review_type = 'course'
WHERE c.status = 'published'
GROUP BY c.id, c.title, c.mentor_id, up.name, c.price, c.duration_hours, c.difficulty;
```

## 5. 存储过程

### 5.1 更新课程统计
```sql
CREATE OR REPLACE FUNCTION update_course_stats()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' OR TG_OP = 'DELETE' THEN
        -- 更新课程学生数量
        UPDATE courses 
        SET current_students = (
            SELECT COUNT(*) 
            FROM learning_records 
            WHERE course_id = COALESCE(NEW.course_id, OLD.course_id)
        )
        WHERE id = COALESCE(NEW.course_id, OLD.course_id);
        
        -- 更新课程评分
        UPDATE courses 
        SET rating = (
            SELECT AVG(rating) 
            FROM reviews 
            WHERE course_id = COALESCE(NEW.course_id, OLD.course_id) 
            AND review_type = 'course'
        ),
        review_count = (
            SELECT COUNT(*) 
            FROM reviews 
            WHERE course_id = COALESCE(NEW.course_id, OLD.course_id) 
            AND review_type = 'course'
        )
        WHERE id = COALESCE(NEW.course_id, OLD.course_id);
    END IF;
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;
```

### 5.2 更新动态统计
```sql
CREATE OR REPRACE FUNCTION update_post_stats()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' OR TG_OP = 'DELETE' THEN
        -- 更新动态点赞数
        UPDATE posts 
        SET like_count = (
            SELECT COUNT(*) 
            FROM post_likes 
            WHERE post_id = COALESCE(NEW.post_id, OLD.post_id)
        )
        WHERE id = COALESCE(NEW.post_id, OLD.post_id);
    END IF;
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;
```

## 6. 触发器

### 6.1 学习记录触发器
```sql
CREATE TRIGGER trigger_update_course_stats
AFTER INSERT OR UPDATE OR DELETE ON learning_records
FOR EACH ROW EXECUTE FUNCTION update_course_stats();
```

### 6.2 评价触发器
```sql
CREATE TRIGGER trigger_update_course_stats_from_reviews
AFTER INSERT OR UPDATE OR DELETE ON reviews
FOR EACH ROW EXECUTE FUNCTION update_course_stats();
```

### 6.3 点赞触发器
```sql
CREATE TRIGGER trigger_update_post_stats
AFTER INSERT OR UPDATE OR DELETE ON post_likes
FOR EACH ROW EXECUTE FUNCTION update_post_stats();
```

## 7. 数据备份和恢复

### 7.1 备份策略
```bash
#!/bin/bash
# 数据库备份脚本

# 设置变量
DB_NAME="master_guide"
BACKUP_DIR="/backup/postgresql"
DATE=$(date +%Y%m%d_%H%M%S)

# 创建备份目录
mkdir -p $BACKUP_DIR

# 执行备份
pg_dump -h localhost -U postgres -d $DB_NAME -F c -f $BACKUP_DIR/${DB_NAME}_${DATE}.dump

# 删除7天前的备份
find $BACKUP_DIR -name "*.dump" -mtime +7 -delete
```

### 7.2 恢复脚本
```bash
#!/bin/bash
# 数据库恢复脚本

# 设置变量
DB_NAME="master_guide"
BACKUP_FILE="$1"

# 检查备份文件是否存在
if [ ! -f "$BACKUP_FILE" ]; then
    echo "备份文件不存在: $BACKUP_FILE"
    exit 1
fi

# 执行恢复
pg_restore -h localhost -U postgres -d $DB_NAME -c $BACKUP_FILE
```

## 8. 性能优化

### 8.1 查询优化
- 使用适当的索引
- 避免全表扫描
- 使用EXPLAIN分析查询计划
- 定期更新统计信息

### 8.2 连接池配置
```go
// 数据库连接池配置
db.DB().SetMaxOpenConns(100)
db.DB().SetMaxIdleConns(10)
db.DB().SetConnMaxLifetime(time.Hour)
db.DB().SetConnMaxIdleTime(15 * time.Minute)
```

### 8.3 缓存策略
```go
// Redis缓存键设计
const (
    UserCacheKey     = "user:%s"
    CourseCacheKey   = "course:%s"
    MentorCacheKey   = "mentor:%s"
    CircleCacheKey   = "circle:%s"
    PostCacheKey     = "post:%s"
)
```

## 9. 数据安全

### 9.1 数据加密
```sql
-- 敏感数据加密
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- 密码哈希存储
UPDATE users SET password_hash = crypt(password, gen_salt('bf'));
```

### 9.2 访问控制
```sql
-- 创建只读用户
CREATE USER readonly WITH PASSWORD 'password';
GRANT CONNECT ON DATABASE master_guide TO readonly;
GRANT USAGE ON SCHEMA public TO readonly;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly;
```

### 9.3 审计日志
```sql
-- 创建审计表
CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    table_name VARCHAR(100) NOT NULL,
    operation VARCHAR(20) NOT NULL,
    old_data JSONB,
    new_data JSONB,
    user_id UUID,
    ip_address INET,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 10. 监控和维护

### 10.1 性能监控
```sql
-- 慢查询监控
SELECT 
    query,
    calls,
    total_time,
    mean_time,
    rows
FROM pg_stat_statements 
ORDER BY mean_time DESC 
LIMIT 10;
```

### 10.2 表大小监控
```sql
-- 表大小统计
SELECT 
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) as size
FROM pg_tables 
WHERE schemaname = 'public'
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;
```

### 10.3 索引使用情况
```sql
-- 索引使用统计
SELECT 
    schemaname,
    tablename,
    indexname,
    idx_scan,
    idx_tup_read,
    idx_tup_fetch
FROM pg_stat_user_indexes 
ORDER BY idx_scan DESC;
```

---

**文档版本**: v1.0  
**最后更新**: 2024-12-01  
**维护者**: Master Guide 技术团队 