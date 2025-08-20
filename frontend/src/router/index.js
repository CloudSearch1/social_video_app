import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue'),
    meta: { keepAlive: true }
  },
  {
    path: '/discover',
    name: 'Discover',
    component: () => import('../views/Discover.vue'),
    meta: { keepAlive: true }
  },
  {
    path: '/upload',
    name: 'Upload',
    component: () => import('../views/Upload.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/message',
    name: 'Message',
    component: () => import('../views/Message.vue'),
    meta: { requiresAuth: true, keepAlive: true }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('../views/Profile.vue'),
    meta: { keepAlive: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Register.vue')
  },
  {
    path: '/video/:id',
    name: 'VideoDetail',
    component: () => import('../views/VideoDetail.vue')
  },
  {
    path: '/user/:id',
    name: 'UserProfile',
    component: () => import('../views/UserProfile.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  
  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router