package services

import (
	"example.com/sarang-apis/dto"
	"example.com/sarang-apis/models"
	"example.com/sarang-apis/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=../mocks/service/mockUserService.go -package=services example.com/sarang-apis/services UserService
type DefaultUserService struct {
	Repo repository.UserRepository
}

type UserService interface {
	UserInsert(user models.User) (*dto.UserDTO, error)
	UserGetAll() ([]models.User, error)
	UserDelete(id primitive.ObjectID) (bool, error)
	UserUpdate(id primitive.ObjectID, update bson.M) (bool, error)
}

func (t DefaultUserService) UserInsert(user models.User) (*dto.UserDTO, error) {

	var res dto.UserDTO

	if len(user.Name) <= 2 {

		res.Status = false
		return &res, nil
	}
	if len(user.LastName) <= 2 {

		res.Status = false
		return &res, nil
	}

	result, err := t.Repo.Insert(user) //send user

	if err != nil || result == false {
		res.Status = false

		return &res, err
	}
	res = dto.UserDTO{Status: result}
	return &res, nil
}

func (t DefaultUserService) UserGetAll() ([]models.User, error) {
	result, err := t.Repo.GetAll()

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultUserService) UserDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)
	if err != nil || result == false {
		return false, err

	}
	return true, nil

}
func (t DefaultUserService) UserUpdate(id primitive.ObjectID, update bson.M) (bool, error) {
	result, err := t.Repo.Update(id, update)
	if err != nil || result == false {
		return false, err
	}
	return true, nil
}

func NewUserService(Repo repository.UserRepository) DefaultUserService {

	return DefaultUserService{Repo: Repo}
} //open a connection from main to UserReposistoryDb
