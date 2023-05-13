package entities

type Branch struct {
	ID          uint `gorm:"primaryKey"`
	App         uint
	Name        string
	Description string
}
