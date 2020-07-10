package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"sort"
)

func prepareOutputData() string {
	totalAmountOfScripts := fmt.Sprintf("Scripts in the project: %s\n", humanize.Comma(int64(len(projectScripts))))
	totalAmountOfLinesOfCode := fmt.Sprintf("Lines of code in the project: %s\n", humanize.Comma(int64(calcLinesOfCode())))

	var biggestScript = findBiggestFile()
	theBiggestScript := fmt.Sprintf("The biggest file in the project: %s lines of code (%s)\n",
		humanize.Comma(int64(biggestScript.LineCount)), biggestScript.Name)

	var smallestScript = findSmallestFile()
	theSmallestScript := fmt.Sprintf("The smallest file in the project: %s lines of code (%s)\n",
		humanize.Comma(int64(smallestScript.LineCount)), smallestScript.Name)

	medianScriptSize := fmt.Sprintf("Median size of script: %s lines of code\n", humanize.Comma(int64(findMedianFile())))

	return totalAmountOfScripts + totalAmountOfLinesOfCode + theBiggestScript + theSmallestScript + medianScriptSize
}

func calcLinesOfCode() (linesOfCode uint) {
	for _, script := range projectScripts {
		linesOfCode += script.LineCount
	}

	return
}

func findBiggestFile() File {
	sort.Slice(projectScripts, func(i, j int) bool {
		return projectScripts[i].LineCount > projectScripts[j].LineCount
	})

	return projectScripts[0]
}

func findSmallestFile() File {
	sort.Slice(projectScripts, func(i, j int) bool {
		return projectScripts[i].LineCount < projectScripts[j].LineCount
	})

	return projectScripts[0]
}

func findMedianFile() uint {
	medianElement := (len(projectScripts) - 1) / 2

	return projectScripts[medianElement].LineCount
}
