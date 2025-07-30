-- Master Guide 数据库初始化脚本
-- 版本: v1.0
-- 创建时间: 2024-12-01
-- 描述: Master Guide平台数据库初始化

-- 创建数据库
CREATE DATABASE master_guide
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- 连接到数据库
\c master_guide;

-- 启用必要的扩展
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_stat_statements";

-- 创建全局ID生成函数
CREATE OR REPLACE FUNCTION generate_table_id(IN prefix VARCHAR(32) DEFAULT 'DEFAULTID_', IN seq_name VARCHAR(50) DEFAULT 'default_id_num_seq') 
RETURNS VARCHAR(32) 
LANGUAGE 'plpgsql' 
VOLATILE AS $BODY$ 
DECLARE 
    next_val BIGINT; 
BEGIN 
    next_val := nextval(seq_name); 
    RETURN prefix || LPAD(next_val::TEXT, 11, '0'); 
END; 
$BODY$;

ALTER FUNCTION generate_table_id(VARCHAR(32), VARCHAR(50)) OWNER TO postgres;
COMMENT ON FUNCTION generate_table_id(VARCHAR(32), VARCHAR(50)) IS '获取ID的全局方法';

-- 创建各表的ID序列
CREATE SEQUENCE IF NOT EXISTS user_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS user_identity_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS user_profile_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS domain_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS course_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS course_content_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS learning_record_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS content_progress_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS circle_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS circle_member_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS post_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS comment_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS post_like_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS comment_like_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS appointment_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS review_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS message_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS notification_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS system_config_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS audit_log_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;

-- 设置序列所有者
ALTER SEQUENCE user_id_num_seq OWNER TO postgres;
ALTER SEQUENCE user_identity_id_num_seq OWNER TO postgres;
ALTER SEQUENCE user_profile_id_num_seq OWNER TO postgres;
ALTER SEQUENCE domain_id_num_seq OWNER TO postgres;
ALTER SEQUENCE course_id_num_seq OWNER TO postgres;
ALTER SEQUENCE course_content_id_num_seq OWNER TO postgres;
ALTER SEQUENCE learning_record_id_num_seq OWNER TO postgres;
ALTER SEQUENCE content_progress_id_num_seq OWNER TO postgres;
ALTER SEQUENCE circle_id_num_seq OWNER TO postgres;
ALTER SEQUENCE circle_member_id_num_seq OWNER TO postgres;
ALTER SEQUENCE post_id_num_seq OWNER TO postgres;
ALTER SEQUENCE comment_id_num_seq OWNER TO postgres;
ALTER SEQUENCE post_like_id_num_seq OWNER TO postgres;
ALTER SEQUENCE comment_like_id_num_seq OWNER TO postgres;
ALTER SEQUENCE appointment_id_num_seq OWNER TO postgres;
ALTER SEQUENCE review_id_num_seq OWNER TO postgres;
ALTER SEQUENCE message_id_num_seq OWNER TO postgres;
ALTER SEQUENCE notification_id_num_seq OWNER TO postgres;
ALTER SEQUENCE system_config_id_num_seq OWNER TO postgres;
ALTER SEQUENCE audit_log_id_num_seq OWNER TO postgres;

-- 创建用户基础表
CREATE TABLE users (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('USER_', 'user_id_num_seq'),
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

-- 创建用户身份表
CREATE TABLE user_identities (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('IDENTITY_', 'user_identity_id_num_seq'),
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_type VARCHAR(20) NOT NULL CHECK (identity_type IN ('master', 'apprentice')),
    domain VARCHAR(100) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'active', 'suspended', 'rejected')),
    verification_status VARCHAR(20) DEFAULT 'unverified' CHECK (verification_status IN ('unverified', 'pending', 'verified', 'rejected')),
    verification_documents TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, domain, identity_type)
);

-- 创建用户档案表
CREATE TABLE user_profiles (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('PROFILE_', 'user_profile_id_num_seq'),
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id VARCHAR(32) NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    avatar VARCHAR(500),
    bio TEXT,
    skills TEXT[],
    experience_years INTEGER DEFAULT 0,
    hourly_rate DECIMAL(10,2),
    location VARCHAR(200),
    website VARCHAR(500),
    social_links JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(identity_id)
);

