package api

import (
	"backup-service/internal/api/backup"
	"net/http"
)

// var (
// 	backupRegex = regexp.MustCompile(`^/backup`)
// 	restoreRegex = regexp.MustCompile(`^/restore`)
// 	authRegex    = regexp.MustCompile(`^/auth`)
// )

func RouteHandlers(w http.ResponseWriter, req *http.Request) {
	http.Handle("/backup/", http.StripPrefix("/backup", http.HandlerFunc(backup.BackupHandler)))
}
