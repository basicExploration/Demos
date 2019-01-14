package salt

import (
	"fmt"
	"testing"
)

var e encryptValue
func TestSalt(t *testing.T) {
	e.saltMethod()
	fmt.Println(e.salt)
}
