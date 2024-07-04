package main

import "fmt"

func main() {
	for {
		fmt.Println(" Grade : F")
		break
	}

	for i := 0; i < 30; i++ { // ลูปจะข้าม 3 ออกไป
		if i%3 == 0{ // ถ้าหาร 3 ลงตัว ข้ามไป
			continue // ข้ามไป
		}
		fmt.Println(i) // แสดงผล
	}

	for i := 0; i < 30; i++ { // ลูปจะหยุดที่ 9 เพราะมี break
		if i%3 == 0{ // ถ้าหาร 3 ลงตัว หยุดที่ 9
			break // หยุดที่ 9
		}
		fmt.Println(i) // แสดงผล
	}

	num := 10 // กำหนดค่า num เป็น 10

	for num > 0 { // ถ้า num มากกว่า 0
		fmt.Println(num) // แสดงผล
		num-- // ลบค่า num ลง 1
	}




	switch num {
	case 10:
		fmt.Println(" Grade : A")
	case 9:
		fmt.Println(" Grade : B")
	default:
		fmt.Println(" Grade : F")
	}

	start:
		fmt.Println("Dead")
		num--
		if num > 0 {
			goto start
		}

}