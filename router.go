package main

import "net/http"

// import "fmt"

func main() {
	r := &Router{}

	r.Route(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The Best Router!"))

	})

	r.Route("GET", `/hello/(?P<Message>\w+)`, func(w http.ResponseWriter, r *http.Request) {
		message := URLParam(r, "Message")
		w.Write([]byte("Hello " + message))
	})

	http.ListenAndServe(":8000", r)
}

// ~~~~~ Router ~~~~~ //

type Router struct {
	routes []RouteEntry
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	e := RouteEntry{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
	}
	rtr.routes = append(rtr.routes, e)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, e := range rtr.routes {
		match := e.Match(r)
		if !match {
			continue
		}

		// We have a match! Call the handler, and return
		e.HandlerFunc.ServeHTTP(w, r)
		return
	}

	// No matches, so it's a 404
	http.NotFound(w, r)
}

// ~~~~~ RouteEntry ~~~~~ //

type RouteEntry struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

func (ent *RouteEntry) Match(r *http.Request) bool {
	println(r.Method, r.URL.Path)
	println(ent.Method, ent.Path)
	if r.Method != ent.Method {
		return false // Method mismatch
	}

	if r.URL.Path != ent.Path {
		return false // Path mismatch
	}

	return true
}

func URLParam(r *http.Request, name string) string {
	ctx := r.Context()

	// ctx.Value returns an `interface{}` type, so we
	// also have to cast it to a map, which is the
	// type we'll be using to store our parameters.
	params := ctx.Value("params").(map[string]string)
	return params[name]
}
