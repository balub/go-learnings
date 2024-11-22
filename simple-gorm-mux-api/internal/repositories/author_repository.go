package repositories

import (
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/database"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	CreateAuthor() models.Author
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(DB *database.Database) AuthorRepository {
	return &authorRepository{
		db: DB.GetDB(),
	}
}

func (a *authorRepository) CreateAuthor() models.Author {
	return models.Author{Id: "fvdf", Name: "Balu Babu", Bio: "I like anime"}
}
