package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    readFile, err := os.Open("2input")

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

    sums := make([]int,len(fileTextLines)-2)
    for i := 0; i < len(fileTextLines)-2; i++ {
        firstValue, err1 := strconv.Atoi(fileTextLines[i])
        secondValue, err2 := strconv.Atoi(fileTextLines[i+1])
        thirdValue, err3 := strconv.Atoi(fileTextLines[i+2])
        if err1 != nil || err2 != nil || err3 != nil {
            log.Fatal(err)
        } else {
            sums[i] = firstValue + secondValue + thirdValue
        }
    }

    prevSum := -1
    numIncreases := 0
    for _, sum := range sums {
        if sum > prevSum && prevSum != -1 {
            numIncreases++
        }
        prevSum = sum
    }
    fmt.Println(numIncreases)
}
