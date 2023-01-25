package main

import (
	"bufio"
	"errors"
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
	rimDozens := [11]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}
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
				err := errors.New("invalid data entry")
				fmt.Print(err)
				//fmt.Println("Calculator is turned off.")
				break
			} else {
				err := errors.New("invalid data entry")
				fmt.Print(err)
				//fmt.Println("Invalid data entry!")
				//continue
				break
			}
		}
		a := upper(strings.TrimSpace(splittedString[0]))    //clears spaces and tabs and uppercase
		b := upper(strings.TrimSpace(splittedString[2]))    //clears spaces and tabs and uppercase
		oper := upper(strings.TrimSpace(splittedString[1])) //clears spaces and tabs and uppercase
		// if Arabic nums are entered
		if Contains(nums, a) && Contains(nums, b) && signContains(signs, oper) {
			a := conversion(a) // conversion to int
			b := conversion(b) // conversion to int
			res := calculation(a, b, oper)
			fmt.Println(res)
			//or if Roman nums are entered
			break
		} else if Contains(rimNums, a) && Contains(rimNums, b) && signContains(signs, oper) {
			a := rimToArabic(rimNums, a)
			b := rimToArabic(rimNums, b)
			if a <= b && oper == "-" {
				err := errors.New("there is no zero and negative numbers in the Roman system")
				fmt.Print(err)
				//fmt.Println("There is no zero and negative numbers in the Roman system!")
				//continue
				break
			} else if a < b && oper == "/" {
				err := errors.New("there is no zero and negative numbers in the Roman system")
				fmt.Print(err)
				//fmt.Println("There is no zero and negative numbers in the Roman system!")
				//continue
				break
			}
			res := calculation(a, b, oper)
			if res%10 == 0 { //if there are no units, do not show them
				fmt.Println(rimDozens[res/10])
				break
			} else { // plus units
				fmt.Println(rimDozens[res/10] + rimNums[res%10-1])
				break
			}
		} else { //if something is wrong with the data entry
			err := errors.New("invalid operation format")
			fmt.Print(err)
			//fmt.Println("Invalid operation format!")
			break
		}
	}
}
