<template>
    <div>
        <p>{{name}}</p>
    </div>
</template>

<script>
export default {
  name: 'child-component',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      truthy: true,
      dogs: ['Pepper', 'Piglet', 'Piqlit']
    }
  },
  computed: {},
  props: {
      name: {
        type: String,
        required: true,
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
