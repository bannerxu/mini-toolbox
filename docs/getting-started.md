# 项目启动指南

本文档详细说明如何启动 Mini Toolbox 项目。

## 📋 环境要求

### 必需工具

- **Node.js**: 版本 >= 16.0.0
- **npm**: 版本 >= 8.0.0（随 Node.js 一起安装）
- **Go**: 版本 >= 1.19（用于后端开发）

### 可选工具

- **Git**: 用于版本控制
- **VS Code**: 推荐的代码编辑器
- **Vue DevTools**: 浏览器扩展，用于 Vue 开发调试

## 🚀 快速启动

### 1. 克隆项目

```bash
git clone <repository-url>
cd mini-toolbox
```

### 2. 前端启动

#### 安装依赖

```bash
cd frontend
npm install
```

#### 启动开发服务器

```bash
npm run dev
```

启动成功后，你会看到类似输出：

```
VITE v5.4.20  ready in 887 ms
➜  Local:   http://localhost:3000/
➜  Network: http://192.168.120.6:3000/
```

在浏览器中访问 `http://localhost:3000` 即可看到应用界面。

#### 其他前端命令

```bash
# 构建生产版本
npm run build

# 预览构建结果
npm run preview
```

### 3. 后端启动（可选）

如果需要完整的图片上传功能，需要启动后端服务：

```bash
cd backend
go run main.go
```

后端服务默认运行在 `http://localhost:8080`。

## 🛠️ 开发环境配置

### Node.js 安装

1. 访问 [Node.js 官网](https://nodejs.org/)
2. 下载 LTS 版本
3. 按照安装向导完成安装
4. 验证安装：
   ```bash
   node --version
   npm --version
   ```

### Go 安装

1. 访问 [Go 官网](https://golang.org/dl/)
2. 下载适合您操作系统的版本
3. 按照安装向导完成安装
4. 验证安装：
   ```bash
   go version
   ```

## 📁 项目结构

```
mini-toolbox/
├── frontend/           # 前端项目
│   ├── src/
│   │   ├── components/  # Vue 组件
│   │   ├── views/      # 页面组件
│   │   ├── composables/ # 组合式函数
│   │   ├── router/     # 路由配置
│   │   └── main.js     # 入口文件
│   ├── package.json    # 依赖配置
│   ├── vite.config.js  # Vite 配置
│   └── index.html      # HTML 模板
├── backend/            # 后端项目
│   └── main.go        # Go 主文件
├── docs/              # 项目文档
└── README.md          # 项目说明
```

## 🔧 Vite 配置说明

前端使用 Vite 作为构建工具，主要配置包括：

- **开发服务器**: 端口 3000
- **代理配置**: `/api` 请求代理到后端 `http://localhost:8080`
- **Vue 插件**: 支持 Vue 3 单文件组件
- **热重载**: 开发时自动刷新

## 🌐 端口配置

| 服务      | 端口   | 用途         |
|---------|------|------------|
| 前端开发服务器 | 3000 | Vite 开发服务器 |
| 后端服务    | 8080 | Go API 服务器 |

## 🎯 功能验证

启动后可以验证以下功能：

### 1. 首页访问

- 访问 `http://localhost:3000`
- 应该看到 Mini Toolbox 首页
- 包含 6 个工具卡片

### 2. 主题切换

- 点击右上角主题切换按钮
- 尝试切换不同主题
- 刷新页面验证主题持久化

### 3. 图片上传（需要后端）

- 点击"图片上传"卡片
- 进入图片上传页面
- 测试拖拽或点击上传功能

### 4. 路由导航

- 在首页和工具页面间切换
- 验证返回按钮功能
- 检查浏览器历史记录

## ⚠️ 常见问题

### 端口被占用

如果 3000 端口被占用，Vite 会自动选择其他端口（如 3001）。

### 依赖安装失败

```bash
# 清理缓存重新安装
rm -rf node_modules package-lock.json
npm install
```

### 热重载不工作

检查文件监听限制：

```bash
# Linux/macOS
echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
```

### 代理请求失败

确保后端服务运行在 `http://localhost:8080`，或修改 `vite.config.js` 中的代理配置。

## 📞 获取帮助

如果遇到启动问题：

1. 检查环境要求是否满足
2. 查看控制台错误信息
3. 参考 [常见问题文档](./faq.md)
4. 在 GitHub 上[创建 Issue](https://github.com/your-repo/mini-toolbox/issues)

---

🎉 恭喜！现在你可以开始使用 Mini Toolbox 了！