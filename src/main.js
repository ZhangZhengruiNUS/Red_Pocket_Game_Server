// import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import { createVuetify } from 'vuetify'
import 'vuetify/styles'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import router from "./router"
import APP from "./route.vue"
// import APP from './Login.vue'

const app = createApp(APP)

const vuetify = createVuetify({
    components,
    directives,
})

app.use(router)
app.use(createPinia())
app.use(vuetify)

app.mount('#app')



