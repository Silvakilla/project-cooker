package file

import (
	"os"
	"path/filepath"
	"project-cooker/constants"
)

// ScanForFiles looks up all files which are found in the root path directory
func ScanForFiles(rootPath string) []string {
	var files []string
	root := rootPath

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		WriteStringToFile(err.Error(), constants.ERROR)
	}

	for _, file := range files {
		WriteScannedFilesToFile(file)
	}

	return files
}
