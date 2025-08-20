<template>
  <div class="video-detail-container">
    <van-nav-bar
      title="视频详情"
      left-arrow
      @click-left="onClickLeft"
    />
    
    <div class="video-container">
      <video
        ref="videoRef"
        class="video-player"
        :src="video.url"
        :poster="video.coverUrl"
        controls
        playsinline
        webkit-playsinline
        x5-video-player-type="h5"
        x5-video-player-fullscreen="true"
        x5-video-orientation="portrait"
      ></video>
    </div>
    
    <div class="video-info">
      <div class="user-info">
        <div class="avatar" @click="goToUserProfile">
          <van-image round width="40" height="40" :src="video.user.avatar" />
        </div>
        <div class="username" @click="goToUserProfile">{{ video.user.username }}</div>
        <div class="follow-btn">
          <van-button 
            size="small" 
            type="primary" 
            :loading="followLoading"
            @click="toggleFollow"
          >
            {{ video.user.isFollowing ? '已关注' : '关注' }}
          </van-button>
        </div>
      </div>
      
      <div class="video-desc">{{ video.description }}</div>
      
      <div class="video-stats">
        <div class="stat-item">
          <van-icon name="like-o" :color="video.isLiked ? '#ff2c55' : '#666'" @click="handleLike" />
          <span>{{ formatCount(video.likes) }}</span>
        </div>
        <div class="stat-item">
          <van-icon name="comment-o" />
          <span>{{ formatCount(video.comments) }}</span>
        </div>
        <div class="stat-item">
          <van-icon name="share-o" @click="handleShare" />
          <span>{{ formatCount(video.shares) }}</span>
        </div>
      </div>
    </div>
    
    <div class="comment-section">
      <div class="section-title">评论 ({{ commentList.length }})</div>
      
      <div class="comment-list">
        <div v-if="commentList.length === 0" class="empty-comment">
          暂无评论，快来抢沙发吧！
        </div>
        
        <div v-for="comment in commentList" :key="comment.id" class="comment-item">
          <div class="comment-avatar">
            <van-image round width="40" height="40" :src="comment.user.avatar" />
          </div>
          <div class="comment-content">
            <div class="comment-user">{{ comment.user.username }}</div>
            <div class="comment-text">{{ comment.content }}</div>
            <div class="comment-footer">
              <span class="comment-time">{{ comment.time }}</span>
              <div class="comment-actions">
                <span class="reply-btn" @click="replyToComment(comment)">回复</span>
                <div class="like-btn" @click="likeComment(comment)">
                  <van-icon name="like-o" :color="comment.isLiked ? '#ff2c55' : '#999'" />
                  <span>{{ formatCount(comment.likes) }}</span>
                </div>
              </div>
            </div>
            
            <!-- 回复列表 -->
            <div v-if="comment.replies && comment.replies.length > 0" class="reply-list">
              <div v-for="reply in comment.replies" :key="reply.id" class="reply-item">
                <div class="reply-content">
                  <span class="reply-user">{{ reply.user.username }}</span>
                  <span v-if="reply.replyTo" class="reply-to">回复 {{ reply.replyTo.username }}：</span>
                  <span>{{ reply.content }}</span>
                </div>
                <div class="reply-footer">
                  <span class="reply-time">{{ reply.time }}</span>
                  <span class="reply-btn" @click="replyToComment(comment, reply)">回复</span>
                </div>
              </div>
              
              <div v-if="comment.replyCount > comment.replies.length" class="more-replies" @click="loadMoreReplies(comment)">
                查看更多{{ comment.replyCount - comment.replies.length }}条回复
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <div class="comment-input">
      <van-field
        v-model="commentText"
        placeholder="添加评论..."
        :border="false"
        :disabled="!userStore.isLoggedIn"
        @keypress.enter="submitComment"
      >
        <template #button>
          <van-button 
            size="small" 
            type="primary" 
            :disabled="!commentText.trim() || !userStore.isLoggedIn"
            @click="submitComment"
          >
            发送
          </van-button>
        </template>
      </van-field>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Toast } from 'vant'
import { useUserStore } from '../store/user'

