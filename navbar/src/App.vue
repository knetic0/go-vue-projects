<template>
  <div id="app">
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
  <a class="navbar-brand ms-3" href="#">Versebyte</a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>
  <div class="collapse navbar-collapse" id="navbarNav">
    <ul class="navbar-nav">
      <li v-for="(data, key) in values" :key="key" class="nav-item active">
        <a class="nav-link" href="#"> {{ data }}</a>
      </li>
    </ul>
  </div>
</nav>
  </div>
</template>
<script>
  export default {
    data() { return {
      socket : "",
      values : [],
      message : "",
    }},
    mounted() {
      this.socket = new WebSocket("ws://localhost:9100/socket")
      this.socket.onmessage = (event) => {
        this.message = event.data
        this.Print(this.message)
      }
    },
    methods : {
      Print(mesg) {
        this.values.push(mesg)
      }
    }
  }
</script>
<style scoped>
</style>