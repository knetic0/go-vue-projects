<template>
  <nav class="navbar navbar-expand-lg navbar-white bg-white">
    <div class="container-fluid">
      <a class="navbar-brand" href="#"> <span class="fs-2 ms-4">{{ Navbar }}</span></a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarText">
        <ul class="navbar-nav me-auto mt-1 mb-2 mb-lg-0 ms-5">
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" href="Anasayfa">Anasayfa</a>
          </li>
          <li class="nav-item">
            <a class="nav-link active" href="Blog">Blog</a>
          </li>
          <li class="nav-item">
            <a class="nav-link active" href="Hakkımızda">Hakkımızda</a>
          </li>
          <li class="nav-item">
            <a class="nav-link active" href="İletişim">İletişim</a>
          </li>
        </ul>
        <span class="navbar-text me-5 mt-1">
          Gerçekçi Bir Vizyon ile...
        </span>
      </div>
    </div>
  </nav>
  </template>
  <script>
    export default {
      name : "NavigationBar",
      data() { return {
        Navbar : "Versebyte",
        NavbarText : "Gerçekçi Bir Vizyon İle...",
      }}
    }
  </script>
  <style scoped>
  </style>