package main

import "testing"

var words = []string{
	"11111111999999993463254355complexWord.#%*@(1233)check2346234539999999911111111",
	"12345678909876543211234567890987654321",
	"qwertyuiopoiuytrewqASDFGHJKL1234567890987654321LKJHGFDSAqwertyuiopoiuytrewq",
}

func getWord(i int) string {
	return words[i%len(words)]
}

func BenchmarkFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find(getWord(i))
	}
}

func BenchmarkFindX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findx(getWord(i))
	}
}

func BenchmarkFindY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findy(getWord(i))
	}
}

func BenchmarkFindZ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findz(getWord(i))
	}
}
