package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err

	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	books, err := s.repository.FindByID(ID)
	return books, err
	// return s.repository.FindAll()
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()

	book := Book{
		Name_product: bookRequest.Name_product,
		Image_url:    bookRequest.Image_url,
		Description:  bookRequest.Description,
		Price:        int(price),
		Name_user:    bookRequest.Name_user,
		Email_user:   bookRequest.Email_user,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
	// return s.repository.FindAll()
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	price, _ := bookRequest.Price.Int64()

	//
	book.Name_product = bookRequest.Name_product
	book.Image_url = bookRequest.Image_url
	book.Description = bookRequest.Description
	book.Price = int(price)
	book.Name_user = bookRequest.Name_user
	book.Email_user = bookRequest.Email_user

	newBook, err := s.repository.Update(book)

	return newBook, err
	// return s.repository.FindAll()
}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	newBook, err := s.repository.Delete(book)

	return newBook, err
	// return s.repository.FindAll()
}
