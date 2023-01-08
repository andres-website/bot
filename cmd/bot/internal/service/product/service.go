package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {

	// TODO: проверять входные данные idx - обрабатывать ошибку, говорить пользователю про ошибку
	// TODO: сериви Get должен проверить границы, в случае если границы не совпадают - должен вернуть
	// 		ошибку
	return &allProducts[idx], nil
}
