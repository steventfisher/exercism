// Package leap implements the IsLeapYear method
package leap

// IsLeapYear returns true if given year is a leap year and false otherwise
func IsLeapYear(n int) bool {
	return (n%4 == 0 && n%100 != 0) || (n%400 == 0)

}
