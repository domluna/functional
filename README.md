Go 1.1 added a bunch of goodies in the "reflect" package to allow for much better implementations of functional patterns. Contents will include memoization wrapping, map, parallel map, reduce, etc...

<h2>1.Using the memoization wrapper</h2>

    package main
    
    import (
	    "github.com/Niessy/functional"
	    "fmt"
    )

    // Method 1, make the function a variable and reassign to itself.
    var fib func(n int64) int64
    func init() {
    	fib = func(n int64) int64 {
    		if n < 2 {return 1}
    		return fib(n-1) + fib(n-2)
    	}
    }

    // Method 2, make a regular function then assign to it.
    func fib2(n int64) int64 {
    	if n < 2 {return 1}
    	return fib(n-1) + fib(n-2)
    }

    func main() {
        // Method 1
    	fib = functional.Memo(fib).(func(int64) int64)
    	fmt.Println(fib(300))
        // Method 2
    	fibMemo := functional.Memo(fib2).(func(int64) int64)
    	fmt.Println(fibMemo(300))
    }
    

Expected output:

    9010910435053616553
    9010910435053616553
    

