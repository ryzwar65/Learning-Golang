// Ada 2 jenis yaitu Integer dan Floating Point

// Spesifik Integer

// int8 min : -128 max: 127
// int16 min : -32768 max: 32767
// int32 min : -2Milyar lebih max: 2Milyar lebih
// int64 min : -12milyar lebih max: 12milyar lebih

// Spesifik Unsigned Integer

// uint8 min : 0 max: 255
// uint16 min : 0 max: 65535
// uint32 min : 0 max: 4milyar lebih
// uint64 min : 0 max: 1milyar lebih

// Floating

// float32
// float64
// Complex biasanya digunakan untuk apps matematika yang komplex
// complex64
// complex128

// Alias Integer

// byte : uint8
// rune : int32
// int : (mengikuti Sistem Operasi) minimal int32
// uint : (mengikuti Sistem Operasi) minimal uint32

package main

import "fmt"

func main() {
	fmt.Println("Integer Satu : ", 1)
	fmt.Println("Float Satu koma 5 : ", 1.5)
}
