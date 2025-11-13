package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/handler"
	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/repository"
	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/service"
)

const (
	port = 8080
)

func main() {
	mux := http.NewServeMux()

	wire(mux)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Stating server at \"%s\"", addr)
	log.Fatalln(http.ListenAndServe(addr, mux))
}

func wire(mux *http.ServeMux) {
	urlRepository := repository.NewURLRepository()

	urlService := service.NewTinyURLService(urlRepository)

	handler.NewShortenURLHandler(urlService)
	mux.Handle("POST /", handler.NewShortenURLHandler(urlService))

	handler.NewRedirectURLHandler(urlService)
	mux.Handle("GET /{code}", handler.NewRedirectURLHandler(urlService))
}
