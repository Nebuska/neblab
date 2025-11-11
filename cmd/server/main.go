package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"task-tracker/api"
	"task-tracker/api/middlewares"
	"task-tracker/config"
	"task-tracker/internal"
	"task-tracker/pkg/database"
	"task-tracker/pkg/jwtAuth"
	"task-tracker/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewRouter(logger *logger.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.GinLogger(logger))
	router.Use(middlewares.RequestLogger(logger))
	router.Use(middlewares.ErrorHandler())

	return router
}

func main() {

	app := fx.New(
		logger.Module,
		fx.Provide(NewRouter),
		config.Module,
		database.Module,
		jwtAuth.Module,
		internal.Module,
		api.Module,
		fx.Invoke(func(lc fx.Lifecycle, r *gin.Engine, cfg *config.Config) {
			server := &http.Server{
				Addr:    ":" + cfg.ServerPort,
				Handler: r,
			}

			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					log.Println("Server running on port:", cfg.ServerPort)
					go func() {
						if err := server.ListenAndServe(); err != nil &&
							!errors.Is(err, http.ErrServerClosed) {
							log.Printf("server error: %v", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					log.Println("Server stopped")
					shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
					defer cancel()
					return server.Shutdown(shutdownCtx)
				},
			})
		}),
	)

	app.Run()
}
