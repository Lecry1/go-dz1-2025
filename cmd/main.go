package main

import (
	"fmt"

	task1 "dz1/internal/task_1"
)

func main() {
	// Я тут просто скопипастил из папки tests и за джипитишл под формат)))
	// А для task2 и task3 не делал
	// Пример 1: нет общих цифр
	a, b := 123, 456
	r1, r2, err := task1.FilterCommonDigits(a, b)
	fmt.Printf("FilterCommonDigits(%d, %d) = (%d, %d, %v), expected: (123, 456, nil)\n", a, b, r1, r2, err)

	// Пример 2: есть общие цифры
	a, b = 12345, 56789
	r1, r2, err = task1.FilterCommonDigits(a, b)
	fmt.Printf("FilterCommonDigits(%d, %d) = (%d, %d, %v), expected: (1234, 6789, nil)\n", a, b, r1, r2, err)

	// Пример 3: все цифры общие — ошибка EmptyNum
	a, b = 111, 111
	r1, r2, err = task1.FilterCommonDigits(a, b)
	fmt.Printf("FilterCommonDigits(%d, %d) = (%d, %d, %v), expected: (0, 0, ErrEmptyNum)\n", a, b, r1, r2, err)

	// Пример 4: отрицательные числа — ошибка NegNums
	a, b = -123, 456
	r1, r2, err = task1.FilterCommonDigits(a, b)
	fmt.Printf("FilterCommonDigits(%d, %d) = (%d, %d, %v), expected: (0, 0, ErrNegNums)\n", a, b, r1, r2, err)

	// Пример 5: нулевые значения
	a, b = 0, 123
	r1, r2, err = task1.FilterCommonDigits(a, b)
	fmt.Printf("FilterCommonDigits(%d, %d) = (%d, %d, %v), expected: (0, 123, nil)\n", a, b, r1, r2, err)

	// Пример 6: большие числа с общими цифрами — ошибка EmptyNum
	a, b = 987654321, 123456789
	r1, r2, err = task1.FilterCommonDigits(a, b)
	fmt.Printf("FilterCommonDigits(%d, %d) = (%d, %d, %v), expected: (0, 0, ErrEmptyNum)\n", a, b, r1, r2, err)

	// Пример 7: оба отрицательные
	a, b = -123, -456
	r1, r2, err = task1.FilterCommonDigits(a, b)
	fmt.Printf("FilterCommonDigits(%d, %d) = (%d, %d, %v), expected: (0, 0, ErrNegNums)\n", a, b, r1, r2, err)

	// Пример 8: частично общие цифры
	a, b = 123456, 456789
	r1, r2, err = task1.FilterCommonDigits(a, b)
	fmt.Printf("FilterCommonDigits(%d, %d) = (%d, %d, %v), expected: (123, 789, nil)\n", a, b, r1, r2, err)
}
