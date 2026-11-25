package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {
	results, err := Solve(200)
	if err != nil {
		log.Fatal(err)
	}
	for _, expression := range results {
		fmt.Println(expression + "=200")
	}
}

func Solve(target int) ([]string, error) {
	var results []string
	var search func(int, int, int, string)

	search = func(digit, sum, number int, expression string) {
		if digit < 0 {
			if sum+number == target {
				results = append(results, expression)
			}
			return
		}

		text := strconv.Itoa(digit)
		search(digit-1, sum+number, digit, expression+"+"+text)
		search(digit-1, sum+number, -digit, expression+"-"+text)

		joined := number*10 + digit
		if number < 0 {
			joined = number*10 - digit
		}
		search(digit-1, sum, joined, expression+text)
	}

	search(8, 0, 9, "9")
	if len(results) == 0 {
		return nil, errors.New("no solutions found")
	}
	return results, nil
}
