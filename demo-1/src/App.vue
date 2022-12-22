<template>
  <div class="w-100 d-flex">
    <Sidebar></Sidebar>
    <RouterView></RouterView>
  </div>
</template>

<script>
  import Sidebar from './components/Sidebar.vue'

  export default {
    name : "App",

    components : {
      Sidebar, 
    },
  }
</script>



