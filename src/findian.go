package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Please, enter a string: ")

	var input_string string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() 
	input_string = scanner.Text()
	a := strings.ToLower(input_string)
	b := strings.TrimSpace(a)
	if strings.HasPrefix(b, "i") && strings.Contains(b[1:len(b)-1], "a") && strings.HasSuffix(b, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}