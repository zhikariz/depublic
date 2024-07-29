package main

import (
	"github.com/zhikariz/depublic/config"
	"github.com/zhikariz/depublic/internal/builder"
	"github.com/zhikariz/depublic/pkg/database"
	"github.com/zhikariz/depublic/pkg/server"
)

type User struct {
	ID   int64
	Name string
}

func (User) TableName() string {
	return "users"
}

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	db, err := database.ConnectToPostgres(cfg)
	checkError(err)

	publicRoutes := builder.BuildPublicRoutes(db)
	_ = builder.BuildPrivateRoutes()

	srv := server.NewServer(cfg, publicRoutes)
	srv.Run()
	srv.GracefulShutdown()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
