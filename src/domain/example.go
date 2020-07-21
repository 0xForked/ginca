package domain

import "time"

type Example struct {
	ID        	uint 		`json:"id" gorm:"primary_key" sql:"index"`
	Name		string		`json:"name" binding:"required"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	DeletedAt 	*time.Time 	`json:"deleted_at" sql:"index"`
}

type ExampleService interface {
	Fetch() ([]Example, error)
	Find(id int) (Example, error)
	Store(example *Example) error
	Update(example *Example) error
	Delete(id int) error
}

type ExampleMySQlRepository interface {
	Fetch() ([]Example, error)
	Find(id int) (Example, error)
	Store(example *Example) error
	Update(example *Example) error
	Delete(example *Example) error
}

type RedisRepository interface {
	Set(key string, value interface{})
	GetObject(key string) *map[string]interface{}
	GetArray(key string) *[]map[string]interface{}
	Ping() string
}