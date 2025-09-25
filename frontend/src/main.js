import { createApp } from "vue"
import App from "./App.vue"
import router from "./router"

//import "./index.css" // optional if you use Tailwind or global css

const app = createApp(App)
app.use(router)
app.mount("#app")
