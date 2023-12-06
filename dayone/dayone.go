package dayone

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"sync"
)

func Run(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

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

	return sum
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func sendFirstWithLastNumberAsInt(s string, numbers chan<- int) {
	numbersMapping := map[string]string{
		"oneight":   "18",
		"twone":     "21",
		"threeight": "38",
		"fiveight":  "58",
		"sevenine":  "79",
		"eightwo":   "82",
		"eighthree": "83",
		"nineight":  "98",
		"one":       "1",
		"two":       "2",
		"three":     "3",
		"four":      "4",
		"five":      "5",
		"six":       "6",
		"seven":     "7",
		"eight":     "8",
		"nine":      "9",
		"0":         "0",
		"1":         "1",
		"2":         "2",
		"3":         "3",
		"4":         "4",
		"5":         "5",
		"6":         "6",
		"7":         "7",
		"8":         "8",
		"9":         "9",
	}

	// Regular expression to find numbers
	re := regexp.MustCompile(`(oneight)|(twone)|(threeight)|(fiveight)|(sevenine)|(eightwo)|(eighthree)|(nineight)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|([0-9])`)

	matches := re.FindAllString(s, -1)

	var numberMatches []rune

	// map matches to numbers
	for _, match := range matches {
		for _, number := range numbersMapping[match] {
			numberMatches = append(numberMatches, number)
		}
	}

	if len(numberMatches) > 0 {

		// get first and last number
		numberString := fmt.Sprint(string(numberMatches[0]), string(numberMatches[len(numberMatches)-1]))

		// convert to int
		number, err := strconv.Atoi(numberString)

		check(err)

		numbers <- number
	}

}
