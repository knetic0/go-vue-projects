<template>
    <input type="text" placeholder="Email" v-model="email" class="form-control mt-5 mb-5">
    <button class="btn btn-primary" @click="Decrypt(email)">Take Password</button>
</template>

<script>
    export default {
        data() { return{
            socket : "",
            email : "",
        }},

        mounted() {
            this.socket = new WebSocket("ws://localhost:9100/Decrypt")
        },

        methods : {
            Decrypt(email) {
                let msg = {"email": email, "password": null};
                this.socket.send(JSON.stringify(msg))
            }
        }
    }
</script>