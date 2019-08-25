<template lang="pug">
  div
    // Кнопка открытия модального окна
    v-btn(@click="open = !open") Изменить
    v-dialog(v-model="open" max-width="400" dark)
      v-card(v-for="item in UserData")
        v-card-title Редактирование
        v-card-actions
          // Input логина
          v-text-field(v-model="item.username" label="Логин")
        v-card-actions
          // Select с возможными вариантами
          v-combobox(v-model="item.country" :items="itemCountry" label="Страна")
        v-card-actions
          // Select с возможными вариантами
          v-combobox(v-model="item.city" :items="itemsCity" label="Город")
        v-card-actions
          // Кнопка не активна при выполнении функции
          // При клике вызываеться функция update в которую передаються динамически изменненые данные о пользователи
          v-btn(@click="update(item)" :loading="loading", :disabled="loading") Изменить

</template>

<script>
  import { mapActions, mapGetters } from 'vuex'
  export default {
    name: "popapUpdate",
    computed: mapGetters(['UserData']),
    data:()=>({
      open:false,
      loading: false,
      itemsCity:[
        'Волгоград','Москва','Санкт-Петербург'
      ],
      itemCountry:[
          'Россия','Казахстан','Белоруссия'
      ]
    }),
    methods:{
      ...mapActions(['updeteUser']),
      update(item){ // Изменение данных пользователя
        this.loading = !this.loading
        this.updeteUser(item) // Вызавет actions updeteUser
            .then(()=> { // Если вернулся результат закрывает окно
                this.open = !this.open
                this.loading = !this.loading
            })
            .catch(()=> this.loading = !this.loading)
      }
    }
  }
</script>

<style scoped>

</style>
