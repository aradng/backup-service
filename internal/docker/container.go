package docker

import "backup-service/internal/model"

type Container struct {
	ID      string
	Names   []string
	Image   string
	State   string
	Status  string
	Project string
	Dir     string

	Volumes  []string
	DB       bool
	Type     *model.DatabaseType
	UserName *string
	Password *string
}
