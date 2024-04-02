// @/main.go
package main

import (
    "log"

    "github.com/Josephglz/Go-API/config"
    "github.com/Josephglz/Go-API/controllers"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    config.Connect()

    app.Get("/empleados", EmployeeController.GetEmployees)
    app.Get("/empleados/:id", EmployeeController.GetEmployee)
    app.Post("/empleados", EmployeeController.CreateEmployee)
    app.Put("/empleados/:id", EmployeeController.UpdateEmployee)
    app.Delete("/empleados/:id", EmployeeController.DeleteEmployee)

    log.Fatal(app.Listen(":3000"))
}