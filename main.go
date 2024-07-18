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
	if err != nil{ // ถ้ามี error
		fmt.Println(err) // แสดงข้อความ error
		panic(err) // จบโปรแกรม
	}

	for { // วนลูปเพื่อให้ผู้ใช้เลือกทำอะไรต่อ
		welcomeText() // แสดงข้อความต้อนรับ
		var choice int // ประกาศตัวแปร choice เป็น int
		fmt.Print(" Enter the number: ") // แสดงข้อความ Enter the number:
		fmt.Scan(&choice) // รับค่าจากผู้ใช้

		switch choice { // ตรวจสอบค่าที่ผู้ใช้เลือก
		case 1: // ถ้าผู้ใช้เลือก 1
			fmt.Println("Your balance is: ", accountBalance) // แสดงค่า accountBalance
			continue // วนลูปต่อไป
		case 2: // ถ้าผู้ใช้เลือก 2
			deposit(&accountBalance) // เรียกฟังก์ชัน deposit
			continue // วนลูปต่อไป
		case 3: // ถ้าผู้ใช้เลือก 3
			withdraw(&accountBalance) // เรียกฟังก์ชัน withdraw
			continue // วนลูปต่อไป
		default: // ถ้าผู้ใช้เลือก 4 หรืออื่นๆ
			fmt.Println("Goodbye!") // แสดงข้อความ Goodbye!
			return // จบโปรแกรม
		}
	}
}

func deposit(accountBalance *float64) { // ฟังก์ชัน deposit รับค่า accountBalance และเพิ่มค่าเข้าไป
	fmt.Print("Enter the amount you want to deposit: ") // แสดงข้อความ Enter the amount you want to deposit:
	var amount float64 // ประกาศตัวแปร amount เป็น float64
	fmt.Scan(&amount) // รับค่าจากผู้ใช้
	*accountBalance += amount // เพิ่มค่า amount เข้าไปใน accountBalance
	fmt.Println("> Your Update Balance :=  ", *accountBalance) // แสดงค่า accountBalance ที่เพิ่มเข้าไป
	writeBalanceToFile(*accountBalance)  // เขียนค่า accountBalance ลงไฟล์ balance.fakebank
}

func withdraw(accountBalance *float64) { // ฟังก์ชัน withdraw รับค่า accountBalance และลดค่าออก
	fmt.Print("Enter the amount you want to withdraw: ")
	var amount float64 // ประกาศตัวแปร amount เป็น float64
	fmt.Scan(&amount) // รับค่าจากผู้ใช้ 

	if amount > *accountBalance { // ถ้า amount มากกว่า accountBalance
		fmt.Println("Insufficient balance") // แสดงข้อความ Insufficient balance
		return
	}

	*accountBalance -= amount // ลดค่า amount ออกจาก accountBalance
	fmt.Println("Your balance is: ", accountBalance) // แสดงค่า accountBalance
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
	balanceText := fmt.Sprint(accountBalance) // แปลงค่า accountBalance จาก float64 เป็น string
	os.WriteFile("balance.fakebank", []byte(balanceText), 0644) // สร้างไฟล์ balance.fakebank และเขียนข้อมูลลงไป
}

func getBalanceFromFile() (accountBalance float64, e error) { // อ่านค่าจากไฟล์ balance.fakebank
	bytes, err := os.ReadFile("balance.fakebank") // อ่านไฟล์ balance.fakebank
	if err != nil { // ถ้าไม่สามารถอ่านไฟล์ได้
		return 500, errors.New("cannot read balance file") // ถ้าไม่สามารถอ่านไฟล์ได้ ให้ return ค่า 500 และข้อความว่า cannot read balance file
	}
	balanceText := string(bytes) // แปลงข้อมูลจาก byte เป็น string
	accountBalance, err = strconv.ParseFloat(balanceText, 64) // แปลงข้อมูลจาก string เป็น float64
	if err != nil { // ถ้าข้อมูลไม่ถูกต้อง
		return 500, errors.New("invalid balance") // ถ้าข้อมูลไม่ถูกต้อง ให้ return ค่า 500 และข้อความว่า invalid balance
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