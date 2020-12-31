package main

import "anysrv"

func main() {
	app := anysrv.New()
	app.Get("/", func(c anysrv.Context) {
		c.String("hello world")
	})
	app.Run("127.0.0.1", 8088)
}
