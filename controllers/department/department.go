package DepartmentController

import (
    "github.com/Josephglz/Go-API/config"
    "github.com/Josephglz/Go-API/models/department"
    "github.com/gofiber/fiber/v2"
)

func GetDepartments(c *fiber.Ctx) error {
    var departments []modelDepartment.Department
    config.Database.Find(&departments)
    
    return c.Status(200).JSON(departments)
}

func GetDepartment(c *fiber.Ctx) error {
    id := c.Params("id")
    var department modelDepartment.Department

    var result = config.Database.Find(&department, id)

    if result.RowsAffected == 0 {
        return c.Status(404).JSON(fiber.Map {
            "message": "Departamento no encontrado",
        })
    }

    return c.Status(200).JSON(department)
}

func CreateDepartment(c *fiber.Ctx) error {
    department := new(modelDepartment.Department)

    if err := c.BodyParser(department); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Error al procesar la solicitud",
        })
    }

    config.Database.Create(&department)
    return c.Status(201).JSON(&department)
}

func UpdateDepartment(c *fiber.Ctx) error {
    id := c.Params("id")
    department := new(modelDepartment.Department)

	if err := c.BodyParser(department); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al procesar la solicitud",
		})
	}

    config.Database.Where("id = ?", id).Updates(&department)
    return c.Status(200).JSON(department)
}

func DeleteDepartment(c *fiber.Ctx) error {
    id := c.Params("id")
    var department modelDepartment.Department

    var result = config.Database.Delete(&department, id)

    if result.RowsAffected == 0 {
        return c.Status(404).JSON(fiber.Map {
            "message": "Empleado no encontrado",
        })
    }

    return c.Status(200).JSON(fiber.Map {
        "message": "Empleado eliminado",
    })
}
