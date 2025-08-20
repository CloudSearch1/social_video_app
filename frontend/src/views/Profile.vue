<template>
  <div class="profile-container">
    <div class="user-header">
      <div class="user-info">
        <van-image
          round
          width="80"
          height="80"
          :src="userInfo.avatar || 'https://example.com/default-avatar.jpg'"
        />
        <h2 class="username">{{ userInfo.username }}</h2>
        <p class="user-id">抖音号：{{ userInfo.id }}</p>
        
        <div class="user-stats">
          <div class="stat-item">
            <div class="stat-value">{{ userInfo.following }}</div>
            <div class="stat-label">关注</div>
          </div>
          <div class="stat-item">
            <div class="stat-value">{{ userInfo.followers }}</div>
            <div class="stat-label">粉丝</div>
          </div>
          <div class="stat-item">
            <div class="stat-value">{{ userInfo.likes }}</div>
            <div class="stat-label">获赞</div>
          </div>
        </div>
        
        <div class="user-actions">
          <van-button v-if="isCurrentUser" type="primary" size="small" @click="editProfile">编辑资料</van-button>
          <van-button v-else type="primary" size="small" :loading="followLoading" @click="toggleFollow">
            {{ userInfo.isFollowing ? '已关注' : '关注' }}
          </van-button>
        </div>
        
        <div class="user-bio">{{ userInfo.bio || '这个人很懒，什么都没留下' }}</div>
      </div>
    </div>
    
    <van-tabs v-model:active="activeTab" animated swipeable>
      <van-tab title="作品">
        <div class="video-grid">
          <div v-if="videos.length === 0" class="empty-tip">
            暂无作品，快去发布你的第一个视频吧！
          </div>
          <div 
            v-for="video in videos" 
            :key="video.id" 
            class="video-item"
            @click="goToVideo(video.id)"
          >
            <van-image
              width="100%"
              height="100%"
              fit="cover"
              :src="video.coverUrl"
            />
            <div class="video-stats">
              <van-icon name="play-circle-o" />
              <span>{{ formatCount(video.views) }}</span>
            </div>
          </div>
        </div>
      </van-tab>
      
      <van-tab title="喜欢">
        <div class="video-grid">
          <div v-if="likedVideos.length === 0" class="empty-tip">
            暂无喜欢的视频
          </div>
          <div 
            v-for="video in likedVideos" 
            :key="video.id" 
            class="video-item"
            @click="goToVideo(video.id)"
          >
            <van-image
              width="100%"
              height="100%"
              fit="cover"
              :src="video.coverUrl"
            />
            <div class="video-stats">
              <van-icon name="play-circle-o" />
              <span>{{ formatCount(video.views) }}</span>
            </div>
          </div>
        </div>
      </van-tab>
    </van-tabs>
    
    <van-dialog
      v-model:show="showEditDialog"
      title="编辑资料"
      show-cancel-button
      @confirm="saveProfile"
    >
      <div class="edit-form">
        <van-field
          v-model="editForm.username"
          label="用户名"
          placeholder="请输入用户名"
        />
        <van-field
          v-model="editForm.bio"
          label="简介"
          type="textarea"
          placeholder="请输入个人简介"
          rows="3"
          autosize
        />
      </div>
    </van-dialog>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Toast } from 'vant'
import { useUserStore } from '../store/user'

