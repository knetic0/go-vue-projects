import { createRouter, createWebHistory } from "vue-router";

import Container from "../components/Container.vue"
import App from "../App.vue"
import Users from "../components/Users.vue"
import Stats from "../components/Stats.vue"

const routes = [
    {
        path: "/charts",
        component: Container
    },

    {
        path: "/",
        component: App,
    },
    {
        path: "/users",
        component: Users
    },
    {
        path: "/stats",
        component: Stats 
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router