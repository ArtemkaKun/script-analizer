package main

import (
	"bufio"
	"os"
	"path/filepath"
)

func findAllScripts(args ConsoleArguments) (scripts []File, err error) {
	err = filepath.Walk(args.DirWithScripts, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			return nil
		}

		matched, err := filepath.Match(args.ExtensionToSearch, filepath.Base(path))
		if err != nil {
			return err
		} else if matched {
			file, err := getFileData(path)
			if err != nil {
				return err
			}

			scripts = append(scripts, file)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return
}

func getFileData(path string) (file File, err error) {
	fileData, err := os.Open(path)
	if err != nil {
		return File{}, err
	}
	defer fileData.Close()

	scanner := bufio.NewScanner(fileData)
	for scanner.Scan() {
		file.LineCount += 1
	}

	if err := scanner.Err(); err != nil {
		return File{}, err
	}

	file.Name = fileData.Name()
	return
}
