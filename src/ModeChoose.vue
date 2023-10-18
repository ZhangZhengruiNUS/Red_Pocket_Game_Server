<script setup>
import {ref} from 'vue'
import router from "@/router";

const username = router.currentRoute.value.params.username

function goGame(){
  router.push("/Game")
}
function goStore(){
  router.push("/catalog")
}
function goWarehouse(){
  router.push(`/WareHouse/${username}`)
}
function goStart(){
  router.push("/")
}

</script>

<template>
  <dev class="base">
    <h1 class="title">Welcome come Back! {{ username }}</h1>
    <h3 class="titleChoose">Choose your battle!</h3>

    <el-button-group class="choose">
      <button class="choosebutton" @click="goGame">start game</button>
      <button class="choosebutton" @click="goStore">go store</button>
      <button class="choosebutton" @click="goWarehouse">my warehouse</button>
      <!--      <a href="#/StartPage">hello</a>-->
      <button class="choosebutton" @click="goStart">log out</button>
    </el-button-group>
  </dev>

</template>

<style scoped>

.base {
  display: flex;
  flex-flow: column; justify-content: center;

  text-align: center;
}

.title {
  font-weight: bold;
  font-size: 60px;
  position: absolute;
  top: 5%;
}

.choose {

  position: absolute;
  top: 32%;
  display: flex;

  flex-flow: column;
}

.choosebutton {

  height: 60px;
  width: 400px;
  line-height: 36px;
  font-weight: bold;
  font-size: 40px;
  background: #accbea;
  color: #2c3e50;
  border: 1px solid transparent;
  padding: 0 10px; margin: 20px;
}
</style>
