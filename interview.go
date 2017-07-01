package notes

import (
	"errors"
	"strings"
)

var allNotes = [12]string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

func Transpose(n string, t int) (string, error) {
	var minor bool
	var sept bool

	if strings.HasSuffix(n, "7") {
		sept = true
		n = strings.TrimSuffix(n, "7")
	}

	if strings.HasSuffix(n, "m") {
		minor = true
		n = strings.TrimSuffix(n, "m")
	}

	var i int = indexOf(n, allNotes)
	if i == -1 {
		return "", errors.New("Non existent note")
	}

	note := allNotes[(i+t)%12]

	if minor {
		note += "m"
	}

  if sept {
  	note += "7"
  }

	return note, nil
}

func indexOf(word string, data [12]string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
