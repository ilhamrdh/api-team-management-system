package repositories

import "github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"

type UserRepository interface {
	FindById(id int) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Update(id int, user *domain.User) (*domain.User, error)
	Delete(id int) error
}
