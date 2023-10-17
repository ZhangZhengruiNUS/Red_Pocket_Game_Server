<script>
  import { ref, computed, onMounted,watch } from 'vue';
  import Hero from '../class/HeroClass.js'
  import ShootAirplane from '../class/ShootAirplane.js'
  import Airplane from '../class/Airplane.js'
  import FastAirplane from '../class/FastAirplane.js'
  import EnemyBullet from '../class/EnemyBullet.js'
  import Bullet from '../class/Bullet.js'
  import Award from '../class/Award.js'
  import MaxFireBullet from '../class/MaxFireBullet.js'

  import HeroView from './HeroView.vue'
  import AirplaneObj from './AirplaneObj.vue';
  import AwardObj from './AwardObj.vue'
  import BulletObj from './BulletObj.vue'
  import EnemyBulletObj from './EnemyBulletObj.vue'
  import MaxFireObj from './MaxFireObj.vue'
  import Info from './Info.vue'
  import CoverBackground from './CoverBackground.vue'


  export default {
    name: 'GameInterface',
    components: {
        HeroView,
        AirplaneObj,
        AwardObj,
        BulletObj,
        EnemyBulletObj,
        MaxFireObj,
        Info,
        CoverBackground,
        CoverBackground
    },
    props: {
      // Your props here
    },
    
    setup(props){
        const airplaneList = ref([]); // Define your data properties as refs
        const pushAirplaneTimer = ref(null);
        const airplaneTimer = ref(null);

        
        const awardList = ref([]);
        const pushAwardTimer = ref(null);
        const awardTimer = ref(null);
        
        const bulletList = ref([]);
        const pushBulletTimer = ref(null);
        const bulletTimer = ref(null);
        
        const enemyBulletList = ref([]);
        const enemyBulletTimer = ref(null);
        
        const maxFireBulletTimer = ref(null);
        const maxFireBulletList = ref([]);
       
        const score = ref(0);
        const power = ref(100);
        const pause = ref(false);        
        const hero = ref(new Hero( 60, 120, window.innerWidth / 2, window.innerHeight / 2));
        // 计算
        
        const difficultLv = computed(() => {
            if (score.value > 1600) {
            return 5;
            } else if (score.value <= 1600 && score.value > 800) {
            return 4;
            } else if (score.value <= 800 && score.value > 300) {
            return 3;
            } else if (score.value <= 300 && score.value > 100) {
            return 2;
            } else {
            return 1;
            }
        });
        const shootSpeed = computed(() => hero.shootSpeed || 1);
        const life = computed(() => hero.value.life || 1);
        
        
        // watch :
        watch(difficultLv, (val) => {
            resetPushAirPlaneTimer(val);
        });
  
        watch(shootSpeed, (val) => {
            resetPushBulletTimer(val);
        });
  
        watch(pause, (val) => {
            if (!val) {
            startTimer();
            } else {
            pauseTimer();
            }
        });
  
        watch(life, (val) => {
            if (val <= 0) {
            // Game over
            console.log('Game over');
            pauseTimer();
            }
        });
        

        onMounted(() => {
            window.addEventListener('mousemove', heroMove)
            window.addEventListener('click', changeShootType);
            window.addEventListener('keyup', keyUp);
            startTimer();
        });


        //具体method:
        const changeShootType = () => {
            console.log('火力改变');
            hero.value.changeShootType();
        };

        const reStartGame = () =>{
            airplaneList.value.splice(0);
            awardList.value.splice(0);
            bulletList.value.splice(0);
            enemyBulletList.value.splice(0);
            maxFireBulletList.value.splice(0);
            score.value=0;
            power.value=100;
            pause.value=false;
            hero.value.reset(60,120);
            startTimer()
        }

        const heroMove = (e) => {
            if(pause.value){
                return;
            }
            hero.value.moveTo(e.pageX,e.pageY);
        }
        
        const maxFire = () => {
            console.log('绝杀！！！');
        };
        
        const keyUp = (event) => {
            console.log(event);
            if (event.code === 'KeyP') {
                console.log('暂停');
                pause.value = true;
            }
            
            if (event.code === 'KeyS') {
                pause.value = false;
            }
            
            if (event.code === 'Space' && power.value === 100) {
                pushMaxFire();
                power.value = 0;
            }
        };
        
      
        const resetPushBulletTimer = (shootSpeed) => {
            clearInterval(pushBulletTimer.value);
            pushBulletTimer.value = setInterval(() => {
            const arr = hero.shoot() || [];
    
            arr.forEach((it) => {
                const idx = bulletList.value.findIndex((item) => item.isEnd());
                if (idx > -1) {
                bulletList.value[idx].reset(it.w * hero.shootPower, it.h * hero.shootPower, it.x, it.y, it.speed, it.xSpeed);
                } else {
                bulletList.value.push(new Bullet(it.w, it.h, it.x, it.y, it.speed, it.xSpeed));
                }
            });
            }, 500 / shootSpeed);
        };

        const resetPushAirPlaneTimer = (difficultLv) => {
            clearInterval(pushAirplaneTimer.value)
            pushAirplaneTimer.value = setInterval(() => {
                let airRandom = Math.floor(Math.random() * 3)
                let oriX = Math.floor(Math.random() * window
                    .innerWidth);

                let x = oriX < 20 ? 20 : oriX > window.innerWidth - 20 ? window.innerWidth -
                    80 :
                    oriX;


                if (airRandom < 1) {
                    let idx = airplaneList.value.findIndex(it => it.type === "ShootAirplane" &&
                        it
                        .run == false)
                    if (idx > -1) {
                        airplaneList.value[idx].start();
                    } else {
                        airplaneList.value.push(new ShootAirplane(60, 60, x, 0))
                    }

                } else {
                    if (airRandom < 2) {
                        let idx = airplaneList.value.findIndex(it => it.type ===
                            "FastAirplane" && it
                            .run == false);

                        if (idx > -1) {
                            airplaneList.value[idx].start();
                        } else {
                            airplaneList.value.push(new FastAirplane(30, 60, x, 0))
                        }

                    } else {


                        let idx = airplaneList.value.findIndex(it => it.type === "Airplane" && it
                            .run == false);

                        if (idx > -1) {
                            airplaneList.value[idx].start();
                        } else {
                            airplaneList.value.push(new Airplane(30, 60, x, 0))
                        }
                    }

                }
            }, 600 / difficultLv)
        }
      
        // 三个功能性 function 
  
        const pushMaxFire = () => {
            for (let i = 0; i < 50; i++) {
                const idx = maxFireBulletList.value.findIndex(item => item.isEnd());
                const x = Math.floor(Math.random() * window.innerWidth);

                if (idx > -1) {
                maxFireBulletList.value[idx].reset(40, 40, x, 200 * Math.random() + window.innerHeight, 40, 40);
                } else {
                maxFireBulletList.value.push(new MaxFireBullet(40, 40, x, 200 * Math.random() + window.innerHeight, 40, 40));
                }
            }
        };

        const pauseTimer = () => {
            clearInterval(pushAirplaneTimer.value);
            clearInterval(pushAwardTimer.value);
            clearInterval(pushBulletTimer.value);
            clearInterval(bulletTimer.value);
            
            clearInterval(airplaneTimer.value);
            clearInterval(awardTimer.value);
            console.log("here we go ")
            clearInterval(enemyBulletTimer.value);
        };

        const startTimer = () => {
            //飞机导入的timer
            pushAirplaneTimer.value = setInterval(() => {
                let airRandom = Math.floor(Math.random() * 4)
                let oriX = Math.floor(Math.random() * window.innerWidth);

                let x = oriX < 20 ? 20 : oriX > window.innerWidth - 20 ? window.innerWidth -80 :oriX;

                if (airRandom < 1) {
                    let idx = airplaneList.value.findIndex(it => it.type === "ShootAirplane" &&
                        it
                        .run == false)
                    if (idx > -1) {
                        airplaneList.value[idx].start();
                    } else {
                        airplaneList.value.push(new ShootAirplane(60, 60, x, 0))
                    }

                } else {
                    if (airRandom < 2) {
                        let idx = airplaneList.value.findIndex(it => it.type ===
                            "FastAirplane" && it
                            .run == false);

                        if (idx > -1) {
                            airplaneList.value[idx].start();
                        } else {
                            airplaneList.value.push(new FastAirplane(30, 60, x, 0))
                        }

                    } else {
                            let idx = airplaneList.value.findIndex(it => it.type === "Airplane" && it
                                .run == false);

                            if (idx > -1) {
                                airplaneList.value[idx].start();
                            } else {
                                airplaneList.value.push(new Airplane(40, 80, x, 0))
                            }
                        }
                    }
                }, 1000 / difficultLv.value)

            //敌方飞机飞行的timer
            airplaneTimer.value = setInterval(() => {
                let oriX = Math.floor(Math.random() * window
                    .innerWidth);

                let x = oriX < 20 ? 20 : oriX > window.innerWidth - 20 ? window.innerWidth -
                    80 :
                    oriX;
                airplaneList.value.forEach((element, index) => {
                    element.step();
                    if (element.outOfWindow()) {
                        // airplaneList.value.splice(index, 1)
                        element.reset(x, element.oriY)
                        element.end();
                    }
                    if (hero.value.hitBy(element)) {
                        // airplaneList.value.splice(index, 1)
                        element.reset(x, element.oriY)
                        element.end();
                        hero.value.subLife(5);
                        hero.value.supShoot()

                    }
                    if (element.canShoot && element.run) {
                        let randomNum = Math.random() * 100;
                        if (randomNum <= 2) {
                            let idx = enemyBulletList.value.findIndex(it => it.isEnd());
                            let obj = element.shoot();
                            if (idx > -1) {
                                enemyBulletList.value[idx].reset(obj.width, obj.height,
                                    obj.x, obj.y, obj.speed, obj.xSpeed)
                            } else {

                                enemyBulletList.value.push(new EnemyBullet(obj.width, obj
                                    .height,
                                    obj.x, obj.y, obj.speed, obj.xSpeed))

                            }
                        }

                    }
                })
            }, 20)

            // hero射击子弹导入的timer
            pushBulletTimer.value = setInterval(() => {
                let arr = hero.value.shoot() || [];

                arr.forEach(it => {
                    let idx = bulletList.value.findIndex(item => item.isEnd());
                    if (idx > -1) {
                        bulletList.value[idx].reset(it.w * hero.value.shootPower, it.h *
                            hero.value.shootPower, it.x, it.y, it.speed, it
                            .xSpeed)
                    } else {
                        bulletList.value.push(new Bullet(it.w * hero.value.shootPower, it
                            .h *
                            hero.value.shootPower, it.x, it.y, it
                            .speed, it
                            .xSpeed))
                    }

                })

            }, 500 / shootSpeed.value)


            //hero子弹飞行的timer
            bulletTimer.value = setInterval(() => {
                bulletList.value.forEach((element, index) => {
                    element.step();

                    if (element.outOfWindow()) {
                        element.end();
                    }

                    airplaneList.value.forEach((ele, ind) => {
                        if (ele.shootBy(element)) {
                            element.end();
                            ele.subLife(hero.value.shootPower);
                            if (ele.isDie()) {
                                score.value += ele.getScore();
                                power.value += ele.getScore();
                                if (power.value >= 100) {
                                    power.value = 100
                                }
                                ele.reset(ele.oriX, ele.oriY)
                                ele.end();
                            }

                        }
                    })
                })
            }, 20)

            maxFireBulletTimer.value = setInterval(() => {
                maxFireBulletList.value.forEach((element, index) => {
                    element.step();
                    if (element.outOfWindow()) {
                        element.end();
                    }

                    airplaneList.value.forEach((ele, ind) => {
                        //对于超快敌机无效
                        if (ele.shootBy(element) && ele.type !=
                            'FastAirplane') {
                            let nowLife = ele.life;
                            ele.subLife(nowLife);
                            element.subLife(nowLife);

                            if (element.isDie()) {
                                element.end()
                            }

                            if (ele.isDie()) {
                                power.value += Math.floor(ele.getScore()/2)
                                ele.reset(ele.oriX, ele.oriY)
                                ele.end();
                            }
                        }
                    })
                })
            }, 20)

            //敌方子弹飞行的timer
            enemyBulletTimer.value = setInterval(() => {
                enemyBulletList.value.forEach((element, index) => {
                    element.step();
                    if (element.outOfWindow()) {
                        element.end()
                    }

                    if (hero.value.hitBy(element)) {
                        element.end()
                        hero.value.subLife(2);
                        hero.value.supShoot()
                    }

                })
            }, 20)


            //奖励导入的timer
            pushAwardTimer.value = setInterval(() => {
                let airRandom = Math.floor(Math.random() * 10)
                let x = Math.floor(Math.random() * window
                    .innerWidth);
                //生命回复
                if (airRandom > 8) {
                    awardList.value.push(new Award(40, 40, x, 0, 1))
                }
                //火力加强
                if (airRandom <= 3) {
                    awardList.value.push(new Award(40, 40, x, 0, 2))
                }
                //射速
                if (airRandom > 3 && airRandom <= 5) {
                    awardList.value.push(new Award(40, 40, x, 0, 3))
                }
                //威力
                if (airRandom > 5 && airRandom <= 8) {
                    awardList.value.push(new Award(40, 40, x, 0, 4))
                }

            }, 3000)

            //奖励飞行的timer
            awardTimer.value = setInterval(() => {
                awardList.value.forEach((element, index) => {
                    element.step();
                    if (element.outOfWindow()) {
                        awardList.value.splice(index, 1)
                    }
                    if (hero.value.hitBy(element)) {
                        hero.value.getAward(element.getType())
                        awardList.value.splice(index, 1)
                    }
                })
            }, 20)
        };
         
        return{
            hero,
            airplaneList,
            awardList,
            bulletList,
            enemyBulletList,
            maxFireBulletList,
            score,difficultLv,power,pause,life,
            startTimer,
            pauseTimer,
            pushMaxFire,
            resetPushAirPlaneTimer,
            resetPushBulletTimer,
            changeShootType,
            heroMove,
            reStartGame
        };
    }
    
  };
</script>



<template>
    <div class = "independent-bg">
        <HeroView :hero="hero"/>
        <AirplaneObj :airplaneList = "airplaneList"/>   
        <AwardObj :awardList ="awardList"/>
        <BulletObj :bulletList ="bulletList"/>
        <EnemyBulletObj :enemyBulletList = "enemyBulletList"/> 
        <MaxFireObj :maxFireBulletList ="maxFireBulletList"/>
        <Info :hero ="hero" :score="score" :difficultLv="difficultLv" :power="power" :pause="pause" />
        <CoverBackground  v-show="life <= 0" :reStartGame="reStartGame" :life="life"/>
    </div>
</template>

<style >
    .independent-bg {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        bottom: 0;
        z-index: 1000;
        background: transparent;
    }
    .fly-obj {
            position: absolute;
            z-index: 500;
            background-size: 100% 100%;
            border-radius: 8px;
    }
</style>