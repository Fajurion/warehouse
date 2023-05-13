package entities

type TFA struct {
	Account     uint `gorm:"primaryKey"`
	Secret      string
	BackupCodes string
}
