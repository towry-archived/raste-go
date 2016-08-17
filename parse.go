
package raste

import (
	"io"
	"unicode"
	"text/scanner"
)

// Parse the css or html.
// if its the css file, it will return a css map.
func Parse(context *Context, in io.Reader) string {
	if context.mode == ModeCss {
		return parseCss(context, in)
	}

	return parseHtml(context, in)
}

// Parse the css.
// generate a map of css class names.
func parseCss(context *Context, in io.Reader) string {
	var preIsDot bool
	var got string

	got = ""
	s := new(scanner.Scanner).Init(in)

	for {
		// scan the next token.
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}

		ch := s.TokenText()

		// if its decimal, just skip.
		if tok == scanner.Int || tok == scanner.Float {
			got += ch
			continue
		}

		if ch == "." {
			preIsDot = true
			got += ch
			continue
		}

		// If it is a class name string,
		// and starts with a underscore.
		if preIsDot && ch[0] == '_' {
			got += transformCssName(context, ch)
			continue
		} else if preIsDot {
			preIsDot = false
		}

		got += ch
	}

	return got
}

// Parse the html.
func parseHtml(context *Context, in io.Reader) string {
	return ""
}

// Transform the css class name,
// return the transformed class name.
func transformCssName(context *Context, name string) string {
	if context.HasName(name) {
		val, _ := context.GetName(name)
		return val
	}

	randStr := RandString(4)
	// Suffix with the lenth of the css class name.
	randStr += string(len(name))

	first := rune(randStr[0])
	// Prevent first character is number
	if unicode.IsNumber(first) {
		randStr = "_" + randStr
	}

	context.SetName(name, randStr)

	return randStr
}
