<template>
  <div id="app" class="container mt-5 ms-5">
    <RouterLink to="/Signin">Signin</RouterLink>
    <RouterLink to="/Decrypt">Decrypt</RouterLink>
    <RouterView></RouterView>
  </div>
</template>

<script>
  export default {
    name: "App",
  }
</script>