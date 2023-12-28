package main

import (
	"fmt"
	"log"
)

func main() {
	errs := evaluate(15)
	log.Println(errs)
}

func evaluate(x int) []error {
	aggErr := aggregateErrors()

	if x > 10 {
		err := fmt.Errorf("%d is over 10\n", x)
		aggErr(err)
	}

	if x > 20 {
		err := fmt.Errorf("%d is over 20\n", x)
		aggErr(err)
	}

	return aggErr(nil)
}

func aggregateErrors() func(err error) []error { // second-order function
	var errs []error

	return func(err error) []error { // closure is here, function going to return
		if err != nil {
			errs = append(errs, err)
		}
		return errs
	}
}
