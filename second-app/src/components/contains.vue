<template>
    <div class="container">
        <div class="float-end border border-success border-5 rounded-pill">
            <h1 class="text-center">Hello World {{ val }}</h1>
            <div class="text-center">
                <img src="../assets/logo.png">
            </div>
        </div>
    </div>
</template>
<script>
    import { toContainesVal } from './leftbar.vue'
    export default {
        name: "Contains",
        data() { return {
            val : toContainesVal,
        }},
        methods : {
            
        },
    }
</script>
<style>
</style>