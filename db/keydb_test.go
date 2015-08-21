package keydb

import (
	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestSet(c *C) {
	d := NewDB()
	value := `{"name": "value"}`
	ok := d.Set("testKey", value)

	c.Assert(ok, Equals, false)
	ok = d.Set("testKey", value)
	c.Assert(ok, Equals, true)
}

func (s *TestSuite) TestSetShouldUpdate(c *C) {
	d := NewDB()
	value := `{"name": "value"}`
	ok := d.Set("testKey", value)

	c.Assert(ok, Equals, false)
	ok = d.Set("testKey", "value changed")
	v, _ := d.Get("testKey")
	c.Assert(v, Equals, "value changed")
	c.Assert(ok, Equals, true)

}

func (s *TestSuite) TestGet(c *C) {
	d := NewDB()
	value := `{"name": "value"}`
	d.Set("testKey", value)
	v, exist := d.Get("testKey")

	c.Assert(exist, Equals, true)
	c.Assert(v, Equals, value)
}

func (s *TestSuite) TestGetNotExist(c *C) {
	d := NewDB()
	_, exist := d.Get("testKey")

	c.Assert(exist, Equals, false)
}
