package main

import (
	"bufio"
	"fmt"
	"os"
)

var characterClasses = map[string]string{
	"1": "0123456789",
	"2": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"3": "abcdefghijklmnopqrstuvwxyz",
	"4": "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"5": "0123456789abcdefghijklmnopqrstuvwxyz",
	"6": "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	"7": "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
}

func getCharacterClasses(choice string) string {
	return characterClasses[choice]
}

func generateCombinations(charClasses, currentCombination string, size int, file *os.File) {
	if size == 0 {
		file.WriteString(currentCombination + "\n")
		return
	}
	for _, char := range charClasses {
		generateCombinations(charClasses, currentCombination+string(char), size-1, file)
	}
}

func generateDictionary(filename, charClasses string, minSize, maxSize int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for size := minSize; size <= maxSize; size++ {
		generateCombinations(charClasses, "", size, file)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter the file name: ")
	filename, _ := reader.ReadString('\n')
	filename = filename[:len(filename)-1] // Removing newline character

	var charClasses string
	for charClasses == "" {
		fmt.Print(
			"1) Numbers\n" +
				"2) Capital Letters\n" +
				"3) Lowercase Letters\n" +
				"4) Numbers + Capital Letters\n" +
				"5) Numbers + Lowercase Letters\n" +
				"6) Numbers + Capital Letters + Lowercase Letters\n" +
				"7) Capital Letters + Lowercase Letters\n" +
				"Please select the character class by number: ",
		)
		choice, _ := reader.ReadString('\n')
		choice = choice[:len(choice)-1] // Removing newline character
		charClasses = getCharacterClasses(choice)
		if charClasses == "" {
			fmt.Println("Invalid choice, please try again.")
		}
	}

	fmt.Print("What is the min size of the word? ")
	var minSize int
	fmt.Scanln(&minSize)

	fmt.Print("What is the max size of the word? ")
	var maxSize int
	fmt.Scanln(&maxSize)

	generateDictionary(filename, charClasses, minSize, maxSize)
}
