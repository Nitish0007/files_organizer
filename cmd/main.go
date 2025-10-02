package main

import (
	"fmt"
	"time"

	"github.com/Nitish0007/files_organizer/internal/app"
)

func main() {
	t := time.Now()
	fmt.Println("Hello, I am your personal files organizer!")
	a := app.NewCliApp()
	a.StartCli()
	fmt.Println("Time taken: ", time.Since(t))
}
