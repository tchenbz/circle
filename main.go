package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func area(radius float64) (float64, error) {
	if radius < 0 {
		err := errors.New("you cannot have a negative radius")
		return -1, err
	}
	result := math.Pi * (radius * radius)
	return result, nil
}

func main() {
	radius := 5
	result, err := area(float64(radius))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}