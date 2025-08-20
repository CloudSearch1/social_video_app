<template>
  <div class="discover-container">
    <div class="search-header">
      <van-search
        v-model="searchText"
        placeholder="搜索用户、视频或话题"
        shape="round"
        background="#121212"
        @search="onSearch"
      />
      <div class="search-tabs">
        <van-tabs v-model:active="activeTab" animated swipeable>
          <van-tab title="推荐"></van-tab>
          <van-tab title="热门"></van-tab>
          <van-tab title="附近"></van-tab>
        </van-tabs>
      </div>
    </div>
    
    <div class="content-container">
      <!-- 热门话题 -->
      <div class="section-header">
        <h3>热门话题</h3>
        <span class="more-link">查看更多</span>
      </div>
      
      <div class="topic-list">
        <div 
          v-for="topic in hotTopics" 
          :key="topic.id" 
          class="topic-item"
          @click="goToTopic(topic.id)"
        >
          <div class="topic-image">
            <van-image width="100%" height="100%" fit="cover" :src="topic.coverUrl" />
          </div>
          <div class="topic-info">
            <div class="topic-name">#{{ topic.name }}</div>
            <div class="topic-count">{{ formatCount(topic.viewCount) }}次播放</div>
          </div>
        </div>
      </div>
      
      <!-- 推荐用户 -->
      <div class="section-header">
        <h3>推荐用户</h3>
        <span class="more-link">查看更多</span>
      </div>
      
      <div class="user-list">
        <div 
          v-for="user in recommendUsers" 
          :key="user.id" 
          class="user-item"
          @click="goToUserProfile(user.id)"
        >
          <div class="user-avatar">
            <van-image round width="100%" height="100%" fit="cover" :src="user.avatar" />
          </div>
          <div class="user-name">{{ user.username }}</div>
          <div class="user-follow">
            <van-button 
              size="mini" 
              type="primary" 
              :loading="user.followLoading"
              @click.stop="toggleFollow(user)"
            >
              {{ user.isFollowing ? '已关注' : '关注' }}
            </van-button>
          </div>
        </div>
      </div>
      
      <!-- 热门视频 -->
      <div class="section-header">
        <h3>热门视频</h3>
        <span class="more-link">查看更多</span>
      </div>
      
      <div class="video-grid">
        <div 
          v-for="video in hotVideos" 
          :key="video.id" 
          class="video-item"
          @click="goToVideo(video.id)"
        >
          <van-image width="100%" height="100%" fit="cover" :src="video.coverUrl" />
          <div class="video-info">
            <div class="video-desc">{{ video.description }}</div>
            <div class="video-user">
              <van-image round width="20" height="20" :src="video.user.avatar" />
              <span>{{ video.user.username }}</span>
            </div>
            <div class="video-stats">
              <van-icon name="play-circle-o" />
              <span>{{ formatCount(video.views) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Toast } from 'vant'
import { useUserStore } from '../store/user'

export default {
  name: 'Discover',
  setup() {
    const router = useRouter()
    const userStore = useUserStore()
    
    const searchText = ref('')
    const activeTab = ref(0)
    
    const hotTopics = ref([])
    const recommendUsers = ref([])
    const hotVideos = ref([])
    
    // 获取热门话题
    const fetchHotTopics = () => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      setTimeout(() => {
        hotTopics.value = Array(10).fill().map((_, i) => ({
          id: `topic-${i}`,
          name: `热门话题${i}`,
          coverUrl: 'https://example.com/topic-cover.jpg',
          viewCount: Math.floor(Math.random() * 100000000)
        }))
      }, 500)
    }
    
    // 获取推荐用户
    const fetchRecommendUsers = () => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      setTimeout(() => {
        recommendUsers.value = Array(10).fill().map((_, i) => ({
          id: `user-${i}`,
          username: `用户${i}`,
          avatar: 'https://example.com/avatar.jpg',
          isFollowing: Math.random() > 0.7,
          followLoading: false
        }))
      }, 500)
    }
    
    // 获取热门视频
    const fetchHotVideos = () => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      setTimeout(() => {
        hotVideos.value = Array(6).fill().map((_, i) => ({
          id: `video-${i}`,
          coverUrl: 'https://example.com/cover.jpg',
          description: `这是一个热门视频描述 #抖音 #热门 ${i}`,
          views: Math.floor(Math.random() * 10000000),
          user: {
            id: `user-${i}`,
            username: `用户${i}`,
            avatar: 'https://example.com/avatar.jpg'
          }
        }))
      }, 500)
    }
    
    // 格式化数字
    const formatCount = (count) => {
      if (count >= 100000000) {
        return (count / 100000000).toFixed(1) + '亿'
      } else if (count >= 10000) {
        return (count / 10000).toFixed(1) + '万'
      } else if (count >= 1000) {
        return (count / 1000).toFixed(1) + 'k'
      }
      return count
    }
    
    // 搜索
    const onSearch = () => {
      if (!searchText.value.trim()) return
      
      Toast('搜索功能开发中')
      // 实际项目中应该跳转到搜索结果页面
    }
    
    // 跳转到话题页面
    const goToTopic = (topicId) => {
      Toast(`话题页面开发中: ${topicId}`)
      // 实际项目中应该跳转到话题详情页面
    }
    
    // 跳转到用户主页
    const goToUserProfile = (userId) => {
      router.push(`/user/${userId}`)
    }
    
    // 跳转到视频详情
    const goToVideo = (videoId) => {
      router.push(`/video/${videoId}`)
    }
    
    // 关注/取消关注
    const toggleFollow = async (user) => {
      if (!userStore.isLoggedIn) {
        router.push('/login')
        return
      }
      
      user.followLoading = true
      
      try {
        let success
        if (user.isFollowing) {
          success = await userStore.unfollowUser(user.id)
        } else {
          success = await userStore.followUser(user.id)
        }
        
        if (success) {
          user.isFollowing = !user.isFollowing
        }
      } catch (error) {
        console.error('操作失败:', error)
        Toast('操作失败，请稍后再试')
      } finally {
        user.followLoading = false
      }
    }
    
    onMounted(() => {
      fetchHotTopics()
      fetchRecommendUsers()
      fetchHotVideos()
    })
    
    return {
      searchText,
      activeTab,
      hotTopics,
      recommendUsers,
      hotVideos,
      formatCount,
      onSearch,
      goToTopic,
      goToUserProfile,
      goToVideo,
      toggleFollow
    }
  }
}
</script>

