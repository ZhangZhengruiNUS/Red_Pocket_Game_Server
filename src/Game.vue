<template>
  <h1> Choose difficult Lv</h1>
  <div v-if="loading"> loading information </div>
  <div v-else>
    <n-carousel 
      class="diff"
      effect="card"
      prev-slide-style="transform: translateX(-150%) translateZ(-800px);"
      next-slide-style="transform: translateX(50%) translateZ(-800px);"
      style="height: 95%"
      :show-dots="false">
         
      <n-carousel-item :style="{ width: '60%' }" 
        v-for="(item, index) in responseData" :key="index"> 
        <img
          class="carousel-img"
          :src="picList[index]"
        >    
        <div class="difflv text">
          difficult level : {{item.diffLv}} 
        </div>
        <div class="award text">
          award density : {{item.awardDensity}}
        </div>
        <div class="enemy text">
          enemy density :{{ item.enemyDensity }}
        </div>
        <div class="car-button">
          <router-link class="linkbtn" 
            :to="{ path:'/GameEquipment/:difflv', 
            query: { difflv: JSON.stringify([userid,item.diffLV,item.awardDensity,item.enemyDensity]) }}"
          >GO!</router-link>
        </div>
      </n-carousel-item>      
    </n-carousel>
  </div>
</template>

<script >
import gameAPI from "./services/gameAPI.js";
import {onMounted,ref} from "vue";
import GameEquipment from "@/Gameview/Equipment.vue";
export default{
  data(){
    let equipt = "/GameEquipment"
    const loading = ref(true);
    const responseData = ref(null);
    const picList = [
      'https://64.media.tumblr.com/3f3fbe121d2ba8f314cb58e2ccf846bd/tumblr_n3ge65mRa11s335jfo1_r7_500.gifv',
      'https://i.pinimg.com/564x/e1/af/be/e1afbe9ebed88dfad68c9e621fabd78e.jpg',
      'https://i.pinimg.com/originals/14/ee/27/14ee27f1ee2970ffb44b1ea0eb156d7e.gif',
      "https://as1.ftcdn.net/v2/jpg/05/85/18/02/1000_F_585180251_GSr7HrbwPdxSMgW5TaOHJenVa26dBu12.jpg",
      'https://i.pinimg.com/originals/f8/b8/f0/f8b8f0e1f66ae0d6416b9e03e54214bd.jpg']
    
    const loadDiff  = async() => {
      try{
        const response = await gameAPI.getDiff()
        responseData.value = response.data
        console.log(responseData.value)
      } catch(err){
        console.log(err)
      } finally{
        loading.value=false
      }
    }  

    onMounted(() => {
      loadDiff();
    });
    
    return{
      responseData,
      picList,
      loading,
      equipt,
      userid: this.$route.params.userid,
      GameEquipment
    }
  }
}

</script>

<style scoped>
.diff {
  position: fixed;
  top: 100px;
}

.carousel-text {
  text-align: center;
  bottom: 0;
  font-size: large;
  z-index: 3;
}
.text{
  font-size: 26px;
  font-family: 'Trebuchet MS';
}
.difflv{
  color: hsl(300, 92%, 46%);
  text-align: center;
  left: 50%;
}
.award{
  color: hsl(145, 73%, 58%);
  text-align: center;
  left: 50%;
}
.enemy{
  color: hsl(246, 60%, 37%);
  text-align: center;
  left: 50%;
}

.carousel-img {
  margin: 0 auto;
  width: 100%;
  height: 60%;
  object-fit: cover;
  position: relative;
  z-index: 2;
}

.car-button {
  top: 90%;
  text-align: center;
  border: none;
  font-size: 30px;
  display: flex;
/* Horizontal centering */
  align-items: center;
}
.linkbtn{
  position: absolute;
  background-color: greenyellow;
  left: 50%;
  transform: translate(50%, 50%);
}
</style>
