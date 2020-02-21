<template>
  <div class="hello">
    <b-jumbotron :header="media.name">
      <hr class="my-4">
      <b-button variant="success" v-if="media.encoding === ''">
        Add New Media
      </b-button>
      <b-button variant="danger">
        Delete
      </b-button>
      <b-button variant="primary" v-if="media.encoding === 'mp4'">
        Watch
      </b-button>
    </b-jumbotron>
    <div class="container">
      <div class="row">
        <div class="col-lg-2" :key="media.id" v-for="media in media.children">
          <router-link :to="'/media/' + media.id">
            <b-card
              img-src="https://cdn.traileraddict.com/content/warner-bros-pictures/blade-runner-2049-poster-4.jpg"
              :title="media.name"
              class="mb-2"
            >
              <b-card-text>{{ media.name }}</b-card-text>
              <template v-slot:footer>
                <em>Footer Slot</em>
              </template>
            </b-card>
          </router-link>
        </div>
      </div>
    </div>
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
  display: block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  margin: 20px 10px;
}
</style>
