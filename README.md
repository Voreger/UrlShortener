# URL Shortener

Сервис для сокращения URL с поддержкой двух типов хранилищ: **in-memory** и **PostgreSQL**.

## Структура проекта

```
UrlShortener/
├── cmd/
│   └── app/
│       └── main.go              # Точка входа
├── internal/
│   ├── handlers/                # HTTP обработчики
│   │   ├── url_handler.go
│   │   └── validator.go
│   ├── models/                  # Модели данных
│   │   └── url.go
│   ├── repository/              # Слой доступа к данным
│   │   ├── repository.go        # Интерфейс
│   │   ├── errors.go
│   │   ├── memory/              # In-memory реализация
│   │   └── postgres/            # PostgreSQL реализация
│   ├── routers/                 # Chi router
│   │   └── router.go
│   └── services/                # Бизнес-логика
│       ├── url_service.go
│       └── generator.go
├── migrations/                  # Goose миграции
│   └── 20260221122843_create_urls_table.sql
├── .env                         # Конфигурация
├── .env.example
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── README.md
```

## Быстрый старт

### 1. Через Docker (рекомендуется)

```bash
# Клонировать репозиторий
git clone https://github.com/YOUR_USERNAME/url-shortener.git
cd url-shortener

# Запустить сервис
docker-compose up -d
```


## Конфигурация

Переменные окружения (файл `.env`):

STORAGE - тип хранилища memory/postgres

API_PORT - Порт api сервера

POSTGRES_DB - имя базы данных

POSTGRES_USER - пользователь postgresql

POSTGRES_PASSWORD - пароль для postgresql

POSTGRES_HOST - хост для базы данных

POSTGRES_PORT - порт для базы данных

POSTGRES_DB_STRING - автоформируемая строка подключения к базе данных

## API

### POST /shorten

Создать короткую ссылку.

**Request:**
```http
POST /shorten
Content-Type: application/json

{
  "url": "https://google.com"
}
```

**Response (201 Created):**
```json
{
  "short": "MFc3jNUIu_"
}
```

**Response (400 Bad Request):**
```
URL is required
```

**Response (500 Internal Server Error):**
```
Failed to create short URL
```

---

### GET /{code}

Получить оригинальный URL по короткому коду.

**Request:**
```http
GET /MFc3jNUIu_
```

**Response (200 OK):**
```json
{
  "url": "https://google.com"
}
```

**Response (400 Bad Request):**
```
Short code is required
```

**Response (404 Not Found):**
```
URL not found
```

---

### GET /health

Health check endpoint.

**Request:**
```http
GET /health
```

**Response (200 OK):**
```
Service health is OK
```


## Требования к коротким ссылкам

| Требование | Реализация |
|------------|------------|
| Длина | 10 символов |
| Символы | `a-zA-Z0-9_` (63 символа) |
| Уникальность | SHA224 хэш + проверка коллизий |
| Идемпотентность | Один URL = один код |

#
## Тесты

```bash
# Запустить все тесты
go test ./... -v

# Запустить с покрытием
go test ./... -cover

# Запустить конкретный пакет
go test ./internal/services/... -v
go test ./internal/handlers/... -v
go test ./internal/repository/memory/... -v
```

## Graceful Shutdown

Сервис корректно обрабатывает сигналы завершения (SIGINT, SIGTERM):

1. Перестает принимать новые соединения
2. Ждет завершения активных запросов (до 10 секунд)
3. Закрывает подключения к БД
4. Останавливает сервер
