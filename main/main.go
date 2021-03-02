package main

import (
	"anysrv"
	_ "net/http/pprof"
)

func main() {
	app := anysrv.New()
	app.Get("/", func(c anysrv.Context) {
		c.String("hello world")
	})
	app.Get("/:d/:a/:b:/c:/:e", func(c anysrv.Context) {
		c.String(c.Param("e"))
	})
	app.Run("127.0.0.1", 8088)
}
