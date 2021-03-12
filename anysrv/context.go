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
	params *[]*param
	req    *http.Request
	resp   http.ResponseWriter
	code   int
}

func (c *context) Path() string {
	return c.req.URL.Path
}
func (c *context) Resp() http.ResponseWriter {
	return c.resp
}
func (c *context) Param(name string) string {
	if c.params != nil {
		l := len(*c.params)
		for i := 0; i < l; i++ {
			n := (*c.params)[i]
			p := c.Path()
			parseReqPath(p)
			if n.name == name {
				return p[mkA[n.deep]:mkB[n.deep]]
			}
		}
	}
	return ""
}

func (c *context) Params() map[string]string {
	m := make(map[string]string)
	p := c.Path()
	parseReqPath(p)
	l := len(*c.params)
	for i := 0; i < l; i++ {
		n := (*c.params)[i]
		m[n.name] = p[mkA[n.deep]:mkB[n.deep]]
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
