<template>
  <div class="navigation">
    <div class="logo">
      <a href="https://github.com/nkoporec/ginker" target="_blank" rel="noopener noreferrer">
        <img class="" src="../assets/images/logo.png" alt="ginker" />
      </a>
    </div>
    <aside class="">
      <ul class="">
        <li class="" title="Terminal">
          <router-link to="/editor">
            <font-awesome-icon icon="terminal" />
          </router-link>
          <router-link to="/settings">
            <font-awesome-icon icon="cog" />
          </router-link>
        </li>
      </ul>
    </aside>
  </div>
</template>

<script>
export default {
  data() {
    return {
    };
  },
  methods: {
  },
  mounted() {
  },
};
</script>
