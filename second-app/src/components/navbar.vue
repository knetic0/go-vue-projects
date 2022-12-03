<template>
    <div>
        <div class="row">
            <div class="col">
                <ul>
                    <li v-for="(key, index) in navbarList " :key="index">
                        <a href="">{{key}}</a>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>
<script>
    export default {
        name : "NavBar",
        data() { return {
            navbarList : ["Home", "News", "Contact", "About"],
        }}
    }
</script>
<style scoped>
    ul {
        list-style-type: none;
        height: 8vh;
        margin: 0;
        padding: 0;
        width: 100%;
        background-color: #37444D;
        color: white;
        display: flex;
        align-items: center;
        flex-direction: row;
        justify-content: space-around;
    }
    a {
        text-decoration: none;
        color: white;
        font-size: 20px;
        padding: 16px 16px;
    }
    li a:hover {
        background-color: #56AB6E;
    }
</style>