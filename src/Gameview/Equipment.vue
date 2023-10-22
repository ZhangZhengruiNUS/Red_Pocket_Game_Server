<template>
  <h1> choose your equipments</h1>
  <div v-if="loading"> loading page </div>
  <div v-else class="display" >
    <n-card :style="{ width: '60%'}" 
        v-for="(item, index) in responseData" :key="index"> 
        <template #cover>
          <img :src=item.picPath>
        </template>
        <div class="cardContent">
          <h3>{{ item.itemName }}</h3>
          <div class="skill_describe">
            {{ item.describe }}
          </div>
          <div class="quantity"> You have: <div class="number">{{ item.quantity }} </div></div>
          <n-checkbox size="small" @update:checked="carryEquip(item.itemId)" label="带上装备" />
        </div>
    </n-card>
  </div>
  <div class="buttonContainer">   
    <router-link class="linkbtn" @click="postEquip"
            :to="{ path:'/GameInterface/:gameinfo', 
            query: { gameinfo: JSON.stringify( linkinfoList.join(',') )}}"
      >GO!</router-link>
  </div>
</template>

<script >
import gameAPI from "../services/gameAPI.js";
import {onMounted,ref} from "vue";
import { useRoute } from 'vue-router'

export default{
  data(){
    const route = useRoute()
    const loading = ref(true)
    let userId = ref(0)
    let equipmentsList = ref([])
    const responseData = ref()
    const difflv = JSON.parse(route.query.difflv)
    userId = parseInt(difflv[0], 10); 
    const diff = parseInt(difflv[1], 10); 
    const awardDensity = parseInt(difflv[2], 10); 
    const enemyDensity = parseInt(difflv[3], 10); 
    let carryEquipList=ref([])
    let reflinkinfoList =ref([userId,diff,awardDensity,enemyDensity,])
    
    const loadEquipment  = async() => {
      try{
        const response = await gameAPI.getEquip(userId)
        responseData.value = response.data
      } catch(err){
        console.log(err)
      } finally{
        loading.value=false
      }
    } 
    onMounted(() => {
      loadEquipment();
    });
    

    const postEquip = async() =>{
      try{
        const response = await gameAPI.postEquip(userId, carryEquipList.value)     
        console.log("jump now")
        console.log(reflinkinfoList.value)
      }catch(err){
        console.log(err)
      }
    }

    const carryEquip = (itemId) => {
      reflinkinfoList.value.push(itemId)
      carryEquipList.value.push(itemId)
    }
    return{
      equipmentsList,
      loading,
      carryEquip,userId,
      carryEquipList,
      responseData,
      postEquip,
      linkinfoList:reflinkinfoList.value
    } 
  }
}
</script>

<style scoped>

.display{
  display: flex;
  justify-content: space-around;
}
.n-card {
  max-width: 300px;
}

.number{
  font-size: 44px;
  font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
.buttonContainer{
  position: absolute;
  left: 10%;
  height: 60px;
  width: 80%;
  background: #4df719;
  align-items: center;
}
.linkBtn{
  width: 80%;
  position: absolute;
  left: 40%;
  line-height: 36px;
  font-weight: bold;
  font-size: 40px;
  color: #2c3e50;
  border: 1px solid transparent;
  padding: 0 10px;
  margin: 20px;
}

</style>