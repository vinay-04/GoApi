package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
%v, which will be replaced with the general string representation of the corresponding argument.
%T, which will be replaced with the type name or type literal of the corresponding argument.
%x, which will be replaced with the hex string representation of the corresponding argument. Note, the hex string representations for values of some kinds of types are not defined. Generally, the corresponding arguments of %x should be strings, integers, integer arrays or integer slices (arrays and slices will be explained in a later article).
%s, which will be replaced with the string representation of the corresponding argument. The corresponding argument should be a string or byte slice.
Format verb %% represents a percent sign.
*/

func main() {
	fmt.Printf("Next pseudo-random number is %v.\n", rand.Uint32());
	fmt.Println(time.Now());

	a, b := 123, "Go"
	fmt.Println("\n======\n")
	fmt.Printf("a == %v == 0x%x, b == %s\n", a, a, b)
	fmt.Printf("type of a: %T, type of b: %T\n", a, b)
	fmt.Printf("1%% 50%% 99%%\n")
}
