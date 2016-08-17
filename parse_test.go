
package raste

import (
	"os"
	"fmt"
	"testing"
)


func TestParseCss(t *testing.T) {
	name := "./__test__/index.less"

	ctx := NewContext(ModeCss)
	reader, err := os.Open(name)
	if err != nil {
		fmt.Println("file not exists")
		return
	}
	defer reader.Close()

	output := Parse(ctx, reader)
	fmt.Println(output)
}
