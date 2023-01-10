package services

import (
	"testing"

	"example.com/sarang-apis/mocks/repository"
	"example.com/sarang-apis/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockRepo *repository.MockUserRepository
var service UserService

var FakeData = []models.User{
	{primitive.NewObjectID(), "Title1", "Content1"},
	{primitive.NewObjectID(), "Title2", "Content2"},
	{primitive.NewObjectID(), "Title3", "Content3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repository.NewMockUserRepository(ct)
	service = NewUserService(mockRepo)

	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultUserService_UserGetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil)
	result, err := service.UserGetAll()
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, result)
}
