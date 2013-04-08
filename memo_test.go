package functional

import (
	"testing"
)

// Test function fibonacci
var fib func(n int64) int64

// Test speed of fib(100), which would take long without memo
func TestMemoTime(t *testing.T) {
	fib = func(n int64) int64 {
		if n < 2 {return 1}
		return fib(n-1) + fib(n-2)
	}
	fib = Memo(fib).(func (int64) int64)
	fib(100)
}




