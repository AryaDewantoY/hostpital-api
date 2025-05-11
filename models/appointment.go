package models

type Appointment struct {
	ID        uint `gorm:"primaryKey"`
	PatientID uint
	DoctorID  uint
	Date      string

	Patient Patient `gorm:"foreignKey:PatientID"`
	Doctor  Doctor  `gorm:"foreignKey:DoctorID"`
}
