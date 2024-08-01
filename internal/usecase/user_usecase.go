package usecase

import (
	"github.com/josevitorrodriguess/e-learning/internal/config/logger"
	"github.com/josevitorrodriguess/e-learning/internal/controllers/response"
	"github.com/josevitorrodriguess/e-learning/internal/models"
	"github.com/josevitorrodriguess/e-learning/internal/repository"
	"github.com/josevitorrodriguess/e-learning/internal/services"
)

type UserUseCase interface {
	GetAllUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	Create(user models.User) (response.User, error)
	Update(user models.User) error
	Delete(id string) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (uc *userUseCase) Create(user models.User) (response.User, error) {
	user.Password = services.Encrypt(user.Password)

	isValidEmail, err := services.ValidateEmail(user.Email)
	if err != nil {
		logger.Error("Error to validate email", err)
		return response.User{}, err
	}
	if !isValidEmail {
		logger.Error("Invalid email format", err)
		return response.User{}, err
	}

	isValidCPF, err := services.ValidateCPF(user.CPF)
	if err != nil {
		logger.Error("Error to validate CPF", err)
		return response.User{}, err
	}
	if !isValidCPF {
		logger.Error("Invalid CPF", err)
		return response.User{}, err
	}

	return uc.userRepo.Create(user)
}

func (uc *userUseCase) GetAllUsers() ([]models.User, error) {
	return uc.userRepo.GetAllUsers()
}

func (uc *userUseCase) GetUser(id string) (models.User, error) {
	return uc.userRepo.GetUser(id)
}

func (uc *userUseCase) Update(user models.User) error {
	return uc.userRepo.Update(user)
}

func (uc *userUseCase) Delete(id string) error {
	return uc.userRepo.Delete(id)
}
