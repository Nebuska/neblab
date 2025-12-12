package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Nebuska/neblab/shared/config"
	"github.com/Nebuska/neblab/shared/database/mysql"
	"github.com/Nebuska/neblab/shared/jwtAuth"
	"github.com/Nebuska/neblab/shared/logger"
	"github.com/Nebuska/neblab/tasker/api"
	"github.com/Nebuska/neblab/tasker/api/middlewares"
	"github.com/Nebuska/neblab/tasker/internal"

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
		mysql.Module,
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
