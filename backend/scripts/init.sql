-- Master Guide 数据库初始化脚本
-- 版本: v1.0
-- 创建时间: 2024-12-01
-- 描述: Master Guide平台数据库初始化

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

ALTER FUNCTION generate_table_id(VARCHAR(32), VARCHAR(50)) OWNER TO master_guide;
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
CREATE SEQUENCE IF NOT EXISTS mentor_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS mentor_review_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS study_session_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS assignment_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS income_transaction_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS withdrawal_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;


-- 设置序列所有者
ALTER SEQUENCE user_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE user_identity_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE user_profile_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE domain_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE course_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE course_content_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE learning_record_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE content_progress_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE circle_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE circle_member_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE post_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE comment_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE post_like_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE comment_like_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE appointment_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE review_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE message_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE notification_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE system_config_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE audit_log_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE mentor_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE mentor_review_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE study_session_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE assignment_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE income_transaction_id_num_seq OWNER TO master_guide;
ALTER SEQUENCE withdrawal_id_num_seq OWNER TO master_guide;


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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建大师表
CREATE TABLE mentors (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('MENTOR_', 'mentor_id_num_seq'),
    identity_id VARCHAR(32) NOT NULL REFERENCES user_identities(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    rating DECIMAL(3,2) DEFAULT 0.00 CHECK (rating >= 0 AND rating <= 5),
    student_count INTEGER DEFAULT 0,
    hourly_rate DECIMAL(10,2) NOT NULL,
    is_online BOOLEAN DEFAULT FALSE,
    experience_years INTEGER DEFAULT 0,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'suspended')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建大师评价表
CREATE TABLE mentor_reviews (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('MENTOR_REVIEW_', 'mentor_review_id_num_seq'),
    mentor_id VARCHAR(32) NOT NULL REFERENCES mentors(id) ON DELETE CASCADE,
    reviewer_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    content TEXT,
    is_anonymous BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



-- 创建领域表
CREATE TABLE domains (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('DOMAIN_', 'domain_id_num_seq'),
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    icon VARCHAR(500),
    is_active BOOLEAN DEFAULT TRUE,
    sort_order INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建课程表
CREATE TABLE courses (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('COURSE_', 'course_id_num_seq'),
    mentor_id VARCHAR(32) NOT NULL REFERENCES mentors(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    cover_image VARCHAR(500),
    price DECIMAL(10,2) NOT NULL,
    duration_hours INTEGER NOT NULL,
    difficulty VARCHAR(20) CHECK (difficulty IN ('beginner', 'intermediate', 'advanced')),
    status VARCHAR(20) DEFAULT 'draft' CHECK (status IN ('draft', 'published', 'archived')),
    max_students INTEGER,
    rating DECIMAL(3,2) DEFAULT 0,
    review_count INTEGER DEFAULT 0,
    tags TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建课程内容表
CREATE TABLE course_contents (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('CONTENT_', 'course_content_id_num_seq'),
    course_id VARCHAR(32) NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    content_type VARCHAR(20) NOT NULL CHECK (content_type IN ('video', 'text', 'quiz')),
    content_url VARCHAR(500),
    content_text TEXT,
    order_index INTEGER NOT NULL,
    duration_minutes INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建学习记录表
CREATE TABLE learning_records (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('LEARNING_', 'learning_record_id_num_seq'),
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id VARCHAR(32) NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    progress_percentage DECIMAL(5,2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'enrolled' CHECK (status IN ('enrolled', 'learning', 'completed', 'dropped', 'paused')),
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    last_accessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    current_chapter VARCHAR(200),
    completed_chapters TEXT[],
    total_study_time INTEGER DEFAULT 0,
    certificate_issued BOOLEAN DEFAULT FALSE,
    certificate_url VARCHAR(500)
);

-- 创建学习会话表
CREATE TABLE study_sessions (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('SESSION_', 'study_session_id_num_seq'),
    learning_record_id VARCHAR(32) NOT NULL REFERENCES learning_records(id) ON DELETE CASCADE,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    duration_minutes INTEGER DEFAULT 0,
    chapter VARCHAR(200),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建作业表
CREATE TABLE assignments (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('ASSIGNMENT_', 'assignment_id_num_seq'),
    learning_record_id VARCHAR(32) NOT NULL REFERENCES learning_records(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    content TEXT,
    attachment_urls TEXT[],
    status VARCHAR(20) DEFAULT 'submitted' CHECK (status IN ('submitted', 'reviewed', 'approved', 'rejected')),
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    reviewed_at TIMESTAMP,
    score DECIMAL(5,2),
    feedback TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建收入交易表
CREATE TABLE income_transactions (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('INCOME_', 'income_transaction_id_num_seq'),
    mentor_id VARCHAR(32) NOT NULL REFERENCES mentors(id) ON DELETE CASCADE,
    student_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    transaction_type VARCHAR(20) NOT NULL CHECK (transaction_type IN ('course_enrollment', 'appointment', 'refund')),
    amount DECIMAL(10,2) NOT NULL,
    platform_fee DECIMAL(10,2) NOT NULL DEFAULT 0,
    net_income DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'failed')),
    description TEXT,
    course_id VARCHAR(32) REFERENCES courses(id) ON DELETE SET NULL,
    appointment_id VARCHAR(32) REFERENCES appointments(id) ON DELETE SET NULL,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建提现表
CREATE TABLE withdrawals (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('WITHDRAWAL_', 'withdrawal_id_num_seq'),
    mentor_id VARCHAR(32) NOT NULL REFERENCES mentors(id) ON DELETE CASCADE,
    amount DECIMAL(10,2) NOT NULL,
    fee DECIMAL(10,2) NOT NULL DEFAULT 0,
    net_amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'failed')),
    bank_account VARCHAR(50) NOT NULL,
    bank_name VARCHAR(100) NOT NULL,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建内容进度表
CREATE TABLE content_progress (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('PROGRESS_', 'content_progress_id_num_seq'),
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content_id VARCHAR(32) NOT NULL REFERENCES course_contents(id) ON DELETE CASCADE,
    is_completed BOOLEAN DEFAULT FALSE,
    progress_percentage DECIMAL(5,2) DEFAULT 0,
    study_time_minutes INTEGER DEFAULT 0,
    last_accessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
    tags TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建圈子成员表
CREATE TABLE circle_members (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('MEMBER_', 'circle_member_id_num_seq'),
    circle_id VARCHAR(32) NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(20) DEFAULT 'member' CHECK (role IN ('admin', 'moderator', 'member')),
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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
    post_type VARCHAR(20) DEFAULT 'text' CHECK (post_type IN ('text', 'image', 'video', 'link')),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'hidden', 'deleted')),
    like_count INTEGER DEFAULT 0,
    comment_count INTEGER DEFAULT 0,
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
    like_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建动态点赞表
CREATE TABLE post_likes (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('LIKE_', 'post_like_id_num_seq'),
    post_id VARCHAR(32) NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(post_id, user_id)
);

-- 创建评论点赞表
CREATE TABLE comment_likes (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('COMMENT_LIKE_', 'comment_like_id_num_seq'),
    comment_id VARCHAR(32) NOT NULL REFERENCES comments(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(comment_id, user_id)
);

-- 创建预约表
CREATE TABLE appointments (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('APPOINTMENT_', 'appointment_id_num_seq'),
    student_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    mentor_id VARCHAR(32) NOT NULL REFERENCES mentors(id) ON DELETE CASCADE,
    appointment_time TIMESTAMP NOT NULL,
    duration_minutes INTEGER NOT NULL,
    meeting_type VARCHAR(20) DEFAULT 'video' CHECK (meeting_type IN ('video', 'voice', 'text')),
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'confirmed', 'completed', 'cancelled')),
    price DECIMAL(10,2) NOT NULL,
    notes TEXT,
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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建消息表
CREATE TABLE messages (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('MESSAGE_', 'message_id_num_seq'),
    from_user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    to_user_id VARCHAR(32) REFERENCES users(id) ON DELETE CASCADE,
    circle_id VARCHAR(32) REFERENCES circles(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    message_type VARCHAR(20) DEFAULT 'text' CHECK (message_type IN ('text', 'image', 'file', 'system')),
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建通知表
CREATE TABLE notifications (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('NOTIFICATION_', 'notification_id_num_seq'),
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    notification_type VARCHAR(50) NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建系统配置表
CREATE TABLE system_configs (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('CONFIG_', 'system_config_id_num_seq'),
    config_key VARCHAR(100) UNIQUE NOT NULL,
    config_value TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建审计日志表
CREATE TABLE audit_logs (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('AUDIT_', 'audit_log_id_num_seq'),
    user_id VARCHAR(32) REFERENCES users(id) ON DELETE SET NULL,
    table_name VARCHAR(100) NOT NULL,
    record_id VARCHAR(32) NOT NULL,
    operation VARCHAR(20) NOT NULL CHECK (operation IN ('CREATE', 'UPDATE', 'DELETE')),
    old_values JSONB,
    new_values JSONB,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建索引
-- 用户表索引
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_users_created_at ON users(created_at);

-- 身份表索引
CREATE INDEX idx_user_identities_user_id ON user_identities(user_id);
CREATE INDEX idx_user_identities_type_domain ON user_identities(identity_type, domain);
CREATE INDEX idx_user_identities_status ON user_identities(status);
CREATE INDEX idx_user_identities_verification_status ON user_identities(verification_status);

-- 档案表索引
CREATE INDEX idx_user_profiles_user_id ON user_profiles(user_id);
CREATE INDEX idx_user_profiles_identity_id ON user_profiles(identity_id);
CREATE INDEX idx_user_profiles_name ON user_profiles(name);
CREATE INDEX idx_user_profiles_skills ON user_profiles USING GIN(skills);

-- 大师表索引
CREATE INDEX idx_mentors_identity_id ON mentors(identity_id);
CREATE INDEX idx_mentors_user_id ON mentors(user_id);
CREATE INDEX idx_mentors_rating ON mentors(rating);
CREATE INDEX idx_mentors_hourly_rate ON mentors(hourly_rate);
CREATE INDEX idx_mentors_is_online ON mentors(is_online);
CREATE INDEX idx_mentors_status ON mentors(status);

-- 大师评价表索引
CREATE INDEX idx_mentor_reviews_mentor_id ON mentor_reviews(mentor_id);
CREATE INDEX idx_mentor_reviews_reviewer_id ON mentor_reviews(reviewer_id);
CREATE INDEX idx_mentor_reviews_rating ON mentor_reviews(rating);
CREATE INDEX idx_mentor_reviews_created_at ON mentor_reviews(created_at);

-- 领域表索引
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

-- 课程内容索引
CREATE INDEX idx_course_contents_course_id ON course_contents(course_id);
CREATE INDEX idx_course_contents_order_index ON course_contents(course_id, order_index);

-- 学习记录索引
CREATE INDEX idx_learning_records_user_id ON learning_records(user_id);
CREATE INDEX idx_learning_records_course_id ON learning_records(course_id);
CREATE INDEX idx_learning_records_status ON learning_records(status);
CREATE INDEX idx_learning_records_enrolled_at ON learning_records(enrolled_at);

-- 内容进度索引
CREATE INDEX idx_content_progress_user_id ON content_progress(user_id);
CREATE INDEX idx_content_progress_content_id ON content_progress(content_id);
CREATE INDEX idx_content_progress_is_completed ON content_progress(is_completed);

-- 学习会话索引
CREATE INDEX idx_study_sessions_learning_record_id ON study_sessions(learning_record_id);
CREATE INDEX idx_study_sessions_start_time ON study_sessions(start_time);
CREATE INDEX idx_study_sessions_chapter ON study_sessions(chapter);

-- 作业索引
CREATE INDEX idx_assignments_learning_record_id ON assignments(learning_record_id);
CREATE INDEX idx_assignments_status ON assignments(status);
CREATE INDEX idx_assignments_submitted_at ON assignments(submitted_at);
CREATE INDEX idx_assignments_score ON assignments(score);

-- 收入交易索引
CREATE INDEX idx_income_transactions_mentor_id ON income_transactions(mentor_id);
CREATE INDEX idx_income_transactions_student_id ON income_transactions(student_id);
CREATE INDEX idx_income_transactions_type ON income_transactions(transaction_type);
CREATE INDEX idx_income_transactions_status ON income_transactions(status);
CREATE INDEX idx_income_transactions_created_at ON income_transactions(created_at);
CREATE INDEX idx_income_transactions_course_id ON income_transactions(course_id);
CREATE INDEX idx_income_transactions_appointment_id ON income_transactions(appointment_id);

-- 提现索引
CREATE INDEX idx_withdrawals_mentor_id ON withdrawals(mentor_id);
CREATE INDEX idx_withdrawals_status ON withdrawals(status);
CREATE INDEX idx_withdrawals_created_at ON withdrawals(created_at);

-- 圈子表索引
CREATE INDEX idx_circles_domain ON circles(domain);
CREATE INDEX idx_circles_status ON circles(status);
CREATE INDEX idx_circles_created_by ON circles(created_by);
CREATE INDEX idx_circles_member_count ON circles(member_count);
CREATE INDEX idx_circles_tags ON circles USING GIN(tags);

-- 圈子成员索引
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

-- 系统配置索引
CREATE INDEX idx_system_configs_key ON system_configs(config_key);

-- 审计日志索引
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
    ui.domain,
    COUNT(DISTINCT c.id) as total_courses,
    COUNT(DISTINCT lr.user_id) as total_students,
    AVG(c.rating) as average_rating,
    COUNT(DISTINCT r.id) as total_reviews,
    SUM(c.price * lr.progress_percentage / 100) as estimated_earnings
FROM user_identities ui
LEFT JOIN user_profiles up ON ui.id = up.identity_id
LEFT JOIN courses c ON ui.id = c.mentor_id
LEFT JOIN learning_records lr ON c.id = lr.course_id
LEFT JOIN reviews r ON ui.id = r.reviewed_id AND r.review_type = 'mentor'
WHERE ui.identity_type = 'master' AND ui.status = 'active'
GROUP BY ui.id, ui.user_id, up.name, ui.domain;

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
    c.rating,
    c.review_count,
    COUNT(DISTINCT lr.user_id) as enrolled_students,
    COUNT(DISTINCT CASE WHEN lr.status = 'completed' THEN lr.user_id END) as completed_students,
    AVG(lr.progress_percentage) as average_progress,
    AVG(r.rating) as average_review_rating,
    COUNT(DISTINCT r.id) as total_reviews
FROM courses c
LEFT JOIN user_profiles up ON c.mentor_id = up.identity_id
LEFT JOIN learning_records lr ON c.id = lr.course_id
LEFT JOIN reviews r ON c.id = r.course_id AND r.review_type = 'course'
WHERE c.status = 'published'
GROUP BY c.id, c.title, c.mentor_id, up.name, c.price, c.duration_hours, c.difficulty, c.rating, c.review_count;

-- 创建触发器函数
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 创建触发器
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_user_identities_updated_at BEFORE UPDATE ON user_identities FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_user_profiles_updated_at BEFORE UPDATE ON user_profiles FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_mentors_updated_at BEFORE UPDATE ON mentors FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_mentor_reviews_updated_at BEFORE UPDATE ON mentor_reviews FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_courses_updated_at BEFORE UPDATE ON courses FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_course_contents_updated_at BEFORE UPDATE ON course_contents FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_circles_updated_at BEFORE UPDATE ON circles FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_posts_updated_at BEFORE UPDATE ON posts FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_comments_updated_at BEFORE UPDATE ON comments FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_appointments_updated_at BEFORE UPDATE ON appointments FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_reviews_updated_at BEFORE UPDATE ON reviews FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_system_configs_updated_at BEFORE UPDATE ON system_configs FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_income_transactions_updated_at BEFORE UPDATE ON income_transactions FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_withdrawals_updated_at BEFORE UPDATE ON withdrawals FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- 创建统计更新函数
CREATE OR REPLACE FUNCTION update_course_stats()
RETURNS TRIGGER AS $$
BEGIN
    -- 更新课程统计
    UPDATE courses 
    SET 
        rating = (
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
    
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_post_stats()
RETURNS TRIGGER AS $$
BEGIN
    -- 更新动态统计
    UPDATE posts 
    SET 
        like_count = (
            SELECT COUNT(*) 
            FROM post_likes 
            WHERE post_id = COALESCE(NEW.post_id, OLD.post_id)
        )
    WHERE id = COALESCE(NEW.post_id, OLD.post_id);
    
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;

-- 创建统计触发器
CREATE TRIGGER trigger_update_course_stats
AFTER INSERT OR UPDATE OR DELETE ON learning_records
FOR EACH ROW EXECUTE FUNCTION update_course_stats();

CREATE TRIGGER trigger_update_course_stats_from_reviews
AFTER INSERT OR UPDATE OR DELETE ON reviews
FOR EACH ROW EXECUTE FUNCTION update_course_stats();

CREATE TRIGGER trigger_update_post_stats
AFTER INSERT OR UPDATE OR DELETE ON post_likes
FOR EACH ROW EXECUTE FUNCTION update_post_stats();

-- 插入初始数据
INSERT INTO domains (code, name, description, icon, sort_order) VALUES
('software_development', '软件开发', '编程、算法、架构设计等软件开发相关领域', '💻', 1),
('ui_design', 'UI设计', '用户界面设计、用户体验设计等', '🎨', 2),
('digital_marketing', '数字营销', '社交媒体营销、SEO、内容营销等', '📱', 3),
('traditional_craft', '传统工艺', '木工、陶艺、编织等传统手工艺', '🏺', 4),
('cooking', '烹饪', '中餐、西餐、烘焙等烹饪技艺', '👨‍🍳', 5),
('music', '音乐', '乐器演奏、作曲、音乐制作等', '🎵', 6),
('fitness', '健身', '力量训练、瑜伽、有氧运动等', '💪', 7),
('photography', '摄影', '人像摄影、风景摄影、商业摄影等', '📸', 8);

-- 插入系统配置
INSERT INTO system_configs (config_key, config_value, description) VALUES
('app_name', 'Master Guide', '应用名称'),
('app_version', '1.0.0', '应用版本'),
('maintenance_mode', 'false', '维护模式'),
('registration_enabled', 'true', '是否允许注册'),
('email_verification_required', 'true', '是否需要邮箱验证'),
('max_file_size', '10485760', '最大文件上传大小（字节）'),
('allowed_file_types', '["image/jpeg","image/png","image/gif","application/pdf","text/plain"]', '允许的文件类型'),
('default_avatar', 'https://example.com/default-avatar.png', '默认头像URL'); 

-- 支付相关表

-- 支付订单表
CREATE TABLE payment_orders (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('PAYORDER_', 'payment_order_id_num_seq'),
    order_type VARCHAR(32) NOT NULL CHECK (order_type IN ('course_enrollment', 'appointment', 'refund')),
    order_ref_id VARCHAR(32) NOT NULL, -- 业务订单ID，如课程报名ID、预约ID等
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'CNY',
    payment_method VARCHAR(32) NOT NULL,
    description TEXT,
    metadata JSONB,
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'failed', 'cancelled')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    UNIQUE(order_type, order_ref_id)
);

-- 支付流水表
CREATE TABLE payment_records (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('PAYREC_', 'payment_record_id_num_seq'),
    order_id VARCHAR(32) NOT NULL REFERENCES payment_orders(id) ON DELETE CASCADE,
    payment_url VARCHAR(500),
    qr_code VARCHAR(500),
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'failed', 'cancelled')),
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'CNY',
    payment_method VARCHAR(32) NOT NULL,
    paid_at TIMESTAMP,
    transaction_id VARCHAR(64),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 退款表
CREATE TABLE payment_refunds (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('REFUND_', 'payment_refund_id_num_seq'),
    payment_id VARCHAR(32) NOT NULL REFERENCES payment_records(id) ON DELETE CASCADE,
    amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'failed', 'cancelled')),
    reason TEXT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    refund_transaction_id VARCHAR(64)
);

-- 支付方式表
CREATE TABLE payment_methods (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    icon VARCHAR(500),
    enabled BOOLEAN DEFAULT TRUE,
    min_amount DECIMAL(10,2) DEFAULT 0.01,
    max_amount DECIMAL(10,2) DEFAULT 100000.00
);

-- 支付相关索引
CREATE INDEX idx_payment_orders_status ON payment_orders(status);
CREATE INDEX idx_payment_orders_created_at ON payment_orders(created_at);
CREATE INDEX idx_payment_records_order_id ON payment_records(order_id);
CREATE INDEX idx_payment_records_status ON payment_records(status);
CREATE INDEX idx_payment_records_paid_at ON payment_records(paid_at);
CREATE INDEX idx_payment_refunds_payment_id ON payment_refunds(payment_id);
CREATE INDEX idx_payment_refunds_status ON payment_refunds(status);
CREATE INDEX idx_payment_methods_enabled ON payment_methods(enabled);

-- 支付相关序列
CREATE SEQUENCE IF NOT EXISTS payment_order_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS payment_record_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS payment_refund_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1; 