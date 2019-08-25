<template>
  <v-app>
    <navbar/>
    <v-content style="background-color: #424242">
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>


import Navbar from "./components/navbar";
export default {
  name: 'App',
  components: {
    Navbar,
  },
  data: () => ({
    //
  }),
};
</script>
