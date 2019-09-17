//ShareWith takes a name and returns One for NAME, one for me.
//If no name given returns One for you, one for me.
package twofer


func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}
