package app

import "backup-service/internal/docker"

type App struct {
	Containers []docker.Container
}

