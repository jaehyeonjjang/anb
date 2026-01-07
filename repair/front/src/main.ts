import { createApp } from 'vue'
import router from './router'
import App from './App.vue'
import store from './store'
import { setupGlobalDirectives } from './directives'
import v3ImgPreview from 'v3-img-preview'

// import '~/styles/element/index.scss'
//


import ElementPlus from 'element-plus'
// import all element css, uncommented next line
import 'element-plus/dist/index.css'

// or use cdn, uncomment cdn link in `index.html`

import '~/styles/index.scss'

// If you want to use ElMessage, import it.
import 'element-plus/theme-chalk/src/message.scss'
import koKR from 'element-plus/es/locale/lang/ko'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'

const app = createApp(App)
setupGlobalDirectives(app)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.use(router)
app.use(store)
app.use(ElementPlus, {
  locale: koKR,
})
app.use(v3ImgPreview, {  
})
app.mount('#app')
