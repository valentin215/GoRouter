package main

import "net/http"

func main() {
	r := &Router{}
	http.ListenAndServe(":8000", r)
}

// ------- Router -------

type RouteEntry struct {
	Path		string
	Method  string
	Handler http.HandlerFunc
}

type Router struct{
	routes []RouteEntry
}

func (sr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
	http.NotFound(w, r)
}

