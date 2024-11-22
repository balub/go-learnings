package routes

import (
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/database"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/repositories"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/services"
)

func initalizeRoutes(DB *database.Database) {
	// initialize repositories
	bookRepository := repositories.NewBookRepository(DB)
	authorRepository := repositories.NewAuthorRepository(DB)

	// initialize services
	bookService := services.BookService(bookRepository)
	authorService := services.AuthorService(authorRepository)

	// initialize route handlers
	book
}
