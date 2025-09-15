# Mini Toolbox Docker 部署指南

## 目录

- [快速开始](#快速开始)
- [环境要求](#环境要求)
- [部署方式](#部署方式)
- [配置说明](#配置说明)
- [常用命令](#常用命令)
- [故障排除](#故障排除)
- [性能优化](#性能优化)
- [安全建议](#安全建议)

## 快速开始

### 一键部署

```bash
# 克隆项目
git clone <your-repo-url>
cd mini-toolbox

# 启动应用
./deploy.sh start
```

访问地址：

- 前端应用：http://localhost
- 后端API：http://localhost:8080
- 健康检查：http://localhost/health

## 环境要求

### 系统要求

- **操作系统**：Linux、macOS 或 Windows (使用 WSL2)
- **内存**：至少 2GB RAM
- **存储**：至少 5GB 可用空间
- **网络**：需要访问互联网下载依赖

### 软件依赖

- **Docker**：版本 20.10 或更高
- **Docker Compose**：版本 2.0 或更高

### 安装 Docker

#### Ubuntu/Debian

```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER
```

#### CentOS/RHEL

```bash
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install docker-ce docker-ce-cli containerd.io
sudo systemctl enable --now docker
```

#### macOS

```bash
# 使用 Homebrew
brew install --cask docker

# 或下载 Docker Desktop
# https://www.docker.com/products/docker-desktop
```

## 部署方式

### Docker Compose 部署

使用 Docker Compose 部署应用：

```bash
./deploy.sh start
```

特点：

- 端口直接暴露：80 (前端)、8080 (后端)
- 实时日志输出
- 较少的资源限制
- 适合调试和开发

## 配置说明

### 环境变量

可以通过创建 `.env` 文件来自定义配置：

```bash
# .env 文件示例
GIN_MODE=release
TZ=Asia/Shanghai
BACKEND_PORT=8080
FRONTEND_PORT=80
```

### 卷挂载

- `uploads_data`：用户上传的文件
- `compressed_data`：压缩后的文件

### 网络配置

使用默认桥接网络

## 常用命令

### 部署脚本命令

```bash
# 查看帮助
./deploy.sh --help

# 启动应用
./deploy.sh start

# 停止所有服务
./deploy.sh stop

# 重启服务
./deploy.sh restart

# 重新构建镜像
./deploy.sh build

# 查看日志
./deploy.sh logs

# 查看服务状态
./deploy.sh status

# 清理所有容器和镜像
./deploy.sh clean

# 备份数据
./deploy.sh backup

# 恢复数据
./deploy.sh restore backup/20231201_120000
```

### Docker Compose 原生命令

```bash
# 启动服务
./deploy.sh start

# 或者直接使用 docker-compose
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down

# 重新构建
docker-compose build --no-cache
```

### 容器管理

```bash
# 进入后端容器
docker exec -it mini-toolbox-backend sh

# 进入前端容器
docker exec -it mini-toolbox-frontend sh

# 查看容器资源使用
docker stats

# 查看容器详细信息
docker inspect mini-toolbox-backend
```

## 故障排除

### 常见问题

#### 1. 端口被占用

```bash
# 查看端口占用
sudo lsof -i :80
sudo lsof -i :8080

# 或使用 netstat
netstat -tlnp | grep :80
```

解决方案：

- 停止占用端口的服务
- 修改 docker-compose.yml 中的端口映射

#### 2. 内存不足

症状：容器频繁重启，OOM 错误

解决方案：

```bash
# 增加系统内存
# 或修改 docker-compose.yml 中的资源限制
deploy:
  resources:
    limits:
      memory: 1G  # 增加内存限制
```

#### 3. 镜像构建失败

```bash
# 清理 Docker 缓存
docker system prune -a

# 重新构建
./deploy.sh build
```

#### 4. 服务无法访问

检查步骤：

1. 确认容器运行状态：`docker ps`
2. 检查网络连接：`docker network ls`
3. 查看服务日志：`./deploy.sh logs`
4. 测试健康检查：`curl http://localhost/health`

### 日志分析

```bash
# 查看所有服务日志
./deploy.sh logs

# 查看特定服务日志
docker-compose logs backend
docker-compose logs frontend

# 实时跟踪日志
docker-compose logs -f --tail=100

# 查看系统日志
journalctl -u docker
```

### 性能监控

```bash
# 查看容器资源使用
docker stats

# 查看系统资源
htop
iotop
```

## 性能优化

### 镜像优化

1. **多阶段构建**：已在 Dockerfile 中实现
2. **最小化基础镜像**：使用 Alpine Linux
3. **层级缓存**：合理安排 COPY 和 RUN 指令顺序

### 资源配置

```yaml
# docker-compose.yml 中的资源限制
deploy:
  resources:
    limits:
      cpus: '2.0'      # 根据实际需求调整
      memory: 1G       # 根据实际需求调整
    reservations:
      cpus: '1.0'
      memory: 512M
```

### 网络优化

1. **使用自定义网络**：减少网络延迟
2. **启用 HTTP/2**：在 nginx 配置中启用
3. **启用 Gzip**：已在 nginx.conf 中配置

## 安全建议

### 容器安全

1. **非 root 用户**：所有容器都使用非 root 用户运行
2. **只读文件系统**：对不需要写入的目录使用只读挂载
3. **最小权限原则**：只暴露必要的端口和卷

### 网络安全

```bash
# 设置防火墙规则（Ubuntu/Debian）
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw deny 8080/tcp  # 只允许内部访问

# 设置防火墙规则（CentOS/RHEL）
sudo firewall-cmd --permanent --add-port=80/tcp
sudo firewall-cmd --permanent --add-port=443/tcp
sudo firewall-cmd --reload
```

### 数据安全

1. **定期备份**：
   ```bash
   # 设置自动备份任务
   echo "0 2 * * * cd /path/to/mini-toolbox && ./deploy.sh backup" | crontab -
   ```

2. **加密存储**：使用加密的卷或文件系统

## 监控和维护

### 健康检查

所有服务都配置了健康检查：

```bash
# 手动检查服务健康状态
curl http://localhost/health
curl http://localhost:8080/health
```

### 日志管理

项目配置了日志轮转：

```yaml
logging:
  driver: "json-file"
  options:
    max-size: "100m"
    max-file: "3"
```

### 自动更新

设置自动更新脚本：

```bash
#!/bin/bash
# auto-update.sh
cd /path/to/mini-toolbox
git pull
./deploy.sh build
./deploy.sh restart
```

## 扩展配置

### 添加新服务

在 docker-compose.yml 中添加新服务：

```yaml
services:
  new-service:
    image: your-image
    container_name: mini-toolbox-new-service
    networks:
      - mini-toolbox-network
    depends_on:
      - backend
```

### 负载均衡

使用 nginx 配置负载均衡：

```nginx
upstream backend {
    server backend1:8080;
    server backend2:8080;
    server backend3:8080;
}
```

## 支持

如果遇到问题，请：

1. 查看 [故障排除](#故障排除) 部分
2. 检查 GitHub Issues
3. 提交新的 Issue，包含：
    - 系统信息
    - 错误日志
    - 复现步骤

---

**最后更新**: 2024年12月
**版本**: 1.0.0