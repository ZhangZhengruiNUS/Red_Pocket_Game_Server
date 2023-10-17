// import './assets/main.css'

import { createApp } from 'vue'
import router from "./router"
import APP from "./route.vue"
// import APP from './Login.vue'

const app = createApp(APP)

app.use(router)

app.mount('#app')



