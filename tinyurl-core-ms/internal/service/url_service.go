package service

import (
	"errors"
	"log"

	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/dto"
	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/entity"
	"github.com/guilherme-or/tinyurl/tinyurl-core-ms/internal/repository"
)

type TinyURLService struct {
	tinyURLRepo repository.TinyURLRepository
}

func NewTinyURLService(tinyURLRepo repository.TinyURLRepository) *TinyURLService {
	return &TinyURLService{
		tinyURLRepo: tinyURLRepo,
	}
}

func (s *TinyURLService) GetTinyURL(code string) (*entity.TinyURL, error) {
	url, err := s.tinyURLRepo.GetByCode(code)

	if err != nil {
		log.Printf("TinyURLService - Error getting tiny URL by code %s: %v", code, err)
		return nil, errors.New("the requested url could not be found for the provided code")
	}

	log.Printf("TinyURLService - Retrieved tinyURL: code=%s url=%s", url.Code, url.RawURL)
	return url, nil
}

func (s *TinyURLService) CreateTinyURL(dto *dto.TinyURLCreateDTO) (*entity.TinyURL, error) {
	url, err := s.tinyURLRepo.Save(dto.RawURL)

	if err != nil {
		log.Printf("TinyURLService - Error saving tiny URL: %v", err)
		return nil, errors.New("there was an error while creating the tiny URL")
	}

	log.Printf("TinyURLService - Created new tinyURL: code=%s url=%s", url.Code, url.RawURL)
	return url, nil
}
