package backup

type BackupInSchema struct {
	Dbs []string `json:"dbs" form:"dbs" binding:"required,min=1"`
}
