
package raste

import "testing"

func TestContextNew(t *testing.T) {
	c := NewContext(ModeHtml)
	if c.mode != ModeHtml {
		t.Fail()
	}

	b := c.New().SetMode(ModeCss)
	if c.mode != ModeHtml && b.mode != ModeCss {
		t.Fail()
	}
}

func TestContextNameSet(t *testing.T) {
	c := NewContext(ModeHtml)

	c.nameSet["index/MainApp"] = "2b"
	c.nameSet["index/ModalComponent"] = "db"

	if c.nameSet["index/MainApp"] != "2b" {
		t.Fail()
	}

	if c.nameSet["index/ModalComponent"] != "db" {
		t.Fail()
	}

	c.nameSet["index/MainApp"] = "true"
	if c.nameSet["index/MainApp"] != "true" {
		t.Fail()
	}

	if c.HasName("index/MainApp") == false {
		t.Fail()
	}
}


func TestContextGetName(t *testing.T) {
	c := NewContext(ModeHtml)

	c.nameSet["a"] = "b"

	// test get
	if val, _ := c.GetName("a"); val != "b" {
		t.Fail()
	}

	// test get none
	name, ok := c.GetName("notExist")
	if ok != false {
		t.Fail()
	} else {
		_ = name
	}
}
