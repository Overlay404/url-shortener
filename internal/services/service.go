package services

import (
	"fmt"
	"log/slog"
	"time"
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

func (ser *Service) SetLink(url string) (urlShort string, err error) {
	urlShortFull := "http://localhost:8000/v1/g/"
	link := models.Link{
		Url:     url,
		Created: time.Now().Format("02.01.2006"),
		Clicks:  0,
	}

	urlShort = pkg.GenShortUrl(6)
	urlShortFull += urlShort
	err = pkg.WriteInFile(urlShort, ser.log)
	if err != nil {
		return "", fmt.Errorf("cannot write short url in file %w", err)
	}
	ser.log.Debug("Service: generate short url", slog.String("urlShort", urlShort))

	err = ser.rep.Set(urlShort, &link)
	if err != nil {
		return "", fmt.Errorf("cannot write object in redis %w", err)
	}
	ser.log.Debug("Service: set link in redis", slog.String("url", url), slog.String("short-url", urlShort))
	return urlShortFull, nil
}

func (ser *Service) Get(urlShort string) (link *models.Link, err error) {
	link, err = ser.rep.Get(urlShort)
	if err != nil {
		return nil, fmt.Errorf("cannot get object from redis %w", err)
	}

	ser.log.Debug("Service: get value in redis", slog.String("urlShort", urlShort))
	return link, nil
}

func (ser *Service) LoadAllShortUrl() (arr []string, err error) {
	arr, err = pkg.ReadFromFile(ser.log)
	if err != nil {
		return nil, fmt.Errorf("load short url list %w", err)
	}
	return arr, nil
}

func (ser *Service) ClickLink(urlShort string) {
	err := ser.rep.Click(urlShort)
	if err != nil{
		ser.log.Error("click not work", slog.String("err", err.Error()))
		return
	}
}
