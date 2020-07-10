package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dustin/go-humanize"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	args := handleConsoleArguments()

	scripts, err := findAllScripts(args)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(prepareOutputData(scripts))
}

func handleConsoleArguments() (args ConsoleArguments){
	dirWithScripts := flag.String("dir", "", "Path to dir with scripts")
	extensionToSearch := flag.String("ext", "", "Scripts extension to search")
	flag.Parse()

	args = ConsoleArguments{
		DirWithScripts: *dirWithScripts,
		ExtensionToSearch: "*" + *extensionToSearch,
	}

	return
}

func findAllScripts(args ConsoleArguments) (scripts []File, err error) {
	err = filepath.Walk(args.DirWithScripts, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			return nil
		}

		if matched, err := filepath.Match(args.ExtensionToSearch, filepath.Base(path)); err != nil {
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

	return
}

func prepareOutputData(scripts []File) string {
	totalAmountOfScripts := fmt.Sprintf("Scripts in the project: %s\n", humanize.Comma(int64(len(scripts))))
	totalAmountOfLinesOfCode := fmt.Sprintf("Lines of code in the project: %s\n", humanize.Comma(int64(calcLinesOfCode(scripts))))
	theBiggestScript := fmt.Sprintf("The biggest file in the project: %s lines of code\n", humanize.Comma(int64(findBiggestFile(scripts).LineCount)))

	return totalAmountOfScripts + totalAmountOfLinesOfCode + theBiggestScript
}

func calcLinesOfCode(scripts []File) (linesOfCode uint) {
	for _, script := range scripts {
		linesOfCode += script.LineCount
	}

	return
}

func findBiggestFile(scripts []File) File {
	sort.Slice(scripts, func(i, j int) bool {
		return scripts[i].LineCount > scripts[j].LineCount
	})

	return scripts[0]
}