package config

import (
    "github.com/Josephglz/Go-API/models/employee"
    "github.com/Josephglz/Go-API/models/access"
    "github.com/Josephglz/Go-API/models/department"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:6KiacKB9O0zF@tcp(localhost:3306)/mybusinessgo?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
    var err error

    Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })

    if err != nil {
        panic(err)
    }

    Database.AutoMigrate(&modelEmployee.Employee{})
    Database.AutoMigrate(&modelAccess.AccessEmployee{})
    Database.AutoMigrate(&modelDepartment.Department{})

    return nil
}