<template>
  <div class="register-container">
    <van-nav-bar
      title="注册"
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
        
        <van-cell title="确认密码" :border="false">
          <template #right-icon>
            <van-field
              v-model="confirmPassword"
              type="password"
              placeholder="请再次输入密码"
              :error="!!errors.confirmPassword"
              :error-message="errors.confirmPassword"
              @focus="clearError('confirmPassword')"
            />
          </template>
        </van-cell>
      </van-cell-group>
      
      <div class="button-container">
        <van-button 
          type="primary" 
          block 
          :loading="userStore.loading" 
          @click="handleRegister"
        >
          注册
        </van-button>
      </div>
      
      <div class="login-link">
        已有账号？<router-link to="/login">立即登录</router-link>
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
  name: 'Register',
  setup() {
    const router = useRouter()
    const userStore = useUserStore()
    
    const username = ref('')
    const password = ref('')
    const confirmPassword = ref('')
    const errors = ref({
      username: '',
      password: '',
      confirmPassword: ''
    })
    
    const validateForm = () => {
      let isValid = true
      
      if (!username.value.trim()) {
        errors.value.username = '请输入用户名'
        isValid = false
      } else if (username.value.length < 3) {
        errors.value.username = '用户名至少需要3个字符'
        isValid = false
      }
      
      if (!password.value) {
        errors.value.password = '请输入密码'
        isValid = false
      } else if (password.value.length < 6) {
        errors.value.password = '密码至少需要6个字符'
        isValid = false
      }
      
      if (!confirmPassword.value) {
        errors.value.confirmPassword = '请确认密码'
        isValid = false
      } else if (confirmPassword.value !== password.value) {
        errors.value.confirmPassword = '两次输入的密码不一致'
        isValid = false
      }
      
      return isValid
    }
    
    const clearError = (field) => {
      errors.value[field] = ''
    }
    
    const handleRegister = async () => {
      if (!validateForm()) return
      
      const success = await userStore.register(username.value, password.value)
      
      if (success) {
        Toast.success('注册成功')
        router.replace('/')
      } else {
        Toast.fail(userStore.error || '注册失败，请稍后再试')
      }
    }
    
    const onClickLeft = () => {
      router.back()
    }
    
    return {
      username,
      password,
      confirmPassword,
      errors,
      userStore,
      clearError,
      handleRegister,
      onClickLeft
    }
  }
}
</script>

<style scoped>
.register-container {
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

.login-link {
  margin-top: 20px;
  text-align: center;
  font-size: 14px;
}

.login-link a {
  color: #1989fa;
  text-decoration: none;
}
</style>