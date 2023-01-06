package repositories

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/models/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) FindById(id int) (*domain.User, error) {
	var user *domain.User
	err := repository.DB.Model(&domain.User{}).Where("user_id = ?", id).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repository *UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user *domain.User
	err := repository.DB.Model(&domain.User{}).Where("username = ?", username).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repository *UserRepositoryImpl) Create(user *domain.User) (*domain.User, error) {
	err := repository.DB.Model(&domain.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repository *UserRepositoryImpl) Update(id int, user *domain.User) (*domain.User, error) {
	err := repository.DB.Model(&domain.User{}).Exec("UPDATE user SET first_name=?, last_name=?, email=?, username=?, password=? WHERE user_id=?",
		user.FirstName,
		user.LastName,
		user.Email,
		user.Username,
		user.Password,
		id,
	).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repository *UserRepositoryImpl) Delete(id int) error {
	code, err := repository.FindById(id)
	if err != nil {
		return err
	}
	response := repository.DB.Model(&domain.User{}).Delete(&domain.User{}, code)
	if response.Error != nil {
		return response.Error
	}
	return nil
}
