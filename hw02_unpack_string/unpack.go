package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inPutStr string) (string, error) {
	var outPutbug strings.Builder
	var prevRun *rune
	ignore := false

	for _, run := range inPutStr {
		if run == '\\' && !ignore {
			ignore = true
			continue
		}
		if unicode.IsDigit(run) && !ignore {
			if prevRun == nil {
				return "", ErrInvalidString
			}

			number, err := strconv.Atoi(string(run))
			if err != nil {
				return "", ErrInvalidString
			}

			outPutbug.WriteString(strings.Repeat(string(*prevRun), number))
			prevRun = nil
			continue
		}

		if prevRun != nil {
			outPutbug.WriteString(string(*prevRun))
		} else {
			prevRun = new(rune)
		}
		*prevRun = run
		ignore = false
	}
	if prevRun != nil {
		outPutbug.WriteString(string(*prevRun))
	}

	return outPutbug.String(), nil
}
