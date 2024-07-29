package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zhikariz/depublic/config"
	"github.com/zhikariz/depublic/internal/http/router"
)

type Server struct {
	*echo.Echo
	cfg *config.Config
}

func NewServer(cfg *config.Config, publicRoutes []*router.Route) *Server {
	e := echo.New()
	for _, v := range publicRoutes {
		e.Add(v.Method, v.Path, v.Handler)
	}
	return &Server{e, cfg}
}

func (s *Server) Run() {
	go func() {
		err := s.Start(fmt.Sprintf(":%s", s.cfg.Port))
		log.Fatal(err)
	}()
}

func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := s.Shutdown(ctx); err != nil {
			s.Logger.Fatal(err)
		}
	}()
}
