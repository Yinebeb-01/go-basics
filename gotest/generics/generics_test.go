package generics

import "testing"

func TestStack(t *testing.T) {
	// t.Run("integer stack", func(t *testing.T) {
	// 	myStackOfInts := new(StackOfInts)

	// 	// check stack is empty
	// 	AssertTrue(t, myStackOfInts.IsEmpty())

	// 	// add a thing, then check it's not empty
	// 	myStackOfInts.Push(123)
	// 	AssertFalse(t, myStackOfInts.IsEmpty())

	// 	// add another thing, pop it back again
	// 	myStackOfInts.Push(456)
	// 	value, _ := myStackOfInts.Pop()
	// 	AssertEqual(t, value, 456)
	// 	value, _ = myStackOfInts.Pop()
	// 	AssertEqual(t, value, 123)
	// 	AssertTrue(t, myStackOfInts.IsEmpty())
	// })

	// t.Run("string stack", func(t *testing.T) {
	// 	myStackOfStrings := new(StackOfStrings)

	// 	// check stack is empty
	// 	AssertTrue(t, myStackOfStrings.IsEmpty())

	// 	// add a thing, then check it's not empty
	// 	myStackOfStrings.Push("123")
	// 	AssertFalse(t, myStackOfStrings.IsEmpty())

	// 	// add another thing, pop it back again
	// 	myStackOfStrings.Push("456")
	// 	value, _ := myStackOfStrings.Pop()
	// 	AssertEqual(t, value, "456")
	// 	value, _ = myStackOfStrings.Pop()
	// 	AssertEqual(t, value, "123")
	// 	AssertTrue(t, myStackOfStrings.IsEmpty())
	// })

	// //to get out of type-safety issue, caller need to assert a type before use.
	// t.Run("interface stack dx is horrid", func(t *testing.T) {
	// 	myStackOfInts := new(StackOfInts)

	// 	myStackOfInts.Push(1)
	// 	myStackOfInts.Push(2)
	// 	firstNum, _ := myStackOfInts.Pop()
	// 	secondNum, _ := myStackOfInts.Pop()

	// 	// get our ints from out interface{}
	// 	reallyFirstNum, ok := firstNum.(int)
	// 	AssertTrue(t, ok) // need to check we definitely got an int out of the interface{}

	// 	reallySecondNum, ok := secondNum.(int)
	// 	AssertTrue(t, ok) // and again!

	// 	AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
	// })
	/*v-2*/

	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int]) // put string for stack of strings

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())

		// can get the numbers we put in as numbers, not untyped interface{}
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}

// helper functions
func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}
