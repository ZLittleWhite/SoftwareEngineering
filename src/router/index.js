import Vue from 'vue'
import Router from 'vue-router'


Vue.use(Router)

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      name: 'AppIndex',
      component: () => import("@/views/index")
    },
    {
      path: '/mine',
      name: 'Mine',
      component: () => import("@/views/mine")
    },
    {
      path: '/detail',
      name: 'Detail',
      component: () => import("@/views/detail")
    },
    {
      path: '/product',
      name: 'Product',
      component: () => import("@/views/productsort")
    },
    {
      path: '/useragree',
      name: 'UserAgree',
      component: () => import("@/views/useragree")
    },
    {
      path: '/mine/codelogin',
      name: 'CodeLogin',
      component: () => import("@/components/mine/codelogin")
    },
    {
      path: '/mine/freeregister',
      name: 'FreeRegister',
      component: () => import("@/components/mine/freeregister")
    },
    {
      path: '/mine/forgetpassword',
      name: 'ForgetPassword',
      component: () => import("@/components/mine/forgetpassword")
    },
    {
      path: '/search',
      name: 'Search',
      component: () => import("@/components/public/search")
    }
  ]
})
