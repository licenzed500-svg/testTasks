package main

import (
	"fmt"
)

func trueDeclination(computersCount *int) string {
	var (
		lastDigit     int = *computersCount % 10
		lastTwoDigits int = *computersCount % 100
		buffer        string
	)
	if lastDigit == 0 && lastTwoDigits != 11 {
		buffer = fmt.Sprintf("%d компьютер", *computersCount)
	} else if lastDigit >= 2 && lastDigit <= 4 && (lastTwoDigits < 10 || lastTwoDigits >= 20) {
		buffer = fmt.Sprintf("%d компьютера", *computersCount)
	} else {
		buffer = fmt.Sprintf("%d компьютеров", *computersCount)
	}
	return buffer
}

func getComputersCount() int {

	var (
		computersCount int
	)

	fmt.Print("enter computer count:")
	fmt.Scan(&computersCount)
	return computersCount
}

func main() {

	var (
		computersCount int = getComputersCount()
	)

	fmt.Println(trueDeclination(&computersCount))
}
