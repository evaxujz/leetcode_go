package main

func main() {

}

// 整数反转

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
func reverse(x int) int {
	isNegative := x < 0
	if isNegative {
		x = 0 - x
	}

	digits := make([]int, 10)

	for d := x; d != 0; d /= 10 {
		digits = append(digits, d%10)
	}

	x = 0
	for i := 0; i < len(digits); i++ {
		if digits[i] != 0 {
			for j := i; j < len(digits); j++ {
				x = 10*x + digits[j]
			}
			break
		}
	}

	if isNegative {
		if x <= 0x7FFFFFFF {
			return -x
		} else {
			return 0
		}
	} else {
		if x <= 0x7FFFFFFE {
			return x
		} else {
			return 0
		}
	}
}
