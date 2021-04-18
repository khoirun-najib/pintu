package pintu

import (
	"fmt"
	"strings"
)

func Lexicograph() string {
	content := ReadFile("randomstring.txt")
	var result string
	array := make(map[string]int)
	for i := 0; i < len(content); i++ {
		if _, ok := array[string(content[i])]; !ok {
			array[string(content[i])] = 1
			result += string(content[i])
		} else {
			array[string(content[i])]++
		}
	}
	fmt.Println(array)
	return result
}

func SmallestLexOrder() string {
	content := ReadFile("randomstring.txt")
	var result string
	var smallest, biggest byte = content[0], content[0]
	for i := 0; i < len(content); i++ {
		// IF string does not exist
		if !strings.Contains(result, string(content[i])) {
			result += string(content[i])

			// Set Biggest & Smallest
			if content[i] < smallest {
				smallest = content[i]
			} else if content[i] > biggest {
				biggest = content[i]
			}
		} else { // IF Exist
			existPos := strings.Index(result, string(content[i]))
			smallestPos := strings.Index(result, string(smallest))
			// Find Position
			if existPos < smallestPos {
				result = strings.Replace(result, string(content[i]), "", -1)
				result += string(content[i])
			}

		}
	}
	return result
}
