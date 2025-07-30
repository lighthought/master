-- Master Guide 数据库迁移文件
-- 版本: 006
-- 描述: 创建数据库视图和触发器
-- 创建时间: 2024-12-01

-- 创建大师统计视图
CREATE OR REPLACE VIEW mentor_stats AS
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

-- 创建课程统计视图
CREATE OR REPLACE VIEW course_stats AS
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

-- 创建用户学习统计视图
CREATE OR REPLACE VIEW user_learning_stats AS
SELECT 
    u.id as user_id,
    u.email,
    COUNT(DISTINCT lr.course_id) as enrolled_courses,
    COUNT(DISTINCT CASE WHEN lr.status = 'completed' THEN lr.course_id END) as completed_courses,
    SUM(lr.total_study_time) as total_study_time,
    AVG(lr.progress_percentage) as avg_progress
FROM users u
LEFT JOIN learning_records lr ON u.id = lr.user_id
GROUP BY u.id, u.email;

-- 创建用户教学统计视图
CREATE OR REPLACE VIEW user_teaching_stats AS
SELECT 
    u.id as user_id,
    u.email,
    COUNT(DISTINCT c.id) as total_courses,
    COUNT(DISTINCT lr.user_id) as total_students,
    SUM(c.price * lr.progress_percentage / 100) as estimated_income,
    AVG(r.rating) as average_rating
FROM users u
LEFT JOIN user_identities ui ON u.id = ui.user_id AND ui.identity_type = 'master'
LEFT JOIN courses c ON ui.id = c.mentor_id
LEFT JOIN learning_records lr ON c.id = lr.course_id
LEFT JOIN reviews r ON ui.id = r.reviewed_id AND r.review_type = 'mentor'
GROUP BY u.id, u.email;

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

-- 创建动态统计更新函数
CREATE OR REPLACE FUNCTION update_post_stats()
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

-- 创建评论统计更新函数
CREATE OR REPLACE FUNCTION update_comment_stats()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' OR TG_OP = 'DELETE' THEN
        -- 更新动态评论数
        UPDATE posts 
        SET comment_count = (
            SELECT COUNT(*) 
            FROM comments 
            WHERE post_id = COALESCE(NEW.post_id, OLD.post_id)
            AND status = 'active'
        )
        WHERE id = COALESCE(NEW.post_id, OLD.post_id);
        
        -- 更新评论点赞数
        UPDATE comments 
        SET like_count = (
            SELECT COUNT(*) 
            FROM comment_likes 
            WHERE comment_id = COALESCE(NEW.comment_id, OLD.comment_id)
        )
        WHERE id = COALESCE(NEW.comment_id, OLD.comment_id);
    END IF;
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;

-- 创建圈子统计更新函数
CREATE OR REPLACE FUNCTION update_circle_stats()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' OR TG_OP = 'DELETE' THEN
        -- 更新圈子成员数
        UPDATE circles 
        SET member_count = (
            SELECT COUNT(*) 
            FROM circle_members 
            WHERE circle_id = COALESCE(NEW.circle_id, OLD.circle_id)
            AND is_active = true
        )
        WHERE id = COALESCE(NEW.circle_id, OLD.circle_id);
        
        -- 更新圈子动态数
        UPDATE circles 
        SET post_count = (
            SELECT COUNT(*) 
            FROM posts 
            WHERE circle_id = COALESCE(NEW.circle_id, OLD.circle_id)
            AND status = 'active'
        )
        WHERE id = COALESCE(NEW.circle_id, OLD.circle_id);
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

CREATE TRIGGER trigger_update_comment_stats
AFTER INSERT OR UPDATE OR DELETE ON comments
FOR EACH ROW EXECUTE FUNCTION update_comment_stats();

CREATE TRIGGER trigger_update_comment_likes_stats
AFTER INSERT OR UPDATE OR DELETE ON comment_likes
FOR EACH ROW EXECUTE FUNCTION update_comment_stats();

CREATE TRIGGER trigger_update_circle_stats
AFTER INSERT OR UPDATE OR DELETE ON circle_members
FOR EACH ROW EXECUTE FUNCTION update_circle_stats();

CREATE TRIGGER trigger_update_circle_post_stats
AFTER INSERT OR UPDATE OR DELETE ON posts
FOR EACH ROW EXECUTE FUNCTION update_circle_stats(); 