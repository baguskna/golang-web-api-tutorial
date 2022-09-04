package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(bookReq BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	book, err := s.repository.FindAll()
	return book, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(id int) (Book, error) {
	book, err := s.repository.FindByID(id)
	return book, err
}

func (s *service) Create(bookReq BookRequest) (Book, error) {
	price, _ := bookReq.Price.Int64()
	rating, _ := bookReq.Rating.Int64()

	book := Book{
		Title:       bookReq.Title,
		Price:       int(price),
		Rating:      int(rating),
		Description: bookReq.Description,
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}
