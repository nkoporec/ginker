package compiler

import (
	"bufio"
	"os/exec"
	"strings"
)

func GetGoBinaryPath() string {
	// @TODO: This should use the viper config.
	binary,_ := whichGoBinary()

	// Make sure we don't have any new lines.
	binary = strings.TrimSuffix(binary, "\n")

	return binary
}

func Execute(arg ...string) (string, error) {
	ginkerDir, err := GetDirConfig()
	if err != nil {
	}

	// Create a cmd.
	cmd := exec.Command(GetGoBinaryPath(), arg...)
	// Set working dir.
	cmd.Dir = ginkerDir.Dir

	// Set up stderr
	stderr, err := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
	}

	// Set up scanner.
	scanner := bufio.NewScanner(stderr)
	errors := ""
	for scanner.Scan() {
		errors = errors + " || " + scanner.Text()
	}

	if errors != "" {
		return errors, nil
	}

	result, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(result), nil
}
