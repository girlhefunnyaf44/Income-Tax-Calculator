package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var max10 float64 = 9875
	var max12 float64 = 40125
	var max22 float64 = 85525
	var max24 float64 = 163300
	var max32 float64 = 207350
	var max35 float64 = 518400
	var tier10_tax float64 = max10 * .1
	var tier12_tax float64 = tier10_tax + ((max12 - max10) * .12)
	var tier22_tax float64 = tier12_tax + ((max22 - max12) * .22)
	var tier24_tax float64 = tier22_tax + ((max24 - max22) * .24)
	var tier32_tax float64 = tier24_tax + ((max32 - max24) * .32)
	var tier35_tax float64 = tier32_tax + ((max35 - max32) * .35)
	var gross, dependents string
	for {
		fmt.Print("Enter your gross income: ")
		fmt.Scan(&gross)

		troll := strings.Contains(gross, "-")
		x := strings.Contains(gross, ".")
		y := len(gross)
		z, err := regexp.MatchString("[a-zA-Z]+", gross)
		if err != nil {
			fmt.Println("Error occurred:", err)
		}

		if (z == true) || (y > 10) || (troll == true) || (x == false) {
			fmt.Println("Please enter a floating point")
			continue
		} else {
			fmt.Println("Your Gross Income is : " + gross)
			break
		}
	}

	fmt.Print("How many dependents do you have? ")
	fmt.Scan(&dependents)
	foo := len(dependents)
	if (foo < 0) || (foo > 8) {
		panic(1)
	}

	income, err := strconv.ParseFloat(gross, 64)
	depend, err := strconv.Atoi(dependents)
	if err != nil {
		fmt.Println(err)
		return
	} else if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("t1: %T\n", income)
	var taxableIncome float64
	var taxDue float64
	taxableIncome = income - 12062 - (2000 * float64(depend))

	
	if taxableIncome <= 0 {
		taxDue = 0
	 } else if taxableIncome <= max10 {
		taxDue = taxableIncome * 0.1
	 } else if taxableIncome <= max12 {
		taxDue = tier10_tax + ((taxableIncome - max10) * .12)
	 } else if taxableIncome <= max22 {
		taxDue = tier12_tax + ((taxableIncome - max12) * .22)
	 } else if taxableIncome <= max24 {
		taxDue = tier22_tax + ((taxableIncome - max22) * .24)
	 } else if taxableIncome <= max32 {
		taxDue = tier24_tax + ((taxableIncome - max24) * .32)
	 } else if taxableIncome <= max35 {
		taxDue = tier32_tax + ((taxableIncome - max32) * .35)
	 } else if taxableIncome > max35 {
		taxDue = tier35_tax + ((taxableIncome - max35) * .37)
	 }
	 fmt.Print("Your gross income is: $")
	 fmt.Println(income)
	 fmt.Print("Your number of dependents is: ")
	 fmt.Println(depend)
	 fmt.Print("Your taxable income is: $")
	 fmt.Println(taxableIncome)
	 fmt.Print("Your tax due is: $")
	 fmt.Println(int(taxDue))


}
