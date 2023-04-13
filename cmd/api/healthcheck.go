package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthcheckHandler godoc
//	@Summary	Retrieves current status of Backend app
//	@Produce	json
//	@Success	200
//	@Router		/v1/healthcheck [get]
func (app *application) healthcheckHandler(c *gin.Context) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	err := app.writeJSON(c.Writer, http.StatusOK, env, nil)
	if err != nil {
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(c.Writer, c.Request, err)
	}
}
