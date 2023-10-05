package helper

import (
	"bufio"
	"fmt"
	"strings"
)

func RemoveSpace(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

func ScanText(title string, scanner *bufio.Scanner) string {
	fmt.Printf(title)
	scanner.Scan()
	return scanner.Text()
}
