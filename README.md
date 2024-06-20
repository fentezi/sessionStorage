# Session Authentication Service

Этот проект представляет собой простой сервис аутентификации с использованием сессий. Пользователи могут регистрироваться, входить в систему и получать доступ к защищенным ресурсам.

## Установка

1. **Клонирование репозитория:**

   ```bash
   git clone https://github.com/fentezi/sessionStorage.git
   cd sessionStorage
   ```

2. **Настройка окружения:**

    Создайте файл .env в корневой директории проекта и добавьте в него следующую переменную для подключения к PostgreSQL:
    ```dotenv
    DB_SQL="host=localhost user=postgres password=pass dbname=dbname port=5432 sslmode=disable"
    ```

3. **Настройка конфигурации**

    Создайте файл config.yaml в корневой директории проекта и добавьте в него следующее содержимое для настройки логгирования и сервера: yaml
    ```yaml
    logging:
        level: "debug"
        format: "json"

    server:
        host: "0.0.0.0"
        port: "8080"
    ```

4. **Установка зависимостей и запуск**

    ```bash
    go mod tidy
    go run cmd/app/app.go
    ```
    Приложение будет доступно по адресу http://localhost:8080.


## Использование

# Регистрация пользователя

Отправьте POST запрос на /signup с JSON телом, содержащим email и password:

```json
{
    "email": "user@example.com",
    "password":"password123"
}
```

# Вход пользователя

Отправьте POST запрос на /signin с JSON телом, содержащим email и password:

```json
{
    "email":"user@example.com",
    "password":"password123"
}
```
После успешного входа вы получите cookie session_id, который используется для аутентификации пользователя.