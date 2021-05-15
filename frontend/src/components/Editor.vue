<template>
  <div class="editor">
    <div class="editor-wrap">
      <div class="brace-editor">
        <div id="brace">
        </div>
      </div>
      <div class="results">
        <p>
          {{ result }}
        </p>
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
      message: " ",
      editor: null,
      result: "",
      typingTimer: null,
      doneTypingInterval: 250
    };
  },
  methods: {
    editorChange: function(value) {
      window.backend.runCompiler(value).then(result => {
        this.result = result;
      });
    }
  },
  mounted() {
    // Set up ace.
    this.editor = ace.edit(document.querySelector('#brace'))
    this.editor.getSession().setMode("ace/mode/golang")
    this.editor.setTheme("ace/theme/monokai")
    this.editor.resize()

    // Set initial value.
    this.editor.session.setValue(`package main\r\n\r\nimport (\r\n "fmt"\r\n)\r\n\r\nfunc main() {\r\n fmt.Println("Ginker is awesome")\r\n}\r\n`)

    // Set change event.
    let self = this;
    this.editor.session.on('change', function() {
      clearTimeout(this.typingTimer);
      if (self.editor.getValue()) {
        self.typingTimer = setTimeout(
          self.editorChange.bind(null, self.editor.getValue()),
          self.doneTypingInterval
        );
      }
    });

    // Set correct height.
    const editor = document.querySelector(".brace-editor");
    editor.style.height = `${window.screen.height}px`;
    const results = document.querySelector(".results");
    results.style.height = `${window.screen.height}px`;

    // Apply settings.
    window.backend.getSettings().then(settings => {
      this.editor.setOptions({
        fontFamily: settings.FontFamily,
      });
      this.editor.setFontSize(Number(settings.FontSize))

      this.editor.container.style.lineHeight = Number(settings.LineHeight)
      this.editor.renderer.updateFontSize()
    });
  },
};
</script>
