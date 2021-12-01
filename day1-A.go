package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	last, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	increased := 0
	for scanner.Scan() {
		next, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if next > last {
			increased = increased + 1
		}
		last = next
	}
	fmt.Println("increased: ", increased)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
