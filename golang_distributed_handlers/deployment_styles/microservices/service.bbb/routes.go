package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/neurotempest/public_bucket/golang_distributed_handlers/deployment_styles/lib"
	bbb_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/get"
	bbb_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/post"
	bbb_aaa_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/aaa/get"
	bbb_aaa_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/aaa/post"
	bbb_bbb_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/bbb/get"
	bbb_bbb_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/bbb/post"
	bbb_ccc_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/ccc/get"
	bbb_ccc_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/ccc/post"
)

func addRoutes(
	router *httprouter.Router,
) {
	router.GET("/readiness", handleGETReadiness())

	router.GET("/bbb", lib.Handle(bbb_get.GetHandler))
	router.POST("/bbb", lib.Handle(bbb_post.PostHandler))
	router.GET("/bbb/aaa", lib.Handle(bbb_aaa_get.GetHandler))
	router.POST("/bbb/aaa", lib.Handle(bbb_aaa_post.PostHandler))
	router.GET("/bbb/bbb", lib.Handle(bbb_bbb_get.GetHandler))
	router.POST("/bbb/bbb", lib.Handle(bbb_bbb_post.PostHandler))
	router.GET("/bbb/ccc", lib.Handle(bbb_ccc_get.GetHandler))
	router.POST("/bbb/ccc", lib.Handle(bbb_ccc_post.PostHandler))
}

func handleGETReadiness() httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	}
}

