package user

import (
	"errors"
	"time"
)

type Service interface {
	GetAllUsers() ([]User, error)
	GetUserByID(ID int) (User, error)
	CreateUser(input CreateUserInput) (User, error)
	UpdateUser(inputID GetUserDetailInput, inputData UpdateUserInput) (User, error)
	DeleteUser(ID int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()

	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) CreateUser(input CreateUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Role = input.Role

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) UpdateUser(inputID GetUserDetailInput, inputData UpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(inputID.ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No User found on with that ID")
	}

	user.Name = inputData.Name
	user.Role = inputData.Role
	user.UpdatedAt = time.Now()

	user, err = s.repository.Update(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) DeleteUser(ID int) (bool, error) {
	user, err := s.GetUserByID(ID)

	if err != nil {
		return false, err
	}

	deleted, err := s.repository.Delete(user)

	if err != nil {
		return false, err
	}
	return deleted, nil
}
