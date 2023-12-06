package main

import (
	"adventofcode/dayone"
	_ "embed"
	"log"
	"strings"
)

//go:embed dayone/input.txt
var dayOneInput string

func main() {
	sum := dayone.Run(strings.NewReader(dayOneInput))

	log.Println(sum)
}
