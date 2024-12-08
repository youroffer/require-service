## Сервис Требований
Обеспечение актуальной аналитики по ключевым словам и навыкам, включая подсчет их частоты встречаемости в вакансиях через API hh.ru.

http://localhost:8081/swagger/index.html

## Запуск
1. Установка ogen и migrate в bin/tools
```bash
make setup
```
2. Важно добавить файл `public.key` в папку ./config/jwt
2. Настройка .env
```bash
cp example.env .env
```
3. Запуск сервисов
```bash
make compose.up
```

## Настройка бд
1. Создание миграции
```bash
make migrate.create name=имя_миграции
```
2. Применение миграций
```bash
make migrate.up
```

## Redis ui

http://localhost:5540

```bash
docker run -d \
  --name redisinsight \
  --network media_default \
  -p 5540:5540 \
  redis/redisinsight:latest
```