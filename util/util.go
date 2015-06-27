package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadImagesFromDir(dir string) []string {
	path := fmt.Sprintf("%s/*.jpg", dir)
	matches, err := filepath.Glob(path)
	if err != nil {
		panic("Error while going over directory")
	}

	return matches
}

func Log(in string) {
	s := os.Getenv("SILENT")
	if len(s) == 0 {
		fmt.Println(in)
	}
}
