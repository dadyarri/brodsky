package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func AskBool(prompt string, defaultValue bool) bool {
	defaultStr := "n"
	if defaultValue {
		defaultStr = "y"
	}

	fmt.Printf("%s [y/n] (default: %s): ", prompt, defaultStr)
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return defaultValue
	}
	input = strings.ToLower(strings.TrimSpace(input))

	if input == "y" {
		return true
	} else if input == "n" {
		return false
	}
	return defaultValue
}

func AskString(prompt string, defaultValue string, validate func(string) bool) string {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s (default: %s): ", prompt, defaultValue)
		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			input = defaultValue
		}
		if validate == nil || validate(input) {
			return input
		}
		fmt.Println("Invalid input. Please try again.")
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return
	}
}

func readKey() (string, error) {
	byteInput := make([]byte, 1)
	_, err := os.Stdin.Read(byteInput)
	if err != nil {
		return "", err
	}
	return string(byteInput), nil
}
