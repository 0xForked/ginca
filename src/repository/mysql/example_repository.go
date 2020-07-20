package mysql

import (
	"github.com/aasumitro/gorest/domain"
	"github.com/jinzhu/gorm"
)

type mysqlExampleRepository struct {
	Connection *gorm.DB
}

func NewMySQLExampleRepository(db *gorm.DB) domain.ExampleRepository {
	return &mysqlExampleRepository{Connection: db}
}

func (mysql mysqlExampleRepository) Fetch() (data []domain.Example, error error) {
	var examples []domain.Example
	if err := mysql.Connection.Find(&examples).Error; err != nil {
		return nil, err
	}
	return examples, nil
}

func (mysql mysqlExampleRepository) Find(id int) (data domain.Example, error error) {
	var example domain.Example
	if err := mysql.Connection.First(&example, id).Error; err != nil {
		return example, err
	}
	return example, nil
}

func (mysql mysqlExampleRepository) Store(example *domain.Example) error {
	if err := mysql.Connection.Create(&example).Error; err != nil {
		return err
	}
	return nil
}

func (mysql mysqlExampleRepository) Update(example *domain.Example) error {
	if err := mysql.Connection.Model(&example).Updates(example).Error; err != nil {
		return err
	}
	return nil
}

func (mysql mysqlExampleRepository) Delete(example *domain.Example) error {
	if err := mysql.Connection.Delete(&example).Error; err != nil {
		return err
	}
	return nil
}