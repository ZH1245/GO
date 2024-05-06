// Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var startTime = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	var endTime = time.Now()
	// fmt.Printf("%s, %s\n", startTime, endTime)
	fmt.Printf("Difference: %s\n", endTime.Sub(startTime))
}
