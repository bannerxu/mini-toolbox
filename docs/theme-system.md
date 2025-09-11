# 主题系统使用指南

Mini Toolbox 提供了强大的主题切换系统，支持多种预设主题和系统跟随功能。

## 🎨 主题概览

### 可用主题

| 主题   | 图标 | 描述         | 特色         |
|------|----|------------|------------|
| 跟随系统 | 🌓 | 根据系统设置自动切换 | 智能适配，护眼友好  |
| 海洋蓝  | 🌊 | 清新的海洋蓝色主题  | 专业、清爽、默认主题 |
| 日落橙  | 🌅 | 温暖的日落橙色主题  | 活力、温暖、创意   |
| 森林绿  | 🌲 | 自然的森林绿色主题  | 自然、平静、环保   |
| 紫罗兰  | 💜 | 优雅的紫色主题    | 优雅、神秘、高贵   |
| 深色模式 | 🌙 | 护眼的深色主题    | 护眼、专业、夜间友好 |

## 🔧 如何使用主题系统

### 1. 切换主题

#### 方法一：使用主题切换器

1. 在页面右上角找到主题切换按钮
2. 点击按钮打开主题选择面板
3. 点击任意主题即可立即应用

#### 方法二：键盘快捷键

- 按 `Escape` 键关闭主题面板

### 2. 主题特性

#### 🌓 跟随系统主题

- **自动适配**: 根据操作系统设置自动切换明暗主题
- **实时更新**: 系统主题变化时自动同步
- **智能选择**: 浅色模式使用海洋蓝，深色模式使用深色主题
- **状态显示**: 面板中显示当前系统偏好

#### 🎯 自定义主题

- **即时预览**: 点击主题立即看到效果
- **颜色展示**: 每个主题显示主要颜色预览
- **持久保存**: 选择后自动保存，下次访问时恢复

### 3. 主题持久化

选择的主题会自动保存到浏览器本地存储中：

```javascript
// 主题设置保存在 localStorage
localStorage.getItem('mini-toolbox-theme') // 获取当前主题
```

#### 存储规则

- **键名**: `mini-toolbox-theme`
- **值**: 主题 ID（如 `ocean`、`sunset`、`auto` 等）
- **作用域**: 单个域名下有效
- **生命周期**: 永久保存，除非用户清理浏览器数据

## 🎨 主题设计规范

### 颜色系统

每个主题包含以下颜色变量：

```css
:root {
  --theme-primary: #667eea;    /* 主色调 */
  --theme-secondary: #764ba2;  /* 辅助色 */
  --theme-accent: #409eff;     /* 强调色 */
  --theme-success: #67c23a;    /* 成功色 */
  --theme-warning: #e6a23c;    /* 警告色 */
  --theme-error: #f56c6c;      /* 错误色 */
}
```

### 深色模式适配

深色模式使用特殊的 CSS 类：

```css
/* 组件在深色模式下的样式 */
:global(.dark-theme) .component {
  background: #2c3e50;
  color: white;
}
```

## 🛠️ 开发者指南

### 为组件添加主题支持

#### 1. 使用 CSS 变量

```vue
<style scoped>
.my-component {
  background: var(--theme-primary);
  color: white;
  border: 1px solid var(--theme-accent);
}
</style>
```

#### 2. 深色模式适配

```vue
<style scoped>
.my-component {
  background: white;
  color: #2c3e50;
}

/* 深色模式样式 */
:global(.dark-theme) .my-component {
  background: #2c3e50;
  color: white;
}
</style>
```

#### 3. 使用主题 Composable

```vue
<script setup>
import { useTheme } from '@/composables/useTheme.js'

const { currentTheme, currentThemeId, setTheme } = useTheme()

// 获取当前主题信息
console.log(currentTheme.value.name) // 当前主题名称
console.log(currentThemeId.value)    // 当前主题 ID

// 程序化切换主题
const switchToOcean = () => {
  setTheme('ocean')
}
</script>
```

### 添加新主题

#### 1. 定义主题配置

```javascript
// 在 useTheme.js 中添加
export const themes = {
  // ... 现有主题
  cosmic: {
    id: 'cosmic',
    name: '宇宙星空',
    icon: '🌌',
    description: '神秘的宇宙星空主题',
    colors: {
      primary: '#4a0e4e',
      secondary: '#81007f',
      accent: '#c732d8',
      success: '#4caf50',
      warning: '#ff9800',
      error: '#f44336'
    }
  }
}
```

#### 2. 更新组件样式

确保所有组件都支持新主题的颜色变量。

### 主题系统 API

#### useTheme() 返回值

| 属性/方法               | 类型          | 描述        |
|---------------------|-------------|-----------|
| `themes`            | Object      | 所有可用主题配置  |
| `currentThemeId`    | ComputedRef | 当前主题 ID   |
| `currentTheme`      | ComputedRef | 当前主题完整信息  |
| `systemPrefersDark` | ComputedRef | 系统是否偏好深色  |
| `setTheme(id)`      | Function    | 设置主题      |
| `loadTheme()`       | Function    | 从本地存储加载主题 |

#### 主题数据结构

```typescript
interface Theme {
  id: string                    // 主题唯一标识
  name: string                  // 主题显示名称
  icon: string                  // 主题图标 emoji
  description: string           // 主题描述
  colors?: {                    // 主题颜色（可选）
    primary: string
    secondary: string
    accent: string
    success: string
    warning: string
    error: string
  }
}
```

## 📱 移动端适配

### 响应式行为

- **桌面端**: 完整的主题选择面板
- **平板端**: 适中的面板尺寸
- **手机端**:
    - 主题按钮仅显示图标
    - 主题面板全屏显示
    - 触摸优化的交互

### 触摸体验

- **点击反馈**: 所有可点击元素有明确的触摸反馈
- **手势支持**: 支持滑动关闭面板（未来功能）
- **大按钮**: 符合移动端 44px 最小触摸目标

## 🔍 常见问题

### Q: 主题设置丢失了怎么办？

A: 检查浏览器是否禁用了本地存储，或者清理了浏览器数据。可以重新选择主题。

### Q: 系统跟随不工作？

A: 确保浏览器支持 `prefers-color-scheme` 媒体查询，且操作系统设置了明暗模式偏好。

### Q: 如何批量应用主题到组件？

A: 使用 CSS 变量和全局样式类，避免在每个组件中重复定义主题样式。

### Q: 主题切换有延迟？

A: 主题切换使用 CSS 变量和过渡动画，如果感觉延迟可以调整 `transition` 时间。

### Q: 深色模式下文字看不清？

A: 确保组件正确实现了深色模式样式，使用 `:global(.dark-theme)` 选择器。

## 🎯 最佳实践

### 1. 设计一致性

- 使用统一的颜色变量
- 保持视觉层次清晰
- 确保足够的对比度

### 2. 用户体验

- 提供清晰的视觉反馈
- 保持主题切换的流畅性
- 考虑不同用户的偏好

### 3. 性能优化

- 使用 CSS 变量而非动态样式
- 避免频繁的 DOM 操作
- 合理使用过渡动画

---

🎨 **主题系统让 Mini Toolbox 更加个性化，为每个用户提供最佳的视觉体验！**