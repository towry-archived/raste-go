
package raste

type Context struct {
	mode modeType 	// The mode of the context
	nameSet map[string]string	// Css class names
}

type modeType int

// modes
const (
	ModeCss modeType = iota
	ModeHtml
)

func NewContext(mode modeType) *Context {
	return &Context{
		mode: mode,
		nameSet: make(map[string]string),
	}
}

// New create a new context with the given one
func (c *Context) New() *Context {
	ctx := &Context{
		mode: c.mode,
		nameSet: c.nameSet,
	}

	return ctx
}

// Set the mode
func (c *Context) SetMode(mod modeType) *Context {
	c.mode = mod

	return c
}

func (c *Context) HasName(name string) bool {
	_, ok := c.nameSet[name]

	return ok
}

// setter
func (c *Context) SetName(name string, val string) {
	c.nameSet[name] = val
}

func (c *Context) RemoveName(name string) {
	delete(c.nameSet, name)
}

// getter
func (c *Context) GetName(name string) (string, bool) {
	if c.HasName(name) {
		return c.nameSet[name], true
	}

	return "", false
}
