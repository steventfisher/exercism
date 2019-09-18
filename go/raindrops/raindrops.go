//Package raindrops implements the Convert method
package raindrops

//Convert takes in an integer and returns some combination
//of Pling Plang Plong, depending if the integer is divisible
//by 3, 5, and 7 respectivly.
func Convert(a int) string {
	rain := ""

	if a%3 == 0 {
		rain += "Pling"
	}
	if a%5 == 0 {
		rain += "Plang"
	}
	if a%7 == 0 {
		rain += "Plong"
	}

	return rain

}
