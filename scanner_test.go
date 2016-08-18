
package raste

import (
	"os"
	"fmt"
	"testing"
	"github.com/towry/raste-go/scanner"
)

// name: less file name
func readLessFile(name string) {
	reader, err := os.Open(name)
	if err != nil {
		fmt.Println("file not exists")
		return
	}
	defer reader.Close()

	got := ""
	s := new(scanner.Scanner).Init(reader)

	for {
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}

		ch := s.TokenText()

		fmt.Print(ch)

		// got += ch;
	}

	// fmt.Println(got)
	_ = got
}

func TestReadLessFile(t *testing.T) {
	readLessFile("./__test__/index.html")
}
