package models

type Doctor struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Specialty string
}
