package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"task-tracker/api"
	"task-tracker/config"
	"task-tracker/internal"
	"task-tracker/pkg/database"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func main() {

	app := fx.New(
		fx.Provide(NewRouter),
		config.Module,
		database.Module,
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
