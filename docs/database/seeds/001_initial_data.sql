-- Master Guide 数据库种子数据
-- 版本: 001
-- 描述: 插入初始测试数据
-- 创建时间: 2024-12-01

-- 插入测试用户
INSERT INTO users (email, password_hash, phone, status, email_verified) VALUES
('admin@masterguide.com', crypt('admin123', gen_salt('bf')), '13800138000', 'active', true),
('mentor1@masterguide.com', crypt('mentor123', gen_salt('bf')), '13800138001', 'active', true),
('mentor2@masterguide.com', crypt('mentor123', gen_salt('bf')), '13800138002', 'active', true),
('student1@masterguide.com', crypt('student123', gen_salt('bf')), '13800138003', 'active', true),
('student2@masterguide.com', crypt('student123', gen_salt('bf')), '13800138004', 'active', true)
ON CONFLICT (email) DO NOTHING;

-- 插入测试身份
INSERT INTO user_identities (user_id, identity_type, domain, status, verification_status) VALUES
((SELECT id FROM users WHERE email = 'admin@masterguide.com'), 'master', 'software_development', 'active', 'verified'),
((SELECT id FROM users WHERE email = 'mentor1@masterguide.com'), 'master', 'software_development', 'active', 'verified'),
((SELECT id FROM users WHERE email = 'mentor1@masterguide.com'), 'apprentice', 'ui_ux_design', 'active', 'verified'),
((SELECT id FROM users WHERE email = 'mentor2@masterguide.com'), 'master', 'ui_ux_design', 'active', 'verified'),
((SELECT id FROM users WHERE email = 'student1@masterguide.com'), 'apprentice', 'software_development', 'active', 'verified'),
((SELECT id FROM users WHERE email = 'student2@masterguide.com'), 'apprentice', 'ui_ux_design', 'active', 'verified')
ON CONFLICT (user_id, domain, identity_type) DO NOTHING;

-- 插入测试档案
INSERT INTO user_profiles (user_id, identity_id, name, bio, skills, experience_years, hourly_rate, location) VALUES
((SELECT id FROM users WHERE email = 'admin@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'admin@masterguide.com') AND identity_type = 'master'), 
 '张大师', '资深全栈开发工程师，拥有8年开发经验', ARRAY['Go', 'Vue.js', 'PostgreSQL', 'Redis'], 8, 300.00, '北京'),
((SELECT id FROM users WHERE email = 'mentor1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor1@masterguide.com') AND identity_type = 'master'), 
 '李大师', 'Go语言专家，专注于后端开发', ARRAY['Go', 'Gin', 'PostgreSQL', 'Docker'], 6, 250.00, '上海'),
((SELECT id FROM users WHERE email = 'mentor1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor1@masterguide.com') AND identity_type = 'apprentice'), 
 '李同学', 'UI设计新手，正在学习Figma', ARRAY['Figma', 'Sketch'], 1, 50.00, '上海'),
((SELECT id FROM users WHERE email = 'mentor2@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor2@masterguide.com') AND identity_type = 'master'), 
 '王大师', 'UI/UX设计专家，擅长用户体验设计', ARRAY['Figma', 'Sketch', 'Adobe XD', '用户研究'], 7, 280.00, '深圳'),
((SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'student1@masterguide.com') AND identity_type = 'apprentice'), 
 '赵同学', '热爱编程的新手，正在学习Go语言', ARRAY['JavaScript', 'HTML', 'CSS'], 1, 30.00, '广州'),
