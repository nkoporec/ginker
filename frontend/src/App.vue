<template>
  <div id="app">
    <div class="container">
      <Navigation />
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import "./assets/css/main.scss";
import Navigation from "./components/Navigation.vue";

export default {
  name: "app",
  components: {
    Navigation,
  },
  methods: {
    windowHeight: function() {
      return window.screen.height;
    }
  },
  mounted() {
    const container = document.querySelector(".container");
    container.style.height = `${this.windowHeight()}px`;
  },
};
</script>
