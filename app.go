package anysrv

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

type App struct {
	router     *router
	middleware []Middleware
	ctx        *context
	cors       []*Cors
}

var useCors bool

func New() *App {
	return &App{
		router:     newRouter(),
		middleware: make([]Middleware, 0, 32),
		ctx:        &context{},
	}
}
func fn(a, b int) bool {
	b = 1 << b
	return a&b == b
}

func (app *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	app.ctx.req = req
	app.ctx.resp = res
	path := req.URL.Path
	if useCors {
		s := 0
		e := len(app.cors)
		for i := (s + e) / 2; i >= s && i < e; {
			c := app.cors[i]
			p := c.Path
			if p > path {
				e = i
				i = (s + e) / 2
			} else if p < path {
				s = i
				i = (s + e) / 2
			} else {
				headers := res.Header()
				headers.Set("Access-Control-Allow-Origin", c.Origin)
				if req.Method == http.MethodOptions {
					if c.AllowMethods != "" {
						headers.Set("Access-Control-Allow-Methods", c.AllowMethods)
					} else {
						headers.Set("Access-Control-Allow-Methods", "*")
					}
					if c.AllowHeaders != "" {
						headers.Set("Access-Control-Allow-Headers", c.AllowHeaders)
					} else {
						headers.Set("Access-Control-Allow-Headers", "*")
					}
					res.WriteHeader(http.StatusNoContent)
					fmt.Printf("%s\t%d\n", req.URL, 204)
					return
				}
				break
			}
		}
	}
	next, params := app.router.Lookup(req.Method, &path)
	app.ctx.params = params
	for _, m := range app.middleware {
		next = m(next)
	}
	if next != nil {
		next(app.ctx)
	} else {
		app.ctx.Error(http.StatusNotFound, errors.New("page not found"))
	}
	fmt.Printf("%s\t%d\n", req.URL, app.ctx.code)
}
func (app *App) Cors(cfg *Cors) {
	for _, a := range app.cors {
		if a.Path == cfg.Path {
			return
		}
	}
	app.cors = append(app.cors, cfg)
}
func (app *App) Add(method, path string, h Handler) {
	app.router.bind(getMethodCode(method), path, h)
}
func (app *App) Ready() {
	if len(app.cors) > 0 {
		useCors = true
		sort.Slice(app.cors, func(i, j int) bool {
			return sort.StringsAreSorted([]string{
				app.cors[i].Path,
				app.cors[j].Path,
			})
		})
	}
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
