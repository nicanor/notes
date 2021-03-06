package notes

import (
	"fmt"
	"testing"
)

func TestNotesAreTransposedOk(t *testing.T) {
	tests := []struct {
		n string
		t int
		e string
	}{
		{n: "A", t: 1, e: "A#"},
		{n: "A", t: 2, e: "B"},
		{n: "A", t: 12, e: "A"},
		{n: "Am", t: 3, e: "Cm"},
		{n: "Bm", t: 4, e: "D#m"},
		{n: "C#m", t: 5, e: "F#m"},
		{n: "C#7", t: 5, e: "F#7"},
		{n: "D7", t: 7, e: "A7"},
		{n: "Em7", t: 8, e: "Cm7"},
		{n: "F5", t: 9, e: "D5"},
		{n: "A", t: -1, e: "G#"},
		{n: "La", t: 2, e: "Si"},
		{n: "Solm7", t: -100, e: "Re#m7"},
	}

	for _, test := range tests {
		result, err := Transpose(test.n, test.t)

		if err != nil {
			t.Errorf("%v %d: should not return error", test.n, test.t)
		}

		if result != test.e {
			t.Errorf("%v %d: %v was not equal to %v", test.n, test.t, result, test.e)
		}
	}
}

func TestWrongNoteReturnsError(t *testing.T) {
	result, err := Transpose("X", 1)
	if result != "" {
		t.Errorf("got %v, but should return an empty string", result)
	}

	if err == nil {
		t.Errorf("got %v, but should return an error", err)
	}
}

func BenchmarkTranspose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Transpose("A#m7", 12)
	}
}

func ExampleTranspositons() {
	n, _ := Transpose("A#m7", 5)
	fmt.Println(n)
	// Output: D#m7
}
