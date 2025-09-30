package utils

import (
	"fmt"
	"os"
)

func LogError(err error, message string) {
	if err == nil {
		return
	}

	messageToPrint := message
	if message != ""{
		messageToPrint = message + " - " + err.Error()
	} else {
		messageToPrint = err.Error()
	}
	fmt.Println("Error: ", messageToPrint)
}

func LogErrorAndExit(err error, message string) {
	if err != nil {
		fmt.Println("Error: ", message + " - " + err.Error())
		os.Exit(1)
		return
	}
}