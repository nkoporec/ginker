package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"

	"github.com/mitchellh/mapstructure"
	"github.com/nkoporec/ginker/compiler"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails"
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
		Width:     1024,
		Height:    768,
		Title:     "ginker",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})
	app.Bind(runCompiler)
	app.Bind(getSettings)
	app.Bind(saveSettings)
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

func getSettings() (compiler.Config, error) {
	dirConfig, err := compiler.GetDirConfig()

	// Send config.
	config, err := compiler.LoadConfig(dirConfig.Dir)
	if err != nil {
		log.Fatal(err)
	}

	return config, nil
}

// Declare settings structure
type Settings struct {
	GolangBinary string
	FontSize     string
	FontFamily   string
	LineHeight   string
}

func saveSettings(data map[string]interface{}) (string, error) {
	settings := Settings{}
	mapstructure.Decode(data, &settings)

	viper.Set("GOLANG_BINARY", settings.GolangBinary)
	viper.Set("FONT_SIZE", settings.FontSize)
	viper.Set("FONT_FAMILY", settings.FontFamily)
	viper.Set("LINE_HEIGHT", settings.LineHeight)
	viper.WriteConfig()

	return "true", nil
}
