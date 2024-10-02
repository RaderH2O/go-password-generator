package main

import (
	"bufio"
	"fmt"
	"os"
	"raderh2o/password_generator/passwordgen"
	"strconv"
)

const defaultCharacterSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+-=1234567890[]{}\"';:,<.>/? "

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a characterset (or none to use the default):")

	scanner.Scan()
	characterSet := scanner.Text()
	if characterSet == "" {
		characterSet = defaultCharacterSet
	}

	length := uint8(0)

	for {
		fmt.Println("Enter the length of the password you want to generate:")

		scanner.Scan()
		lengthInput, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Please enter a number (without spaces/extra characters)")
			continue
		}
		length = uint8(lengthInput)
		break
	}

	fmt.Println(passwordgen.Generate([]byte(characterSet), length))
}
