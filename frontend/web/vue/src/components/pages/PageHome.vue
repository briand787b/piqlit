<template>
  <div class="hello">
    <h3>Media</h3>
    <button v-if="!creating_media" @click="creating_media = true">New Media</button>
    <div v-if="creating_media">
      <label for="name">Name</label>
      <input v-model="media_name_input" type="text" name="name" />
      <label for="name">File</label>
      <input ref="rootMediaFileInput" @change="newUploadFile" type="file" name="file" />
      <div>
        <button @click="creating_media = false">Cancel</button>
        <button @click="addNewMedia">Submit</button>
      </div>
    </div>
    <ul>
      <li :key="media.id" v-for="media in media_list">
        <router-link :to="'/media/' + media.id">{{ media.name }}</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

const backendHost = process.env.VUE_APP_BACKEND_HOST;
const instance = axios.create({
  baseURL: "http://" + backendHost,
  timeout: 1000,
  crossdomain: true,
  withCredentials: false
});

export default {
  name: "PageHome",
  data() {
    return {
      creating_media: false,
      loading: true,
      media_list: [],
      media_name_input: undefined,
      media_file_input: undefined
    };
  },
  computed: {
    inputFileEncoding() {
      let arr = this.media_file_input.name.split(".");
      return arr[arr.length - 1];
    }
  },
  mounted() {
    instance.get("/media").then(response => {
      for (const i in response.data.media) {
        this.media_list.push(response.data.media[i]);
      }
    });
  },
  methods: {
    newUploadFile() {
      this.media_file_input = this.$refs["rootMediaFileInput"].files[0];
    },
    addNewMedia() {
      const body = {
        name: this.media_name_input,
        length: this.media_file_input.size,
        encoding: this.inputFileEncoding
      };
      // console.log("POSTing " + JSON.stringify(body) + " to: " + url);
      instance
        .post("/media", body)
        .then(response => {
          console.log(response);
          if (response.status > 299) {
            throw "non-2XX response status code"
          }

          this.media_list.push(response.data)
        })
        .then(() => {
          let formdata = new FormData()
          formdata.append('file', this.media_file_input)
          let lastId = this.media_list[this.media_list.length -1].id
          return instance.put('/media/'+lastId+'/upload/raw', formdata, {
            headers: {
              'Content-Type': 'multipart/form-data'
            }
          })
        })
        .then((response) => {
          console.log(response);
          if (response.status > 299) {
            throw "non-2XX response status code"
          }
        })
        .catch(e => {
          console.log(e);
        });
    }
  }
};
</script>

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
  margin: 10px 10px;
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
