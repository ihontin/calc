package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// upper case for text
func upper(str1 string) string {
	newStr := strings.ToUpper(str1)
	return newStr
}

// conversion rim num to int
func rimToArabic(a [10]string, x string) int {
	for ind, n := range a {
		if x == n {
			return ind + 1
		}
	}
	return 0
}

// Contains checking if an element exists in an array (Roman and Arabic numbers)
func Contains(a [10]string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// checking if an element exists in an array (sings: *, -, /, +)
func signContains(a [4]string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// conversion string to int
func conversion(strnum string) int {
	intnum, err := strconv.Atoi(strnum)
	if err != nil {
		log.Fatal(err)
	}
	return intnum
}

// calculation (a operator b)
func calculation(a, b int, oper string) int {
	var res int
	switch oper {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "/":
		res = a / b
	case "*":
		res = a * b
	}
	return res
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	rimNums := [10]string{"|", "||", "|||", "|V", "V", "V|", "V||", "V|||", "|X", "X"}
	nums := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	signs := [4]string{"+", "-", "*", "/"}

	for {
		//fmt.Println("Enter the expression or type 'e' to Exit")
		text, _ := reader.ReadString('\n')
		splittedString := strings.Split(text, " ")
		textlen := len(splittedString)
		// error, more than 3 characters sent, repeat input
		if textlen != 3 {
			//To exit input "e" or "E"
			if textlen == 1 && splittedString[0] == "e\n" || textlen == 1 && splittedString[0] == "E\n" {
				fmt.Println("Calculator is turned off.")
				break
			}
			fmt.Println("Invalid data entry!")
			continue
		}
		a := upper(strings.TrimSpace(splittedString[0]))    //clears spaces and tabs and uppercase
		b := upper(strings.TrimSpace(splittedString[2]))    //clears spaces and tabs and uppercase
		oper := upper(strings.TrimSpace(splittedString[1])) //clears spaces and tabs and uppercase
		if Contains(nums, a) && Contains(nums, b) && signContains(signs, oper) {
			a := conversion(a) // conversion to int
			b := conversion(b) // conversion to int
			res := calculation(a, b, oper)
			fmt.Println(res)
		} else if Contains(rimNums, a) && Contains(rimNums, b) && signContains(signs, oper) {
			a := rimToArabic(rimNums, a)
			b := rimToArabic(rimNums, b)
			if a <= b && oper == "-" {
				fmt.Println("There is no zero and negative numbers in the Roman system!")
				continue
			} else if a < b && oper == "/" {
				fmt.Println("There is no zero and negative numbers in the Roman system!")
				continue
			}
			res := calculation(a, b, oper)
			fmt.Println(res)
		} else {
			fmt.Println("Invalid operation format!")
		}
		//fmt.Println("("+a+")", " ("+b+") ", "("+oper+")")
	}
}