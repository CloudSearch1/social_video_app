<template>
  <div class="upload-container">
    <van-nav-bar
      title="上传视频"
      left-arrow
      @click-left="onClickLeft"
    />
    
    <div class="upload-content">
      <div 
        class="upload-area" 
        :class="{ 'has-video': videoFile }"
        @click="triggerFileInput"
      >
        <input 
          ref="fileInput"
          type="file"
          accept="video/*"
          class="file-input"
          @change="onFileSelected"
        />
        
        <template v-if="videoFile">
          <video 
            ref="videoPreview"
            class="video-preview"
            :src="videoPreviewUrl"
            muted
            loop
          ></video>
          <div class="video-actions">
            <van-button type="danger" size="small" @click.stop="removeVideo">删除</van-button>
          </div>
        </template>
        
        <template v-else>
          <van-icon name="plus" size="24" />
          <p>点击选择视频</p>
          <p class="upload-tip">支持mp4、mov等格式，文件大小不超过100MB</p>
        </template>
      </div>
      
      <div class="form-container">
        <van-field
          v-model="description"
          rows="3"
          autosize
          type="textarea"
          maxlength="150"
          placeholder="添加视频描述..."
          show-word-limit
        />
        
        <div class="topic-container">
          <div class="topic-header">
            <span>添加话题</span>
            <van-icon name="plus" @click="showTopicSelector = true" />
          </div>
          <div class="topic-list" v-if="selectedTopics.length > 0">
            <div 
              v-for="(topic, index) in selectedTopics" 
              :key="index"
              class="topic-tag"
            >
              #{{ topic }}
              <van-icon name="cross" @click.stop="removeTopic(index)" />
            </div>
          </div>
        </div>
        
        <div class="privacy-setting">
          <span>谁可以看这个视频</span>
          <van-dropdown-menu>
            <van-dropdown-item v-model="privacyValue" :options="privacyOptions" />
          </van-dropdown-menu>
        </div>
        
        <div class="allow-comment">
          <span>允许评论</span>
          <van-switch v-model="allowComment" size="24" />
        </div>
      </div>
      
      <div class="submit-container">
        <van-button 
          type="primary" 
          block 
          :disabled="!videoFile || !description.trim()" 
          :loading="uploading"
          @click="uploadVideo"
        >
          发布
        </van-button>
      </div>
    </div>
    
    <van-popup v-model:show="showTopicSelector" position="bottom" round>
      <div class="topic-selector">
        <div class="topic-selector-header">
          <span>选择话题</span>
          <van-icon name="cross" @click="showTopicSelector = false" />
        </div>
        
        <van-search
          v-model="topicSearchText"
          placeholder="搜索话题"
          shape="round"
        />
        
        <div class="topic-selector-list">
          <div 
            v-for="topic in filteredTopics" 
            :key="topic"
            class="topic-selector-item"
            @click="addTopic(topic)"
          >
            #{{ topic }}
          </div>
        </div>
      </div>
    </van-popup>
  </div>
</template>

<script>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { Toast } from 'vant'
import { useUserStore } from '../store/user'

