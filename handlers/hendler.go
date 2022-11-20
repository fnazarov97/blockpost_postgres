package handlers

import (
	"article/config"
	"article/storage"
)

type Handler struct {
	IM   storage.StorageI
	Conf config.Config
}
