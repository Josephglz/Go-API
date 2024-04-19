// @/main.go
package main

import (
    "log"

    "github.com/Josephglz/Go-API/config"
    "github.com/Josephglz/Go-API/controllers/employee"
    "github.com/Josephglz/Go-API/controllers/department"
    "github.com/Josephglz/Go-API/controllers/access"
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
    app.Post("/empleados/login", EmployeeController.ValidateEmployee)

    app.Get("/departamentos", DepartmentController.GetDepartment)
    app.Get("/departamentos/:id", DepartmentController.GetDepartment)
    app.Post("/departamentos", DepartmentController.CreateDepartment)
    app.Put("/departamentos/:id", DepartmentController.UpdateDepartment)
    app.Delete("/departamentos/:id", DepartmentController.DeleteDepartment)

    app.Get("/accesos", AccessController.GetAccesses)
    app.Get("/accesos/:id", AccessController.GetAccess)
    app.Post("/accesos", AccessController.CreateAccess)
    app.Put("/accesos/:id", AccessController.UpdateAccess)
    app.Delete("/accesos/:id", AccessController.DeleteAccess)


    log.Fatal(app.Listen(":3000"))
}