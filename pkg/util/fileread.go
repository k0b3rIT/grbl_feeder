package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(filePath string) []string {
	readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

	return fileLines
}

func WriteFile(filePath string, data []string) {
    justString := strings.Join(data,"\n")
    err := os.WriteFile(filePath, []byte(justString), 0644)   

    if err != nil {
        panic(err)
    }
}