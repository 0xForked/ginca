package domain

import "time"

type Example struct {
	ID        	uint 		`json:"id" gorm:"primary_key" sql:"index"`
	Name		string		`json:"name" binding:"required"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	DeletedAt 	*time.Time 	`json:"deleted_at" sql:"index"`
}

type ExampleServiceContract interface {
	Fetch() ([]Example, error)
	Find(id int) (Example, error)
	Store(example *Example) error
	Update(example *Example) error
	Delete(id int) error
}

type ExampleMySQlRepositoryContract interface {
	Fetch() ([]Example, error)
	Find(id int) (Example, error)
	Store(example *Example) error
	Update(example *Example) error
	Delete(example *Example) error
}
