package talk

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello go!!"
	if got := SayHello(); got != want {
		t.Errorf("SayHello() = %q, want %q: ", got, want)
	}
}

func TestCurse(t *testing.T) {
	want := "Fuck off!!"
	if got := Curse(); got != want {
		t.Errorf("Curse() = %q, want %q: ", got, want)
	}
}