package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	v1 "practice/api/routes/v1"
)

func Init(router *httprouter.Router) {
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, "{ \"message\":\"Hello world!. I am Employee Service.\",\"success\":true,\"version\": \"1.0.0\" }")
	})
	router.GET("/elb-check", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, "{ \"message\":\"Hello world AWS ALB!. I am Employee Service.\",\"success\":true,\"version\": \"1.0.0\" }")
	})
	router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(404)
		_, _ = fmt.Fprint(rw, "{ \"message\":\"Not Found.\",\"success\":false,\"version\": \"1.0.0\" }")
	})

	v1.Init(router)
}
