<template>
  <div id="app" class="container">
    <input type="text" ref="input" v-model="header" />
    <button v-on:click="headerYazdir" @click="listAppend">Yazdir</button>
    <h1> {{ header }}</h1>
    <ul>
      <li v-for="(data, key) in itemList" :key="key">
        {{ data }} 
        <a href="#" v-if="data[0] === 'M'">{{ data }}</a>
      </li>
    </ul>
  </div>
</template>
<script>
  export default {
    name : "App",
    data() { return {
      header : "",
      itemList : [],
      search : "",
      socket : "",
      message : "",
    }},
    mounted() {
      this.socket = new WebSocket("ws://localhost:9100/socket")
      this.socket.onmessage = (event) => { 
        this.message = event.data
        this.fromDatabase(this.message)     
      }
    },
    created() {
      //console.log(this.socket.onmessage)
    },
    methods: {
      headerYazdir() {
        this.search = this.$refs.input.value
      },
      listAppend() {
        let msg = { "title" : this.search, "done" : false}
        this.socket.send(JSON.stringify(msg))
        this.itemList.push(this.$refs.input.value)
        this.header = ""
      },
      fromDatabase(data) {
        this.itemList.push(data)
      }
    }
  }
</script>