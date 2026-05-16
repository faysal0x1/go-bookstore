package routes

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type MiddlewareFunc func(http.Handler) http.Handler

type RouteGroup struct {
	router      *mux.Router
	middlewares []mux.MiddlewareFunc
	registry    map[string]interface{} // Can be MiddlewareFunc or func(param string) MiddlewareFunc
}

func NewRouteGroup(router *mux.Router, registry map[string]interface{}) *RouteGroup {
	return &RouteGroup{
		router:   router,
		registry: registry,
	}
}

func (rg *RouteGroup) Middleware(names ...string) *RouteGroup {
	newMiddlewares := make([]mux.MiddlewareFunc, len(rg.middlewares))
	copy(newMiddlewares, rg.middlewares)

	for _, name := range names {
		mw, params := rg.parseMiddlewareName(name)
		if mw != nil {
			if params != "" {
				// Handle parameterized middleware like "role:Admin"
				if factory, ok := mw.(func(string) MiddlewareFunc); ok {
					newMiddlewares = append(newMiddlewares, mux.MiddlewareFunc(factory(params)))
				}
			} else {
				// Handle standard middleware like "auth"
				if handler, ok := mw.(MiddlewareFunc); ok {
					newMiddlewares = append(newMiddlewares, mux.MiddlewareFunc(handler))
				}
			}
		}
	}

	return &RouteGroup{
		router:      rg.router,
		middlewares: newMiddlewares,
		registry:    rg.registry,
	}
}

func (rg *RouteGroup) parseMiddlewareName(name string) (interface{}, string) {
	parts := strings.SplitN(name, ":", 2)
	mwName := parts[0]
	params := ""
	if len(parts) > 1 {
		params = parts[1]
	}

	if rg.registry != nil {
		if mw, ok := rg.registry[mwName]; ok {
			return mw, params
		}
	}
	return nil, ""
}

func (rg *RouteGroup) Prefix(prefix string) *RouteGroup {
	subrouter := rg.router.PathPrefix(prefix).Subrouter()
	for _, mw := range rg.middlewares {
		subrouter.Use(mw)
	}
	return &RouteGroup{
		router:   subrouter,
		registry: rg.registry,
	}
}

func (rg *RouteGroup) Group(fn func(rg *RouteGroup)) {
	fn(rg)
}

func (rg *RouteGroup) Handle(path string, handler http.HandlerFunc) *mux.Route {
	return rg.router.HandleFunc(path, handler)
}

func (rg *RouteGroup) Post(path string, handler http.HandlerFunc) *mux.Route {
	return rg.Handle(path, handler).Methods("POST")
}

func (rg *RouteGroup) Get(path string, handler http.HandlerFunc) *mux.Route {
	return rg.Handle(path, handler).Methods("GET")
}

func (rg *RouteGroup) Put(path string, handler http.HandlerFunc) *mux.Route {
	return rg.Handle(path, handler).Methods("PUT")
}

func (rg *RouteGroup) Delete(path string, handler http.HandlerFunc) *mux.Route {
	return rg.Handle(path, handler).Methods("DELETE")
}
