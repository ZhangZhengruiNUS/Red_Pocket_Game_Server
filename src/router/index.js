import {createRouter, createWebHashHistory, createWebHistory} from "vue-router";

import StartPage from "@/StartPage.vue";
import ModeChoose from "@/ModeChoose.vue";
import WareHouse from "@/WareHouse.vue";
import GameInterface from "@/view/GameInterface.vue";
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
            path: '/store',
            name: 'CatalogAdmin',
            component: CatalogAdmin
        },
        {
            path: '/store/product/:id',
            name: 'ProductView',
            component: ProductDetail
        },
        {
            path: '/store/productUser/:id',
            name: 'ProductViewUser',
            component: ProductDetailUser
        },
        {
            path: '/store/cart',
            name: 'CartView',
            component: Cart
        },
        {
            path: '/store/Catalog',
            name: 'Catalog',
            component: Catalog
        },
        {
            path: '/store/Edit',
            name: 'EditView',
            component: Edit
        },
        {
            path: '/store/ProductDetailUser',
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