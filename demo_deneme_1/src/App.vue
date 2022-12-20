<template>
  <div class="bg-dark text-white">
    <Sidebar></Sidebar>
    <RouterView></RouterView>
  </div>
</template>

<script>
  import Sidebar from './components/Sidebar.vue'
  import Container from './components/Container.vue'

  export default {
    name : "App",

    components : {
      Sidebar, 
      Container,
    }
  }
</script>

