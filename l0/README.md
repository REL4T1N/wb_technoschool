# WB TECHNOSCHOOL L0

## Структура проекта

```
l0
├─ database
│  └─ demo_orders
│     ├─ order_json
│     │  ├─ order1.json
│     │  ├─ order2.json
│     │  └─ order3.json
│     ├─ sql_inserts
│     │  ├─ insert_order1.sql
│     │  ├─ insert_order2.sql
│     │  └─ insert_order3.sql
│     ├─ docker_compose.yml
│     └─ init.sql
├─ app
│  ├─ internal
│  │  ├─ db
│  │  │  ├─ db.go
│  │  │  └─ repository.go
│  │  ├─ api/handler.go
│  │  ├─ cache/cache.go
│  │  ├─ config/config.go
│  │  ├─ kafka/consumer.go
|  │  └─ models/order.go
│  ├─ static/index.html
│  ├─ go.mod
│  ├─ go.sum
│  └─ main.go
```

## Локальный запуск сервиса

1. Перейти в корень проекта `l0`.
2. Перейти в папку `database` и запустить Docker Compose:
   ```bash
   cd database
   docker-compose up -d
   ```
3. Выйти из `database` и перейти в `app`:
   ```bash
   cd ../app
   go mod tidy
   go run main.go
   ```

## Техническое задание

### Развернуть локально базу данных
1) Создать новую базу данных для сервиса.
2) Настроить пользователя: заведите пользователя и выдайте права на созданную БД.
3) Создать таблицы: спроектируйте структуру для хранения полученных данных о заказах, ориентируясь на прилагаемую модель данных.

### Разработать сервис
1) Написать приложение на Go, реализующее описанные ниже функции.
2) Разработать простейший интерфейс для отображения полученных данных по ID заказа.
3) Подключиться и подписаться на канал сообщений: настроить получение данных из брокера сообщений (Kafka).
4) Сохранять полученные данные в БД: при приходе нового сообщения о заказе парсить его и вставлять соответствующую запись(и) в PostgreSQL.
5) Реализовать кэширование данных в сервисе: хранить последние полученные данные заказов в памяти (например, в map), чтобы быстро выдавать их по запросу.
6) При перезапуске восстанавливать кеш из БД: при старте сервиса заполнять кеш актуальными данными из базы, чтобы продолжить обслуживание запросов без задержек.
7) Запустить HTTP-сервер для выдачи данных по ID: реализовать HTTP-эндпоинт, который по `order_id` будет возвращать данные заказа из кеша (JSON API). Если в кеше данных нет, можно подтягивать из БД.

## Пример заказа

Файл `model.json`:

```json
{
   "order_uid": "b563feb7b2b84b6test",
   "track_number": "WBILMTESTTRACK",
   "entry": "WBIL",
   "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
   },
   "payment": {
      "transaction": "b563feb7b2b84b6test",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
   },
   "items": [
      {
         "chrt_id": 9934930,
         "track_number": "WBILMTESTTRACK",
         "price": 453,
         "rid": "ab4219087a764ae0btest",
         "name": "Mascaras",
         "sale": 30,
         "size": "0",
         "total_price": 317,
         "nm_id": 2389212,
         "brand": "Vivienne Sabo",
         "status": 202
      }
   ],
   "locale": "en",
   "internal_signature": "",
   "customer_id": "test",
   "delivery_service": "meest",
   "shardkey": "9",
   "sm_id": 99,
   "date_created": "2021-11-26T06:22:19Z",
   "oof_shard": "1"
}
```

Другие примеры заказов в формате JSON находятся по пути:
```
./database/demo_orders/order_json
```
SQL-запросы для их добавления в базу:
```
./database/demo_orders/sql_inserts
```

## Требования
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL 15
- Kafka 7.4.0
- Доступ к локальной сети: http://localhost:8080 (можно изменить в конфиге сервиса)
- Подключение к PostgreSQL указано в `app/internal/config/config.go`

## Использование

1. Убедиться, что Docker и Docker Compose установлены.
2. Перейти в папку `database` и запустить контейнеры:
   ```bash
   cd database
   docker-compose up -d
   ```
3. Перейти в `app`, установить зависимости и запустить сервис:
   ```bash
   cd ../app
   go mod tidy
   go run main.go
   ```
4. Проверить работу HTTP API по адресу: `http://localhost:8080/orders/{order_uid}`

