package app

import (
	"net/http/httptest"
	"testing"

	services "example.com/sarang-apis/mocks/service"
	"example.com/sarang-apis/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var td UserHandler

func TestUserHandler_GetAllUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := services.NewMockUserService(ctrl)

	td = UserHandler(mockService)
	router := fiber.New()

	router.Get("/api/users", td.GetAllUser)

	var FakeDataForHandler = []models.User{
		{primitive.NewObjectID(), "Title1", "Content1"},
		{primitive.NewObjectID(), "Title2", "Content2"},
		{primitive.NewObjectID(), "Title3", "Content3"},
	}

	mockService.EXPECT().TodoGetAll().Return(FakeDataForHandler, nil)
	req := httptest.NewRequest("GET", "/api/users", nil)

	resp, _ := router.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}