export default {
  name: 'Profile',
  setup() {
    const route = useRoute()
    const router = useRouter()
    const userStore = useUserStore()
    
    const userId = computed(() => route.params.id || userStore.user?.id)
    const isCurrentUser = computed(() => !route.params.id || route.params.id === userStore.user?.id)
    
    const userInfo = ref({
      id: '',
      username: '',
      avatar: '',
      bio: '',
      following: 0,
      followers: 0,
      likes: 0,
      isFollowing: false
    })
    
    const videos = ref([])
    const likedVideos = ref([])
    const activeTab = ref(0)
    const followLoading = ref(false)
    const showEditDialog = ref(false)
    const editForm = ref({
      username: '',
      bio: ''
    })
    
    // 获取用户信息
    const fetchUserInfo = async () => {
      try {
        // 模拟API请求
        // 实际项目中应该调用后端API
        setTimeout(() => {
          if (isCurrentUser.value && userStore.user) {
            userInfo.value = {
              ...userStore.user,
              following: 42,
              followers: 128,
              likes: 1024
            }
          } else {
            userInfo.value = {
              id: userId.value,
              username: `用户${userId.value}`,
              avatar: 'https://example.com/avatar.jpg',
              bio: '这是一个示例简介',
              following: Math.floor(Math.random() * 500),
              followers: Math.floor(Math.random() * 1000),
              likes: Math.floor(Math.random() * 10000),
              isFollowing: Math.random() > 0.5
            }
          }
        }, 500)
      } catch (error) {
        console.error('获取用户信息失败:', error)
        Toast('获取用户信息失败，请稍后再试')
      }
    }
    
    // 获取用户视频
    const fetchUserVideos = async () => {
      try {
        // 模拟API请求
        // 实际项目中应该调用后端API
        setTimeout(() => {
          videos.value = Array(9).fill().map((_, i) => ({
            id: `video-${i}`,
            coverUrl: 'https://example.com/cover.jpg',
            views: Math.floor(Math.random() * 10000)
          }))
        }, 500)
      } catch (error) {
        console.error('获取用户视频失败:', error)
        Toast('获取用户视频失败，请稍后再试')
      }
    }
    
    // 获取用户喜欢的视频
    const fetchLikedVideos = async () => {
      try {
        // 模拟API请求
        // 实际项目中应该调用后端API
        setTimeout(() => {
          likedVideos.value = Array(6).fill().map((_, i) => ({
            id: `liked-video-${i}`,
            coverUrl: 'https://example.com/cover.jpg',
            views: Math.floor(Math.random() * 10000)
          }))
        }, 500)
      } catch (error) {
        console.error('获取喜欢的视频失败:', error)
        Toast('获取喜欢的视频失败，请稍后再试')
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
    
    // 跳转到视频详情
    const goToVideo = (videoId) => {
      router.push(`/video/${videoId}`)
    }
    
    // 关注/取消关注
    const toggleFollow = async () => {
      if (!userStore.isLoggedIn) {
        router.push('/login')
        return
      }
      
      followLoading.value = true
      
      try {
        let success
        if (userInfo.value.isFollowing) {
          success = await userStore.unfollowUser(userId.value)
        } else {
          success = await userStore.followUser(userId.value)
        }
        
        if (success) {
          userInfo.value.isFollowing = !userInfo.value.isFollowing
          userInfo.value.followers += userInfo.value.isFollowing ? 1 : -1
        }
      } catch (error) {
        console.error('操作失败:', error)
        Toast('操作失败，请稍后再试')
      } finally {
        followLoading.value = false
      }
    }
    
    // 编辑资料
    const editProfile = () => {
      editForm.value.username = userInfo.value.username
      editForm.value.bio = userInfo.value.bio
      showEditDialog.value = true
    }
    
    // 保存资料
    const saveProfile = async () => {
      try {
        const success = await userStore.updateProfile({
          username: editForm.value.username,
          bio: editForm.value.bio
        })
        
        if (success) {
          userInfo.value.username = editForm.value.username
          userInfo.value.bio = editForm.value.bio
          Toast.success('资料更新成功')
        } else {
          Toast.fail(userStore.error || '更新失败，请稍后再试')
        }
      } catch (error) {
        console.error('更新资料失败:', error)
        Toast.fail('更新失败，请稍后再试')
      }
    }
    
    onMounted(() => {
      fetchUserInfo()
      fetchUserVideos()
      fetchLikedVideos()
    })
    
    return {
      userInfo,
      videos,
      likedVideos,
      activeTab,
      followLoading,
      showEditDialog,
      editForm,
      isCurrentUser,
      formatCount,
      goToVideo,
      toggleFollow,
      editProfile,
      saveProfile
    }
  }
}
</script>

<style scoped>
.profile-container {
  min-height: 100vh;
  background-color: #f7f8fa;
  padding-bottom: 50px;
}

.user-header {
  background-color: #fff;
  padding: 20px;
  text-align: center;
}

.username {
  margin: 10px 0 5px;
  font-size: 18px;
}

.user-id {
  color: #999;
  font-size: 14px;
  margin: 0 0 15px;
}

.user-stats {
  display: flex;
  justify-content: center;
  margin-bottom: 15px;
}

.stat-item {
  padding: 0 20px;
  position: relative;
}

.stat-item:not(:last-child)::after {
  content: '';
  position: absolute;
  right: 0;
  top: 5px;
  height: 70%;
  width: 1px;
  background-color: #eee;
}

.stat-value {
  font-weight: bold;
  font-size: 16px;
}

.stat-label {
  color: #999;
  font-size: 12px;
}

.user-actions {
  margin-bottom: 15px;
}

.user-bio {
  color: #666;
  font-size: 14px;
  padding: 0 20px;
}

.video-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2px;
  padding: 2px;
}

.video-item {
  position: relative;
  aspect-ratio: 9/16;
}

.video-stats {
  position: absolute;
  bottom: 5px;
  left: 5px;
  color: #fff;
  font-size: 12px;
  display: flex;
  align-items: center;
  text-shadow: 0 0 2px rgba(0, 0, 0, 0.5);
}

.video-stats .van-icon {
  margin-right: 3px;
}

.empty-tip {
  grid-column: span 3;
  text-align: center;
  padding: 50px 0;
  color: #999;
}

.edit-form {
  padding: 20px;
}
</style>