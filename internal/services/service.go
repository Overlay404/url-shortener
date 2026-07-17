package services

import "url-shortener/internal/repositories"

type Service struct {
	rep *repositories.Repos
}