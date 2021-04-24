<template>
  <div class="editor">
    <div class="editor-wrap">
      <div class="brace-editor">
        <div id="brace">
        </div>
      </div>
      <div class="results">
        <p>Ginker is awesome!</p>
      </div>
    </div>
  </div>
</template>

<script>
import * as ace from 'brace';
import 'brace/mode/golang';
import 'brace/theme/monokai';

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
  },
  mounted() {
    console.log(document.querySelector('#brace'))
    let editor = ace.edit(document.querySelector('#brace'))
    editor.getSession().setMode("ace/mode/golang")
    editor.setTheme("ace/theme/monokai")
    editor.resize()
  },
};
</script>
