package passwordgen

import (
	rand "math/rand/v2"
)

func Generate(charlist []byte, length uint8) string {
	output := ""

	for range length {
		output += string(charlist[rand.IntN(len(charlist))])
	}
	return output
}