export default {
  name: 'Upload',
  setup() {
    const router = useRouter()
    const userStore = useUserStore()
    
    const fileInput = ref(null)
    const videoPreview = ref(null)
    const videoFile = ref(null)
    const videoPreviewUrl = ref('')
    const description = ref('')
    const selectedTopics = ref([])
    const privacyValue = ref(0)
    const allowComment = ref(true)
    const uploading = ref(false)
    
    const showTopicSelector = ref(false)
    const topicSearchText = ref('')
    
    const privacyOptions = [
      { text: '公开', value: 0 },
      { text: '好友可见', value: 1 },
      { text: '仅自己可见', value: 2 }
    ]
    
    const popularTopics = [
      '抖音', '搞笑', '美食', '旅行', '音乐', '舞蹈', '宠物', '游戏',
      '穿搭', '美妆', '健身', '知识', '生活', '情感', '电影', '动漫'
    ]
    
    const filteredTopics = computed(() => {
      if (!topicSearchText.value) {
        return popularTopics
      }
      return popularTopics.filter(topic => 
        topic.toLowerCase().includes(topicSearchText.value.toLowerCase())
      )
    })
    
    // 触发文件选择
    const triggerFileInput = () => {
      if (!videoFile.value) {
        fileInput.value.click()
      }
    }
    
    // 选择文件
    const onFileSelected = (event) => {
      const file = event.target.files[0]
      if (!file) return
      
      // 检查文件类型
      if (!file.type.startsWith('video/')) {
        Toast('请选择视频文件')
        return
      }
      
      // 检查文件大小（100MB）
      if (file.size > 100 * 1024 * 1024) {
        Toast('文件大小不能超过100MB')
        return
      }
      
      videoFile.value = file
      videoPreviewUrl.value = URL.createObjectURL(file)
      
      // 自动播放预览
      setTimeout(() => {
        if (videoPreview.value) {
          videoPreview.value.play().catch(err => {
            console.error('视频预览播放失败:', err)
          })
        }
      }, 100)
    }
    
    // 删除视频
    const removeVideo = () => {
      if (videoPreviewUrl.value) {
        URL.revokeObjectURL(videoPreviewUrl.value)
      }
      
      videoFile.value = null
      videoPreviewUrl.value = ''
      fileInput.value.value = ''
    }
    
    // 添加话题
    const addTopic = (topic) => {
      if (!selectedTopics.value.includes(topic)) {
        selectedTopics.value.push(topic)
      }
      showTopicSelector.value = false
    }
    
    // 删除话题
    const removeTopic = (index) => {
      selectedTopics.value.splice(index, 1)
    }
    
    // 上传视频
    const uploadVideo = async () => {
      if (!userStore.isLoggedIn) {
        router.push('/login')
        return
      }
      
      if (!videoFile.value || !description.value.trim()) {
        Toast('请选择视频并添加描述')
        return
      }
      
      uploading.value = true
      
      try {
        // 模拟上传过程
        // 实际项目中应该调用后端API上传视频
        Toast.loading({
          message: '上传中...',
          forbidClick: true,
          duration: 0
        })
        
        // 模拟上传延迟
        await new Promise(resolve => setTimeout(resolve, 2000))
        
        Toast.clear()
        Toast.success('上传成功')
        
        // 上传成功后跳转到首页
        router.replace('/')
      } catch (error) {
        console.error('上传视频失败:', error)
        Toast.fail('上传失败，请稍后再试')
      } finally {
        uploading.value = false
      }
    }
    
    // 返回上一页
    const onClickLeft = () => {
      router.back()
    }
    
    onBeforeUnmount(() => {
      // 清理资源
      if (videoPreviewUrl.value) {
        URL.revokeObjectURL(videoPreviewUrl.value)
      }
    })
    
    return {
      fileInput,
      videoPreview,
      videoFile,
      videoPreviewUrl,
      description,
      selectedTopics,
      privacyValue,
      privacyOptions,
      allowComment,
      uploading,
      showTopicSelector,
      topicSearchText,
      filteredTopics,
      triggerFileInput,
      onFileSelected,
      removeVideo,
      addTopic,
      removeTopic,
      uploadVideo,
      onClickLeft
    }
  }
}
</script>

<style scoped>
.upload-container {
  min-height: 100vh;
  background-color: #f7f8fa;
  display: flex;
  flex-direction: column;
}

.upload-content {
  flex: 1;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.upload-area {
  width: 100%;
  height: 200px;
  background-color: #eee;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.upload-area.has-video {
  height: 300px;
}

.file-input {
  display: none;
}

.upload-tip {
  font-size: 12px;
  color: #999;
  margin-top: 10px;
}

.video-preview {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.video-actions {
  position: absolute;
  bottom: 10px;
  right: 10px;
}

.form-container {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.topic-container {
  background-color: #fff;
  padding: 10px;
  border-radius: 8px;
}

.topic-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.topic-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.topic-tag {
  background-color: #f0f0f0;
  padding: 5px 10px;
  border-radius: 16px;
  font-size: 12px;
  display: flex;
  align-items: center;
}

.topic-tag .van-icon {
  margin-left: 5px;
}

.privacy-setting, .allow-comment {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #fff;
  padding: 10px;
  border-radius: 8px;
}

.submit-container {
  margin-top: auto;
  padding: 10px 0;
}

.topic-selector {
  padding: 20px;
}

.topic-selector-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.topic-selector-list {
  margin-top: 15px;
  max-height: 300px;
  overflow-y: auto;
}

.topic-selector-item {
  padding: 10px;
  border-bottom: 1px solid #eee;
}
</style>