package task1

func FilterCommonDigits(a, b int) (int, int, error){
	if a < 0 || b < 0 {
		return 0, 0, ErrNegNums
	} 
	
	// --=-=-=-=-=-=-=-=-=-
	var digitA[10] bool
	var digitB[10] bool

	tmpA := a
	for tmpA > 0 {
		d := tmpA % 10
		digitA[d] = true
		tmpA /= 10
	}

	tmpB := b
	for tmpB > 0 {
		d := tmpB % 10
		digitB[d] = true
		tmpB /= 10
	}
	// -=-=-=-=--=-=-=-=-=-

	var common[10] bool
	for i := 0; i < 10; i++ {
		if digitA[i] && digitB[i] {
			common[i] = true
		}
	}
	// -=-=-=-=-=-=-=-=-=-

	newA := 0
	tmpA = a
	multA := 1
	for  tmpA > 0 {
		d := tmpA % 10
		if !common[d] {
			newA = multA * d + newA
			multA *= 10
		}
		tmpA /= 10
	}

	newB := 0
	tmpB = b
	multB := 1
	for tmpB > 0{
		d := tmpB % 10
		if !common[d] {
			newB += multB * d
			multB *= 10
		}
		tmpB /= 10
	}
	if ( newA == 0 && newA != a) || ( newB == 0 && newB != b){
		return 0,0, ErrEmptyNum
	}

	return newA, newB, nil
}
