package intkey

import (
	"math/rand"
	"testing"
)

// Tested on M1 Macbook pro.
//
// go test -bench=. -benchtime=10s
// goos: darwin
// goarch: arm64
// pkg: github.com/greedy52/golang-switch-slice-map-compare/intkey
// BenchmarkBaseline-8   	860302086	        13.52 ns/op
// BenchmarkSwitch1-8    	888747649	        13.50 ns/op
// BenchmarkSwitch5-8    	828560742	        14.49 ns/op
// BenchmarkSwitch10-8   	489604197	        24.49 ns/op
// BenchmarkSwitch20-8   	398235208	        30.13 ns/op
// BenchmarkSlice-8      	887899017	        13.49 ns/op
// BenchmarkMap-8        	339221132	        35.84 ns/op
//
// BenchmarkBaseline is a baseline where rand number is generated and a
// function call is made but not using any containers for searching the key.
//
// BenchmarkSwitch1/5/10/20 are using switches with 1/5/10/20 cases
// respectively.
//
// BenchmarkSlice and BenchmarkMap are using a slice and a map respectively.

const dataSize = 20

var dataMap map[int]int
var dataSlice []int

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
		key := rand.Int() % dataSize
		dataBaseline(key)
	}
}

func BenchmarkSwitch1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := rand.Int() % dataSize
		dataSwitch1(key)
	}
}

func BenchmarkSwitch5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := rand.Int() % dataSize
		dataSwitch5(key)
	}
}

func BenchmarkSwitch10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := rand.Int() % dataSize
		dataSwitch10(key)
	}
}

func BenchmarkSwitch20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := rand.Int() % dataSize
		dataSwitch20(key)
	}
}

func BenchmarkSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := rand.Int() % dataSize
		dataSliceFunc(key)
	}
}

func BenchmarkMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := rand.Int() % dataSize
		dataMapFunc(key)
	}
}

func init() {
	dataMap = make(map[int]int, dataSize)
	dataSlice = make([]int, dataSize)

	for i := 0; i < dataSize; i++ {
		dataMap[i] = i
		dataSlice[i] = i
	}
}
