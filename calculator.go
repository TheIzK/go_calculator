package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readNumber(instructions string) int {
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Println(instructions)
	number, err := consolereader.ReadString('\n')
	check(err)
    number = strings.Replace(number, "\n", "", -1)

	numToReturn, err := strconv.Atoi(number)
	check(err)

	return numToReturn
}

func calculationType() string {
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Println("add, subtract, divide or multiply?")
	i, err := consolereader.ReadString('\n')
	check(err)
	i = strings.Replace(i, "\n", "", -1)
	return i
}

func calculate(c string, i1 int, i2 int) string {
    switch {
    	case c == "add":
    		sum := i1 + i2
    		return strconv.Itoa(sum)
    	case c == "subtract":
			difference := i1 - i2
        	return strconv.Itoa(difference)
    	case c == "multiply":
			product := i1 * i2
        	return strconv.Itoa(product)
    	case c == "divide":
			quotient := i1 / i2
        	return strconv.Itoa(quotient)
    }
    return "Invalid"
}

func total(s string) string {
	if s == "Invalid" {
	final := "Invalid entry. Exiting program"
	return final
	}
	final := "The total is " + s
	return final
}

func check(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

func main() {
	num1 := readNumber("Input first number : ")
	num2 := readNumber("Input second number : ")
	calculationType := calculationType()
	calculation := calculate(calculationType, num1, num2)
	finalTotal := total(calculation)

	fmt.Println(finalTotal)
}
