package transaction

type Service interface {
	FindAll() ([]Transaction, error)
	FindByID(ID int) (Transaction, error)
	Create(TransactionRequest TransactionRequest) (Transaction, error)
	Update(ID int, TransactionRequest TransactionRequest) (Transaction, error)
	Delete(ID int) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Transaction, error) {
	Transactions, err := s.repository.FindAll()
	return Transactions, err

	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Transaction, error) {
	Transactions, err := s.repository.FindByID(ID)
	return Transactions, err
	// return s.repository.FindAll()
}

func (s *service) Create(TransactionRequest TransactionRequest) (Transaction, error) {
	price, _ := TransactionRequest.Price.Int64()

	Transaction := Transaction{
		Name_product: TransactionRequest.Name_product,
		Image_url:    TransactionRequest.Image_url,
		Description:  TransactionRequest.Description,
		Price:        int(price),
		Name_user:    TransactionRequest.Name_user,
		Email_user:   TransactionRequest.Email_user,
		Name_buyer:   TransactionRequest.Name_buyer,
		Email_buyer:  TransactionRequest.Email_buyer,
	}

	newTransaction, err := s.repository.Create(Transaction)

	return newTransaction, err
	// return s.repository.FindAll()
}

func (s *service) Update(ID int, TransactionRequest TransactionRequest) (Transaction, error) {
	Transaction, _ := s.repository.FindByID(ID)

	price, _ := TransactionRequest.Price.Int64()

	//
	Transaction.Name_product = TransactionRequest.Name_product
	Transaction.Image_url = TransactionRequest.Image_url
	Transaction.Description = TransactionRequest.Description
	Transaction.Price = int(price)
	Transaction.Name_user = TransactionRequest.Name_user
	Transaction.Email_user = TransactionRequest.Email_user
	Transaction.Name_buyer = TransactionRequest.Name_buyer
	Transaction.Email_buyer = TransactionRequest.Email_buyer

	newTransaction, err := s.repository.Update(Transaction)

	return newTransaction, err
	// return s.repository.FindAll()
}

func (s *service) Delete(ID int) (Transaction, error) {
	Transaction, _ := s.repository.FindByID(ID)

	newTransaction, err := s.repository.Delete(Transaction)

	return newTransaction, err
	// return s.repository.FindAll()
}
