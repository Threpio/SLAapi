package main

import "github.com/go-chi/chi"

func main() {
	
	r := chi.NewRouter()
	
	r.Route("/search", func(r chi.Router) {
		r.Get("/flexible", )
	}

}