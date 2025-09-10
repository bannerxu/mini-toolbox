<template>
  <div class="theme-switcher">
    <!-- 主题切换按钮 -->
    <button
        class="theme-toggle-btn"
        @click="showThemePanel = !showThemePanel"
        :title="currentTheme.description"
    >
      <span class="theme-icon">{{ currentTheme.icon }}</span>
      <span class="theme-name">{{ currentTheme.name }}</span>
    </button>

    <!-- 主题选择面板 -->
    <Transition name="theme-panel">
      <div v-if="showThemePanel" class="theme-panel">
        <div class="theme-panel-header">
          <h3>选择主题</h3>
          <button class="close-btn" @click="showThemePanel = false">×</button>
        </div>

        <div class="theme-options">
          <div
              v-for="theme in themes"
              :key="theme.id"
              class="theme-option"
              :class="{ 'active': currentThemeId === theme.id }"
              @click="selectTheme(theme.id)"
          >
            <div class="theme-preview">
              <div class="theme-preview-icon">{{ theme.icon }}</div>
              <div
                  v-if="theme.colors"
                  class="theme-colors"
              >
                <div
                    class="color-dot"
                    v-for="(color, key) in theme.colors"
                    :key="key"
                    :style="{ backgroundColor: color }"
                    :title="key"
                ></div>
              </div>
              <div v-else class="auto-indicator">
                <div class="auto-split">
                  <div class="light-side"></div>
                  <div class="dark-side"></div>
                </div>
              </div>
            </div>

            <div class="theme-info">
              <div class="theme-title">{{ theme.name }}</div>
              <div class="theme-desc">{{ theme.description }}</div>
            </div>

            <div v-if="currentThemeId === theme.id" class="check-icon">✓</div>
          </div>
        </div>

        <!-- 系统信息 -->
        <div v-if="currentThemeId === 'auto'" class="system-info">
          <div class="system-status">
            <span class="status-icon">{{ systemPrefersDark ? '🌙' : '☀️' }}</span>
            <span>当前系统偏好: {{ systemPrefersDark ? '深色' : '浅色' }}</span>
          </div>
        </div>
      </div>
    </Transition>

    <!-- 背景遮罩 -->
    <Transition name="overlay">
      <div
          v-if="showThemePanel"
          class="theme-overlay"
          @click="showThemePanel = false"
      ></div>
    </Transition>
  </div>
</template>

<script setup>
import {onMounted, onUnmounted, ref} from 'vue'
import {useTheme} from '../composables/useTheme.js'

const {themes, currentThemeId, currentTheme, systemPrefersDark, setTheme} = useTheme()

const showThemePanel = ref(false)

const selectTheme = (themeId) => {
  setTheme(themeId)
  showThemePanel.value = false
}

// 点击外部关闭面板
const handleClickOutside = (event) => {
  if (showThemePanel.value && !event.target.closest('.theme-switcher')) {
    showThemePanel.value = false
  }
}

// ESC键关闭面板
const handleKeydown = (event) => {
  if (event.key === 'Escape' && showThemePanel.value) {
    showThemePanel.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  document.removeEventListener('keydown', handleKeydown)
})
</script>

<style scoped>
.theme-switcher {
  position: relative;
  z-index: 1000;
}

.theme-toggle-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
  font-size: 0.9rem;
}

.theme-toggle-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: translateY(-1px);
}

.theme-icon {
  font-size: 1.2rem;
}

.theme-name {
  font-weight: 500;
}

.theme-panel {
  position: absolute;
  top: calc(100% + 0.5rem);
  right: 0;
  background: white;
  border-radius: 16px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  min-width: 320px;
  max-width: 400px;
  z-index: 1001;
}

.theme-panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.5rem;
  border-bottom: 1px solid #f0f0f0;
}

.theme-panel-header h3 {
  margin: 0;
  color: #2c3e50;
  font-size: 1.1rem;
  font-weight: 600;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  color: #999;
  cursor: pointer;
  padding: 0;
  width: 2rem;
  height: 2rem;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: all 0.2s ease;
}

.close-btn:hover {
  background: #f5f5f5;
  color: #666;
}

.theme-options {
  padding: 1rem;
  max-height: 400px;
  overflow-y: auto;
}

.theme-option {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
  position: relative;
  border: 2px solid transparent;
}

.theme-option:hover {
  background: #f8f9fa;
}

.theme-option.active {
  background: var(--theme-primary, #667eea);
  color: white;
  border-color: var(--theme-primary, #667eea);
}

.theme-preview {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  flex-shrink: 0;
}

.theme-preview-icon {
  font-size: 1.5rem;
}

.theme-colors {
  display: flex;
  gap: 2px;
}

.color-dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.auto-indicator {
  width: 24px;
  height: 12px;
  border-radius: 6px;
  overflow: hidden;
  border: 1px solid #ddd;
}

.auto-split {
  display: flex;
  height: 100%;
}

.light-side {
  flex: 1;
  background: linear-gradient(135deg, #667eea, #764ba2);
}

.dark-side {
  flex: 1;
  background: linear-gradient(135deg, #2c3e50, #34495e);
}

.theme-info {
  flex: 1;
}

.theme-title {
  font-weight: 600;
  margin-bottom: 0.2rem;
}

.theme-desc {
  font-size: 0.8rem;
  opacity: 0.7;
}

.theme-option.active .theme-desc {
  opacity: 0.9;
}

.check-icon {
  font-size: 1.2rem;
  font-weight: bold;
  color: currentColor;
}

.system-info {
  padding: 1rem 1.5rem;
  border-top: 1px solid #f0f0f0;
  background: #f8f9fa;
  border-radius: 0 0 16px 16px;
}

.system-status {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  color: #666;
}

.status-icon {
  font-size: 1rem;
}

.theme-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.1);
  z-index: 999;
}

/* 动画效果 */
.theme-panel-enter-active,
.theme-panel-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.theme-panel-enter-from,
.theme-panel-leave-to {
  opacity: 0;
  transform: translateY(-10px) scale(0.95);
}

.overlay-enter-active,
.overlay-leave-active {
  transition: opacity 0.3s ease;
}

.overlay-enter-from,
.overlay-leave-to {
  opacity: 0;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .theme-panel {
    right: -1rem;
    left: -1rem;
    min-width: auto;
  }

  .theme-toggle-btn {
    padding: 0.4rem 0.8rem;
    font-size: 0.8rem;
  }

  .theme-name {
    display: none;
  }
}

/* 深色主题适配 */
:global(.dark-theme) .theme-panel {
  background: #2c3e50;
  border-color: rgba(255, 255, 255, 0.1);
}

:global(.dark-theme) .theme-panel-header {
  border-bottom-color: rgba(255, 255, 255, 0.1);
}

:global(.dark-theme) .theme-panel-header h3 {
  color: white;
}

:global(.dark-theme) .theme-option:hover {
  background: rgba(255, 255, 255, 0.1);
}

:global(.dark-theme) .system-info {
  background: rgba(255, 255, 255, 0.05);
  border-top-color: rgba(255, 255, 255, 0.1);
}

:global(.dark-theme) .system-status {
  color: #bbb;
}
</style>