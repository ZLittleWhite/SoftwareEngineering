import Vue from 'vue'
import App from './App.vue'
import router from './router'
import plugin from "./plugin";

import Vant from 'vant';
import axios from 'axios'
import 'vant/lib/index.css';
import {ShareSheet} from 'vant';
import Koacors from "koa2-cors"

axios.defaults.baseURL = '/api'
axios.defaults.headers.post['Content-Type'] = 'application/json';

Vue.use(plugin)
Vue.use(ShareSheet);
Vue.use(Vant);
Vue.use(Koacors);

Vue.prototype.$axios=axios
Vue.config.productionTip = false

new Vue({
    router,
    render: h => h(App),
}).$mount('#app')
