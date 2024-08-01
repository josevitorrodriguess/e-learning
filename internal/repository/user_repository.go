package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/josevitorrodriguess/e-learning/internal/config/logger"
	"github.com/josevitorrodriguess/e-learning/internal/controllers/response"
	"github.com/josevitorrodriguess/e-learning/internal/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	Create(user models.User) (response.User, error)
	Update(user models.User) error
	Delete(id string) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	query := "SELECT * FROM users"
	err := ur.db.Select(&users, query)
	if err != nil {
		logger.Error("error selecting users from db", err)
		return nil, err
	}
	logger.Info("successfully fetched all users")
	return users, nil
}

func (ur *userRepository) GetUser(id string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE id = $1"
	err := ur.db.Get(&user, query, id)
	if err != nil {
		logger.Error("error fetching user from db", err)
		return models.User{}, err
	}
	logger.Info("successfully fetched user")
	return user, nil
}

func (ur *userRepository) Create(user models.User) (response.User, error) {
	query := "INSERT INTO users (name, email, password, login, cpf,role, wallet) VALUES (?, ?, ?, ?, ?, ?, ?);"
	result, err := ur.db.Exec(query, user.Name, user.Email, user.Password, user.Login, user.CPF, user.Role,user.Wallet)
	if err != nil {
		logger.Error("error inserting user into db", err)
		return response.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error fetching user ID after insert", err)
		return response.User{}, err
	}
	userResponse := response.User{
		ID:     id,
		Name:   user.Name,
		Email:  user.Email,
		Login:  user.Login,
		CPF:    user.CPF,
		Role:   user.Role,
		Wallet: user.Wallet,
	}

	logger.Info("successfully created user")
	return userResponse, nil
}

func (ur *userRepository) Update(user models.User) error {
	query := `UPDATE users SET name = :name, email = :email, password = :password WHERE id = :id`
	_, err := ur.db.NamedExec(query, user)
	if err != nil {
		logger.Error("error updating user in db", err)
		return err
	}
	logger.Info("successfully updated user")
	return nil
}

func (ur *userRepository) Delete(id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := ur.db.Exec(query, id)
	if err != nil {
		logger.Error("error deleting user from db", err)
		return err
	}
	logger.Info("successfully deleted user")
	return nil
}
