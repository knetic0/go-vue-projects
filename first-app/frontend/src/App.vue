<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Generate a thumbnail of a website</h1>
        <form v-on:submit.prevent="makeWebsiteThumbnail">
          <div class="form-group">
            <input v-model="websiteUrl" placeholder="Enter a website" type="text" id="website-input" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Generate!</button>
          </div>
          <img :src="thumbnailUrl" />
        </form>
      </div>
    </div>
  </div>
</template>
<script>
  import axios from 'axios';
  export default {
    name : 'App',
    data() { return {
      websiteUrl : '',
      thumbnailUrl: '',
    } },
    methods : {
      makeWebsiteThumbnail() {
        axios.post("http://localhost:3000/api/thumbnail", {
          url: this.websiteUrl,
      })
      .then((response) => {
        this.thumbnailUrl = response.data.screenshot;
      })
      .catch((error) => {
        window.alert(`The API returned an error: ${error}`);
      })
      }
    }
  }
</script>