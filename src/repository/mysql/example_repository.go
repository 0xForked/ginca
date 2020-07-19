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
	mysql.Connection.Find(&examples)
	return examples, nil
}

func (mysql mysqlExampleRepository) Find(id int) (data domain.Example, error error) {
	var example domain.Example
	mysql.Connection.First(&example, id)
	return example, nil
}