-- Master Guide 数据库迁移文件
-- 版本: 003
-- 描述: 创建社群相关表
-- 创建时间: 2024-12-01

CREATE SEQUENCE IF NOT EXISTS circle_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS circle_member_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS post_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS comment_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS post_like_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;
CREATE SEQUENCE IF NOT EXISTS comment_like_id_num_seq INCREMENT BY 1 START 1 MINVALUE 1 MAXVALUE 99999999999 CACHE 1;

-- 设置序列所有者
ALTER SEQUENCE circle_id_num_seq OWNER TO postgres;
ALTER SEQUENCE circle_member_id_num_seq OWNER TO postgres;
ALTER SEQUENCE post_id_num_seq OWNER TO postgres;
ALTER SEQUENCE comment_id_num_seq OWNER TO postgres;
ALTER SEQUENCE post_like_id_num_seq OWNER TO postgres;
ALTER SEQUENCE comment_like_id_num_seq OWNER TO postgres;

-- 创建圈子表
CREATE TABLE IF NOT EXISTS circles (
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
CREATE TABLE IF NOT EXISTS circle_members (
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
CREATE TABLE IF NOT EXISTS posts (
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
CREATE TABLE IF NOT EXISTS comments (
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
CREATE TABLE IF NOT EXISTS post_likes (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('POSTLIKE_', 'post_like_id_num_seq'),
    post_id VARCHAR(32) NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(post_id, user_id)
);

-- 创建评论点赞表
CREATE TABLE IF NOT EXISTS comment_likes (
    id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('COMMENTLIKE_', 'comment_like_id_num_seq'),
    comment_id VARCHAR(32) NOT NULL REFERENCES comments(id) ON DELETE CASCADE,
    user_id VARCHAR(32) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(comment_id, user_id)
);

-- 创建社群索引
CREATE INDEX IF NOT EXISTS idx_circles_domain ON circles(domain);
CREATE INDEX IF NOT EXISTS idx_circles_status ON circles(status);
CREATE INDEX IF NOT EXISTS idx_circles_created_by ON circles(created_by);
CREATE INDEX IF NOT EXISTS idx_circles_member_count ON circles(member_count);
CREATE INDEX IF NOT EXISTS idx_circles_tags ON circles USING GIN(tags);

CREATE INDEX IF NOT EXISTS idx_circle_members_circle_id ON circle_members(circle_id);
CREATE INDEX IF NOT EXISTS idx_circle_members_user_id ON circle_members(user_id);
CREATE INDEX IF NOT EXISTS idx_circle_members_role ON circle_members(role);

CREATE INDEX IF NOT EXISTS idx_posts_circle_id ON posts(circle_id);
CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);
CREATE INDEX IF NOT EXISTS idx_posts_identity_id ON posts(identity_id);
CREATE INDEX IF NOT EXISTS idx_posts_status ON posts(status);
CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_posts_like_count ON posts(like_count DESC);
CREATE INDEX IF NOT EXISTS idx_posts_post_type ON posts(post_type);

CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);
CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);
CREATE INDEX IF NOT EXISTS idx_comments_parent_id ON comments(parent_id);
CREATE INDEX IF NOT EXISTS idx_comments_created_at ON comments(created_at);

CREATE INDEX IF NOT EXISTS idx_post_likes_post_id ON post_likes(post_id);
CREATE INDEX IF NOT EXISTS idx_post_likes_user_id ON post_likes(user_id);
CREATE INDEX IF NOT EXISTS idx_comment_likes_comment_id ON comment_likes(comment_id);
CREATE INDEX IF NOT EXISTS idx_comment_likes_user_id ON comment_likes(user_id); 