package getEvenNumbers

import (
	"strconv"
	"strings"
)

func GetEvenNumbers(array []int) string {
	var evenNumbers []int
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			evenNumbers = append(evenNumbers, array[i])
		}
	}

	for i := 0; i < len(evenNumbers); i++ {
		for j := i + 1; j < len(evenNumbers); j++ {
			if evenNumbers[i] > evenNumbers[j] {
				evenNumbers[i], evenNumbers[j] = evenNumbers[j], evenNumbers[i]
			}
		}
	}

	var evenNumbersStr []string
	for _, num := range evenNumbers {
		evenNumbersStr = append(evenNumbersStr, strconv.Itoa(num))
	}

	return strings.Join(evenNumbersStr, " - ")
}