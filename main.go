package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// var text = readflie()
	// fmt.Println("Enter the text you want to write: ")
	// fmt.Scan(&text)
	// writeText(float64(text))
	// fmt.Print("Your text is: ", text)
	var accountBalance, err = getBalanceFromFile() // อ่านค่าจากไฟล์ balance.fakebank
	if err != nil{
		fmt.Println(err)
		panic(err)
	}

	for {
		welcomeText()
		var choice int
		fmt.Print(" Enter the number: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
			continue
		case 2:
			deposit(&accountBalance)
			continue
		case 3:
			withdraw(&accountBalance)
			continue
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
}

func deposit(accountBalance *float64) {
	fmt.Print("Enter the amount you want to deposit: ")
	var amount float64
	fmt.Scan(&amount)
	*accountBalance += amount
	fmt.Println("> Your Update Balance :=  ", *accountBalance) // แสดงค่า accountBalance ที่เพิ่มเข้าไป
	writeBalanceToFile(*accountBalance)                       // เขียนค่า accountBalance ลงไฟล์ balance.fakebank
}

func withdraw(accountBalance *float64) {
	fmt.Print("Enter the amount you want to withdraw: ")
	var amount float64
	fmt.Scan(&amount)

	if amount > *accountBalance {
		fmt.Println("Insufficient balance")
		return
	}

	*accountBalance -= amount
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

func writeBalanceToFile(accountBalance float64) { // เขียนค่าลงไฟล์ balance.fakebank
	balanceText := fmt.Sprint(accountBalance)                   // เอาค่า accountBalance มาเปลี่ยนเป็น string
	os.WriteFile("balance.fakebank", []byte(balanceText), 0644) // สร้างไฟล์ balance.fakebank และเขียนข้อมูลลงไป
}

func getBalanceFromFile() (accountBalance float64, e error) { // อ่านค่าจากไฟล์ balance.fakebank
	bytes, err := os.ReadFile("balance.fakebank") // อ่านไฟล์ balance.fakebank
	if err != nil {
		return 500, errors.New("cannot read balance file") // ถ้าไม่สามารถอ่านไฟล์ได้ ให้ return ค่า 500 และข้อความว่า cannot read balance file
	}
	balanceText := string(bytes)                              // แปลงข้อมูลจาก bytes เป็น string
	accountBalance, err = strconv.ParseFloat(balanceText, 64) // แปลงข้อมูลจาก string เป็น float64
	if err != nil {
		return 500, errors.New("invalid balance")
	}
	return
}

// func writeText(text float64) {
// 	textfile := fmt.Sprint(text) // เอาค่า textflie มาเปลี่ยนเป็น string
// 	os.WriteFile("eiei.eiei", []byte(textfile), 0644)
// }

// func readflie () (text float64) {
// 	bytes, _ := os.ReadFile("eiei.eiei") // อ่านไฟล์ balance.fakebank
// 	textfile := string(bytes)  
// 	text, _ = strconv.ParseFloat(textfile, 64) // แปลงข้อมูลจาก string เป็น float64
// 	return 
// }