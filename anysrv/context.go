package anysrv

import "net/http"

type Context interface {
	Path() string
	Status(int) *context
	Error(int, error)
	Write(b []byte)
	String(str string)
	Params() map[string]string
	Param(name string) string
	Resp() http.ResponseWriter
}

type context struct {
	node *node
	req  *http.Request
	resp http.ResponseWriter
	code int
}

func (c *context) Path() string {
	return c.req.URL.Path
}
func (c *context) Resp() http.ResponseWriter {
	return c.resp
}
func (c *context) Param(name string) string {
	p := c.node
	pt := c.Path()
	e := len(pt)
	if pt[e-1] == '/' {
		e--
	}
	s := e - 1
	for p.parent != nil {
		for pt[s] != '/' {
			s--
		}
		if p.path[0] == ':' && p.path[1:] == name {
			return pt[s+1 : e]
		}
		e = s
		s--
		p = p.parent
	}
	return ""
}

func (c *context) Params() map[string]string {
	m := make(map[string]string)
	p := c.node
	pt := c.Path()
	e := len(pt)
	if pt[e-1] == '/' {
		e--
	}
	s := e - 1
	for p.parent != nil {
		for pt[s] != '/' {
			s--
		}
		if p.path[0] == ':' {
			m[p.path[1:]] = pt[s+1 : e]
		}
		e = s
		p = p.parent
	}
	return m
}

func (c *context) String(str string) {
	c.resp.Write([]byte(str))
}

func (c *context) Error(code int, e error) {
	c.Status(code)
	c.Write([]byte(e.Error()))
}

func (c *context) Write(b []byte) {
	_, err := c.resp.Write(b)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
	}
}

func (c *context) Status(s int) *context {
	c.resp.WriteHeader(s)
	return c
}
