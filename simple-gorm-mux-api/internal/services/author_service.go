package services

import (
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/models"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/repositories"
)

type AuthorService interface {
	CreateAuthor() models.Author
}

type authorService struct {
	repo repositories.AuthorRepository
}

func NewAuthorService(repo repositories.AuthorRepository) AuthorService {
	return &authorService{
		repo: repo,
	}
}

func (a *authorService) CreateAuthor() models.Author {
	return models.Author{Id: "fvdf", Name: "Balu Babu", Bio: "I like anime"}
}
