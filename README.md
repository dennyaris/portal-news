# News Portal — Clean Architecture (Go + Fiber + GORM)

A pragmatic Clean Architecture starter for a news portal. Entities:
- **User** — writers/admins
- **Category** — section taxonomy
- **Content** — news article

## .env
```
PORT=8080
APP_ENV=development

DB_DRIVER=mysql
DB_DSN=root:password@tcp(127.0.0.1:3306)/news_portal?parseTime=true&loc=Local
```

## Quickstart
```bash
make bootstrap
make run
```

The app auto-migrates the MySQL schema on start.
