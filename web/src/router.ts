import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Authorized from './views/Authorized.vue'
import Login from './views/Login.vue'
import auth from './services/auth-service'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/callback',
      name: 'callback',
      component: Authorized
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.path === '/' || to.path === '/callback' || auth.isAuthenticated()) {
    return next()
  }
  auth.login(to.path, process.env.VUE_APP_AUDIENCE)
})

export default router
