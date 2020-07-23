package mysql

import (
	"github.com/aasumitro/gorest/logs"
	"github.com/aasumitro/gorest/src/domain"
	"github.com/jinzhu/gorm"
)

type mysqlExampleRepository struct {
	connection *gorm.DB
}

func NewMySQLExampleRepository(db *gorm.DB) domain.ExampleMySQlRepositoryContract {
	return &mysqlExampleRepository{connection: db}
}

func (mysql mysqlExampleRepository) Fetch() (data []domain.Example, error error) {
	var examples []domain.Example
	if err := mysql.connection.Find(&examples).Error; err != nil {
		logs.AppError.Println(err)
		return nil, err
	}
	return examples, nil
}

func (mysql mysqlExampleRepository) Find(id int) (data domain.Example, error error) {
	var example domain.Example
	if err := mysql.connection.First(&example, id).Error; err != nil {
		logs.AppError.Println(err)
		return example, err
	}
	return example, nil
}

func (mysql mysqlExampleRepository) Store(example *domain.Example) error {
	if err := mysql.connection.Create(&example).Error; err != nil {
		logs.AppError.Println(err)
		return err
	}
	return nil
}

func (mysql mysqlExampleRepository) Update(example *domain.Example) error {
	if err := mysql.connection.Model(&example).Updates(example).Error; err != nil {
		logs.AppError.Println(err)
		return err
	}
	return nil
}

func (mysql mysqlExampleRepository) Delete(example *domain.Example) error {
	if err := mysql.connection.Delete(&example).Error; err != nil {
		logs.AppError.Println(err)
		return err
	}
	return nil
}