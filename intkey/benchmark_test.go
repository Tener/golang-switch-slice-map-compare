package intkey

import (
	"math/rand"
	"testing"
)

// Tested on M1 Macbook pro.
//
// go test -bench=. -benchtime=30s
// goos: darwin
// goarch: arm64
// pkg: github.com/greedy52/golang-switch-slice-map-compare/intkey
// BenchmarkBaseline-10    	1000000000	         0.6536 ns/op
// BenchmarkSwitch1-10     	1000000000	         0.6528 ns/op
// BenchmarkSwitch5-10     	1000000000	         2.313 ns/op
// BenchmarkSwitch10-10    	1000000000	         3.108 ns/op
// BenchmarkSwitch20-10    	1000000000	         4.259 ns/op
// BenchmarkSlice-10       	1000000000	         1.392 ns/op
// BenchmarkMap-10         	1000000000	        25.22 ns/op
//
// BenchmarkBaseline is a baseline where rand number is generated and a
// function call is made but not using any containers for searching the key.
//
// BenchmarkSwitch1/5/10/20 are using switches with 1/5/10/20 cases
// respectively.
//
// BenchmarkSlice and BenchmarkMap are using a slice and a map respectively.

const dataSize = 250

var dataMap map[int]int
var dataSlice []int
var randArr []int

func dataBaseline(key int) int {
	return 0
}

func dataMapFunc(key int) int {
	value, ok := dataMap[key]
	if ok {
		return value
	}
	return 0
}

func dataSliceFunc(key int) int {
	if key < len(dataSlice) {
		return dataSlice[key]
	}
	return 0
}

func dataSwitch1(key int) int {
	switch key {
	case 0:
		return 0
	}
	return 0
}

func dataSwitch5(key int) int {
	switch key {
	case 0:
		return 0
	case 4:
		return 4
	case 8:
		return 8
	case 12:
		return 12
	case 16:
		return 16
	}
	return 0
}

func dataSwitch10(key int) int {
	switch key {
	case 0:
		return 0
	case 2:
		return 2
	case 4:
		return 4
	case 6:
		return 6
	case 8:
		return 8
	case 10:
		return 10
	case 12:
		return 12
	case 14:
		return 14
	case 16:
		return 16
	case 18:
		return 18
	}
	return 0
}

func dataSwitch20(key int) int {
	switch key {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return 4
	case 5:
		return 5
	case 6:
		return 6
	case 7:
		return 7
	case 8:
		return 8
	case 9:
		return 9
	case 10:
		return 10
	case 11:
		return 11
	case 12:
		return 12
	case 13:
		return 13
	case 14:
		return 14
	case 15:
		return 15
	case 16:
		return 16
	case 17:
		return 17
	case 18:
		return 18
	case 19:
		return 19
	}
	return 0
}

func BenchmarkBaseline(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := randArr[n%len(randArr)] % dataSize
		dataBaseline(key)
	}
}

func BenchmarkSwitch1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := randArr[n%len(randArr)] % dataSize
		dataSwitch1(key)
	}
}

func BenchmarkSwitch5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := randArr[n%len(randArr)] % dataSize
		dataSwitch5(key)
	}
}

func BenchmarkSwitch10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := randArr[n%len(randArr)] % dataSize
		dataSwitch10(key)
	}
}

func BenchmarkSwitch20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := randArr[n%len(randArr)] % dataSize
		dataSwitch20(key)
	}
}

func BenchmarkSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := randArr[n%len(randArr)] % dataSize
		dataSliceFunc(key)
	}
}

func BenchmarkMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := randArr[n%len(randArr)] % dataSize
		dataMapFunc(key)
	}
}

func init() {
	dataMap = make(map[int]int, dataSize)
	dataSlice = make([]int, dataSize)

	randArr = make([]int, 1000*1000)
	for i := 0; i < len(randArr); i++ {
		randArr[i] = rand.Int() % dataSize
	}

	for i := 0; i < dataSize; i++ {
		dataMap[i] = i
		dataSlice[i] = i
	}
}
