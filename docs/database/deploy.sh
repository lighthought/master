#!/bin/bash

# Master Guide 数据库部署脚本
# 版本: v1.0
# 创建时间: 2024-12-01

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 配置变量
DB_NAME="master_guide"
DB_USER="postgres"
DB_HOST="localhost"
DB_PORT="5432"
MIGRATIONS_DIR="migrations"
SEEDS_DIR="seeds"

# 日志函数
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查PostgreSQL是否安装
check_postgres() {
    log_info "检查PostgreSQL安装状态..."
    if ! command -v psql &> /dev/null; then
        log_error "PostgreSQL未安装，请先安装PostgreSQL"
        exit 1
    fi
    log_success "PostgreSQL已安装"
}

# 检查数据库连接
check_connection() {
    log_info "检查数据库连接..."
    if ! psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "SELECT 1;" &> /dev/null; then
        log_error "无法连接到PostgreSQL数据库"
        log_error "请检查以下配置："
        log_error "  - 数据库主机: $DB_HOST"
        log_error "  - 数据库端口: $DB_PORT"
        log_error "  - 数据库用户: $DB_USER"
        exit 1
    fi
    log_success "数据库连接正常"
}

# 创建数据库
create_database() {
    log_info "创建数据库: $DB_NAME"
    
    # 检查数据库是否已存在
    if psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -lqt | cut -d \| -f 1 | grep -qw $DB_NAME; then
        log_warning "数据库 $DB_NAME 已存在"
        read -p "是否要删除现有数据库并重新创建？(y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            log_info "删除现有数据库..."
            psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "DROP DATABASE IF EXISTS $DB_NAME;"
            log_success "数据库已删除"
        else
            log_info "跳过数据库创建"
            return 0
        fi
    fi
    
    # 创建数据库
    psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "CREATE DATABASE $DB_NAME;"
    log_success "数据库创建成功"
}

# 运行迁移文件
run_migrations() {
    log_info "运行数据库迁移..."
    
    # 创建迁移记录表
    psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
        CREATE TABLE IF NOT EXISTS migrations (
            id SERIAL PRIMARY KEY,
            filename VARCHAR(255) NOT NULL UNIQUE,
            executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    "
    
    # 获取所有迁移文件
    migration_files=($(ls -1 $MIGRATIONS_DIR/*.sql | sort))
    
    for file in "${migration_files[@]}"; do
        filename=$(basename "$file")
        
        # 检查是否已执行
        if psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c "SELECT 1 FROM migrations WHERE filename = '$filename';" | grep -q 1; then
            log_info "跳过已执行的迁移: $filename"
            continue
        fi
        
        log_info "执行迁移: $filename"
        
        # 执行迁移文件
        if psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$file"; then
            # 记录迁移执行
            psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "INSERT INTO migrations (filename) VALUES ('$filename');"
            log_success "迁移执行成功: $filename"
        else
            log_error "迁移执行失败: $filename"
            exit 1
        fi
    done
    
    log_success "所有迁移执行完成"
}

# 运行种子数据
run_seeds() {
    log_info "运行种子数据..."
    
    # 获取所有种子文件
    seed_files=($(ls -1 $SEEDS_DIR/*.sql | sort))
    
    for file in "${seed_files[@]}"; do
        filename=$(basename "$file")
        log_info "执行种子文件: $filename"
        
        if psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$file"; then
            log_success "种子数据执行成功: $filename"
        else
            log_error "种子数据执行失败: $filename"
            exit 1
        fi
    done
    
    log_success "所有种子数据执行完成"
}

# 创建应用用户
create_app_user() {
    log_info "创建应用用户..."
    
    # 创建只读用户
    psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
        DO \$\$
        BEGIN
            IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'readonly') THEN
                CREATE USER readonly WITH PASSWORD 'readonly_password';
            END IF;
        END
        \$\$;
    "
    
    # 创建应用用户
    psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
        DO \$\$
        BEGIN
            IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'app_user') THEN
                CREATE USER app_user WITH PASSWORD 'app_password';
            END IF;
        END
        \$\$;
    "
    
    # 授权
    psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
        GRANT CONNECT ON DATABASE $DB_NAME TO readonly, app_user;
        GRANT USAGE ON SCHEMA public TO readonly, app_user;
        GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly;
        GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO app_user;
        GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO app_user;
        ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO readonly;
        ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO app_user;
        ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO app_user;
    "
    
    log_success "应用用户创建完成"
}

# 验证部署
verify_deployment() {
    log_info "验证部署..."
    
    # 检查表是否存在
    tables=("users" "user_identities" "user_profiles" "courses" "circles" "posts")
    
    for table in "${tables[@]}"; do
        if psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c "SELECT 1 FROM information_schema.tables WHERE table_name = '$table';" | grep -q 1; then
            log_success "表 $table 存在"
        else
            log_error "表 $table 不存在"
            exit 1
        fi
    done
    
    # 检查数据是否插入
    user_count=$(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c "SELECT COUNT(*) FROM users;" | xargs)
    course_count=$(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c "SELECT COUNT(*) FROM courses;" | xargs)
    
    log_info "用户数量: $user_count"
    log_info "课程数量: $course_count"
    
    log_success "部署验证完成"
}

# 显示连接信息
show_connection_info() {
    log_info "数据库连接信息："
    echo "  数据库名称: $DB_NAME"
    echo "  数据库主机: $DB_HOST"
    echo "  数据库端口: $DB_PORT"
    echo "  管理员用户: $DB_USER"
    echo "  应用用户: app_user"
    echo "  只读用户: readonly"
    echo ""
    log_info "测试连接命令："
    echo "  psql -h $DB_HOST -p $DB_PORT -U app_user -d $DB_NAME"
    echo ""
    log_info "测试只读连接命令："
    echo "  psql -h $DB_HOST -p $DB_PORT -U readonly -d $DB_NAME"
}

# 主函数
main() {
    log_info "开始部署 Master Guide 数据库..."
    
    # 检查当前目录
    if [ ! -d "$MIGRATIONS_DIR" ]; then
        log_error "未找到迁移目录: $MIGRATIONS_DIR"
        log_error "请确保在正确的目录中运行此脚本"
        exit 1
    fi
    
    # 执行部署步骤
    check_postgres
    check_connection
    create_database
    run_migrations
    run_seeds
    create_app_user
    verify_deployment
    
    log_success "数据库部署完成！"
    echo ""
    show_connection_info
}

# 显示帮助信息
show_help() {
    echo "Master Guide 数据库部署脚本"
    echo ""
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  -h, --help     显示此帮助信息"
    echo "  -H, --host     数据库主机 (默认: localhost)"
    echo "  -p, --port     数据库端口 (默认: 5432)"
    echo "  -U, --user     数据库用户 (默认: postgres)"
    echo "  -d, --database 数据库名称 (默认: master_guide)"
    echo ""
    echo "示例:"
    echo "  $0"
    echo "  $0 -H 192.168.1.100 -p 5432 -U postgres -d master_guide"
}

# 解析命令行参数
while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            exit 0
            ;;
        -H|--host)
            DB_HOST="$2"
            shift 2
            ;;
        -p|--port)
            DB_PORT="$2"
            shift 2
            ;;
        -U|--user)
            DB_USER="$2"
            shift 2
            ;;
        -d|--database)
            DB_NAME="$2"
            shift 2
            ;;
        *)
            log_error "未知选项: $1"
            show_help
            exit 1
            ;;
    esac
done

# 运行主函数
main 