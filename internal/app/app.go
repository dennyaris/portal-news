package app

import (
	"log"
	"time"

	"github.com/dennyaris/portal-news/internal/config"
	"github.com/dennyaris/portal-news/internal/delivery/http/handler"
	"github.com/dennyaris/portal-news/internal/delivery/http/router"
	"github.com/dennyaris/portal-news/internal/infra/db"
	sqlrepo "github.com/dennyaris/portal-news/internal/repository/sql"
	"github.com/dennyaris/portal-news/internal/usecase/category"
	"github.com/dennyaris/portal-news/internal/usecase/content"
	"github.com/dennyaris/portal-news/internal/usecase/user"
	"github.com/dennyaris/portal-news/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
)

type Application struct{ Fiber *fiber.App }

func New(cfg *config.Config) *Application {
	app := fiber.New()

	gormDB, err := db.InitDB(db.Config{Driver: cfg.DBDriver, DSN: cfg.DSN})
	if err != nil {
		log.Fatalf("db open: %v", err)
	}

	// Auto-migrations using entity structs
	// if err := gormDB.AutoMigrate(
	// 	&entity.User{},
	// 	&entity.Category{},
	// 	&entity.Content{}); err != nil {
	// 	log.Fatalf("auto-migrate: %v", err)
	// }

	if err := db.SafeAutoMigrate(gormDB); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	// Repos
	uRepo := sqlrepo.NewUserRepo(gormDB)
	cRepo := sqlrepo.NewCategoryRepo(gormDB)
	nRepo := sqlrepo.NewContentRepo(gormDB)

	vd := validator.New()
	uid := func() string { return xid.New().String() }
	now := time.Now

	uSvc := user.New(uRepo, uid, now, vd.Validate)
	catSvc := category.New(cRepo, uid, now, vd.Validate)
	newsSvc := content.New(nRepo, uid, now, vd.Validate)

	uh := handler.NewUserHandler(uSvc)
	ch := handler.NewCategoryHandler(catSvc)
	nh := handler.NewContentHandler(newsSvc)

	router.Register(app, uh, ch, nh)
	return &Application{Fiber: app}
}
