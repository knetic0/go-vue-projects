<template>
    <div class="row container">
        <div class="float-start">
            <div class="leftbar">
                <h1 class="text-center mt-3 mb-3">Languages</h1>
                <ul>
                    <li v-for="(key, index) in values" :key="index">
                        <a v-on:click="pressFunction(key)" href="#"> {{ key }} </a>
                    </li>
                </ul>
                <button class="btn btn-sm btn-primary" v-on:click="pressMeFunction">Press on me!</button>
            </div>
        </div>
    </div>
</template>
<script>
    export default {
        name : "LeftBar",
        data() { return {
            toContainsVal : "",
            values : ["Python", "C++", "HTML", "CSS", "JavaScript", "GoLang", "C#"],
        }},
        methods : {
            pressFunction(key) {
                this.toContainsVal = key
                alert(key)
            },
            pressMeFunction() {
                alert("Welcome to the System!")
            }
        }
    }
</script>
<style scoped>
    .float-start {
        float: left;
    }
    ul {
        list-style-type: none;
        height: 100%;
        margin-top: 10px;
        padding: 0;
        width: 100%;
        background-color: white;
        display: flex;
        align-items: center;
        flex-direction: column;
        justify-content: space-around;
        border-radius: 10px;
        border: 5px solid #56AB6E;
    }
    a {
        text-decoration: none;
        color: #37444D;
        font-size: 16px;
        padding: 10px 12px;
    }
    li {
        margin: 20px;
        border-bottom: 3px solid black;
    }
    li a:hover {
        background-color: #56AB6E;
        border-bottom: 5px solid black;
    }
    button {
        margin-left: 40px;
    }
</style>