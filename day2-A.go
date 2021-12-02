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
	file, err := os.Open("day2.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hpos := 0
	vpos := 0 //depth
	for scanner.Scan() {
		tokens := strings.SplitN(scanner.Text(), " ", 2)

		direction := tokens[0]
		distance, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			hpos += distance
		case "up":
			vpos -= distance
		case "down":
			vpos += distance
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", hpos*vpos)
}
