package task1

func GetDigit(num int) [10]bool {
	var digits[10] bool
	for tmp:= num; tmp > 0; tmp /= 10 {
		d := tmp % 10
		digits[d] = true
	}
	return digits
}

func NewDigit(num int, digits *[10] bool) int {
	newNum := 0
	multNew := 1
	for tmpNum := num; tmpNum > 0; tmpNum /= 10 {
		digit := tmpNum % 10
		if !(*digits)[digit] {
			newNum = multNew * digit + newNum
			multNew *= 10
		}
	}
	return newNum
}


func FilterCommonDigits(a, b int) (int, int, error){
	if a < 0 || b < 0 {
		return 0, 0, ErrNegNums
	} 
	
	// --=-=-=-=-=-=-=-=-=-

	digitA := GetDigit(a)
	digitB := GetDigit(b)

	// -=-=-=-=--=-=-=-=-=-

	var common[10] bool
	for i := range 10 {
		if digitA[i] && digitB[i] {
			common[i] = true
		}
	}
	// -=-=-=-=-=-=-=-=-=-

	newA := NewDigit(a, &common)
	newB := NewDigit(b, &common)

	if ( newA == 0 && newA != a) || ( newB == 0 && newB != b){
		return 0,0, ErrEmptyNum
	}

	return newA, newB, nil
}
