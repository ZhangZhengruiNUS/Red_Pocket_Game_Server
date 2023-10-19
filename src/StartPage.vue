<script setup>
import {ref} from 'vue'
import LogIn from './components/LogIn.vue'
import SignUp from './components/SignUp.vue'

// import axios from "axios";
// import router from "@/router";
//
// axios.get('https://jsonplaceholder.typicode.com/posts/1',)
//     .then(function (response){
//       console.log(response);
//       console.log("ok");
//       console.log(response.data)
//       console.log("id: " + response.data.id)
//
//     })
//     .catch(function (error){
//       console.log(error);
//       console.log("error");
//       })

const picked = ref('Log In')

function changePicked(choose) {
  picked.value = choose;
}

</script>

<template>
  <dev class="Startbase">

    <h1 class="Starttitle">FlyFight</h1>
    <br>
    <button-group class="Startchoose" type="text">
      <button class="Startchoosebutton" v-on:click="changePicked('Log In')">Log In</button>
      <button class="Startchoosebutton" v-on:click="changePicked('Sign Up')">Sign up</button>
    </button-group>
    <br>
    <sub class="Startsignpage">
      <LogIn  v-if="picked == 'Log In'"/>
      <SignUp v-if="picked == 'Sign Up'"/>
    </sub>

<!--    <p @SignUpResponse="(Username) => username = Username"></p>-->
<!--    <p @SignInResponse="(Username) => username = Username"></p>-->

  </dev>

</template>

<style scoped>

.Startbase {
  display: flex;
  justify-content: center;
  text-align: center;
}

.Starttitle {
  font-weight: bold;
  font-size: 75px;
  position: absolute;
  top: 3%;
}

.Startchoose {
  position: absolute;
  top: 32%;
}

.Startchoosebutton {
  height: 60px;
  width: 400px;
  line-height: 36px;
  font-weight: bold;
  font-size: 40px;
  background: #accbea;
  color: #2c3e50;
  border: 1px solid transparent;
//border: 1px solid #164bde; padding: 0 10px;
  margin: 2px;
}

.Startsignpage {
  position: absolute;
  top: 40%;
  font-size: 20px;
}

</style>
