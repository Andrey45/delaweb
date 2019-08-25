<template lang="pug">
  div
    // Кнопка открытия модального окна
    v-btn( outlined @click="open = !open") Войти
    v-dialog(v-model="open" max-width="400" dark)
      v-card
        v-card-actions
          // Кнопка перехода на авторизацию
          v-btn(@click="window = 1" text) Авторизация
          v-divider(inset vertical)
          // Кнопка перехода на ругустрацию
          v-btn(@click="window = 2" text) Регестрация
        v-window(v-model="window")
          // Отрисовываеться при авторизации
          v-window-item( :value="1")
            login( @close="open = !open")
          // Отрисовываеться при регустрации
          v-window-item( :value="2")
            check(@window="window = 1")

</template>
<script>
  import Login from "./componentPopap/login";
  import Check from "./componentPopap/check";
  export default {
    name: "popap",
    components: {Check, Login},
    data:() => ({
      open: false,
      window: 1,
    }),
  }
</script>

<style scoped>

</style>
