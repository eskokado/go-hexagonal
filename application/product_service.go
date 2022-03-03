package application

import "log"

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		log.Fatalln("***** error service get - " + err.Error())
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()
	if err != nil {
		log.Fatalln("***** error service product isvalid - " + err.Error())
		return &Product{}, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		log.Fatalln("********** error save persistence(Create) - " + err.Error())
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		log.Fatalln("******* error enable product service - " + err.Error())
		return &Product{}, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		log.Fatalln("********** error save persistence(Enable) - " + err.Error())
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		log.Fatalln("******* error disable product service - " + err.Error())
		return &Product{}, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		log.Fatalln("********** error save persistence(Disable) - " + err.Error())
		return &Product{}, err
	}
	return result, nil
}
