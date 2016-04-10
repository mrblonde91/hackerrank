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
	size, _ := reader.ReadString('\n')
	count, _ := strconv.ParseInt(strings.TrimSpace(size), 0, 64)
	results, _ := reader.ReadString('\n')
	//	stringConvertToInt(results, int(count))
	addArray(results, int(count))
}

/*func stringConvertToInt(results string, size int) {
integers := strings.Fields(results)
for i := size - 1; 0 <= i; i-- {
	fmt.Printf("%s ", integers[i])
}*/

func addArray(results string, size int) {
	integers := strings.Fields(results)
	var sum int64
	for i := 0; i < size; i++ {
		curval, _ := strconv.ParseInt(integers[i], 0, 64)
		sum = sum + curval
		fmt.Println("current", curval)
	}
	fmt.Print(sum)
}
