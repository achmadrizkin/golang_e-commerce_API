package hoodie

type Service interface {
	FindAll() ([]Hoodie, error)
	FindByID(ID int) (Hoodie, error)
	Create(HoodieRequest HoodieRequest) (Hoodie, error)
	Update(ID int, HoodieRequest HoodieRequest) (Hoodie, error)
	Delete(ID int) (Hoodie, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Hoodie, error) {
	Hoodies, err := s.repository.FindAll()
	return Hoodies, err

	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Hoodie, error) {
	Hoodies, err := s.repository.FindByID(ID)
	return Hoodies, err
	// return s.repository.FindAll()
}

func (s *service) Create(HoodieRequest HoodieRequest) (Hoodie, error) {
	price, _ := HoodieRequest.Price.Int64()

	Hoodie := Hoodie{
		Name_product: HoodieRequest.Name_product,
		Image_url:    HoodieRequest.Image_url,
		Description:  HoodieRequest.Description,
		Price:        int(price),
		Name_user:    HoodieRequest.Name_user,
		Email_user:   HoodieRequest.Email_user,
	}

	newHoodie, err := s.repository.Create(Hoodie)

	return newHoodie, err
	// return s.repository.FindAll()
}

func (s *service) Update(ID int, HoodieRequest HoodieRequest) (Hoodie, error) {
	Hoodie, _ := s.repository.FindByID(ID)

	price, _ := HoodieRequest.Price.Int64()

	//
	Hoodie.Name_product = HoodieRequest.Name_product
	Hoodie.Image_url = HoodieRequest.Image_url
	Hoodie.Description = HoodieRequest.Description
	Hoodie.Price = int(price)
	Hoodie.Name_user = HoodieRequest.Name_user
	Hoodie.Email_user = HoodieRequest.Email_user

	newHoodie, err := s.repository.Update(Hoodie)

	return newHoodie, err
	// return s.repository.FindAll()
}

func (s *service) Delete(ID int) (Hoodie, error) {
	Hoodie, _ := s.repository.FindByID(ID)

	newHoodie, err := s.repository.Delete(Hoodie)

	return newHoodie, err
	// return s.repository.FindAll()
}
