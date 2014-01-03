package ape

var (
	verbsGet     = []string{VerbGet, VerbHead}
	verbsPost    = []string{VerbPost}
	verbsPut     = []string{VerbPut, VerbPatch}
	verbsDelete  = []string{VerbDelete}
	verbsHead    = []string{VerbHead}
	verbsOptions = []string{VerbOptions}
	verbsTrace   = []string{VerbTrace}
)

type App struct {
	router *Router
}

func NewApp() *App {
	return &App{NewRouter()}
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
