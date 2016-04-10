package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	expectedReturn, _ := reader.ReadString('\n')
	returned, _ := reader.ReadString('\n')
	fmt.Print(getFine(returned, expectedReturn))
}

func getFine(returned string, expectedReturn string) int {
	actualReturned := strings.Fields(returned)
	dexpectedReturn := strings.Fields(expectedReturn)
	dateFormat := "2 1 2006"
	expected, _ := time.Parse(dateFormat, dexpectedReturn[0]+" "+dexpectedReturn[1]+" "+dexpectedReturn[2])
	dReturned, _ := time.Parse(dateFormat, actualReturned[0]+" "+actualReturned[1]+" "+actualReturned[2])
	delta := expected.Sub(dReturned)
	if delta > 0 {
		daysDifference := int(delta.Hours() / 24)
		if dReturned.Year() == expected.Year() {
			if expected.Month() == dReturned.Month() {
				return 15 * daysDifference
			} else if expected.Month() != dReturned.Month() {
				return 500 * daysDifference
			}
		} else {
			return 10000
		}
	}
	return 0
}
