import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import PageHome from './components/pages/PageHome.vue'
import PageMedia from './components/pages/PageMedia.vue'
import BootstrapVue from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false
// Vue.prototype.$http = axios
Vue.use(VueRouter)
Vue.use(BootstrapVue)

const router = new VueRouter({
  routes: [
    { path: "/", component: PageHome, props: true },
    { path: "/media/:id", component: PageMedia, props: true }
  ]
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
