package books

import (
	"github.com/delgus/swagger-rest/internal/models"
)

type Service struct {
	r Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		r: repo,
	}
}

func (s *Service) GetAll(author *string, title *string) models.ArrayOfBooks {
	switch {
	case author != nil && title != nil:
		return s.r.GetAllByAuthorAndTitle(author, title)
	case author != nil:
		return s.r.GetAllByAuthor(author)
	case title != nil:
		return s.r.GetAllByTitle(title)
	default:
		return s.r.GetAll()
	}
}

func (s *Service) GetByID(id int64) *models.Book {
	return s.r.GetByID(id)
}

func (s *Service) Save(book *models.Book) bool {
	return s.r.Save(book)
}

func (s *Service) Delete(id int64) bool {
	return s.r.Delete(id)
}

func (s *Service) Update(book *models.Book) bool {
	return s.r.Update(book)
}
