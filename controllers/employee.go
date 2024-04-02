package EmployeeController

import (
    "github.com/Josephglz/Go-API/config"
    "github.com/Josephglz/Go-API/models"
    "github.com/gofiber/fiber/v2"
)

func GetEmployees(c *fiber.Ctx) error {
	var employees []models.Employee

	config.Database.Find(&employees)
	return c.Status(200).JSON(employees)
}

func GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee

	var result = config.Database.Find(&employee, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Empleado no encontrado",
		})
	}

	return c.Status(200).JSON(employee)
}

func CreateEmployee(c *fiber.Ctx) error {
	employee := new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al procesar la solicitud",
		})
	}

	config.Database.Create(&employee)
	return c.Status(201).JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employee := new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al procesar la solicitud",
		})
	}

	config.Database.Where("id = ?", id).Updates(&employee)
	return c.Status(200).JSON(employee)	
}

func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee

	var result = config.Database.Delete(&employee, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Empleado no encontrado",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Empleado eliminado",
	})
}