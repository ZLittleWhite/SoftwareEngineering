import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'Index',
            component: () => import('@/views/')
        },
        {
            path: '/buytickets',
            name: 'BuyTickets',
            component: () => import('@/views/buytickets')
        },
        {
            path: '/ticketsdetails',
            name: 'TicketsDetails',
            component: () => import('@/components/ticketsdetails')
        },
        {
            path: '/mine',
            name: 'Mine',
            component: () => import('@/views/mine'),
        },
        {
            path: '/collection',
            name: 'MineCollection',
            component: () => import('@/components/mine/components/collection')
        },
        {
            path: '/shopcar',
            name: 'ShopCar',
            component: () => import('@/components/mine/components/shopcar')
        },
        {
            path: '/login',
            name: 'Login',
            component: () => import('@/components/mine/components/login')
        },
        {
            path: '/register',
            name: 'Register',
            component: () => import('@/components/mine/components/register')
        }
    ]
})
