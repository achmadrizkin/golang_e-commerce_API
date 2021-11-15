package laptop

type Service interface {
	FindAll() ([]Laptop, error)
	FindByID(ID int) (Laptop, error)
	Create(LaptopRequest LaptopRequest) (Laptop, error)
	Update(ID int, LaptopRequest LaptopRequest) (Laptop, error)
	Delete(ID int) (Laptop, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Laptop, error) {
	Laptops, err := s.repository.FindAll()
	return Laptops, err

	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Laptop, error) {
	Laptops, err := s.repository.FindByID(ID)
	return Laptops, err
	// return s.repository.FindAll()
}

func (s *service) Create(LaptopRequest LaptopRequest) (Laptop, error) {
	price, _ := LaptopRequest.Price.Int64()

	Laptop := Laptop{
		Name_product: LaptopRequest.Name_product,
		Image_url:    LaptopRequest.Image_url,
		Description:  LaptopRequest.Description,
		Price:        int(price),
		Name_user:    LaptopRequest.Name_user,
		Email_user:   LaptopRequest.Email_user,
	}

	newLaptop, err := s.repository.Create(Laptop)

	return newLaptop, err
	// return s.repository.FindAll()
}

func (s *service) Update(ID int, LaptopRequest LaptopRequest) (Laptop, error) {
	Laptop, _ := s.repository.FindByID(ID)

	price, _ := LaptopRequest.Price.Int64()

	//
	Laptop.Name_product = LaptopRequest.Name_product
	Laptop.Image_url = LaptopRequest.Image_url
	Laptop.Description = LaptopRequest.Description
	Laptop.Price = int(price)
	Laptop.Name_user = LaptopRequest.Name_user
	Laptop.Email_user = LaptopRequest.Email_user

	newLaptop, err := s.repository.Update(Laptop)

	return newLaptop, err
	// return s.repository.FindAll()
}

func (s *service) Delete(ID int) (Laptop, error) {
	Laptop, _ := s.repository.FindByID(ID)

	newLaptop, err := s.repository.Delete(Laptop)

	return newLaptop, err
	// return s.repository.FindAll()
}
