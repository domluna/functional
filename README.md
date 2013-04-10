Go 1.1 added a bunch of goodies in the "reflect" package to allow for much better implementations of functional patterns. Contents will include memoization wrapping, map, reduce, etc...

Note that this is more of a play around with the relfect package type of thing, for the best performance you should probably implement the function suited to your need directly, due to the reflect overhead. These here are more for convenience and educational purposes.

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
    func fibMemo(n int64) int64 {
    	if n < 2 {return 1}
    	return fibMemo(n-1) + fibMemo(n-2)
    }

    func main() {
        // Method 1
    	fib = functional.Memo(fib).(func(int64) int64)
    	fmt.Println(fib(300))
        // Method 2
    	fibMemo := functional.Memo(fibMemo).(func(int64) int64)
    	fmt.Println(fibMemo(300))
    }
    

Expected output:

    9010910435053616553
    9010910435053616553
    
    
<h2>2.Using Map function</h2>

	
	func cube(x int64) int64 {
		return x * x * x
	}
	
	cs := []int64{1,2,3,4,5}
    cs := functional.Map(cube, cs)
    
Output:
	
	1, 8, 27, 64, 125
	
	
    

