package ALLFunction

import (
	"fmt"
	"math"
	"strings"
)

func PrintSquare(n int) string {

	
	if n == 0 {
		return ""
	}

	n = int(math.Abs(float64(n)))

	var builder strings.Builder

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("#")
		}
		if i < n-1 {
			builder.WriteString("\n")
		}
		fmt.Println()
	}

	return builder.String()

	
}

func Add (a , b int) int{
	return a + b
}

func Deposit_interest(interest_rate, deposit_amout float64) float64{
	return deposit_amout * interest_rate / 100
}

func Strating (strating_amount , years , return_rate float64) float64{
	return strating_amount * math.Pow((1+return_rate/100), float64(years))
}

