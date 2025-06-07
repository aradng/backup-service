package docker

import (
	"backup-service/internal/model"
	"crypto/md5"
	"encoding/hex"
)

func (container Container) Hash() string {
	input := container.Project
	for _, volume := range container.Volumes {
		input += volume
	}
	if container.Type != nil {
		input += model.SecretSchemas[*container.Type].Name
	}
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
