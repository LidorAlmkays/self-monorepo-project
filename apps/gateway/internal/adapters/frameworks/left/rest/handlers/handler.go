package handlers

import "libs/golang/logger"

type Handler struct {
	l logger.CustomLogger
}

func NewHandler(l logger.CustomLogger)*Handler{
	return &Handler{l:l}
}