package model

type Employee struct {
	ID       int64  `gorm:"primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Position string `gorm:"size:50"`
	Status   int    `gorm:"default:0"` // 0: Inactive, 1: Active
	At
}

type EmployeeListCond struct {
	Name     *string
	Position []string
	Status   *int
}

type EmployeeUpdateCond struct {
	ID       int64
	Name     *string
	Position *string
	Status   *int
}
