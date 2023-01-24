// crc8.go
// description: Calculate CRC8
// details:
// A cyclic redundancy check (CRC) is an error-detecting code commonly used in digital networks
// and storage devices to detect accidental changes to raw data.
// See more [CRC](https://en.wikipedia.org/wiki/Cyclic_redundancy_check)

package checksum

import (
	"fmt"
)

// CRCModel contains the necessary parameters for calculating the DRC algorithm
type CRCModel struct {
	Poly   uint8
	Init   uint8
	RefIn  bool
	RefOut bool
	XorOut uint8
	Name   string
}

// CRC8 calculates CRC8 checksum of the given data.

func CRC8(data []byte, model CRCModel) uint8 {
	table := getTable(model)
	fmt.Println("tt", table)
	return 12
}

func getTable(model CRCModel) []uint8 {
	table := make([]uint8, 256)
	for i := 0; i < 256; i++ {
		crc := uint8(i)
		for j := 0; j < 8; j++ {
			isSetBit := (crc & 0x80) != 0
			crc <<= 1
			if isSetBit {
				crc ^= model.Poly
			}
		}
		table[i] = crc
	}
	return table
}
