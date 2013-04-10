package functional

import (
	"testing"
)

func square(x int64) int64 {
	return x * x
}

func TestMap(t *testing.T) {
	tSlice := []int64{1,2,3,4,5}
	other := []int64{1,2,3,4,5}
	tSlice = Map(square, tSlice).([]int64)
	for i, v := range tSlice {
		if square(other[i]) != v {
			t.Errorf("Testing map using square func, int %d should be %d", v, square(other[i]))
		}
	}
}
