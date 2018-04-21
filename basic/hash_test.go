package basic

//
//func TestAdler32(t *testing.T) {
//	sum := adler32.Checksum([]byte("This is a good day!"))
//	fmt.Printf("Adler32 sum : %v\n", sum)
//	calculator := adler32.New()
//	calculator.Write([]byte("This is"))
//	calculator.Write([]byte(" a good day!"))
//	fmt.Printf("Adler32 sum (mannual): %v\n", calculator.Sum32())
//}
//
//// x³²+ x³¹+ x²⁴+ x²²+ x¹⁶+ x¹⁴+ x⁸+ x⁷+ x⁵+ x³+ x¹+ x⁰
//// 0b11010101100000101000001010000001=0xD5828281
//func TestCRC32(t *testing.T) {
//	crc32q := crc32.MakeTable(0xD5828281)
//	content := []byte("hello world\x00\x00\x00\x00")
//	fmt.Printf("%08x\n", crc32.Checksum(content, crc32q))
//	content = []byte("hello world\x0E\x1C\xC6\x45")
//	fmt.Printf("%08x\n", crc32.Checksum(content, crc32q))
//}
//
//func TestCRC64(t *testing.T) {
//	img, err := imgio.Open("../1.jpg")
//	assert.Equal(t, err, nil)
//	iso := crc64.MakeTable(crc64.ISO)
//	writer := crc64.New(iso)
//	err = imgio.Encode(writer, img, imgio.JPEG)
//	assert.Equal(t, err, nil)
//	fmt.Printf("sum64 %v\n", writer.Sum64())
//
//	writer32 := crc32.NewIEEE()
//	err = imgio.Encode(writer32, img, imgio.JPEG)
//	assert.Equal(t, err, nil)
//	fmt.Printf("sum32 %v\n", writer32.Sum32())
//
//}
//
//func TestAdler32OnWriter(t *testing.T) {
//	img, err := imgio.Open("../1.jpg")
//	assert.Equal(t, err, nil)
//	writer := adler32.New()
//	err = imgio.Encode(writer, img, imgio.JPEG)
//	assert.Equal(t, err, nil)
//	fmt.Printf("adler32 sum %v\n", writer.Sum32())
//}
