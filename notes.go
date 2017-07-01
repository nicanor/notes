// Package notes transposes musical notes.
package notes

import (
	"errors"
	"strings"
)

var allNotes = [12]string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var allSuffixes = [4]string{"m7", "7", "5", "m"}

// Transpose returns a new note, given a note and a number of semitones.
func Transpose(n string, t int) (string, error) {
	var note, suffix string = split(n)
	var i int = index(note)

	if i == -1 {
		return "", errors.New("Non existent note")
	}

	note = get(i + t)

	return note + suffix, nil
}

// get returns note by position in allNotes. If negative index, it goes backwards.
func get(i int) string {
	if i < 0 {
		i = 12 - (-i % 12)
	}

	return allNotes[i%12]
}

// index returns position of note in allNotes.
func index(n string) int {
	for i, v := range allNotes {
		if n == v {
			return i
		}
	}
	return -1
}

// split separates prefix from suffix.
func split(n string) (string, string) {
	for _, s := range allSuffixes {
		if strings.HasSuffix(n, s) {
			p := strings.TrimSuffix(n, s)
			return p, s
		}
	}
	return n, ""
}
