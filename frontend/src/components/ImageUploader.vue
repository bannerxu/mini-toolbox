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
    <div v-if="uploadedImageData" class="result-container">
      <h3>上传成功!</h3>
      <div class="uploaded-info">
        <p><strong>文件名:</strong> {{ uploadedImageData.fileName }}</p>
        <p><strong>文件大小:</strong> {{ formatFileSize(uploadedImageData.fileSize) }}</p>
        <p><strong>上传时间:</strong> {{ formatTime(uploadedImageData.uploadTime) }}</p>
      </div>
      <img :src="uploadedImageUrl" alt="已上传图片" class="uploaded-image">
      <div class="action-buttons">
        <a :href="uploadedImageUrl" target="_blank" class="view-btn">查看图片</a>
        <button @click="goToCompress" class="compress-btn">去压缩</button>
      </div>
    </div>
  </div>
</template>

<script>
import {ref} from 'vue'
import axios from 'axios'
import {useRouter} from 'vue-router'

export default {
  name: 'ImageUploader',
  setup() {
    const router = useRouter()
    const fileInput = ref(null)
    const selectedFile = ref(null)
    const previewUrl = ref('')
    const isUploading = ref(false)
    const uploadStatus = ref('')
    const uploadedImageUrl = ref('')
    const uploadedImageData = ref(null)
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

        // 假设后端返回 { success: true, message: "上传成功", data: { filePath: "...", fileName: "...", fileSize: ... } }
        if (response.data.success) {
          uploadStatus.value = '上传成功!'
          statusClass.value = 'success'
          uploadedImageData.value = response.data.data
          uploadedImageUrl.value = `/api/uploads/${response.data.data.filePath}`
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

    // 格式化时间显示
    const formatTime = (timestamp) => {
      return new Date(timestamp * 1000).toLocaleString()
    }

    // 跳转到压缩页面
    const goToCompress = () => {
      router.push({
        name: 'ImageCompress',
        query: {
          filename: uploadedImageData.value.filePath
        }
      })
    }

    return {
      fileInput,
      selectedFile,
      previewUrl,
      isUploading,
      uploadStatus,
      uploadedImageUrl,
      uploadedImageData,
      statusClass,
      isDragOver,
      triggerFileInput,
      handleFileSelect,
      handleDrop,
      handleDragOver,
      handleDragEnter,
      handleDragLeave,
      uploadFile,
      formatFileSize,
      formatTime,
      goToCompress
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

.uploaded-info {
  margin: 15px 0;
  text-align: left;
  background: #f9f9f9;
  padding: 15px;
  border-radius: 4px;
  border-left: 4px solid #409eff;
}

.uploaded-info p {
  margin: 5px 0;
  color: #666;
}

.uploaded-image {
  max-width: 100%;
  max-height: 400px;
  border-radius: 4px;
  margin: 15px 0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.action-buttons {
  margin-top: 15px;
  display: flex;
  gap: 10px;
  justify-content: center;
}

.view-btn, .compress-btn {
  padding: 8px 16px;
  border-radius: 4px;
  text-decoration: none;
  border: none;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.3s;
}

.view-btn {
  background-color: #f0f0f0;
  color: #666;
}

.view-btn:hover {
  background-color: #e0e0e0;
}

.compress-btn {
  background-color: #52c41a;
  color: white;
}

.compress-btn:hover {
  background-color: #73d13d;
}
</style>