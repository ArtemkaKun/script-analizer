package main

import (
	"flag"
	"fmt"
	"log"
)

var projectScripts []File

func main() {
	var err error
	args := handleConsoleArguments()

	projectScripts, err = findAllScripts(args)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(prepareOutputData())
}

func handleConsoleArguments() ConsoleArguments {
	dirWithScripts := flag.String("dir", "", "Path to dir with scripts")
	extensionToSearch := flag.String("ext", "", "Scripts extension to search")
	flag.Parse()

	return ConsoleArguments{
		DirWithScripts: *dirWithScripts,
		ExtensionToSearch: "*" + *extensionToSearch,
	}
}