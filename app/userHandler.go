package app

import (
	"net/http"

	"example.com/sarang-apis/models"
	"example.com/sarang-apis/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	Service services.UserService
}

func (h UserHandler) CreateUser(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	} //convert json to struct
	result, err := h.Service.UserInsert(user)

	if err != nil || result.Status == false {
		return err
	}
	return c.Status(http.StatusCreated).JSON(result)

}

func (h UserHandler) GetAllUser(c *fiber.Ctx) error {

	result, err := h.Service.UserGetAll()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())

	}
	return c.Status(http.StatusOK).JSON(result)

}

func (h UserHandler) DeleteUser(c *fiber.Ctx) error {
	query := c.Params("id")

	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.UserDelete(cnv)

	if err != nil || result == false {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}

func (h UserHandler) UpdateUser(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	var update bson.M
	if err := c.BodyParser(&update); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	result, err := h.Service.UserUpdate(cnv, update)
	if err != nil || result == false {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}
