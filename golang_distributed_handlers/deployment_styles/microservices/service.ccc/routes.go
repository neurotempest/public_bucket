package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/neurotempest/public_bucket/golang_distributed_handlers/deployment_styles/lib"
	ccc_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/get"
	ccc_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/post"
	ccc_aaa_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/aaa/get"
	ccc_aaa_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/aaa/post"
	ccc_bbb_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/bbb/get"
	ccc_bbb_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/bbb/post"
	ccc_ccc_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/ccc/get"
	ccc_ccc_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/ccc/post"
)

func addRoutes(
	router *httprouter.Router,
) {
	router.GET("/readiness", handleGETReadiness())

	router.GET("/ccc", lib.Handle(ccc_get.GetHandler))
	router.POST("/ccc", lib.Handle(ccc_post.PostHandler))
	router.GET("/ccc/aaa", lib.Handle(ccc_aaa_get.GetHandler))
	router.POST("/ccc/aaa", lib.Handle(ccc_aaa_post.PostHandler))
	router.GET("/ccc/bbb", lib.Handle(ccc_bbb_get.GetHandler))
	router.POST("/ccc/bbb", lib.Handle(ccc_bbb_post.PostHandler))
	router.GET("/ccc/ccc", lib.Handle(ccc_ccc_get.GetHandler))
	router.POST("/ccc/ccc", lib.Handle(ccc_ccc_post.PostHandler))
}

func handleGETReadiness() httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	}
}

