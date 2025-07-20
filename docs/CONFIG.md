## Файл конфигурации

```yaml
http:
  host: 0.0.0.0
  port: 8080

postgres:
  host: postgres
  port: 5432
  db: marketplace_db
  user: admin
  password: adminpass

jwt:
  issuer: marketplace.AuthService
  expiration_time: 30m

crypt:
  hashing_cost: 14

service:
  page_size: 5
  max_price: 10000000
  max_title_length: 100
  max_description_length: 2000
  max_image_size: 1024

  min_login_length: 3
  max_login_length: 30
  min_password_length: 8
  max_password_length: 30
  special_symbols: "!@#$%^&*?/"

  debug_mode: true

paths:
  api: /api/v1
  register: /auth/register
  login: /auth/login
  create_ad: /ads/create
  get_feed: /ads/feed
```

| Параметр | Значение |
| ----- | ----- |
| ```http``` | Хост и порт, которые будет прослушивать приложение. |
| ```postgres``` | Данные для подключения к PostgreSQL. |
| ```jwt.issuer``` | Имя издателя JWT (поле ```iss``` в теле).
| ```jwt.expiration_time``` | Время жизни JWT. По истечению этого времени токен становится невалидным. |
| ```crypt.hashing_cost``` | Цена хэширования пароля. Время хэширования увеличивается с ростом цены. |
| ```service``` | Настройки приложения. |
| ```service.debug_mode``` | Указывает, должен ли быть включен debug-режим. В частности, при включенном Debug Mode в теле HTTP-ответа возвращается полная структура ошибки, произошедшей на сервере в ходе неудачного выполнения запроса. |
| ```paths``` | Эндпойнты. |

## Переменные окружения

Следующие параметры нельзя задать через файл конфигурации (в целях безопасности).

| Переменная | Значение |
| --- | --- |
| PUBLIC_KEY_PEM | Публичный ключ ed25519, закодированный в PKIX. |
| PRIVATE_KEY_PEM | Приватный ключ ed25519, закодированный в PKCS #8. |
