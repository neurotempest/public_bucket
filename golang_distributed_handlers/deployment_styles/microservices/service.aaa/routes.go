package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/neurotempest/public_bucket/golang_distributed_handlers/deployment_styles/lib"
	aaa_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/aaa/get"
	aaa_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/aaa/post"
	aaa_aaa_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/aaa/aaa/get"
	aaa_aaa_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/aaa/aaa/post"
	aaa_bbb_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/aaa/bbb/get"
	aaa_bbb_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/aaa/bbb/post"
	aaa_ccc_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/aaa/ccc/get"
	aaa_ccc_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/aaa/ccc/post"
)

func addRoutes(
	router *httprouter.Router,
) {
	router.GET("/readiness", handleGETReadiness())

	router.GET("/aaa", lib.Handle(aaa_get.GetHandler))
	router.POST("/aaa", lib.Handle(aaa_post.PostHandler))
	router.GET("/aaa/aaa", lib.Handle(aaa_aaa_get.GetHandler))
	router.POST("/aaa/aaa", lib.Handle(aaa_aaa_post.PostHandler))
	router.GET("/aaa/bbb", lib.Handle(aaa_bbb_get.GetHandler))
	router.POST("/aaa/bbb", lib.Handle(aaa_bbb_post.PostHandler))
	router.GET("/aaa/ccc", lib.Handle(aaa_ccc_get.GetHandler))
	router.POST("/aaa/ccc", lib.Handle(aaa_ccc_post.PostHandler))
}

func handleGETReadiness() httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	}
}

