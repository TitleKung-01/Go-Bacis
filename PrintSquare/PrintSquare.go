package PrintSquare

import (
	"fmt"
	"math"
)

func PrintSquare(n int) {
	fmt.Println("ใส่เลขที่ต้องการ :")
	fmt.Scan(&n)

	if n == 0 {
		return
	}

	n = int(math.Abs(float64(n)))

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
