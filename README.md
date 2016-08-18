raste-go
--------

## Usage

```go
package main

import "os"
import "github.com/towry/raste-go"

func main() {
	cssfiles := [2]string{"a.css", "b.css"}
	htmlfiles := [1]string{"a.html"}

	ctx := raste.NewContext(raste.ModeCss)
	// first we parse all css/less files
	reader, _ := os.Open(cssfiles[0])
	defer reader.Close()
	// get the parsed output
	output := raste.Parse(ctx, reader)
	// you write the parsed output to a file
	// ... 
	// a lot code here

	// after all the css files are parsed
	// parse the html file use the same context
	// but change the context mode to html first
	ctx.SetMode(raste.ModeHtml)

	htmlreader, _ := os.Open(htmlfiles[0])
	defer htmlreader.Close()
	// get the html parsed string
	htmloutput := raste.Parse(ctx, htmlreader)
	// you write the htmloutput to a file
}

```

After the css files is first parsed, the will be a css class names map in the 
context, you use the context to continue parse the html files, replace the css
class names.

The follow css class name is replaced:

* class name starts with a underscore, like `_nsaApp`, or `_0RandomJ3`

---

&copy; 2016 Towry Wang
