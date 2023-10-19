<template>
  <n-card title="带封面的卡片">
    <template #cover>
      <img
          src="https://static.vecteezy.com/system/resources/previews/015/281/932/original/cannon-icon-cartoon-style-vector.jpg">
    </template>
    技能描述
    <n-checkbox size="small" @update:checked="clock" label="带上装备" />
  </n-card>

  <n-card title="带封面的卡片">
    <template #cover>
      <img
          src="https://s3-eu-central-1.amazonaws.com/cartoons-s3/styles/product_detail_image/s3/tsar%20cannon.jpg?itok=GxjsPcuQ">
    </template>
    技能描述
    <n-checkbox size="small" @change="clock" label="带上装备" />
  </n-card>

  <p>{{equipments}}</p>

  <div v-for="equipment in equipments">
    <n-card :title="equipment.itemName">
      <template #cover>
        <img :src="equipment.picPath"/>
      </template>
      {{ equipment.describe }}
      <n-checkbox size="small" label="带上装备" />
    </n-card>
  </div>

  <router-link class="router_button" to="/GameInterface">GO!</router-link>
</template>

<script setup>
import axios from "axios";
import {ref} from "vue";

let GameInterface = "/GameInterface"
let userId = ref(0)
let equipments = ref([])

function clock(){
  console.log("checked")
}

axios.get("http://localhost:5173/#/", {
  params: {
    userId: userId.value
  }
}).then(function (response) {
  let equipt;
  for (equipt in response.data) {
    equipments.value.push({
      itemId: equipt.itemId,
      itemName: equipt.itemName,
      describe: equipt.describe,
      quantity: equipt.quantity,
      picPath: equipt.picPath
    });
  }
}).catch(error => console.log(error));

</script>

<style scoped>
.n-card {
  max-width: 300px;
}

.router_button {
  height: 60px;
  width: 400px;
  line-height: 36px;
  font-weight: bold;
  font-size: 40px;
  background: #4df719;
  color: #2c3e50;
  border: 1px solid transparent;
  padding: 0 10px;
  margin: 20px;
}

</style>