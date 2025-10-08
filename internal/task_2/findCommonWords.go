package task2

import (
	"fmt"
	"os"
	"strings"
)

func FindCommonWords(outputFilename string, inputFilenames ...string) error {

	if len(inputFilenames) == 0{
		return  nil
	}
	// var m map[string]int  расширение vs code заставляет такой вид объявления переписывать, окак
	// m = make(map[string]int)
	m := make(map[string]int)
	for _, filename := range inputFilenames {
		data, err := os.ReadFile(filename) // вообще это плохое решение, весь файл в память плохо пихать
		if err != nil {
			return ErrOpenFile
		}
		text := string(data)
		for _ , word := range strings.Fields(text) {
			m[word] += 1
		}
	}
	fileOut, err := os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil{
		return ErrOpenFile
	}
	for key, value := range m {
		if len(inputFilenames) == value  {
			fmt.Fprint(fileOut ,key)
		}
	}
	defer fileOut.Close()

	return nil
}
