import { createRouter, createWebHistory } from "vue-router";

import Container from "../components/Container.vue"
import Users from "../components/Users.vue"
import Stats from "../components/Stats.vue"
import Dashboard from "../components/Dashboard.vue"
import Register from "../components/Register.vue"

const routes = [
    {
        path: "/charts",
        component: Container
    },

    {
        path: "/dashboard",
        component: Dashboard,
    },
    {
        path: "/users",
        component: Users
    },
    {
        path: "/stats",
        component: Stats 
    },
    {
        path: "/register",
        component: Register
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router