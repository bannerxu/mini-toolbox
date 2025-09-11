# Mini Toolbox

> 🧰 开源的迷你工具箱 - 轻量级、便捷的工具集合

![Mini Toolbox](https://img.shields.io/badge/Mini_Toolbox-v1.0.0-blue.svg)
![Vue](https://img.shields.io/badge/Vue-3.4.0-green.svg)
![Vite](https://img.shields.io/badge/Vite-5.0.0-yellow.svg)
![License](https://img.shields.io/badge/License-MIT-red.svg)

Mini Toolbox 是一个现代化的 Web 工具集合，为开发人员和普通用户提供日常所需的实用工具。采用前后端分离架构，支持多主题切换，具有出色的响应式设计。

## ✨ 特性

- 🎨 **多主题支持** - 6种精美主题，支持系统跟随
- 📱 **响应式设计** - 完美适配桌面端、平板和移动端
- 🔧 **模块化架构** - 组件化设计，易于扩展
- ⚡ **现代化技术栈** - Vue 3 + Vite + Go
- 🚀 **快速启动** - 一键启动开发环境
- 💾 **本地存储** - 用户偏好设置自动保存

## 🛠️ 工具列表

### 已实现

- 🖼️ **图片上传** - 支持拖拽上传，多种格式支持

### 开发中

- 📝 文本格式化 - JSON、XML、HTML 格式化工具
- 📱 二维码生成 - 快速生成各种类型的二维码
- 🎨 颜色选择器 - RGB、HEX、HSL 颜色转换
- 🔗 链接缩短 - 生成短链接，方便分享
- 🔐 密码生成器 - 生成安全可靠的随机密码

## 🚀 快速开始

### 环境要求

- Node.js >= 16.0.0
- npm >= 8.0.0
- Go >= 1.19 (可选，用于后端功能)

### 启动前端

```bash
# 克隆项目
git clone <repository-url>
cd mini-toolbox

# 进入前端目录
cd frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

访问 `http://localhost:3000` 即可使用。

### 启动后端（可选）

```bash
# 在另一个终端中
cd backend
go run main.go
```

后端服务运行在 `http://localhost:8080`。

## 📚 文档

详细的项目文档位于 [`docs/`](./docs/) 目录：

### 📖 用户指南

- [📋 项目启动指南](./docs/getting-started.md) - 详细的安装和启动说明
- [🎨 主题系统使用指南](./docs/theme-system.md) - 主题切换功能说明
- [❓ 常见问题](./docs/faq.md) - 问题解答和故障排除

### 🛠️ 开发者文档

- [🏗️ 项目架构说明](./docs/architecture.md) - 系统架构和技术选型
- [📝 开发规范](./docs/development.md) - 代码规范和最佳实践
- [🧩 组件文档](./docs/components.md) - 组件使用说明

### 📋 其他文档

- [🚀 构建部署](./docs/deployment.md) - 项目构建和部署指南
- [📊 Wiki 首页](./docs/README.md) - 完整的文档索引

## 🎨 主题预览

| 主题      | 预览                                                              | 描述             |
|---------|-----------------------------------------------------------------|----------------|
| 🌓 跟随系统 | 自动切换                                                            | 根据系统设置自动切换明暗主题 |
| 🌊 海洋蓝  | ![#667eea](https://via.placeholder.com/15/667eea/000000?text=+) | 清新的海洋蓝色主题（默认）  |
| 🌅 日落橙  | ![#ff6b6b](https://via.placeholder.com/15/ff6b6b/000000?text=+) | 温暖的日落橙色主题      |
| 🌲 森林绿  | ![#4caf50](https://via.placeholder.com/15/4caf50/000000?text=+) | 自然的森林绿色主题      |
| 💜 紫罗兰  | ![#9c27b0](https://via.placeholder.com/15/9c27b0/000000?text=+) | 优雅的紫色主题        |
| 🌙 深色模式 | ![#bb86fc](https://via.placeholder.com/15/bb86fc/000000?text=+) | 护眼的深色主题        |

## 🏗️ 项目结构

```
mini-toolbox/
├── frontend/           # 前端项目 (Vue 3 + Vite)
│   ├── src/
│   │   ├── components/  # Vue 组件
│   │   ├── views/      # 页面组件
│   │   ├── composables/ # 组合式函数
│   │   ├── router/     # 路由配置
│   │   └── main.js     # 入口文件
│   ├── package.json    # 前端依赖
│   └── vite.config.js  # Vite 配置
├── backend/            # 后端项目 (Go)
│   └── main.go        # Go 主文件
├── docs/              # 项目文档
│   ├── README.md      # Wiki 首页
│   ├── getting-started.md
│   ├── architecture.md
│   └── theme-system.md
└── README.md          # 项目说明
```

## 🤝 贡献指南

我们欢迎各种形式的贡献！

### 如何贡献

1. Fork 这个项目
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的修改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开一个 Pull Request

### 贡献类型

- 🐛 Bug 修复
- ✨ 新功能
- 📝 文档改进
- 🎨 UI/UX 优化
- ⚡ 性能优化
- 🧪 测试用例

## 📊 开发状态

- [x] 项目基础架构
- [x] 首页和路由系统
- [x] 主题切换系统
- [x] 图片上传功能
- [x] 响应式设计
- [x] 项目文档
- [ ] 文本格式化工具
- [ ] 二维码生成工具
- [ ] 颜色选择器
- [ ] 链接缩短工具
- [ ] 密码生成器
- [ ] 单元测试
- [ ] 部署脚本

## 🔗 相关链接

- [在线预览](https://your-demo-url.com) - 在线体验（待部署）
- [问题反馈](https://github.com/your-repo/mini-toolbox/issues) - 报告 Bug 或建议
- [项目讨论](https://github.com/your-repo/mini-toolbox/discussions) - 社区讨论

## 📄 许可证

本项目采用 [MIT 许可证](LICENSE) - 查看 LICENSE 文件了解详情。

## 🙏 致谢

感谢所有为这个项目做出贡献的开发者！

特别感谢以下开源项目：

- [Vue.js](https://vuejs.org/) - 渐进式 JavaScript 框架
- [Vite](https://vitejs.dev/) - 下一代前端构建工具
- [Go](https://golang.org/) - 简洁、快速、可靠的编程语言

---

⭐ 如果这个项目对你有帮助，请给我们一个 Star！
