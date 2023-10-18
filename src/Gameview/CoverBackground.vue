
<script>
import { computed, reactive, ref ,toRefs,onMounted} from 'vue';
export default {
    name: 'CoverBackground',
    components:{
        
    },
    props: {
        life: Number,
        reStartGame:Function
    },
    setup(props) {
        onMounted(() => {
        // You can access the props here too if needed
            console.log("life is " + props.life);
        });
        return{
            //airplaneList
        }
    }
};
</script>


<template>
    <div  class="cover-bg">
        <div @click="reStartGame()">
            Restart Game 
        </div>
    </div>
</template>

<style scoped>
    .cover-bg {
        position: absolute;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, .65);
        z-index: 2000;
    }

    .cover-bg>div {
        position: absolute;
        left: 50%;
        top: 50%;
        width: 200px;
        height: 60px;
        text-align: center;
        transform: translate(-50%, -50%);
        line-height: 60px;
        background: #5eb662;
        box-shadow: 0px 0px 8px #5eb662;
        border-radius: 8px;
        color: #fff;
        font-size: 28px;
    }
</style>