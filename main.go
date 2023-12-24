package main

import (
    "fmt"
    "strconv"
)

func main() {
    usedCoupons := make(map[int]bool)

    for {
        fmt.Print("Enter the Coupon Number: ")
        couponString := ""
        fmt.Scanln(&couponString)

        coupon, err := strconv.Atoi(couponString)
        if err != nil {
            break
        }

        if usedCoupons[coupon] {
            fmt.Println("\x1b[31mThis has been already used :-(\x1b[0m") // Use ANSI escape codes for colors
        } else {
            usedCoupons[coupon] = true
            fmt.Println("\x1b[32mThank You! :-)\x1b[0m")
        }
    }
}



//Use code with caution. Learn more
//Key points and explanations:

//map[int]bool for HashSet: Go doesn't have a built-in HashSet, so we use a map[int]bool to efficiently track used coupons.
//strconv.Atoi for parsing: The strconv.Atoi function is used to convert the string input to an integer, similar to C#'s int.TryParse.
//Error handling: The error from strconv.Atoi is checked to break the loop if the input is not a valid integer.
//ANSI escape codes for colors: Go doesn't have a built-in ConsoleColor type, so we use ANSI escape codes to achieve similar color formatting in the console output.
//fmt.Scanln for input: The fmt.Scanln function is used to read a line of input from the console.
//Infinite loop with for {}: Go's for {} construct is used to create an infinite loop, equivalent to C#'s while (true).
//Additional notes:

//Consider using a third-party package like github.com/emirpasic/gods/sets if you need more advanced set operations.
//For complex console output formatting, explore libraries like github.com/mattn/go-colorable or github.com/fatih/color.
