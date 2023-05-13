package entities

type Version struct {
	ID     string `gorm:"primaryKey"` // 24-character long string
	Branch uint
	Name   string
}
