//Package twofer implements ShareWith takes a name and returns
//One for NAME, one for me.
//If no name given returns One for you, one for me.
package twofer

import (
	"fmt"
)

//ShareWith takes a name and returns One for Name, one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