((SELECT id FROM users WHERE email = 'student2@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'student2@masterguide.com') AND identity_type = 'apprentice'), 
 '孙同学', '设计爱好者，正在学习UI设计', ARRAY['Photoshop', 'Illustrator'], 1, 40.00, '杭州')
ON CONFLICT (identity_id) DO NOTHING;

-- 插入测试课程
INSERT INTO courses (mentor_id, title, description, price, duration_hours, difficulty, status, tags, prerequisites, learning_objectives) VALUES
((SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor1@masterguide.com') AND identity_type = 'master'), 
 'Go Web开发实战', '从零开始学习Go Web开发，掌握Gin框架和PostgreSQL数据库', 299.00, 20, 'intermediate', 'published', 
 ARRAY['Go', 'Web开发', 'Gin', 'PostgreSQL'], '需要基本的编程基础', 
 ARRAY['掌握Go语言基础语法', '学会使用Gin框架', '掌握PostgreSQL数据库操作', '能够独立开发Web应用']),
((SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor2@masterguide.com') AND identity_type = 'master'), 
 'UI设计从入门到精通', '系统学习UI设计理论和方法，掌握Figma等设计工具', 399.00, 25, 'beginner', 'published', 
 ARRAY['UI设计', 'Figma', '用户体验'], '无特殊要求', 
 ARRAY['理解UI设计基本原则', '掌握Figma设计工具', '学会用户研究方法', '能够独立完成UI设计项目']),
((SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'admin@masterguide.com') AND identity_type = 'master'), 
 '全栈开发实战', 'Vue.js + Go全栈开发完整项目实战', 499.00, 30, 'advanced', 'published', 
 ARRAY['Vue.js', 'Go', '全栈开发', '项目实战'], '需要Go和JavaScript基础', 
 ARRAY['掌握Vue.js前端开发', '学会Go后端开发', '理解前后端分离架构', '能够独立开发完整项目'])
ON CONFLICT DO NOTHING;

-- 插入测试课程内容
INSERT INTO course_contents (course_id, title, content_type, content_text, order_index, duration_minutes, is_free) VALUES
((SELECT id FROM courses WHERE title = 'Go Web开发实战'), '第一章：Go语言基础', 'text', 'Go语言简介和基础语法介绍', 1, 45, true),
((SELECT id FROM courses WHERE title = 'Go Web开发实战'), '第二章：Gin框架入门', 'video', 'Gin框架基础知识和路由配置', 2, 60, false),
((SELECT id FROM courses WHERE title = 'Go Web开发实战'), '第三章：数据库操作', 'video', 'PostgreSQL数据库连接和CRUD操作', 3, 75, false),
((SELECT id FROM courses WHERE title = 'UI设计从入门到精通'), '第一章：设计基础理论', 'text', 'UI设计的基本原则和设计思维', 1, 30, true),
((SELECT id FROM courses WHERE title = 'UI设计从入门到精通'), '第二章：Figma工具使用', 'video', 'Figma界面介绍和基本操作', 2, 90, false),
((SELECT id FROM courses WHERE title = 'UI设计从入门到精通'), '第三章：用户研究方法', 'video', '用户访谈、问卷调查等研究方法', 3, 60, false)
ON CONFLICT DO NOTHING;

-- 插入测试学习记录
INSERT INTO learning_records (user_id, course_id, progress_percentage, status, total_study_time) VALUES
((SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 (SELECT id FROM courses WHERE title = 'Go Web开发实战'), 65.5, 'learning', 780),
((SELECT id FROM users WHERE email = 'student2@masterguide.com'), 
 (SELECT id FROM courses WHERE title = 'UI设计从入门到精通'), 45.0, 'learning', 540)
ON CONFLICT (user_id, course_id) DO NOTHING;

-- 插入测试圈子
INSERT INTO circles (name, description, domain, created_by, status, tags) VALUES
('Go开发交流圈', 'Go语言开发技术交流，分享开发经验和最佳实践', 'software_development', 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor1@masterguide.com') AND identity_type = 'master'), 
 'active', ARRAY['Go', '开发', '技术交流']),
('UI设计学习圈', 'UI设计学习交流，分享设计作品和设计理念', 'ui_ux_design', 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor2@masterguide.com') AND identity_type = 'master'), 
 'active', ARRAY['UI设计', '用户体验', '设计交流'])
ON CONFLICT DO NOTHING;

-- 插入测试圈子成员
INSERT INTO circle_members (circle_id, user_id, identity_id, role) VALUES
((SELECT id FROM circles WHERE name = 'Go开发交流圈'), 
 (SELECT id FROM users WHERE email = 'mentor1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor1@masterguide.com') AND identity_type = 'master'), 
 'admin'),
((SELECT id FROM circles WHERE name = 'Go开发交流圈'), 
 (SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'student1@masterguide.com') AND identity_type = 'apprentice'), 
 'member'),
((SELECT id FROM circles WHERE name = 'UI设计学习圈'), 
 (SELECT id FROM users WHERE email = 'mentor2@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor2@masterguide.com') AND identity_type = 'master'), 
 'admin'),
((SELECT id FROM circles WHERE name = 'UI设计学习圈'), 
 (SELECT id FROM users WHERE email = 'student2@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'student2@masterguide.com') AND identity_type = 'apprentice'), 
 'member')
ON CONFLICT (circle_id, user_id) DO NOTHING;

-- 插入测试动态
INSERT INTO posts (user_id, identity_id, circle_id, content, post_type, status) VALUES
((SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'student1@masterguide.com') AND identity_type = 'apprentice'), 
 (SELECT id FROM circles WHERE name = 'Go开发交流圈'), 
 '今天学习了Go的并发编程，感觉收获很大！goroutine和channel的概念很清晰，希望以后能和大家多交流学习心得。', 'text', 'active'),
