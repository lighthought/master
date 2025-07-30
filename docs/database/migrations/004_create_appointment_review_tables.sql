-- Master Guide 数据库迁移文件
-- 版本: 004
-- 描述: 创建预约和评价相关表
-- 创建时间: 2024-12-01

CREATE SEQUENCE IF NOT EXISTS appointment_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS review_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;

-- 设置序列所有者
ALTER SEQUENCE appointment_id_num_seq OWNER TO postgres;
ALTER SEQUENCE review_id_num_seq OWNER TO postgres;

-- 创建预约表
CREATE TABLE IF NOT EXISTS appointments (
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
CREATE TABLE IF NOT EXISTS reviews (
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

-- 创建预约索引
CREATE INDEX IF NOT EXISTS idx_appointments_student_id ON appointments(student_id);
CREATE INDEX IF NOT EXISTS idx_appointments_mentor_id ON appointments(mentor_id);
CREATE INDEX IF NOT EXISTS idx_appointments_appointment_time ON appointments(appointment_time);
CREATE INDEX IF NOT EXISTS idx_appointments_status ON appointments(status);
CREATE INDEX IF NOT EXISTS idx_appointments_meeting_type ON appointments(meeting_type);

-- 创建评价索引
CREATE INDEX IF NOT EXISTS idx_reviews_reviewer_id ON reviews(reviewer_id);
CREATE INDEX IF NOT EXISTS idx_reviews_reviewed_id ON reviews(reviewed_id);
CREATE INDEX IF NOT EXISTS idx_reviews_course_id ON reviews(course_id);
CREATE INDEX IF NOT EXISTS idx_reviews_appointment_id ON reviews(appointment_id);
CREATE INDEX IF NOT EXISTS idx_reviews_rating ON reviews(rating);
CREATE INDEX IF NOT EXISTS idx_reviews_review_type ON reviews(review_type);
CREATE INDEX IF NOT EXISTS idx_reviews_created_at ON reviews(created_at); 