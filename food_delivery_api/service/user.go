package service

import (
	"food_delivery/repository"
	"food_delivery/server/request"
	"food_delivery/server/response"
)

type UserService struct {
	userRepositoryI repository.UserRepositoryI
}

type UserServiceI interface {
	GetAll() ([]*response.User, error)
	RegisterUser(u request.RegisterRequest) (*response.User, error)
	GetUserByEmail(email string) (*response.User, error)
	GetUserByID(id int) (*response.User, error)
	UpdateUserPasswordById(id int, password string) error
	UpdateUserProfile(ID int, req request.UpdateUserRequest) error
	//ChangePassword(req request.ChangePasswordRequest, cfg *config.Config) (*response.TokenResponse, error)
}

func NewUserService(userRepositoryI repository.UserRepositoryI) UserServiceI {
	return &UserService{
		userRepositoryI: userRepositoryI,
	}
}

func (r *UserService) GetAll() ([]*response.User, error) {
	users, err := r.userRepositoryI.GetAll()
	if err != nil {
		return nil, err
	}

	var userResponses []*response.User
	for _, user := range users {
		userResponses = append(userResponses, &response.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Password:  user.Password,
			Age:       user.Age,
			Email:     user.Email,
			Phone:     user.Phone,
			Address:   user.Address,
		})
	}

	return userResponses, nil
}

func (r *UserService) RegisterUser(u request.RegisterRequest) (*response.User, error) {
	user, err := r.userRepositoryI.RegisterUser(u)
	if err != nil {
		return nil, err
	}

	return &response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  user.Password,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
	}, nil
}

func (r *UserService) GetUserByEmail(email string) (*response.User, error) {
	user, err := r.userRepositoryI.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return &response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  user.Password,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
	}, nil
}

func (r *UserService) GetUserByID(id int) (*response.User, error) {

	user, err := r.userRepositoryI.GetUserByID(int64(id))
	if err != nil {
		return nil, err
	}

	return &response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  user.Password,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
	}, nil
}

func (r *UserService) UpdateUserPasswordById(id int, password string) error {

	err := r.userRepositoryI.UpdateUserPasswordByID(int64(id), password)
	if err != nil {
		return err
	}

	return nil

}

func (r *UserService) UpdateUserProfile(userID int, req request.UpdateUserRequest) error {

	user, err := r.userRepositoryI.GetUserByID(int64(userID))
	if err != nil {
		return err
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Age != 0 {
		user.Age = req.Age
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Address != "" {
		user.Address = req.Address
	}

	isUserNameChanged := false
	if req.Username != "" {
		isUserNameChanged = true
	}

	err = r.userRepositoryI.UpdateUserProfileByID(user, isUserNameChanged)
	if err != nil {
		return err
	}

	return nil
}
