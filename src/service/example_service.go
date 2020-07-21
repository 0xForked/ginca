package service

import (
	"github.com/aasumitro/gorest/src/domain"
)

type exampleService struct {
	exampleMySqlRepository domain.ExampleMySQlRepository
}

func NewExampleService(mysql domain.ExampleMySQlRepository) domain.ExampleService {
	return &exampleService{exampleMySqlRepository: mysql}
}

func (service exampleService) Fetch() (data []domain.Example, err error) {
	data, err = service.exampleMySqlRepository.Fetch()
	return data, err
}

func (service exampleService) Find(id int) (data domain.Example, err error) {

	data, err = service.exampleMySqlRepository.Find(id)
	return data, err
}

func (service exampleService) Store(example *domain.Example) (err error) {
	if err := service.exampleMySqlRepository.Store(example); err != nil {
		return err
	}
	return
}

func (service exampleService) Update(example *domain.Example) (err error) {
	_, err = service.exampleMySqlRepository.Find(int(example.ID))
	if err != nil {
		return err
	}
	if err := service.exampleMySqlRepository.Update(example); err != nil {
		return err
	}
	return nil
}

func (service exampleService) Delete(id int) (err error) {
	example, err := service.exampleMySqlRepository.Find(id)
	if err != nil {
		return err
	}
	if err := service.exampleMySqlRepository.Delete(&example); err != nil {
		return err
	}
	return
}