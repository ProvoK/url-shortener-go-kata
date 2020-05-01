package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gopinionated/internal/shortener"
	"gopinionated/internal/shortener/backend"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func putURL(controller shortener.Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		req := struct {
			URL string `json:"url"`
		}{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		id, err := controller.Shorten(context.TODO(), req.URL)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"id": id,
		})

	}
}

func redirect(controller shortener.Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ID := chi.URLParam(r, "ID")
		url, err := controller.ResolveID(context.TODO(), ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("Can't find given ID: %s", ID),
			})
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusMovedPermanently)
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	conf := backend.BoltConfig{Path: "bolt.db"}
	bk, err := backend.NewBoltBackend(conf)
	if err != nil {
		log.Fatalf("Fatal: %s", err)
	}
	controller := shortener.NewController(bk)

	r.Put("/urls", putURL(controller))
	r.Get("/r/{ID}", redirect(controller))
	http.ListenAndServe(":8080", r)
}
