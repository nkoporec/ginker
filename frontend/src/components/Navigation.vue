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
          <a href="index" class="active">
            <font-awesome-icon icon="terminal" />
          </a>
        </li>
        <li class="" title="Settings">
          <a href="settings" class="">
            <font-awesome-icon icon="cog" />
          </a>
        </li>
      </ul>
    </aside>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: " "
    };
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.basic().then(result => {
        self.message = result;
      });
    }
  }
};
</script>