export default {
  name: 'VideoDetail',
  setup() {
    const route = useRoute()
    const router = useRouter()
    const userStore = useUserStore()
    
    const videoRef = ref(null)
    const videoId = route.params.id
    const video = ref({
      id: videoId,
      url: 'https://example.com/video.mp4',
      coverUrl: 'https://example.com/cover.jpg',
      description: '这是一个示例视频描述 #抖音 #示例',
      likes: 1024,
      comments: 256,
      shares: 128,
      isLiked: false,
      user: {
        id: '100',
        username: '用户100',
        avatar: 'https://example.com/avatar.jpg',
        isFollowing: false
      }
    })
    
    const commentList = ref([])
    const commentText = ref('')
    const followLoading = ref(false)
    const replyingTo = ref(null)
    
    // 获取视频详情
    const fetchVideoDetail = () => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      setTimeout(() => {
        // 视频数据已经在上面初始化了
      }, 500)
    }
    
    // 获取评论列表
    const fetchComments = () => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      setTimeout(() => {
        commentList.value = Array(5).fill().map((_, i) => ({
          id: `comment-${i}`,
          content: `这是一条评论内容 ${i}`,
          time: '3小时前',
          likes: Math.floor(Math.random() * 100),
          isLiked: Math.random() > 0.7,
          replyCount: Math.floor(Math.random() * 10),
          user: {
            id: `user-${200 + i}`,
            username: `用户${200 + i}`,
            avatar: 'https://example.com/avatar.jpg'
          },
          replies: Array(Math.min(3, Math.floor(Math.random() * 5))).fill().map((_, j) => ({
            id: `reply-${i}-${j}`,
            content: `这是一条回复内容 ${j}`,
            time: '2小时前',
            user: {
              id: `user-${300 + j}`,
              username: `用户${300 + j}`,
              avatar: 'https://example.com/avatar.jpg'
            },
            replyTo: j > 0 ? {
              id: `user-${300 + j - 1}`,
              username: `用户${300 + j - 1}`
            } : null
          }))
        }))
      }, 500)
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
    
    // 点赞视频
    const handleLike = () => {
      if (!userStore.isLoggedIn) {
        router.push('/login')
        return
      }
      
      video.value.isLiked = !video.value.isLiked
      video.value.likes += video.value.isLiked ? 1 : -1
      
      // 实际项目中应该调用API更新点赞状态
    }
    
    // 分享视频
    const handleShare = () => {
      Toast('分享功能开发中')
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
        if (video.value.user.isFollowing) {
          success = await userStore.unfollowUser(video.value.user.id)
        } else {
          success = await userStore.followUser(video.value.user.id)
        }
        
        if (success) {
          video.value.user.isFollowing = !video.value.user.isFollowing
        }
      } catch (error) {
        console.error('操作失败:', error)
        Toast('操作失败，请稍后再试')
      } finally {
        followLoading.value = false
      }
    }
    
    // 点赞评论
    const likeComment = (comment) => {
      if (!userStore.isLoggedIn) {
        router.push('/login')
        return
      }
      
      comment.isLiked = !comment.isLiked
      comment.likes += comment.isLiked ? 1 : -1
      
      // 实际项目中应该调用API更新点赞状态
    }
    
    // 回复评论
    const replyToComment = (comment, reply = null) => {
      if (!userStore.isLoggedIn) {
        router.push('/login')
        return
      }
      
      const replyToUser = reply ? reply.user.username : comment.user.username
      commentText.value = `回复 @${replyToUser}: `
      replyingTo.value = {
        commentId: comment.id,
        replyId: reply ? reply.id : null,
        userId: reply ? reply.user.id : comment.user.id
      }
      
      // 聚焦输入框
      setTimeout(() => {
        document.querySelector('.comment-input input').focus()
      }, 100)
    }
    
    // 加载更多回复
    const loadMoreReplies = (comment) => {
      // 模拟API请求
      // 实际项目中应该调用后端API
      Toast.loading({
        message: '加载中...',
        forbidClick: true,
        duration: 1000
      })
      
      setTimeout(() => {
        const newReplies = Array(Math.min(5, comment.replyCount - comment.replies.length)).fill().map((_, i) => ({
          id: `reply-${comment.id}-${comment.replies.length + i}`,
          content: `这是加载的更多回复内容 ${i}`,
          time: '1天前',
          user: {
            id: `user-${400 + i}`,
            username: `用户${400 + i}`,
            avatar: 'https://example.com/avatar.jpg'
          },
          replyTo: i > 0 ? {
            id: `user-${400 + i - 1}`,
            username: `用户${400 + i - 1}`
          } : null
        }))
        
        comment.replies = [...comment.replies, ...newReplies]
      }, 1000)
    }
    
    // 提交评论
    const submitComment = () => {
      if (!userStore.isLoggedIn) {
        router.push('/login')
        return
      }
      
      if (!commentText.value.trim()) return
      
      // 模拟API请求
      // 实际项目中应该调用后端API
      Toast.loading({
        message: '发送中...',
        forbidClick: true,
        duration: 0
      })
      
      setTimeout(() => {
        if (replyingTo.value) {
          // 回复评论
          const comment = commentList.value.find(c => c.id === replyingTo.value.commentId)
          if (comment) {
            const newReply = {
              id: `reply-new-${Date.now()}`,
              content: commentText.value.replace(/^回复 @[\w]+: /, ''),
              time: '刚刚',
              user: {
                id: userStore.user.id,
                username: userStore.user.username,
                avatar: userStore.user.avatar || 'https://example.com/avatar.jpg'
              },
              replyTo: replyingTo.value.replyId ? {
                id: replyingTo.value.userId,
                username: commentText.value.match(/@([\w]+)/)[1]
              } : null
            }
            
            if (!comment.replies) {
              comment.replies = []
            }
            
            comment.replies.unshift(newReply)
            comment.replyCount = (comment.replyCount || 0) + 1
          }
        } else {
          // 新评论
          const newComment = {
            id: `comment-new-${Date.now()}`,
            content: commentText.value,
            time: '刚刚',
            likes: 0,
            isLiked: false,
            replyCount: 0,
            user: {
              id: userStore.user.id,
              username: userStore.user.username,
              avatar: userStore.user.avatar || 'https://example.com/avatar.jpg'
            },
            replies: []
          }
          
          commentList.value.unshift(newComment)
          video.value.comments += 1
        }
        
        Toast.clear()
        Toast.success('评论成功')
        
        commentText.value = ''
        replyingTo.value = null
      }, 1000)
    }
    
    // 跳转到用户主页
    const goToUserProfile = () => {
      router.push(`/user/${video.value.user.id}`)
    }
    
    // 返回上一页
    const onClickLeft = () => {
      router.back()
    }
    
    onMounted(() => {
      fetchVideoDetail()
      fetchComments()
      
      // 自动播放视频
      if (videoRef.value) {
        videoRef.value.play().catch(err => {
          console.error('视频播放失败:', err)
        })
      }
    })
    
    onBeforeUnmount(() => {
      // 暂停视频
      if (videoRef.value) {
        videoRef.value.pause()
      }
    })
    
    return {
      videoRef,
      video,
      commentList,
      commentText,
      followLoading,
      userStore,
      formatCount,
      handleLike,
      handleShare,
      toggleFollow,
      likeComment,
      replyToComment,
      loadMoreReplies,
      submitComment,
      goToUserProfile,
      onClickLeft
    }
  }
}
</script>

