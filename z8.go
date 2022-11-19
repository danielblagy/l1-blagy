package main

import "fmt"

func main() {
	var data int64 = 452323
	fmt.Printf("%b\n", data)

	modifiedData := setBit(data, 0, false)
	modifiedData = setBit(modifiedData, 1, false)
	modifiedData = setBit(modifiedData, 2, true)
	modifiedData = setBit(modifiedData, 3, true)
	fmt.Printf("%b\n", modifiedData)
}

func setBit(value int64, pos int, bit bool) int64 {
	if bit {
		// shift the number 1 the specified number of spaces in the integer,
		// then OR it with the original input
		return value | (1 << pos)
	} else {
		// shift the number 1 the specified number of spaces in the integer,
		// then flip every bit in the mask with the ^ operator,
		// finally use a bitwise AND, which doesn't touch the numbers AND'ed with 1,
		// but which will unset the value in the mask which is set to 0.
		return value & ^(1 << pos)
	}
}
