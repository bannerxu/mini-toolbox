<template>
  <div class="compress-container">
    <h2>图片压缩</h2>

    <!-- 文件选择区域 -->
    <div v-if="!selectedImage" class="file-selector">
      <div class="upload-area" @click="triggerFileInput">
        <input
            type="file"
            ref="fileInput"
            class="file-input"
            accept="image/*"
            @change="handleFileSelect"
        >
        <div class="upload-placeholder">
          <span>选择要压缩的图片</span>
        </div>
      </div>
    </div>

    <!-- 压缩配置区域 -->
    <div v-if="selectedImage" class="compress-config">
      <div class="image-preview">
        <h3>原始图片</h3>
        <img :src="originalImageUrl" alt="原始图片" class="preview-image">
        <div class="image-info">
          <p><strong>文件名:</strong> {{ selectedImage.fileName }}</p>
          <p><strong>文件大小:</strong> {{ formatFileSize(selectedImage.fileSize) }}</p>
        </div>
      </div>

      <div class="config-panel">
        <h3>压缩设置</h3>
        <div class="config-group">
          <label>压缩质量 (1-100):</label>
          <input
              type="range"
              v-model="compressOptions.quality"
              min="1"
              max="100"
              class="slider"
          >
          <span class="value-display">{{ compressOptions.quality }}%</span>
        </div>

        <div class="config-group">
          <label>
            <input
                type="checkbox"
                v-model="enableResize"
            > 调整尺寸
          </label>
        </div>

        <div v-if="enableResize" class="resize-options">
          <div class="size-input-group">
            <label>宽度:</label>
            <input
                type="number"
                v-model="compressOptions.width"
                @input="handleWidthChange"
                placeholder="自动"
                min="1"
            >
            <span>px</span>
          </div>
          <div class="size-input-group">
            <label>高度:</label>
            <input
                type="number"
                v-model="compressOptions.height"
                @input="handleHeightChange"
                placeholder="自动"
                min="1"
            >
            <span>px</span>
          </div>
          <div class="config-group">
            <label>
              <input
                  type="checkbox"
                  v-model="compressOptions.keepAspect"
              > 保持宽高比
            </label>
          </div>
          <div v-if="originalImageDimensions.width" class="original-size-info">
            <p><strong>原始尺寸:</strong> {{ originalImageDimensions.width }} x {{ originalImageDimensions.height }} px
            </p>
            <p><strong>宽高比:</strong> {{ aspectRatio.toFixed(2) }}</p>
          </div>
        </div>

        <button
            @click="compressImage"
            :disabled="isCompressing"
            class="compress-btn"
        >
          {{ isCompressing ? '压缩中...' : '开始压缩' }}
        </button>

        <div v-if="compressStatus" class="status-message" :class="statusClass">
          {{ compressStatus }}
        </div>
      </div>
    </div>

    <!-- 压缩结果对比 -->
    <div v-if="compressResult" class="result-comparison">
      <h3>压缩结果对比</h3>
      <div class="comparison-container">
        <div class="comparison-item">
          <h4>压缩前</h4>
          <img :src="originalImageUrl" alt="原始图片" class="comparison-image">
          <div class="comparison-info">
            <p><strong>大小:</strong> {{ formatFileSize(compressResult.originalSize) }}</p>
          </div>
        </div>

        <div class="comparison-arrow">
          <span>→</span>
        </div>

        <div class="comparison-item">
          <h4>压缩后</h4>
          <img :src="compressedImageUrl" alt="压缩后图片" class="comparison-image">
          <div class="comparison-info">
            <p><strong>大小:</strong> {{ formatFileSize(compressResult.compressedSize) }}</p>
            <p><strong>压缩比:</strong> {{ formatCompressionRatio(compressResult.compressionRatio) }}</p>
            <p><strong>节省:</strong> {{ formatFileSize(compressResult.originalSize - compressResult.compressedSize) }}
            </p>
          </div>
        </div>
      </div>

      <div class="result-actions">
        <a :href="compressedImageUrl" download class="download-btn">下载压缩图片</a>
        <button @click="resetCompress" class="reset-btn">重新压缩</button>
      </div>
    </div>
  </div>
</template>

<script>
import {computed, nextTick, onMounted, ref} from 'vue'
import {useRoute} from 'vue-router'
import axios from 'axios'

