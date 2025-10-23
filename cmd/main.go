package main

import (
	"fmt"
	"os"
	task1 "dz1/internal/task_1"
	task2 "dz1/internal/task_2"
	task3 "dz1/internal/task_3"
)

func main() {
	// Я тут просто скопипастил из папки tests и загптшил под формат)))

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

	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

	dir := "./internal/task_2/files"

	// --- Пример вызова 1: обычный случай ---
	output := dir + "/result1.txt"
	_ = os.Remove(output)
	err = task2.FindCommonWords(output, dir+"/file1.txt", dir+"/file2.txt", dir+"/file3.txt")
	data, _ := os.ReadFile(output)
	content := string(data)
	fmt.Printf("FindCommonWords(file1.txt, file2.txt, file3.txt) = (%q, %v), expected: (contains \"world\", \"go\", nil)\n", content, err)

	// --- Пример вызова 2: файлы без общих слов ---
	output = dir + "/result2.txt"
	_ = os.Remove(output)
	err = task2.FindCommonWords(output, dir+"/no_common1.txt", dir+"/no_common2.txt")
	data, _ = os.ReadFile(output)
	content = string(data)
	fmt.Printf("FindCommonWords(no_common1.txt, no_common2.txt) = (%q, %v), expected: (\"\", nil)\n", content, err)

	// --- Пример вызова 3: один файл ---
	output = dir + "/result3.txt"
	_ = os.Remove(output)
	err = task2.FindCommonWords(output, dir+"/single.txt")
	data, _ = os.ReadFile(output)
	content = string(data)
	fmt.Printf("FindCommonWords(single.txt) = (%q, %v), expected: (contains \"single\", \"file\", \"example\", \"words\", nil)\n", content, err)

	// --- Пример вызова 4: пустые файлы ---
	output = dir + "/result4.txt"
	_ = os.Remove(output)
	err = task2.FindCommonWords(output, dir+"/empty1.txt", dir+"/empty2.txt")
	data, _ = os.ReadFile(output)
	content = string(data)
	fmt.Printf("FindCommonWords(empty1.txt, empty2.txt) = (%q, %v), expected: (\"\", nil)\n", content, err)

	// --- Пример вызова 5: одинаковое содержимое ---
	output = dir + "/result5.txt"
	_ = os.Remove(output)
	err = task2.FindCommonWords(output, dir+"/same1.txt", dir+"/same2.txt")
	data, _ = os.ReadFile(output)
	content = string(data)
	fmt.Printf("FindCommonWords(same1.txt, same2.txt) = (%q, %v), expected: (contains \"apple\", \"banana\", \"cherry\", nil)\n", content, err)

	// --- Пример вызова 6: учет регистра ---
	output = dir + "/result6.txt"
	_ = os.Remove(output)
	err = task2.FindCommonWords(output, dir+"/case1.txt", dir+"/case2.txt")
	data, _ = os.ReadFile(output)
	content = string(data)
	fmt.Printf("FindCommonWords(case1.txt, case2.txt) = (%q, %v), expected: (\"\", nil)\n", content, err)

	// --- Пример вызова 7: несуществующий файл ---
	output = dir + "/result7.txt"
	_ = os.Remove(output)
	err = task2.FindCommonWords(output, dir+"/file1.txt", dir+"/nonexistent.txt")
	fmt.Printf("FindCommonWords(file1.txt, nonexistent.txt) = (n/a, %v), expected: (error)\n", err)

	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

	slice := []int{1, 2, 3}
	err = task3.ScaleSlice(&slice, 3)
	fmt.Printf("ScaleSlice([1,2,3], 3) = %v, %v, expected: [1 2 3 1 2 3 1 2 3], nil\n", slice, err)

	// --- Пример вызова 2: коэффициент 0 - очистка среза ---
	slice = []int{1, 2, 3}
	err = task3.ScaleSlice(&slice, 0)
	fmt.Printf("ScaleSlice([1,2,3], 0) = %v, %v, expected: [], nil\n", slice, err)

	// --- Пример вызова 3: коэффициент 1 - без изменений ---
	slice = []int{1, 2, 3}
	err = task3.ScaleSlice(&slice, 1)
	fmt.Printf("ScaleSlice([1,2,3], 1) = %v, %v, expected: [1 2 3], nil\n", slice, err)

	// --- Пример вызова 4: пустой срез ---
	slice = []int{}
	err = task3.ScaleSlice(&slice, 5)
	fmt.Printf("ScaleSlice([], 5) = %v, %v, expected: [], nil\n", slice, err)

	// --- Пример вызова 5: nil срез ---
	var nilSlice []int = nil
	err = task3.ScaleSlice(&nilSlice, 5)
	fmt.Printf("ScaleSlice(nil, 5) = nil, %v, expected: nil, nil\n", err)

	// --- Пример вызова 6: переполнение (ошибка) ---
	slice = make([]int, 1000000) // 1,000,000 элементов
	err = task3.ScaleSlice(&slice, 5000) // слишком большой коэффициент
	fmt.Printf("ScaleSlice([1_000_000 элементов], 5000) = n/a, %v, expected: n/a, ErrOverflow\n", err)
}


