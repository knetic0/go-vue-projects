<template>
  <div class="row container">
    <div class="col-6">
      <form class="mt-5">
        <h1 class="mb-3"> Sign in </h1>
        <input v-model="email" type="text" class="form-control mt-2 w-50" placeholder="Email">
        <input v-model="password" type="password" class="form-control mt-2 w-50" placeholder="Password">
        <button @click="SendInformation" class="btn btn-primary mt-2"><a href="" style="color: white; text-decoration: none; width: 100%; height: 100%">Sign in</a></button>
      </form>
    </div>
  </div>
</template>

<script>
  export default {
    name : "Container",

    data() { return {
      email : "",
      password : "",
      socket : "",
    }},

    mounted() {
      this.socket = new WebSocket("ws://localhost:9100/Signin")
    },

    methods : {
      SendInformation() {
        let info = {"email": this.email, "password": this.password};
        this.socket.send(JSON.stringify(info))
      },
    }
  }
</script>