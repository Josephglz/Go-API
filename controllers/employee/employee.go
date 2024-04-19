package EmployeeController

import (
    "github.com/Josephglz/Go-API/config"
    "github.com/Josephglz/Go-API/models/employee"
    "github.com/gofiber/fiber/v2"
)

func GetEmployees(c *fiber.Ctx) error {
	var employees []modelEmployee.Employee

	config.Database.Find(&employees)
	return c.Status(200).JSON(employees)
}

func GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee modelEmployee.Employee

	var result = config.Database.Find(&employee, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Empleado no encontrado",
		})
	}

	return c.Status(200).JSON(employee)
}

func CreateEmployee(c *fiber.Ctx) error {
	employee := new(modelEmployee.Employee)

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
	employee := new(modelEmployee.Employee)

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
	var employee modelEmployee.Employee

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

func ValidateEmployee(c *fiber.Ctx) error {
	//Esta es un metodo que se llamará desde un método POST que recibirá en el body un email y un password
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al procesar la solicitud",
		})
	}

	email := body["email"]
	password := body["password"]

	var employee modelEmployee.Employee

	var result = config.Database.Where("email = ? AND password = ?", email, password).First(&employee)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Empleado no encontrado",
			"found": 0, 
		})
	}

	return c.Status(200).JSON(employee)
}