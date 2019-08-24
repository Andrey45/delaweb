<template lang="pug">
  div
    v-btn(@click="open = !open") Войти
    v-dialog(v-model="open" max-width="400" dark)
      v-card
        v-card-actions
          v-btn(@click="window = 1" text) Авторизация
          v-divider(inset vertical)
          v-btn(@click="window = 2" text) Регестрация
        v-window(v-model="window")
          v-window-item( :value="1")
            v-card-actions.px-4
              v-text-field(
                v-model="email" label="Электронная почта"
                :rules="[ rules.emaill, rules.required]")
            v-card-actions.px-4
              v-text-field(
                v-model="password"
                :append-icon="show1 ? 'visibility' : 'visibility_off'"
                :rules="[rules.required, rules.minLength]"
                :type="show1 ? 'text' : 'password'"
                name="input-10-1"
                label="Пароль"
                hint="At least 8 characters"
                counter
                @click:append="show1 = !show1")
            v-card-actions
              v-btn(@click="auth()" outlined) Войти
          v-window-item( :value="2")
            v-card-actions.px-4
              v-text-field(
                v-model="email" label="Электронная почта"
                :rules="[ rules.emaill, rules.required]")
            v-card-actions.px-4
              v-text-field(
                v-model="password"
                :append-icon="show1 ? 'visibility' : 'visibility_off'"
                :rules="[rules.required, rules.minLength]"
                :type="show1 ? 'text' : 'password'"
                name="input-10-1"
                label="Пароль"
                hint="At least 8 characters"
                counter
                @click:append="show1 = !show1")

</template>

<script>
  import axios from "axios"
  export default {
    name: "popap",
    data:() => ({
      open: false,
      loading: false,
      window: 1,
      show1: false,
      email: '',
      password: '',
      rules:{
        required: value => !!value || 'Пожалуйста заполните.',
        minLength: value => value.length >= 8 || 'Не корректный пароль',
        emaill: value => {
          const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          return pattern.test(value) || 'Не корректный E-mail.'
        }
      }
    }),
    computed: {
      currentTitle() {
        switch (this.window) {
          case 1:
            return 'Авторизация'
          case 2:
            return 'Регистрация'
        }
      }
    },
    methods:{
      auth() {
        this.loading = true
        axios.post('https://delaweb.herokuapp.com/autorization', {
          email: this.email,
          password: this.password
        }).then(res => {
          console.log(res.status)
          if (res.data.email !== null && res.data.password !== null) {
            this.loading = false
            this.open = false
            this.$router.push({path: '/profile'})
          } else {
            this.loading = false
          }
        })
      }
    }
  }
</script>

<style scoped>

</style>
