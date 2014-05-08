// Inline decoder of yenc encoded content
package yenc

import "fmt"

func Decode(data []byte) []byte {
	i, j := 0, 0
	awaitingSpecial := false

	// Very basic decoder
	for ; i < len(data); i, j = i+1, j+1 {

		//fmt.Printf("i = %v, j = %v, data[i] = '%v' '%c'", i, j, data[i], data[i])

		if awaitingSpecial {
			// escaped chars yenc42+yenc64
			data[j] = (((data[i] - 42) & 255) - 64) & 255
			awaitingSpecial = false
		} else if data[i] == '=' {
			// if escape char - then skip and backtrack j
			awaitingSpecial = true
			j--
		} else if data[i] == 10 {
			// skip line feeds
			j--
		} else if data[i] == 13 {
			// skip carriage returns
			j--
		} else {
			// normal char, yenc42
			data[j] = (data[i] - 42) & 255
		}

		fmt.Printf(" ,data[j] = '%v' %c\n", data[j], data[j])

	}

	return data[:len(data)-(i-j)]
}
