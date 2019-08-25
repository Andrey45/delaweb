import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/store'
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false

router.beforeEach((to,from, next)=>{ // Хук ловит неавторизованных потзователей и отправляет на главную
  let login = store.getters.Login
  if (to.matched.some(url => url.meta.login)) {
    if (login === false) {
      next({path: '/'})
    }else {
      next()
    }
  } else {
    next()
  }
})
new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
