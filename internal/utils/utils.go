package utils

import (
	"strconv"
)

/*
Converts an hexadecimal value in string format to
byte array
*/
func HexToByteArray(hex string) ([]uint8, error) {
	bits := make([]uint8, len(hex)*4)
	for i, c := range hex {
		n, err := strconv.ParseUint(string(c), 16, 4)
		if err != nil {
			return nil, err
		}
		bits[i*4] = uint8(n & 8 >> 3)
		bits[i*4+1] = uint8(n & 4 >> 2)
		bits[i*4+2] = uint8(n & 2 >> 1)
		bits[i*4+3] = uint8(n & 1)
	}
	return bits, nil
}

/*
Converts a boolean to integer, where true is equal to one
and false is equal to zero
*/
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
