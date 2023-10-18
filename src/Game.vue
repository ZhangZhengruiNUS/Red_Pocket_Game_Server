<script setup>
import {RouterView} from "vue-router";
import GameInterface from './Gameview/GameInterface.vue'
let LogOutPath = "/GameInterface"
let equipt = "/GameEquipment"

</script>

<template>

  <!--
    <Hero :x="1000" :y="1000" :height="60" :width="120" />

    <div class="container">
      <div class="cards-container">
        <div class="diffLV" v-for="info in diffList">
          <p> </p>
          <p> </p>
        </div>
      </div>
    </div>
  -->
  <!--
  <router-link class="choosebutton" to="/store">Choose equipment</router-link>
  <GameInterface/>
  -->
  <h1> Choose difficult Lv</h1>
  <n-carousel
      class = "diff"
      effect="card"
      prev-slide-style="transform: translateX(-150%) translateZ(-800px);"
      next-slide-style="transform: translateX(50%) translateZ(-800px);"
      style="height: 95%"
      :show-dots="false"
  >
    <n-carousel-item :style="{ width: '60%' }">
      <img
          class="carousel-img"
          src="./Gameview/background/bg1.png"
      >
      <div class="carousel-text">
        LV4
        <router-link class="car-button" :to="equipt">GO!</router-link>
      </div>
    </n-carousel-item>
    <n-carousel-item :style="{ width: '60%' }">
      <img
          class="carousel-img"
          src="./Gameview/background/bg2.jpg"
      >
      <div class="carousel-text">
        Difficult LV 3 asdfasd
        <router-link class="car-button" :to="equipt">GO!</router-link>
      </div>
    </n-carousel-item>
    <n-carousel-item :style="{ width: '60%' }">
      <img
          class="carousel-img"
          src="./Gameview/background/bg3.jpg"
      >

      <div class="carousel-text">
        This is some text.
        <router-link class="car-button" :to="equipt">GO!</router-link>

      </div>
    </n-carousel-item>

  </n-carousel>

</template>

<style scoped>
.diff{
  position: fixed;
  top: 100px;
}
.carousel-text{
  text-align: center;
  bottom: 0;
  font-size: large;
  z-index: 3;
}
.carousel-img {
  margin: 0 auto;
  width: 100%;
  height: 60%;
  object-fit: cover;
  position: relative;
  z-index: 2;
}

.car-button{
  background-color: greenyellow;
  border: none;
  font-size: larger;
  width: 80%;
  height: max-content;
  left: 0%;
}
</style>
