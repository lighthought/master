-- Master Guide 数据库迁移文件
-- 版本: 001
-- 描述: 创建初始数据库架构
-- 创建时间: 2024-12-01

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

-- 设置序列所有者
ALTER SEQUENCE user_id_num_seq OWNER TO postgres;
ALTER SEQUENCE user_identity_id_num_seq OWNER TO postgres;
ALTER SEQUENCE user_profile_id_num_seq OWNER TO postgres;
ALTER SEQUENCE domain_id_num_seq OWNER TO postgres;

-- 创建用户基础表
CREATE TABLE IF NOT EXISTS users (
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
CREATE TABLE IF NOT EXISTS user_identities (
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
CREATE TABLE IF NOT EXISTS user_profiles (
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
CREATE TABLE IF NOT EXISTS domains (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('DOMAIN_', 'domain_id_num_seq'),
    name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    icon VARCHAR(500),
    sort_order INTEGER DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建基础索引
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_phone ON users(phone);
CREATE INDEX IF NOT EXISTS idx_users_status ON users(status);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at);

CREATE INDEX IF NOT EXISTS idx_user_identities_user_id ON user_identities(user_id);
CREATE INDEX IF NOT EXISTS idx_user_identities_type_domain ON user_identities(identity_type, domain);
CREATE INDEX IF NOT EXISTS idx_user_identities_status ON user_identities(status);
CREATE INDEX IF NOT EXISTS idx_user_identities_verification_status ON user_identities(verification_status);

CREATE INDEX IF NOT EXISTS idx_user_profiles_user_id ON user_profiles(user_id);
CREATE INDEX IF NOT EXISTS idx_user_profiles_identity_id ON user_profiles(identity_id);
CREATE INDEX IF NOT EXISTS idx_user_profiles_name ON user_profiles(name);
CREATE INDEX IF NOT EXISTS idx_user_profiles_skills ON user_profiles USING GIN(skills);

CREATE INDEX IF NOT EXISTS idx_domains_code ON domains(code);
CREATE INDEX IF NOT EXISTS idx_domains_is_active ON domains(is_active);
CREATE INDEX IF NOT EXISTS idx_domains_sort_order ON domains(sort_order);

-- 插入初始专业领域数据
INSERT INTO domains (name, code, description, sort_order) VALUES
('软件开发', 'software_development', '包括各种编程语言和开发技术', 1),
('UI/UX设计', 'ui_ux_design', '用户界面和用户体验设计', 2),
('数字营销', 'digital_marketing', '网络营销和推广策略', 3),
('传统工艺', 'traditional_crafts', '传统手工艺和技艺传承', 4),
('音乐艺术', 'music_arts', '音乐创作和表演艺术', 5),
('摄影摄像', 'photography_videography', '摄影和视频制作技术', 6),
('商业管理', 'business_management', '企业管理和商业策略', 7),
('语言学习', 'language_learning', '外语学习和语言技能', 8)
ON CONFLICT (code) DO NOTHING; 