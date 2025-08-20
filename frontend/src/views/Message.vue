<template>
  <div class="message-container">
    <van-nav-bar title="消息" />
    
    <div class="message-tabs">
      <van-tabs v-model:active="activeTab" animated swipeable>
        <van-tab title="私信">
          <div class="chat-list">
            <div v-if="chatList.length === 0" class="empty-tip">
              暂无私信消息
            </div>
            <div 
              v-for="chat in chatList" 
              :key="chat.id" 
              class="chat-item"
              @click="goToChat(chat.id)"
            >
              <div class="chat-avatar">
                <van-image round width="50" height="50" :src="chat.avatar" />
                <div class="online-status" v-if="chat.online"></div>
              </div>
              <div class="chat-info">
                <div class="chat-header">
                  <div class="chat-name">{{ chat.name }}</div>
                  <div class="chat-time">{{ chat.lastTime }}</div>
                </div>
                <div class="chat-message">{{ chat.lastMessage }}</div>
              </div>
              <div class="chat-badge" v-if="chat.unread > 0">
                {{ chat.unread > 99 ? '99+' : chat.unread }}
              </div>
            </div>
          </div>
        </van-tab>
        
        <van-tab title="评论">
          <div class="comment-list">
            <div v-if="commentList.length === 0" class="empty-tip">
              暂无评论消息
            </div>
            <div 
              v-for="comment in commentList" 
              :key="comment.id" 
              class="comment-item"
              @click="goToVideo(comment.videoId)"
            >
              <div class="comment-avatar">
                <van-image round width="40" height="40" :src="comment.user.avatar" />
              </div>
              <div class="comment-info">
                <div class="comment-user">
                  {{ comment.user.username }}
                  <span class="comment-time">{{ comment.time }}</span>
                </div>
                <div class="comment-content">{{ comment.content }}</div>
                <div class="comment-video">
                  <van-image width="40" height="40" fit="cover" :src="comment.videoCover" />
                  <div class="video-title">{{ comment.videoTitle }}</div>
                </div>
              </div>
            </div>
          </div>
        </van-tab>
        
        <van-tab title="赞">
          <div class="like-list">
            <div v-if="likeList.length === 0" class="empty-tip">
              暂无点赞消息
            </div>
            <div 
              v-for="like in likeList" 
              :key="like.id" 
              class="like-item"
              @click="goToVideo(like.videoId)"
            >
              <div class="like-avatar">
                <van-image round width="40" height="40" :src="like.user.avatar" />
              </div>
              <div class="like-info">
                <div class="like-user">
                  {{ like.user.username }}
                  <span class="like-time">{{ like.time }}</span>
                </div>
                <div class="like-content">赞了你的视频</div>
                <div class="like-video">
                  <van-image width="40" height="40" fit="cover" :src="like.videoCover" />
                </div>
              </div>
            </div>
          </div>
        </van-tab>
        
        <van-tab title="关注">
          <div class="follow-list">
            <div v-if="followList.length === 0" class="empty-tip">
              暂无关注消息
            </div>
            <div 
              v-for="follow in followList" 
              :key="follow.id" 
              class="follow-item"
            >
              <div class="follow-avatar">
                <van-image round width="40" height="40" :src="follow.user.avatar" />
              </div>
              <div class="follow-info">
                <div class="follow-user">
                  {{ follow.user.username }}
                  <span class="follow-time">{{ follow.time }}</span>
                </div>
                <div class="follow-content">关注了你</div>
              </div>
              <div class="follow-action">
                <van-button 
                  size="small" 
                  type="primary" 
                  :loading="follow.followLoading"
                  @click.stop="toggleFollow(follow)"
                >
                  {{ follow.isFollowing ? '已关注' : '关注' }}
                </van-button>
              </div>
            </div>
          </div>
        </van-tab>
      </van-tabs>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Toast } from 'vant'
import { useUserStore } from '../store/user'

