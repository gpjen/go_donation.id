package user

import "gorm.io/gorm"

type UserRepository interface {
	FindAll() ([]User, error)
	CreateNewUser(user User) (User, error)
	// FindById(ID uint64) (User, error)
	// FindByEmail(email string) (User, error)
	// Delete(ID uint64) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateNewUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindAll() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error
	return user, err
}
