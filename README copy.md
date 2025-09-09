# News Portal — Clean Architecture (Go + Fiber)

A pragmatic Clean Architecture starter for a news portal. Entities:
- **User** — writers/admins
- **Category** — section taxonomy
- **Content** — news article

Memory repositories are provided to keep bootstrapping simple. Replace them with GORM/SQL adapters later without touching use cases/handlers.

## Quickstart
```bash
make bootstrap
make run
# PORT=9898 make run
```

## REST Endpoints
- Health: `GET /healthz`
- Users:
  - `POST /api/v1/users` (name, email)
  - `GET  /api/v1/users/:id`
  - `GET  /api/v1/users` (list)
- Categories:
  - `POST /api/v1/categories` (name, slug)
  - `GET  /api/v1/categories/:id`
  - `GET  /api/v1/categories?limit=&page=` (list)
  - `PUT  /api/v1/categories/:id`
  - `DELETE /api/v1/categories/:id`
- Contents:
  - `POST /api/v1/contents` (title, slug, body, status, author_id, category_id)
  - `GET  /api/v1/contents/:id`
  - `GET  /api/v1/contents?status=&cat=&author=&q=&limit=&page=` (list + filters)
  - `PUT  /api/v1/contents/:id`
  - `DELETE /api/v1/contents/:id`

Status for content: `draft|published`.
