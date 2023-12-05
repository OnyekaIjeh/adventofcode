package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	// get file path from arg
	if len(os.Args) != 2 {
		check(errors.New("please pass in only one argument as the file path"))
	}

	filePath := os.Args[1]

	fmt.Println(filePath)

	// Get the directory of the current executable
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	// Join the base path of the executable with the relative file path
	fullPath := filepath.Join(basepath, filePath)

	// read the file line by line
	file, err := os.Open(fullPath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Create a channel for integers
	numbers := make(chan int)
	var wg sync.WaitGroup

	for scanner.Scan() {
		line := scanner.Text()

		wg.Add(1)

		go func(line string) {
			defer wg.Done()
			sendFirstWithLastNumberAsInt(line, numbers) // Process each line
		}(line)
	}

	check(scanner.Err())

	// Close the channel when all goroutines have finished
	go func() {
		wg.Wait()
		close(numbers)
	}()

	// Receive numbers and calculate the sum
	sum := 0
	for num := range numbers { // Range over the channel to receive values
		sum += num
	}

	log.Println(sum)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func sendFirstWithLastNumberAsInt(s string, numbers chan<- int) {
	numbersMapping := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	// Regular expression to find numbers
	re := regexp.MustCompile(`(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|([0-9])`)

	matches := re.FindAllString(s, -1)

	numberMatches := make([]string, len(matches))

	// map matches to numbers
	for i, match := range matches {
		numberMatches[i] = numbersMapping[match]
	}

	// get first and last number
	numberString := fmt.Sprint(numberMatches[0], numberMatches[len(numberMatches)-1])

	// convert to int
	number, err := strconv.Atoi(numberString)

	log.Println(number)

	check(err)

	numbers <- number
}
