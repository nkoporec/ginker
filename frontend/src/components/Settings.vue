<template>
  <div class="settings">
    <section>
      <div>
        <label for="go-binary">
          <span>Golang Binary:</span>
        </label>
        <input type="textfield" id="go-binary" name="go-binary" v-model="GolangBinary" @change="updateSettings"/>
      </div>
      <div>
        <label for="font-family">
          <span>Font Family:</span>
        </label>
        <input type="textfield" id="font-family" name="font-family" v-model="FontFamily" @change="updateSettings" />
      </div>
      <div>
        <label for="font-size">
          <span>Font size:</span>
        </label>
        <input type="number" id="font-size" name="font-size" v-model="FontSize" @change="updateSettings" />
      </div>
      <div>
        <label for="line-height">
          <span>Line height:</span>
        </label>
        <input type="number" id="line-height" name="line-height" v-model="LineHeight" @change="updateSettings" />
      </div>
    </section>
  </div>
</template>

<script>
export default {
  data() {
    return {
      'GolangBinary': "",
      'FontSize': 12,
      'FontFamily': "Monospace",
      'LineHeight': 1,
    };
  },
  methods: {
    updateSettings: function() {
      const settings = {
        'GolangBinary': this.GolangBinary,
        'FontSize': this.FontSize,
        'FontFamily': this.FontFamily,
        'LineHeight': this.LineHeight,
      };

      window.backend.saveSettings(settings).then(settings => {
        console.log(settings)
      });
    }
  },
  mounted() {
    window.backend.getSettings().then(settings => {
      this.GolangBinary = settings.GolangBinary
      this.FontSize = settings.FontSize
      this.FontFamily = settings.FontFamily
      this.LineHeight = settings.LineHeight
    });
  },
};
</script>