export default {
  name: 'ImageCompress',
  setup() {
    const route = useRoute()
    const fileInput = ref(null)
    const selectedImage = ref(null)
    const originalImageUrl = ref('')
    const compressedImageUrl = ref('')
    const compressResult = ref(null)
    const isCompressing = ref(false)
    const compressStatus = ref('')
    const statusClass = ref('')
    const enableResize = ref(false)
    const originalImageDimensions = ref({width: 0, height: 0})
    const isUpdatingDimensions = ref(false)

    const compressOptions = ref({
      quality: 85,
      width: '',
      height: '',
      keepAspect: true
    })

    // 计算宽高比
    const aspectRatio = computed(() => {
      if (originalImageDimensions.value.width && originalImageDimensions.value.height) {
        return originalImageDimensions.value.width / originalImageDimensions.value.height
      }
      return 1
    })

    // 初始化 - 检查是否有预设的文件名
    onMounted(() => {
      const filename = route.query.filename
      if (filename) {
        loadImageFromFilename(filename)
      }
    })

    // 从文件名加载图片
    const loadImageFromFilename = (filename) => {
      selectedImage.value = {
        fileName: filename,
        filePath: filename,
        fileSize: 0 // 这里可以通过API获取实际大小
      }
      originalImageUrl.value = `/api/uploads/${filename}`
      // 加载图片尺寸
      loadImageDimensions(`/api/uploads/${filename}`)
    }

    // 加载图片尺寸
    const loadImageDimensions = (imageUrl) => {
      const img = new Image()
      img.onload = () => {
        originalImageDimensions.value = {
          width: img.width,
          height: img.height
        }
      }
      img.src = imageUrl
    }

    // 触发文件选择
    const triggerFileInput = () => {
      fileInput.value?.click()
    }

    // 处理文件选择
    const handleFileSelect = (event) => {
      const file = event.target.files[0]
      if (!file) return

      // 验证文件类型
      if (!file.type.startsWith('image/')) {
        alert('请选择图片文件!')
        return
      }

      // 先上传文件
      uploadImage(file)
    }

    // 上传图片
    const uploadImage = async (file) => {
      try {
        const formData = new FormData()
        formData.append('image', file)

        const response = await axios.post('/api/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })

        if (response.data.success) {
          selectedImage.value = response.data.data
          originalImageUrl.value = `/api/uploads/${response.data.data.filePath}`
          // 加载图片尺寸
          await nextTick()
          loadImageDimensions(`/api/uploads/${response.data.data.filePath}`)
        } else {
          throw new Error(response.data.message || '上传失败')
        }
      } catch (error) {
        console.error('上传错误:', error)
        alert(error.response?.data?.message || error.message || '上传失败，请重试')
      }
    }

    // 处理宽度变化
    const handleWidthChange = () => {
      if (compressOptions.value.keepAspect && !isUpdatingDimensions.value && originalImageDimensions.value.width) {
        isUpdatingDimensions.value = true
        const width = parseInt(compressOptions.value.width)
        if (width > 0) {
          compressOptions.value.height = Math.round(width / aspectRatio.value)
        } else {
          compressOptions.value.height = ''
        }
        isUpdatingDimensions.value = false
      }
    }

    // 处理高度变化
    const handleHeightChange = () => {
      if (compressOptions.value.keepAspect && !isUpdatingDimensions.value && originalImageDimensions.value.height) {
        isUpdatingDimensions.value = true
        const height = parseInt(compressOptions.value.height)
        if (height > 0) {
          compressOptions.value.width = Math.round(height * aspectRatio.value)
        } else {
          compressOptions.value.width = ''
        }
        isUpdatingDimensions.value = false
      }
    }

    // 压缩图片
    const compressImage = async () => {
      if (!selectedImage.value) return

      isCompressing.value = true
      compressStatus.value = '正在压缩...'
      statusClass.value = 'info'

      try {
        const formData = new FormData()
        formData.append('filename', selectedImage.value.filePath)
        formData.append('quality', compressOptions.value.quality)

        if (enableResize.value) {
          if (compressOptions.value.width) {
            formData.append('width', compressOptions.value.width)
          }
          if (compressOptions.value.height) {
            formData.append('height', compressOptions.value.height)
          }
          formData.append('keepAspect', compressOptions.value.keepAspect)
        }

        const response = await axios.post('/api/compress', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })

        if (response.data.success) {
          compressStatus.value = '压缩成功!'
          statusClass.value = 'success'
          compressResult.value = response.data.data
          compressedImageUrl.value = response.data.data.compressedUrl
        } else {
          throw new Error(response.data.message || '压缩失败')
        }
      } catch (error) {
        console.error('压缩错误:', error)
        compressStatus.value = error.response?.data?.message || error.message || '压缩失败，请重试'
        statusClass.value = 'error'
      } finally {
        isCompressing.value = false
      }
    }

    // 重置压缩
    const resetCompress = () => {
      compressResult.value = null
      compressedImageUrl.value = ''
      compressStatus.value = ''
      compressOptions.value = {
        quality: 85,
        width: '',
        height: '',
        keepAspect: true
      }
      enableResize.value = false
      originalImageDimensions.value = {width: 0, height: 0}
    }

    // 格式化文件大小显示
    const formatFileSize = (bytes) => {
      if (bytes === 0) return '0 B'
      const k = 1024
      const sizes = ['B', 'KB', 'MB', 'GB']
      const i = Math.floor(Math.log(bytes) / Math.log(k))
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
    }

    // 格式化压缩比显示
    const formatCompressionRatio = (ratio) => {
      // 如果已经是字符串格式（带%号），直接返回
      if (typeof ratio === 'string' && ratio.includes('%')) {
        return ratio
      }
      // 如果是数字，进行计算
      if (typeof ratio === 'number') {
        return (ratio * 100).toFixed(1) + '%'
      }
      // 如果是字符串数字，先转换为数字
      const numRatio = parseFloat(ratio)
      if (!isNaN(numRatio)) {
        return numRatio.toFixed(1) + '%'
      }
      return '无法计算'
    }

    return {
      fileInput,
      selectedImage,
      originalImageUrl,
      compressedImageUrl,
      compressResult,
      isCompressing,
      compressStatus,
      statusClass,
      enableResize,
      compressOptions,
      originalImageDimensions,
      aspectRatio,
      triggerFileInput,
      handleFileSelect,
      handleWidthChange,
      handleHeightChange,
      compressImage,
      resetCompress,
      formatFileSize,
      formatCompressionRatio
    }
  }
}
</script>

