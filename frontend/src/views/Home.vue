<template>
  <div class="home-container">
    <!-- 头部区域 -->
    <header class="header">
      <div class="header-content">
        <div class="title-section">
          <h1 class="title">Mini Toolbox</h1>
          <p class="subtitle">迷你工具箱 - 轻量级、便捷的工具集合</p>
        </div>
        <div class="header-actions">
          <ThemeSwitcher/>
        </div>
      </div>
    </header>

    <!-- 工具网格 -->
    <main class="tools-grid">
      <ToolCard
          v-for="tool in tools"
          :key="tool.id"
          :tool="tool"
          @click="navigateToTool(tool)"
      />
    </main>

    <!-- 底部信息 -->
    <footer class="footer">
      <p>开源项目 | 持续更新中...</p>
    </footer>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {useRouter} from 'vue-router'
import ToolCard from '../components/ToolCard.vue'
import ThemeSwitcher from '../components/ThemeSwitcher.vue'
import {useTheme} from '../composables/useTheme.js'

const router = useRouter()
const {currentTheme} = useTheme()

// 工具列表配置
const tools = ref([
  {
    id: 'image-upload',
    name: '图片上传',
    description: '支持拖拽上传，多种格式支持',
    icon: '🖼️',
    route: '/image-upload',
    category: '文件处理',
    color: '#409eff'
  },
  {
    id: 'text-formatter',
    name: '文本格式化',
    description: 'JSON、XML、HTML格式化工具',
    icon: '📝',
    route: '/text-formatter',
    category: '文本处理',
    color: '#67c23a',
    disabled: true
  },
  {
    id: 'qr-generator',
    name: '二维码生成',
    description: '快速生成各种类型的二维码',
    icon: '📱',
    route: '/qr-generator',
    category: '实用工具',
    color: '#e6a23c',
    disabled: true
  },
  {
    id: 'color-picker',
    name: '颜色选择器',
    description: 'RGB、HEX、HSL颜色转换',
    icon: '🎨',
    route: '/color-picker',
    category: '设计工具',
    color: '#f56c6c',
    disabled: true
  },
  {
    id: 'url-shortener',
    name: '链接缩短',
    description: '生成短链接，方便分享',
    icon: '🔗',
    route: '/url-shortener',
    category: '网络工具',
    color: '#909399',
    disabled: true
  },
  {
    id: 'password-generator',
    name: '密码生成器',
    description: '生成安全可靠的随机密码',
    icon: '🔐',
    route: '/password-generator',
    category: '安全工具',
    color: '#606266',
    disabled: true
  }
])

// 导航到工具页面
const navigateToTool = (tool) => {
  if (tool.disabled) {
    // 显示开发中提示
    alert(`${tool.name} 功能正在开发中，敬请期待！`)
    return
  }
  router.push(tool.route)
}
</script>

<style scoped>
.home-container {
  min-height: 100vh;
  background: linear-gradient(135deg, var(--theme-primary, #667eea) 0%, var(--theme-secondary, #1C1D1F) 100%);
  padding: 2rem;
  transition: background 0.3s ease;
}

.header {
  margin-bottom: 3rem;
  color: white;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  max-width: 1200px;
  margin: 0 auto;
  gap: 2rem;
}

.title-section {
  text-align: center;
  flex: 1;
}

.header-actions {
  flex-shrink: 0;
  margin-top: 0.5rem;
}

.title {
  font-size: 3rem;
  font-weight: bold;
  margin-bottom: 0.5rem;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}

.subtitle {
  font-size: 1.2rem;
  opacity: 0.9;
  margin-bottom: 0;
}

.tools-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 2rem;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.footer {
  text-align: center;
  margin-top: 4rem;
  color: white;
  opacity: 0.8;
}

.footer p {
  margin: 0;
  font-size: 0.9rem;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .home-container {
    padding: 1rem;
  }

  .header-content {
    flex-direction: column;
    align-items: center;
    text-align: center;
    gap: 1rem;
  }

  .title {
    font-size: 2rem;
  }

  .subtitle {
    font-size: 1rem;
  }

  .tools-grid {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
}

@media (max-width: 480px) {
  .title {
    font-size: 1.8rem;
  }

  .tools-grid {
    padding: 0;
  }
}

/* 深色主题适配 */
:global(.dark-theme) .home-container {
  background: linear-gradient(135deg, #2c3e50 0%, #34495e 100%);
}

:global(.dark-theme) .footer {
  color: #bbb;
}
</style>