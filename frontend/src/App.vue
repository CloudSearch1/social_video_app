<template>
  <div class="app-container">
    <router-view v-slot="{ Component }">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </router-view>
    
    <van-tabbar v-model="activeTab" route>
      <van-tabbar-item to="/" icon="home-o">首页</van-tabbar-item>
      <van-tabbar-item to="/discover" icon="search">发现</van-tabbar-item>
      <van-tabbar-item to="/upload" icon="plus">上传</van-tabbar-item>
      <van-tabbar-item to="/message" icon="chat-o">消息</van-tabbar-item>
      <van-tabbar-item to="/profile" icon="user-o">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  name: 'App',
  setup() {
    const activeTab = ref(0)
    
    return {
      activeTab
    }
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body {
  height: 100%;
  width: 100%;
  overflow: hidden;
  background-color: #000;
  color: #fff;
  font-family: 'PingFang SC', 'Helvetica Neue', Helvetica, Arial, sans-serif;
}

.app-container {
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
}

/* 自定义Vant主题 */
:root {
  --van-black: #000;
  --van-white: #fff;
  --van-gray-8: #323233;
  --van-tabbar-background-color: rgba(0, 0, 0, 0.8);
  --van-tabbar-item-text-color: #999;
  --van-tabbar-item-active-color: #fff;
}
</style>