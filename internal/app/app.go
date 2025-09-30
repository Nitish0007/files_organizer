package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"github.com/Nitish0007/personal_files_organizer/utils"
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
	Run(rootPath string)
	Stop()
	Perform(rootPath string)
}

type CliApp struct {}

func NewCliApp() *CliApp {
	return &CliApp{}
}

func (a *CliApp) Run(rootPath string) {
	a.Perform(rootPath)
}

func (a *CliApp) Stop() {
	fmt.Println("GracefullyShutting down CLI App")
	os.Exit(0)
}

func (a *CliApp) Perform(rootPath string) {
	fmt.Println("Looking at your files in Downloads...")
	files, err := os.ReadDir(rootPath + "/Downloads")
	if err != nil {
		utils.LogError(err, "unable to read Downloads directory")
		return
	}
	fmt.Println("Found ", len(files), " files in Downloads")
	
	numWorkers := 10
	filesQueue := make(chan os.DirEntry, len(files))

	wg := sync.WaitGroup{}
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range filesQueue {
				if file.IsDir() {
					fmt.Println("Skipping directory: ", file.Name())
					// continue
				}
				extension := strings.ToLower(filepath.Ext(file.Name()))
				category, exists := ExtensionTypes[extension]
				if !exists {
					category = "Others"
				}
				addToDirectory(file, category)
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

func addToDirectory(file os.DirEntry, folder string) {
	srcPath, err := utils.DirToOrganize()
	if err != nil {
		utils.LogError(err, "unable to get directory to organize")
		return
	}
	srcPath = srcPath + "/" + file.Name()
	destPath, err := utils.OrganizedDirPath()
	if err != nil {
		utils.LogError(err, "unable to get organized directory path")
		return
	}
	destPath = destPath + "/" + folder + "/" + file.Name()
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