<script setup>
import {ref} from 'vue'
import SignIn from './components/SignIn.vue'
import SignUp from './components/SignUp.vue'

const picked = ref('Sign in')
let username = ref("username")
let password = ref("password")

function changePicked(choose) {
  picked.value = choose;
}

</script>

<template>
  <dev class="Startbase">

    <h1 class="Starttitle">FlyFight</h1>
    <br>
    <button-group class="Startchoose" type="text">
      <button class="Startchoosebutton" v-on:click="changePicked('Sign in')">Sign in</button>
      <button class="Startchoosebutton" v-on:click="changePicked('Sign up')">Sign up</button>
    </button-group>
    <br>
    <sub class="Startsignpage">
      <SignIn @SignInResponse="(Username) => username = Username" v-if="picked == 'Sign in'"/>
      <SignUp @SignUpResponse="(Username) => username = Username" v-if="picked == 'Sign up'"/>
    </sub>

<!--    <p @SignUpResponse="(Username) => username = Username"></p>-->
<!--    <p @SignInResponse="(Username) => username = Username"></p>-->
    <p>{{username}}</p>

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
