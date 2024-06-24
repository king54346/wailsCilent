import { createRouter, createWebHashHistory } from 'vue-router';

//路由
const routes = [
    {
        path: "/",
        redirect: '/game',
        component: () => import('@/layout/index.vue'),
        name: 'Game',
        meta: { title: 'Game', icon: 'Aim' },
        children: [
            {
                path: "/game",
                component: () => import('@/views/game/game.vue'),
                meta: { title: "game管理" }
            }
        ]
    },
    {
        path: '/annon',
        component: () => import('@/layout/index.vue'),
        name: 'Annon',
        meta: { title: 'Annon', icon: 'Aim' },
        children: [
            {
                path: "manage",
                name: 'AnnonManage',
                component: () => import('@/views/game/annon.vue'),
                meta: { title: "Annon管理" }
            }
        ]
    },
    {
        path: '/order',
        component: () => import('@/layout/index.vue'),
        name: 'Order',
        meta: { title: 'Order', icon: 'Management' },
        children: [
            {
                path: 'manage',
                name: 'OrderManage',
                component: () => import('@/views/game/order.vue'),
                meta: { title: '工单管理', icon: 'Tools' }
            },
        ]
    },
    {
        path: '/404',
        name: '404',
        component: () => import('@/views/exception/404.vue'),
    },
    {
        path: '/:pathMatch(.*)*',
        redirect: '/404'
    }
]


//创建路由
const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})

export default router;