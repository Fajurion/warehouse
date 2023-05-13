package entities

type App struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}
