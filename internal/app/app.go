package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"github.com/Nitish0007/files_organizer/utils"
)

var totalSize int64 = 0

var ExtensionTypes = map[string]string {
	// Images
	".jpg":  "Images",
	".jpeg": "Images", 
	".png":  "Images",
	".gif":  "Images",
	".bmp":  "Images",
	".svg":  "Images",
	".webp": "Images",
	
	// Documents
	".pdf":  "Documents",
	".doc":  "Documents",
	".docx": "Documents",
	".txt":  "Documents",
	".rtf":  "Documents",
	".odt":  "Documents",
	".xls":  "Documents",
	".csv":  "Documents",
	
	// Videos
	".mp4":  "Videos",
	".avi":  "Videos",
	".mov":  "Videos",
	".wmv":  "Videos",
	".flv":  "Videos",
	".mkv":  "Videos",
	".webm": "Videos",
	
	// Audio
	".mp3":  "Audio",
	".wav":  "Audio",
	".flac": "Audio",
	".aac":  "Audio",
	".ogg":  "Audio",
	
	// Archives
	".zip":  "Archives",
	".rar":  "Archives",
	".7z":   "Archives",
	".tar":  "Archives",
	".gz":   "Archives",
	
	// Code
	".go":   "Code",
	".js":   "Code",
	".py":   "Code",
	".java": "Code",
	".cpp":  "Code",
	".c":    "Code",
	".html": "Code",
	".css":  "Code",
	".php":  "Code",
	".rb":   "Code",
	
	// Executables
	".exe":  "Executables",
	".app":  "Executables",
	".dmg":  "Executables",
	".pkg":  "Executables",
}

type App interface {
	StartCli()
	Perform(rootPath, pathToOrganize string)
}

type CliApp struct {}

func NewCliApp() *CliApp {
	return &CliApp{}
}

func (a *CliApp) StartCli() {
	rootPath, err := utils.GetRootPath()
	defaultDir := "Downloads"
	pathToOrganize := rootPath + "/" + defaultDir
	utils.LogErrorAndExit(err, "unable to get root path")
	fmt.Println("Your root path is: ", rootPath)
	for {
		fmt.Println("Default directory to organize is ", defaultDir)
		input  := utils.TakeInputYesOrNo("Want to change the directory?")
		if input == "y" || input == "Y" {
			for {
				dir := utils.TakeCommandAsInput("Enter the Directory")
				dir = strings.TrimPrefix(strings.TrimSpace(dir), "/")
				fullPath := rootPath + "/" + dir
				isValid := utils.ValidateDirectoryPath(fullPath)
				if isValid {
					defaultDir = dir
					pathToOrganize = fullPath
					break
				}
			}
		}

		a.Perform(rootPath, pathToOrganize)
	}
}

func (a *CliApp) Perform(rootPath, pathToOrganize string) {
	defer atomic.StoreInt64(&totalSize, 0)
	fmt.Println("Organizing files in " + pathToOrganize + "...")
	// utils.CreateOrganizedDirectory(pathToOrganize)
	utils.CreateDirectory(pathToOrganize + "/Organized")
	files, err := os.ReadDir(pathToOrganize)
	if err != nil {
		utils.LogError(err, "unable to read " + pathToOrganize + " directory")
		return
	}
	fmt.Println("Found ", len(files), " files in " + pathToOrganize)
	
	numWorkers := 10
	filesQueue := make(chan os.DirEntry, len(files))

	wg := sync.WaitGroup{}
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range filesQueue {
				if file.IsDir() || !file.Type().IsRegular() {
					fmt.Println("Skipping non-regular file or directory: ", file.Name())
					continue
				}
				extension := strings.ToLower(filepath.Ext(file.Name()))
				category, exists := ExtensionTypes[extension]
				if !exists {
					fmt.Println("Extension not found: ", extension, file.Name())
					category = "Others"
				}
				addToDirectory(file, category, pathToOrganize)
			}
 		}()
	}

	for _, file := range files {
		if file.Type().IsRegular() {
			filesQueue <- file
		}
	}
	close(filesQueue)
	wg.Wait()
	fmt.Printf("%d files copied\n", len(files))
	fmt.Println("Total size: ", utils.GetStandardizedSize(totalSize))
}

func addToDirectory(file os.DirEntry, folder, pathToOrganize string) {
	if file.IsDir() || !file.Type().IsRegular() {
		return
	}
	utils.CreateDirectory(pathToOrganize + "/Organized" + "/" + folder)
	srcPath := pathToOrganize + "/" + file.Name()
	destPath := pathToOrganize + "/Organized" + "/" + folder + "/" + file.Name()
	// utils.CopyFile(srcPath, destPath) # if you want to copy instead of move
	fileInfo, err := os.Stat(srcPath)
	if err != nil {
		utils.LogError(err, "unable to get file info")
		return
	}
	atomic.AddInt64(&totalSize, fileInfo.Size())

	err = utils.MoveFile(srcPath, destPath)
	if err != nil {
		utils.LogError(err, "unable to move file")
		atomic.AddInt64(&totalSize, -fileInfo.Size())
		return
	}
}