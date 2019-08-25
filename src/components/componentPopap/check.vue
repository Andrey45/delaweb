<template lang="pug">
  v-form(v-model="valid")
    // Ошибка
    v-alert(:value="alert" type="error")
      | Пользователь не зарегистрирован или введены некорректные данные
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
    v-card-actions.px-4
      // Input для подтверждения пароля
      v-text-field(
        v-model="password2"
        :append-icon="show1 ? 'visibility' : 'visibility_off'"
        :rules="[rules.required, rules.minLength, current]"
        :type="show1 ? 'text' : 'password'"
        label="Пароль"
        hint="Не менее 8 символов"
        @click:append="show1 = !show1")
    v-card-actions
      // Кнопка активна если все поля заполены, если нет то нажать невозожно
      // При клике вызывает функцию chec()
      v-btn( @click="chec()" outlined, :loading="loading", :disabled="!valid || loading") Зарегистрироваться
</template>

<script>
  import {mapActions} from 'vuex'
  export default {
    name: "check",
    data:()=>({
      valid: false,
      show1: false,
      rules:{
        // валидация input-ов
        required: value => !!value || 'Пожалуйста заполните!',
        minLength: value => value.length >= 8 || 'Пароль не менее 8 символов!',
        emaill: value => {
          const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          return pattern.test(value) || 'Некорректный E-mail.'
        }
      },
      email: '',
      password: '',
      password2: '',
      loading: false,
      alert: false
    }),
    computed:{
      current(){
        // Вычисляемое свойство проверки паролей на идентичность
        if(this.password === this.password2){ // Если пароли идентичны
          return true
        } else { // Если пароль различны
          return 'Пароли не совпадают'
        }
      }
    },
    methods:{
      ...mapActions(['check']),
      chec(){ // Функция регистрации
        this.check({email: this.email, password: this.password2}) // Вызывает actions 'check'
            .then(()=>{ // Если Promise вернул resolve то переводит на авторизацию
              this.loading = false
              this.$emit('window')
            })
            .catch(()=>{ // Возвращяет alert c ошибкой
              this.loading = false
              this.alert = true
            })
      }
    }
  }
</script>

<style scoped>

</style>
