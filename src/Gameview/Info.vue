<script>
import {onMounted} from 'vue';
export default {
    name: 'Info',
    props: {
        hero: Object,
        score: Number,
        difficultLv: Number ,
        coupon:Number,
        power:Number, 
        pause:Boolean
    },
    setup(props) {
        onMounted(() => {
        // You can access the props here too if needed
        // console.log("try to get here");
        // console.log(props.hero);
        });
        return{
            //airplaneList
        }
    }
};
</script>

<template>
    <div class="hero-info">
                <div>Life: 
                    <progress :value="hero.life" max="100"></progress>
                </div>
                <div >HotFire Level: {{hero.shootLv}}</div>
                <div>Bullet Power:Lv{{hero.shootSpeed}}</div>
                <div>Bullet Speed:Lv{{hero.shootPower}}</div>
                <div>射击模式：{{hero.shootType===1?'集':hero.shootType===2?'散':'乱'}}</div>
            </div>

            <div class="score-info">
                <div>得分：{{score}}</div>
                <div>难度：{{difficultLv}}</div>
                <div>coupon: {{ coupon }}</div>
            </div>

            <div class="power-info">
                <div class="title">power</div>
                <div class="con">
                    <div :style="{width:`${power}%`}" class="bg">

                    </div>
                </div>
            </div>

            <div class="oprate-info">
                <div style="font-size:12px;">S键开始<br />P键暂停<br />鼠标左键切换弹道<br />空格键绝杀</div>
                <div class="container" :class="{conActive:pause}">
                    <div class="inner" :class="{innActive:pause}">
                        {{pause?'停':''}}
                    </div>
                </div>
            </div>

</template>

<style scoped>
    .hero-info {
        position: absolute;
        top: 5px;
        right: 5px;
        z-index: 1000;
        background: #ddd;
        padding: 8px;
        border-radius: 8px;
        font-size:18px;
    }

    .score-info {
        position: absolute;
        top: 5px;
        left: 5px;
        z-index: 1000;
        background: #ddd;
        padding: 8px;
        border-radius: 8px;
        font-size:18px;
    }

    .power-info {
        position: absolute;
        bottom: 5px;
        right: 5px;
        z-index: 1000;
        background: #ddd;
        padding: 8px;
        border-radius: 8px;
        display: flex;
        align-items: center;

    }

    .power-info .con {
        height: 20px;
        width: 120px;
        box-sizing: border-box;
        border: 2px solid #000;
        margin-left: 4px;

    }

    .power-info .con .bg {
        width: 100%;
        height: 100%;
        background-image: linear-gradient(to right, #66f, #f66);
        background-size: 116px 100%;
    }

    .oprate-info {
        position: absolute;
        bottom: 5px;
        left: 5px;
        z-index: 1000;
        background: #ddd;
        padding: 8px;
        border-radius: 8px;
    }

    .oprate-info .container {
        height: 30px;
        width: 60px;
        background-color: #4fc08d;
        border-radius: 25px;
        position: relative;
        border: 2px solid #666;
        box-sizing: border-box;
    }

    .oprate-info .inner {


        position: absolute;
        height: 100%;
        width: 50%;
        background-color: #fff;
        border-radius: 50%;
        text-align: center;
        font-size: 12px;
        line-height: 25px;
    }


    .oprate-info .conActive {
        background-color: #d14;
    }

    .oprate-info .innActive {
        left: 50%;
    }

    .oprate-info .conActive {
        background-color: #d14;
    }

    .oprate-info .innActive {
        left: 50%;
    }
</style>