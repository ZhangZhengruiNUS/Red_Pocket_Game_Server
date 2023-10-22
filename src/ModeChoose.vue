<script setup>
import {ref} from 'vue'
import router from "@/router";
import axios from "axios";

let userid = ref(router.currentRoute.value.params.userid)

//for catalog
let catalog_equipments = ref([])


function goGame() {
  router.push(`/Game/${userid.value}`)//how to transfer the id is next to solve
}


function goCatalog() {
  router.push(`/Catalog/${userid.value}`)//how to transfer the id is next to solve
}

function goWarehouse() {
  router.push(`/WareHouse/${userid.value}`);
}

function goStart() {
  router.push("/")
}

</script>

<template>
  <dev class="base">
    <h1 class="title">Welcome come Back!</h1>
    <h3 class="titleChoose">Choose your battle!</h3>

    <el-button-group class="choose">
      <button class="choosebutton" @click="goGame">start game</button>
      <button class="choosebutton" @click="goCatalog">go store</button>
      <button class="choosebutton" @click="goWarehouse">my warehouse</button>
      <!--      <a href="#/StartPage">hello</a>-->
      <button class="choosebutton" @click="goStart">log out</button>
    </el-button-group>
  </dev>

</template>

<style scoped>

.base {
  display: flex;
  flex-flow: column;
  justify-content: center;
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
  padding: 0 10px;
  margin: 20px;
}
</style>

//border: 1px solid #164bde; 
