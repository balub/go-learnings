package handlers

import (
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/services"
)

type AuthorHandler struct {
	authorService services.AuthorService
}

func NewAuthorHandler(authorService services.AuthorService) *AuthorHandler {
	return &AuthorHandler{
		authorService: authorService,
	}
}
