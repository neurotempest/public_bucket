package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewServer(
) http.Handler {

	router := httprouter.New()
	addRoutes(
		router,
	)
	var handler http.Handler = router
	//handler = someMiddleware(handler)
	//handler = someMiddleware2(handler)
	//handler = someMiddleware3(handler)
	return handler
}
