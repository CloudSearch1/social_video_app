<template>
  <div class="video-container">
    <video
      ref="videoRef"
      class="video-player"
      :src="video.url"
      :poster="video.coverUrl"
      loop
      muted
      playsinline
      webkit-playsinline
      x5-video-player-type="h5"
      x5-video-player-fullscreen="true"
      x5-video-orientation="portrait"
      @click="togglePlay"
    ></video>
    
    <div class="video-overlay">
      <div class="video-info">
        <div class="user-info">
          <div class="avatar" @click="goToUserProfile">
            <van-image round width="40" height="40" :src="video.user.avatar" />
            <div v-if="!video.user.isFollowing" class="follow-icon" @click.stop="handleFollow">+</div>
          </div>
          <div class="username" @click="goToUserProfile">{{ video.user.username }}</div>
        </div>
        <div class="video-desc">{{ video.description }}</div>
      </div>
      
      <div class="action-bar">
        <div class="action-item" @click="handleLike">
          <div class="action-icon" :class="{ active: video.isLiked }">
            <van-icon name="like" :color="video.isLiked ? '#fe2c55' : '#fff'" size="28" />
          </div>
          <div class="action-count">{{ formatCount(video.likes) }}</div>
        </div>
        <div class="action-item" @click="handleComment">
          <div class="action-icon">
            <van-icon name="chat" color="#fff" size="28" />
          </div>
          <div class="action-count">{{ formatCount(video.comments) }}</div>
        </div>
        <div class="action-item" @click="handleShare">
          <div class="action-icon">
            <van-icon name="share" color="#fff" size="28" />
          </div>
          <div class="action-count">{{ formatCount(video.shares) }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, watch, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'VideoPlayer',
  props: {
    video: {
      type: Object,
      required: true
    },
    isPlaying: {
      type: Boolean,
      default: false
    }
  },
  emits: ['like', 'comment', 'share', 'follow'],
  setup(props, { emit }) {
    const router = useRouter()
    const videoRef = ref(null)
    
    // 播放视频
    const playVideo = () => {
      if (videoRef.value) {
        videoRef.value.play().catch(err => {
          console.error('视频播放失败:', err)
        })
      }
    }
    
    // 暂停视频
    const pauseVideo = () => {
      if (videoRef.value) {
        videoRef.value.pause()
      }
    }
    
    // 监听isPlaying变化，控制视频播放/暂停
    watch(() => props.isPlaying, (newVal) => {
      if (newVal) {
        playVideo()
      } else {
        pauseVideo()
      }
    }, { immediate: true })
    
    // 切换播放/暂停
    const togglePlay = () => {
      if (videoRef.value) {
        if (videoRef.value.paused) {
          playVideo()
        } else {
          pauseVideo()
        }
      }
    }
    
    // 格式化数字
    const formatCount = (count) => {
      if (count >= 10000) {
        return (count / 10000).toFixed(1) + 'w'
      } else if (count >= 1000) {
        return (count / 1000).toFixed(1) + 'k'
      }
      return count
    }
    
    // 点赞
    const handleLike = () => {
      emit('like', props.video.id)
    }
    
    // 评论
    const handleComment = () => {
      emit('comment', props.video.id)
    }
    
    // 分享
    const handleShare = () => {
      emit('share', props.video.id)
    }
    
    // 关注
    const handleFollow = () => {
      emit('follow', props.video.user.id)
    }
    
    // 跳转到用户主页
    const goToUserProfile = () => {
      router.push(`/user/${props.video.user.id}`)
    }
    
    onMounted(() => {
      if (props.isPlaying) {
        playVideo()
      }
    })
    
    onBeforeUnmount(() => {
      pauseVideo()
    })
    
    return {
      videoRef,
      playVideo,
      pauseVideo,
      togglePlay,
      formatCount,
      handleLike,
      handleComment,
      handleShare,
      handleFollow,
      goToUserProfile
    }
  }
}
</script>

<style scoped>
.video-container {
  position: relative;
  height: 100%;
  width: 100%;
  overflow: hidden;
}

.video-player {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.video-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  background: linear-gradient(transparent, rgba(0, 0, 0, 0.5));
}

.video-info {
  flex: 1;
  padding-right: 80px;
}

.user-info {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.avatar {
  position: relative;
  margin-right: 10px;
}

.follow-icon {
  position: absolute;
  bottom: -2px;
  right: -2px;
  width: 20px;
  height: 20px;
  background-color: #fe2c55;
  color: #fff;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 16px;
  font-weight: bold;
}

.username {
  font-weight: bold;
  font-size: 16px;
}

.video-desc {
  font-size: 14px;
  margin-bottom: 10px;
  word-break: break-word;
}

.action-bar {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 15px;
}

.action-icon {
  margin-bottom: 5px;
}

.action-icon.active {
  animation: pulse 0.5s;
}

.action-count {
  font-size: 12px;
}

@keyframes pulse {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.2);
  }
  100% {
    transform: scale(1);
  }
}
</style>