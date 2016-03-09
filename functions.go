package main

import (
	"strconv"
	"strings"
	"time"
)

func findStrInSlice(str string, slice []string) int {
	for i, val := range slice {
		if strings.EqualFold(str, val) {
			return i
		}
	}
	return -1
}

func getCurrentDate() int {
	currentTime := time.Now().Local()
	currentDate := currentTime.Format("20060102")
	intDate, err := strconv.Atoi(currentDate)
	checkError(err, "functions-getCurrentDate")
	return intDate
}

func isValidPrice(price float64) bool {
	return price > 0
}

func isUnit(unit string, unitOption string) bool {
	return strings.EqualFold(unit, unitOption)
}

func isMatching(a string, b string) bool {
	return strings.EqualFold(a, b)
}
