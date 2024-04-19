package modelDepartment

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name		string		`json:"name"`
	Manager		int			`json:"manager_id"`
}