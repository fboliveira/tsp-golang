package util

import (
	"fmt"
	"strings"
)

func GetLine() string {
	return strings.Repeat("-", 80)
}

func PrintLine() {
	fmt.Println(GetLine())
}

func PrintHeader() {

	PrintLine()
	fmt.Println("Travelling salesperson problem - TSP")
	fmt.Println("# GOLang version #")
	PrintLine()

}
