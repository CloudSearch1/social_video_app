<template>
  <div class="user-profile">
    <van-nav-bar title="用户资料" left-text="返回" left-arrow @click-left="onClickLeft" />
    <div class="profile-content">
      <van-cell-group>
        <van-cell title="用户名" :value="user.username" />
        <van-cell title="粉丝数" :value="user.followers" />
        <van-cell title="关注数" :value="user.following" />
        <van-cell title="获赞数" :value="user.likes" />
      </van-cell-group>
    </div>
  </div>
</template>

<script>
import { useRouter } from 'vue-router'

export default {
  name: 'UserProfile',
  setup() {
    const router = useRouter()
    
    const user = {
      username: '示例用户',
      followers: '1000',
      following: '500',
      likes: '5000'
    }
    
    const onClickLeft = () => {
      router.go(-1)
    }
    
    return {
      user,
      onClickLeft
    }
  }
}
</script>

<style scoped>
.user-profile {
  background-color: #000;
  min-height: 100vh;
}

.profile-content {
  padding: 10px;
}
</style>