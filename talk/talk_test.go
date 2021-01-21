package talk

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello, world."
	if got := SayHello(); got != want {
		t.Errorf("SayHello() = %q, want %q: ", got, want)
	}
}

func TestCurse(t *testing.T) {
	want := curse
	if got := Curse(); got != want {
		t.Errorf("Curse() = %q, want %q: ", got, want)
	}
}
