package routes

import (
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/database"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/handlers"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/repositories"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/services"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/utils"
	"github.com/gorilla/mux"
)

func InitializeRoutes(DB *database.Database, s *mux.Router) {
	// initialize repositories
	bookRepository := repositories.NewBookRepository(DB)
	// authorRepository := repositories.NewAuthorRepository(DB)

	// initialize services
	bookService := services.BookService(bookRepository)
	// authorService := services.AuthorService(authorRepository)

	// initialize route handlers
	bookHandler := handlers.NewBookHandler(bookService)
	// authorHandler := handlers.NewAuthorHandler(authorService)

	// initialize route handlers

	s.HandleFunc("/ping", utils.Pong)

	//TODO: route to create a book
	s.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")

	//TODO: route to retrieve all books and filter by author and get book by ID
	// s.HandleFunc("/books/{bookID}", bookHandler.GetBooks).Methods("GET")

	//TODO: route to update a specific books details
	// /	s.HandleFunc("/books/{bookID}", bookHandler.UpdateBookDetails).Methods("PUT")

	//TODO: route to delete a book
	// s.HandleFunc("/books/{bookID}", bookHandler.DeleteBook).Methods("DELETE")

	//TODO: route to create a new author
	// s.HandleFunc("/authors", authorHandler.CreateAuthor).Methods("POST")

	//TODO: route to retrieve all authors
	// s.HandleFunc("/authors/{authorID}", authorHandler.GetAuthors).Methods("GET")

	//TODO: route to update author details
	// s.HandleFunc("/authors/{authorID}", authorHandler.UpdateAuthorDetails).Methods("PUT")

	//TODO: route to delete a author
	// s.HandleFunc("/authors/{authorID}", authorHandler.DeleteAuthor).Methods("DELETE")
}
