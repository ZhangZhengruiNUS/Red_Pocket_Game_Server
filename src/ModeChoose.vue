<script setup>
import {ref} from 'vue'
import StartPage from './StartPage.vue'

const username = ref('')
let LogOutPath = ref("/StartPage")

let WarehousePath = ref("/WareHouse/")

</script>

<template>
  <dev class="base">
    <h1 class="title">Welcome come Back! {{ $route.params.username }}</h1>
    <h3 class="titleChoose">Choose your battle!</h3>

    <el-button-group class="choose">
      <router-link class="choosebutton" to="/Game">start game</router-link>
      <button class="choosebutton">go store</button>
      <router-link class="choosebutton" :to="{path: WarehousePath + $route.params.username}">my warehouse</router-link>
      <button class="choosebutton" @click="">turn table</button>
      <!--      <a href="#/StartPage">hello</a>-->
      <router-link class="choosebutton" :to="LogOutPath">log out</router-link>
    </el-button-group>
  </dev>

</template>

<style scoped>

.base {
  display: flex;
//flex-flow: column; justify-content: center;
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
//border: 1px solid #164bde; padding: 0 10px; margin: 20px;
}
</style>
