package main

import (
	"log"
	"net/http"
)

type Route struct {
	Name    string
	Path    string
	Handler func(r Response)
}

type Middleware struct {
	name           string
	fn             func(r Response) bool
	exceptionNames []string
}

func (m Middleware) mustSkip(name string) bool {
	if len(m.exceptionNames) == 0 {
		return false
	}

	for _, exceptionName := range m.exceptionNames {
		if exceptionName == name {
			return true
		}
	}

	return false
}

type Router struct {
	routes      map[string]Route
	middlewares []Middleware
}

func (router Router) registerRoute(route Route) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		response := newResponse(w, r)
		log.Println("Conected to ", route.Name)
		for _, middleware := range router.middlewares {
			mustSkip := middleware.mustSkip(route.Name)
			if !mustSkip {
				continue
			}

			if !middleware.fn(response) {
				return
			}
		}

		route.Handler(response)
	}

	http.HandleFunc(route.Path, handler)
}
