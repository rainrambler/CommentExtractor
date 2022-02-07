package main

import (
	"log"
	"os"
	"path/filepath"
)

// https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html
func parseFiles(dir string) {
	allres := ""

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) != ".c" {
			return nil
		}

		res := parseFile(path)
		allres += res
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	WriteTextFile(filepath.Join(dir, "allcomments.txt"), allres)
}

// https://golangr.com/rename-file/
func renameFile(src, dst string) {
	// rename file
	os.Rename(src, dst)
}

const NewLineChar = '\n'

func parseFile(filename string) string {
	content, err := ReadTextFile(filename)
	if err != nil {
		return ""
	}

	contentnew := formatString(content)
	return ">>> Parsing " + filename + "\n\n" + contentnew

	//WriteTextFile(filename+".tmp", contentnew)
}