((SELECT id FROM users WHERE email = 'mentor1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor1@masterguide.com') AND identity_type = 'master'), 
 (SELECT id FROM circles WHERE name = 'Go开发交流圈'), 
 '分享一个Go项目的最佳实践：使用依赖注入和接口设计可以让代码更加清晰和易于测试。', 'text', 'active'),
((SELECT id FROM users WHERE email = 'student2@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'student2@masterguide.com') AND identity_type = 'apprentice'), 
 (SELECT id FROM circles WHERE name = 'UI设计学习圈'), 
 '刚完成了一个移动端UI设计项目，感觉配色和布局都很满意，想和大家分享一下设计思路。', 'text', 'active')
ON CONFLICT DO NOTHING;

-- 插入测试评论
INSERT INTO comments (post_id, user_id, identity_id, content) VALUES
((SELECT id FROM posts WHERE content LIKE '%Go的并发编程%'), 
 (SELECT id FROM users WHERE email = 'mentor1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor1@masterguide.com') AND identity_type = 'master'), 
 '很棒！Go的并发编程确实很强大，建议你可以尝试用channel实现生产者消费者模式。'),
((SELECT id FROM posts WHERE content LIKE '%Go项目的最佳实践%'), 
 (SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'student1@masterguide.com') AND identity_type = 'apprentice'), 
 '谢谢分享！这些经验对我们新手很有帮助。'),
((SELECT id FROM posts WHERE content LIKE '%移动端UI设计项目%'), 
 (SELECT id FROM users WHERE email = 'mentor2@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor2@masterguide.com') AND identity_type = 'master'), 
 '设计得很不错！建议可以考虑一下无障碍设计，让更多用户能够使用。')
ON CONFLICT DO NOTHING;

-- 插入测试预约
INSERT INTO appointments (student_id, mentor_id, appointment_time, duration_minutes, meeting_type, status, price, notes) VALUES
((SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor1@masterguide.com') AND identity_type = 'master'), 
 '2024-12-02 14:00:00', 60, 'video', 'confirmed', 250.00, '想请教Go并发编程的问题'),
((SELECT id FROM users WHERE email = 'student2@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor2@masterguide.com') AND identity_type = 'master'), 
 '2024-12-03 10:00:00', 90, 'video', 'pending', 280.00, '想请教UI设计配色方案')
ON CONFLICT DO NOTHING;

-- 插入测试评价
INSERT INTO reviews (reviewer_id, reviewed_id, course_id, rating, content, review_type) VALUES
((SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor1@masterguide.com') AND identity_type = 'master'), 
 (SELECT id FROM courses WHERE title = 'Go Web开发实战'), 5, '课程内容很实用，老师讲解很详细，收获很大！', 'course'),
((SELECT id FROM users WHERE email = 'student2@masterguide.com'), 
 (SELECT id FROM user_identities WHERE user_id = (SELECT id FROM users WHERE email = 'mentor2@masterguide.com') AND identity_type = 'master'), 
 (SELECT id FROM courses WHERE title = 'UI设计从入门到精通'), 4, '课程设计很系统，从理论到实践都有涉及，推荐！', 'course')
ON CONFLICT DO NOTHING;

-- 插入测试消息
INSERT INTO messages (from_user_id, to_user_id, content, message_type) VALUES
((SELECT id FROM users WHERE email = 'mentor1@masterguide.com'), 
 (SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 '你好！关于Go并发编程的问题，我们可以详细讨论一下。', 'text'),
((SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 (SELECT id FROM users WHERE email = 'mentor1@masterguide.com'), 
 '好的，谢谢老师！我有很多问题想请教。', 'text')
ON CONFLICT DO NOTHING;

-- 插入测试通知
INSERT INTO notifications (user_id, title, content, notification_type) VALUES
((SELECT id FROM users WHERE email = 'student1@masterguide.com'), 
 '课程更新通知', 'Go Web开发实战课程新增了第三章内容', 'course_update'),
((SELECT id FROM users WHERE email = 'student2@masterguide.com'), 
 '预约确认通知', '您的预约已确认，请按时参加', 'appointment_confirmed')
ON CONFLICT DO NOTHING;

-- 更新统计信息
UPDATE circles SET member_count = (SELECT COUNT(*) FROM circle_members WHERE circle_id = circles.id AND is_active = true);
UPDATE circles SET post_count = (SELECT COUNT(*) FROM posts WHERE circle_id = circles.id AND status = 'active');
UPDATE courses SET current_students = (SELECT COUNT(*) FROM learning_records WHERE course_id = courses.id);
UPDATE courses SET rating = (SELECT AVG(rating) FROM reviews WHERE course_id = courses.id AND review_type = 'course');
UPDATE courses SET review_count = (SELECT COUNT(*) FROM reviews WHERE course_id = courses.id AND review_type = 'course'); 