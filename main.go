package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	const filename = "text.txt"
	findWords(filename)
}

func findWords(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("помилка читання документу: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	pattern := regexp.MustCompile(`^[бвгґджзклмнпрстфхцчшщБВГҐДЖЗКЛМНПРСТФХЦЧШЩ][а-яА-Я]*[бвгґджзклмнпрстфхцчшщБВГҐДЖЗКЛМНПРСТФХЦЧШЩ]$`)
	for scanner.Scan() {
		word := scanner.Text()
		matches := pattern.FindAllString(word, -1)
		for _, match := range matches {
			fmt.Println("Слово починається та закінчується приголосною:", match)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("помилка сканування документу: %v", err)
	}

	return nil
}
