package main
//let we write a simple test code
import (
	"testing"
)
func TestHello(t *testing.T) {
	got := Hello("Chris", "es")
	want := "Hola, Chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}