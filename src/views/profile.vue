<template lang="pug">
 v-parallax(src="https://cdn.vuetifyjs.com/images/parallax/material.jpg")
   v-container(style="width: max-content")
     v-card(v-for="prop in UserData" max-width="900")
        v-card-title
          // Если пользоваетль уже имеит логин
          v-container(v-if="prop.username !== ''") Логин: {{prop.username}}
          // Если ользователь не заполнил или удалил логин
          v-container(v-else) Логин: Незаполнен
        v-card-text
          // Данные почты
          v-container(v-if="prop.email !== ''") E-mail: {{prop.email}}
        v-card-text
          // Если пользоваетль уже указал страну
          v-container(v-if="prop.country !== ''") Страна: {{prop.country}}
          // Если не указал страну
          v-container(v-else)  Страна: Не указана
        v-card-text
          // Если пользоваетль уже указал город
          v-container(v-if="prop.city !== ''") Город: {{prop.city}}
          // Если не указал город
          v-container(v-else) Город: Не указан
     v-container(style="width: max-content")
        // Попап изменения данных
        popap-update

</template>
<script>
  import { mapGetters } from "vuex"
  import PopapUpdate from "../components/popapUpdate";
  export default {
    components: {PopapUpdate},
    computed: mapGetters(['UserData']),
  }
</script>
