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

	gormDB, err := db.Open(db.Config{Driver: cfg.DBDriver, DSN: cfg.DSN})
	if err != nil {
		log.Fatalf("db open: %v", err)
	}

	// Auto-migrations using entity structs (simple mapping)
	if err := gormDB.AutoMigrate(&userModel{}, &categoryModel{}, &contentModel{}); err != nil {
		log.Fatalf("auto-migrate: %v", err)
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

// Local GORM model mapping mirroring entity fields
// Keeping these here avoids coupling entity package to gorm tags if you prefer that separation.
type userModel struct {
	ID        string `gorm:"primaryKey;size:32"`
	Name      string
	Email     string `gorm:"uniqueIndex;size:191"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type categoryModel struct {
	ID        string `gorm:"primaryKey;size:32"`
	Name      string
	Slug      string `gorm:"uniqueIndex;size:191"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type contentModel struct {
	ID          string `gorm:"primaryKey;size:32"`
	Title       string
	Slug        string `gorm:"uniqueIndex;size:191"`
	Body        string `gorm:"type:longtext"`
	Excerpt     string `gorm:"type:longtext"`
	Image       string `gorm:"type:longtext"`
	Status      string `gorm:"index;size:32"`
	AuthorID    string `gorm:"index;size:32"`
	CategoryID  string `gorm:"index;size:32"`
	Tags        string `gorm:"type:longtext"`
	PublishedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