-- 创建专业领域表
CREATE TABLE domains (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('DOMAIN_', 'domain_id_num_seq'),
    name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    icon VARCHAR(500),
    sort_order INTEGER DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建课程表
CREATE TABLE courses (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('COURSE_', 'course_id_num_seq'),
    mentor_id VARCHAR(32) NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
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
    tags TEXT[],
    prerequisites TEXT,
    learning_objectives TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建课程内容表
CREATE TABLE course_contents (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('CONTENT_', 'course_content_id_num_seq'),
    course_id VARCHAR(32) NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
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

-- 创建学习记录表
CREATE TABLE learning_records (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('LEARNING_', 'learning_record_id_num_seq'),
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id VARCHAR(32) NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    progress_percentage DECIMAL(5,2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'enrolled' CHECK (status IN ('enrolled', 'learning', 'completed', 'dropped')),
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    last_accessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_study_time INTEGER DEFAULT 0,
    UNIQUE(user_id, course_id)
);

-- 创建内容学习进度表
CREATE TABLE content_progress (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('PROGRESS_', 'content_progress_id_num_seq'),
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content_id VARCHAR(32) NOT NULL REFERENCES course_contents(id) ON DELETE CASCADE,
    progress_percentage DECIMAL(5,2) DEFAULT 0,
    is_completed BOOLEAN DEFAULT FALSE,
    completed_at TIMESTAMP,
    last_accessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    study_time INTEGER DEFAULT 0,
    UNIQUE(user_id, content_id)
);

-- 创建圈子表
CREATE TABLE circles (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('CIRCLE_', 'circle_id_num_seq'),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    domain VARCHAR(100) NOT NULL,
    created_by VARCHAR(32) NOT NULL REFERENCES user_identities(id),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'archived')),
    member_count INTEGER DEFAULT 0,
    post_count INTEGER DEFAULT 0,
    avatar VARCHAR(500),
    rules TEXT,
    tags TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建圈子成员表
CREATE TABLE circle_members (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('MEMBER_', 'circle_member_id_num_seq'),
    circle_id VARCHAR(32) NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id VARCHAR(32) NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    role VARCHAR(20) DEFAULT 'member' CHECK (role IN ('member', 'moderator', 'admin')),
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    UNIQUE(circle_id, user_id)
);

-- 创建动态表
CREATE TABLE posts (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('POST_', 'post_id_num_seq'),
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id VARCHAR(32) NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    circle_id VARCHAR(32) NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    media_urls TEXT[],
    post_type VARCHAR(20) DEFAULT 'text' CHECK (post_type IN ('text', 'image', 'video', 'link', 'poll')),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'hidden', 'deleted')),
    like_count INTEGER DEFAULT 0,
    comment_count INTEGER DEFAULT 0,
    share_count INTEGER DEFAULT 0,
    view_count INTEGER DEFAULT 0,
    is_top BOOLEAN DEFAULT FALSE,
    is_essence BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建评论表
CREATE TABLE comments (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('COMMENT_', 'comment_id_num_seq'),
    post_id VARCHAR(32) NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    identity_id VARCHAR(32) NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    parent_id VARCHAR(32) REFERENCES comments(id) ON DELETE CASCADE,
    reply_to_user_id VARCHAR(32) REFERENCES users(id),
    like_count INTEGER DEFAULT 0,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'hidden', 'deleted')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建动态点赞表
CREATE TABLE post_likes (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('POSTLIKE_', 'post_like_id_num_seq'),
    post_id VARCHAR(32) NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(post_id, user_id)
);

-- 创建评论点赞表
CREATE TABLE comment_likes (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('COMMENTLIKE_', 'comment_like_id_num_seq'),
    comment_id VARCHAR(32) NOT NULL REFERENCES comments(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(comment_id, user_id)
);

-- 创建预约表
CREATE TABLE appointments (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('APPOINTMENT_', 'appointment_id_num_seq'),
    student_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    mentor_id VARCHAR(32) NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    appointment_time TIMESTAMP NOT NULL,
    duration_minutes INTEGER NOT NULL,
    meeting_type VARCHAR(20) DEFAULT 'video' CHECK (meeting_type IN ('video', 'voice', 'text')),
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'confirmed', 'completed', 'cancelled', 'no_show')),
    price DECIMAL(10,2) NOT NULL,
    notes TEXT,
    meeting_url VARCHAR(500),
    meeting_id VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建评价表
CREATE TABLE reviews (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('REVIEW_', 'review_id_num_seq'),
    reviewer_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    reviewed_id VARCHAR(32) NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    course_id VARCHAR(32) REFERENCES courses(id) ON DELETE CASCADE,
    appointment_id VARCHAR(32) REFERENCES appointments(id) ON DELETE CASCADE,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    content TEXT,
    review_type VARCHAR(20) NOT NULL CHECK (review_type IN ('course', 'mentor', 'appointment')),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'hidden', 'deleted')),
    like_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建消息表
CREATE TABLE messages (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('MESSAGE_', 'message_id_num_seq'),
    from_user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    to_user_id VARCHAR(32) REFERENCES users(id),
    circle_id VARCHAR(32) REFERENCES circles(id),
    content TEXT NOT NULL,
    message_type VARCHAR(20) DEFAULT 'text' CHECK (message_type IN ('text', 'image', 'file', 'system')),
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建通知表
CREATE TABLE notifications (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('NOTIFICATION_', 'notification_id_num_seq'),
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    content TEXT,
    notification_type VARCHAR(50) NOT NULL,
    related_id VARCHAR(32),
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建系统配置表
CREATE TABLE system_configs (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('CONFIG_', 'system_config_id_num_seq'),
    config_key VARCHAR(100) NOT NULL UNIQUE,
    config_value TEXT,
    config_type VARCHAR(20) DEFAULT 'string' CHECK (config_type IN ('string', 'number', 'boolean', 'json')),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建审计日志表
CREATE TABLE audit_logs (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('AUDIT_', 'audit_log_id_num_seq'),
    table_name VARCHAR(100) NOT NULL,
    operation VARCHAR(20) NOT NULL,
    old_data JSONB,
    new_data JSONB,
    user_id VARCHAR(32),
    ip_address INET,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建索引
-- 用户表索引
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_users_created_at ON users(created_at);

-- 用户身份表索引
CREATE INDEX idx_user_identities_user_id ON user_identities(user_id);
CREATE INDEX idx_user_identities_type_domain ON user_identities(identity_type, domain);
CREATE INDEX idx_user_identities_status ON user_identities(status);
CREATE INDEX idx_user_identities_verification_status ON user_identities(verification_status);

-- 用户档案表索引
CREATE INDEX idx_user_profiles_user_id ON user_profiles(user_id);
CREATE INDEX idx_user_profiles_identity_id ON user_profiles(identity_id);
CREATE INDEX idx_user_profiles_name ON user_profiles(name);
CREATE INDEX idx_user_profiles_skills ON user_profiles USING GIN(skills);

-- 专业领域表索引
CREATE INDEX idx_domains_code ON domains(code);
CREATE INDEX idx_domains_is_active ON domains(is_active);
CREATE INDEX idx_domains_sort_order ON domains(sort_order);

-- 课程表索引
CREATE INDEX idx_courses_mentor_id ON courses(mentor_id);
CREATE INDEX idx_courses_status ON courses(status);
CREATE INDEX idx_courses_difficulty ON courses(difficulty);
CREATE INDEX idx_courses_price ON courses(price);
CREATE INDEX idx_courses_rating ON courses(rating);
CREATE INDEX idx_courses_created_at ON courses(created_at);
CREATE INDEX idx_courses_tags ON courses USING GIN(tags);

-- 课程内容表索引
CREATE INDEX idx_course_contents_course_id ON course_contents(course_id);
CREATE INDEX idx_course_contents_order_index ON course_contents(course_id, order_index);

-- 学习记录表索引
CREATE INDEX idx_learning_records_user_id ON learning_records(user_id);
CREATE INDEX idx_learning_records_course_id ON learning_records(course_id);
CREATE INDEX idx_learning_records_status ON learning_records(status);
CREATE INDEX idx_learning_records_enrolled_at ON learning_records(enrolled_at);

-- 内容学习进度表索引
CREATE INDEX idx_content_progress_user_id ON content_progress(user_id);
CREATE INDEX idx_content_progress_content_id ON content_progress(content_id);
CREATE INDEX idx_content_progress_is_completed ON content_progress(is_completed);

-- 圈子表索引
CREATE INDEX idx_circles_domain ON circles(domain);
CREATE INDEX idx_circles_status ON circles(status);
CREATE INDEX idx_circles_created_by ON circles(created_by);
CREATE INDEX idx_circles_member_count ON circles(member_count);
CREATE INDEX idx_circles_tags ON circles USING GIN(tags);

-- 圈子成员表索引
CREATE INDEX idx_circle_members_circle_id ON circle_members(circle_id);
CREATE INDEX idx_circle_members_user_id ON circle_members(user_id);
CREATE INDEX idx_circle_members_role ON circle_members(role);

-- 动态表索引
CREATE INDEX idx_posts_circle_id ON posts(circle_id);
CREATE INDEX idx_posts_user_id ON posts(user_id);
CREATE INDEX idx_posts_identity_id ON posts(identity_id);
CREATE INDEX idx_posts_status ON posts(status);
CREATE INDEX idx_posts_created_at ON posts(created_at DESC);
CREATE INDEX idx_posts_like_count ON posts(like_count DESC);
CREATE INDEX idx_posts_post_type ON posts(post_type);

-- 评论表索引
CREATE INDEX idx_comments_post_id ON comments(post_id);
CREATE INDEX idx_comments_user_id ON comments(user_id);
CREATE INDEX idx_comments_parent_id ON comments(parent_id);
CREATE INDEX idx_comments_created_at ON comments(created_at);

-- 点赞表索引
CREATE INDEX idx_post_likes_post_id ON post_likes(post_id);
CREATE INDEX idx_post_likes_user_id ON post_likes(user_id);
CREATE INDEX idx_comment_likes_comment_id ON comment_likes(comment_id);
CREATE INDEX idx_comment_likes_user_id ON comment_likes(user_id);

-- 预约表索引
CREATE INDEX idx_appointments_student_id ON appointments(student_id);
CREATE INDEX idx_appointments_mentor_id ON appointments(mentor_id);
CREATE INDEX idx_appointments_appointment_time ON appointments(appointment_time);
CREATE INDEX idx_appointments_status ON appointments(status);
CREATE INDEX idx_appointments_meeting_type ON appointments(meeting_type);

-- 评价表索引
CREATE INDEX idx_reviews_reviewer_id ON reviews(reviewer_id);
CREATE INDEX idx_reviews_reviewed_id ON reviews(reviewed_id);
CREATE INDEX idx_reviews_course_id ON reviews(course_id);
CREATE INDEX idx_reviews_appointment_id ON reviews(appointment_id);
CREATE INDEX idx_reviews_rating ON reviews(rating);
CREATE INDEX idx_reviews_review_type ON reviews(review_type);
CREATE INDEX idx_reviews_created_at ON reviews(created_at);

-- 消息表索引
CREATE INDEX idx_messages_from_user_id ON messages(from_user_id);
CREATE INDEX idx_messages_to_user_id ON messages(to_user_id);
CREATE INDEX idx_messages_circle_id ON messages(circle_id);
CREATE INDEX idx_messages_is_read ON messages(is_read);
CREATE INDEX idx_messages_created_at ON messages(created_at);

-- 通知表索引
CREATE INDEX idx_notifications_user_id ON notifications(user_id);
CREATE INDEX idx_notifications_is_read ON notifications(is_read);
CREATE INDEX idx_notifications_created_at ON notifications(created_at);

-- 系统配置表索引
CREATE INDEX idx_system_configs_key ON system_configs(config_key);

-- 审计日志表索引
CREATE INDEX idx_audit_logs_table_name ON audit_logs(table_name);
CREATE INDEX idx_audit_logs_operation ON audit_logs(operation);
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);

-- 创建视图
-- 大师统计视图
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

-- 课程统计视图
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

-- 插入初始数据
-- 插入专业领域数据
INSERT INTO domains (name, code, description, sort_order) VALUES
('软件开发', 'software_development', '包括各种编程语言和开发技术', 1),
('UI/UX设计', 'ui_ux_design', '用户界面和用户体验设计', 2),
('数字营销', 'digital_marketing', '网络营销和推广策略', 3),
('传统工艺', 'traditional_crafts', '传统手工艺和技艺传承', 4),
('音乐艺术', 'music_arts', '音乐创作和表演艺术', 5),
('摄影摄像', 'photography_videography', '摄影和视频制作技术', 6),
('商业管理', 'business_management', '企业管理和商业策略', 7),
('语言学习', 'language_learning', '外语学习和语言技能', 8);

-- 插入系统配置数据
INSERT INTO system_configs (config_key, config_value, config_type, description) VALUES
('site_name', 'Master Guide', 'string', '网站名称'),
('site_description', '技艺传承平台', 'string', '网站描述'),
('max_file_size', '10485760', 'number', '最大文件上传大小（字节）'),
('allowed_file_types', '["jpg","jpeg","png","gif","mp4","mov","pdf","doc","docx"]', 'json', '允许上传的文件类型'),
('course_commission_rate', '0.1', 'number', '课程佣金比例'),
('appointment_commission_rate', '0.15', 'number', '预约佣金比例'),
('max_course_students', '1000', 'number', '课程最大学生数'),
('min_course_price', '9.99', 'number', '课程最低价格'),
('max_course_price', '9999.99', 'number', '课程最高价格'),
('verification_required', 'true', 'boolean', '大师身份是否需要认证');

-- 创建只读用户
CREATE USER readonly WITH PASSWORD 'readonly_password';
GRANT CONNECT ON DATABASE master_guide TO readonly;
GRANT USAGE ON SCHEMA public TO readonly;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly;
GRANT SELECT ON ALL SEQUENCES IN SCHEMA public TO readonly;

-- 创建应用用户
CREATE USER app_user WITH PASSWORD 'app_password';
GRANT CONNECT ON DATABASE master_guide TO app_user;
GRANT USAGE ON SCHEMA public TO app_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO app_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO app_user;

-- 设置默认权限
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO readonly;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO app_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO app_user;

-- 创建更新时间戳触发器函数
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 为需要更新时间戳的表创建触发器
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_user_identities_updated_at BEFORE UPDATE ON user_identities FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_user_profiles_updated_at BEFORE UPDATE ON user_profiles FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_courses_updated_at BEFORE UPDATE ON courses FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_course_contents_updated_at BEFORE UPDATE ON course_contents FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_circles_updated_at BEFORE UPDATE ON circles FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_posts_updated_at BEFORE UPDATE ON posts FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_comments_updated_at BEFORE UPDATE ON comments FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_appointments_updated_at BEFORE UPDATE ON appointments FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_reviews_updated_at BEFORE UPDATE ON reviews FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_system_configs_updated_at BEFORE UPDATE ON system_configs FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- 创建统计更新函数
CREATE OR REPLACE FUNCTION update_course_stats()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' OR TG_OP = 'DELETE' THEN
        UPDATE courses 
        SET current_students = (
            SELECT COUNT(*) 
            FROM learning_records 
            WHERE course_id = COALESCE(NEW.course_id, OLD.course_id)
        )
        WHERE id = COALESCE(NEW.course_id, OLD.course_id);
        
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

CREATE OR REPLACE FUNCTION update_post_stats()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' OR TG_OP = 'DELETE' THEN
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

-- 创建触发器
CREATE TRIGGER trigger_update_course_stats
AFTER INSERT OR UPDATE OR DELETE ON learning_records
FOR EACH ROW EXECUTE FUNCTION update_course_stats();

CREATE TRIGGER trigger_update_course_stats_from_reviews
AFTER INSERT OR UPDATE OR DELETE ON reviews
FOR EACH ROW EXECUTE FUNCTION update_course_stats();

CREATE TRIGGER trigger_update_post_stats
AFTER INSERT OR UPDATE OR DELETE ON post_likes
FOR EACH ROW EXECUTE FUNCTION update_post_stats();

-- 创建测试数据（可选）
-- 插入测试用户
INSERT INTO users (email, password_hash, phone, status) VALUES
('admin@masterguide.com', crypt('admin123', gen_salt('bf')), '13800138000', 'active'),
('test@masterguide.com', crypt('test123', gen_salt('bf')), '13800138001', 'active');

-- 插入测试身份
INSERT INTO user_identities (user_id, identity_type, domain, status, verification_status) VALUES
((SELECT id FROM users WHERE email = 'admin@masterguide.com'), 'master', 'software_development', 'active', 'verified'),
((SELECT id FROM users WHERE email = 'test@masterguide.com'), 'apprentice', 'software_development', 'active', 'verified');

-- 插入测试档案
INSERT INTO user_profiles (user_id, identity_id, name, bio, skills, experience_years, hourly_rate) VALUES
((SELECT id FROM users WHERE email = 'admin@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'admin@masterguide.com')), 
 '张大师', '资深全栈开发工程师', ARRAY['Go', 'Vue.js', 'PostgreSQL'], 8, 200.00),
((SELECT id FROM users WHERE email = 'test@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'test@masterguide.com')), 
 '李同学', '热爱学习的新手', ARRAY['JavaScript', 'Vue.js'], 1, 50.00);

-- 完成初始化
COMMIT;

-- 显示初始化结果
SELECT 'Database initialization completed successfully!' as status;
SELECT COUNT(*) as total_tables FROM information_schema.tables WHERE table_schema = 'public';
SELECT COUNT(*) as total_indexes FROM pg_indexes WHERE schemaname = 'public'; 