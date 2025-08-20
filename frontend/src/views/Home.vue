<template>
  <div class="home-container">
    <van-swipe class="video-swipe" :loop="false" vertical :show-indicators="false" @change="onSwipeChange">
      <van-swipe-item v-for="(video, index) in videoList" :key="video.id">
        <video-player 
          :video="video" 
          :is-playing="currentIndex === index" 
          @like="handleLike" 
          @comment="handleComment"
          @share="handleShare"
          @follow="handleFollow"
        />
      </van-swipe-item>
    </van-swipe>
    
    <van-pull-refresh v-model="refreshing" @refresh="onRefresh" success-text="刷新成功">
      <div class="refresh-placeholder"></div>
    </van-pull-refresh>
  </div>
</template>

<script>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { Toast } from 'vant'
import VideoPlayer from '../components/VideoPlayer.vue'
import { useUserStore } from '../store/user'

export default {
  name: 'Home',
  components: {
    VideoPlayer
  },
  setup() {
    const router = useRouter()
    const userStore = useUserStore()
    
    const videoList = ref([])
    const currentIndex = ref(0)
    const refreshing = ref(false)
    const loading = ref(false)
    const finished = ref(false)
    const page = ref(1)
    
    // 获取视频列表
    const fetchVideos = async () => {
      try {
        loading.value = true
        // 模拟API请求
        // 实际项目中应该调用后端API
        setTimeout(() => {
          const newVideos = Array(5).fill().map((_, i) => ({
            id: Date.now() + i,
            url: 'https://example.com/video.mp4', // 示例URL，实际项目中应该使用真实视频URL
            coverUrl: 'https://example.com/cover.jpg',
            description: `这是一个示例视频描述 #抖音 #示例 ${page.value}-${i}`,
            likes: Math.floor(Math.random() * 10000),
            comments: Math.floor(Math.random() * 1000),
            shares: Math.floor(Math.random() * 500),
            isLiked: false,
            user: {
              id: 100 + i,
              username: `用户${100 + i}`,
              avatar: 'https://example.com/avatar.jpg',
              isFollowing: false
            }
          }))
          
          videoList.value = [...videoList.value, ...newVideos]
          page.value += 1
          loading.value = false
          
          if (page.value > 3) {
            finished.value = true
          }
        }, 1000)
      } catch (error) {
        console.error('获取视频失败:', error)
        Toast('获取视频失败，请稍后再试')
        loading.value = false
      }
    }
    
    // 下拉刷新
    const onRefresh = () => {
      videoList.value = []
      finished.value = false
      page.value = 1
      fetchVideos().then(() => {
        refreshing.value = false
        Toast('刷新成功')
      })
    }
    
    // 滑动切换视频
    const onSwipeChange = (index) => {
      currentIndex.value = index
      
      // 当滑动到倒数第二个视频时，加载更多
      if (index >= videoList.value.length - 2 && !loading.value && !finished.value) {
        fetchVideos()
      }
    }
    
    // 点赞
    const handleLike = (videoId) => {
      const video = videoList.value.find(v => v.id === videoId)
      if (!video) return
      
      if (userStore.isLoggedIn) {
        video.isLiked = !video.isLiked
        video.likes += video.isLiked ? 1 : -1
        // 实际项目中应该调用API更新点赞状态
      } else {
        router.push('/login')
      }
    }
    
    // 评论
    const handleComment = (videoId) => {
      router.push(`/video/${videoId}`)
    }
    
    // 分享
    const handleShare = (videoId) => {
      Toast('分享功能开发中')
    }
    
    // 关注
    const handleFollow = (userId) => {
      const videos = videoList.value.filter(v => v.user.id === userId)
      
      if (userStore.isLoggedIn) {
        videos.forEach(video => {
          video.user.isFollowing = !video.user.isFollowing
        })
        // 实际项目中应该调用API更新关注状态
      } else {
        router.push('/login')
      }
    }
    
    onMounted(() => {
      fetchVideos()
    })
    
    return {
      videoList,
      currentIndex,
      refreshing,
      loading,
      finished,
      onRefresh,
      onSwipeChange,
      handleLike,
      handleComment,
      handleShare,
      handleFollow
    }
  }
}
</script>

<style scoped>
.home-container {
  height: 100vh;
  width: 100%;
  position: relative;
  background-color: #000;
}

.video-swipe {
  height: 100%;
  width: 100%;
}

.refresh-placeholder {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 50px;
  z-index: 100;
  pointer-events: none;
}
</style>