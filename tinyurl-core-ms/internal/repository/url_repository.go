package repository

import (
	"errors"
	"net/url"
	"time"

	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/entity"
	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/util"
)

type TinyURLRepository interface {
	GetByCode(code string) (*entity.TinyURL, error)
	Save(rawOriginalURL string) (*entity.TinyURL, error)
}

type TinyURLRepositoryImpl struct {
	urls      map[string]*entity.TinyURL
	currentId int64
}

func NewURLRepository() TinyURLRepository {
	return &TinyURLRepositoryImpl{
		urls:      make(map[string]*entity.TinyURL),
		currentId: 0,
	}
}

func (r *TinyURLRepositoryImpl) GetByCode(code string) (*entity.TinyURL, error) {
	tinyURL, ok := r.urls[code]
	if !ok || tinyURL == nil {
		return nil, errors.New("tiny url not found or null in memory repository")
	}
	return tinyURL, nil
}

func (r *TinyURLRepositoryImpl) Save(rawOriginalURL string) (*entity.TinyURL, error) {
	originalURL, err := url.ParseRequestURI(rawOriginalURL)
	if err != nil {
		return nil, errors.New("invalid URL format")
	}

	for _, url := range r.urls {
		if url.RawURL == originalURL.String() {
			return url, nil
		}
	}

	code := util.GenerateCode()
	if _, exists := r.urls[code]; exists {
		return nil, errors.New("generated code already exists, try again")
	}

	r.currentId++
	tinyURL := &entity.TinyURL{
		ID:        r.currentId,
		Code:      code,
		URL:       originalURL,
		RawURL:    rawOriginalURL,
		CreatedAt: time.Now(),
	}

	r.urls[code] = tinyURL

	return tinyURL, nil
}
