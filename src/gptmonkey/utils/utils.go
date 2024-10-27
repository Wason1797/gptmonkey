package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func ReadInput(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg)
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	result := strings.Replace(text, "\n", "", -1)
	return result
}

func ClampIndex(i int, n int) int {
	return int(math.Max(0, math.Min(float64(i), float64(n-1))))
}
