package main

import (
	"fmt"
)

func main() {
	accountBalance := 0.0

	for {
		welcomeText()
		var choice int
		fmt.Print(" Enter the number: ")
		fmt.Scan(&choice)

		// if choice == 1 {
			
		// } else if choice == 2 {
		// 	deposit(accountBalance)
		// 	continue
		// } else if choice == 3 {
		// 	withdraw(accountBalance)
		// 	continue
		// } else {
		// 	fmt.Println("Thank you for using FakeBank")
		// 	break
		// }

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
			continue
		case 2:
			deposit(accountBalance)
			continue
		case 3:
			withdraw(accountBalance)
			continue
		default:
			fmt.Println("Goodbye!")
			return
		}

	}
}

func deposit(accountBalance float64) {
	fmt.Print("Enter the amount you want to deposit: ")
	var amount float64
	fmt.Scan(&amount)
	accountBalance += amount
	fmt.Println("Your balance is: ", accountBalance)
}

func withdraw(accountBalance float64) {
	fmt.Print("Enter the amount you want to withdraw: ")
	var amount float64
	fmt.Scan(&amount)

	if amount > accountBalance {
		fmt.Println("Insufficient balance")
		return
	}

	accountBalance -= amount
	fmt.Println("Your balance is: ", accountBalance)
}

func welcomeText() {
	fmt.Println(`
############################
#                          #
#    Welcome to FakeBank   #
#                          #
############################
What do you want to do? 
  1. Check balance
  2. Deposit
  3. Withdraw
  4. Exit
 `)
}
