# Blacklist Test Task
### Локальный запуск:
___
#### Параметры среды в файле `.env` в корневой директории
```shell
JWT_SECRET=
ADMIN_USERNAME=
ADMIN_PASSWORD=

PG_USERNAME=
PG_PASSWORD=
DATABASE_NAME=
DRIVER_NAME=
DATABASE_HOST_PORT=

SERVER_HOST=
SERVER_PORT=
```
#### Запуск БД и сервиса через docker compose
```shell
docker-compose up
```

### Описание работы:
___
Микросервис предназначен для хранения информации о пользователях, которые были добавлены в черный список. Микросервис имеет возможность добавления, удаления и поиска пользователей в черном списке. Данные хранятся в СУБД Postgres.

Стандартный адрес работы http-сервера `localhost:8283`

Методы API доступны в Swagger `localhost:8283/swagger/index.html`

Для выполнения методов необходима авторизация через Bearer токен, который можно получить после Basic Auth реквеста на адрес `/api/v1/token`. Пара `username:password` по умолчанию `admin:admin`. 

---
## API запросы к сервису:
```
// POST /api/v1/blacklist/
// DELETE /api/v1/blacklist/{id}
// GET /api/v1/blacklist?user_name={user_name}
// GET /api/v1/blacklist?phone_number={phone_number}
```
---

### Получение JWT токена

`GET /api/v1/auth/token`

Требует заголовок `Authorization: Basic <credentials>`.

Возвращает JSON с Bearer токеном в случае успешной аутентификации.

    {
        "token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNlY3JldC1pZCIsInR5cCI6IkpXVCJ9.eyJFeHRlbnNpb25zIjpudWxsLCJHcm91cHMiOm51bGwsIklEIjoiMSIsIk5hbWUiOiJyb290IiwiYXVkIjpbIiJdLCJleHAiOjE2NzcwMTMyMDksImlhdCI6MTY3NzAxMjkwOSwibmJmIjoxNjc3MDEyOTA5LCJzdWIiOiIxIn0.BQ0nZcjEQ68IYXC2IIAVncGvd3x9HuGHcx3jlxRIfVg"
    }

## Создание
### Добавить запись в черный список
#### Request

`POST /api/v1/blacklist/`

Требует `Authorization: Bearer <token>` header.

Содержит raw body с json данными о внесенном в черный лист пользователе. JSON содержит следующие параметры:
- `phone_number` Номер телефона пользователя (обязательный параметр)
- `user_name` Имя пользователя (обязательный параметр)
- `ban_reason` Причина добавления пользователя в черный список (обязательный параметр)
- `username_who_banned` Имя пользователя, который добавил в черный список (обязательный параметр)
```
{
  "phone_number": "79990004422",
  "user_name": "Test Testov",
  "ban_reason": "Ban Reason",
  "username_who_banned": "SomeUser"
}
```

#### Response
Ответ содержит в себе json с созданным id. Возвращает статус 201:

    HTTP/1.1 201 Created
    Content-Length: 46 
    Content-Type: application/json
    Date: Tue,21 Feb 2023 21:09:02 GMT


    {"id": "6696a8a2-cb97-47f3-9ce3-95522c64f218"}


## Получение
### Получение по номеру телефона
#### Request
`GET /api/v1/blacklist?phone_number={phone_number}`

Требует `Authorization: Bearer <token>` header.

Получение записей из черного листа для данного номера телефона.

#### Response

Возвращает статус 200 при успешно найденной записи и json со списком полученных данных.

    HTTP/1.1 200 OK
    Date: Tue,21 Feb 2023 21:08:34 GMT
    Status: 200 OK
    Content-Type: application/json
    Content-Length: 220

    [{"id": "be2d9702-b63c-4e5a-a6e5-71da42b8e643","phone_number": "79992221133","user_name": "Testov Test Test1","ban_reason": "Some ban reason","date_banned": "2023-01-11T13:24:49Z","username_who_banned": "Somebody Whobanned"}]

Возвращает статус 404 если запись не найдена и текст ошибки.

        content-length: 17
        content-type: text/plain; charset=utf-8
        date: Tue,20 Feb 2023 14:55:55 GMT
        x-content-type-options: nosniff

        record not found

### Получение по имени пользователя
#### Request
`GET /api/v1/blacklist?user_name={user_name}`

Требует `Authorization: Bearer <token>` header.

Получение записей из черного листа для данного  имени пользователя.

#### Response

Возвращает статус 200 при успешно найденной записи и json со списком полученных данных.

    HTTP/1.1 200 OK
    Date: Tue,21 Feb 2023 21:08:34 GMT
    Status: 200 OK
    Content-Type: application/json
    Content-Length: 220

    [{"id": "be2d9702-b63c-4e5a-a6e5-71da42b8e643","phone_number": "79992221133","user_name": "Testov Test Test1","ban_reason": "Some ban reason","date_banned": "2023-01-11T13:24:49Z","username_who_banned": "Somebody Whobanned"}]

Возвращает статус 404 если запись не найдена и текст ошибки.

        content-length: 17
        content-type: text/plain; charset=utf-8
        date: Tue,20 Feb 2023 14:55:55 GMT
        x-content-type-options: nosniff

        record not found



## Удаление
### Удаление по id
#### Request
`DELETE /api/v1/blacklist/{id}`

Требует `Authorization: Bearer <token>` header.

Удаление по id записи с пользователем.

#### Response

Возвращает статус 204 при успешно удаленной записи.

    HTTP/1.1 204 No Content
    Date: Fri, 13 Jan 2023 20:45:15 GMT

Возвращает статус 404 если запись не найдена и текст ошибки

    content-length: 17
    content-type: text/plain; charset=utf-8
    date: Tue,21 Feb 2023 21:10:04 GMT
    x-content-type-options: nosniff

    record not found