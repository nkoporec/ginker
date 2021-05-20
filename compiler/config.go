package compiler

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	GolangBinary string `mapstructure:"GOLANG_BINARY"`
	FontSize     string `mapstructure:"FONT_SIZE"`
	FontFamily   string `mapstructure:"FONT_FAMILY"`
	LineHeight   string `mapstructure:"LINE_HEIGHT"`
}

func LoadConfig(path string) (config Config, err error) {
	if _, err := os.Stat(filepath.Join(path, "ginker.env")); os.IsNotExist(err) {
		os.Create(filepath.Join(path, "ginker.env"))
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("ginker")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	if len(viper.AllKeys()) < 4 {
		binary, err := whichGoBinary()
		if err != nil {
		}

		viper.Set("GOLANG_BINARY", binary)
		viper.Set("FONT_SIZE", "12")
		viper.Set("FONT_FAMILY", "monospace")
		viper.Set("LINE_HEIGHT", "1")
		viper.WriteConfig()
		viper.ReadInConfig()
	}

	err = viper.Unmarshal(&config)
	return
}

func SaveConfig(path string, config *Config) {
	viper.Set("GOLANG_BINARY", config.GolangBinary)
	viper.Set("FONT_SIZE", config.FontSize)
	viper.Set("FONT_FAMILY", config.FontFamily)
	viper.Set("LINE_HEIGHT", config.LineHeight)
	viper.WriteConfig()
}

func whichGoBinary() (string, error) {
	goBinary := ""

	// If we don't have (on first startup) then
	// try to get it, by running $(which)
	// @TODO: How to do this in Windows ?
	switch os := runtime.GOOS; os {
	case "darwin", "linux":
		cmd := exec.Command("which", "go")
		stdout, err := cmd.Output()

		if err != nil {
		}
		goBinary = string(stdout)
	}

	return goBinary, nil
}
