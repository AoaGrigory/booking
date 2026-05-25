Поднять базу

```bash
docker-compose up -d
```

Запустить сервер

```bash
go run main/main.go
```

Эндпоинты

Получить все комнаты \
GET /rooms

Получить комнату по айди \
GET /rooms/:id

Создать комнату(только админ) \
POST /rooms \
Headers X-User-Role: admin \
тело запроса
```json
{
    "class": "Lux",
    "price": 5000,
    "description": "Номер с видом на море"
}
```

Удалить комнату(только админ) \
DELETE /rooms/:id \
Headers X-User-Role: admin

Забронировать комнату \
POST /bookings \
Headers X-User-ID: 1 
тело запроса
```json
{
    "room_id": 1,
    "start_date": "2026-06-01",
    "end_date": "2026-06-05"
}
```
Получить информацию о бронировании \
GET /bookings/:id 
