
package raste

// Parse the css or html.
// if its the css file, it will return a css map.
func Parse(context *Context) {
	if context.mode == ModeCss {
		parseCss(context)
		return
	}

	parseHtml(context)
}

// Parse the css.
// generate a map of css class names.
func parseCss(context *Context) {

}

// Parse the html.
func parseHtml(context *Context) {

}
