<template>
  <div id="app" class="container">
    <center>
      <br />
      <input class="form-control w-50" placeholder="Enter a todo" type="text" ref="input" v-model="header" />
      <br />
      <button class="btn btn-primary" v-on:click="headerYazdir" @click="listAppend">Yazdir</button> 
      <h1> {{ header }}</h1>
      <br />
      <ul class="list-group w-25 list-group-flush">
        <li class="list-group-item" v-for="(data, key) in itemList" :key="key">
          {{ data }} 
          <button class="btn btn-primary" @click="isDone(key)">✓</button>
          <a href="#" v-if="data[0] === 'M'">{{ data }}</a>
        </li>
      </ul>
    </center>
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
        if (this.message != "")
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
        this.itemList.push(msg)
        this.header = ""
      },
      fromDatabase(data) {
        this.itemList.push(data)
      },
      isDone(index) {
        this.itemList.splice(index, 1)
      }
    }
  }
</script>