<template>
  <div id="app">
    <h1 class="title">Let's vote your favorite Gopher!!</h1>
    <div class="imgWrap">
      <div class="imgdom" v-for="(item, index) in imageData" v-bind:key="index">
        <img :src="'img/'+item.Img" alt="gopher" class="image" />
        <div class="name">{{item.Name}}</div>
        <div class="likeWrap">
          <div>
            <img src="img/good.png" alt="good" class="good" />
          </div>
          <div class="balloon">12</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "app",
  data: () => ({
    imageData: [],
    info: null
  }),
  mounted() {
    this.getGopher();
  },
  methods: {
    getGopher() {
      axios
        .get("http://localhost:8081/api/v1/gophers")
        .then(response => {
          // this.images = response.gophers.map(item => item.img);
          this.imageData = response.data.gophers;
          // this.imageData = response.data.gophers.map(item => "img/" + item.Img);
          // this.names = response.data.gophers.map(item => item.Name);
        })
        .catch(error => (this.info = error));
    }
  }
};
</script>

<style>
body {
  background: #7fd5ea;
  height: 100%;
  margin: 0;
  padding: 0;
}
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  width: 90%;
  height: 100vh;
  margin: 10px auto;
  background: #fff;
}
.title {
  padding-top: 30px;
}
.imgWrap {
  margin: 0 auto;
  padding: 5px 30px;
  background-color: #fff;
  display: flex;
  flex-wrap: wrap;
}
.imgdom {
  margin-right: 10px;
  display: flex;
  flex-direction: column;
}
.image {
  width: 150px;
  height: 150px;
}
.name {
  margin: 10px 0 0;
  font-size: 20px;
}
.likeWrap {
  display: flex;
  margin: 10px auto;
}
.good {
  width: 33px;
  height: 33px;
}
.balloon {
  margin-top: 3px;
  margin-left: 6px;
  position: relative;
  display: inline-block;
  width: 32px;
  height: 32px;
  line-height: 32px;
  vertical-align: middle;
  text-align: center;
  color: rgb(55, 55, 55);
  font-size: 11px;
  font-family: arial, sans-serif;
  font-weight: bold;
  box-sizing: border-box;
  margin: 5px 0px 10px 15px;
  padding: 0px 5px;
  background: rgb(162, 236, 255);
  border-radius: 50%;
}
.balloon:before {
  content: "";
  position: absolute;
  top: 80%;
  left: -9px;
  margin-top: -15px;
  border: 5px solid transparent;
  border-right: 5px solid #a2ecff;
  z-index: 0;
}
</style>
