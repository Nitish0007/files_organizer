package utils

import (
	"fmt"
	"os"
	"io"
)

func GetRootPath() (string, error) {
	rootPath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return rootPath, nil
}


func CreateOrganizedDirectory(rootPath string) string {
	CreateDirectory(rootPath + "/Downloads/Organized/Documents")	
	CreateDirectory(rootPath + "/Downloads/Organized/Images")
	CreateDirectory(rootPath + "/Downloads/Organized/Videos")
	CreateDirectory(rootPath + "/Downloads/Organized/Audio")
	CreateDirectory(rootPath + "/Downloads/Organized/Archives")
	CreateDirectory(rootPath + "/Downloads/Organized/Code")
	CreateDirectory(rootPath + "/Downloads/Organized/Executables")
	// CreateDirectory(rootPath + "/Downloads/Organized/Folders") # NOTE: Might need to handle nested copying/moving of folders, will do it later
	CreateDirectory(rootPath + "/Downloads/Organized/Others")
	return rootPath + "/Downloads/Organized"
}

func DirToOrganize() (string, error) {
	rp, err := GetRootPath()
	if err != nil {
		return "", err
	}
	return rp + "/Downloads", nil
}

func OrganizedDirPath() (string, error) {
	dirToOrganize, err := DirToOrganize()
	if err != nil {
		return "", err
	}
	return dirToOrganize + "/Organized", nil
}

func CreateDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
}

func GetStandardizedSize(size int64) string {
	finalValue := "unknown"

	switch {
		case size > 0 && size < 1024:
			finalValue = fmt.Sprintf("%d B", size)
		case size >= 1024 && size < (1024 * 1024):
			finalValue = fmt.Sprintf("%d KB", size / 1024)
		case size >= (1024 * 1024) && size < (1024 * 1024 * 1024):
			finalValue = fmt.Sprintf("%d MB", size / (1024 * 1024))
		case size >= (1024 * 1024 * 1024) && size < (1024 * 1024 * 1024 * 1024):
			finalValue = fmt.Sprintf("%d GB", size / (1024 * 1024 * 1024))
		default:
			finalValue = "unknown"
	}
	return finalValue
}

func MoveFile(srcPath, destPath string) error {
	fmt.Println("Moving file from ", srcPath, " to ", destPath)
	err := os.Rename(srcPath, destPath)
	return err
}

func CopyFile(srcPath, destPath string) error {
	fmt.Println("Copying file from ", srcPath, " to ", destPath)
	srcFile, err := os.Open(srcPath)
	LogError(err, "unable to open source file")
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	LogError(err, "unable to create destination file")
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		LogError(err, "unable to copy file")
		return err
	}
	return nil
}