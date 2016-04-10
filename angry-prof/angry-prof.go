package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	amount, _ := reader.ReadString('\n')
	amounts := strings.Fields(amount)
	tests, _ := strconv.ParseInt(amounts[0], 0, 64)
	for i := 0; i < int(tests); i++ {
		students, times := readValues(reader)
		studentInt := strings.Fields(students)
		timesInt := strings.Fields(times)
		fmt.Println(isItCancelled(studentInt, timesInt))
	}
}

func readValues(reader *bufio.Reader) (students, times string) {
	students, _ = reader.ReadString('\n')
	times, _ = reader.ReadString('\n')
	return
}

func isItCancelled(students []string, times []string) string {
	late := 0
	amount, _ := strconv.ParseInt(students[0], 10, 64)
	cancellation, _ := strconv.ParseInt(students[1], 10, 64)
	for i := int64(0); i < amount; i++ {
		result, _ := strconv.ParseInt(times[i], 10, 64)
		if result > int64(0) {
			late++
		}
	}
	if late > int(amount)-int(cancellation) {
		return "YES"
	} else {
		return "NO"
	}
}
