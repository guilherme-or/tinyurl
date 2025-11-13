package handler

import (
	"log"
	"net/http"

	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/dto"
	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/service"
	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/util"
)

type ShortenURLHandler struct {
	urlService *service.TinyURLService
}

func NewShortenURLHandler(urlService *service.TinyURLService) http.Handler {
	return &ShortenURLHandler{
		urlService: urlService,
	}
}

func (h *ShortenURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var dto dto.TinyURLCreateDTO

	err := util.FromJSON(r, &dto)
	if err != nil {
		log.Printf("ShortenURLHandler - Error decoding request body: %v", err)
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}

	tinyURL, err := h.urlService.CreateTinyURL(&dto)
	if err != nil {
		log.Printf("ShortenURLHandler - Error creating tiny URL: %v", err)
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}

	if err := util.ToJSON(w, http.StatusCreated, tinyURL); err != nil {
		log.Printf("ShortenURLHandler - Error encoding response: %v", err)
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
}
