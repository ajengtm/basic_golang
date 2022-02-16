package public

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	gormio "gorm.io/gorm"

	"basic_golang/config"

	zaplogger "basic_golang/internal/adapter/zap"
	"basic_golang/internal/domain/auth"

	"github.com/go-chi/chi"

	"go.uber.org/zap"
)

/* File ini untuk dependency injection */
type Server struct {
	Cfg        config.MainConfig
	AuthDomain auth.AuthDomainInterface
	router     *chi.Mux
}

func NewServer(
	cfg config.MainConfig,
	gormioMasterClient *gormio.DB,
	gormioSlaveClient *gormio.DB,
) *Server {
	authDomain := auth.NewAuthDomain()

	return &Server{
		Cfg:        cfg,
		AuthDomain: authDomain,
		router:     chi.NewRouter(),
	}
}

func (s *Server) Run() (err error) {
	logger := zaplogger.GetLogger()
	// config chi
	s.routes()

	srv := &http.Server{
		Addr:    s.Cfg.Server.PublicPort,
		Handler: s.router,
	}
	logger.Info(fmt.Sprintf("Server is running in port %s", s.Cfg.Server.PublicPort))
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal("cannot start server", zap.Error(err))
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	tenSecond := time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), tenSecond)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Fatal("server shutdown failed", zap.Error(err))
	}

	if err == http.ErrServerClosed {
		err = nil
	}

	logger.Info("Shutting Down")
	return nil
}
