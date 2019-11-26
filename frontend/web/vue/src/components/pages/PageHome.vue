<template>
  <div class="hello">
    <h3>Media</h3>
    <button v-if="!creating_media" @click="creating_media = true">New Media</button>
    <div v-if="creating_media">
      <label for="name">Name</label>
      <input type="text" name="name" />
      <label for="name">File</label>
      <input type="file" name="file" />
      <div>
        <button @click="creating_media = false">Cancel</button>
        <button @click="creating_media = false">Submit</button>
      </div>
    </div>
    <ul>
      <li :key="media.id" v-for="media in media_list">
        <router-link :to="'/media/' + media.id" >{{ media.name }}</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "PageHome",
  data() {
    return {
      creating_media: false,
      loading: true,
      media_list: []
    };
  },
  mounted() {
    axios.get("http://localhost:8000/media").then((response) => {
      for (const i in response.data.media) {
        this.media_list.push(response.data.media[i]);
      }
    });
  },
  methods: {
    addNewMedia() {
      this.media_list.unshift({});
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  margin: 40px 10px;
}
input {
  margin: 10px;
}
</style>
