<template lang="pug">
  v-app-bar(dark app)
    v-toolbar-title
      v-btn(text,router, to="/") DeLaWeb
    div.flex-grow-1
    // Попап авторизации и регистрации
    popap(v-if="!Login")
    // Роут на гравную и функция выхода из профиля
    v-btn( text v-if="Login" @click="logout", router, to="/") Выйти
    // Роут на страницу профиля, виден только при условии авторизации
    v-btn.mx-2(v-if="Login", router, to="/profile", icon, max-width="36" max-height="36")
      v-icon account_circle
</template>

<script>
  import {mapGetters, mapMutations} from "vuex"
  import Popap from "./popap";
  export default {
    name: "navbar",
    components: {Popap},
    data:()=>({
      open: false
    }),
    computed: mapGetters(['Login']),
    methods:{
      ...mapMutations(['Logout']),
      logout(){
        this.Logout()
      }
    }
  }
</script>

<style scoped>

</style>
