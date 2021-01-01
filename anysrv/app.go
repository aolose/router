package anysrv

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type App struct {
	router     *router
	middleware []Middleware
	ctx        *context
}

func New() *App {
	return &App{
		router:     newRouter(8, 8),
		middleware: make([]Middleware, 0, 32),
		ctx:        &context{},
	}
}

func (app *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	app.ctx.req = req
	app.ctx.resp = res
	next, node := app.router.Lookup(req.Method, req.URL.Path)
	app.ctx.node = node
	for _, m := range app.middleware {
		next = m(next)
	}
	if next != nil {
		next(app.ctx)
	} else {
		app.ctx.Error(http.StatusNotFound, errors.New("page not found"))
	}
}
func (app *App) Add(method, path string, h Handler) {
	app.router.bind(getMethodCode(method), path, h)
}

func (app *App) Ready() {
	app.router.ready()
}
func (app *App) Run(addr string, port int) {
	app.Ready()
	fmt.Printf("Server running on: http://%s:%d", addr, port)
	http.ListenAndServe(addr+":"+strconv.Itoa(port), app)
}

func (app *App) Any(path string, h Handler) {
	for i := 0; i < 7; i++ {
		app.router.bind(i, path, h)
	}
}

func (app *App) Use(m ...Middleware) {
	l := len(app.middleware)
	c := cap(app.middleware)
	n := len(m)
	if l+n == c {
		v := make([]Middleware, l+n, (l+n)*2)
		copy(v, app.middleware)
		app.middleware = v
	}
	copy(app.middleware[l:], m)
}

func (app *App) Get(path string, h Handler) {
	app.router.bind(GET, path, h)
}

func (app *App) Post(path string, h Handler) {
	app.router.bind(POST, path, h)
}

func (app *App) Put(path string, h Handler) {
	app.router.bind(PUT, path, h)
}

func (app *App) Head(path string, h Handler) {
	app.router.bind(HEAD, path, h)
}

func (app *App) Delete(path string, h Handler) {
	app.router.bind(DELETE, path, h)
}

func (app *App) Patch(path string, h Handler) {
	app.router.bind(PATCH, path, h)
}

func (app *App) Options(path string, h Handler) {
	app.router.bind(OPTIONS, path, h)
}
