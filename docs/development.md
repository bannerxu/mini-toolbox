# 开发规范

本文档定义了 Mini Toolbox 项目的开发规范和最佳实践。

## 📝 代码规范

### JavaScript/Vue 规范

#### 1. 命名规范

```javascript
// 变量和函数：小驼峰命名
const userName = 'john'
const getUserInfo = () => {}

// 常量：大写下划线分隔
const API_BASE_URL = 'https://api.example.com'
const MAX_FILE_SIZE = 5 * 1024 * 1024

// 组件名：大驼峰命名
// 文件名：PascalCase.vue
// components/ImageUploader.vue
// views/HomePage.vue
```

#### 2. Vue 组件规范

```vue
<!-- 推荐的组件结构 -->
<template>
  <div class="component-name">
    <!-- 模板内容 -->
  </div>
</template>

<script setup>
// 1. 导入声明
import { ref, computed, onMounted } from 'vue'
import SomeComponent from './SomeComponent.vue'

// 2. Props 定义
const props = defineProps({
  title: {
    type: String,
    required: true
  },
  count: {
    type: Number,
    default: 0
  }
})

// 3. Emits 定义
const emit = defineEmits(['update', 'change'])

// 4. 响应式数据
const isLoading = ref(false)
const userList = ref([])

// 5. 计算属性
const filteredUsers = computed(() => {
  return userList.value.filter(user => user.active)
})

// 6. 方法
const handleSubmit = () => {
  emit('update', formData)
}

// 7. 生命周期
onMounted(() => {
  fetchUserList()
})
</script>

<style scoped>
/* 组件样式 */
.component-name {
  /* 样式定义 */
}
</style>
```

#### 3. Composables 规范

```javascript
// composables/useCounter.js
import { ref, computed } from 'vue'

export function useCounter(initialValue = 0) {
  // 内部状态
  const count = ref(initialValue)
  
  // 计算属性
  const doubleCount = computed(() => count.value * 2)
  
  // 方法
  const increment = () => {
    count.value++
  }
  
  const decrement = () => {
    count.value--
  }
  
  const reset = () => {
    count.value = initialValue
  }
  
  // 返回公开的 API
  return {
    count: readonly(count),
    doubleCount,
    increment,
    decrement,
    reset
  }
}
```

### CSS 规范

#### 1. 选择器命名

```css
/* BEM 命名规范 */
.block {}
.block__element {}
.block--modifier {}
.block__element--modifier {}

/* 示例 */
.card {}
.card__header {}
.card__body {}
.card--large {}
.card__header--highlighted {}
```

#### 2. CSS 变量使用

```css
/* 主题变量 */
:root {
  --theme-primary: #667eea;
  --theme-secondary: #1C1D1F;
  --theme-text: #2c3e50;
}

/* 组件中使用 */
.component {
  background: var(--theme-primary);
  color: var(--theme-text);
  border: 1px solid var(--theme-secondary);
}
```

#### 3. 响应式设计

```css
/* 移动端优先 */
.component {
  /* 基础样式（移动端） */
  padding: 1rem;
}

/* 平板端 */
@media (min-width: 768px) {
  .component {
    padding: 1.5rem;
  }
}

/* 桌面端 */
@media (min-width: 1024px) {
  .component {
    padding: 2rem;
  }
}
```

## 📁 文件组织规范

### 目录结构

```
src/
├── components/          # 可复用组件
│   ├── common/         # 通用组件
│   ├── forms/          # 表单组件
│   └── layout/         # 布局组件
├── views/              # 页面组件
├── composables/        # 组合式函数
├── utils/              # 工具函数
├── stores/             # 状态管理（如需要）
├── router/             # 路由配置
├── assets/             # 静态资源
│   ├── images/
│   ├── icons/
│   └── styles/
└── types/              # TypeScript 类型定义（如需要）
```

### 文件命名

- **组件文件**: `PascalCase.vue`
- **页面文件**: `PascalCase.vue`
- **工具文件**: `camelCase.js`
- **样式文件**: `kebab-case.css`
- **图片文件**: `kebab-case.png/jpg/svg`

## 🔧 Git 规范

### 分支命名

```bash
# 功能分支
feature/user-authentication
feature/image-upload
feature/theme-system

# 修复分支
fix/upload-bug
fix/theme-switching-issue

# 热修复分支
hotfix/security-patch
hotfix/critical-bug

# 发布分支
release/v1.0.0
release/v1.1.0
```

### 提交信息规范

```bash
# 格式：<type>(<scope>): <description>

# 功能
feat(upload): add drag and drop support
feat(theme): implement dark mode

# 修复
fix(router): resolve navigation bug
fix(upload): handle large file error

# 文档
docs(readme): update installation guide
docs(api): add theme system documentation

# 样式
style(component): improve button hover effect
style(layout): adjust mobile spacing

# 重构
refactor(utils): simplify file validation logic
refactor(theme): extract color constants

# 测试
test(upload): add unit tests for file validation
test(theme): add integration tests

# 构建
build(deps): update vue to 3.4.0
build(config): optimize vite configuration

# 其他
chore(docs): update contributing guidelines
ci(actions): add automated testing
```

### 提交类型说明

