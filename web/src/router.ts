import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Authorized from './views/Authorized.vue'
import auth from './services/auth-service'

Vue.use(Router)
process.env.audience = 'https://meizam-2020.herokuapp.com'

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
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
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

  auth.login(to.path, process.env.audience)
})

export default router
