package service

import (
	"github.com/loctodale/go-ecommerce-backend-api/internal/repo"
	"github.com/loctodale/go-ecommerce-backend-api/pkg/reponse"
)

//type UserService struct {
//	userRepo *repo.UserRepo
//}
//
//func NewUserService() *UserService {
//	return &UserService{
//		userRepo: repo.NewUserRepo(),
//	}
//}
//
//func (us *UserService) GetInfoUserService() string {
//	return us.userRepo.GetInfoUser()
//}

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return reponse.ErrorCodeUserHasExited
	}
	return reponse.ErrInvalidToken
}
