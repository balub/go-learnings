package services

import (
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/models"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/repositories"
)

type BookService interface {
	CreateBook() models.Book
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{
		bookRepo: repo,
	}
}

func (b *bookService) CreateBook() models.Book {
	return models.Book{Id: "12", Title: "Some random title", Genre: "sci-fi"}
}
