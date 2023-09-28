package main

func main() {
	// This anonymous function has no parameters
	// but has two results.
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // Call it. No arguments are needed.

	// The following anonymous function have no results.

	func(a, b int) {
		// The following line prints: a*a + b*b = 25
		println("a*a + b*b =", a*a + b*b)
	}(x, y) // pass argument x and y to parameter a and b.

	func(x int) {
		// The parameter x shadows the outer x.
		// The following line prints: x*x + y*y = 32
		println("x*x + y*y =", x*x + y*y)
	}(y) // pass argument y to parameter x.

	func() {
		// The following line prints: x*x + y*y = 25
		println("x*x + y*y =", x*x + y*y)
	}()// no arguments are needed.

	// c is a untyped complex constant.
	const c = complex(1.6, 3.3)

	// The results of real(c) and imag(c) are both
	// untyped floating-point values. They are both
	// deduced as values of type float32 below.
	var a, b float32 = real(c), imag(c)

	// d is deduced as a typed value of type complex64.
	// The results of real(d) and imag(d) are both
	// typed values of type float32.
	var d = complex(a, b)

	// e is deduced as a typed value of type complex128.
	// The results of real(e) and imag(e) are both
	// typed values of type float64.
	var e = c
	println(a, b, d, e);
}