<script setup>
import {ref} from "vue";
import router from "@/router";

const Username = ref('')
const Password = ref('')
const errorMessage =ref("Sorry! The username and the password don't match!")

function log(){
  router.push(`/ModeChoose/${Username.value}`)
}
</script>

<template>
  <h2>Log in</h2>
  <br>
  <h4>Continue your exciting fight fight!</h4>
  <br>
  <p> Username: <input v-model="Username" placeholder="Username"></p>
<!--  <p>{{Username}}</p>-->
  <p> Password: <input type="password" v-model="Password" placeholder="Password"></p>
  <br>
<!--  <router-link :to="{path: '/ModeChoose/' + Username}">-->
    <button class="Signbutton" @click="log" is="false">Log in</button>
<!--  </router-link>-->
  <button class="Signbutton" @click="">Quit</button>
</template>

<style>
.Signbutton{
  height: 35px;
  width: 150px;
  line-height: 23px;
  font-weight: bold;
  font-size: 22px;
  background: #accbea;
  color: #2c3e50;
  border: 1px solid transparent;
//border: 1px solid #164bde;
  padding: 0 10px;
  margin: 2px;
}
</style>


