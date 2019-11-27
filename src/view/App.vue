<template>
  <div id="app">
    <div class v-for="(image, index) in images" v-bind:key="index">
      <img src="image" alt="gopher" />
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "app",
  data: () => ({
    images: [],
    info: null
  }),
  mounted() {
    this.getGopher();
  },
  methods: {
    getGopher() {
      axios
        .get("http://localhost:8081/api/v1/gopher")
        .then(response => {
          this.images = response.data;
        })
        .catch(error => (this.info = error));
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
