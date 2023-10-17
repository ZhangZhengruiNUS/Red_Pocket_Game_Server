<script setup>
import {ref} from "vue";
import {RouterLink} from "vue-router";

const Username = ref('')
const Password = ref('')
const errorMessage =ref("Sorry! The username and the password don't match!")

const emit = defineEmits(['SignInResponse'])
emit("SignInResponse", Username)
</script>

<script>
</script>

<template>
  <h2>Sign in</h2>
  <br>
  <h4>Continue your exciting fight fight!</h4>
  <br>
  <p> Username: <input v-model="Username" placeholder="Username"></p>
<!--  <p>{{Username}}</p>-->
  <p> Password: <input v-model="Password" placeholder="Password"></p>
  <br>
  <router-link :to="{path: '/ModeChoose/' + Username}">
    <button class="Signbutton" @click="">Sign in</button>
  </router-link>
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


