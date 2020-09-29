// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

//引入Vant 组件
import Vant from 'vant';
import 'vant/lib/index.css';

Vue.use(Vant);

Vue.config.productionTip = false

//去掉默认样式
import "@/assets/css/deleteorigincss.css"

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
