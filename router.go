package ape

type Router struct {
	routes []*Route
}

func (r *Router) Add(route *Route) {
	r.routes = append(r.routes, route)
}

func (r *Router) MatchedRoutes(verb string, path string) []*Route {
	routes := make([]*Route, 0)

	for _, route := range r.routes {
		if route.Match(verb, path) {
			routes = append(routes, route)
		}
	}

	return routes
}
