# Oly Server - Go Version

Сервер мессенджера Oly, переписанный на Go.

## Структура проекта

```
go-server/
├── cmd/
│   └── server/
│       └── main.go          # Точка входа приложения
├── internal/
│   ├── config/
│   │   └── config.go        # Конфигурация приложения
│   ├── database/
│   │   └── database.go      # Подключение к базе данных
│   ├── handlers/
│   │   ├── auth.go          # Обработчики аутентификации
│   │   └── chat.go          # Обработчики чатов
│   ├── middleware/
│   │   └── jwt.go           # JWT middleware
│   ├── models/
│   │   └── models.go        # Модели данных
│   └── services/
│       ├── auth.go          # Сервис аутентификации
│       └── chat.go          # Сервис чатов
├── go.mod
├── go.sum
└── README.md
```

## Требования

- Go 1.19+
- PostgreSQL
- Переменные окружения (см. ниже)

## Установка зависимостей

```bash
cd go-server
go mod tidy
```

## Сборка

```bash
go build -o oly-server ./cmd/server
```

## Запуск

```bash
./oly-server
```

Сервер запустится на порту 3000.

## Переменные окружения

Создайте файл `.env` или установите переменные окружения:

```env
DB_USER=lyola
DB_PASSWORD=13lyola
DB_NAME=lyoladb
DB_PORT=5432
DB_HOST=localhost
JWT_SECRET=L7pxkiaY/eZHEXHcvgHcZx93G8hqwvNDJamBIKyn1E4=
```

## API Endpoints

### Аутентификация

#### POST /auth/sign-up
Регистрация нового пользователя.

**Request Body:**
```json
{
  "username": "string",
  "email": "string",
  "password": "string"
}
```

#### POST /auth/sign-in
Вход в систему.

**Request Body:**
```json
{
  "email": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Вы успешно вошли в систему",
  "token": "jwt_token",
  "user": {
    "id": "string",
    "username": "string",
    "email": "string"
  }
}
```

### Чаты

#### GET /chat
Получить список чатов пользователя.

**Headers:**
- `X-User-Id: string` - ID текущего пользователя

#### POST /chat/create
Создать новый чат.

**Request Body:**
```json
{
  "creatorId": "string"
}
```

#### POST /chat/:chatId/participants
Добавить участника в чат.

**Request Body:**
```json
{
  "userId": "string"
}
```

## Схема базы данных

Проект использует ту же схему базы данных, что и оригинальный TypeScript проект:

- `user` - пользователи
- `chat` - чаты
- `chat_participant` - участники чатов
- `message` - сообщения
- `voice_call` - голосовые звонки
- `voice_participant` - участники голосовых звонков

## Отличия от TypeScript версии

1. **Фреймворк**: Вместо Elysia.js используется стандартная библиотека `net/http`
2. **База данных**: Вместо Drizzle ORM используется стандартный `database/sql` с драйвером `lib/pq`
3. **Хеширование паролей**: Вместо Bun.password используется `golang.org/x/crypto/bcrypt`
4. **JWT**: Вместо @elysia/jwt используется `github.com/golang-jwt/jwt/v5`
5. **Генерация ID**: Вместо @paralleldrive/cuid2 используется простая функция на основе `crypto/rand`
