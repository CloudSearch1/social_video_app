<template>
  <div class="login-container">
    <van-nav-bar
      title="登录"
      left-arrow
      @click-left="onClickLeft"
    />
    
    <div class="form-container">
      <van-cell-group inset>
        <van-cell title="用户名" :border="false">
          <template #right-icon>
            <van-field
              v-model="username"
              placeholder="请输入用户名"
              :error="!!errors.username"
              :error-message="errors.username"
              @focus="clearError('username')"
            />
          </template>
        </van-cell>
        
        <van-cell title="密码" :border="false">
          <template #right-icon>
            <van-field
              v-model="password"
              type="password"
              placeholder="请输入密码"
              :error="!!errors.password"
              :error-message="errors.password"
              @focus="clearError('password')"
            />
          </template>
        </van-cell>
      </van-cell-group>
      
      <div class="button-container">
        <van-button 
          type="primary" 
          block 
          :loading="userStore.loading" 
          @click="handleLogin"
        >
          登录
        </van-button>
      </div>
      
      <div class="register-link">
        还没有账号？<router-link to="/register">立即注册</router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Toast } from 'vant'
import { useUserStore } from '../store/user'

export default {
  name: 'Login',
  setup() {
    const router = useRouter()
    const userStore = useUserStore()
    
    const username = ref('')
    const password = ref('')
    const errors = ref({
      username: '',
      password: ''
    })
    
    const validateForm = () => {
      let isValid = true
      
      if (!username.value.trim()) {
        errors.value.username = '请输入用户名'
        isValid = false
      }
      
      if (!password.value) {
        errors.value.password = '请输入密码'
        isValid = false
      }
      
      return isValid
    }
    
    const clearError = (field) => {
      errors.value[field] = ''
    }
    
    const handleLogin = async () => {
      if (!validateForm()) return
      
      const success = await userStore.login(username.value, password.value)
      
      if (success) {
        Toast.success('登录成功')
        router.replace('/')
      } else {
        Toast.fail(userStore.error || '登录失败，请稍后再试')
      }
    }
    
    const onClickLeft = () => {
      router.back()
    }
    
    return {
      username,
      password,
      errors,
      userStore,
      clearError,
      handleLogin,
      onClickLeft
    }
  }
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  background-color: #f7f8fa;
  display: flex;
  flex-direction: column;
}

.form-container {
  padding: 20px;
  margin-top: 20px;
}

.button-container {
  margin-top: 30px;
}

.register-link {
  margin-top: 20px;
  text-align: center;
  font-size: 14px;
}

.register-link a {
  color: #1989fa;
  text-decoration: none;
}
</style>