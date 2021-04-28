package main

import (
  _ "embed"
	"os"
	"os/exec"
  "github.com/wailsapp/wails"
  "github.com/nkoporec/ginker/compiler"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	// Create working dir.
	_, err := compiler.CreateWorkingDir()
	if err != nil {
	}

  // Create module.
	dirConfig, err := compiler.GetDirConfig()
	if err != nil {
	}
	compiler.Execute("mod", "init", dirConfig.ModuleName)

	// Load config.
	compiler.LoadConfig(dirConfig.Dir)

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "ginker",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(runCompiler)
  app.Run()
}

func runCompiler(value string) (string, error) {
	ginkerDir, err := compiler.GetDirConfig()
	if err != nil {
	}

	// write the whole body at once
	file, _ := os.Create(ginkerDir.File)
	file.Write([]byte(value))
	file.Close()

	// Run tidy.
	tidy, err := compiler.Execute("mod", "tidy")
	if tidy != "" {
		return tidy, nil
	}

	// Run fmt.
	fmt, err := compiler.Execute("fmt")
	if fmt != "" {
		return fmt, nil
	}

	// Test run and display any potential errors.
	testRun, err := compiler.Execute("run", ginkerDir.File)
	if testRun != "" {
		return testRun, nil
	}

	// Execute for real.
	// @TODO: This should work with execute()
	result, err := exec.Command(compiler.GetGoBinaryPath(), "run", ginkerDir.File).Output()
	if err != nil {
    return err.Error(), nil
	}

	return string(result), nil
}
