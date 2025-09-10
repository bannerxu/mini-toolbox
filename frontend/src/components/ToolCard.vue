<template>
  <div
      class="tool-card"
      :class="{ 'disabled': tool.disabled }"
      :style="{ '--primary-color': tool.color }"
      @click="handleClick"
  >
    <!-- 卡片图标 -->
    <div class="tool-icon">
      {{ tool.icon }}
    </div>

    <!-- 卡片内容 -->
    <div class="tool-content">
      <h3 class="tool-name">{{ tool.name }}</h3>
      <p class="tool-description">{{ tool.description }}</p>

      <!-- 分类标签 -->
      <div class="tool-category">
        <span class="category-tag">{{ tool.category }}</span>
      </div>
    </div>

    <!-- 状态指示器 -->
    <div class="tool-status" v-if="tool.disabled">
      <span class="status-text">开发中</span>
    </div>

    <!-- 悬停效果 -->
    <div class="hover-overlay">
      <span class="action-text">{{ tool.disabled ? '敬请期待' : '点击使用' }}</span>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  tool: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['click'])

const handleClick = () => {
  emit('click', props.tool)
}
</script>

<style scoped>
.tool-card {
  position: relative;
  background: white;
  border-radius: 16px;
  padding: 2rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  overflow: hidden;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

/* 深色主题适配 */
:global(.dark-theme) .tool-card {
  background: rgba(44, 62, 80, 0.9);
  border-color: rgba(255, 255, 255, 0.1);
  color: white;
}

:global(.dark-theme) .tool-name {
  color: white;
}

:global(.dark-theme) .tool-description {
  color: #bbb;
}

.tool-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.tool-card:active {
  transform: translateY(-4px);
}

.tool-card.disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.tool-card.disabled:hover {
  transform: translateY(-2px);
}

.tool-icon {
  font-size: 3rem;
  text-align: center;
  margin-bottom: 1rem;
  height: 4rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.tool-content {
  text-align: center;
}

.tool-name {
  font-size: 1.5rem;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 0.5rem;
  margin-top: 0;
}

.tool-description {
  color: #666;
  font-size: 0.95rem;
  line-height: 1.5;
  margin-bottom: 1rem;
}

.tool-category {
  display: flex;
  justify-content: center;
}

.category-tag {
  background: var(--primary-color, #409eff);
  color: white;
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.tool-status {
  position: absolute;
  top: 1rem;
  right: 1rem;
  background: #f56c6c;
  color: white;
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.7rem;
  font-weight: 500;
}

.hover-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, var(--primary-color, #409eff), rgba(0, 0, 0, 0.1));
  opacity: 0;
  transition: opacity 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 16px;
}

.tool-card:hover .hover-overlay {
  opacity: 0.9;
}

.action-text {
  color: white;
  font-size: 1.1rem;
  font-weight: 600;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}

/* 响应式设计 */
@media (max-width: 480px) {
  .tool-card {
    padding: 1.5rem;
  }

  .tool-icon {
    font-size: 2.5rem;
    height: 3rem;
  }

  .tool-name {
    font-size: 1.3rem;
  }

  .tool-description {
    font-size: 0.9rem;
  }
}
</style>