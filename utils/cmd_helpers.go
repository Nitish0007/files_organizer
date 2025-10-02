package utils

import (
	"bufio"
	"os"
	"strings"
	"errors"
	"fmt"
)

func ExitCommand(){
	fmt.Println("Gracefully shutting down...")
	os.Exit(0)
}

func TakeInputYesOrNo(message string) string {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdout)
	input, err := reader.ReadString('\n')
	LogErrorAndExit(err, "unable to read input")
	input = strings.TrimSpace(input)
	if input == "y" || input == "n" || input == "Y" || input == "N" {
		return input
	}
	if input == "exit" {
		ExitCommand()
	}
	fmt.Printf("Invalid input: %s. Please enter y or n\n", input)
	return TakeInputYesOrNo(message)
}

func TakeCommandAsInput(message string) string {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdout)
	input, err := reader.ReadString('\n')
	LogErrorAndExit(err, "unable to read input")
	input = strings.TrimSpace(input)
	if input == "" {
		LogError(errors.New("empty string"), "Invalid input")
		return TakeCommandAsInput(message)
	}
	if input == "exit" {
		ExitCommand()
	}
	return input
}

func ValidateDirectoryPath(path string) bool {
	if path == "" {
		LogError(errors.New("empty string"), "Invalid input")
		return false
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		LogError(err, "unable to get dir/file info")
	}
	return fileInfo.IsDir() 
}