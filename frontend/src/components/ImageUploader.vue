<template>
  <div class="upload-container">
    <h2>图片上传</h2>

    <!-- 文件选择区域 -->
    <div
        class="upload-area"
        :class="{ 'drag-over': isDragOver }"
        @click="triggerFileInput"
        @drop="handleDrop"
        @dragover="handleDragOver"
        @dragenter="handleDragEnter"
        @dragleave="handleDragLeave"
    >
      <input
          type="file"
          ref="fileInput"
          class="file-input"
          accept="image/*"
          @change="handleFileSelect"
      >
      <div v-if="!selectedFile" class="upload-placeholder">
        <span>{{ isDragOver ? '松开鼠标以上传文件' : '点击选择图片或拖拽到此处' }}</span>
      </div>
      <div v-else class="file-info">
        <p>已选择: {{ selectedFile.name }}</p>
        <p>大小: {{ formatFileSize(selectedFile.size) }}</p>
      </div>
    </div>

    <!-- 图片预览 -->
    <div v-if="previewUrl" class="preview-container">
      <img :src="previewUrl" alt="预览图" class="preview-image">
    </div>

    <!-- 上传按钮和状态 -->
    <div class="action-area">
      <button
          @click="uploadFile"
          :disabled="!selectedFile || isUploading"
          class="upload-btn"
      >
        {{ isUploading ? '上传中...' : '开始上传' }}
      </button>

      <div v-if="uploadStatus" class="status-message" :class="statusClass">
        {{ uploadStatus }}
      </div>
    </div>

    <!-- 成功上传后的显示 -->
    <div v-if="uploadedImageUrl" class="result-container">
      <h3>上传成功!</h3>
      <img :src="uploadedImageUrl" alt="已上传图片" class="uploaded-image">
      <p class="image-link">
        <a :href="uploadedImageUrl" target="_blank">查看图片</a>
      </p>
    </div>
  </div>
</template>

<script>
import {ref} from 'vue'
import axios from 'axios'

export default {
  name: 'ImageUploader',
  setup() {
    const fileInput = ref(null)
    const selectedFile = ref(null)
    const previewUrl = ref('')
    const isUploading = ref(false)
    const uploadStatus = ref('')
    const uploadedImageUrl = ref('')
    const statusClass = ref('')
    const isDragOver = ref(false)

    // 触发文件选择
    const triggerFileInput = () => {
      fileInput.value?.click()
    }

    // 处理文件选择
    const handleFileSelect = (event) => {
      const file = event.target.files[0]
      processFile(file)
    }

    // 处理拖拽放置
    const handleDrop = (event) => {
      event.preventDefault()
      event.stopPropagation()
      isDragOver.value = false

      const files = event.dataTransfer.files
      if (files.length > 0) {
        processFile(files[0])
      }
    }

    // 处理拖拽悬停
    const handleDragOver = (event) => {
      event.preventDefault()
      event.stopPropagation()
    }

    // 处理拖拽进入
    const handleDragEnter = (event) => {
      event.preventDefault()
      event.stopPropagation()
      isDragOver.value = true
    }

    // 处理拖拽离开
    const handleDragLeave = (event) => {
      event.preventDefault()
      event.stopPropagation()
      // 只有当离开整个拖拽区域时才设置为false
      if (!event.currentTarget.contains(event.relatedTarget)) {
        isDragOver.value = false
      }
    }

    // 统一处理文件（选择或拖拽）
    const processFile = (file) => {
      if (!file) return

      // 验证文件类型
      if (!file.type.startsWith('image/')) {
        uploadStatus.value = '请选择图片文件!'
        statusClass.value = 'error'
        return
      }

      // 验证文件大小 (限制5MB)
      if (file.size > 5 * 1024 * 1024) {
        uploadStatus.value = '文件大小不能超过5MB!'
        statusClass.value = 'error'
        return
      }

      selectedFile.value = file
      uploadStatus.value = ''

      // 创建预览
      const reader = new FileReader()
      reader.onload = (e) => {
        previewUrl.value = e.target.result
      }
      reader.readAsDataURL(file)
    }

    // 上传文件到服务器
    const uploadFile = async () => {
      if (!selectedFile.value) return

      isUploading.value = true
      uploadStatus.value = '正在上传...'
      statusClass.value = 'info'

      try {
        const formData = new FormData()
        formData.append('image', selectedFile.value)  // 'image' 与后端保持一致

        const response = await axios.post('/api/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          },
          // 可选的上传进度监听
          onUploadProgress: (progressEvent) => {
            const percentCompleted = Math.round(
                (progressEvent.loaded * 100) / progressEvent.total
            )
            uploadStatus.value = `上传中: ${percentCompleted}%`
          }
        })

        // 假设后端返回 { success: true, message: "上传成功", data: { filePath: "..." } }
        if (response.data.success) {
          uploadStatus.value = '上传成功!'
          statusClass.value = 'success'
          uploadedImageUrl.value = `/api/static/${response.data.data.filePath}`
        } else {
          throw new Error(response.data.message || '上传失败')
        }
      } catch (error) {
        console.error('上传错误:', error)
        uploadStatus.value = error.response?.data?.message || error.message || '上传失败，请重试'
        statusClass.value = 'error'
      } finally {
        isUploading.value = false
      }
    }

    // 格式化文件大小显示
    const formatFileSize = (bytes) => {
      if (bytes === 0) return '0 B'
      const k = 1024
      const sizes = ['B', 'KB', 'MB', 'GB']
      const i = Math.floor(Math.log(bytes) / Math.log(k))
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
    }

    return {
      fileInput,
      selectedFile,
      previewUrl,
      isUploading,
      uploadStatus,
      uploadedImageUrl,
      statusClass,
      isDragOver,
      triggerFileInput,
      handleFileSelect,
      handleDrop,
      handleDragOver,
      handleDragEnter,
      handleDragLeave,
      uploadFile,
      formatFileSize
    }
  }
}
</script>

<style scoped>
.upload-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

.upload-area {
  border: 2px dashed #ccc;
  border-radius: 8px;
  padding: 40px 20px;
  text-align: center;
  cursor: pointer;
  transition: border-color 0.3s;
  margin-bottom: 20px;
}

.upload-area:hover {
  border-color: #409eff;
}

.upload-area.drag-over {
  border-color: #409eff;
  background-color: #f0f9ff;
  transform: scale(1.02);
}

.file-input {
  display: none;
}

.preview-container {
  margin: 20px 0;
  text-align: center;
}

.preview-image {
  max-width: 100%;
  max-height: 300px;
  border-radius: 4px;
}

.action-area {
  margin: 20px 0;
  text-align: center;
}

.upload-btn {
  padding: 12px 24px;
  background-color: #409eff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s;
}

.upload-btn:hover:not(:disabled) {
  background-color: #66b1ff;
}

.upload-btn:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.status-message {
  margin-top: 10px;
  padding: 8px;
  border-radius: 4px;
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

.result-container {
  margin-top: 30px;
  padding: 20px;
  border: 1px solid #e8e8e8;
  border-radius: 8px;
  text-align: center;
}

.uploaded-image {
  max-width: 100%;
  max-height: 400px;
  border-radius: 4px;
  margin: 15px 0;
}

.image-link {
  margin-top: 10px;
}
</style>