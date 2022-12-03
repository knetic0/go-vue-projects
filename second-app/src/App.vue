<template>
  <Navbar></Navbar>
  <div class="d-flex w-75">
    <Leftbar></Leftbar>
    <Contains></Contains>
  </div>
</template>
<script>
  import Navbar from './components/navbar.vue';
  import Leftbar from './components/leftbar.vue'
  import Contains from './components/contains.vue'
  export default {
    name : "App",
    data() { return {

    }},
    components : {
      Navbar,
      Leftbar,
      Contains,
    },
    methods : {

    }
  }
</script>