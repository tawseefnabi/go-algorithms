package checksum_test

import (
	"fmt"
	"github.com/tawseefnabi/go-algorithms/checksum"
	"testing"
)

var (
	CRC8 = checksum.CRCModel{0x07, 0x00, false, false, 0x00, "CRC-8"}
)

func TestCalculateCRC8(t *testing.T) {
	data := []byte("123456789")
	test := struct {
		name  string
		data  []byte
		model checksum.CRCModel
		want  uint8
	}{
		CRC8.Name, data, CRC8, 0xF4,
	}
	t.Run(CRC8.Name, func(t *testing.T) {
		fmt.Println(checksum.CRC8(test.data, test.model))
	})
}
