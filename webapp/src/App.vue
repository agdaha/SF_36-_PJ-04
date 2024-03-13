<template>
  <v-app>
    <v-main>
      <News-page/>
    </v-main>
  </v-app>
</template>

<script>

import NewsPage from "./components/News.vue"
export default {
  name: 'App',

  components: {
    NewsPage
  },

  data: () => ({
    //
  }),
}
</script>
