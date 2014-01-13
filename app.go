package ape

import (
	"github.com/rosylilly/ape/encoder"
)

var (
	verbsGet     = []string{VerbGet, VerbHead}
	verbsPost    = []string{VerbPost}
	verbsPut     = []string{VerbPut, VerbPatch}
	verbsDelete  = []string{VerbDelete}
	verbsHead    = []string{VerbHead}
	verbsOptions = []string{VerbOptions}
	verbsTrace   = []string{VerbTrace}
	verbsAll     = []string{
		VerbDelete, VerbGet, VerbHead, VerbOptions, VerbPatch,
		VerbPost, VerbPut, VerbTrace,
	}
)

type App struct {
	Encoders       map[string]Encoder
	DefaultFormat  string
	ContentTypes   map[string]string
	ErrorHandler   ErrorHandler
	Prefix         string
	router         *Router
	beforeHandlers []Handler
	afterHandlers  []Handler
}

func NewApp() *App {
	return &App{
		Encoders:       map[string]Encoder{"json": encoder.JSONEncoder},
		DefaultFormat:  "json",
		ContentTypes:   map[string]string{"json": "application/json; charset=utf-8"},
		ErrorHandler:   defaultErrorHandler,
		router:         NewRouter(),
		beforeHandlers: make([]Handler, 0),
		afterHandlers:  make([]Handler, 0),
	}
}

func (a *App) Get(path string, handler HandlerFunc) *Route {
	return a.router.Add(NewRoute(verbsGet, path, handler))
}

func (a *App) Post(path string, handler HandlerFunc) *Route {
	return a.router.Add(NewRoute(verbsPost, path, handler))
}

func (a *App) Put(path string, handler HandlerFunc) *Route {
	return a.router.Add(NewRoute(verbsPut, path, handler))
}

func (a *App) Delete(path string, handler HandlerFunc) *Route {
	return a.router.Add(NewRoute(verbsDelete, path, handler))
}

func (a *App) Head(path string, handler HandlerFunc) *Route {
	return a.router.Add(NewRoute(verbsHead, path, handler))
}

func (a *App) Options(path string, handler HandlerFunc) *Route {
	return a.router.Add(NewRoute(verbsOptions, path, handler))
}

func (a *App) Trace(path string, handler HandlerFunc) *Route {
	return a.router.Add(NewRoute(verbsTrace, path, handler))
}

func (a *App) Mount(prefix string, app *App) *Route {
	app.Prefix = prefix
	prefix += "/:path"
	r := a.router.Add(NewRoute(verbsAll, prefix, app))
	r.Constrain("path", ".*?")
	return r
}

func (a *App) Before(hander HandlerFunc) {
	a.beforeHandlers = append(a.beforeHandlers, hander)
}

func (a *App) After(hander HandlerFunc) {
	a.afterHandlers = append(a.afterHandlers, hander)
}
