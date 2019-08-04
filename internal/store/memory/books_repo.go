package memory

import (
	"github.com/delgus/swagger-rest/internal/models"
)

type BooksRepo struct {
	Books models.ArrayOfBooks
}

func (r *BooksRepo) Fill() {
	var id1 int64 = 1
	var id2 int64 = 2
	author := "Dolgov"
	title1 := "First Book"
	title2 := "Second Book"
	r.Books = append(r.Books, &models.Book{
		ID:     &id1,
		Author: &author,
		Title:  &title1,
	})
	r.Books = append(r.Books, &models.Book{
		ID:     &id2,
		Author: &author,
		Title:  &title2,
	})
}

func (r *BooksRepo) GetAll() models.ArrayOfBooks {
	return r.Books
}

func (r *BooksRepo) GetAllByAuthor(author *string) models.ArrayOfBooks {
	var result models.ArrayOfBooks
	for _, b := range r.Books {
		if *b.Author == *author {
			result = append(result, b)
		}
	}
	return result
}

func (r *BooksRepo) GetAllByTitle(title *string) models.ArrayOfBooks {
	var result models.ArrayOfBooks
	for _, b := range r.Books {
		if *b.Title == *title {
			result = append(result, b)
		}
	}
	return result
}

func (r *BooksRepo) GetAllByAuthorAndTitle(author *string, title *string) models.ArrayOfBooks {
	var result models.ArrayOfBooks
	for _, b := range r.Books {
		if *b.Author == *author && *b.Title == *title {
			result = append(result, b)
		}
	}
	return result
}

func (r *BooksRepo) GetByID(id int64) *models.Book {
	for _, b := range r.Books {
		if *b.ID == id {
			return b
		}
	}
	return nil
}

func (r *BooksRepo) Save(book *models.Book) bool {
	r.Books = append(r.Books, book)
	return true
}

func (r *BooksRepo) Delete(id int64) bool {
	for _, b := range r.Books {
		if *b.ID == id {
			return true
		}
	}
	return false
}

func (r *BooksRepo) Update(book *models.Book) bool {
	for _, b := range r.Books {
		if *b.ID == *book.ID {
			b = book
			return true
		}
	}
	return false
}
