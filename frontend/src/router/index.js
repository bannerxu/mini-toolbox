import {createRouter, createWebHistory} from 'vue-router'
import Home from '../views/Home.vue'
import ImageUpload from '../views/ImageUpload.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        meta: {
            title: 'Mini Toolbox - 首页'
        }
    },
    {
        path: '/image-upload',
        name: 'ImageUpload',
        component: ImageUpload,
        meta: {
            title: 'Mini Toolbox - 图片上传'
        }
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return {top: 0}
        }
    }
})

// 路由守卫，更新页面标题
router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})

export default router