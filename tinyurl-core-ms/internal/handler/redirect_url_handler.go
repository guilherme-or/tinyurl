package handler

import (
	"log"
	"net/http"

	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/service"
	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/util"
)

type RedirectURLHandler struct {
	urlService *service.TinyURLService
}

func NewRedirectURLHandler(urlService *service.TinyURLService) http.Handler {
	return &RedirectURLHandler{
		urlService: urlService,
	}
}

func (h *RedirectURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if err := util.ValidateCode(code); err != nil {
		log.Printf("RedirectURLHandler - Error validating code: %v", err)
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}

	tinyURL, err := h.urlService.GetTinyURL(code)
	if err != nil {
		log.Printf("RedirectURLHandler - Error retrieving tiny URL: %v", err)
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, tinyURL.RawURL, http.StatusFound)
}
