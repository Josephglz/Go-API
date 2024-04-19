package AccessController

import (
	"github.com/Josephglz/Go-API/config"
	"github.com/Josephglz/Go-API/models/access"
	"github.com/gofiber/fiber/v2"
)

func GetAccesses(c *fiber.Ctx) error {
	var accesses []modelAccess.AccessEmployee

	config.Database.Find(&accesses)
	return c.Status(200).JSON(accesses)
}

func GetAccess(c *fiber.Ctx) error {
	id := c.Params("id")
	var access modelAccess.AccessEmployee

	var result = config.Database.Find(&access, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Acceso no encontrado",
		})
	}

	return c.Status(200).JSON(access)
}

func CreateAccess(c *fiber.Ctx) error {
	access := new(modelAccess.AccessEmployee)

	if err := c.BodyParser(access); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al procesar la solicitud",
		})
	}

	config.Database.Create(&access)
	return c.Status(201).JSON(access)
}

func UpdateAccess(c *fiber.Ctx) error {
	id := c.Params("id")
	access := new(modelAccess.AccessEmployee)

	if err := c.BodyParser(access); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al procesar la solicitud",
		})
	}

	config.Database.Where("id = ?", id).Updates(&access)
	return c.Status(200).JSON(access)
}

func DeleteAccess(c *fiber.Ctx) error {
	id := c.Params("id")
	var access modelAccess.AccessEmployee

	config.Database.Where("id = ?", id).Delete(&access)
	return c.Status(200).JSON(fiber.Map{
		"message": "Acceso eliminado",
	})
}