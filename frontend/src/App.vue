<template>
  <div id="app">
    <div class="container">
      <Navigation />
      <Editor />
    </div>
  </div>
</template>

<script>
import "./assets/css/main.scss";
import Navigation from "./components/Navigation.vue";
import Editor from "./components/Editor.vue";

export default {
  name: "app",
  components: {
    Navigation,
    Editor,
  }
};
</script>
