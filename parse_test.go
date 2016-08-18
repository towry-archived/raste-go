
package raste

import (
	"os"
	"fmt"
	"testing"
)


func TestParseCss(t *testing.T) {
	var output string
	css := "./__test__/index.less"
	html := "./__test__/index.html"

	ctx := NewContext(ModeCss)
	reader, err := os.Open(css)
	if err != nil {
		fmt.Println("file not exists")
		return
	}
	defer reader.Close()

	output = Parse(ctx, reader)
	fmt.Println(output)

	ctx.SetMode(ModeHtml)
	htmlreader, err := os.Open(html)
	if err != nil {
		fmt.Println("file not exists")
		return
	}
	defer htmlreader.Close()
	output = Parse(ctx, htmlreader)

	fmt.Println("\nHTML:\n")
	fmt.Println(output)
}
