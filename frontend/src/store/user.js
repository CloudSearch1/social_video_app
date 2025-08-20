import { defineStore } from 'pinia'
import axios from 'axios'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    user: JSON.parse(localStorage.getItem('user')) || null,
    loading: false,
    error: null
  }),
  
  getters: {
    isLoggedIn: (state) => !!state.token,
    currentUser: (state) => state.user
  },
  
  actions: {
    async login(username, password) {
      try {
        this.loading = true
        this.error = null
        
        // 实际项目中应该调用后端API
        const response = await axios.post('/api/users/login', { username, password })
        
        const { token, user } = response.data
        
        this.token = token
        this.user = user
        
        localStorage.setItem('token', token)
        localStorage.setItem('user', JSON.stringify(user))
        
        return true
      } catch (error) {
        this.error = error.response?.data?.message || '登录失败，请稍后再试'
        return false
      } finally {
        this.loading = false
      }
    },
    
    async register(username, password) {
      try {
        this.loading = true
        this.error = null
        
        // 实际项目中应该调用后端API
        const response = await axios.post('/api/users/register', { username, password })
        
        const { token, user } = response.data
        
        this.token = token
        this.user = user
        
        localStorage.setItem('token', token)
        localStorage.setItem('user', JSON.stringify(user))
        
        return true
      } catch (error) {
        this.error = error.response?.data?.message || '注册失败，请稍后再试'
        return false
      } finally {
        this.loading = false
      }
    },
    
    logout() {
      this.token = null
      this.user = null
      
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    },
    
    async updateProfile(userData) {
      try {
        this.loading = true
        this.error = null
        
        // 实际项目中应该调用后端API
        const response = await axios.put('/api/users/profile', userData, {
          headers: { Authorization: `Bearer ${this.token}` }
        })
        
        const updatedUser = response.data
        
        this.user = updatedUser
        localStorage.setItem('user', JSON.stringify(updatedUser))
        
        return true
      } catch (error) {
        this.error = error.response?.data?.message || '更新个人资料失败，请稍后再试'
        return false
      } finally {
        this.loading = false
      }
    },
    
    async followUser(userId) {
      try {
        if (!this.isLoggedIn) return false
        
        // 实际项目中应该调用后端API
        await axios.post(`/api/users/${userId}/follow`, {}, {
          headers: { Authorization: `Bearer ${this.token}` }
        })
        
        return true
      } catch (error) {
        this.error = error.response?.data?.message || '关注用户失败，请稍后再试'
        return false
      }
    },
    
    async unfollowUser(userId) {
      try {
        if (!this.isLoggedIn) return false
        
        // 实际项目中应该调用后端API
        await axios.delete(`/api/users/${userId}/follow`, {
          headers: { Authorization: `Bearer ${this.token}` }
        })
        
        return true
      } catch (error) {
        this.error = error.response?.data?.message || '取消关注失败，请稍后再试'
        return false
      }
    }
  }
})