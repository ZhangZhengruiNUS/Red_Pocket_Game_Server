import {createRouter, createWebHashHistory, createWebHistory} from "vue-router";

import StartPage from "@/StartPage.vue";
import ModeChoose from "@/ModeChoose.vue";
import WareHouse from "@/WareHouse.vue";
import GameInterface from "@/view/GameInterface.vue";

const router = createRouter({
    history : createWebHashHistory(import.meta.env.BASE_URL),
    routes:[
        {
            path: "/",
            name:"StartPage",
            component: StartPage
        },
        {
            path: "/ModeChoose/:username",
            name:"ModeChoose",
            component: ModeChoose,
            // meta:{
            //     data:{
            //         username
            //     }
            // }
        },
        {
            path: "/WareHouse/:username",
            name:"WareHouse",
            component: WareHouse,
            // meta:{
            //     data:{
            //         username
            //     }
            // }
        },
        {
            path: "/Game",
            name: GameInterface,
            component: GameInterface,
        },
        {
            path: "/StartPage",
            redirect:"/"
        }
    ]
})

export default router