<style scoped>
.discover-container {
  min-height: 100vh;
  background-color: #121212;
  padding-bottom: 50px;
}

.search-header {
  position: sticky;
  top: 0;
  z-index: 100;
  background-color: #121212;
}

.search-tabs {
  background-color: #121212;
}

.content-container {
  padding: 10px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 15px 0 10px;
}

.section-header h3 {
  font-size: 16px;
  color: #fff;
  margin: 0;
}

.more-link {
  font-size: 12px;
  color: #999;
}

.topic-list {
  display: flex;
  overflow-x: auto;
  padding: 5px 0;
  scrollbar-width: none; /* Firefox */
}

.topic-list::-webkit-scrollbar {
  display: none; /* Chrome, Safari, Edge */
}

.topic-item {
  flex: 0 0 120px;
  margin-right: 10px;
  border-radius: 8px;
  overflow: hidden;
  background-color: #1e1e1e;
}

.topic-image {
  height: 80px;
  overflow: hidden;
}

.topic-info {
  padding: 5px;
}

.topic-name {
  font-size: 14px;
  color: #fff;
  margin-bottom: 2px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.topic-count {
  font-size: 12px;
  color: #999;
}

.user-list {
  display: flex;
  overflow-x: auto;
  padding: 5px 0;
  scrollbar-width: none;
}

.user-list::-webkit-scrollbar {
  display: none;
}

.user-item {
  flex: 0 0 80px;
  margin-right: 15px;
  text-align: center;
}

.user-avatar {
  width: 60px;
  height: 60px;
  margin: 0 auto 5px;
}

.user-name {
  font-size: 12px;
  color: #fff;
  margin-bottom: 5px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.video-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 10px;
}

.video-item {
  border-radius: 8px;
  overflow: hidden;
  background-color: #1e1e1e;
  position: relative;
}

.video-item .van-image {
  aspect-ratio: 9/16;
}

.video-info {
  padding: 8px;
}

.video-desc {
  font-size: 12px;
  color: #fff;
  margin-bottom: 5px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}

.video-user {
  display: flex;
  align-items: center;
  font-size: 12px;
  color: #999;
  margin-bottom: 3px;
}

.video-user .van-image {
  margin-right: 5px;
}

.video-stats {
  font-size: 12px;
  color: #999;
  display: flex;
  align-items: center;
}

.video-stats .van-icon {
  margin-right: 3px;
}
</style>