export default {
  name: 'Message',
  setup() {
    const router = useRouter()
    const userStore = useUserStore()
    
    const activeTab = ref(0)
    const chatList = ref([])
    const commentList = ref([])
    const likeList = ref([])
    const followList = ref([])
    
    // 获取私信列表
    const fetchChatList = () => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      setTimeout(() => {
        chatList.value = Array(10).fill().map((_, i) => ({
          id: `chat-${i}`,
          name: `用户${i}`,
          avatar: 'https://example.com/avatar.jpg',
          lastMessage: `这是最后一条消息内容 ${i}`,
          lastTime: '12:30',
          unread: Math.floor(Math.random() * 10),
          online: Math.random() > 0.5
        }))
      }, 500)
    }
    
    // 获取评论列表
    const fetchCommentList = () => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      setTimeout(() => {
        commentList.value = Array(10).fill().map((_, i) => ({
          id: `comment-${i}`,
          content: `这是一条评论内容，评论了你的视频 ${i}`,
          time: '3小时前',
          videoId: `video-${i}`,
          videoCover: 'https://example.com/cover.jpg',
          videoTitle: `视频标题 ${i}`,
          user: {
            id: `user-${i}`,
            username: `用户${i}`,
            avatar: 'https://example.com/avatar.jpg'
          }
        }))
      }, 500)
    }
    
    // 获取点赞列表
    const fetchLikeList = () => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      setTimeout(() => {
        likeList.value = Array(10).fill().map((_, i) => ({
          id: `like-${i}`,
          time: '5小时前',
          videoId: `video-${i}`,
          videoCover: 'https://example.com/cover.jpg',
          user: {
            id: `user-${i}`,
            username: `用户${i}`,
            avatar: 'https://example.com/avatar.jpg'
          }
        }))
      }, 500)
    }
    
    // 获取关注列表
    const fetchFollowList = () => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      setTimeout(() => {
        followList.value = Array(10).fill().map((_, i) => ({
          id: `follow-${i}`,
          time: '1天前',
          isFollowing: Math.random() > 0.5,
          followLoading: false,
          user: {
            id: `user-${i}`,
            username: `用户${i}`,
            avatar: 'https://example.com/avatar.jpg'
          }
        }))
      }, 500)
    }
    
    // 跳转到聊天页面
    const goToChat = (chatId) => {
      Toast('聊天功能开发中')
      // 实际项目中应该跳转到聊天详情页面
    }
    
    // 跳转到视频详情
    const goToVideo = (videoId) => {
      router.push(`/video/${videoId}`)
    }
    
    // 关注/取消关注
    const toggleFollow = async (follow) => {
      if (!userStore.isLoggedIn) {
        router.push('/login')
        return
      }
      
      follow.followLoading = true
      
      try {
        let success
        if (follow.isFollowing) {
          success = await userStore.unfollowUser(follow.user.id)
        } else {
          success = await userStore.followUser(follow.user.id)
        }
        
        if (success) {
          follow.isFollowing = !follow.isFollowing
        }
      } catch (error) {
        console.error('操作失败:', error)
        Toast('操作失败，请稍后再试')
      } finally {
        follow.followLoading = false
      }
    }
    
    onMounted(() => {
      fetchChatList()
      fetchCommentList()
      fetchLikeList()
      fetchFollowList()
    })
    
    return {
      activeTab,
      chatList,
      commentList,
      likeList,
      followList,
      goToChat,
      goToVideo,
      toggleFollow
    }
  }
}
</script>

<style scoped>
.message-container {
  min-height: 100vh;
  background-color: #f7f8fa;
  padding-bottom: 50px;
}

.message-tabs {
  background-color: #fff;
}

.empty-tip {
  padding: 50px 0;
  text-align: center;
  color: #999;
}

/* 私信样式 */
.chat-list {
  background-color: #fff;
}

.chat-item {
  display: flex;
  align-items: center;
  padding: 15px;
  position: relative;
  border-bottom: 1px solid #f5f5f5;
}

.chat-avatar {
  position: relative;
  margin-right: 12px;
}

.online-status {
  position: absolute;
  bottom: 0;
  right: 0;
  width: 10px;
  height: 10px;
  background-color: #4caf50;
  border-radius: 50%;
  border: 2px solid #fff;
}

.chat-info {
  flex: 1;
  overflow: hidden;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 5px;
}

.chat-name {
  font-weight: bold;
  font-size: 16px;
}

.chat-time {
  color: #999;
  font-size: 12px;
}

.chat-message {
  color: #666;
  font-size: 14px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.chat-badge {
  background-color: #ff2c55;
  color: #fff;
  font-size: 12px;
  min-width: 18px;
  height: 18px;
  border-radius: 9px;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 5px;
}

/* 评论样式 */
.comment-list {
  background-color: #fff;
}

.comment-item {
  display: flex;
  padding: 15px;
  border-bottom: 1px solid #f5f5f5;
}

.comment-avatar {
  margin-right: 12px;
}

.comment-info {
  flex: 1;
}

.comment-user {
  font-weight: bold;
  margin-bottom: 5px;
}

.comment-time {
  font-weight: normal;
  color: #999;
  font-size: 12px;
  margin-left: 5px;
}

.comment-content {
  margin-bottom: 10px;
  font-size: 14px;
}

.comment-video {
  display: flex;
  align-items: center;
  background-color: #f5f5f5;
  padding: 5px;
  border-radius: 4px;
}

.video-title {
  margin-left: 10px;
  font-size: 12px;
  color: #666;
}

/* 点赞样式 */
.like-list {
  background-color: #fff;
}

.like-item {
  display: flex;
  padding: 15px;
  border-bottom: 1px solid #f5f5f5;
}

.like-avatar {
  margin-right: 12px;
}

.like-info {
  flex: 1;
}

.like-user {
  font-weight: bold;
  margin-bottom: 5px;
}

.like-time {
  font-weight: normal;
  color: #999;
  font-size: 12px;
  margin-left: 5px;
}

.like-content {
  margin-bottom: 10px;
  font-size: 14px;
}

.like-video {
  width: 40px;
  height: 40px;
  border-radius: 4px;
  overflow: hidden;
}

/* 关注样式 */
.follow-list {
  background-color: #fff;
}

.follow-item {
  display: flex;
  align-items: center;
  padding: 15px;
  border-bottom: 1px solid #f5f5f5;
}

.follow-avatar {
  margin-right: 12px;
}

.follow-info {
  flex: 1;
}

.follow-user {
  font-weight: bold;
  margin-bottom: 5px;
}

.follow-time {
  font-weight: normal;
  color: #999;
  font-size: 12px;
  margin-left: 5px;
}

.follow-content {
  font-size: 14px;
}

.follow-action {
  margin-left: 10px;
}
</style>