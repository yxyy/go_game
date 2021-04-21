import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '@/views/login'
import Layout from '@/views/layout'
import Home from '@/views/home'
import Menu from '@/views/menu/index.vue'
import User from '@/views/user'
import UserInfo from '@/views/userInfo'

Vue.use(VueRouter)

const routes = [
  {
      path: '/login',
      name: 'login',
      component:Login
  },
  {
      path: '/',
      component: Layout,
      children:[
        {
          path: '/',
          component: Home
        },
        {
          path: '/home',
          name: 'home',
          component: Home
        },
        {
          path: '/menu',
          name: 'menu',
          component: Menu
        },
        {
          path: '/user',
          name: '管理员管理',
          component: User
        },
        {
          path: '/info',
          name: 'userInfo',
          component: UserInfo
        }
      ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  console.log(to.path)
  if(to.path=="/login"){
    next()
  }else{
    let access_token = window.localStorage.getItem("access_token")
    if (!access_token) {
        next('/login?redirect='+ to.path)
        return
      }
    // 登陆
    next()
  }
})

export default router
