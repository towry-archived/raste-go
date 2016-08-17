
package raste

import (
	"os"
	"fmt"
	"testing"
	"text/scanner"
)

// name: less file name
func readLessFile(name string) {
	reader, err := os.Open(name)
	if err != nil {
		fmt.Println("file not exists")
		return
	}
	defer reader.Close()

	// got := ""
	s := new(scanner.Scanner).Init(reader)

	for {
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}

		ch := s.TokenText()

		if tok == scanner.Int || tok == scanner.Float {
			fmt.Println("got a number", ch)
		} else {
			fmt.Println(ch)
		}
	}

	// fmt.Println(got)
}

func TestReadLessFile(t *testing.T) {
	readLessFile("./__test__/index.less")
}
