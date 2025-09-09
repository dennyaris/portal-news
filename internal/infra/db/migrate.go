package db

import (
	"fmt"
	"log"
	"time"

	"github.com/dennyaris/portal-news/internal/entity"
	"gorm.io/gorm"
)

// SafeAutoMigrate: aman dipanggil berkali-kali, tambah tabel/kolom/index,
// hindari konflik nama index di Postgres, dan tidak ganggu data existing.
func SafeAutoMigrate(db *gorm.DB) error {
	// Hindari race saat banyak instance start (khusus Postgres)
	if db.Dialector.Name() == "postgres" {
		if err := db.Exec("SELECT pg_advisory_lock(?)", 987654321).Error; err != nil {
			return fmt.Errorf("advisory lock: %w", err)
		}
		defer db.Exec("SELECT pg_advisory_unlock(?)", 987654321)
	}

	// Tambah tabel/kolom baru; tidak drop/rename
	if err := db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Content{}); err != nil {
		return fmt.Errorf("auto-migrate: %w", err)
	}

	// Pastikan index bernama unik per tabel (hindari “already exists” Postgres)
	m := db.Migrator()
	ensure := func(model any, name string) error {
		if m.HasIndex(model, name) {
			return nil
		}
		return m.CreateIndex(model, name)
	}

	if err := ensure(&entity.User{}, "uniq_users_email"); err != nil {
		return err
	}
	if err := ensure(&entity.Category{}, "uniq_categories_slug"); err != nil {
		return err
	}
	for _, ix := range []string{"uniq_contents_slug", "idx_contents_status", "idx_contents_author", "idx_contents_category"} {
		if err := ensure(&entity.Content{}, ix); err != nil {
			return err
		}
	}

	log.Printf("migrate ok %s", time.Now().Format(time.RFC3339))
	return nil
}
