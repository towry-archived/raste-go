
package raste

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// 	"unicode/utf8"
// )

// className is a css class name returned from the scanner
// in the html mode or css mode
type className struct {
	pos Pos 	// The starting position, in bytes
	val string 	// The value of this className
}

type Pos int
