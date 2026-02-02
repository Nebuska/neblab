package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Nebuska/neblab/account/api"
	"github.com/Nebuska/neblab/account/api/middlewares"
	"github.com/Nebuska/neblab/account/internal/auth"
	"github.com/Nebuska/neblab/shared/config"
	"github.com/Nebuska/neblab/shared/database/postgres"
	"github.com/Nebuska/neblab/shared/jwtAuth"
	"github.com/Nebuska/neblab/shared/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewRouter(l *logger.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.GinLogger(l))
	router.Use(middlewares.RequestLogger(l))
	router.Use(middlewares.ErrorHandler())

	return router
}

func GinStarter(lc fx.Lifecycle, r *gin.Engine, cfg *config.Config) {
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
}

func main() {
	app := fx.New(
		logger.Module,
		config.Module,
		postgres.Module,
		jwtAuth.Module,
		auth.Module,

		fx.Provide(NewRouter),
		api.Module,

		fx.Invoke(GinStarter),
	)

	app.Run()
}
