import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'

import 'element-ui/lib/theme-chalk/index.css';
import { Pagination } from 'element-ui';

// css
import './assets/css/pace.css'
import './assets/css/custom.css'
import "./assets/css/other.css"

// js

Vue.use(Pagination);
Vue.config.productionTip = false
Vue.prototype.$http = axios // 修改内部的$http为axios

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
