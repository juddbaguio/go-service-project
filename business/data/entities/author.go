package entities

import "go-service-project/business/domain"

type Author struct {
	ID        int    `gorm:"primaryKey;column:author_id"`
	FirstName string `gorm:"column:first_name;type:nvarchar(50)"`
	LastName  string `gorm:"column:last_name;type:nvarchar(50)"`
}

func (a *Author) ToDomain() domain.Author {
	return domain.Author{
		ID:        a.ID,
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
}
