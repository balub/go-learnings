package repositories

import (
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/database"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(DB *database.Database) BookRepository {
	return &bookRepository{
		db: DB.GetDB(),
	}
}

func (b *bookRepository) CreateBook(book *models.Book) error {
	newBook := models.Book{Id: book.Id, Title: book.Title, Genre: book.Genre}
	return b.db.Create(&newBook).Error
}
