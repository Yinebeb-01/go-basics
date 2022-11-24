package main

import (
	"errors"
	"fmt"
	"math"
)

// function outside main

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined negetive input")
	}
	return math.Sqrt(x), nil
}

// more on func, they are built on a stack.
func f1() (r int) {
	r = f2()
	return
}
func f2() int {
	return 4
}

// variadic functon
func add(arg ...int) int {
	total := 0
	for _, v := range arg {
		total += v
	}
	return total
}
func areaCircle(x, y int, r float64) float64 {
	return math.Pi * r * r
}

// a function that return another function- closure
func evengenerator() func() uint { //if the function is used in a return value or to be assigned to avarible,
	//no need to mention a name.
	i := uint(0)               // to make 0 unit- type casting from int to uint.
	return func() (ret uint) { // the first return value is 0.
		ret = i
		i += 2
		return
	}
}

// recursion func.
func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}

func fibon(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	}
	return fibon(num-1) + fibon(num-2)
}
func fir() {
	fmt.Println("first")
}
func sec() {
	fmt.Println("second")
}

// function with a pointer: to demonstrate that a value will copied appropriatelly.
func zero(xptr *int) {
	*xptr = 0
}

// more on pointers:
func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	temp := *y // temp will intialized as a pointer and store value of y = *y
	*y = *x    // assign value of x to y
	*x = temp  // assign value stored on pointer temp which is y's inotoal vaue to pointer x.
	//*px, *py = *py, *px   //this approach seems by appying concurrency, go assign both value at a time; the change will reflect after this line goes run.. I think
	// but this is not the fact; rather it is due right-side expressions will evaluate first, then assignment goes right to left.
}

// Refactoring long-ifs into a better way:function which greets name in a particular language
var greetings = map[string]string{ // globally dcalarations need a var keyword.
	"en":  "Hello",
	"fr":  "Bounjar",
	"it":  "itali",
	"or":  "Akam",
	"amh": "ሰለመ",
}

func greet(name string) {
	for id := range greetings {
		if welcom, ok := greetings[id]; ok {
			fmt.Println(welcom, name)
		} else {
			errors.New("Language not found!")
		}
	}
}


// more on struct
type circle struct {
	x, y int
	r    float64
}

// passing variable-struct circle, by value.
func circleArea(c circle) float64 {
	return math.Pi * c.r * c.r
}

// passing variable-circle, by reference.
func circleArea_v1(c *circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *circle) circleArea_v2() float64 { //by passing a receiver btw func and funcName,
	//this will allows us to call a function using the . operator.
	//But if we can via dot, could it support definition of more than one method?
	return math.Pi * c.r * c.r

}

