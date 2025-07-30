# 数据库ID生成方式修改说明

## 修改概述

根据用户要求，我们将数据库中的所有表的 `id` 列从 `UUID` 类型和 `gen_random_uuid()` 默认值改为 `VARCHAR(32)` 类型并使用自定义的 `generate_table_id` 函数。

## 修改内容

### 1. 新增全局ID生成函数

在所有数据库脚本中添加了 `generate_table_id` 函数：

```sql
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
```

### 2. 为每个表创建对应的序列

为每个表创建了对应的序列，例如：
- `user_id_num_seq` - 用户表
- `course_id_num_seq` - 课程表
- `post_id_num_seq` - 动态表
- 等等...

### 3. 修改的表结构

所有表的 `id` 列都从：
```sql
id UUID PRIMARY KEY DEFAULT gen_random_uuid()
```

改为：
```sql
id VARCHAR(32) PRIMARY KEY DEFAULT generate_table_id('TABLE_PREFIX_', 'table_id_num_seq')
```

### 4. 修改的外键引用

所有外键引用都从 `UUID` 类型改为 `VARCHAR(32)` 类型。

## 修改的文件列表

### 主要文件
- `database/init.sql` - 主初始化脚本
- `database/migrations/001_create_initial_schema.sql` - 初始架构迁移
- `database/migrations/002_create_course_tables.sql` - 课程表迁移
- `database/migrations/003_create_community_tables.sql` - 社群表迁移
- `database/migrations/004_create_appointment_review_tables.sql` - 预约评价表迁移
- `database/migrations/005_create_message_notification_tables.sql` - 消息通知表迁移

### 种子数据文件
- `database/seeds/001_initial_data.sql` - 无需修改，使用子查询获取ID

## ID格式示例

修改后的ID格式为：`PREFIX_00000000001`

例如：
- 用户ID：`USER_00000000001`
- 课程ID：`COURSE_00000000001`
- 动态ID：`POST_00000000001`
- 评论ID：`COMMENT_00000000001`

## 优势

1. **可读性**：ID包含表前缀，便于识别和调试
2. **顺序性**：使用序列生成，保证ID的顺序性
3. **唯一性**：每个表使用独立的序列，避免ID冲突
4. **扩展性**：支持最多11位数字，足够应对大规模数据

## 注意事项

1. 所有外键引用都已相应更新为 `VARCHAR(32)` 类型
2. 索引和约束保持不变
3. 种子数据使用子查询，无需硬编码ID
4. 迁移脚本按顺序执行，确保依赖关系正确

## 部署说明

1. 如果是从零开始部署，直接运行 `database/deploy.sh` 脚本
2. 如果是现有数据库升级，需要按顺序执行迁移脚本
3. 确保在运行迁移脚本前，`generate_table_id` 函数已创建 