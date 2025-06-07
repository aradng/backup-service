package backup

import (
	"backup-service/internal/docker"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.RouterGroup

func RegisterRoutes(router *gin.RouterGroup) {
	Router = router.Group("/backup")
	Router.POST("/", BackupDBs)
}

// @Summary backup dbs
// @Schemes
// @Param request body []string list "asd"
// @Tags Backup
// @Accept json
// @Produce json
// @Success 204
// @Failure 404 {array} string
// @Router /backup/ [post]
func BackupDBs(c *gin.Context) {
	var dbs []string
	status := make(map[string]bool)
	failed := make([]string, 0)
	err := c.BindJSON(&dbs)
	if err != nil {
		return
	}
	for _, db := range dbs {
		status[db] = false
	}
	for _, container := range docker.Containers {
		if container.DB && !status[container.ID] {
			status[container.ID] = true
			container.Backup()
		}
	}
	for _, db := range dbs {
		if !status[db] {
			failed = append(failed, db)
		}
	}
	if len(failed) == 0 {
		c.Status(http.StatusNoContent)
		return
	}
	c.JSON(http.StatusNotFound, failed)
}
