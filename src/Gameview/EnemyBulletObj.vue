<script>
import { onMounted} from 'vue';
export default {
    name: 'EnemyBulletObj',
    props: {
        enemyBulletList: Array,
    },
    setup(props) {
        onMounted(() => {
        // You can access the props here too if needed
        // console.log("try to get here");
        // console.log(props.enemyBulletList.length);
        });
        return{
            //airplaneList
        }
    }
};
</script>

<template>
    <div class="enemy-bullet fly-obj" v-for="item in enemyBulletList" :style="{transform:`translate(${item.x}px, ${item.y}px)`,width:`${item.width}px`,
        height:`${item.height}px`}">
    </div>
</template>

<style scoped>
    .enemy-bullet {
        background: #bb250d;
        border-radius: 50%;
    }
</style>