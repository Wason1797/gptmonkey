package utils

import (
	"bufio"
	"fmt"
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
