-- Master Guide 数据库迁移文件
-- 版本: 002
-- 描述: 创建课程相关表
-- 创建时间: 2024-12-01

-- 创建课程相关序列
CREATE SEQUENCE IF NOT EXISTS course_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS course_content_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS learning_record_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS content_progress_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;

-- 设置序列所有者
ALTER SEQUENCE course_id_num_seq OWNER TO postgres;
ALTER SEQUENCE course_content_id_num_seq OWNER TO postgres;
ALTER SEQUENCE learning_record_id_num_seq OWNER TO postgres;
ALTER SEQUENCE content_progress_id_num_seq OWNER TO postgres;

-- 创建课程表
CREATE TABLE IF NOT EXISTS courses (
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
CREATE TABLE IF NOT EXISTS course_contents (
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
CREATE TABLE IF NOT EXISTS learning_records (
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
CREATE TABLE IF NOT EXISTS content_progress (
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

-- 创建课程索引
CREATE INDEX IF NOT EXISTS idx_courses_mentor_id ON courses(mentor_id);
CREATE INDEX IF NOT EXISTS idx_courses_status ON courses(status);
CREATE INDEX IF NOT EXISTS idx_courses_difficulty ON courses(difficulty);
CREATE INDEX IF NOT EXISTS idx_courses_price ON courses(price);
CREATE INDEX IF NOT EXISTS idx_courses_rating ON courses(rating);
CREATE INDEX IF NOT EXISTS idx_courses_created_at ON courses(created_at);
CREATE INDEX IF NOT EXISTS idx_courses_tags ON courses USING GIN(tags);

CREATE INDEX IF NOT EXISTS idx_course_contents_course_id ON course_contents(course_id);
CREATE INDEX IF NOT EXISTS idx_course_contents_order_index ON course_contents(course_id, order_index);

CREATE INDEX IF NOT EXISTS idx_learning_records_user_id ON learning_records(user_id);
CREATE INDEX IF NOT EXISTS idx_learning_records_course_id ON learning_records(course_id);
CREATE INDEX IF NOT EXISTS idx_learning_records_status ON learning_records(status);
CREATE INDEX IF NOT EXISTS idx_learning_records_enrolled_at ON learning_records(enrolled_at);

CREATE INDEX IF NOT EXISTS idx_content_progress_user_id ON content_progress(user_id);
CREATE INDEX IF NOT EXISTS idx_content_progress_content_id ON content_progress(content_id);
CREATE INDEX IF NOT EXISTS idx_content_progress_is_completed ON content_progress(is_completed); 