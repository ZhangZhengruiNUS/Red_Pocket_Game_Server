<script>
import { computed, reactive, ref ,toRefs,onMounted} from 'vue';
export default {
    name: 'BulletObj',
    props: {
        bulletList: Array,
    },
    setup(props) {
        onMounted(() => {
        // You can access the props here too if needed
            console.log("try to get here");
            console.log(props.bulletList.length);
        });
        return{
            //airplaneList
        }
    }
};
</script>


<template>
    <div class="bullet fly-obj" v-for="item in bulletList" :style="{transform:`translate(${item.x}px, ${item.y}px)`,width:`${item.width}px`,
            height:`${item.height}px`}">

    </div>
</template>

<style scoped>
    .bullet {
        background: #256945;
        border-radius: 50%;
    }
</style>