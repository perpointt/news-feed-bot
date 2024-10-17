# База данных

## Запуск
`docker-compose -f docker-compose.dev.yaml up -d`

## Миграции
### Статус
`goose postgres "host=localhost user=postgres dbname=news_feed_bot password=postgres  sslmode=disable" status`

### Накатить
`goose postgres "host=localhost user=postgres dbname=news_feed_bot password=postgres  sslmode=disable" status`