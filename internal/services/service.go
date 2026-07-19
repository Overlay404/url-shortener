package services

import (
	"fmt"
	"log/slog"
	"url-shortener/internal/models"
	"url-shortener/internal/repositories"
	"url-shortener/pkg"
)

type Service struct {
	rep *repositories.Repos
	log *slog.Logger
}

func New(repos *repositories.Repos, logger *slog.Logger) (ser *Service) {
	return &Service{
		rep: repos,
		log: logger,
	}
}

func (ser *Service) Set(link *models.Link) (urlShort string, err error) {
	urlShort = pkg.GenShortUrl(6)
	err = pkg.WriteInFile(urlShort, ser.log)
	if err != nil {
		return "", fmt.Errorf("cannot write short url in file %w", err)
	}

	err = ser.rep.Set(urlShort, link)
	if err != nil {
		return "", fmt.Errorf("cannot write object in redis %w", err)
	}

	ser.log.Debug("Service: generate short url", slog.String("urlShort", urlShort))
	return urlShort, nil
}

func (ser *Service) Get(urlShort string) (link *models.Link, err error) {
	link, err = ser.rep.Get(urlShort)
	if err != nil {
		return nil, fmt.Errorf("cannot get object from redis %w", err)
	}

	ser.log.Debug("Service: get value in redis", slog.String("urlShort", urlShort))
	return link, nil
}
