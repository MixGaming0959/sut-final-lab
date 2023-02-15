package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~ต้องไม่เป็นค่าว่าง"`
	Email      string
	EmployeeID string `valid:"matches(^[J || M || S]\\d{8}$)~รหัสพนักงานขึนต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว"`
	// รหัสพนักงานขึนต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว
}
