package arc

import "net/http"

type Handler func(c *Context) error

type Router struct {
	Routes []*Route
}

type Route struct {
	Method  string
	Path    string
	Handler Handler
}

func (r *Router) Run(port string) {
	http.ListenAndServe(port, r)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.Routes {
		if route.Method == req.Method && route.Path == req.URL.Path {
			route.Handler(newContext(w, req))
		}
	}
}

func New() *Router {
	return &Router{
		Routes: make([]*Route, 0),
	}
}

func (r *Router) addRoute(method, path string, handler Handler) {
	r.Routes = append(r.Routes, &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

func (r *Router) GET(path string, handler Handler) {
	r.addRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler Handler) {
	r.addRoute("POST", path, handler)
}

func (r *Router) PUT(path string, handler Handler) {
	r.addRoute("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler Handler) {
	r.addRoute("DELETE", path, handler)
}
