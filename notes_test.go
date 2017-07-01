package notes

import (
	"errors"
	"strings"
)

var allNotes = [12]string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

var allSuffixes = [4]string{"m7", "7", "5", "m"}

func Transpose(n string, t int) (string, error) {
	note, suffix := split(n)

	var i int = index(note)
	if i == -1 {
		return "", errors.New("Non existent note")
	}

	note = allNotes[(i+t)%12]
	note += suffix

	return note, nil
}

func index(n string) int {
	for i, v := range allNotes {
		if n == v {
			return i
		}
	}
	return -1
}

func split(n string) (string, string) {
	for _, s := range allSuffixes {
		if strings.HasSuffix(n, s) {
			p := strings.TrimSuffix(n, s)
			return p, s
		}
	}
	return n, ""
}
