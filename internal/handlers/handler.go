package handlers

import "url-shortener/internal/services"

type handler struct {
	serv *services.Service
}