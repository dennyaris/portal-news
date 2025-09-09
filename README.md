# News Portal — Clean Architecture (Go + Fiber + GORM)
This is the GORM-backed variant of the news portal skeleton

## .env
```
PORT=8080
APP_ENV=development

# MySQL DSN: user:pass@tcp(host:3306)/dbname?parseTime=true&loc=Local
DB_DRIVER=mysql
DB_DSN=root:password@tcp(127.0.0.1:3306)/news_portal?parseTime=true&loc=Local
```

## Run
```bash
make bootstrap
make run
```

The app auto-migrates the MySQL schema on start.
