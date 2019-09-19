// Package proverb implements the Proverb method
package proverb

import(
	"fmt"
)

// Proverb takes a string array as input and returns the proverb string array
func Proverb(rhyme []string) []string {
	prov [(len(rhyme)/2) + 1]string

	for i := 0; i < len(rhyme); i++ {
		prov[i] = fmt.Sprintf("For want of a %s the %s was lost", rhyme[i], rhyme[i+1])
	}

	append(rhyme, fmt.Sprintf("And all for the want of a %s", prov[0]))

	return prov
}
