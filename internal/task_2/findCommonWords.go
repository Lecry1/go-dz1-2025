package task2

import (
	"fmt"
	"os"
	"bufio"
)

func FindCommonWords(outputFilename string, inputFilenames ...string) error {

	if len(inputFilenames) == 0{
		return  nil
	}
	// var m map[string]int  расширение vs code заставляет такой вид объявления переписывать, окак
	// m = make(map[string]int)
	m := make(map[string]int)
	for _, filename := range inputFilenames {
		file, err := os.Open(filename)
		if err != nil {
			return ErrOpenFile
		}
		defer file.Close()
		
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		
		for scanner.Scan() {
			word := scanner.Text()
			m[word]++
		}
		
		if err := scanner.Err(); err != nil {
			return err
		}
	}

	fileOut, err := os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil{
		return ErrOpenFile
	}
	defer fileOut.Close()


	for key, value := range m {
		if len(inputFilenames) == value  {
			fmt.Fprint(fileOut ,key)
		}
	}
	

	return nil
}
