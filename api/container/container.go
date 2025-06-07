package container

import (
	"backup-service/internal/docker"
	"backup-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.RouterGroup

func RegisterRoutes(router *gin.RouterGroup) {
	Router = router.Group("/container")
	Router.POST("/", GetContainers)
}

// @Summary get containers
// @Param is_db query bool false "Filter to only show databases"
// @Tags Container
// @Accept json
// @Produce json
// @Success 200 {array} docker.Container
// @Router /backup/containers [get]
func GetContainers(c *gin.Context) {
	var dbInSchema DBInSchema
	if err := c.ShouldBind(&dbInSchema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if dbInSchema.IsDB != nil && *dbInSchema.IsDB {
		c.JSON(http.StatusOK, utils.Filter(docker.Containers, func(c *docker.Container) bool {
			return c.DB
		}))
		return
	}
	c.JSON(http.StatusOK, docker.Containers)
}
