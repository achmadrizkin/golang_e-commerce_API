package allproducts

type Service interface {
	FindAll() ([]AllProduct, error)
	FindByID(ID int) (AllProduct, error)
	FindByCategory(category string) ([]AllProduct, error)
	FindByUser(email_user string) ([]AllProduct, error)
	FindByNameProduct(name_product string, price string, email_user string) ([]AllProduct, error)
	Create(allProductRequest AllProductRequest) (AllProduct, error)
	Update(ID int, allProductRequest AllProductRequest) (AllProduct, error)
	Delete(ID int) (AllProduct, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]AllProduct, error) {
	books, err := s.repository.FindAll()
	return books, err

	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (AllProduct, error) {
	books, err := s.repository.FindByID(ID)
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByCategory(category string) ([]AllProduct, error) {
	books, err := s.repository.FindByCategory(category)
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByUser(email_user string) ([]AllProduct, error) {
	books, err := s.repository.FindByUser(email_user)
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByNameProduct(name_product string, price string, email_user string) ([]AllProduct, error) {
	books, err := s.repository.FindByNameProduct(name_product, price, email_user)
	return books, err
	// return s.repository.FindAll()
}

func (s *service) Create(allProductRequest AllProductRequest) (AllProduct, error) {
	price, _ := allProductRequest.Price.Int64()

	book := AllProduct{
		Name_product: allProductRequest.Name_product,
		Image_url:    allProductRequest.Image_url,
		Description:  allProductRequest.Description,
		Price:        int(price),
		Name_user:    allProductRequest.Name_user,
		Email_user:   allProductRequest.Email_user,
		Category:     allProductRequest.Category,
	}

	newAllProduct, err := s.repository.Create(book)

	return newAllProduct, err
	// return s.repository.FindAll()
}

func (s *service) Update(ID int, allProductRequest AllProductRequest) (AllProduct, error) {
	book, _ := s.repository.FindByID(ID)

	price, _ := allProductRequest.Price.Int64()

	//
	book.Name_product = allProductRequest.Name_product
	book.Image_url = allProductRequest.Image_url
	book.Description = allProductRequest.Description
	book.Price = int(price)
	book.Name_user = allProductRequest.Name_user
	book.Email_user = allProductRequest.Email_user
	book.Category = allProductRequest.Category

	newAllProduct, err := s.repository.Update(book)

	return newAllProduct, err
	// return s.repository.FindAll()
}

func (s *service) Delete(ID int) (AllProduct, error) {
	book, _ := s.repository.FindByID(ID)

	newAllProduct, err := s.repository.Delete(book)

	return newAllProduct, err
	// return s.repository.FindAll()
}
