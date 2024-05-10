package note

import (
	"janstupica/StickyNotes/internal/app/note"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RegisterHTTPEndpoints(
	router *gin.RouterGroup,
	log *zerolog.Logger,
	useCase note.UseCase,
) {
	handler := New(log, useCase)

	routes := router.Group("/note")
	{
		routes.POST("/add", handler.Create)
		routes.PUT("/update", handler.Update)
		routes.DELETE("/delete/:id", handler.Delete)
		routes.GET("/get/:id", handler.Get)
		routes.GET("/get/all/:tripId", handler.GetAll)
	}
}
