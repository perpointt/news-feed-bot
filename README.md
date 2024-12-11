# База данных

## Запуск
`docker-compose -f docker-compose.dev.yaml up -d`

## Миграции
### Статус
`goose postgres "host=localhost user=postgres dbname=news_feed_bot password=postgres  sslmode=disable" status`

### Накатить

`goose postgres "host=localhost user=postgres dbname=news_feed_bot password=postgres  sslmode=disable" up`

### Подключиться к бд

`docker exec -it <container_name> psql -U postgres -d news_feed_bot`

### Получить миграции

`SELECT * FROM goose_db_version;`

### Удалить версию миграции

`DELETE FROM goose_db_version WHERE version_id = <version_name>;`
