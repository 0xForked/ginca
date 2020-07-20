package service

import (
	"github.com/aasumitro/gorest/domain"
)

type exampleService struct {
	exampleRepository	domain.MySQlRepository
}

func NewExampleService(repository domain.MySQlRepository) domain.ExampleService {
	return &exampleService{exampleRepository: repository}
}

func (service exampleService) Fetch() (data []domain.Example, err error) {
	data, err = service.exampleRepository.Fetch()
	return data, err
}

func (service exampleService) Find(id int) (data domain.Example, err error) {
	data, err = service.exampleRepository.Find(id)
	return data, err
}

func (service exampleService) Store(example *domain.Example) (err error) {
	if err := service.exampleRepository.Store(example); err != nil {
		return err
	}
	return
}

func (service exampleService) Update(example *domain.Example) (err error) {
	_, err = service.exampleRepository.Find(int(example.ID))
	if err != nil {
		return err
	}
	if err := service.exampleRepository.Update(example); err != nil {
		return err
	}
	return nil
}

func (service exampleService) Delete(id int) (err error) {
	example, err := service.exampleRepository.Find(id)
	if err != nil {
		return err
	}
	if err := service.exampleRepository.Delete(&example); err != nil {
		return err
	}
	return
}