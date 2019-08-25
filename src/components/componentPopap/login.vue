<template lang="pug">
  v-form(v-model="valid")
    // Ошибка
    v-alert(:value="alert" type="error")
      | Пользователя не существует или введены некорректные данные
    v-card-actions.px-4
      // Input для email
      v-text-field(
        v-model="email" label="Электронная почта"
        :rules="[ rules.emaill, rules.required]")
    v-card-actions.px-4
      // Input для пароля
      v-text-field(
        v-model="password"
        :append-icon="show1 ? 'visibility' : 'visibility_off'"
        :rules="[rules.required, rules.minLength]"
        :type="show1 ? 'text' : 'password'"
        label="Пароль"
        hint="Не менее 8 символов"
        @click:append="show1 = !show1")
    v-card-actions
      // Кнопка активна если все поля заполены, если нет то нажать невозожно
      // При клике вызывает функцию log()
      v-btn(@click="log" outlined, :loading="loading" :disabled="!valid || loading") Войти
</template>

<script>
  import {mapActions} from 'vuex'
  export default {
    name: "login",
    data:()=>({
      valid: false,
      show1: false,
      rules:{
        // Валидация input-ов
        required: value => !!value || 'Пожалуйста заполните!',
        minLength: value => value.length >= 8 || 'Пароль не менее 8 символов!',
        emaill: value => {
          const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          return pattern.test(value) || 'Некорректный E-mail.'
        }
      },
      email: '',
      password: '',
      loading: false,
      alert: false
    }),
    methods:{
      ...mapActions(['login']),
      log() {
        this.loading = !this.loading
        this.login({ email: this.email, password: this.password }) // Вызывает actions 'login'
            .then(()=>{ // Если Promise вернул resolve то переводит на страницу профиля
              this.loading = !this.loading
              this.$emit('close')
              this.$router.push({path:'/profile'})
            })
            .catch(()=>{ // Возвращяет alert c ошибкой и очищает поля ввода
              this.alert = !this.alert
              this.loading = !this.loading
              this.password = ''
              this.email = ''
            })
      },
    }
  }
</script>

<style scoped>

</style>
