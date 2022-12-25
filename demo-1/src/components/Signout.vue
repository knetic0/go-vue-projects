<template>
    <h1 class="text-center">Cıkış Başarılı!</h1>
</template>

<script>
    export default {
        name : "Signout",

        data() {
            return {
                socket : "",
            }
        },

        created() {
            this.socket = new WebSocket("ws://localhost:9100/signout")
        }
    }
</script>