- `feat`: 新功能
- `fix`: 错误修复
- `docs`: 文档更新
- `style`: 代码格式化（不影响代码运行）
- `refactor`: 重构（既不是新增功能，也不是修复错误）
- `test`: 测试相关
- `build`: 构建系统或外部依赖变更
- `ci`: CI 配置文件和脚本变更
- `chore`: 其他变更

## 📋 代码审查清单

### Pull Request 检查项

- [ ] 代码符合项目规范
- [ ] 包含适当的注释
- [ ] 添加或更新了测试
- [ ] 更新了相关文档
- [ ] 通过了所有测试
- [ ] 没有引入新的警告
- [ ] 性能没有明显下降
- [ ] 兼容目标浏览器

### Vue 组件检查项

- [ ] 使用 `<script setup>` 语法
- [ ] Props 有类型定义和默认值
- [ ] 事件使用 `defineEmits` 定义
- [ ] 样式使用 `scoped` 作用域
- [ ] 支持主题切换（使用 CSS 变量）
- [ ] 响应式设计适配
- [ ] 可访问性考虑（aria 标签等）

## 🧪 测试规范

### 单元测试

```javascript
// tests/components/Button.test.js
import { mount } from '@vue/test-utils'
import Button from '@/components/Button.vue'

describe('Button Component', () => {
  it('renders correctly', () => {
    const wrapper = mount(Button, {
      props: {
        label: 'Click me'
      }
    })
    
    expect(wrapper.text()).toContain('Click me')
  })
  
  it('emits click event', async () => {
    const wrapper = mount(Button)
    
    await wrapper.trigger('click')
    
    expect(wrapper.emitted()).toHaveProperty('click')
  })
})
```

### E2E 测试

```javascript
// tests/e2e/theme-switching.spec.js
describe('Theme Switching', () => {
  it('should switch theme correctly', () => {
    cy.visit('/')
    
    // 点击主题切换按钮
    cy.get('[data-testid="theme-toggle"]').click()
    
    // 选择深色主题
    cy.get('[data-testid="theme-dark"]').click()
    
    // 验证主题已切换
    cy.get('html').should('have.class', 'dark-theme')
  })
})
```

## 📊 性能规范

### 代码分割

```javascript
// 路由级别的代码分割
const routes = [
  {
    path: '/',
    component: () => import('@/views/Home.vue')
  },
  {
    path: '/upload',
    component: () => import('@/views/ImageUpload.vue')
  }
]
```

### 图片优化

- 使用适当的图片格式（WebP、AVIF）
- 提供多种尺寸的图片
- 实现懒加载
- 压缩图片文件

### 代码优化

- 避免不必要的重新渲染
- 使用 `computed` 而非 `watch`（当可以时）
- 合理使用 `v-memo` 和 `v-once`
- 优化大列表渲染

## 🔒 安全规范

### 输入验证

```javascript
// 文件上传验证
const validateFile = (file) => {
  // 检查文件类型
  const allowedTypes = ['image/jpeg', 'image/png', 'image/gif']
  if (!allowedTypes.includes(file.type)) {
    throw new Error('不支持的文件类型')
  }
  
  // 检查文件大小
  const maxSize = 5 * 1024 * 1024 // 5MB
  if (file.size > maxSize) {
    throw new Error('文件大小超过限制')
  }
  
  return true
}
```

### XSS 防护

- 使用 Vue 的内置转义
- 谨慎使用 `v-html`
- 验证和清理用户输入

### CSRF 防护

- 使用 CSRF 令牌
- 验证请求来源
- 实现适当的 CORS 策略

## 📚 文档规范

### 组件文档

```vue
<!--
@component ImageUploader
@description 支持拖拽的图片上传组件
@example
<ImageUploader 
  :max-size="5242880"
  @upload-success="handleSuccess"
  @upload-error="handleError"
/>
-->
<template>
  <!-- 组件模板 -->
</template>

<script setup>
/**
 * 图片上传组件的 Props
 */
const props = defineProps({
  /** 最大文件大小（字节） */
  maxSize: {
    type: Number,
    default: 5 * 1024 * 1024
  },
  /** 允许的文件类型 */
  acceptedTypes: {
    type: Array,
    default: () => ['image/jpeg', 'image/png', 'image/gif']
  }
})

/**
 * 组件事件
 */
const emit = defineEmits({
  /** 上传成功时触发 */
  'upload-success': (file, response) => true,
  /** 上传失败时触发 */
  'upload-error': (error) => true
})
</script>
```

### API 文档

- 使用 JSDoc 格式
- 描述参数类型和返回值
- 提供使用示例
- 注明注意事项

## 🎯 最佳实践

### 1. 性能优化

- 使用 `shallowRef` 和 `shallowReactive` 当适合时
- 避免在模板中使用复杂表达式
- 合理使用 `keep-alive`
- 实现虚拟滚动（大列表）

### 2. 用户体验

- 提供加载状态指示
- 实现错误边界
- 优化首屏加载时间
- 支持键盘导航

### 3. 可维护性

- 保持组件单一职责
- 使用 TypeScript（推荐）
- 编写测试用例
- 定期重构代码

### 4. 可访问性

- 使用语义化 HTML
- 提供 ARIA 标签
- 支持键盘导航
- 确保足够的颜色对比度

---

📖 **遵循这些规范有助于保持代码质量和项目的长期可维护性。如有疑问，请参考具体的技术文档或向团队咨询。**