// import './assets/main.css'
// import './assets/base.css'
import '@/assets/body.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'

import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

router.beforeEach((to, from, next) => {
    NProgress.start()
    next()
})

router.afterEach(() => {
    NProgress.done()
})

const app = createApp(App)

app.use(router)

app.use(ElementPlus,{ locale: zhCn })

app.mount('#app')
