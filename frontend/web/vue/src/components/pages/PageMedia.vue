<template>
  <div class="hello">
    <p>{{ media.name }}</p>
    <ul>
      <li :key="child.id" v-for="child in media.children">{{ child.name }}</li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "HelloWorld",
  props: {
    id: {
      type: Number,
      required: true
    }
  },
  data() {
    return {
      media: undefined
    };
  },
  mounted() {
    const backendHost = process.env.VUE_APP_BACKEND_HOST;
    axios.get("http://" + backendHost + "/media/" + this.id).then(response => {
      this.media = response.data;
    });
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
</style>
