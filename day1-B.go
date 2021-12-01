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
	var m []int

	for scanner.Scan() {
		next, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		m = append(m, next)
	}

	increased := 0
	last := -1
	for i := 0; i < len(m)-2; i++ {
		if i == 0 {
			last = m[i] + m[i+1] + m[i+2]
			continue
		}

		next := m[i] + m[i+1] + m[i+2]
		if next > last {
			increased++
		}
		last = next
	}

	fmt.Println("increased: ", increased)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
