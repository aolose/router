package main

import "anysrv"

func main() {
	app := anysrv.New()
	app.Get("/", func(c anysrv.Context) {
		c.String("hello world")
	})
	app.Get("/:d", func(c anysrv.Context) {
		c.String(c.Param("d"))
	})
	app.Run("127.0.0.1", 8088)
}
