package remotecollaber_test

import (
	"testing"

	"github.com/kidoda/remotecollaber"
)
// Wont pass atm; requires user input to pick hw option.
func TestGetDevNum(t testing.T) {
	var n int
	v = getDevNum()
	if v.Type != int {
		fmt.Println("Expected int got %s", v.Type)
	}

}
