package services

import (

	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/models"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/repositories"
)

type BookService interface {
	CreateBook(book *models.Book) error
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{
		bookRepo: repo,
	}
}

func (b *bookService) CreateBook(book *models.Book) error {
	err := b.bookRepo.CreateBook(book)
	if err != nil {
		return err
	}

	return nil
}
