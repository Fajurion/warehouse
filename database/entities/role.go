package entities

type Role struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	PermissionLevel uint
}
