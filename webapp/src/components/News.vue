<template>
  <div>
    <h2 class="mx-5 my-5">GoNews - агрегатор новостей.</h2>
    <v-layout row wrap>
      <v-flex xs12 md6 lg3>
        <div v-for="post in news" :key="post.Id">

          <v-card elevation="10" outlined class="mx-5 my-5">
            <v-card-title>
              <a :href="post.Link" target="_blank"> {{ post.Title }} </a>
            </v-card-title>
            <v-card-text>
              {{ post.Description }}
              <v-card-subtitle>
                {{ post.Author }} {{ new Date(post.PubDate * 1000) }}
              </v-card-subtitle>
            </v-card-text>
          </v-card>

        </div>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
export default {
  name: "News-page",
  data() {
    return {
      news: [],
    };
  },
  mounted() {
    let url = "http://" + window.location.hostname + ":" + window.location.port + "/news/10";
    fetch(url)
      .then((response) => response.json())
      .then((data) => (this.news = data));
  },
};
</script>

<style scoped></style>