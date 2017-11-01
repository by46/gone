package basic

import (
	"testing"
	"hash/adler32"
	"fmt"
	"hash/crc32"
)

func TestAdler32(t *testing.T) {
	sum := adler32.Checksum([]byte("This is a good day!"))
	fmt.Printf("Adler32 sum : %v\n", sum)
	calculator := adler32.New()
	calculator.Write([]byte("This is"))
	calculator.Write([]byte(" a good day!"))
	fmt.Printf("Adler32 sum (mannual): %v\n", calculator.Sum32())
}

// x³²+ x³¹+ x²⁴+ x²²+ x¹⁶+ x¹⁴+ x⁸+ x⁷+ x⁵+ x³+ x¹+ x⁰
// 0b11010101100000101000001010000001=0xD5828281
func TestCRC32(t *testing.T) {
	crc32q := crc32.MakeTable(0xD5828281)
	content := []byte("hello world\x00\x00\x00\x00")
	fmt.Printf("%08x\n", crc32.Checksum(content, crc32q))
	content  = []byte("hello world\x0E\x1C\xC6\x45")
	fmt.Printf("%08x\n", crc32.Checksum(content, crc32q))
}
