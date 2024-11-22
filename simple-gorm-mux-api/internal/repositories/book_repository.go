package repositories

import (
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/database"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook() models.Book
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(DB *database.Database) BookRepository {
	return &bookRepository{
		db: DB.GetDB(),
	}
}

func (b *bookRepository) CreateBook() models.Book {
	return models.Book{Id: "12", Title: "Some random title", Genre: "sci-fi"}
}
