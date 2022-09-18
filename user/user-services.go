package user

import "golang.org/x/crypto/bcrypt"

type UserService interface {
	FindAll() ([]User, error)
	CreateNew(user RegisterUserInput) (User, error)
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) FindAll() ([]User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) CreateNew(user RegisterUserInput) (User, error) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	newUser := User{
		Name:         user.Name,
		Occupation:   user.Occupation,
		Email:        user.Email,
		PasswordHash: string(hash),
	}

	return s.userRepository.CreateNewUser(newUser)
}
