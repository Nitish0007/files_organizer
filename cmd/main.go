package main

import (
	"fmt"
	"time"

	"github.com/Nitish0007/personal_files_organizer/internal/app"
	"github.com/Nitish0007/personal_files_organizer/utils"
)

func main() {
	t := time.Now()
	fmt.Println("Hello, I am your personal files organizer!")
	rootPath, err := utils.GetRootPath()
	utils.LogErrorAndExit(err, "unable to get root path")
	utils.CreateOrganizedDirectory(rootPath)
	a := app.NewCliApp()
	a.Run(rootPath)
	fmt.Println("Time taken: ", time.Since(t))
}