<template>
  <div id="app">
    <h1>Find a law</h1>
    <button type="button" @click="getLaw()">Get a new hacker law</button>
    <div v-if="law != null">
      <h1>{{ law.name }}</h1>
      <p>{{ law.definition }}</p>
    </div>
  </div>
</template>

<script>

export default {
  name: "App",
  components: {},
  data() {
    return {
      law: null,
    };
  },
  methods: {
    getLaw() {
      this.axios.get("/law").then((response) => (this.law = response.data));
    },
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
