package main

import (
	"fmt"
	"math"
	"strconv"
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func main() {
	var mealCost float64
	var tipPercent int
	var taxPercent int
	var tip float64
	var tax float64
	var totalCost string
	var tipAmount float64
	var taxAmount float64

	fmt.Scan(&mealCost)
	fmt.Scan(&tipPercent)
	fmt.Scan(&taxPercent)

	tip = float64(tipPercent) / float64(100)
	tax = float64(taxPercent) / float64(100)

	tipAmount = mealCost * tip
	taxAmount = mealCost * tax
	totalCost = strconv.Itoa(int(Round(mealCost+tipAmount+taxAmount, .5, 0)))

	fmt.Println("The total meal cost is " + totalCost + " dollars.")
}
