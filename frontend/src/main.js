import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import 'element-plus/dist/index.css'
import router from '@/router'
import store from '@/store'
import plugin from '@/utils/plugin'
import '@/styles/index.scss' // global css
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import setupComponent from '@/components'
// import 'highlight.js/styles/atom-one-dark.css'
// import 'highlight.js/lib/common'
const app = createApp(App)

import './permission.js'

//注册所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}


app.use(plugin)
app.use(router)
app.use(store)
app.use(ElementPlus, { size: 'default',locale: zhCn });
//全局组件注册
setupComponent(app)

app.mount('#app')