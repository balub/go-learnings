package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/models"
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

func (b *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	// Decode json post body into valid book object
	var bookData models.Book
	err := json.NewDecoder(r.Body).Decode(&bookData)
	if err != nil {
		log.Fatal("Error: coudld not decode newBook JSON")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: coudld not decode newBook JSON"))
	}

	// Do input validation
	if len(bookData.Title) <= 3 {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Title too short"))
	}

	// Create new book
	err = b.bookService.CreateBook(&bookData)

	// Marshal JSON
	res, _ := json.Marshal(bookData)

	// Return newly created book
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
