<!-- This isn't being used right now -->
<template>
 <div>
      <label for="name">Name</label>
      <input v-model="media_name_input" type="text" name="name" />
      <label for="name">File</label>
      <input ref="rootMediaFileInput" @change="newUploadFile" type="file" name="file" />
      <div>
        <button @click="creating_media = false">Cancel</button>
        <button @click="addNewMedia">Submit</button>
      </div>
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
  name: "NewMedia",
  data() {
    return {
      media_name_input: undefined,
      media_file_input: undefined
    };
  },
  props: {
      id: {
          type: Number
      }
  },
  computed: {
    inputFileEncoding() {
      let arr = this.media_file_input.name.split(".");
      return arr[arr.length - 1];
    }
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