//Package hamming implements the Distance method
package hamming

import (
	"errors"
)

//Distance takes two DNA strings and returns
//the hamming number.
func Distance(a, b string) (int, error) {
	hammingNumber := 0

	if len(a) != len(b) {
		return 0, errors.New("Strings are different lengths")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingNumber++
		}
	}
	return hammingNumber, nil

}
