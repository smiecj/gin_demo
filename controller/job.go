// package controller
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// job controller
type jobController struct{}

// model
// todo: move to model folder
type Job struct {
	Name string
	Need int
}

// GetJobs godoc
// @Summary      Get jobs
// @Description  get jobs
// @Tags         job
// @Accept       json
// @Produce      json
// @Success      200  {array}   controller.Job
// @Router       /job/list [get]
func (c *jobController) GetJobs(context *gin.Context) {
	context.JSON(http.StatusOK, []Job{
		{Name: "software engineer", Need: 1000},
		{Name: "project manager", Need: 100},
	})
}

func NewJobController() *jobController {
	return new(jobController)
}
