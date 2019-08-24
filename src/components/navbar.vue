<template lang="pug">
  v-app-bar(dark app)
    v-toolbar-title DeLaWeb
    div.flex-grow-1
    popap
</template>

<script>
  import Popap from "./popap";
  export default {
    name: "navbar",
    components: {Popap},
    data:()=>({
      open: false
    })
  }
</script>

<style scoped>

</style>
