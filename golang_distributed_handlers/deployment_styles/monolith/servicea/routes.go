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
	bbb_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/get"
	bbb_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/post"
	bbb_aaa_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/aaa/get"
	bbb_aaa_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/aaa/post"
	bbb_bbb_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/bbb/get"
	bbb_bbb_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/bbb/post"
	bbb_ccc_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/ccc/get"
	bbb_ccc_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/bbb/ccc/post"
	ccc_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/get"
	ccc_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/post"
	ccc_aaa_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/aaa/get"
	ccc_aaa_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/aaa/post"
	ccc_bbb_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/bbb/get"
	ccc_bbb_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/bbb/post"
	ccc_ccc_get "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/ccc/get"
	ccc_ccc_post "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/ccc/ccc/post"
	users_list "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/users/list"
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
	router.GET("/bbb", lib.Handle(bbb_get.GetHandler))
	router.POST("/bbb", lib.Handle(bbb_post.PostHandler))
	router.GET("/bbb/aaa", lib.Handle(bbb_aaa_get.GetHandler))
	router.POST("/bbb/aaa", lib.Handle(bbb_aaa_post.PostHandler))
	router.GET("/bbb/bbb", lib.Handle(bbb_bbb_get.GetHandler))
	router.POST("/bbb/bbb", lib.Handle(bbb_bbb_post.PostHandler))
	router.GET("/bbb/ccc", lib.Handle(bbb_ccc_get.GetHandler))
	router.POST("/bbb/ccc", lib.Handle(bbb_ccc_post.PostHandler))
	router.GET("/ccc", lib.Handle(ccc_get.GetHandler))
	router.POST("/ccc", lib.Handle(ccc_post.PostHandler))
	router.GET("/ccc/aaa", lib.Handle(ccc_aaa_get.GetHandler))
	router.POST("/ccc/aaa", lib.Handle(ccc_aaa_post.PostHandler))
	router.GET("/ccc/bbb", lib.Handle(ccc_bbb_get.GetHandler))
	router.POST("/ccc/bbb", lib.Handle(ccc_bbb_post.PostHandler))
	router.GET("/ccc/ccc", lib.Handle(ccc_ccc_get.GetHandler))
	router.POST("/ccc/ccc", lib.Handle(ccc_ccc_post.PostHandler))

	router.GET("/users/list", lib.Handle(users_list.GetHandler))
}

func handleGETReadiness() httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	}
}

