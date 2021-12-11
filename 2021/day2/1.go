package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

    horizontalPos := 0;
    depth := 0;
    for _, eachline := range fileTextLines {
        splitline := strings.Split(eachline, " ")
        direction := splitline[0]
        distance, err := strconv.Atoi(splitline[1])
        if err != nil {
            log.Fatal(err)
        }
        if direction == "forward" {
            horizontalPos += distance
        } else if direction == "down" {
            depth += distance
        } else if direction == "up" {
            depth -= distance
        }
    }
    fmt.Println(horizontalPos * depth)
}
