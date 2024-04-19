package modelAccess

import (
	"gorm.io/gorm"
	"time"
)

type AccessEmployee struct {
	gorm.Model
	ID				uint		`json:"id"`
	EmployeeID		int			`json:"employee_id"`
	CreatedAt		time.Time	`json:"created_at"`
	Result			int			`json:"result"`
}
