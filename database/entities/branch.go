package entities

type Branch struct {
	ID            uint `gorm:"primaryKey"`
	App           uint
	Name          string
	LatestVersion string `gorm:"column:latest_version, unique"`
	Description   string
}
