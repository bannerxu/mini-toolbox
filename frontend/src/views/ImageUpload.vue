<template>
  <div class="image-upload-page">
    <!-- 头部导航 -->
    <header class="page-header">
      <div class="header-content">
        <button class="back-btn" @click="goBack">
          ← 返回首页
        </button>
        <h1 class="page-title">图片上传工具</h1>
        <div class="header-actions">
          <ThemeSwitcher/>
        </div>
      </div>
    </header>

    <!-- 主要内容区域 -->
    <main class="page-main">
      <div class="upload-wrapper">
        <ImageUploader/>
      </div>
    </main>
  </div>
</template>

<script setup>
import {useRouter} from 'vue-router'
import ImageUploader from '../components/ImageUploader.vue'
import ThemeSwitcher from '../components/ThemeSwitcher.vue'

const router = useRouter()

const goBack = () => {
  router.push('/')
}
</script>

<style scoped>
.image-upload-page {
  min-height: 100vh;
  background: linear-gradient(135deg, var(--theme-primary, #667eea) 0%, var(--theme-secondary, #764ba2) 100%);
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
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 0.9rem;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: translateX(-2px);
}

.page-title {
  color: white;
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}

.page-main {
  padding: 2rem;
}

.upload-wrapper {
  max-width: 800px;
  margin: 0 auto;
  background: rgba(255, 255, 255, 0.95);
  border-radius: 16px;
  padding: 2rem;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(10px);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .page-header {
    padding: 1rem;
  }

  .header-content {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }

  .page-title {
    font-size: 1.3rem;
  }

  .page-main {
    padding: 1rem;
  }

  .upload-wrapper {
    padding: 1.5rem;
  }
}

@media (max-width: 480px) {
  .back-btn {
    padding: 0.4rem 0.8rem;
    font-size: 0.8rem;
  }

  .page-title {
    font-size: 1.2rem;
  }

  .upload-wrapper {
    padding: 1rem;
  }
}

/* 深色主题适配 */
:global(.dark-theme) .image-upload-page {
  background: linear-gradient(135deg, #2c3e50 0%, #34495e 100%);
}
</style>