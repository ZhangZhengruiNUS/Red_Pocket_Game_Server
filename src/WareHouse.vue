<script setup>
import {ref} from 'vue'
import RollTable from './components/RollTable.vue';
import axios from 'axios'
let userId = ref(0)
let equipments = ref([])
let coupons = ref(-1)
let credit = ref(-1)
let credits = ref(0)

let startroll = ref(false)

axios.get('http://localhost:5173/#/ModeChoose?username=Username', {
  params: {
    userId: userId.value,
  }
}).then(function (response) {
  coupons.value = response.data.coupons;
  credit.value = response.data.credit;
  let equipt;
  for (equipt in response.data) {
    equipments.value.push({
      itemId: equipt.itemId,
      itemName: equipt.itemName,
      describe: equipt.describe,
      quantity: equipt.quantity,
      price: equipt.price,
      picPath: equipt.picPath
    });
  }
}).catch(error => console.log(error))

function getRandomColor() {
  return "hsl(" + Math.random() * 360 + "), 100%, 75%";
}


function turntable() {
//   about the jumping out turntable
  if (coupons.value == 0) {
    startroll.value = true
    couponWarning.value = "You don't have any coupons, you can get one by gaming!";
    return
  } else {
    startroll.value = true
    couponsNumber.value = couponsNumber.value - 1
    couponWarning.value = ''
  }
}
</script>

<template>
  <div v-if="startroll" class="overlay">
    <RollTable
      :prizeList="[
        { name: '手机', src: 'https://www.apple.com/newsroom/images/product/iphone/geo/Apple-iPhone-14-iPhone-14-Plus-hero-220907-geo.jpg.og.jpg?202308290218' },
        { name: '手表', src: 'https://img1.baidu.com/it/u=2631716577,1296460670&fm=253&fmt=auto&app=120&f=JPEG' },
        { name: '苹果', src: 'https://img2.baidu.com/it/u=2611478896,137965957&fm=253&fmt=auto&app=138&f=JPEG' },
        { name: '棒棒糖', src: 'https://img2.baidu.com/it/u=576980037,1655121105&fm=253&fmt=auto&app=138&f=PNG' },
        { name: '娃娃', src: 'https://img2.baidu.com/it/u=4075390137,3967712457&fm=253&fmt=auto&app=138&f=PNG' },
        { name: '木马', src: 'https://img1.baidu.com/it/u=2434318933,2727681086&fm=253&fmt=auto&app=120&f=JPEG' },
        { name: '德芙', src: 'https://img0.baidu.com/it/u=1378564582,2397555841&fm=253&fmt=auto&app=120&f=JPEG' },
        { name: '玫瑰', src: 'https://img1.baidu.com/it/u=1125656938,422247900&fm=253&fmt=auto&app=120&f=JPEG' }
        ]"
    />
    <button @click="startroll=false">close</button>

  </div>
  
  <div>
    <h1>Hi! Welcome to your warehouse!</h1>

    <div>
      <h3>Your credits are: {{ credits }}</h3>
      <br>
      <h3>You have {{ coupons }} coupons!</h3>
      <button @click="turntable">roll the turntable!</button>
      <p v-if="couponWarning" class="ErrorMsg">{{ coupons }}</p>
      <p v-else>{{ }}</p>
    </div>

    <div class="equipmentContainer">
      <!--      for the card in the equipments     -->
      <div v-for="equipment in equipments">
        <p class="equipment" :id="equipment.itemId">
          <p>picture</p>
          <p>{{ equipment.itemName }}</p>
          <p>{{ equipment.describe }}</p>
          <p>{{ equipment.quantity }}</p>
        </p>
      </div>
    </div>

    <div>
      <router-link :to="{path: '/ModeChoose/' + $route.params.username}">back to menu</router-link>
    </div>

  </div>
</template>

<style scoped>
.overlay {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.77);
  z-index: 10;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.equipmentContainer {
  max-width: 1000px; 
  padding: 10px; 
  margin: 0 auto; display: flex; flex-wrap: wrap;
}

.equipment {
  width: 225px;
  height: 225px;
  background-color: #98d3fc;
  padding: 10px;
  border-radius: 15px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  margin-right: 20px;
  margin-bottom: 20px;
}

.jump {
  width: 750px;
  background-color: cyan;
  border-radius: 10px;
  padding: 30px;
  position: relative;
  display: flex;
  flex-direction: column;
}

.ErrorMsg {
  color: red;
}
</style>