<style scoped>
.video-detail-container {
  min-height: 100vh;
  background-color: #fff;
  display: flex;
  flex-direction: column;
}

.video-container {
  width: 100%;
  height: 56.25vw; /* 16:9 宽高比 */
  background-color: #000;
  position: relative;
}

.video-player {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.video-info {
  padding: 15px;
  border-bottom: 8px solid #f5f5f5;
}

.user-info {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.avatar {
  margin-right: 10px;
}

.username {
  flex: 1;
  font-weight: bold;
}

.video-desc {
  margin-bottom: 15px;
  font-size: 14px;
  line-height: 1.5;
}

.video-stats {
  display: flex;
  justify-content: space-around;
  padding-top: 10px;
  border-top: 1px solid #f5f5f5;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-item .van-icon {
  font-size: 24px;
  margin-bottom: 5px;
}

.comment-section {
  flex: 1;
  padding: 15px;
}

.section-title {
  font-weight: bold;
  margin-bottom: 15px;
}

.empty-comment {
  text-align: center;
  color: #999;
  padding: 20px 0;
}

.comment-item {
  display: flex;
  margin-bottom: 20px;
}

.comment-avatar {
  margin-right: 10px;
}

.comment-content {
  flex: 1;
}

.comment-user {
  font-weight: bold;
  margin-bottom: 5px;
}

.comment-text {
  font-size: 14px;
  margin-bottom: 5px;
}

.comment-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
  color: #999;
}

.comment-actions {
  display: flex;
  align-items: center;
}

.reply-btn {
  margin-right: 15px;
  cursor: pointer;
}

.like-btn {
  display: flex;
  align-items: center;
}

.like-btn .van-icon {
  margin-right: 3px;
}

.reply-list {
  margin-top: 10px;
  padding: 10px;
  background-color: #f9f9f9;
  border-radius: 4px;
}

.reply-item {
  margin-bottom: 10px;
}

.reply-content {
  font-size: 13px;
  margin-bottom: 3px;
}

.reply-user {
  font-weight: bold;
  margin-right: 5px;
}

.reply-to {
  color: #666;
}

.reply-footer {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  color: #999;
}

.more-replies {
  text-align: center;
  color: #1989fa;
  font-size: 12px;
  padding: 5px 0;
  cursor: pointer;
}

.comment-input {
  padding: 10px;
  border-top: 1px solid #f5f5f5;
  background-color: #fff;
  position: sticky;
  bottom: 0;
}
</style>