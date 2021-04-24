<template>
  <div class="navigation">
    <div class="logo">
      <a href="https://github.com/nkoporec/ginker" target="_blank" rel="noopener noreferrer">
        <img class="" src="../assets/images/logo.png" alt="ginker" />
      </a>
    </div>
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
