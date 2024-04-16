# TODO List

## Описание

Инструменты:

- веб-сервер: `gin`
- валидация: `validator`
- документация: `swaggo`
- база данных: `mongo-driver`
- конфиг: `viper`

## Запуск

Создать конфиг файл

```shell
cp .env.example .env
```

Указать параметры в `.env`

Запустить в докере

```shell
make docker-build
```

## Структура приложения

```
├── Makefile
├── README.md [<- вы здесь]
├── api
│   └── swagger [документация]
│       ├── docs.go
│       ├── swagger-config.yaml [конфиг документации]
│       ├── swagger.json
│       └── swagger.yaml
├── cmd
│   └── server [точка входа]
│       └── main.go
├── deploy [конфигурация билда]
│   └── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal [внутренние сервисы]
│   ├── config [конфигурация проекта]
│   │   ├── app_config.go
│   │   ├── config.go
│   │   ├── db_config.go
│   │   └── swager_config.go
│   ├── db [база дынных]
│   │   ├── db.go
│   │   ├── models [коллекции mongodb]
│   │   │   └── task.go
│   │   └── repository [обращение к коллекциям]
│   │       └── task
│   │           └── repository.go
│   ├── dto [запросы ответы апи]
│   │   ├── base.go
│   │   └── todo.go
│   ├── logger [логгер]
│   │   └── logger.go
│   ├── middleware [мидлвейр]
│   │   └── middleware.go
│   ├── todo [сервис задач]
│   │   ├── handlers.go
│   │   └── service.go
│   ├── validate [валидатор]
│   │   └── validate.go
│   └── values [статические данные]
│       ├── dictionary.go
│       └── errors.go
└── requests.http [примеры запросов]
```