package task3

import (
	"math"
)

func ScaleSlice(slice *[]int, scaleFactor uint32) error{
	if slice == nil {
		return nil
	}
	if len(*slice) * int(scaleFactor) > math.MaxUint32 {
		return ErrOverflow
	}
	if scaleFactor == 0 {
		//slice = &([]int{}) почему-то так не работает
		*slice = (*slice)[:0]
	}

	origin := make([]int, len(*slice))
	copy(origin, *slice)
	for i := 0; i < int(scaleFactor) - 1; i++ { 
		*slice = append(*slice, origin...)
	}
	return  nil
}

