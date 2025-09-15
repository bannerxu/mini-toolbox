#!/bin/bash

# Mini Toolbox Docker 部署脚本
# 作者: Mini Toolbox Team
# 版本: 1.0.0

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 配置变量
PROJECT_NAME="mini-toolbox"
COMPOSE_FILE="docker-compose.yml"

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

# 显示帮助信息
show_help() {
    echo "Mini Toolbox Docker 部署脚本"
    echo ""
    echo "用法: $0 [命令] [选项]"
    echo ""
    echo "命令:"
    echo "  start       启动应用"
    echo "  stop        停止所有服务"
    echo "  restart     重启所有服务"
    echo "  build       重新构建镜像"
    echo "  logs        查看日志"
    echo "  status      查看服务状态"
    echo "  clean       清理所有容器和镜像"
    echo "  backup      备份数据"
    echo "  restore     恢复数据"
    echo ""
    echo "选项:"
    echo "  -h, --help  显示此帮助信息"
    echo "  -v          详细输出"
}

# 检查Docker和Docker Compose
check_dependencies() {
    log_info "检查依赖..."
    
    if ! command -v docker &> /dev/null; then
        log_error "Docker 未安装，请先安装 Docker"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        log_error "Docker Compose 未安装，请先安装 Docker Compose"
        exit 1
    fi
    
    log_success "依赖检查通过"
}

# 部署应用
start_app() {
    log_info "启动应用..."
    
    check_dependencies
    
    # 停止可能运行的容器
    docker-compose -f $COMPOSE_FILE down 2>/dev/null || true
    
    # 构建并启动服务
    docker-compose -f $COMPOSE_FILE up --build -d
    
    # 等待服务启动
    log_info "等待服务启动..."
    sleep 10
    
    # 检查服务状态
    if docker-compose -f $COMPOSE_FILE ps | grep -q "Up"; then
        log_success "应用启动成功！"
        echo ""
        log_info "访问地址:"
        log_info "  前端: http://localhost"
        log_info "  后端: http://localhost:8080"
        log_info "  健康检查: http://localhost/health"
        echo ""
        log_info "查看日志: $0 logs"
        log_info "停止服务: $0 stop"
    else
        log_error "服务启动失败，请检查日志"
        docker-compose -f $COMPOSE_FILE logs
        exit 1
    fi
}


# 停止服务
stop_services() {
    log_info "停止服务..."
    
    docker-compose -f $COMPOSE_FILE down 2>/dev/null || true
    
    log_success "服务已停止"
}

# 重启服务
restart_services() {
    log_info "重启服务..."
    
    if docker-compose -f $COMPOSE_FILE ps -q | grep -q .; then
        docker-compose -f $COMPOSE_FILE restart
        log_success "服务已重启"
    else
        log_warning "没有运行中的服务"
    fi
}

# 重新构建镜像
rebuild_images() {
    log_info "重新构建镜像..."
    
    # 停止服务
    stop_services
    
    # 删除旧镜像
    docker images | grep $PROJECT_NAME | awk '{print $3}' | xargs -r docker rmi -f
    
    # 重新构建
    docker-compose -f $COMPOSE_FILE build --no-cache
    docker-compose -f $COMPOSE_FILE up -d
    
    log_success "镜像重建完成"
}

# 查看日志
show_logs() {
    if docker-compose -f $COMPOSE_FILE ps -q | grep -q .; then
        docker-compose -f $COMPOSE_FILE logs -f
    else
        log_warning "没有运行中的服务"
    fi
}

# 查看状态
show_status() {
    log_info "服务状态:"
    echo ""
    
    if docker-compose -f $COMPOSE_FILE ps -q | grep -q .; then
        docker-compose -f $COMPOSE_FILE ps
    else
        log_warning "没有运行中的服务"
    fi
    
    echo ""
    log_info "Docker 系统信息:"
    docker system df
}

# 清理容器和镜像
clean_all() {
    log_warning "这将删除所有 Mini Toolbox 相关的容器、镜像和卷"
    read -p "确定要继续吗? (y/N): " confirm
    
    if [[ $confirm =~ ^[Yy]$ ]]; then
        log_info "清理容器和镜像..."
        
        # 停止所有服务
        stop_services
        
        # 删除容器
        docker ps -a | grep $PROJECT_NAME | awk '{print $1}' | xargs -r docker rm -f
        
        # 删除镜像
        docker images | grep $PROJECT_NAME | awk '{print $3}' | xargs -r docker rmi -f
        
        # 删除卷
        docker volume ls | grep $PROJECT_NAME | awk '{print $2}' | xargs -r docker volume rm
        
        # 清理无用的资源
        docker system prune -f
        
        log_success "清理完成"
    else
        log_info "取消清理"
    fi
}

# 备份数据
backup_data() {
    log_info "备份数据..."
    
    backup_dir="backup/$(date +%Y%m%d_%H%M%S)"
    mkdir -p $backup_dir
    
    # 备份卷数据
    docker run --rm -v ${PROJECT_NAME}_uploads_data:/data -v $(pwd)/$backup_dir:/backup alpine tar czf /backup/uploads.tar.gz -C /data .
    docker run --rm -v ${PROJECT_NAME}_compressed_data:/data -v $(pwd)/$backup_dir:/backup alpine tar czf /backup/compressed.tar.gz -C /data .
    
    log_success "数据备份到: $backup_dir"
}

# 恢复数据
restore_data() {
    if [ -z "$2" ]; then
        log_error "请指定备份目录，例如: $0 restore backup/20231201_120000"
        exit 1
    fi
    
    backup_dir="$2"
    
    if [ ! -d "$backup_dir" ]; then
        log_error "备份目录不存在: $backup_dir"
        exit 1
    fi
    
    log_info "从 $backup_dir 恢复数据..."
    
    # 恢复卷数据
    if [ -f "$backup_dir/uploads.tar.gz" ]; then
        docker run --rm -v ${PROJECT_NAME}_uploads_data:/data -v $(pwd)/$backup_dir:/backup alpine tar xzf /backup/uploads.tar.gz -C /data
    fi
    
    if [ -f "$backup_dir/compressed.tar.gz" ]; then
        docker run --rm -v ${PROJECT_NAME}_compressed_data:/data -v $(pwd)/$backup_dir:/backup alpine tar xzf /backup/compressed.tar.gz -C /data
    fi
    
    log_success "数据恢复完成"
}

# 主函数
main() {
    case "${1:-}" in
        "start")
            start_app
            ;;
        "stop")
            stop_services
            ;;
        "restart")
            restart_services
            ;;
        "build")
            rebuild_images
            ;;
        "logs")
            show_logs
            ;;
        "status")
            show_status
            ;;
        "clean")
            clean_all
            ;;
        "backup")
            backup_data
            ;;
        "restore")
            restore_data "$@"
            ;;
        "-h"|"--help"|"")
            show_help
            ;;
        *)
            log_error "未知命令: $1"
            show_help
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@"