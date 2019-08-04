package books

import "github.com/delgus/swagger-rest/internal/models"

type Repo interface {
	GetAll() models.ArrayOfBooks
	GetAllByAuthor(author *string) models.ArrayOfBooks
	GetAllByTitle(title *string) models.ArrayOfBooks
	GetAllByAuthorAndTitle(author *string, title *string) models.ArrayOfBooks
	GetByID(id int64) *models.Book
	Save(book *models.Book) bool
	Delete(id int64) bool
	Update(book *models.Book) bool
}
