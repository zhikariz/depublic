package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/zhikariz/depublic/common"
	"github.com/zhikariz/depublic/config"
	"github.com/zhikariz/depublic/internal/http/router"
)

type Server struct {
	*echo.Echo
	cfg *config.Config
}

func NewServer(cfg *config.Config, publicRoutes, privateRoutes []*router.Route) *Server {
	e := echo.New()
	for _, v := range publicRoutes {
		e.Add(v.Method, v.Path, v.Handler)
	}

	for _, v := range privateRoutes {
		e.Add(v.Method, v.Path, v.Handler, JWTMiddleware(cfg.JWTSecretKey))
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

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(common.JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
	})
}
