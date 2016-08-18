
package raste

import (
	"io"
	"strings"
	"unicode"
	"fmt"
	"github.com/towry/raste-go/scanner"
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
	var preIsClass bool
	var lastTokText string

	got := ""
	s := new(scanner.Scanner).Init(in)

	for {
		// scan the next token
		tok := s.Scan()

		if tok == scanner.EOF {
			break
		}

		ch := s.TokenText()

		if !preIsClass && ch == "class" {
			preIsClass = true 
		} else if preIsClass && ch != "=" && lastTokText != "=" {
			preIsClass = false
		}

		if preIsClass && lastTokText == "=" {
			// its the class names
			got += replaceClassName(context, ch)
		} else {
			got += ch;
		}

		lastTokText = ch
	}

	return got
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
	randStr = fmt.Sprintf("%s%d", randStr, len(name))

	first := rune(randStr[0])
	// Prevent first character is number
	if unicode.IsNumber(first) {
		randStr = "_" + randStr
	}

	context.SetName(name, randStr)

	return randStr
}

// Replace the css class names.
// names is string like "a-class anotherClass"
func replaceClassName(context *Context, names string) string {
	names = strings.Replace(names, "\"", "", 2)
	got := "'"
	// Split it by white space
	list := strings.Split(names, " ");

	for _, name := range list {
		if name[0] != '_' {
			got += name 
			continue
		}

		// The class name is prefixed with a underscore,
		// look it up in the name set.
		nameInSet, ok := context.GetName(name)
		if !ok {
			// If it is not in the name set, 
			// we just leave it alone.
			got += " "
			got += name
		} else {
			got += " "
			got += nameInSet
		}
	}

	got += "'"

	return got
}
