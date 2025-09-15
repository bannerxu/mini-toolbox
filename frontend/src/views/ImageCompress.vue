<template>
  <div class="image-compress-page">
    <!-- 头部导航 -->
    <header class="page-header">
      <div class="header-content">
        <button class="back-btn" @click="goBack">
          ← 返回首页
        </button>
        <h1 class="page-title">图片压缩工具</h1>
        <div class="header-actions">
          <ThemeSwitcher/>
        </div>
      </div>
    </header>

    <!-- 主要内容区域 -->
    <main class="page-main">
      <div class="compress-wrapper">
        <ImageCompress/>
      </div>
    </main>
  </div>
</template>

<script setup>
import {useRouter} from 'vue-router'
import ImageCompress from '../components/ImageCompress.vue'
import ThemeSwitcher from '../components/ThemeSwitcher.vue'

const router = useRouter()

const goBack = () => {
  router.push('/')
}
</script>

<style scoped>
.image-compress-page {
  min-height: 100vh;
  background: linear-gradient(135deg, var(--theme-primary, #667eea) 0%, var(--theme-secondary, #1C1D1F) 100%);
  transition: background 0.3s ease;
}

.page-header {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  padding: 1rem 2rem;
  position: sticky;
  top: 0;
  z-index: 100;
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.header-actions {
  flex-shrink: 0;
}

.back-btn {
  background: rgba(255, 255, 255, 0.2);
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: translateY(-1px);
}

.page-title {
  color: white;
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  flex: 1;
  text-align: center;
}

.page-main {
  padding: 2rem;
}

.compress-wrapper {
  max-width: 1200px;
  margin: 0 auto;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 2rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

@media (max-width: 768px) {
  .page-header {
    padding: 1rem;
  }

  .header-content {
    flex-direction: column;
    gap: 0.5rem;
  }

  .page-title {
    text-align: center;
  }

  .page-main {
    padding: 1rem;
  }

  .compress-wrapper {
    padding: 1rem;
    border-radius: 12px;
  }
}
</style>