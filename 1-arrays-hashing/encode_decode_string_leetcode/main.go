package main

import (
	"fmt"
	"strconv"
)

func encode(strings []string) string {
	str := ""
	for _, s := range strings {
		str += fmt.Sprintf("%d/%s", len(s), s)
	}
	return str
}

func decode(input string) ([]string, error) {
	var strings []string
	i := 0
	for i < len(input) {
		slashIndex := i
		for slashIndex < len(input) && input[slashIndex] != '/' {
			slashIndex++
		}
		if slashIndex == len(input) {
			return nil, fmt.Errorf("invalid input: missing '/'")
		}
		strLength, err := strconv.Atoi(input[i:slashIndex])
		if err != nil {
			return nil, err
		}
		if slashIndex+1+strLength > len(input) {
			return nil, fmt.Errorf("invalid input: string length out of bounds")
		}
		str := input[slashIndex+1 : slashIndex+1+strLength]
		strings = append(strings, str)
		i = slashIndex + 1 + strLength
	}
	return strings, nil
}

func main() {
	strings := []string{
		"hello",
		"world",
		"this",
		"is",
		"a",
		"random",
		"slice",
		"of",
		"words",
		"!",
		"123456789011121314151617181920",
	}
	encoded := encode(strings)
	decoded, err := decode(encoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings)
	fmt.Println(encoded)
	fmt.Println(decoded)

	for i := range strings {
		if decoded[i] != strings[i] {
			fmt.Printf("error: %s != %s\n", decoded[i], strings[i])
		}
	}
	fmt.Println("Correct")
}
