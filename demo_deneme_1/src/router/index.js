import { createRouter, createWebHistory } from "vue-router";

import Container from "../components/Container.vue"
import App from "../App.vue"

const routes = [
    {
        path: "/dashboard",
        component: Container
    },

    {
        path: "/",
        component: App
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router