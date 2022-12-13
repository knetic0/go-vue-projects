import { createRouter, createWebHistory } from 'vue-router'

import Container from "../components/Container.vue"
import Decrypt from "../components/Dec.vue"

const routes = [
    {
        path: "/Signin",
        component: Container
    },

    {
        path: "/Decrypt",
        component: Decrypt
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router 