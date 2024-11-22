package handlers

import (

	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/services"
)

type BookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}


