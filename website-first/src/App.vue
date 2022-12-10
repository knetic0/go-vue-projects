<template>
  <Navbar></Navbar>
  <Container></Container>
  <Footer></Footer>
</template>
<script>
  import Navbar from "./components/NavigationBar.vue"
  import Container from "./components/Container.vue"
  import Footer from "./components/Footer.vue"
  export default {
    name : "App",
    components : {
      Navbar,
      Container,
      Footer,
    },
    data() {
      return {
      
      }
    },
    data() { return {
      background_img : "https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fgetwallpapers.com%2Fwallpaper%2Ffull%2Fc%2F5%2Fd%2F688482-large-software-wallpapers-1920x1080-hd-1080p.jpg&f=1&nofb=1&ipt=cb321fada2ded5225ca308e5dc18ba7ddf2501aba2ce81cd25cc58b8ea012aa6&ipo=images",
    }},
  }
</script>