package entities

type Key struct {
	App    uint `gorm:"primaryKey"`
	Key    string
	Secret string
}
