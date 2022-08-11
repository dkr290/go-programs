// Package sumutil ask if you are ready to rock
package sumutil

// Sum adds an unlimited values of type int
func Sum(xi ...int) int {
	s := 0

	for _, v := range xi {
		s += v
	}
	return s
}
