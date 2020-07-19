package service

import (
	"github.com/aasumitro/gorest/domain"
)

type exampleService struct {
	exampleRepository	domain.ExampleRepository
}

func NewExampleService(repository domain.ExampleRepository) domain.ExampleService {
	return &exampleService{exampleRepository: repository}
}

func (service exampleService) Fetch() (data []domain.Example, error error) {
	data, error = service.exampleRepository.Fetch()
	if error != nil {
		return nil, error
	}
	return
}

func (service exampleService) Find(id int) (data domain.Example, error error) {
	data, error = service.exampleRepository.Find(id)
	if error != nil {
		return domain.Example{}, error
	}
	return
}
