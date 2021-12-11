package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    readFile, err := os.Open("1input")

    if err != nil {
        log.Fatalf("Failed to open file; %s", err)
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileTextLines []string

    for fileScanner.Scan() {
        fileTextLines = append(fileTextLines, fileScanner.Text())
    }

    readFile.Close()

    prevValue := -1
    numIncreases := 0
    for _, eachline := range fileTextLines {
        convertedValue, err := strconv.Atoi(eachline)
        if err != nil {
            log.Fatal(err)
        } else {
            if convertedValue > prevValue && prevValue != -1 {
                numIncreases++
            }
            prevValue = convertedValue
        }
    }
    fmt.Println(numIncreases)
}
