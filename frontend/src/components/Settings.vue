<template>
  <div class="settings">
    <section>
      <div>
        <label for="go-binary">
          <span>Golang Binary:</span>
        </label>
        <input type="textfield" id="go-binary" name="go-binary" value="" />
      </div>
      <div>
        <label for="theme">
          <span>Theme:</span>
        </label>
        <select id="theme">
          <option value="monokai" selected>Monokai</option>
        </select>
      </div>
      <div>
        <label for="font">
          <span>Font Family:</span>
        </label>
        <select id="font">
          <option value="" selected>Monospace</option>
        </select>
      </div>
      <div>
        <label for="font-size">
          <span>Font size:</span>
        </label>
        <input type="number" id="font-size" name="font-size" value="12" />
      </div>
      <div>
        <label for="line-height">
          <span>Line height:</span>
        </label>
        <input type="number" id="line-height" name="line-height" value="1.2" />
      </div>
    </section>
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
