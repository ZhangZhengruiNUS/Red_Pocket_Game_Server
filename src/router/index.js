import {createRouter, createWebHashHistory, createWebHistory} from "vue-router";

import StartPage from "@/StartPage.vue";
import ModeChoose from "@/ModeChoose.vue";
import WareHouse from "@/WareHouse.vue";

import GameInterface from "@/Gameview/GameInterface.vue";
import Game from "@/Game.vue";
import GameEquipment from "@/Gameview/Equipment.vue";

import Catalog from "@/views/Catalog.vue";
import ProductDetail from "@/views/ProductDetail.vue";
import ProductDetailUser from "@/views/ProductDetailUser.vue";
import Cart from "@/views/Cart.vue";
import CatalogAdmin from "@/views/CatalogAdmin.vue";
import Edit from "@/views/Edit.vue";


const router = createRouter({
    history : createWebHashHistory(import.meta.env.BASE_URL),
    routes:[
        {
            path: "/",
            name:"StartPage",
            component: StartPage
        },
        {
            path: "/ModeChoose/:userid",
            name:"ModeChoose",
            component: ModeChoose,
        },
        {
            path: "/WareHouse/:userid",
            name:"WareHouse",
            component: WareHouse,
        },
        {
            path: "/Game",
            name: Game,
            component: Game,
        },
        {
            path: "/GameInterface",
            name: GameInterface,
            component: GameInterface,
        },
        {
            path:"/GameEquipment",
            name:GameEquipment,
            component:GameEquipment
        },
        {
            path: '/CatalogAdmin',
            name: 'CatalogAdmin',
            component: CatalogAdmin
        },
        {
            path: '/Catalog/product/:id',
            name: 'ProductView',
            component: ProductDetail
        },
        {
            path: '/Catalog/productUser/:id',
            name: 'ProductViewUser',
            component: ProductDetailUser
        },
        {
            path: '/Catalog/cart',
            name: 'CartView',
            component: Cart
        },
        {
            path: '/Catalog/',
            name: 'Catalog',
            component: Catalog
        },
        {
            path: '/Catalog/Edit',
            name: 'EditView',
            component: Edit
        },
        {
            path: '/Catalog/ProductDetailUser',
            name: 'ProductDetailUser',
            component: ProductDetailUser
        },
        {
            path: "/StartPage",
            redirect:"/"
        }
    ]
})

export default router