<style scoped>
.compress-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.file-selector {
  margin-bottom: 30px;
}

.upload-area {
  border: 2px dashed #ccc;
  border-radius: 8px;
  padding: 40px 20px;
  text-align: center;
  cursor: pointer;
  transition: border-color 0.3s;
}

.upload-area:hover {
  border-color: #409eff;
}

.file-input {
  display: none;
}

.compress-config {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 30px;
  margin-bottom: 30px;
}

.image-preview {
  text-align: center;
}

.preview-image {
  max-width: 100%;
  max-height: 300px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.image-info {
  margin-top: 15px;
  text-align: left;
  background: #f9f9f9;
  padding: 15px;
  border-radius: 4px;
}

.config-panel {
  background: #f9f9f9;
  padding: 20px;
  border-radius: 8px;
}

.config-group {
  margin-bottom: 20px;
}

.config-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
}

.slider {
  width: 100%;
  margin-right: 10px;
}

.value-display {
  font-weight: bold;
  color: #409eff;
}

.resize-options {
  margin-left: 20px;
  padding: 15px;
  background: white;
  border-radius: 4px;
  border: 1px solid #e0e0e0;
}

.size-input-group {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}

.size-input-group input[type="number"] {
  width: 80px;
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.original-size-info {
  margin-top: 15px;
  padding: 10px;
  background: #f0f9ff;
  border-radius: 4px;
  border-left: 3px solid #409eff;
}

.original-size-info p {
  margin: 5px 0;
  font-size: 13px;
  color: #666;
}

.compress-btn {
  width: 100%;
  padding: 12px;
  background-color: #409eff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s;
}

.compress-btn:hover:not(:disabled) {
  background-color: #66b1ff;
}

.compress-btn:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.status-message {
  margin-top: 15px;
  padding: 10px;
  border-radius: 4px;
  text-align: center;
}

.status-message.info {
  background-color: #e6f7ff;
  color: #1890ff;
}

.status-message.success {
  background-color: #f6ffed;
  color: #52c41a;
}

.status-message.error {
  background-color: #fff2f0;
  color: #ff4d4f;
}

.result-comparison {
  margin-top: 40px;
}

.comparison-container {
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  gap: 20px;
  align-items: center;
  margin: 20px 0;
}

.comparison-item {
  text-align: center;
  padding: 20px;
  background: #f9f9f9;
  border-radius: 8px;
}

.comparison-arrow {
  font-size: 24px;
  color: #409eff;
  font-weight: bold;
}

.comparison-image {
  max-width: 100%;
  max-height: 250px;
  border-radius: 4px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.comparison-info {
  margin-top: 15px;
  text-align: left;
}

.comparison-info p {
  margin: 5px 0;
  color: #666;
}

.result-actions {
  display: flex;
  gap: 15px;
  justify-content: center;
  margin-top: 30px;
}

.download-btn, .reset-btn {
  padding: 12px 24px;
  border-radius: 4px;
  text-decoration: none;
  border: none;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.3s;
}

.download-btn {
  background-color: #52c41a;
  color: white;
}

.download-btn:hover {
  background-color: #73d13d;
}

.reset-btn {
  background-color: #f0f0f0;
  color: #666;
}

.reset-btn:hover {
  background-color: #e0e0e0;
}

@media (max-width: 768px) {
  .compress-config {
    grid-template-columns: 1fr;
  }

  .comparison-container {
    grid-template-columns: 1fr;
    gap: 15px;
  }

  .comparison-arrow {
    transform: rotate(90deg);
  }

  .result-actions {
    flex-direction: column;
    align-items: center;
  }
}
</style>