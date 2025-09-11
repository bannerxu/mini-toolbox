# 常见问题解答 (FAQ)

本文档收录了 Mini Toolbox 使用过程中的常见问题和解决方案。

## 🚀 启动相关问题

### Q: npm install 失败怎么办？

**A:** 尝试以下解决方案：

```bash
# 1. 清理缓存
npm cache clean --force

# 2. 删除 node_modules 和 package-lock.json
rm -rf node_modules package-lock.json

# 3. 重新安装
npm install

# 4. 如果还是失败，尝试使用 yarn
npm install -g yarn
yarn install
```

### Q: 端口 3000 被占用怎么办？

**A:** Vite 会自动选择其他可用端口（如 3001、3002），或者你可以手动指定端口：

```bash
npm run dev -- --port 3001
```

### Q: 后端服务启动失败？

**A:** 检查以下几点：

1. 确保 Go 环境已正确安装：`go version`
2. 检查端口 8080 是否被占用
3. 查看错误日志确定具体原因

## 🎨 主题相关问题

### Q: 主题设置丢失了？

**A:** 可能的原因和解决方案：

- **浏览器隐私模式**: 隐私模式下 localStorage 不会持久保存
- **浏览器数据清理**: 重新选择喜欢的主题即可
- **浏览器兼容性**: 确保浏览器支持 localStorage

### Q: "跟随系统" 主题不工作？

**A:** 检查以下条件：

1. 浏览器支持 `prefers-color-scheme` 媒体查询
2. 操作系统已设置明暗模式偏好
3. 浏览器版本较新（推荐 Chrome 76+、Firefox 67+、Safari 12.1+）

### Q: 自定义主题色彩不正确？

**A:** 确保组件正确使用了 CSS 变量：

```css
/* 正确使用方式 */
.my-component {
  background: var(--theme-primary);
  color: var(--theme-text);
}
```

## 📱 响应式问题

### Q: 移动端显示异常？

**A:** 检查以下几点：

1. 浏览器缩放比例是否为 100%
2. 设备是否开启了特殊的显示模式
3. 尝试刷新页面或清除浏览器缓存

### Q: 主题切换面板在移动端显示不完整？

**A:** 这是正常现象，移动端的主题面板经过了优化设计：

- 主题按钮只显示图标，节省空间
- 主题面板会适配屏幕宽度
- 可以通过滑动查看所有主题选项

## 🖼️ 图片上传问题

### Q: 拖拽上传不工作？

**A:** 确认以下几点：

1. 浏览器支持 HTML5 拖拽 API
2. 文件是图片格式（jpg、png、gif、webp 等）
3. 文件大小不超过限制（默认 5MB）

### Q: 上传后显示错误？

**A:** 可能的原因：

- **后端未启动**: 确保 Go 后端服务运行在 8080 端口
- **网络问题**: 检查网络连接
- **文件格式**: 确认是支持的图片格式
- **文件大小**: 检查是否超过大小限制

### Q: 上传成功但无法预览？

**A:** 检查：

1. 后端静态文件服务是否正常
2. 图片文件是否成功保存
3. 浏览器控制台是否有错误信息

## 🔧 开发相关问题

### Q: 热重载不工作？

**A:** 尝试以下解决方案：

```bash
# 1. 重启开发服务器
npm run dev

# 2. 清理缓存
rm -rf node_modules/.vite

# 3. 检查文件监听限制（Linux/macOS）
echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
```

### Q: 如何添加新的工具？

**A:** 按照以下步骤：

1. 在 `src/components/` 创建工具组件
2. 在 `src/views/` 创建对应页面（如需要）
3. 在 `src/router/index.js` 添加路由
4. 在 `Home.vue` 的工具列表中添加配置
5. 更新相关文档

### Q: 如何自定义主题？

**A:** 参考 [主题系统开发指南](./theme-system.md#开发者指南)：

1. 在 `useTheme.js` 中添加主题配置
2. 定义主题颜色变量
3. 确保所有组件支持新主题

## 🌐 浏览器兼容性

### 支持的浏览器版本

| 浏览器     | 最低版本 | 推荐版本 |
|---------|------|------|
| Chrome  | 87+  | 最新版  |
| Firefox | 78+  | 最新版  |
| Safari  | 14+  | 最新版  |
| Edge    | 88+  | 最新版  |

### Q: IE 浏览器支持吗？

**A:** 不支持。Mini Toolbox 使用了现代 Web 技术，建议升级到现代浏览器。

### Q: 某些功能在旧浏览器中不工作？

**A:** 检查浏览器版本，主要功能需要：

- CSS Grid 支持
- CSS 自定义属性（变量）支持
- ES2020+ JavaScript 特性支持
- Fetch API 支持

## 📊 性能问题

### Q: 页面加载慢？

**A:** 优化建议：

1. 检查网络连接
2. 清除浏览器缓存
3. 确保没有运行过多的后台程序
4. 使用现代浏览器

### Q: 主题切换有延迟？

**A:** 这是正常现象：

- 主题切换包含 CSS 过渡动画（0.3s）
- 如果感觉延迟过长，检查浏览器性能

## 🐛 错误报告

### 如何报告 Bug？

1. 访问 [GitHub Issues](https://github.com/your-repo/mini-toolbox/issues)
2. 点击 "New Issue"
3. 选择 "Bug Report" 模板
4. 详细描述问题，包括：
    - 操作步骤
    - 预期结果
    - 实际结果
    - 浏览器和操作系统信息
    - 控制台错误信息（如有）

### Bug 报告模板

```
**问题描述**
简明扼要地描述问题。

**复现步骤**
1. 打开 '...'
2. 点击 '....'
3. 滚动到 '....'
4. 看到错误

**预期行为**
描述你期望发生的情况。

**实际行为**
描述实际发生的情况。

**截图**
如果适用，添加截图来帮助解释问题。

**环境信息**
- 操作系统: [例如 macOS 12.0]
- 浏览器: [例如 Chrome 95.0]
- 项目版本: [例如 v1.0.0]

**附加信息**
其他可能有助于解决问题的信息。
```

## 📞 获取帮助

如果这个 FAQ 没有解决你的问题：

1. 🔍 [搜索已有 Issues](https://github.com/your-repo/mini-toolbox/issues)
2. 💬 [参与社区讨论](https://github.com/your-repo/mini-toolbox/discussions)
3. 📝 [创建新的 Issue](https://github.com/your-repo/mini-toolbox/issues/new)
4. 📧 联系维护者（如果适用）

---

❓ **没找到你的问题？** 请在 [GitHub Discussions](https://github.com/your-repo/mini-toolbox/discussions)
中提出，我们会及时回复并更新此文档。