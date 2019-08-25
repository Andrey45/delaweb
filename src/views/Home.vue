<template lang="pug">
  div
    // Слайдер с картинками
    v-carousel(cycle, hide-delimiter-background, show-arrows-on-hover)
      v-carousel-item(v-for='(item,i) in items', :key='i', :src='item.src')

</template>

<script>

import { mapGetters } from 'vuex'
export default {
  computed: mapGetters(['Login']),
  data:()=>({
    items:[
      {
        src: 'https://cdn.vuetifyjs.com/images/carousel/squirrel.jpg',
      },
      {
        src: 'https://cdn.vuetifyjs.com/images/carousel/sky.jpg',
      },
      {
        src: 'https://cdn.vuetifyjs.com/images/carousel/bird.jpg',
      },
      {
        src: 'https://cdn.vuetifyjs.com/images/carousel/planet.jpg',
      }
    ]
  })
};
</script>
