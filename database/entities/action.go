package entities

type Action struct {
	ID      string `gorm:"primaryKey"`
	Version string
	Action  string
	Hash    string // File hash
	Path    string // File path
}

const (
	ActionUpdate = "update"
	ActionDelete = "delete"
)
