package server

import (
	"context"
	"janstupica/StickyNotes/docs"
	"janstupica/StickyNotes/internal/app/database"
	note "janstupica/StickyNotes/internal/app/note"
	notehttp "janstupica/StickyNotes/internal/app/note/delivery/http"
	notestorage "janstupica/StickyNotes/internal/app/note/storage/database"
	noteusecase "janstupica/StickyNotes/internal/app/note/usecase"
	"janstupica/StickyNotes/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Log        *zerolog.Logger
	HttpServer *http.Server
	DB         *database.DB

	NoteUseCase note.UseCase
}

func New() *Server {
	log := logger.Init("debug")
	database := database.New()

	noteStorage := notestorage.New(database.Client)

	return &Server{
		Log: &log,
		DB:  database,

		NoteUseCase: noteusecase.New(noteStorage),
	}
}

// @BasePath		/api/v1

// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "Hello World!")
}

// @BasePath		/api/v1
// @Title StickyNote API
// @Version 1.0
// @Description This is the API for StickyNotes App.
func (s *Server) Run() error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	docs.SwaggerInfo.BasePath = "/api/v1"

	api := router.Group("/api/v1")
	{
		eg := api.Group("/example")
		{
			eg.GET("/helloworld", HelloWorld)
		}
	}

	notehttp.RegisterHTTPEndpoints(api, s.Log, s.NoteUseCase)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := ":8080" // todo: change
	s.HttpServer = &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil {
			s.Log.Fatal().Msgf("Failed to listen and serve: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.HttpServer.Shutdown(ctx)
}
