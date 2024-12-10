package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/indraexyt2/web-forum-go/internal/configs"
	"github.com/indraexyt2/web-forum-go/internal/handlers/memberships"
	"github.com/indraexyt2/web-forum-go/internal/handlers/posts"
	membershipRepo "github.com/indraexyt2/web-forum-go/internal/repository/memberships"
	postRepo "github.com/indraexyt2/web-forum-go/internal/repository/posts"
	membershipSvc "github.com/indraexyt2/web-forum-go/internal/service/memberships"
	postSvc "github.com/indraexyt2/web-forum-go/internal/service/posts"
	"github.com/indraexyt2/web-forum-go/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
