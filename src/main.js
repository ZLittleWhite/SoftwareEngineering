import Vue from 'vue'
import App from './App.vue'
import Vant from 'vant';

import router from './router'
import plugin from "./plugin";

import 'vant/lib/index.css';
import {ShareSheet} from 'vant';

import axios from 'axios'

axios.defaults.headers.post['Content-Type'] = 'application/json';

Vue.use(plugin)
Vue.use(ShareSheet);
Vue.use(Vant);

Vue.prototype.$axios=axios
Vue.config.productionTip = false

new Vue({
    router,
    render: h => h(App),
}).$mount('#app')
