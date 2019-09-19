// Package proverb implements the Proverb method
package proverb

import (
	"fmt"
)

// Proverb takes a string array as input and returns the proverb string array
func Proverb(rhyme []string) []string {
	var prov []string

	if len(rhyme) == 0 {
		return []string{}
	}

	for i := 0; i < len(rhyme)-1; i++ {
		prov = append(prov, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}

	prov = append(prov, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return prov
}
