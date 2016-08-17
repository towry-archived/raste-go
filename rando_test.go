
package raste

import "testing"
import "fmt"

func TestRando(t *testing.T) {
	i := 0
	for i < 20 {
		b := RandString(5)
		fmt.Println(b)
		i += 1
	}
}
