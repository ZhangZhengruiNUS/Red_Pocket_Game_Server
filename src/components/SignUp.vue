<script setup>
import {ref} from "vue";

const Username = ref('')
const Password1 = ref('')
const Password2 = ref('')
const unmerrorMessage =ref("Sorry! The username have been taken!")
const pwderrorMessage =ref("Sorry! The password don't match!")
const emit = defineEmits(['SignUpResponse'])
emit("SignUpResponse", Username)
</script>

<script>
</script>

<template>
  <h2>Sign up</h2>
  <br>
  <h4>Start your exciting fight journey by signing up!</h4>
  <br>
  <p> Username: <input v-model="Username" placeholder="Username"></p>
  <p> Input your Password: <input v-model="Password1" placeholder="Password"></p>
  <p> Config your Password: <input v-model="Password2" placeholder="Password"></p>
  <p class="ErrorMsg" v-if="Password1!=Password2">{{pwderrorMessage}}</p>
  <br>
  <button class="Signbutton" @click="">Sign up</button>
  <button class="Signbutton" @click="">Quit</button>
</template>

<style scoped>
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
.ErrorMsg{
  color:red;
}
</style>