import {createApp} from 'vue'
import App from './App.vue'
import axios from "axios";
import VueAxios from "vue-axios";

const client = axios.create({
    baseURL: "/api/v1",
});

const app = createApp(App)
app.use(VueAxios, client)

app.mount('#app')
