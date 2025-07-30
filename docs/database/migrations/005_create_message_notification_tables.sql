-- Master Guide 数据库迁移文件
-- 版本: 005
-- 描述: 创建消息和通知相关表
-- 创建时间: 2024-12-01



CREATE SEQUENCE IF NOT EXISTS message_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS notification_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS system_config_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS audit_log_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;

-- 设置序列所有者
ALTER SEQUENCE message_id_num_seq OWNER TO postgres;
ALTER SEQUENCE notification_id_num_seq OWNER TO postgres;
ALTER SEQUENCE system_config_id_num_seq OWNER TO postgres;
ALTER SEQUENCE audit_log_id_num_seq OWNER TO postgres;

-- 创建消息表
CREATE TABLE IF NOT EXISTS messages (
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
CREATE TABLE IF NOT EXISTS notifications (
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
CREATE TABLE IF NOT EXISTS system_configs (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('CONFIG_', 'system_config_id_num_seq'),
    config_key VARCHAR(100) NOT NULL UNIQUE,
    config_value TEXT,
    config_type VARCHAR(20) DEFAULT 'string' CHECK (config_type IN ('string', 'number', 'boolean', 'json')),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建审计日志表
CREATE TABLE IF NOT EXISTS audit_logs (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('AUDIT_', 'audit_log_id_num_seq'),
    table_name VARCHAR(100) NOT NULL,
    operation VARCHAR(20) NOT NULL,
    old_data JSONB,
    new_data JSONB,
    user_id VARCHAR(32),
    ip_address INET,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建消息索引
CREATE INDEX IF NOT EXISTS idx_messages_from_user_id ON messages(from_user_id);
CREATE INDEX IF NOT EXISTS idx_messages_to_user_id ON messages(to_user_id);
CREATE INDEX IF NOT EXISTS idx_messages_circle_id ON messages(circle_id);
CREATE INDEX IF NOT EXISTS idx_messages_is_read ON messages(is_read);
CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages(created_at);

-- 创建通知索引
CREATE INDEX IF NOT EXISTS idx_notifications_user_id ON notifications(user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_is_read ON notifications(is_read);
CREATE INDEX IF NOT EXISTS idx_notifications_created_at ON notifications(created_at);

-- 创建系统配置索引
CREATE INDEX IF NOT EXISTS idx_system_configs_key ON system_configs(config_key);

-- 创建审计日志索引
CREATE INDEX IF NOT EXISTS idx_audit_logs_table_name ON audit_logs(table_name);
CREATE INDEX IF NOT EXISTS idx_audit_logs_operation ON audit_logs(operation);
CREATE INDEX IF NOT EXISTS idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX IF NOT EXISTS idx_audit_logs_created_at ON audit_logs(created_at);

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
('verification_required', 'true', 'boolean', '大师身份是否需要认证')
ON CONFLICT (config_key) DO NOTHING; 